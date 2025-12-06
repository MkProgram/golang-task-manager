# Repository Guidelines

## Project Structure & Module Organization
- Go module lives at the repo root (`go.mod`). The current entrypoint is `main.go`.
- Prefer package code under `internal/<feature>/` for app logic and `pkg/<name>/` only for code meant to be reused externally.
- Place tests alongside code as `*_test.go`; fixtures can live under `testdata/`.
- Keep generated or built artifacts out of the repo; if you need binaries, use `bin/` locally and add it to `.gitignore`.

## Build, Test, and Development Commands
- `go fmt ./...` format all Go files; required before committing.
- `go vet ./...` static analysis to catch common issues.
- `go test ./...` run the full test suite; add `-race` when touching concurrency.
- `go run main.go` quick local run for the CLI entrypoint.
- `go build -o bin/task-manager ./...` produce a binary for manual testing.

## Coding Style & Naming Conventions
- Follow standard Go style: tabs for indentation; `gofmt` is the source of truth.
- Package names should be short, lower-case, and singular; files use lower-case with underscores only when clarifying (`task_service.go`).
- Exported identifiers use PascalCase only when needed outside the package; keep internals unexported (`camelCase`) where possible.
- Errors: use `err` suffix (`taskErr`), wrap with context using `fmt.Errorf("action: %w", err)`; prefer sentinel errors over booleans.
- Keep functions small and focused; favor pure functions for business logic.

## Testing Guidelines
- Use the standard `testing` package; structure tests as table-driven where meaningful.
- Name tests `Test<Thing>` and helpers `new<Thing>ForTest` or `must<Thing>`.
- Aim for coverage of edge cases (empty input, invalid states, I/O failures). Include race detector runs for concurrent code paths.

## Commit & Pull Request Guidelines
- Commit messages should be short, imperative, and scoped (e.g., `Add task validation`, `Refine CLI flags`).
- In PRs, include: goal/summary, key changes, tests performed (`go test ./...`, `go vet ./...`), and any follow-ups.
- Link issues or tickets when available; add screenshots or sample commands when behavior changes.

## Security & Configuration Tips
- Do not commit secrets; use environment variables or `.env.local` files kept out of git.
- Validate inputs from files/CLI flags; default to least-privilege when adding external calls or file writes.

## Tutoring Guidance
- Give hints first; guide the user’s thinking instead of providing direct fixes.
- Never show code unless the user explicitly requests it.
- Do not add code to this repository until the user says otherwise.
- Scope tutoring to Go topics; avoid unrelated language examples.
