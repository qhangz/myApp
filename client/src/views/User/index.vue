<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';
import { onMounted, ref } from 'vue';
import { getUserByUsername } from '@/api/user'

const login = () => {
    console.log("login");
    useUserStore().login({ username: 'HANG', password: '123456' })
}
const logout = () => {
    useUserStore().logout()
    console.log("logout");
}

interface MyUser {
    ID: Int16Array;
    createTime: Date;
    updateTime: Date;
    deleteTime: Date;
    username: any;
    password: any;
    email: string;
    age: Int16Array;
    summary: string;
    avatar_iamge: string;
}
let userInfo = ref<MyUser[]>([])

const getUserInfo = async () => {
    let myUserName = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
    const res = await getUserByUsername(myUserName)
    userInfo.value = res.data
}
onMounted(() => {
    getUserInfo()
})
</script>

<template>
    <div class="user">
        <div class="this-user">
            <div class="inner">
                <div class="top">
                    <img class="avatar-image" :src="userInfo.avatar_image" alt="avatar image">
                    <div class="uid">UID&nbsp;:&nbsp;{{ userInfo.ID }}</div>
                </div>
                <div class="bottom">
                    <div class="username">用户名&nbsp;:&nbsp;{{ userInfo.username }}</div>
                    <div class="email">邮箱&nbsp;:&nbsp;{{ userInfo.email }}</div>
                    <div class="age">年龄&nbsp;:&nbsp;{{ userInfo.age }}</div>
                    <div class="summary">简介&nbsp;:&nbsp;{{ userInfo.summary }}</div>
                </div>
            </div>
        </div>
        <div class="login" @click="login()">
            登录
        </div>
        <div class="logout" @click="logout()">
            登出
        </div>
    </div>
</template>

<style lang="scss" scoped>
.user {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    // min-height: calc(100vh - 188px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 1180px;

    .this-user {
        height: 230px;
        cursor: pointer;
        width: 25%;
        padding-left: 10px;
        padding-right: 10px;
        box-sizing: border-box;
        position: relative;
        margin-bottom: 30px;
        color: var(--text-color1);
        margin-top: 50px;

        .inner {
            background-color: var(--bg1);
            box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
            transition: all .2s linear;
            border-radius: 5%;
            border-top-left-radius: 20%;
            padding: 20px 15px 10px;

            .top {
                display: flex;
                flex-direction: row;
                margin-bottom: 15px;

                .avatar-image {
                    background-size: cover;
                    background-repeat: no-repeat;
                    width: 30%;
                    height: 30%;
                    // cursor: pointer;
                    transition: transform 0.5s ease-in-out;
                    border-radius: 50%;
                    margin-top: -7%;
                    margin-left: -5%;
                }

                .avatar-image:hover {
                    transform: rotate(360deg);
                }

                .uid {
                    font-size: 14px;
                    font-weight: 600;
                    margin-top: 10px;
                    margin-bottom: 10px;
                    margin-left: 20px;
                }

            }

            .username {
                font-size: 16px;
                font-weight: 600;
                margin-top: 10px;
                margin-bottom: 10px;
                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

            .email {
                font-size: 14px;
                font-weight: 600;
                margin-top: 10px;
                margin-bottom: 10px;
                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

            .age {
                font-size: 14px;
                font-weight: 600;
                margin-top: 10px;
                margin-bottom: 10px;
            }

            .summary {
                font-size: 14px;
                font-weight: 600;
                margin-top: 10px;
                margin-bottom: 10px;

                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

            &:hover {
                content: "";
                transform: translate3d(0, -3.5px, 0);
                box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
            }
        }
    }

    .login {
        width: 100px;
        height: 50px;
        background-color: var(--primary-100);
        color: var(--text-color2);
        box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.25);
        border-radius: 10px;
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        margin-top: 10px;
    }

    .logout {
        width: 100px;
        height: 50px;
        background-color: var(--primary-100);
        color: var(--text-color2);
        box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.25);
        border-radius: 10px;
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        margin-top: 10px;
    }
}
</style>