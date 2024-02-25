import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

import router from '~/router'
app.use(router)

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
app.use(ElementPlus)


import 'virtual:windi.css'

app.mount('#app')
