import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: true,
    port: 3000,
    fs: {
      strict: false
    }
  },
  base: '/account/',
  plugins: [vue()]
})
