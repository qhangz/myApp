// 引入初始化样式文件
import '@/styles/base.scss'
// import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
// import { ElMessage } from 'element-plus'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
