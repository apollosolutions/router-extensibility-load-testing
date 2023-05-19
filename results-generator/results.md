### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a Coprocessor.

| Type     | Min    | Mean    | p50    | p90     | p95      | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | -------- | -------- | -------- |
| baseline | 2.88ms | 10.87ms | 4.59ms | 10.45ms | 37.73ms  | 162.75ms | 261.58ms |
| go       | 0.39ms | 2.57ms  | 1.56ms | 5.10ms  | 7.87ms   | 15.75ms  | 44.53ms  |
| node     | 0.37ms | 2.14ms  | 1.44ms | 4.21ms  | 6.32ms   | 11.53ms  | 24.87ms  |
| rhai     | 3.59ms | 21.45ms | 5.42ms | 47.71ms | 150.16ms | 254.28ms | 311.06ms |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min    | Mean   | p50    | p90    | p95     | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------- | ------- | -------- |
| baseline | 3.54ms | 8.18ms | 5.10ms | 8.53ms | 21.68ms | 80.60ms | 178.48ms |
| go       | 0.27ms | 1.98ms | 1.34ms | 3.66ms | 5.44ms  | 10.51ms | 23.33ms  |
| node     | 0.48ms | 2.03ms | 1.31ms | 3.78ms | 5.58ms  | 12.69ms | 46.27ms  |

### Using local go subgraphs

This allows you to not have to connect to the open internet to talk to a subgraph

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max      |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | -------- |
| baseline | 2.83ms | 6.03ms | 4.28ms | 6.40ms | 7.38ms | 63.91ms | 276.56ms |
| local    | 2.73ms | 4.85ms | 4.05ms | 6.12ms | 7.45ms | 13.75ms | 148.52ms |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 3.12ms | 22.93ms | 4.84ms | 16.81ms | 91.02ms | 481.46ms | 622.79ms |
| config   | 3.03ms | 10.68ms | 5.15ms | 12.88ms | 43.32ms | 121.11ms | 281.80ms |
| go       | 0.38ms | 1.89ms  | 1.39ms | 3.04ms  | 4.28ms  | 11.52ms  | 23.25ms  |
| node     | 0.47ms | 1.82ms  | 1.39ms | 2.84ms  | 4.00ms  | 9.25ms   | 24.81ms  |
| rhai     | 3.24ms | 12.07ms | 5.16ms | 8.95ms  | 30.47ms | 208.74ms | 335.60ms |

