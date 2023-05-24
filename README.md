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

| Type     | Min              | Mean              | p50              | p90               | p95               | p99                  | Max                  |
| -------- | ---------------- | ----------------- | ---------------- | ----------------- | ----------------- | -------------------- | -------------------- |
| baseline | 1.31ms           | 5.29ms            | 4.57ms           | 8.27ms            | 10.01ms           | 16.68ms              | 91.33ms              |
| csharp   | 2.05ms (+0.74ms) | 6.63ms (+1.34ms)  | 4.92ms (+0.35ms) | 8.92ms (+0.65ms)  | 12.22ms (+2.21ms) | 36.09ms (+19.41ms)   | 212.01ms (+120.68ms) |
| go       | 1.87ms (+0.56ms) | 6.59ms (+1.30ms)  | 5.73ms (+1.16ms) | 10.22ms (+1.95ms) | 12.27ms (+2.26ms) | 20.38ms (+3.70ms)    | 92.06ms (+0.73ms)    |
| java     | 2.82ms (+1.51ms) | 13.37ms (+8.08ms) | 4.71ms (+0.14ms) | 7.44ms (-0.83ms)  | 12.28ms (+2.27ms) | 398.98ms (+382.30ms) | 692.91ms (+601.58ms) |
| node     | 2.24ms (+0.93ms) | 6.92ms (+1.63ms)  | 5.44ms (+0.87ms) | 10.04ms (+1.77ms) | 13.39ms (+3.38ms) | 37.34ms (+20.66ms)   | 150.21ms (+58.88ms)  |
| rhai     | 1.44ms (+0.13ms) | 5.34ms (+0.05ms)  | 4.57ms (0.00ms)  | 8.32ms (+0.05ms)  | 10.29ms (+0.28ms) | 19.78ms (+3.10ms)    | 89.14ms (-2.19ms)    |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min              | Mean               | p50              | p90              | p95                | p99                  | Max                    |
| -------- | ---------------- | ------------------ | ---------------- | ---------------- | ------------------ | -------------------- | ---------------------- |
| baseline | 1.35ms           | 4.53ms             | 3.78ms           | 6.69ms           | 7.91ms             | 18.88ms              | 76.30ms                |
| csharp   | 1.49ms (+0.14ms) | 5.33ms (+0.80ms)   | 3.15ms (-0.63ms) | 4.95ms (-1.74ms) | 6.29ms (-1.62ms)   | 50.87ms (+31.99ms)   | 332.08ms (+255.78ms)   |
| go       | 2.01ms (+0.66ms) | 5.05ms (+0.52ms)   | 4.25ms (+0.47ms) | 7.29ms (+0.60ms) | 9.34ms (+1.43ms)   | 19.33ms (+0.45ms)    | 66.79ms (-9.51ms)      |
| java     | 2.52ms (+1.17ms) | 33.58ms (+29.05ms) | 5.34ms (+1.56ms) | 9.56ms (+2.87ms) | 35.80ms (+27.89ms) | 969.13ms (+950.25ms) | 1365.42ms (+1289.12ms) |
| node     | 2.74ms (+1.39ms) | 7.02ms (+2.49ms)   | 5.78ms (+2.00ms) | 9.99ms (+3.30ms) | 13.04ms (+5.13ms)  | 35.70ms (+16.82ms)   | 108.69ms (+32.39ms)    |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min              | Mean              | p50              | p90               | p95                | p99                  | Max                  |
| -------- | ---------------- | ----------------- | ---------------- | ----------------- | ------------------ | -------------------- | -------------------- |
| baseline | 1.31ms           | 4.85ms            | 4.05ms           | 7.30ms            | 8.55ms             | 17.51ms              | 83.64ms              |
| config   | 1.36ms (+0.05ms) | 4.83ms (-0.02ms)  | 4.30ms (+0.25ms) | 7.36ms (+0.06ms)  | 8.23ms (-0.32ms)   | 16.44ms (-1.07ms)    | 65.63ms (-18.01ms)   |
| csharp   | 1.97ms (+0.66ms) | 7.83ms (+2.98ms)  | 6.26ms (+2.21ms) | 11.73ms (+4.43ms) | 15.04ms (+6.49ms)  | 37.98ms (+20.47ms)   | 206.87ms (+123.23ms) |
| go       | 1.86ms (+0.55ms) | 5.72ms (+0.87ms)  | 5.34ms (+1.29ms) | 8.43ms (+1.13ms)  | 9.58ms (+1.03ms)   | 16.22ms (-1.29ms)    | 80.92ms (-2.72ms)    |
| java     | 2.15ms (+0.84ms) | 14.65ms (+9.80ms) | 6.26ms (+2.21ms) | 12.12ms (+4.82ms) | 19.40ms (+10.85ms) | 365.91ms (+348.40ms) | 652.42ms (+568.78ms) |
| node     | 2.04ms (+0.73ms) | 6.53ms (+1.68ms)  | 5.90ms (+1.85ms) | 9.86ms (+2.56ms)  | 12.63ms (+4.08ms)  | 24.27ms (+6.76ms)    | 79.73ms (-3.91ms)    |
| rhai     | 1.34ms (+0.03ms) | 4.95ms (+0.10ms)  | 4.42ms (+0.37ms) | 7.40ms (+0.10ms)  | 8.20ms (-0.35ms)   | 13.81ms (-3.70ms)    | 119.59ms (+35.95ms)  |


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
