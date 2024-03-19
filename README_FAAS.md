
#### Vercel

需要科学上网

https://ssr-svelte-benchmark.vercel.app/

    import adapter from '@sveltejs/adapter-vercel';

#### Netlify

    import adapterNetlify from "@sveltejs/adapter-netlify";
    const adapterNetlifyDisableEdge = adapterNetlify({
        edge: false,
        split: false,
    });
    const adapterNetlifyEdge = adapterNetlify({
        edge: true,
    });

    pnpm run build

#### Cloudfare Pages


    pnpm i -D @sveltejs/adapter-cloudflare

免费版的 cloudare pages都有 限制一天
Requests today 61,694 / 100,000

#### Cloudfare Workers

    pnpm i -D @sveltejs/adapter-cloudflare-workers

    pnpm i -g wrangler
    wrangler login
    wrangler deploy
    
https://ssr-svelte-benchmark.gakaki.workers.dev/
需要科学上网 国内访问不到 速度和Pages差不多
部署需要安装命令行wangle 稍微麻烦

wangle login
wangle deploy

#### Deno Deploy(待续)

https://dbushell.com/2023/06/26/sveltekit-oauth-deno-deploy/

#### WinterJS (Warmer) 待续

给定Demo是静态文件生成,非SSR,无可比性.

https://gakaki-ssr-svelte-benchmark.wasmer.app/

#### AWS Lambda(待续)

本地运行报错 看起来要解决rollup编译为ES2000的问题
llrt build/index.js
SyntaxError: unexpected 'await' keyword
at build/handler.js:1167:0



#### AWS Lambda(待续) LLRT

#### 阿里云函数计算(待续)


