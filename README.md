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
* [C#](./coprocessors/csharp)
* [Java](./coprocessors/java)

With more to come.

## Results

For the below tables, each section corresponds to the related test name. Each type relates to either the baseline (meaning no Router configuration), or the extensibility option. Languages imply a coprocessor.

The tests were run at 100 requests per second for 30 seconds against an Apollo Router version 1.19.0.

To help with consistency, there are resource limits for both the router and the coprocessors when using Docker--currently 1 CPU core and 1GB of RAM

### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a Coprocessor.

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.31ms | 5.29ms  | 4.57ms | 8.27ms  | 10.01ms | 16.68ms  | 91.33ms  |
| csharp   | 2.05ms | 6.63ms  | 4.92ms | 8.92ms  | 12.22ms | 36.09ms  | 212.01ms |
| go       | 1.87ms | 6.59ms  | 5.73ms | 10.22ms | 12.27ms | 20.38ms  | 92.06ms  |
| java     | 2.82ms | 13.37ms | 4.71ms | 7.44ms  | 12.28ms | 398.98ms | 692.91ms |
| node     | 2.24ms | 6.92ms  | 5.44ms | 10.04ms | 13.39ms | 37.34ms  | 150.21ms |
| rhai     | 1.44ms | 5.34ms  | 4.57ms | 8.32ms  | 10.29ms | 19.78ms  | 89.14ms  |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min    | Mean    | p50    | p90    | p95     | p99      | Max       |
| -------- | ------ | ------- | ------ | ------ | ------- | -------- | --------- |
| baseline | 1.35ms | 4.53ms  | 3.78ms | 6.69ms | 7.91ms  | 18.88ms  | 76.30ms   |
| csharp   | 1.49ms | 5.33ms  | 3.15ms | 4.95ms | 6.29ms  | 50.87ms  | 332.08ms  |
| go       | 2.01ms | 5.05ms  | 4.25ms | 7.29ms | 9.34ms  | 19.33ms  | 66.79ms   |
| java     | 2.52ms | 33.58ms | 5.34ms | 9.56ms | 35.80ms | 969.13ms | 1365.42ms |
| node     | 2.74ms | 7.02ms  | 5.78ms | 9.99ms | 13.04ms | 35.70ms  | 108.69ms  |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.31ms | 4.85ms  | 4.05ms | 7.30ms  | 8.55ms  | 17.51ms  | 83.64ms  |
| config   | 1.36ms | 4.83ms  | 4.30ms | 7.36ms  | 8.23ms  | 16.44ms  | 65.63ms  |
| csharp   | 1.97ms | 7.83ms  | 6.26ms | 11.73ms | 15.04ms | 37.98ms  | 206.87ms |
| go       | 1.86ms | 5.72ms  | 5.34ms | 8.43ms  | 9.58ms  | 16.22ms  | 80.92ms  |
| java     | 2.15ms | 14.65ms | 6.26ms | 12.12ms | 19.40ms | 365.91ms | 652.42ms |
| node     | 2.04ms | 6.53ms  | 5.90ms | 9.86ms  | 12.63ms | 24.27ms  | 79.73ms  |
| rhai     | 1.34ms | 4.95ms  | 4.42ms | 7.40ms  | 8.20ms  | 13.81ms  | 119.59ms |

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
