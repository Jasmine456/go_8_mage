import {fileURLToPath, URL} from "node:url";

import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
    server: {
        //利用后端服务 解决跨域问题
        proxy: {
            //http://localhost:5173/vblog/api/v1/blog
            //http://localhost:7080/vblog/api/v1/blog
            "/vblog/api/v1": "http://localhost:7080/",
        },
    },
});
