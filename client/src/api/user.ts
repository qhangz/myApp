import httpInstance from '@/utils/http'

// user register
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

// user list
export const getUserList = () => {
    return httpInstance({
        url: '/api/user/list',
        method: 'GET',
    })
}

// get user by username
export const getUserByUsername = (username: string) => {
    return httpInstance({
        url: `/api/user/info/${username}`,
        method: 'GET'
    });
}

// user delete
export const userDelete = (username: string) => {
    return httpInstance({
        url: '/api/user/delete',
        method: 'POST',
        data: {
            username: username,
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