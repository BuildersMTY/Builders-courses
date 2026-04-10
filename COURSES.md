# BuildersMTY — Courses Repository Spec

## Purpose

This repository contains all course content for the BuildersMTY learning platform. It is the single source of truth for course structure, stubs, tests, resources, and reference solutions. The platform (`builders-platform`) reads this repo to bootstrap student workspaces — it never writes here.

---

## Repository Structure

```
courses/
  {course-slug}/
    {language}/
      course.yaml       ← defines the entire course: modules, submodules, tests, resources
      src/              ← stubs: source files with empty function bodies for the student to fill
      solution/         ← complete working reference implementation, never exposed to students
      resources/        ← markdown docs organized by module/submodule
      makefiles/        ← one .mk file per submodule, used by the runner to execute tests
```

---

## course.yaml Schema

```yaml
meta:
  slug: string # unique, kebab-case, matches directory name
  title: string
  description: string
  language: string # go | c | python | rust | javascript
  difficulty: beginner | intermediate | advanced
  runner: string # docker image, e.g. buildersmty/runner-go:latest
  estimated_hours:
    junior: int
    mid: int
    senior: int

modules:
  - id: string # short, snake_case
    title: string
    description: string
    submodules:
      - id: string # short, snake_case
        title: string
        spec: string # technical description of what the student must implement
        stubs:
          - path: string # relative to src/, e.g. server.go
        makefile: string # relative to makefiles/, e.g. tcp-listen.mk
        tests:
          - type: submodule | module_integration
            stdin: string # input fed to the program, empty string if none
            expected_stdout: string # exact expected output
            timeout_ms: int
        resources:
          - title: string
            file: string # relative to resources/
            type: doc | spec | signature | hint
            visible_to:
              - junior | mid | senior
```

---

## Stubs

Stubs are source files with function signatures and empty bodies. They define what the student must implement without revealing how. Every function the student needs to fill must be present with its correct signature and a comment describing its expected behavior.

```go
// Listen opens a TCP socket on the given port and returns a net.Listener.
// It must return an error if the port is already in use.
func Listen(port int) (net.Listener, error) {
    // TODO: implement
}
```

---

## Solution

The `solution/` directory contains a complete, clean, well-commented implementation. It is used by:

- CI to validate that all tests pass before merging a course PR
- Course authors as the reference when writing stubs and tests

The solution is never served to students by the platform. It is excluded from enrollment repo initialization.

---

## Resources

Resources are markdown files. Each resource targets a specific submodule and a specific difficulty level. The platform filters resources based on the student's enrolled difficulty.

Resource types:

- `doc` — language or stdlib documentation relevant to the submodule
- `spec` — technical specification of the system being built (RFC, man page, protocol spec)
- `signature` — function signatures and type definitions relevant to the implementation
- `hint` — step-by-step implementation guidance, junior only in most cases

---

## Makefiles

Each submodule has its own `.mk` file that defines a `test` target. The runner executes `make -f {makefile} test` inside the student's workspace. The makefile must:

- Compile the code if the language requires it
- Run the test binary or script
- Exit 0 on pass, non-zero on failure
- Print output that matches `expected_stdout` exactly on success

```makefile
test:
	go build -o server . && \
	./server & sleep 0.5 && \
	curl -s http://localhost:8080/health && \
	kill %1
```

---

## Runner

Each language has a Docker image pre-built with all required tooling. The image is defined in a `Dockerfile` inside the course directory and published to Docker Hub under `buildersmty/runner-{language}`.

The runner is stateless — it receives the student's files via a tmpfs mount, executes the makefile, returns stdout/stderr, and is returned to the pool. No state persists between runs.

---

## Contributing a Course

1. Fork this repo
2. Copy an existing course directory as a template
3. Fill `course.yaml`, `src/`, `solution/`, `resources/`, `makefiles/`
4. Ensure `solution/` passes all tests defined in `course.yaml`
5. Open a PR — CI validates automatically

A PR must include:

- A complete `course.yaml`
- All stubs with correct signatures and TODO comments
- A working `solution/` that passes every test
- At least one resource per submodule visible to each difficulty level
- A `Dockerfile` for the runner if introducing a new language

---

## License

Content in this repository is licensed under **CC BY-NC 4.0**. Free to use, modify, and redistribute for non-commercial purposes. Commercial use requires explicit permission from BuildersMTY.
