import {
  defineConfig
} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

import {
  ElementPlusResolver
} from 'unplugin-vue-components/resolvers'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),

  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src") // map '@' to './src'
    },
  },
  define: {
    "process.env": {
      "BASE_URL": "http://localhost:9000"
    }
  },

  server: {
    host: '127.0.0.1',
    port: 7000,
    proxy: {
      // 选项写法
      '/qqapi': {
        target: 'https://zy.xywlapi.cc', // 所要代理的目标地址，也就是后台接口地址
        //rewrite: path => path.replace(/^\/cc/g, ''),
        changeOrigin: true, // true/false, Default: false - changes the origin of the host header to the target URL
      },
        '/ip': {
          target: "https://api.wrdan.com",
          changeOrigin: true,
        }
      }
  }
})