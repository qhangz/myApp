import axios from 'axios'

// 创建axios实例
const httpInstance = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 5000
})

// axios请求拦截器
httpInstance.interceptors.request.use(
    config => {
        // 设置axios为form-data
        if (config.method === 'post') {
            config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
        }

        return config
    },
    error => {
        Promise.reject(error)
    }
)

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res.data, e => {
    return Promise.reject(e)
})


export default httpInstance