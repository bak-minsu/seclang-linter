# Parse Syntax

## Status

Accepted

## Context

We need to parse SecLang syntax in a way that is maintainable for the linter contributors.

## Decision

We will be using an AST (abstract syntax tree) to parse SecLang, to parse the syntax before running analysis.

## Consequences

Pros:
- ASTs are a well understood way to represent programming language syntax.

Cons:
- There may be lighter-weight ways to target certain types of syntax errors.