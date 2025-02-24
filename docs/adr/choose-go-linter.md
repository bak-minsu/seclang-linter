# Choose a linter for Go

## Status

Accepted

## Context

We should add a linter for Go, as linters catch a lot of common mistakes that may not be caught by the compiler, before release.

## Decision

We propose to use golangci-lint as the linter tool.

## Consequences

Pros:
- This linter is a common CI linting tool for Go and has large community support.
- This linter has frequent updates, decommissioning old, abandoned linters, and adding new and improved features for static security checks.

Cons:
- No known downsides.