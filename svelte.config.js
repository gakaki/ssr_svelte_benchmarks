import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

// static
import adapterStatic from "@sveltejs/adapter-static";
const adapterWasmerOrStatic = adapterStatic({
  pages: "build",
  assets: "build",
  fallback: undefined,
  precompress: false,
  strict: false,
});

// Node
import adapterNode from "@sveltejs/adapter-node";
const adapterNodeJS = adapterNode({
  pages: "build",
  assets: "build",
  fallback: undefined,
  precompress: false,
  strict: false,
});
// for bun
// import adapter from "svelte-adapter-bun";

// for node.js
// pnpm i -D @sveltejs/adapter-node
// import adapter from '@sveltejs/adapter-node';

// for deno someissuse 500 can not build success
// import adapter from 'sveltekit-adapter-deno';
// pnpm run build
// deno run --allow-env --allow-read --allow-net mod.ts

// import adapter from '@sveltejs/adapter-vercel';

import adapterNetlify from "@sveltejs/adapter-netlify";
const adapterNetlifyDisableEdge = adapterNetlify({
  edge: false,
  split: false,
});
const adapterNetlifyEdge = adapterNetlify({
  edge: true,
});

// for cloudflare-pages
// pnpm i -D @sveltejs/adapter-cloudflare
import adapterCloudflarePage from "@sveltejs/adapter-cloudflare";
const cloudflarePageAdapter = adapterCloudflarePage({
  routes: {
    include: ["/*"],
    exclude: ["<all>"],
  },
});

// for cloudflare-workers
// pnpm i -D @sveltejs/adapter-cloudflare-workers
// pnpm i -g wrangler
// wrangler login
// wrangler deploy

import adapterCloudflareWorker from "@sveltejs/adapter-cloudflare-workers";
const cloudflareWorkerAdapter = adapterCloudflareWorker({
  routes: {
    include: ["/*"],
    exclude: ["<all>"],
  },
});

const config = {
  kit: {
    adapter: adapterWasmerOrStatic,
  },
  preprocess: vitePreprocess(),
};

export default config;
