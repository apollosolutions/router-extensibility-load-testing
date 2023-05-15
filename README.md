# Router Extensibility Load Testing

**The code in this repository is experimental and has been provided for reference purposes only. Community feedback is welcome but this project may not be supported in the same way that repositories in the official [Apollo GraphQL GitHub organization](https://github.com/apollographql) are. If you need help you can file an issue on this repository, [contact Apollo](https://www.apollographql.com/contact-sales) to talk to an expert, or create a ticket directly in Apollo Studio.**

> Note: The Apollo Router is made available under the Elastic License v2.0 (ELv2).
> Read [our licensing page](https://www.apollographql.com/docs/resources/elastic-license-v2-faq/) for more details.

## Overview

This repository is a simple way to test the overhead of the three customization points of the Apollo Router:

* [Coprocessors](https://www.apollographql.com/docs/router/customizations/coprocessor)
* [Rhai](https://www.apollographql.com/docs/router/customizations/rhai)
* Configuration options

The current tests are:

* Setting a static header to subgraphs (Config, Rhai, Coprocessor)
* Setting 10 GUID headers on response to clients (Rhai, Coprocessor)
* JWT-based client awareness (Coprocessor)

The coprocessors are currently written in: 
* [Go](./coprocessors/go/)
* [Node](./coprocessors/node)

With more to come.

## Results

For the below tables, each section corresponds to the related test name. Each type relates to either the baseline (meaning no Router configuration), or the extensibility option. Languages imply a coprocessor. 

The tests were run at 100 requests per second for 30 seconds against an Apollo Router version 1.18.0. 

To help with consistency, there are resource limits for both the router and the coprocessors when using Docker. 

* The router is configured with .33 CPU cycles and 1GB RAM.
* Coprocessors are configured with .25 CPU cycles and 1GB RAM. 

### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a Coprocessor.

| Type     | Min    | Mean   | p50    | p90    | p95    | p99      | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | -------- | -------- |
| baseline | 1.32ms | 4.21ms | 3.33ms | 6.44ms | 8.11ms | 19.57ms  | 66.23ms  |
| go       | 1.99ms | 6.19ms | 4.22ms | 7.15ms | 8.42ms | 51.76ms  | 316.66ms |
| node     | 2.31ms | 8.66ms | 4.34ms | 6.93ms | 8.35ms | 197.68ms | 482.84ms |
| rhai     | 1.34ms | 4.22ms | 3.25ms | 6.13ms | 7.79ms | 17.14ms  | 155.39ms |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.30ms | 5.93ms  | 4.17ms | 8.50ms  | 12.61ms | 39.88ms  | 162.83ms |
| go       | 2.20ms | 10.61ms | 6.67ms | 14.72ms | 21.08ms | 107.63ms | 354.13ms |
| node     | 2.76ms | 8.64ms  | 6.40ms | 11.20ms | 13.71ms | 75.25ms  | 237.10ms |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | -------- |
| baseline | 1.35ms | 7.31ms | 3.94ms | 7.09ms | 9.14ms | 99.44ms | 510.60ms |
| config   | 1.37ms | 4.35ms | 3.49ms | 6.37ms | 7.60ms | 19.20ms | 123.73ms |
| go       | 1.87ms | 5.26ms | 4.33ms | 6.84ms | 8.39ms | 22.69ms | 176.16ms |
| node     | 2.02ms | 6.90ms | 4.80ms | 7.39ms | 9.99ms | 81.72ms | 272.95ms |
| rhai     | 1.32ms | 5.03ms | 3.29ms | 6.14ms | 7.66ms | 37.28ms | 305.04ms |

## Prerequisites

You will need to have installed:

* [Vegeta](https://github.com/tsenart/vegeta)
* [Task](https://github.com/go-task/task) (for `Taskfile` support)
* A copy of the [Retail Supergraph demo](https://github.com/apollosolutions/retail-supergraph) running on port 4001

_Note: `vegeta` and `go-task` can both can be installed via `brew`._

Next, you'll also need an Apollo Graph Reference and Apollo Key. For the testing, we are using a local supergraph (located at `./router/supergraph.graphql`), but [the Coprocessor feature is restricted to enterprise customers only](https://www.apollographql.com/docs/router/customizations/coprocessor).

## Usage

Once you have the necessary requirements:

* Copy the `.sample_env` file to `.env` and fill in the fields
* Run `task test-all` to run the available tests within the project.

## Contributing

### Coprocessor

To add new coprocessors, you will need to:
- Add a new folder to the [coprocessors](./coprocessors/)
- Write the coprocessor to use the three static endpoints. Refer to [the Go implementation](./coprocessors/go/main.go) for more details:
  - `/static-subgraph`
  - `/guid-response`
  - `/client-awareness`
- Add a Dockerfile to build and host the image
- Add a new `setup-` and `cleanup-` command within [`Taskfile.Shared.yml`](./Taskfile.Shared.yml)
- Update the [Taskfile.Test.yml](./Taskfile.Test.yml) to run the new coprocessor and report on it

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
