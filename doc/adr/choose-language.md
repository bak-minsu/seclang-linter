# Choose Language for Project

## Status

Accepted

## Context

This project needs to be written in a language that allows this linting tool to run in commit pipelines.

## Decision

This linting project will be written in Go (Golang).

## Consequences

Pros:

- This allows the same community of people who may be contributring to to Coraza, also written in Go, to contribute to this project.
- Go compiles natively to support many architectures and operating systems, allowing for use in many popular systems.
- Go has a simple testing framework which allows for quick unit testing without dependencies, besides the go binary.

Cons:

- Contributors will need to learn Go to contribute to this project, likely limiting the contributor pool to Go developers.
- Go is a general purpose language, and does not specifically target syntax parsing, therefore does not have quality-of-life features such as a pattern matching syntax for regex.