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

| Type     | Min (ms)        | Mean (ms)        | p50 (ms)        | p90 (ms)         | p95 (ms)         | p99 (ms)            | Max (ms)            |
| -------- | --------------- | ---------------- | --------------- | ---------------- | ---------------- | ------------------- | ------------------- |
| baseline | 1.31            | 5.29             | 4.57            | 8.27             | 10.01            | 16.68               | 91.33               |
| csharp   | 2.05<br>(+0.74) | 6.63<br>(+1.34)  | 4.92<br>(+0.35) | 8.92<br>(+0.65)  | 12.22<br>(+2.21) | 36.09<br>(+19.41)   | 212.01<br>(+120.68) |
| go       | 1.87<br>(+0.56) | 6.59<br>(+1.30)  | 5.73<br>(+1.16) | 10.22<br>(+1.95) | 12.27<br>(+2.26) | 20.38<br>(+3.70)    | 92.06<br>(+0.73)    |
| java     | 2.82<br>(+1.51) | 13.37<br>(+8.08) | 4.71<br>(+0.14) | 7.44<br>(-0.83)  | 12.28<br>(+2.27) | 398.98<br>(+382.30) | 692.91<br>(+601.58) |
| node     | 2.24<br>(+0.93) | 6.92<br>(+1.63)  | 5.44<br>(+0.87) | 10.04<br>(+1.77) | 13.39<br>(+3.38) | 37.34<br>(+20.66)   | 150.21<br>(+58.88)  |
| rhai     | 1.44<br>(+0.13) | 5.34<br>(+0.05)  | 4.57<br>(0.00)  | 8.32<br>(+0.05)  | 10.29<br>(+0.28) | 19.78<br>(+3.10)    | 89.14<br>(-2.19)    |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min (ms)        | Mean (ms)         | p50 (ms)        | p90 (ms)        | p95 (ms)          | p99 (ms)            | Max (ms)              |
| -------- | --------------- | ----------------- | --------------- | --------------- | ----------------- | ------------------- | --------------------- |
| baseline | 1.35            | 4.53              | 3.78            | 6.69            | 7.91              | 18.88               | 76.30                 |
| csharp   | 1.49<br>(+0.14) | 5.33<br>(+0.80)   | 3.15<br>(-0.63) | 4.95<br>(-1.74) | 6.29<br>(-1.62)   | 50.87<br>(+31.99)   | 332.08<br>(+255.78)   |
| go       | 2.01<br>(+0.66) | 5.05<br>(+0.52)   | 4.25<br>(+0.47) | 7.29<br>(+0.60) | 9.34<br>(+1.43)   | 19.33<br>(+0.45)    | 66.79<br>(-9.51)      |
| java     | 2.52<br>(+1.17) | 33.58<br>(+29.05) | 5.34<br>(+1.56) | 9.56<br>(+2.87) | 35.80<br>(+27.89) | 969.13<br>(+950.25) | 1365.42<br>(+1289.12) |
| node     | 2.74<br>(+1.39) | 7.02<br>(+2.49)   | 5.78<br>(+2.00) | 9.99<br>(+3.30) | 13.04<br>(+5.13)  | 35.70<br>(+16.82)   | 108.69<br>(+32.39)    |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min (ms)        | Mean (ms)        | p50 (ms)        | p90 (ms)         | p95 (ms)          | p99 (ms)            | Max (ms)            |
| -------- | --------------- | ---------------- | --------------- | ---------------- | ----------------- | ------------------- | ------------------- |
| baseline | 1.31            | 4.85             | 4.05            | 7.30             | 8.55              | 17.51               | 83.64               |
| config   | 1.36<br>(+0.05) | 4.83<br>(-0.02)  | 4.30<br>(+0.25) | 7.36<br>(+0.06)  | 8.23<br>(-0.32)   | 16.44<br>(-1.07)    | 65.63<br>(-18.01)   |
| csharp   | 1.97<br>(+0.66) | 7.83<br>(+2.98)  | 6.26<br>(+2.21) | 11.73<br>(+4.43) | 15.04<br>(+6.49)  | 37.98<br>(+20.47)   | 206.87<br>(+123.23) |
| go       | 1.86<br>(+0.55) | 5.72<br>(+0.87)  | 5.34<br>(+1.29) | 8.43<br>(+1.13)  | 9.58<br>(+1.03)   | 16.22<br>(-1.29)    | 80.92<br>(-2.72)    |
| java     | 2.15<br>(+0.84) | 14.65<br>(+9.80) | 6.26<br>(+2.21) | 12.12<br>(+4.82) | 19.40<br>(+10.85) | 365.91<br>(+348.40) | 652.42<br>(+568.78) |
| node     | 2.04<br>(+0.73) | 6.53<br>(+1.68)  | 5.90<br>(+1.85) | 9.86<br>(+2.56)  | 12.63<br>(+4.08)  | 24.27<br>(+6.76)    | 79.73<br>(-3.91)    |
| rhai     | 1.34<br>(+0.03) | 4.95<br>(+0.10)  | 4.42<br>(+0.37) | 7.40<br>(+0.10)  | 8.20<br>(-0.35)   | 13.81<br>(-3.70)    | 119.59<br>(+35.95)  |


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
