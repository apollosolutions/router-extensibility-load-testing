### guid-response

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.36ms | 10.40ms | 5.03ms | 11.41ms | 16.88ms | 163.76ms | 656.39ms |
| go       | 2.10ms | 9.81ms  | 5.58ms | 13.72ms | 18.30ms | 120.42ms | 398.76ms |
| rhai     | 1.34ms | 7.56ms  | 3.70ms | 7.57ms  | 12.36ms | 122.52ms | 455.06ms |

### jwt-client-awareness

| Type     | Min    | Mean   | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------ | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.32ms | 9.15ms | 4.75ms | 12.43ms | 16.89ms | 133.67ms | 439.93ms |
| go       | 2.01ms | 8.56ms | 5.76ms | 12.21ms | 15.61ms | 48.62ms  | 300.25ms |

### static-subgraph

| Type     | Min    | Mean    | p50    | p90     | p95     | p99      | Max      |
| -------- | ------ | ------- | ------ | ------- | ------- | -------- | -------- |
| baseline | 1.45ms | 8.32ms  | 5.34ms | 11.95ms | 17.38ms | 68.28ms  | 305.95ms |
| config   | 1.23ms | 7.43ms  | 4.82ms | 8.59ms  | 10.66ms | 78.89ms  | 380.00ms |
| go       | 2.00ms | 10.70ms | 5.70ms | 11.63ms | 16.92ms | 165.81ms | 504.05ms |
| rhai     | 1.37ms | 6.77ms  | 3.62ms | 8.31ms  | 11.78ms | 81.23ms  | 426.96ms |
