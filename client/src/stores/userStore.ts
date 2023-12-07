// 管理用户数据相关
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userLogin } from '@/api/user'
import { createPinia } from 'pinia'

export const useUserStore = defineStore('user', () => {
    // 管理用户登录状态的state
    const userState = {
        isLogin: false,  //是否登录
        // userInfo: {},    //用户信息
        token:'',       //用户token
    }
    // 管理用户数据的state
    const userInfo = ref({})

    // 登录
    const login = async ({ username, password }: { username: string, password: any }) => {
        const res = await userLogin(username, password)
        // userState.isLogin = true
        userState.token = res.data.token
        userInfo.value = res
        console.log("res of userstore",res);

    }

    // 登出
    const logout = () => {
        userState.isLogin = false
        userInfo.value = {}
        // this.token = ''
    }

    // 改变登录状态
    const changeLoginState = (state: boolean) => {
        userState.isLogin = state
    }

    // 获取接口数据的action函数
    const getUserInfo = async ({ username, password }: { username: string, password: any }) => {
        const res = await userLogin(username, password)
        userInfo.value = res
    }
    // 退出时清除用户信息
    const clearUserInfo = () => {
        userInfo.value = {}
        userState.isLogin = false
    }

    return {
        userState,
        userInfo,
        login,
        logout,
        changeLoginState,
        getUserInfo,
        clearUserInfo,
    }
})
