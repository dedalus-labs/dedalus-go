# Changelog

## 0.0.4 (2026-03-25)

Full Changelog: [v0.0.3...v0.0.4](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.3...v0.0.4)

### Bug Fixes

* **api:** rename StreamStatus to Watch, remove wake_if_needed params, update workspace states ([5a28130](https://github.com/dedalus-labs/dedalus-go/commit/5a28130ee20552b3dbf3a1097ab8574b551f36e7))


### Chores

* **ci:** skip lint on metadata-only changes ([c17f53f](https://github.com/dedalus-labs/dedalus-go/commit/c17f53f08433c343cd3a5212be62e071a04d7252))
* **tests:** bump steady to v0.19.7 ([715d3e7](https://github.com/dedalus-labs/dedalus-go/commit/715d3e78d302a6fd046d3cca7bd931b3ba5492e9))

## 0.0.3 (2026-03-23)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.2...v0.0.3)

### Bug Fixes

* **api:** add workspace status streaming, terminal events types ([cb4ce8c](https://github.com/dedalus-labs/dedalus-go/commit/cb4ce8cc93f61d1022efe7bffca296ff7b4dc269))


### Chores

* **internal:** update gitignore ([2792731](https://github.com/dedalus-labs/dedalus-go/commit/2792731cdc4918c80d06b4c1476e0180fc69cacb))
* **tests:** bump steady to v0.19.4 ([4d8986e](https://github.com/dedalus-labs/dedalus-go/commit/4d8986ea15ca11981b9453150877bd6c6e6acaa7))
* **tests:** bump steady to v0.19.5 ([4da56b4](https://github.com/dedalus-labs/dedalus-go/commit/4da56b4f0566678cf86c80fa27f666f9c45e15f9))
* **tests:** bump steady to v0.19.6 ([d0822f9](https://github.com/dedalus-labs/dedalus-go/commit/d0822f939fa85458477a46ade3b5a53ccb295e4e))


### Refactors

* **tests:** switch from prism to steady ([58bc744](https://github.com/dedalus-labs/dedalus-go/commit/58bc744da3f352bc4a43654635bcea1ed7dd8898))

## 0.0.2 (2026-03-18)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.1...v0.0.2)

### Bug Fixes

* **api:** consolidate pagination & disable websockets ([dc840eb](https://github.com/dedalus-labs/dedalus-go/commit/dc840ebddc4763dc2648c42ac4ad28b7f3fc1e2a))


### Chores

* **api:** update homebrew tap and code samples ([d2cc1e5](https://github.com/dedalus-labs/dedalus-go/commit/d2cc1e5993bf3ef0a767e342997f6baff6c66bf1))

## 0.0.1 (2026-03-18)

Full Changelog: [v0.0.1...v0.0.1](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.1...v0.0.1)

### Features

* **api:** config update for dedalus-ai/dev ([efb7cd4](https://github.com/dedalus-labs/dedalus-go/commit/efb7cd447555b5843cb0ca09b527c487d648a445))
* **api:** stable beta ([625182e](https://github.com/dedalus-labs/dedalus-go/commit/625182e9a4c61cfd499e68bb9d720b5e854a6bca))
* **client:** add a convenient param.SetJSON helper ([320a14a](https://github.com/dedalus-labs/dedalus-go/commit/320a14a2596de8e00a963d9f6eae99d7f480c1f2))
* **encoder:** support bracket encoding form-data object members ([efe15cb](https://github.com/dedalus-labs/dedalus-go/commit/efe15cb0999a864413fe24070af36a119a251258))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([127e012](https://github.com/dedalus-labs/dedalus-go/commit/127e012c1ab252eb9d34593e03f59bd19460e2c7))
* **api:** update flags ([709491c](https://github.com/dedalus-labs/dedalus-go/commit/709491c0c86d00eb4157f67a1cc2d1b77a365953))
* bugfix for setting JSON keys with special characters ([08eea40](https://github.com/dedalus-labs/dedalus-go/commit/08eea40b01dc984c6fa049caed18201d61ab1f84))
* **client:** correctly specify Accept header with */* instead of empty ([f5c0379](https://github.com/dedalus-labs/dedalus-go/commit/f5c0379bfd745bfb7d3297fedb3e89484ddaed7e))
* **docs:** add missing pointer prefix to api.md return types ([ed51631](https://github.com/dedalus-labs/dedalus-go/commit/ed516312f3a47756e0067a31848f572e925171c7))
* **docs:** fix mcp installation instructions for remote servers ([a8278fb](https://github.com/dedalus-labs/dedalus-go/commit/a8278fb807ab106f815c108cb48802e44d0313f6))
* **encoder:** correctly serialize NullStruct ([f17f3cb](https://github.com/dedalus-labs/dedalus-go/commit/f17f3cb1659b0652553f87ec736316ee0532d477))
* fix request delays for retrying to be more respectful of high requested delays ([aa84cfa](https://github.com/dedalus-labs/dedalus-go/commit/aa84cfaa6b89d78332e5a4b8b93c5a42c368802a))
* **mcp:** correct code tool API endpoint ([80fb27b](https://github.com/dedalus-labs/dedalus-go/commit/80fb27bfb28f390202d005b231a6aa8a53cf8df1))
* rename param to avoid collision ([1fb7da1](https://github.com/dedalus-labs/dedalus-go/commit/1fb7da10556985335542581a31eac5a90c308857))
* skip usage tests that don't work with Prism ([aed2ecd](https://github.com/dedalus-labs/dedalus-go/commit/aed2ecd3c72ac3c80744e59d974d8af60089707c))


### Chores

* add float64 to valid types for RegisterFieldValidator ([53fb6de](https://github.com/dedalus-labs/dedalus-go/commit/53fb6de7fa499b0646e4d9920893be583fbcdff3))
* **api:** resolving merge conflicts ([0a6027f](https://github.com/dedalus-labs/dedalus-go/commit/0a6027fce0e7816ac621007a5236e489794bff8f))
* bump gjson version ([f5a2dd8](https://github.com/dedalus-labs/dedalus-go/commit/f5a2dd8e4ea6fbf1e63121636fb4c7d80adcc890))
* **ci:** add build step ([07af61f](https://github.com/dedalus-labs/dedalus-go/commit/07af61f2c458064711e34c77d390fbbf90b26da1))
* **ci:** skip uploading artifacts on stainless-internal branches ([d5cdb79](https://github.com/dedalus-labs/dedalus-go/commit/d5cdb79b06e66b348ebbc9149f1fb162a6c78a46))
* configure new SDK language ([a1b0f8f](https://github.com/dedalus-labs/dedalus-go/commit/a1b0f8fcc5469e666e3e1bc58e54c5a0702b1343))
* **docs:** add missing descriptions ([823effe](https://github.com/dedalus-labs/dedalus-go/commit/823effe108ba4324fd67ee0a465b131e9d878e4a))
* elide duplicate aliases ([53b49ac](https://github.com/dedalus-labs/dedalus-go/commit/53b49ac116aa55c94c65b1b67005bd8b1f46a8d3))
* **internal:** codegen related update ([65de220](https://github.com/dedalus-labs/dedalus-go/commit/65de220ffec39712ed95f844dbbf9307f3a4f40e))
* **internal:** codegen related update ([60dc5cf](https://github.com/dedalus-labs/dedalus-go/commit/60dc5cfb9ea894cec3f77d37642ba50cc20275c6))
* **internal:** grammar fix (it's -&gt; its) ([33780b1](https://github.com/dedalus-labs/dedalus-go/commit/33780b1eb6304a52de6af841711c94deed6e3f71))
* **internal:** minor cleanup ([16fbfe9](https://github.com/dedalus-labs/dedalus-go/commit/16fbfe9c0f1545d004a67ae62d34c873db619e2e))
* **internal:** move custom custom `json` tags to `api` ([2e71ce5](https://github.com/dedalus-labs/dedalus-go/commit/2e71ce58b0f6da528a63501077a476100aaafee7))
* **internal:** remove mock server code ([407fd7e](https://github.com/dedalus-labs/dedalus-go/commit/407fd7e94e36924020e5f45205f710415c095b94))
* **internal:** tweak CI branches ([e5e1a32](https://github.com/dedalus-labs/dedalus-go/commit/e5e1a32732a987631af7b76be55b0b67d3ce5dea))
* **internal:** update `actions/checkout` version ([5de9dde](https://github.com/dedalus-labs/dedalus-go/commit/5de9ddebe612319feff93f9a8d719617f6a29bb9))
* **internal:** use explicit returns ([3710b82](https://github.com/dedalus-labs/dedalus-go/commit/3710b8216f26ccdf11c137a7de7a4747190db5e9))
* **internal:** use explicit returns in more places ([631cf20](https://github.com/dedalus-labs/dedalus-go/commit/631cf2079ca51f71e378e09c3e332e0f6f87d16a))
* update mock server docs ([87edf2e](https://github.com/dedalus-labs/dedalus-go/commit/87edf2e08d8d70b1ea4e3122ad7f5de159a235c6))
* update placeholder string ([9be6e6d](https://github.com/dedalus-labs/dedalus-go/commit/9be6e6d8e1033b92f0250cdaa1af86dd79f0e3ff))
* update SDK settings ([5fd47f5](https://github.com/dedalus-labs/dedalus-go/commit/5fd47f54bbb7e62e334a322e096ed750d6062b3d))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([e95906b](https://github.com/dedalus-labs/dedalus-go/commit/e95906b60f27d9cbc1ea121a174264fe529d8e18))
