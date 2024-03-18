## test data on faas cloud


## Vercel
pnpm i -D @sveltejs/adapter-vercel

bun -v

curl -fsSL https://bun.sh/install | bash

can not visit in china , need 科学上网,使用科学上网后测试的数据

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark-nf5uu0c60-warlbor.vercel.app
Running 10s test @ https://ssr-svelte-benchmark-nf5uu0c60-warlbor.vercel.app
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   105.03ms  131.03ms   1.42s    93.83%
    Req/Sec   354.65    149.30     0.96k    69.44%
  35899 requests in 10.10s, 464.40MB read
  Non-2xx or 3xx responses: 35899
Requests/sec:   3555.25
Transfer/sec:     45.99MB


wrk -t12 -c400 -d10s https://ssr-svelte-benchmark-nf5uu0c60-warlbor.vercel.app
Running 10s test @ https://ssr-svelte-benchmark-nf5uu0c60-warlbor.vercel.app
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   105.29ms  136.76ms   1.54s    94.13%
    Req/Sec   351.36    162.55   818.00     69.32%
  35185 requests in 10.12s, 420.12MB read
  Socket errors: connect 125, read 265, write 0, timeout 0
  Non-2xx or 3xx responses: 35185
Requests/sec:   3475.50
Transfer/sec:     41.50MB


## Netlify edge mode

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.netlify.app/
Running 10s test @ https://ssr-svelte-benchmark.netlify.app/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   752.49ms  669.20ms   1.99s    50.14%
    Req/Sec    40.97     44.63   205.00     86.61%
  796 requests in 10.10s, 73.22MB read
  Socket errors: connect 652, read 836, write 0, timeout 68
  Non-2xx or 3xx responses: 284
Requests/sec:     78.82
Transfer/sec:      7.25MB


wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.netlify.app/
Running 10s test @ https://ssr-svelte-benchmark.netlify.app/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   750.98ms  660.58ms   2.00s    48.00%
    Req/Sec    44.49     42.38   202.00     76.52%
  848 requests in 10.10s, 80.83MB read
  Socket errors: connect 860, read 414, write 0, timeout 73
  Non-2xx or 3xx responses: 307
Requests/sec:     84.00
Transfer/sec:      8.01MB


## Netlify disable edge mode 特别容易访问限制

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.netlify.app/
Running 10s test @ https://ssr-svelte-benchmark.netlify.app/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   608.65ms  338.17ms   1.98s    82.20%
    Req/Sec    45.82     30.64   173.00     69.25%
  2356 requests in 10.10s, 321.46MB read
  Socket errors: connect 1792, read 476, write 0, timeout 13
Requests/sec:    233.38
Transfer/sec:     31.84MB

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.netlify.app
Running 10s test @ https://ssr-svelte-benchmark.netlify.app
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.12s   404.78ms   1.98s    65.24%
    Req/Sec    18.30     16.52    70.00     75.65%
  518 requests in 10.09s, 81.56MB read
  Socket errors: connect 2164, read 348, write 0, timeout 3
Requests/sec:     51.36
Transfer/sec:      8.09MB

 wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.netlify.app
Running 10s test @ https://ssr-svelte-benchmark.netlify.app
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.00us    0.00us   0.00us     nan%
    Req/Sec     0.00      0.00     0.00       nan%
  0 requests in 10.10s, 20.74KB read
  Socket errors: connect 158, read 120, write 0, timeout 0
Requests/sec:      0.00
Transfer/sec:      2.05KB

