import path from "path"
import react from "@vitejs/plugin-react"
import {defineConfig, loadEnv} from "vite"

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');

  return {
    plugins: [react()],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    define: {
      'process.env.REACT_APP_MQTT_URL': JSON.stringify(env.REACT_APP_MQTT_URL),
      'process.env.REACT_APP_MQTT_PASSWORD': JSON.stringify(env.REACT_APP_MQTT_PASSWORD),
      'process.env.REACT_APP_MQTT_USERNAME': JSON.stringify(env.REACT_APP_MQTT_USERNAME),
    },
  }
})

