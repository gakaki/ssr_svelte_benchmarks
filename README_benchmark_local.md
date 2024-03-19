## Node 21.7.1
export NODE_ENV=production
node build/index.js

    wrk -t12 -c400 -d10s http://127.0.0.1:3000
    Running 10s test @ http://127.0.0.1:3000
      12 threads and 400 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency   151.36ms   92.26ms   1.99s    94.49%
        Req/Sec    39.79     22.79   181.00     71.01%
      4388 requests in 10.09s, 587.96MB read
      Socket errors: connect 155, read 500, write 2, timeout 52
    Requests/sec:    434.81
    Transfer/sec:     58.26MB

    wrk -t12 -c400 -d10s http://127.0.0.1:3000
    Running 10s test @ http://127.0.0.1:3000
      12 threads and 400 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency   146.49ms   92.85ms   1.94s    95.06%
        Req/Sec    39.13     25.22   212.00     74.82%
      4465 requests in 10.08s, 598.25MB read
      Socket errors: connect 155, read 458, write 4, timeout 55
    Requests/sec:    442.79
    Transfer/sec:     59.33MB

    wrk -t12 -c400 -d10s http://127.0.0.1:3000
    Running 10s test @ http://127.0.0.1:3000
      12 threads and 400 connections

      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency   145.81ms   90.60ms   1.95s    93.57%
        Req/Sec    43.62     29.48   220.00     75.32%
      4592 requests in 10.09s, 615.30MB read
      Socket errors: connect 155, read 413, write 2, timeout 54
    Requests/sec:    455.25
    Transfer/sec:     61.00MB




## Bun v1.0.33
bun -v

curl -fsSL https://bun.sh/install | bash

export NODE_ENV=production
bun build/index.js
wrk -t12 -c400 -d10s http://127.0.0.1:3000

Running 10s test @ http://127.0.0.1:3000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   672.37ms  172.87ms   1.49s    78.16%
    Req/Sec    30.99     18.12   120.00     70.35%
  3466 requests in 10.09s, 464.27MB read
  Socket errors: connect 155, read 239, write 0, timeout 0
Requests/sec:    343.43
Transfer/sec:     46.00MB

wrk -t12 -c400 -d10s http://127.0.0.1:3000
Running 10s test @ http://127.0.0.1:3000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   653.40ms  161.92ms   1.42s    84.46%
    Req/Sec    36.20     26.63   130.00     70.44%
  3564 requests in 10.10s, 477.39MB read
  Socket errors: connect 155, read 335, write 0, timeout 0
Requests/sec:    352.84
Transfer/sec:     47.26MB


wrk -t12 -c400 -d10s http://127.0.0.1:3000
Running 10s test @ http://127.0.0.1:3000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   686.59ms  180.70ms   1.60s    88.71%
    Req/Sec    31.89     23.32   140.00     79.55%
  3402 requests in 10.11s, 455.66MB read
  Socket errors: connect 155, read 238, write 0, timeout 0
Requests/sec:    336.66
Transfer/sec:     45.09MB

wrk -t12 -c400 -d10s http://127.0.0.1:3000
Running 10s test @ http://127.0.0.1:3000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   642.21ms  154.88ms   1.39s    85.38%
    Req/Sec    32.52     20.25   120.00     69.14%
  3638 requests in 10.09s, 487.30MB read
  Socket errors: connect 155, read 247, write 0, timeout 0
Requests/sec:    360.47
Transfer/sec:     48.28MB


## Deno could not issue 500
curl -fsSL https://deno.land/install.sh | sh // so slow need speed faster network
deno --version
deno 1.41.3 (release, aarch64-apple-darwin)
v8 12.3.219.9
typescript 5.3.3

no officel support on deno , 3rd party issue 500
https://github.com/dbushell/sveltekit-adapter-deno
pnpm install --save-dev sveltekit-adapter-deno

## WinterJS mac can not compole success ,colud not test
