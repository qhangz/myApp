<script setup lang="ts">
import { onMounted, ref, reactive } from 'vue'
import { Search } from '@element-plus/icons-vue'    // 引入搜索图标
import { useUserStore } from '@/stores/userStore';

import { useScroll } from '@vueuse/core'
const { y } = useScroll(window)

// user info interface
useUserStore().isUserLogin()
const isLogin = ref(useUserStore().userState.isLogin)

// user avatar image
const avatarImage = reactive(localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).avatar_image : '')
// console.log("avatarImage", avatarImage);
// console.log("locals", localStorage.getItem('userInfo'));
// 自动补全输入框
interface RestaurantItem {
    value: string
    link: string
}

const state1 = ref('')

const restaurants = ref<RestaurantItem[]>([])
const querySearch = (queryString: string, cb: any) => {
    const results = queryString
        ? restaurants.value.filter(createFilter(queryString))
        : restaurants.value
    // call callback function to return suggestions
    cb(results)
}
const createFilter = (queryString: string) => {
    return (restaurant: RestaurantItem) => {
        return (
            restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        )
    }
}
const loadAll = () => {
    return [
        { value: 'qhangz', link: 'https://github.com/qhangz' },
        { value: 'HANG', link: 'https://github.com/qhangz' },
        { value: 'user', link: 'https://github.com/qhangz' },
    ]
}
const handleSelect = (item: RestaurantItem) => {
    console.log(item)
}
onMounted(() => {
    restaurants.value = loadAll()
})
</script>
<template>
    <div class='app-header' :class="{ hidden: y > 200 }">
        <div class="container">
            <!-- logo -->
            <!-- <h1 class="logo">
                <RouterLink to="/">LOGO</RouterLink>
            </h1> -->
            <!-- header nav -->
            <ul class="app-header-nav">
                <li class="home">
                    <RouterLink to="/" :class="{ 'active': $route.path === '/' }">用户列表</RouterLink>
                </li>
                <li>
                    <RouterLink to="/userdelete" :class="{ 'active': $route.path === '/userdelete' }">用户删除</RouterLink>
                </li>
                <li>
                    <RouterLink to="/userupdate" :class="{ 'active': $route.path === '/userupdate' }">信息更新</RouterLink>
                </li>
                <li>
                    <RouterLink to="/about" :class="{ 'active': $route.path === '/about' }">关于</RouterLink>
                </li>
            </ul>
            <!-- searcbox -->
            <div class="search">
                <el-col :span="200">
                    <el-autocomplete v-model="state1" :fetch-suggestions="querySearch" clearable class="inline-input w-50"
                        placeholder="search" @select="handleSelect" />
                    <el-button style="margin-left: 5px;" :icon="Search" circle />
                </el-col>
            </div>
            <!-- user info -->
            <div class="user" v-if="isLogin">
                <div class="is-login">
                    <div class="notice">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-xiaoxi"></use>
                        </svg>
                        <div class="dropdown">
                            <ul class="drop-item">
                            </ul>
                        </div>
                    </div>
                    <div class="avatar">
                        <RouterLink to="/user">
                            <img class="icon" :src="avatarImage" />
                        </RouterLink>
                    </div>
                </div>
            </div>
            <div class="user" v-else>
                <div class="not-login">
                    <RouterLink to="/login"><span>登录&nbsp;</span>&nbsp;&nbsp;注册</RouterLink>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.app-header {
    background-color: var(--bg1);
    height: 60px;
    width: 100%;
    padding: 5px 0;
    position: fixed;
    top: 0;
    left: 0;
    transition: top 0.2s ease-in-out;
    display: flex;
    z-index: 999;

    &.hidden {
        top: -64px;
    }

    .container {
        display: flex;
        align-items: center;
        justify-content: center;

        gap: 20px;
        width: 100%;

        // .logo {
        //     width: 50px;
        //     margin-left: 0px;

        //     a {
        //         display: block;
        //         height: 50px;
        //         width: 100%;
        //         text-indent: -9999px;
        //         background: url('@/assets/images/logo.png') no-repeat center 0px / contain;
        //     }
        // }

        .app-header-nav {
            width: 560px;
            display: flex;
            padding-left: 20px;
            position: relative;
            z-index: 998;

            li {
                margin-right: 30px;
                width: 70px;
                text-align: center;

                a {
                    font-size: 16px;
                    line-height: 32px;
                    height: 32px;
                    display: inline-block;
                    color: var(--text-color3);

                    &:hover {
                        color: var(--primary-100);
                        border-bottom: 1px solid var(--primary-100);
                    }
                }

                .active {
                    color: var(--primary-100);
                }
            }
        }

        .search {
            display: flex;
        }

        .user {
            margin-left: 300px;

            .is-login {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                // width: 100px;
                gap: 20px;

                .notice {
                    width: 32px;
                    height: 32px;

                    .icon {
                        width: 100%;
                        height: 100%;
                    }
                }

                .notice:hover {
                    cursor: pointer;
                }

                .avatar {
                    width: 40px;
                    height: 40px;
                    cursor: pointer;
                    transition: transform 0.5s ease-in-out;

                    .icon {
                        width: 100%;
                        height: 100%;
                        // background: url('@/assets/images/myavatar.jpg') no-repeat center 0px / contain;
                        border-radius: 50%;
                    }
                }

                .avatar:hover {

                    transform: rotate(360deg);
                }
            }

            .not-login {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                width: 100px;
                background-color: var(--primary-500);
                border-radius: 4px;
                border: 1px solid var(--primary-300);

                a {
                    color: $primary-100;
                    width: 100%;
                    display: flex;
                    justify-content: center;
                    align-items: center;

                    span {
                        border-right: 1px solid var(--primary-300);
                    }
                }
            }

            .not-login:hover {
                background-color: var(--primary-400);
            }
        }
    }
}
</style>
