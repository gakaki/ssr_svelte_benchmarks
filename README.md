## ssr_svelte_benchmarks

#### 目录结构

    ├── app.yaml                      # 云函数配置
    ├── crawler_wallstreetcn
    ├── layout.md
    ├── package.json
    ├── src
    │   ├── app.css
    │   ├── app.html
    │   ├── lib
    │   │   ├── redis.ts
    │   │   └── type.ts
    │   └── routes
    │       ├── +layout.svelte
    │       ├── +page.server.ts
    │       ├── +page.svelte
    ├── svelte.config.js
    ├── vite.config.js
    └── wasmer.toml

### 项目简介

本地与云端的serverless JS Runtime 环境 + Svelte SSR 的性能测试: Demo华尔街见闻首页(80%的页面内容)

### 数据来源

crawler_wallstreetcn 该目录为Golang编写的爬虫程序,华尔街见闻首页的数据API,包括最新文章,快讯、股票上证指数、大师课等.存储到Redis中用以缓存数据

### 本地测试环境 总结数据

wrk 压力测试工具 测试如下本地
MacBook Air M1 2020 16G 256G
Node.js、Deno、GraalVM_JS、Aws LLRT 、BunWasmer(WinterJS)

#### Node.js

#### Bun

    wrk -t12 -c400 -d10s http://127.0.0.1:3000

https://bun.sh/guides/ecosystem/sveltekit

    bun add -D svelte-adapter-bun

    <!-- import adapter from "@sveltejs/adapter-auto"; -->
    import adapter from "svelte-adapter-bun";
    import { vitePreprocess } from "@sveltejs/kit/vite";

    /** @type {import('@sveltejs/kit').Config} */
    const config = {
    kit: {
        adapter: adapter(),
    },
    preprocess: vitePreprocess(),
    };

    export default config;

    bun run build
    bun ./build/index.js

### winterjs(待续)

官方项目编译不通过本地测试不能

### Faas测试环境 总结数据

#### Serverless FaaS 测试环境 (待续)

#### Vercel

需要科学上网
https://ssr-svelte-benchmark.vercel.app/

#### Netlify

pnpm i -D @sveltejs/adapter-netlify
https://ssr-svelte-benchmark.netlify.app/
可直接访问

#### Deno Land(待续)

#### WinterJS (Warmer)

#### Cloudfare Pages

https://ssr-svelte-benchmark.pages.dev/
免费版的 cloudare pages都有 限制一天
Requests today 61,694 / 100,000

#### Cloudfare Workers

https://ssr-svelte-benchmark.gakaki.workers.dev/
需要科学上网 国内访问不到 速度和Pages差不多
部署需要安装命令行wangle 稍微麻烦

wangle login
wangle deploy

#### AWS Lambda(待续)

## aws llrt的测试本地了只能 估计node staic吧

#### AWS Lambda(待续) LLRT

#### 阿里云函数计算(待续)

#### Demo site use Wintercg

### webp和avif格式可以减少图片大小到60% 70%

### nodejs ssr with React

### rust-ssr with React

跑不起来 只能切换到React再次尝试

1 wasmer
2 ssr-rs
3 aws llrt本地
4 aws 待续
