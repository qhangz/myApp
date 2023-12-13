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

// username update
// username, newusername
export const usernameUpdate = (username: string, newusername: string) => {
    return httpInstance({
        url: '/api/user/update/username',
        method: 'POST',
        data: {
            username: username,
            newusername: newusername,
        }
    })
}
// password update
// username, password, newpassword
export const passwordUpdate = (username: string, password: any, newpassword: any) => {
    return httpInstance({
        url: '/api/user/update/password',
        method: 'POST',
        data: {
            username: username,
            password: password,
            newpassword: newpassword,
        }
    })
}
// email update
// username, newemail
export const emailUpdate = (username: string, newemail: string) => {
    return httpInstance({
        url: '/api/user/update/email',
        method: 'POST',
        data: {
            username: username,
            newemail: newemail,
        }
    })
}
// age update
// username, newage
export const ageUpdate = (username: string, newage: number) => {
    return httpInstance({
        url: '/api/user/update/age',
        method: 'POST',
        data: {
            username: username,
            newage: newage,
        }
    })
}
// summary update
// username, newsummary
export const summaryUpdate = (username: string, newsummary: string) => {
    return httpInstance({
        url: '/api/user/update/summary',
        method: 'POST',
        data: {
            username: username,
            newsummary: newsummary,
        }
    })
}