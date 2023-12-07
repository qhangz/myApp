import httpInstance from '@/utils/http'

// user register
// export const userRegister = (parma: any) => {
//     return httpInstance({
//         url: '/api/user/register',
//         method: 'POST',
//         data: parma
//     })
// }

// export const userRegister=(parma:any)=>{
//     return httpInstance({
//         url:'/api/user/register',
//         method:'POST',
//         data:{
//             username:parma.username,
//             password:parma.password,
//             email:parma.email,
//         }
//     })
// }

export const userRegister = (username: string, password: any, email: any) => {
    return httpInstance({
        url: '/api/user/register',
        method: 'POST',
        data: {
            username: username,
            password: password,
            email: email,
        }
    })
}

// user login
export const userLogin = (username: string, password: any) => {
    return httpInstance({
        url: '/api/user/login',
        method: 'POST',
        data: {
            username,
            password,
        }
    })
}

// user info
// export const userInfo = () => {
//     return httpInstance({
//         url: '/api/user/info',
//         method: 'GET',
//     })
// }