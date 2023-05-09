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

With more to come.

## Results

For the below tables, each section corresponds to the related test name. Each type relates to either the baseline (meaning no Router configuration), or the extensibility option. Languages imply a coprocessor. 

The tests were run at 100 requests per second for 30 seconds against an Apollo Router version 1.18.0. 

To help with consistency, there are resource limits for both the router and the coprocessors when using Docker. 

* The router is configured with .33 CPU cycles and 1GB RAM.
* Coprocessors are configured with .25 CPU cycles and 1GB RAM. 


### guid-response

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | -------- |
| baseline | 1.99ms | 4.12ms | 2.61ms | 3.70ms | 4.74ms | 42.72ms | 283.38ms |
| go       | 2.76ms | 4.94ms | 3.51ms | 4.77ms | 5.95ms | 32.19ms | 272.03ms |
| rhai     | 2.27ms | 3.91ms | 2.96ms | 4.20ms | 5.42ms | 31.00ms | 135.80ms |

### jwt-client-awareness

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | -------- |
| baseline | 2.14ms | 4.21ms | 3.08ms | 4.58ms | 6.04ms | 33.27ms | 178.97ms |
| go       | 2.79ms | 4.89ms | 3.64ms | 4.96ms | 6.11ms | 35.71ms | 244.76ms |

### static-subgraph

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | -------- |
| baseline | 2.09ms | 3.97ms | 3.04ms | 4.61ms | 6.00ms | 26.56ms | 92.05ms  |
| config   | 2.02ms | 3.40ms | 2.64ms | 3.58ms | 4.48ms | 18.17ms | 132.82ms |
| go       | 2.71ms | 4.34ms | 3.52ms | 4.82ms | 5.84ms | 22.48ms | 95.44ms  |
| rhai     | 2.15ms | 3.62ms | 2.91ms | 4.39ms | 5.92ms | 17.12ms | 70.93ms  |


## Prerequisites

You will need to have installed:

* [Vegeta](https://github.com/tsenart/vegeta)
* [Task (for Taskfile support)](https://github.com/go-task/task)
* [A copy of the Retail Supergraph demo and it running](https://github.com/apollosolutions/retail-supergraph)

Both can be installed by `brew`.

Next, you'll also need an Apollo Graph Reference and Apollo Key. For the testing, we using a local supergraph (located at `./router/supergraph.graphql`), but [the Coprocessor feature is restricted to enterprise customers only](https://www.apollographql.com/docs/router/customizations/coprocessor). 

## Usage

Once you have the necessary requirements: 

* Copy the `.sample_env` file to `.env` and fill in the fields
* Run `task test-all` to run the available tests within the project. 

## Contributing

### Coprocessor

To add new coprocessors, you will need to:
- Add a new folder in the [coprocessors folder](./coprocessors/)
- Write the coprocessor to use the three static endpoints. Refer to [the Go implementation for more details](./coprocessors/go/main.go): 
  - `/static-subgraph`
  - `guid-response`
  - `client-awareness`
- Add a Dockerfile to build and host the image
- Add a new `setup-` and `cleanup-` command within [`Shared.yml`](./Shared.yml)
- Update the test files [(e.g. ClientAwareness.yml)](./ClientAwareness.yml) to run the new test and report on it

### Tests

To create new tests: 

- Determine what you would like to benchmark against (Rhai, Config, and/or Coprocessors)
- Implement the test within all coprocessors and related extension points 
- Following the format of the [`static-subgraph`](./tests/static-subgraph/) folder, create a new folder for the test and associated Router configurations
- Create a new Taskfile, similar to [ClientAwareness.yml](./ClientAwareness.yml)

See current tests for reference.

## Licensing

Source code in this repository is covered by the Elastic License 2.0. The
default throughout the repository is a license under the Elastic License 2.0,
unless a file header or a license file in a subdirectory specifies another
license. [See the LICENSE](./LICENSE) for the full license text.