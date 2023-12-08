import { useUserStore } from "@/stores/userStore"
import { userLogin } from "@/api/user"

const alreadyLoginMsg = () => {
    ElMessage({
        message: '已经登录',
        type: 'warning',
    })
}
const loginSuccessMsg = () => {
    ElMessage({
        message: '登录成功',
        type: 'success',
    })
}
const loginWarningMsg = () => {
    ElMessage({
        message: '登录有误',
        type: 'warning',
    })
}
const errorMsg = () => {
    ElMessage.error('Oops. 出错了')
}

export const myLogin = (username: string, password: string) => {
    const login = async (username: string, password: string) => {
        console.log("userLogin hook");
        if (useUserStore().userState.isLogin === true) {
            console.log("already login");
            alreadyLoginMsg()   //提示已经登录
        } else {
            const res = await userLogin(username, password)
            console.log("logining");
            if (res.code == 200) {
                loginSuccessMsg()   //提示登录成功
                // 更新useUserStore中的信息
                useUserStore().userState.token = res.data.token
                useUserStore().userState.isLogin = true
                useUserStore().userInfo.value = res.data.userInfo
                // localstorage存储登录状态
                localStorage.setItem('isLogin', 'true')
                localStorage.setItem('token', res.data.token)
                localStorage.setItem('userInfo', JSON.stringify(useUserStore().userInfo))
            } else if (res.code == 400) {
                loginWarningMsg()   //提示登录失败
            } else {
                errorMsg()   //提示登录失败
            }
        }
        console.log("userlogin hook end");
    }

    return login(username, password)
}
