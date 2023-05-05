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

TODO

## Prerequisites

You will need to have installed:

* [Vegeta](https://github.com/tsenart/vegeta)
* [Task (for Taskfile support)](https://github.com/go-task/task)

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