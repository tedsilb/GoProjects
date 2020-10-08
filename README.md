# Go Projects

[![CodeFactor](https://www.codefactor.io/repository/github/tedsilb/goprojects/badge)](https://www.codefactor.io/repository/github/tedsilb/goprojects)

Various Go projects I work on from time to time.

## Building

Projects are built using [Bazel](https://bazel.build).

- To build all projects:
  - `bazel build ...`
- To run a specific project:
  - `bazel run projects/{project} {args}`
- For example:
  - `bazel run projects/DescriptiveStats`
