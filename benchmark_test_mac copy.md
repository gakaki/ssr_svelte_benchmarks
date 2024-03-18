## test data on faas cloud

## Cloudflare Workers
有cron 有kv 
wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.gakaki.workers.dev/

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.gakaki.workers.dev/
Running 10s test @ https://ssr-svelte-benchmark.gakaki.workers.dev/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   580.46ms  344.36ms   2.00s    72.37%
    Req/Sec    58.54     33.32   166.00     65.78%
  5087 requests in 10.10s, 704.90MB read
  Socket errors: connect 0, read 0, write 0, timeout 187
Requests/sec:    503.79
Transfer/sec:     69.81MB

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.gakaki.workers.dev/
Running 10s test @ https://ssr-svelte-benchmark.gakaki.workers.dev/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   542.12ms  356.47ms   2.00s    73.91%
    Req/Sec    63.42     34.27   181.00     67.16%
  5573 requests in 10.09s, 775.58MB read
  Socket errors: connect 0, read 0, write 0, timeout 160
Requests/sec:    552.37
Transfer/sec:     76.87MB

wrk -t12 -c400 -d10s https://ssr-svelte-benchmark.gakaki.workers.dev/
Running 10s test @ https://ssr-svelte-benchmark.gakaki.workers.dev/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   617.83ms  353.64ms   1.99s    72.27%
    Req/Sec    55.13     33.44   163.00     66.51%
  4741 requests in 10.09s, 662.85MB read
  Socket errors: connect 0, read 0, write 0, timeout 209
Requests/sec:    470.01
Transfer/sec:     65.71MB

## Cloudflare Pages 无论是外网服务器还是本地都能访问到 就是部署之后要等上5分钟 

Requests today
32,027 / 100,000

可以连接github的repo


wrk -t12 -c400 -d10s https://1ba3fd77.ssr-svelte-benchmark.pages.dev
Running 10s test @ https://1ba3fd77.ssr-svelte-benchmark.pages.dev
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   554.68ms  358.75ms   2.00s    71.59%
    Req/Sec    60.03     34.78   202.00     66.97%
  5129 requests in 10.08s, 714.48MB read
  Socket errors: connect 0, read 0, write 0, timeout 237
Requests/sec:    508.63
Transfer/sec:     70.85MB

 wrk -t12 -c400 -d10s https://1ba3fd77.ssr-svelte-benchmark.pages.dev
Running 10s test @ https://1ba3fd77.ssr-svelte-benchmark.pages.dev
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   576.83ms  358.59ms   2.00s    70.78%
    Req/Sec    57.96     31.41   190.00     66.46%
  4934 requests in 10.08s, 690.09MB read
  Socket errors: connect 0, read 0, write 0, timeout 211
Requests/sec:    489.54
Transfer/sec:     68.47MB


Running 10s test @ https://1ba3fd77.ssr-svelte-benchmark.pages.dev
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   555.81ms  357.65ms   1.99s    70.76%
    Req/Sec    62.84     34.17   194.00     68.45%
  5860 requests in 10.09s, 810.87MB read
  Socket errors: connect 0, read 0, write 0, timeout 73
Requests/sec:    580.79
Transfer/sec:     80.37MB


wrk -t12 -c400 -d10s https://1ba3fd77.ssr-svelte-benchmark.pages.dev

Running 10s test @ https://1ba3fd77.ssr-svelte-benchmark.pages.dev
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   489.10ms  217.70ms   1.87s    79.29%
    Req/Sec    44.52     31.65   191.00     63.27%
  4434 requests in 10.08s, 601.97MB read
  Socket errors: connect 158, read 0, write 0, timeout 12
Requests/sec:    439.97
Transfer/sec:     59.73MB

wrk -t12 -c400 -d10s https://1ba3fd77.ssr-svelte-benchmark.pages.dev

Running 10s test @ https://1ba3fd77.ssr-svelte-benchmark.pages.dev
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   497.88ms  242.09ms   1.93s    75.65%
    Req/Sec    42.12     33.50   202.00     79.35%
  4218 requests in 10.10s, 579.00MB read
  Socket errors: connect 158, read 0, write 0, timeout 8
Requests/sec:    417.80
Transfer/sec:     57.35MB

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

