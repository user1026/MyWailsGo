import {createApp} from 'vue'
import App from './App.vue'
import router from "@/router/index.js";
//import * as ElementPlusIconsVue from '@element-plus/icons-vue'



const app = createApp(App)
// for (const [key, Component] of Object.entries(ElementPlusIconsVue)) {
//     app.component(key, Component)
// }
app.use(router)
app.mount('#app')
