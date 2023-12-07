<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';

const login = () => {
    console.log("login");
    useUserStore().login({ username: 'HANG', password: '123456' })
}
const logout = () => {
    useUserStore().logout()
    console.log("logout");
}
</script>

<template>
    <div class="user">
        this is user page
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
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

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