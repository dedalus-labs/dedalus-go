// File generated from our OpenAPI spec by Scalar. See README.md for details.

// Package redact masks the credentials this SDK sends wherever the SDK prints
// a request: the debug log, an API error's message and dumps, a transport
// error's URL, and the headers a redirect may not carry to another origin.
package redact

import (
	"encoding/hex"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// SensitiveHeaders are the headers a credential travels in, lower-cased: the
// conventional credential headers, plus every header this SDK's auth schemes
// send a credential in. Their values are masked before a request or response
// is printed, and they are dropped from a redirect to another origin.
var SensitiveHeaders = []string{"authorization", "proxy-authorization", "api-key", "x-api-key", "cookie", "set-cookie", "api_key"}

// SensitiveQueryParams are the query parameters this SDK's auth schemes send a
// credential in. Their values are masked wherever printed text quotes them: the
// request line, a header such as a `Location` or `Link` that echoes the URL,
// or the body.
var SensitiveQueryParams = []string{}

// Mask is what a redacted value is replaced with.
const Mask = "***"

// IsSensitiveHeader reports whether name is one of [SensitiveHeaders], in any
// case.
func IsSensitiveHeader(name string) bool {
	for _, sensitive := range SensitiveHeaders {
		if strings.EqualFold(name, sensitive) {
			return true
		}
	}
	return false
}

// Headers returns headers with the value of every [SensitiveHeaders] entry
// replaced, cloning only when one is present, so the headers passed in are
// never changed.
func Headers(headers http.Header) http.Header {
	var redacted http.Header
	for _, name := range SensitiveHeaders {
		values := headers.Values(name)
		if len(values) == 0 {
			continue
		}
		if redacted == nil {
			redacted = headers.Clone()
		}
		redacted.Del(name)
		for range values {
			redacted.Add(name, Mask)
		}
	}
	if redacted == nil {
		return headers
	}
	return redacted
}

// userinfoPattern matches the userinfo of a URL quoted in text, where a token
// often rides (`https://user:token@host`), keeping the `//` before it.
// The `//` may also be JSON-escaped as `\/\/`, which an encoder that
// escapes slashes writes for every URL in a body; the matched spelling is kept.
// Besides what ends an authority, a quote, `,`, `<` or `>` ends the run, since
// a body or header quotes the URL between them: without that stop the match
// would run on from one JSON string into the next and swallow everything up
// to an `@` there, as in `"https://example.com","email":"jane@example.com"`.
var userinfoPattern = sync.OnceValue(func() *regexp.Regexp {
	return regexp.MustCompile(`(//|\\/\\/)[^/?#@\s"',<>]+@`)
})

// queryPattern is the pattern [Text] masks credential query values with, built
// once on first use; nil when this SDK places no credential in the query.
var queryPattern = sync.OnceValue(func() *regexp.Regexp {
	names := make([]string, 0, len(SensitiveQueryParams))
	for _, name := range SensitiveQueryParams {
		// An empty name would build a pattern that masks every `?=`.
		if name != "" {
			names = append(names, namePattern(name))
		}
	}
	if len(names) == 0 {
		return nil
	}
	// A name counts after a separator, `?`, `&` or `;`, either literal or
	// percent-encoded any number of times over (`%3F`, `%253F`), since a
	// server can echo a whole URL encoded inside another parameter's value, as
	// a redirect's `return_to` does. A JSON-escaped `&` (`\u0026`) counts too:
	// Go's `encoding/json` writes every `&` in a string that way, so a URL in a
	// JSON body reads `?page=2\u0026api_key=…`. Its hex digits hold no letter,
	// so there is no case to fold. The `=` after the name can be encoded the
	// same way as the separator. The value then runs to the next literal `&`, `#` or whitespace (a
	// dump's line ends in `\r\n`): an encoded `%26` may end a nested value, but
	// stopping there would leak the rest of a value a sloppy encoder left
	// half-encoded, and masking too much is harmless. `<`, `>` and `"` end it
	// too, since no valid URL carries them unencoded, so a `Link` header's
	// `<url>` or a quoted URL keeps its closing delimiter. `'` does not: a
	// query value can hold it raw. RE2 matches in linear time, so a hostile
	// response cannot stall the log.
	return regexp.MustCompile(`((?:[?&;]|%(?:25)*(?i:3f|26|3b)|\\u0026)(?:` + strings.Join(names, "|") + `)(?:=|%(?:25)*(?i:3d)))[^&#\s<>"]*`)
})

// namePattern matches one credential query name in every spelling a URL can
// carry it in: each character as written or percent-encoded, at any depth and
// in either hex case (`k`, `%6B`, `%6b`, `%256B`), and a space also as `+`.
// Only the hex ignores case, not the name. A multi-byte character is matched
// whole or as its encoded UTF-8 bytes, since Go's regexp matches runes.
func namePattern(name string) string {
	var pattern strings.Builder
	for _, r := range name {
		var encoded strings.Builder
		for _, b := range []byte(string(r)) {
			encoded.WriteString(`%(?:25)*(?i:` + hex.EncodeToString([]byte{b}) + `)`)
		}
		pattern.WriteString(`(?:` + regexp.QuoteMeta(string(r)) + `|` + encoded.String())
		if r == ' ' {
			pattern.WriteString(`|\+`)
		}
		pattern.WriteString(`)`)
	}
	return pattern.String()
}

// Text masks URL userinfo and every credential query value in text, wherever
// they appear in it: the request line, the values of other query parameters
// that nest an encoded URL, header values, and the body.
func Text(text []byte) []byte {
	text = userinfoPattern().ReplaceAll(text, []byte("${1}"+Mask+"@"))
	if pattern := queryPattern(); pattern != nil {
		text = pattern.ReplaceAll(text, []byte("${1}"+Mask))
	}
	return text
}

// String is [Text] for a string.
func String(text string) string {
	return string(Text([]byte(text)))
}

// URL is the text of u with its userinfo and credential query values masked.
func URL(u *url.URL) string {
	if u == nil {
		return ""
	}
	return String(u.String())
}

// URLError returns err with the URL a [*url.Error] names masked, keeping its
// operation and cause, so [errors.Is] and [errors.As] still reach the cause. The
// standard client names the full request URL, query credentials and all, in
// every transport error it returns. Any other error is returned as it is.
func URLError(err error) error {
	var urlErr *url.Error
	if err == nil || !errors.As(err, &urlErr) || urlErr != err {
		return err
	}
	return &url.Error{Op: urlErr.Op, URL: String(urlErr.URL), Err: urlErr.Err}
}

// maxMessageBody is how much of a response body an error message quotes.
// The whole body stays readable on the error itself.
const maxMessageBody = 2048

// MessageBody is body made fit for a one-line error message: credential
// values are masked, the text is cut after maxMessageBody bytes, and control
// characters (line breaks, terminal escapes, bidirectional overrides) are
// written as Go escapes, so a hostile body cannot forge log lines or restyle a
// terminal. Masking comes first, on the whole body: a credential the cut would
// split, such as userinfo whose closing `@` lies past it, is then still
// recognised.
func MessageBody(body string) string {
	body = String(body)
	truncated := false
	if len(body) > maxMessageBody {
		cut := maxMessageBody
		// Step back to a character boundary rather than end on half of one.
		for cut > 0 && !utf8.RuneStart(body[cut]) {
			cut--
		}
		body = body[:cut]
		truncated = true
	}
	var out strings.Builder
	for _, r := range body {
		if r == utf8.RuneError || unicode.IsControl(r) || unicode.Is(unicode.Bidi_Control, r) {
			quoted := strconv.QuoteRuneToASCII(r)
			out.WriteString(quoted[1 : len(quoted)-1])
			continue
		}
		out.WriteRune(r)
	}
	if truncated {
		out.WriteString("…(truncated)")
	}
	return out.String()
}
