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

