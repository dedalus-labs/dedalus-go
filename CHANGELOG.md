# Changelog

## [0.6.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.5.0...v0.6.0) (2026-09-24)


### ⚠ BREAKING CHANGES

* **api:** 3 breaking changes to the SDK surface.
    - Removed `bearer` auth scheme `Bearer`.
    - Removed operation `machines.terminals.connect` (`GET /v1/machines/{machine_id}/terminals/{terminal_id}/stream`).
    - Removed optional property `RetrieveResponseHeaders.ETag`.
* **api:** 72 breaking changes to the SDK surface.
    - Configuration of `bearer` auth scheme `BearerAuth` changed.
    - Removed header param `X-Dedalus-Org-Id` from `machines.list`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.create`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.retrieve`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.update`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.delete`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.sleep`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.wake`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.ssh.list`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.ssh.create`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.ssh.retrieve`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.ssh.delete`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.list`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.create`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.retrieve`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.delete`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.output`.
    - Removed header param `X-Dedalus-Org-Id` from `machines.executions.events`.
    - Serialization or defaults of path param `machine_id` on `machines.terminals.connect` changed.
    - Serialization or defaults of path param `terminal_id` on `machines.terminals.connect` changed.
    - Removed header param `X-Dedalus-Org-Id` from `machines.terminals.connect`.
    - Removed operation `machines.watch` (`GET /v1/machines/{machine_id}/status/stream`).
    - Removed operation `machines.network.retrieve` (`GET /v1/machines/{machine_id}/network`).
    - Removed operation `machines.artifacts.list` (`GET /v1/machines/{machine_id}/artifacts`).
    - Removed operation `machines.artifacts.retrieve` (`GET /v1/machines/{machine_id}/artifacts/{artifact_id}`).
    - Removed operation `machines.artifacts.delete` (`DELETE /v1/machines/{machine_id}/artifacts/{artifact_id}`).
    - Removed operation `machines.ports.list` (`GET /v1/machines/{machine_id}/ports`).
    - Removed operation `machines.ports.create` (`POST /v1/machines/{machine_id}/ports`).
    - Removed operation `machines.ports.retrieve` (`GET /v1/machines/{machine_id}/ports/{port_id}`).
    - Removed operation `machines.ports.delete` (`DELETE /v1/machines/{machine_id}/ports/{port_id}`).
    - Removed operation `machines.terminals.list` (`GET /v1/machines/{machine_id}/terminals`).
    - Removed operation `machines.terminals.create` (`POST /v1/machines/{machine_id}/terminals`).
    - Removed operation `machines.terminals.retrieve` (`GET /v1/machines/{machine_id}/terminals/{terminal_id}`).
    - Removed operation `machines.terminals.delete` (`DELETE /v1/machines/{machine_id}/terminals/{terminal_id}`).
    - Removed operation `networks.retrieve` (`GET /v1/networks/{network_id}`).
    - Removed operation `usage.retrieve` (`GET /v1/usage`).
    - Removed operation `usage.machineCompute` (`GET /v1/usage/machines/compute`).
    - Removed operation `usage.machineStorage` (`GET /v1/usage/machines/storage`).
    - Added required property `execution.log_capture`.
    - Property `execution.machine_id` type changed from `string` to `string<uuid>`.
    - Property `machine.machine_id` type changed from `string` to `string<uuid>`.
    - Added required property `lifecycle_status.memory_configured_mib`.
    - Property `machine_detail_response.machine_id` type changed from `string` to `string<uuid>`.
    - Schema `machine_id_path_segment` shape changed.
    - Property `machine_list_item.machine_id` type changed from `string` to `string<uuid>`.
    - Property `ssh_session.machine_id` type changed from `string` to `string<uuid>`.
    - Property `ssh_session.status` type changed from `enum(wake_in_progress | ready | closed | …)` to `enum(wake_in_progress | ssh_in_progress | ready | …)`.
    - Removed optional property `CreateResponseHeaders.ETag`.
    - Removed optional property `CreateResponseHeaders.X-Dedalus-Storage-Operation-Id`.
    - Removed schema `artifact_list`.
    - Removed schema `artifact`.
    - Removed schema `port_create_params`.
    - Removed schema `terminal_create_params`.
    - Removed schema `machine_compute_usage`.
    - Removed schema `machine_compute_usage_row`.
    - Removed schema `machine_network`.
    - Removed schema `machine_storage_usage`.
    - Removed schema `machine_storage_usage_row`.
    - Removed schema `network_gateway`.
    - Removed schema `network`.
    - Removed schema `port_list`.
    - Removed schema `port`.
    - Removed schema `terminal_client_event`.
    - Removed schema `terminal_closed_event`.
    - Removed schema `terminal_error_event`.
    - Removed schema `terminal_input_event`.
    - Removed schema `terminal_list`.
    - Removed schema `terminal_output_event`.
    - Removed schema `terminal_resize_event`.
    - Removed schema `terminal`.
    - Removed schema `terminal_server_event`.
    - Removed schema `org_usage`.

### Features

* **api:** initial SDK generation ([9022c5d](https://github.com/dedalus-labs/dedalus-go/commit/9022c5da91f3c8cf9e3b9456032a54dbf19461d7))
* **api:** remove auth scheme Bearer (+25 more changes) ([3cc365d](https://github.com/dedalus-labs/dedalus-go/commit/3cc365d036a91426924757459165fe9a931a22b1))
* **api:** update auth scheme BearerAuth (+122 more changes) ([42e1cd5](https://github.com/dedalus-labs/dedalus-go/commit/42e1cd5dc56f30599748e0e0327dfbce774488dc))


### Chores

* **api:** update generated SDK content ([e8ebb95](https://github.com/dedalus-labs/dedalus-go/commit/e8ebb9529a06e4628cbd12b3d7e0054281fb1d4a))

## 0.5.0 (2026-09-01)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.4.0...v0.5.0)

### ⚠ BREAKING CHANGES

* **api:** regenerate SDKs from the current public DCS contract

### Features

* **api:** regenerate SDKs from the current public DCS contract ([3062ff5](https://github.com/dedalus-labs/dedalus-go/commit/3062ff59c8f1dc066bc47884773ddf55339c24d4))
* **client:** optimize json encoder for internal types ([165d9dd](https://github.com/dedalus-labs/dedalus-go/commit/165d9dd185ca38e330862cde49a6df1c9ed49562))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([c239ff7](https://github.com/dedalus-labs/dedalus-go/commit/c239ff7504684070a127e69862a63bce3e4af594))


### Chores

* **internal:** allow the mock server port to be set with STAINLESS_MOCK_PORT ([68b8030](https://github.com/dedalus-labs/dedalus-go/commit/68b803017b6272654d7ba5f820a5e16214df3e9e))

## 0.4.0 (2026-05-12)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** move usage to top-level resource, remove Orgs, rename methods/types ([52c734c](https://github.com/dedalus-labs/dedalus-go/commit/52c734cfd7789be6f708a96fd2bf0757a92bdee0))

## 0.3.0 (2026-05-12)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** add org usage endpoints, autosleep to machines, remove IfMatch header ([9cac88d](https://github.com/dedalus-labs/dedalus-go/commit/9cac88dbd5e5ab39c511a3cf886729378bb1a660))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([db937b3](https://github.com/dedalus-labs/dedalus-go/commit/db937b3c7440745d0bdffe637284ceef4641f25b))


### Chores

* avoid embedding reflect.Type for dead code elimination ([693bce7](https://github.com/dedalus-labs/dedalus-go/commit/693bce7245bc89741966acbe5956c199939b9f0f))
* redact api-key headers in debug logs ([85071cf](https://github.com/dedalus-labs/dedalus-go/commit/85071cfcc13bd6e0e224097107cf0b208850e1ef))

## 0.2.0 (2026-04-27)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.1.0...v0.2.0)

### Features

* **go:** add default http client with timeout ([0274dc8](https://github.com/dedalus-labs/dedalus-go/commit/0274dc855748725fde1204fd708d4043e8cbb276))
* support setting headers via env ([8db7f42](https://github.com/dedalus-labs/dedalus-go/commit/8db7f420775ef3ead170a022c03e5c3067b405e4))


### Chores

* **internal:** codegen related update ([9bacd0e](https://github.com/dedalus-labs/dedalus-go/commit/9bacd0e8816fd837eba3ec48c111eeb5a885bbb5))
* **internal:** more robust bootstrap script ([cdecb20](https://github.com/dedalus-labs/dedalus-go/commit/cdecb20135dad2b47f1c7f83058a9a8e8bd364df))
* **tests:** bump steady to v0.22.1 ([59b57d6](https://github.com/dedalus-labs/dedalus-go/commit/59b57d6852aac7819fe8ff68604e73aec283891c))

## 0.1.0 (2026-04-02)

Full Changelog: [v0.0.5...v0.1.0](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.5...v0.1.0)

### Features

* **api:** add sleep & wake methods ([867c576](https://github.com/dedalus-labs/dedalus-go/commit/867c5762ec96afec74583964b5d904ea956c639a))

## 0.0.5 (2026-04-01)

Full Changelog: [v0.0.4...v0.0.5](https://github.com/dedalus-labs/dedalus-go/compare/v0.0.4...v0.0.5)

### Features

* **internal:** support comma format in multipart form encoding ([b6e4a66](https://github.com/dedalus-labs/dedalus-go/commit/b6e4a66de5e3dca7e6e3019770f9f605413d052f))


### Bug Fixes

* fix issue with unmarshaling in some cases ([67e8dd2](https://github.com/dedalus-labs/dedalus-go/commit/67e8dd2cd5611bd68d17be2ab8879b4d92be775e))
* prevent duplicate ? in query params ([499b0b9](https://github.com/dedalus-labs/dedalus-go/commit/499b0b9e2dacc422cc3d9f9151bb6e1a715d3524))


### Chores

* **api:** rename workspaces to machines ([dff4de2](https://github.com/dedalus-labs/dedalus-go/commit/dff4de28ecb059854de3d8022b77e765b4d1ca56))
* **ci:** support opting out of skipping builds on metadata-only commits ([5bbd0c1](https://github.com/dedalus-labs/dedalus-go/commit/5bbd0c13b84be6b815fdddeb29f525e4f73b9474))
* **client:** fix multipart serialisation of Default() fields ([2fb659b](https://github.com/dedalus-labs/dedalus-go/commit/2fb659bf2dcb63ab6b0ae4a17255950443c86609))
* **internal:** support default value struct tag ([04b200b](https://github.com/dedalus-labs/dedalus-go/commit/04b200b04f54bb3af1b068d5ecb2e1500d16a10a))
* remove unnecessary error check for url parsing ([ebb626d](https://github.com/dedalus-labs/dedalus-go/commit/ebb626d7f0225deea8a7465a6219ba955019f8e8))
* **tests:** bump steady to v0.20.1 ([c0aae27](https://github.com/dedalus-labs/dedalus-go/commit/c0aae2735eb5ee29ef92b00f038ca6adb872d271))
* **tests:** bump steady to v0.20.2 ([be7314e](https://github.com/dedalus-labs/dedalus-go/commit/be7314e7b99d42cde5052e4ad3e1e46cfde93efb))
* update docs for api:"required" ([128f066](https://github.com/dedalus-labs/dedalus-go/commit/128f066dedd518f68cf1df91a1153604d3b763c1))

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
