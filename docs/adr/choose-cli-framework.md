# Choose CLI Framework

## Status

Accepted

## Context

We need a framework for CLI input parsing and auto-completion, to match modern expectations for shell tooling.

## Decision

The CLI will be built using the [Cobra](https://github.com/spf13) framework.

## Consequences

Pros:
- Cobra receives very frequent updates, meaning any security vulnerabilities will likely be patched quicker, and will more likely implement features that are introduced in newer versions of Go quicker that other, less often maintained libraries.
- Cobra is popular and contributors are likely familiar with its tooling.

Cosn:
- Cobra does not fit neatly into Go's idioms, causing readability to suffer until contributor has read a bit into its stylistic guidelines.