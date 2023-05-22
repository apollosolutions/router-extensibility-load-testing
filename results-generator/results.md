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

