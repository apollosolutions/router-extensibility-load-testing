### GUID Response

This tests the overhead of setting 10 GUID headers on the response to the client using the `RouterResponse` stage. This is only available via Rhai or a Coprocessor.

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max     |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | ------- |
| baseline | 0.44ms | 2.09ms | 1.63ms | 3.60ms | 4.81ms | 8.17ms  | 20.09ms |
| csharp   | 0.51ms | 2.05ms | 1.51ms | 3.21ms | 4.85ms | 11.71ms | 26.33ms |
| go       | 0.53ms | 1.99ms | 1.51ms | 3.27ms | 4.38ms | 11.64ms | 45.55ms |
| java     | 0.47ms | 1.98ms | 1.53ms | 3.17ms | 4.33ms | 10.30ms | 22.81ms |
| node     | 0.53ms | 2.01ms | 1.48ms | 3.12ms | 4.70ms | 11.79ms | 32.48ms |
| rhai     | 0.45ms | 2.58ms | 1.68ms | 5.23ms | 7.10ms | 12.22ms | 30.52ms |

### Client Awareness using a JWT

This tests the overhead of validating a JWT, and using the JWT body to set the `apollographql-client-name` and `apollographql-client-version` headers. Those headers are then used for client identification within Apollo Studio.
This is only available via a coprocessor.

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max     |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | ------- |
| baseline | 0.55ms | 1.94ms | 1.52ms | 3.16ms | 4.15ms | 8.68ms  | 29.60ms |
| csharp   | 0.46ms | 2.27ms | 1.56ms | 4.57ms | 6.12ms | 9.72ms  | 22.30ms |
| go       | 0.46ms | 2.22ms | 1.52ms | 3.97ms | 7.45ms | 11.59ms | 24.99ms |
| java     | 0.44ms | 2.01ms | 1.53ms | 3.25ms | 4.32ms | 11.28ms | 24.06ms |
| node     | 0.51ms | 1.91ms | 1.50ms | 2.80ms | 3.76ms | 10.52ms | 24.27ms |

### Using local go subgraphs

This allows you to not have to connect to the open internet to talk to a subgraph

| Type     | Min    | Mean   | p50    | p90     | p95     | p99     | Max     |
| -------- | ------ | ------ | ------ | ------- | ------- | ------- | ------- |
|          | 1.12ms | 5.47ms | 4.27ms | 9.94ms  | 13.55ms | 21.02ms | 34.01ms |
| baseline | 1.05ms | 5.35ms | 4.26ms | 9.58ms  | 11.75ms | 18.64ms | 36.03ms |
| local    | 1.06ms | 5.41ms | 3.72ms | 10.36ms | 13.15ms | 21.20ms | 39.10ms |

### Static Subgraph Header

This tests the overhead of setting a static header to each subgraph request. The header is named `source` with a value matching the extensibility option. This is available via all three extensibility options.

| Type     | Min    | Mean   | p50    | p90    | p95    | p99     | Max     |
| -------- | ------ | ------ | ------ | ------ | ------ | ------- | ------- |
| baseline | 0.47ms | 2.06ms | 1.50ms | 3.38ms | 4.73ms | 11.27ms | 25.20ms |
| config   | 0.44ms | 1.96ms | 1.53ms | 3.14ms | 4.29ms | 9.59ms  | 16.18ms |
| csharp   | 0.34ms | 2.90ms | 1.70ms | 6.83ms | 8.60ms | 12.40ms | 37.88ms |
| go       | 0.31ms | 2.10ms | 1.53ms | 3.65ms | 4.95ms | 12.13ms | 31.34ms |
| java     | 0.39ms | 2.57ms | 1.49ms | 6.44ms | 8.22ms | 12.33ms | 26.43ms |
| node     | 0.42ms | 2.32ms | 1.49ms | 4.89ms | 6.46ms | 11.37ms | 23.17ms |
| rhai     | 0.52ms | 2.04ms | 1.47ms | 2.99ms | 5.38ms | 12.47ms | 32.19ms |

