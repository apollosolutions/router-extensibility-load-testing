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

With more to come?

## Results

For the below tables, each section corresponds to the related test name. Each type relates to either the baseline (meaning no Router configuration), or the extensibility option. Languages imply a coprocessor.

The tests were run at 100 requests per second for 60 seconds against an Apollo Router version 1.19.0.

To help with consistency, there are resource limits for both the router and the coprocessors when using Docker--currently 1 CPU core and 1GB of RAM

### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a Coprocessor.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)        | Max (ms)          |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- | ----------------- |
| baseline | 2.11            | 4.25            | 4.21            | 5.00            | 5.39            | 6.84            | 30.34             |
| csharp   | 2.67<br>(+0.56) | 4.50<br>(+0.25) | 4.47<br>(+0.26) | 5.25<br>(+0.25) | 5.58<br>(+0.19) | 6.89<br>(+0.05) | 40.71<br>(+10.37) |
| go       | 2.52<br>(+0.41) | 4.62<br>(+0.37) | 4.48<br>(+0.27) | 5.44<br>(+0.44) | 5.95<br>(+0.56) | 7.62<br>(+0.78) | 45.12<br>(+14.78) |
| java     | 2.74<br>(+0.63) | 4.19<br>(-0.06) | 3.88<br>(-0.33) | 5.03<br>(+0.03) | 5.55<br>(+0.16) | 8.58<br>(+1.74) | 52.14<br>(+21.80) |
| node     | 2.78<br>(+0.67) | 4.57<br>(+0.32) | 4.48<br>(+0.27) | 5.35<br>(+0.35) | 5.75<br>(+0.36) | 7.59<br>(+0.75) | 38.48<br>(+8.14)  |
| python   | 2.81<br>(+0.70) | 4.64<br>(+0.39) | 4.58<br>(+0.37) | 5.51<br>(+0.51) | 5.97<br>(+0.58) | 8.07<br>(+1.23) | 28.23<br>(-2.11)  |
| rhai     | 2.04<br>(-0.07) | 4.32<br>(+0.07) | 4.24<br>(+0.03) | 5.12<br>(+0.12) | 5.58<br>(+0.19) | 7.32<br>(+0.48) | 34.93<br>(+4.59)  |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)        | Max (ms)          |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | --------------- | ----------------- |
| baseline | 2.03            | 4.29            | 4.25            | 5.01            | 5.36            | 6.10            | 31.34             |
| csharp   | 2.86<br>(+0.83) | 4.51<br>(+0.22) | 4.41<br>(+0.16) | 5.39<br>(+0.38) | 5.78<br>(+0.42) | 7.25<br>(+1.15) | 19.18<br>(-12.16) |
| go       | 2.44<br>(+0.41) | 4.62<br>(+0.33) | 4.46<br>(+0.21) | 5.54<br>(+0.53) | 6.16<br>(+0.80) | 9.09<br>(+2.99) | 38.66<br>(+7.32)  |
| java     | 2.95<br>(+0.92) | 4.55<br>(+0.26) | 4.18<br>(-0.07) | 5.51<br>(+0.50) | 6.04<br>(+0.68) | 8.64<br>(+2.54) | 49.45<br>(+18.11) |
| node     | 3.31<br>(+1.28) | 4.63<br>(+0.34) | 4.52<br>(+0.27) | 5.33<br>(+0.32) | 5.65<br>(+0.29) | 6.44<br>(+0.34) | 35.20<br>(+3.86)  |
| python   | 2.86<br>(+0.83) | 4.59<br>(+0.30) | 4.53<br>(+0.28) | 5.37<br>(+0.36) | 5.76<br>(+0.40) | 7.57<br>(+1.47) | 36.77<br>(+5.43)  |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min (ms)        | Mean (ms)       | p50 (ms)        | p90 (ms)        | p95 (ms)        | p99 (ms)          | Max (ms)          |
| -------- | --------------- | --------------- | --------------- | --------------- | --------------- | ----------------- | ----------------- |
| baseline | 2.08            | 4.30            | 4.25            | 5.04            | 5.36            | 6.50              | 28.74             |
| config   | 1.99<br>(-0.09) | 4.24<br>(-0.06) | 4.20<br>(-0.05) | 5.00<br>(-0.04) | 5.34<br>(-0.02) | 6.14<br>(-0.36)   | 30.11<br>(+1.37)  |
| csharp   | 3.14<br>(+1.06) | 4.66<br>(+0.36) | 4.51<br>(+0.26) | 5.41<br>(+0.37) | 5.77<br>(+0.41) | 7.17<br>(+0.67)   | 38.30<br>(+9.56)  |
| go       | 3.07<br>(+0.99) | 4.58<br>(+0.28) | 4.43<br>(+0.18) | 5.45<br>(+0.41) | 5.70<br>(+0.34) | 6.80<br>(+0.30)   | 36.33<br>(+7.59)  |
| java     | 3.20<br>(+1.12) | 4.94<br>(+0.64) | 4.41<br>(+0.16) | 5.77<br>(+0.73) | 6.44<br>(+1.08) | 16.76<br>(+10.26) | 66.46<br>(+37.72) |
| node     | 3.31<br>(+1.23) | 4.61<br>(+0.31) | 4.44<br>(+0.19) | 5.50<br>(+0.46) | 5.85<br>(+0.49) | 6.73<br>(+0.23)   | 41.97<br>(+13.23) |
| python   | 1.69<br>(-0.39) | 4.00<br>(-0.30) | 3.90<br>(-0.35) | 4.63<br>(-0.41) | 5.94<br>(+0.58) | 8.10<br>(+1.60)   | 23.28<br>(-5.46)  |
| rhai     | 1.99<br>(-0.09) | 4.25<br>(-0.05) | 4.20<br>(-0.05) | 5.07<br>(+0.03) | 5.48<br>(+0.12) | 7.65<br>(+1.15)   | 25.71<br>(-3.03)  |

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
