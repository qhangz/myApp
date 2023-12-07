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
        // 在请求发送前做一些处理，如添加token等
        return config
    },
    error => {
        // 请求错误时的处理逻辑
        return Promise.reject(error)
    }
)

// axios响应式拦截器
httpInstance.interceptors.response.use(
    response => {
        // 对相应数据做一些处理，例如根据状态码进行错误处理等
        return response.data
    },
    error => {
        // 对相应错误做一些处理
        return Promise.reject(error)
    }
)


export default httpInstance