# Router Extensibility Load Testing

**The code in this repository is experimental and has been provided for reference purposes only. Community feedback is welcome but this project may not be supported in the same way that repositories in the official [Apollo GraphQL GitHub organization](https://github.com/apollographql) are. If you need help you can file an issue on this repository, [contact Apollo](https://www.apollographql.com/contact-sales) to talk to an expert, or create a ticket directly in Apollo Studio.**

> Note: The Apollo Router is made available under the Elastic License v2.0 (ELv2).
> Read [our licensing page](https://www.apollographql.com/docs/resources/elastic-license-v2-faq/) for more details.

## Overview

This repository is a simple way to test the overhead of the three customization points of the Apollo Router:

- [Coprocessors](https://www.apollographql.com/docs/router/customizations/coprocessor)
- [Rhai](https://www.apollographql.com/docs/router/customizations/rhai)
- Configuration options

The current tests are:

- Setting a static header to subgraphs (Config, Rhai, Coprocessor)
- Setting 10 GUID headers on response to clients (Rhai, Coprocessor)
- JWT-based client awareness (Coprocessor)

The coprocessors are currently written in:

- [Go](./coprocessors/go/)
- [Node](./coprocessors/node)
- [C#](./coprocessors/csharp)
- [Java](./coprocessors/java)
- [Python](./coprocessors/python)

## Results

For the below tables, each section corresponds to the related test name. Each type relates to either the baseline (meaning no Router configuration), or the extensibility option. Languages imply a coprocessor.

The tests were run at 100 requests per second for 60 seconds against an Apollo Router version 1.19.0 on a Windows machine using WSL2 with Ubuntu. 

To help with consistency, there are resource limits for both the router and the coprocessors when using Docker--currently 1 CPU core and 1GB of RAM.

### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a coprocessor.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)        | Max (ms)        |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| baseline | 0.48            | 0.68            | 0.65            | 0.88            | 0.93            | 1.10            | 4.68            |
| csharp   | 0.69<br>(+0.21) | 0.90<br>(+0.22) | 0.87<br>(+0.22) | 1.09<br>(+0.21) | 1.15<br>(+0.22) | 1.27<br>(+0.17) | 6.10<br>(+1.42) |
| go       | 0.66<br>(+0.18) | 0.83<br>(+0.15) | 0.79<br>(+0.14) | 1.02<br>(+0.14) | 1.09<br>(+0.16) | 1.21<br>(+0.11) | 4.98<br>(+0.30) |
| java     | 0.72<br>(+0.24) | 0.99<br>(+0.31) | 0.94<br>(+0.29) | 1.24<br>(+0.36) | 1.36<br>(+0.43) | 1.64<br>(+0.54) | 7.11<br>(+2.43) |
| node     | 0.74<br>(+0.26) | 0.95<br>(+0.27) | 0.91<br>(+0.26) | 1.14<br>(+0.26) | 1.21<br>(+0.28) | 1.33<br>(+0.23) | 5.82<br>(+1.14) |
| python   | 0.81<br>(+0.33) | 1.00<br>(+0.32) | 0.96<br>(+0.31) | 1.20<br>(+0.32) | 1.27<br>(+0.34) | 1.37<br>(+0.27) | 5.78<br>(+1.10) |
| rhai     | 0.53<br>(+0.05) | 0.75<br>(+0.07) | 0.72<br>(+0.07) | 0.96<br>(+0.08) | 1.03<br>(+0.10) | 1.19<br>(+0.09) | 5.02<br>(+0.34) |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.

This is only available via a coprocessor.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)        | Max (ms)          |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- | ----------------- |
| baseline | 0.49            | 0.67            | 0.64            | 0.87            | 0.93            | 1.11            | 4.90              |
| csharp   | 0.76<br>(+0.27) | 1.02<br>(+0.35) | 0.98<br>(+0.34) | 1.17<br>(+0.30) | 1.25<br>(+0.32) | 1.42<br>(+0.31) | 32.42<br>(+27.52) |
| go       | 0.68<br>(+0.19) | 0.84<br>(+0.17) | 0.81<br>(+0.17) | 1.03<br>(+0.16) | 1.10<br>(+0.17) | 1.20<br>(+0.09) | 5.06<br>(+0.16)   |
| java     | 0.82<br>(+0.33) | 1.13<br>(+0.46) | 1.07<br>(+0.43) | 1.42<br>(+0.55) | 1.62<br>(+0.69) | 2.04<br>(+0.93) | 7.92<br>(+3.02)   |
| node     | 1.03<br>(+0.54) | 1.28<br>(+0.61) | 1.25<br>(+0.61) | 1.50<br>(+0.63) | 1.57<br>(+0.64) | 1.72<br>(+0.61) | 6.19<br>(+1.29)   |
| python   | 0.84<br>(+0.35) | 1.03<br>(+0.36) | 1.00<br>(+0.36) | 1.19<br>(+0.32) | 1.27<br>(+0.34) | 1.38<br>(+0.27) | 5.77<br>(+0.87)   |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)        | Max (ms)        |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| baseline | 0.48            | 0.69            | 0.65            | 0.88            | 0.94            | 1.12            | 4.72            |
| config   | 0.49<br>(+0.01) | 0.68<br>(-0.01) | 0.65<br>(0.00)  | 0.88<br>(0.00)  | 0.93<br>(-0.01) | 1.09<br>(-0.03) | 4.92<br>(+0.20) |
| csharp   | 0.80<br>(+0.32) | 1.06<br>(+0.37) | 1.04<br>(+0.39) | 1.23<br>(+0.35) | 1.29<br>(+0.35) | 1.39<br>(+0.27) | 5.77<br>(+1.05) |
| go       | 0.74<br>(+0.26) | 0.97<br>(+0.28) | 0.95<br>(+0.30) | 1.12<br>(+0.24) | 1.18<br>(+0.24) | 1.28<br>(+0.16) | 5.22<br>(+0.50) |
| java     | 0.81<br>(+0.33) | 1.14<br>(+0.45) | 1.08<br>(+0.43) | 1.41<br>(+0.53) | 1.58<br>(+0.64) | 2.09<br>(+0.97) | 7.58<br>(+2.86) |
| node     | 0.92<br>(+0.44) | 1.13<br>(+0.44) | 1.09<br>(+0.44) | 1.32<br>(+0.44) | 1.40<br>(+0.46) | 1.61<br>(+0.49) | 6.60<br>(+1.88) |
| python   | 0.60<br>(+0.12) | 0.73<br>(+0.04) | 0.69<br>(+0.04) | 0.86<br>(-0.02) | 0.90<br>(-0.04) | 1.01<br>(-0.11) | 5.08<br>(+0.36) |
| rhai     | 0.53<br>(+0.05) | 0.72<br>(+0.03) | 0.68<br>(+0.03) | 0.92<br>(+0.04) | 0.98<br>(+0.04) | 1.16<br>(+0.04) | 5.11<br>(+0.39) |

## Prerequisites

You will need to have installed:

- [Vegeta](https://github.com/tsenart/vegeta)
- [Task](https://github.com/go-task/task) (for `Taskfile` support)

_Note: `go-task` can be installed via `brew`._

Next, you'll also need an Apollo Graph Reference and Apollo Key. For the testing, we are using a local supergraph (located at `./router/supergraph.graphql`), but [the Coprocessor feature is restricted to enterprise customers only](https://www.apollographql.com/docs/router/customizations/coprocessor).

## Usage

Once you have the necessary requirements:

- Copy the `.sample_env` file to `.env` and fill in the fields
- Run `task test-all` to run the available tests within the project.

### Note

During the development of this project, it was discovered that running these tests on MacOS may result in inconsistent results. We strongly recommend running these tests on a Windows machine if possible to ensure the results are consistent from run to run. 

## Contributing

### Coprocessor

To add new coprocessors, you will need to:

- Add a new folder to the [coprocessors](./coprocessors/)
- Write the coprocessor to use the three static endpoints. Refer to [the Go implementation](./coprocessors/go/main.go) for more details:
  - `/static-subgraph`
  - `/guid-response`
  - `/client-awareness`
- Add a Dockerfile to build and host the image
- Update the [Taskfile.Test.yml](./Taskfile.Test.yml) to run the new coprocessor and report on it
- Add coprocessor to test tasks in [Taskfile.yml](./Taskfile.yml) (i.e. under `tasks.static.cmds`)

### Tests

To create new tests:

- Determine what you would like to benchmark against (Rhai, Config, and/or Coprocessors)
- Implement the test within all coprocessors and related extension points
- Following the format of the [`static-subgraph`](./tests/static-subgraph/) folder, create a new folder for the test and associated Router configurations
- Create a new test setup under `includes` in [Taskfile.yml](./Taskfile.yml) follow the pattern of `includes.static`
- Create a new test task in [Taskfile.yml](./Taskfile.yml) follow the pattern of `tasks.static`

See current tests for reference.

## Licensing

Source code in this repository is covered by the Elastic License 2.0. The
default throughout the repository is a license under the Elastic License 2.0,
unless a file header or a license file in a subdirectory specifies another
license. [See the LICENSE](./LICENSE) for the full license text.
