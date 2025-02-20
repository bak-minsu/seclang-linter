# Choose a Pipeline

## Status

Approved

## Context

We need a pipeline that is able to run generic pieces automation on committed and merged code. The automation that this pipeline supports includes but is not limited to:
- Testing the code
- Creating versioned release tags
- Update README.md badges

## Decision

We choose GitHub Actions as our pipeline.

## Consequences

Pros:
- GitHub Actions run when commits and merges are made within GitHub, and can be set to not require manual kickoff. 
- GitHub Actions has a free tier that allows contributors to run its actions themselves.
- No additional dependencies are needed to use GitHub actions, keeping contributions simple.

Cons:
- GitHub Actions have significant limits within its free tier, which is likely the tier that we will use for this project