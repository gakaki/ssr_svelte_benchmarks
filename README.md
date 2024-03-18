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

### 本地测试环境  总结数据

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

### Faas测试环境  总结数据


#### Serverless FaaS 测试环境 (待续)

#### Deno Land(待续)

#### WinterJS (Warmer)(待续)

#### Cloudfare Pages(待续)

#### Cloudfare Workers(待续)

#### AWS Lambda(待续) LLRT

#### 阿里云函数计算(待续)

#### Demo site use Wintercg
