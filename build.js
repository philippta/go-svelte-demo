import esbuild from "esbuild";
import sveltePlugin from "esbuild-svelte";

esbuild
  .build({
    entryPoints: ["templates/components/index.js"],
    mainFields: ["svelte", "browser", "module", "main"],
    bundle: true,
    minify: true,
    outfile: "./static/components.js",
    plugins: [sveltePlugin()],
    logLevel: "info",
  })
  .catch(() => process.exit(1));
