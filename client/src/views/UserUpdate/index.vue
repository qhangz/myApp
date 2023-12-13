<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { getUserByUsername, usernameUpdate, passwordUpdate, emailUpdate, ageUpdate, summaryUpdate } from '@/api/user';
import { useUserStore } from '@/stores/userStore';

// form info
const form = reactive({
    username: '',
    password: '',
    originpassword: '',
    email: '',
    age: '',
    summary: '',
})
let myUserName = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
// change function
const changeUsername = () => {
    console.log("changeUsername");
    console.log(form.username);
    if (form.username != '') {
        usernameUpdate(myUserName, form.username)
        // 登出，需要重新登录
        useUserStore().logout()
    }
}
const changePassword = () => {
    console.log("changePassword");
    console.log(form.password);
}
const changeEmail = () => {
    console.log("changeEmail");
    console.log(form.email);
}
const changeAge = () => {
    console.log("changeAge");
    console.log(form.age);
}
const changeSummary = () => {
    console.log("changeSummary");
    console.log(form.summary);
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
    <div class="user-update">
        <div class="title">
            This is user update
        </div>
        <div class="container">
            <div class="user-card">
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
            </div>
            <div class="update-box">
                <div class="form update-username">
                    <input type="text" placeholder="修改用户名" v-model="form.username">
                    <button class="changebtn" @click="changeUsername">修改</button>
                </div>
                <div class="form update-password">
                    <input class="password-input" type="password" placeholder="原密码" v-model="form.password">
                    <input class="password-input" type="password" placeholder="新密码" v-model="form.originpassword">
                    <button class="changebtn" @click="changePassword">修改</button>
                </div>
                <div class="form update-email">
                    <input type="text" placeholder="修改邮箱" v-model="form.email">
                    <button class="changebtn" @click="changeEmail">修改</button>
                </div>
                <div class="form update-age">
                    <input type="text" placeholder="修改年龄" v-model="form.age">
                    <button class="changebtn" @click="changeAge">修改</button>
                </div>
                <div class="form update-summary">
                    <input type="text" placeholder="修改简介" v-model="form.summary">
                    <button class="changebtn" @click="changeSummary">修改</button>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.user-update {
    display: flex;
    align-items: center;
    flex-direction: column;
    min-height: calc(100vh - 188px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 1280px;

    .title {
        font-size: 20px;
        color: var(--text-color1);
        width: 100%;
        height: 50px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .container {
        display: flex;
        width: 100%;
        justify-content: center;
        align-items: center;
        flex-direction: column;

        .user-card {
            display: flex;
            width: 300px;

            // background-color: aqua;
            .this-user {
                // height: 230px;
                cursor: pointer;
                width: 100%;
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
        }

        .update-box {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            color: var(--text-color1);
            width: 100%;

            .form {
                width: 100%;
                padding: 10px 0;
                display: flex;
                align-items: center;
                justify-content: center;

                input {
                    width: 17%;
                    height: 35px;
                    border: none;
                    outline: none;
                    border-radius: 10px;
                    padding-left: 2em;
                    // background-color: #f0f0f0;
                    box-shadow: inset 2px 2px 4px #d1d9e6, inset -2px -2px 4px #f9f9f9;
                    // box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
                    margin: 0 20px;
                }

                .password-input {
                    height: 35px;
                    width: 8.5%;
                    margin: 0 10px;
                }

                .changebtn {
                    width: 7%;
                    height: 40px;
                    border-radius: 24px;
                    border: none;
                    outline: none;
                    background-color: $primary-200;
                    color: #fff;
                    font-size: 0.9em;
                    cursor: pointer;
                    box-shadow: 8px 8px 16px #d1d9e6, -8px -8px 16px #f9f9f9;
                }
            }
        }
    }
}
</style>