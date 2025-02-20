# Design

## Problem Statement

[OWASP Coraza WAF](https://github.com/corazawaf/coraza) uses a configuration language called SecLang, which extends the OWASP Core Ruleset. This linter looks at given files written in Coraza's SecLang and validates the syntax. 

## Design Specification

For specifying design, this project uses Architecture decision records ([ADRs](https://github.com/joelparkerhenderson/architecture-decision-record?tab=readme-ov-file#how-to-start-using-adrs)).

Rather than specifying a large document that may eventually become obsolete, this system allows for an accumulation of decisions that are made for the project, which organically follows lean software development processes.

## Template for ADRs

This is the template in [Documenting architecture decisions - Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions).

In each ADR file, write these sections:

```markdown

# Title

## Status

What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?
```