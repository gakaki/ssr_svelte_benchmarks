#### 统一步骤

修改svelte.config.js 使用不同 adapter 配置传入, pnpm run build

    import adapterNode from "@sveltejs/adapter-node";
    const adapterNodeJS = adapterNode({
        pages: "build",
        assets: "build",
        fallback: undefined,
        precompress: false,
        strict: false,
    });


#### Deno(待续)

for deno someissuse 500 can not build success

    import adapter from 'sveltekit-adapter-deno';
    pnpm run build
    deno run --allow-env --allow-read --allow-net mod.ts

报错
#### Node.js

    pnpm i -D @sveltejs/adapter-node
    import adapterNode from "@sveltejs/adapter-node";
    const adapterNodeJS = adapterNode({
        pages: "build",
        assets: "build",
        fallback: undefined,
        precompress: false,
        strict: false,
    });

#### [Bun]((https://bun.sh/guides/ecosystem/sveltekit))

    
    pnpm i -D @sveltejs/adapter-bun

    import adapterNode from "svelte-adapter-bun";
    const adapterNodeJS = adapterNode({
        pages: "build",
        assets: "build",
        fallback: undefined,
        precompress: false,
        strict: false,
    });
   

    bun run build
    bun ./build/index.js

#### Winterjs(待续)

官方项目编译不通过本地测试不能,给定Demo是静态文件生成,非SSR

#### Aws LLRT(待续)


#### rust-ssr with React


#### nodejs ssr with React

