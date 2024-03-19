import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

// static
import adapterStatic from "@sveltejs/adapter-static";
const adapterStaticOrWasmer = adapterStatic({
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

import adapterNetlify from "@sveltejs/adapter-netlify";
const adapterNetlifyDisableEdge = adapterNetlify({
  edge: false,
  split: false,
});
const adapterNetlifyEdge = adapterNetlify({
  edge: true,
});



import adapterCloudflarePage from "@sveltejs/adapter-cloudflare";
const cloudflarePageAdapter = adapterCloudflarePage({
  routes: {
    include: ["/*"],
    exclude: ["<all>"],
  },
});


import adapterCloudflareWorker from "@sveltejs/adapter-cloudflare-workers";
const cloudflareWorkerAdapter = adapterCloudflareWorker({
  routes: {
    include: ["/*"],
    exclude: ["<all>"],
  },
});

const config = {
  kit: {
    adapter: adapterNetlifyEdge,
  },
  preprocess: vitePreprocess(),
};

export default config;
