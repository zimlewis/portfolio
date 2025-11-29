import { defineConfig } from "vite";
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    build: {
        // where Vite outputs your JS/CSS
        outDir: "public/assets",

        // don’t put inside nested /assets folder
        assetsDir: "",

        // your JS entrypoint(s)
        rollupOptions: {
            input: {
                main: "frontend/scripts/main.ts"   // <--- your entry
            },
            output: {
                entryFileNames: `[name].js`,
                assetFileNames: `[name].[ext]`,
            }
        },

        emptyOutDir: true,
    },
    plugins: [
        tailwindcss()
    ]
});
