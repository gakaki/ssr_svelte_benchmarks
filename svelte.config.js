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
import adapter from '@sveltejs/adapter-netlify';



import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter({
			// if true, will create a Netlify Edge Function rather
			// than using standard Node-based functions
			edge: true,

			// if true, will split your app into multiple functions
			// instead of creating a single one for the entire app.
			// if `edge` is true, this option cannot be used
			// split: false
		})
    // adapter: adapter({
    //   // default options are shown. On some platforms
    //   // these options are set automatically — see below
    //   pages: "build",
    //   assets: "build",
    //   fallback: undefined,
    //   precompress: false,
    //   strict: false,
    // }),
  },
  preprocess: vitePreprocess(),
};

export default config;
