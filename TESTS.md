# BuildersMTY — Test System Spec

## Overview

Tests in BuildersMTY courses are defined declaratively in `course.yaml`. The course author never writes a test runner — they describe the expected behavior and the platform executes it. A global tester reads the `type` field of each test and dispatches to the correct execution handler.

This design keeps contribution fast: most courses only need `unit`, `stdout`, or `http` tests. Advanced courses can opt into protocol-level or scripted testing without changing the schema.

---

## Execution Model

The platform builds the student's project **once per submodule run** and then executes all tests in the submodule against that build.

```
submodule run
  ├─ build_cmd        ← compile (per-language default, see table below)
  ├─ artifact at $BUILDERSMTY_BINARY
  └─ for each test in tests[]:
       ├─ http  → spawn binary, wait for port, send request, kill
       ├─ tcp   → spawn binary, wait for port, send bytes, kill
       ├─ stdout → spawn binary, pipe stdin, capture stdout, kill
       ├─ unit  → run unit_cmd with test name filter, capture exit code
       └─ script → spawn binary (default), run script with env vars, kill
```

**Build artifact contract.** After `build_cmd` runs, an executable must exist at `$BUILDERSMTY_BINARY` (default: `/tmp/program`). All test types that need to run code use this path. Scripts get the path via the env var so they can invoke it themselves.

**Program lifecycle within a submodule.** For tests that need a running program, the platform starts it **once at the beginning of the submodule run**, executes all tests in order against the same process, and kills it at the end. Tests do NOT get a fresh process per test. This lets contributors write stateful test sequences (e.g. POST then GET) without managing lifecycle.

**Failures.** Build failure → submodule fails immediately, no tests run. Spawn failure (binary crashes before becoming ready) → reported as a distinct error to the student, not a test timeout.

---

## Language Defaults

The platform has built-in defaults per language. Contributors only override these in `course.yaml` when they have custom needs.

| Language | `build_cmd` | `run_cmd` | `unit_cmd` |
|---|---|---|---|
| `go` | `go build -o $BUILDERSMTY_BINARY .` | `$BUILDERSMTY_BINARY` | `go test -run {match} -v -count=1 .` |
| `rust` | `cargo build --release && cp target/release/* $BUILDERSMTY_BINARY` | `$BUILDERSMTY_BINARY` | `cargo test {match}` |
| `python` | (no build step) | `python main.py` | `pytest -k {match}` |
| `c` | `make` (or `cc -o $BUILDERSMTY_BINARY *.c`) | `$BUILDERSMTY_BINARY` | (no convention; use `script`) |
| `javascript` | `npm install && npm run build` | `node dist/index.js` | `vitest run -t {match}` |

`{match}` is interpolated from each `unit` test's `match` field.

**Overriding defaults.** Add to `meta` in `course.yaml`:

```yaml
meta:
  language: go
  build_cmd: "go build -tags integration -o $BUILDERSMTY_BINARY ./cmd/server"
  run_cmd: "$BUILDERSMTY_BINARY --port 8080"
  unit_cmd: "go test -tags integration -run {match} -v ./..."
```

A submodule can also override `build_cmd`/`run_cmd`/`unit_cmd` for one specific submodule. Most courses never override anything.

---

## Test Types

### `unit`

Runs the language's native test runner against a named test in the student's project. The platform invokes `unit_cmd` (per-language default) with the test name interpolated, captures the exit code, and passes if exit is 0.

Use for: pure libraries (parsers, lexers, regex engines), helper functions, anything that's naturally tested with the language's own test framework. The platform never needs to spawn the student's binary for `unit` tests.

```yaml
tests:
  - type: unit
    match: TestParseRequestLine
    timeout_ms: 5000
```

`match` is interpolated into `unit_cmd`. For Go that becomes `go test -run TestParseRequestLine -v -count=1 .`. For Rust, `cargo test test_parse_request_line`. For Python, `pytest -k test_parse_request_line`.

The student writes test files in their language's standard location (`*_test.go`, `tests/`, `test_*.py`, etc.). These files are committed to the course repo as part of the stubs and are visible to the student so they can read what's being tested.

---

### `stdout`

The simplest test type for one-shot CLIs. The runner executes the student's program, optionally feeds it stdin, and compares stdout against the expected output.

Use for: compilers, interpreters, shells, CLIs, memory allocators, any program that communicates via stdio.

```yaml
tests:
  - type: stdout
    stdin: "hello world"
    expected_stdout: "hello world\n"
    timeout_ms: 3000
```

**Matching behavior:** exact match by default. Trailing newline differences are normalized. Use `contains` mode for dynamic output:

```yaml
tests:
  - type: stdout
    stdin: ""
    expected_stdout_contains: "listening on port"
    timeout_ms: 3000
```

---

### `http`

Starts the student's program, waits for it to be ready, then sends an HTTP request and validates the response. Tests at the HTTP semantic level — status code, headers, body — without caring about byte-level formatting, header ordering, or whitespace.

Use for: HTTP servers, REST APIs, web frameworks, reverse proxies.

```yaml
tests:
  - type: http
    request:
      method: GET
      path: /health
      headers:
        Host: localhost
    expected:
      status: 200
      body_contains: '"status":"ok"'
      timeout_ms: 3000
```

**Available assertions:**

```yaml
expected:
  status: 200 # exact HTTP status code
  body_contains: "string" # body includes this string
  body_equals: '{"status":"ok"}' # exact body match
  headers:
    Content-Type: "application/json" # header value contains this string
```

Multiple assertions in a single test all must pass.

---

### `tcp`

Sends raw bytes to the student's program over a TCP connection and validates the raw response. Operates at the protocol level — no HTTP abstraction, no parsing.

Use for: Redis (RESP protocol), custom protocols, binary protocols, anything where HTTP semantics don't apply.

Bytes are expressed as hex strings for readability:

```yaml
tests:
  - type: tcp
    port: 6379
    send_hex: "2b 50 49 4e 47 0d 0a" # RESP: +PING\r\n
    expected_hex: "2b 50 4f 4e 47 0d 0a" # RESP: +PONG\r\n
    timeout_ms: 3000
```

Plain strings are also supported for text-based protocols:

```yaml
tests:
  - type: tcp
    port: 6379
    send: "*1\r\n$4\r\nPING\r\n"
    expected: "+PONG\r\n"
    timeout_ms: 3000
```

---

### `script`

Runs an external test script or binary against the student's program. The escape hatch for cases that cannot be expressed declaratively — concurrency tests, filesystem state verification, multi-step stateful interactions, CLI tools invoked many times with different args.

Use for: Git (run `$BUILDERSMTY_BINARY init`, `$BUILDERSMTY_BINARY add foo`, then verify repo state), Docker (verify namespace isolation), DB persistence (kill+restart the binary), concurrency tests, anything requiring programmatic assertions.

```yaml
tests:
  - type: script
    file: tests/concurrency_test.sh
    timeout_ms: 10000
```

**Default lifecycle**: the platform spawns the student's binary before running the script (so HTTP/TCP scripts can hit it directly) and kills it when the script exits.

**Manual lifecycle** (for tests that need to spawn/kill the binary themselves, like DB persistence):

```yaml
tests:
  - type: script
    file: tests/persistence_test.sh
    manages_lifecycle: true   # platform does NOT spawn the binary
    timeout_ms: 10000
```

When `manages_lifecycle: true`, the script gets `$BUILDERSMTY_BINARY` (path to the built binary) and is fully responsible for spawning, killing, and restarting it.

**Environment variables passed to scripts:**

| Variable | Always set | Description |
|---|---|---|
| `BUILDERSMTY_WORKSPACE_DIR` | yes | Path to the student's mounted workspace (read-write tmpfs) |
| `BUILDERSMTY_BINARY` | yes | Path to the built artifact (e.g. `/tmp/program`) |
| `BUILDERSMTY_PROGRAM_PID` | only when `manages_lifecycle: false` | PID of the spawned binary |

The script must exit 0 on pass and non-zero on failure. Stdout is streamed to the student in real time.

Scripts are stored in the course directory under `tests/` and are version-controlled alongside the course content. They can be written in any language available in the course runner image. The runner image guarantees `bash`, `curl`, `nc`, and the language toolchain are available.

---

## Test Execution Flow

```
student clicks "Run Tests" for submodule X
  → platform reads tests[] from course.yaml for that submodule
  → runner container is pulled from the pool
  → student's files are mounted as tmpfs
  → platform runs build_cmd (per-language default or override)
      → if build fails: report build error, stop
  → platform spawns binary if any test in submodule needs it (http/tcp/script with default lifecycle)
      → wait for readiness (port bound, or short delay if no port declared)
  → for each test in order:
      → dispatch to handler based on type
      → capture result: passed | failed + diff
      → stream output to browser in real time
  → kill spawned binary (SIGTERM, then SIGKILL after grace period)
  → if all tests pass: submodule marked as passed, git commit
  → if any test fails: show which test failed and why, student stays on submodule
  → container returned to pool, tmpfs destroyed
```

---

## Module Integration Tests

When all submodules in a module pass, the platform automatically runs the module's integration test. This test verifies that all components work together, not just individually.

Integration tests use the same four types. They are defined at the module level in `course.yaml`:

```yaml
modules:
  - id: tcp
    title: "TCP Foundation"
    integration_test:
      type: http
      request:
        method: GET
        path: /
      expected:
        status: 200
    submodules:
      - ...
```

---

## Stateful Tests (v2)

For courses where a test depends on the state left by a previous test (e.g. Redis SET then GET), a `depends_on` field will be supported in a future version:

```yaml
tests:
  - id: set_key
    type: tcp
    send: "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
    expected: "+OK\r\n"

  - id: get_key
    type: tcp
    depends_on: set_key # runs in the same session, state is preserved
    send: "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n"
    expected: "$3\r\nbar\r\n"
```

Not implemented in v1. For now, stateful scenarios should be handled via `type: script`.

---

## Coverage by Course

| Course           | Primary type | Edge cases                                                    |
| ---------------- | ------------ | ------------------------------------------------------------- |
| HTTP Server      | `http`       | `unit` for parsing helpers, `script` for path traversal       |
| Memory Allocator | `stdout`     | `script` for heap state inspection                            |
| Claude Code      | `stdout`     | `script` for tool-use verification                            |
| Redis            | `tcp`        | `script` (manages_lifecycle) for persistence tests            |
| Shell            | `stdout`     | none                                                          |
| Git              | `script`     | binary invoked many times via `$BUILDERSMTY_BINARY`           |
| Docker           | `script`     | none — namespace/cgroup tests require programmatic assertions |
| Compiler         | `stdout`     | `unit` for lexer/parser components                            |
| Load Balancer    | `http`       | `script` for distribution verification                        |
| Auth Server      | `http`       | none                                                          |
| Regex Engine     | `unit`       | none                                                          |
| React Clone      | `unit`       | none                                                          |
| Build Your DB    | `tcp`        | `script` (manages_lifecycle) for persistence/crash recovery   |
