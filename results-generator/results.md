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

