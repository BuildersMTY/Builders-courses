# BuildersMTY — Test System Spec

## Overview

Tests in BuildersMTY courses are defined declaratively in `course.yaml`. The course author never writes a test runner — they describe the expected behavior and the platform executes it. A global tester reads the `type` field of each test and dispatches to the correct execution handler.

This design keeps contribution fast: most courses only need `stdout` or `http` tests. Advanced courses can opt into protocol-level or scripted testing without changing the schema.

---

## Test Types

### `stdout`

The simplest test type. The runner executes the student's program, optionally feeds it stdin, and compares stdout against the expected output.

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

Runs an external test script or binary against the student's program. The escape hatch for cases that cannot be expressed declaratively — concurrency tests, filesystem state verification, multi-step stateful interactions.

Use for: Git (verify repo state after commands), Docker (verify namespace isolation), concurrency tests, anything requiring programmatic assertions.

```yaml
tests:
  - type: script
    file: tests/concurrency_test.go
    timeout_ms: 10000
```

The script receives the student's workspace path via `BUILDERSMTY_WORKSPACE_DIR` and the running program's PID via `BUILDERSMTY_PROGRAM_PID`. It must exit 0 on pass and non-zero on failure. Stdout is streamed to the student in real time.

Scripts are stored in the course directory under `tests/` and are version-controlled alongside the course content. They can be written in any language available in the course runner image.

---

## Test Execution Flow

```
student clicks "Run Tests" for submodule X
  → platform reads tests[] from course.yaml for that submodule
  → runner container is pulled from the pool
  → student's files are mounted as tmpfs
  → for each test in order:
      → dispatch to handler based on type
      → capture result: passed | failed + diff
      → stream output to browser in real time
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
| HTTP Server      | `http`       | none                                                          |
| Memory Allocator | `stdout`     | `script` for heap state inspection                            |
| Claude Code      | `stdout`     | `script` for tool-use verification                            |
| Redis            | `tcp`        | `script` for persistence tests                                |
| Shell            | `stdout`     | none                                                          |
| Git              | `stdout`     | `script` for repo state verification                          |
| Docker           | `script`     | none — namespace/cgroup tests require programmatic assertions |
| Compiler         | `stdout`     | none                                                          |
| Load Balancer    | `http`       | `script` for distribution verification                        |
| Auth Server      | `http`       | none                                                          |
