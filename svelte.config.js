// import adapter from "@sveltejs/adapter-static";

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


// for cloudflare-pages
// pnpm i -D @sveltejs/adapter-cloudflare
// for cloudflare-workers

import adapter from '@sveltejs/adapter-netlify';

import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";
const adapterCloudflarePages = adapter({
  // See below for an explanation of these options
  routes: {
    include: ['/*'],
    exclude: ['<all>']
  }
})
const adapterNetlifyDisableEdge =  adapter({
  // if true, will create a Netlify Edge Function rather
  // than using standard Node-based functions
  edge: false,
  split: false
})
const adapterNetlifyEdge =  adapter({
  edge: true
})

const config = {
  kit: {
    adapter:adapterCloudflarePages()
  },
  preprocess: vitePreprocess(),
};

export default config;
