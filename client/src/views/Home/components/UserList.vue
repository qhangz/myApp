<script setup lang="ts">
import { getUserList } from '@/api/user';
import { onMounted, ref } from 'vue';
// interface User {
//                   avatar_image: string;
//                   // Add other properties here if needed
//                 }

//                 let userList = ref<User[]>([])
//                 const getUserInfo = async () => {
//                     const res = await getUserList()
//                     userList.value = res.data
//                     console.log(userList.value);
//                 }
//                 onMounted(() => {
//                     getUserInfo()
//                 })

interface User {
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
let userList = ref<User[]>([])
const getUserInfo = async () => {
    const res = await getUserList()
    userList.value = res.data
    // console.log(userList.value);
}
onMounted(() => {
    getUserInfo()
})

// let userList = ref([])
// const getUserInfo = async () => {
//     const res = await getUserList()
//     userList.value = res.data
//     console.log(userList.value);
// }
// onMounted(() => {
//     getUserInfo()
// })
</script>

<template>
    <div class="user-list">
        <div class="title">
            This is user list
        </div>
        <div class="user-item" v-for="(item, index) in userList" :key="index">
            <div class="inner">
                <div class="top">
                    <img class="avatar-image" :src="item.avatar_image" />
                    <div class="uid">UID&nbsp;:&nbsp;{{ item.ID }}</div>
                </div>
                <div class="username">username&nbsp;:&nbsp;{{ item.username }}</div>
                <div class="email">email&nbsp;:&nbsp;{{ item.email }}</div>
                <div class="age">age&nbsp;:&nbsp;{{ item.age }}</div>
                <div class="summary">summary&nbsp;:&nbsp;{{ item.summary }}</div>
            </div>
        </div>


    </div>
</template>
<style scoped lang="scss">
.user-list {
    width: 100%;
    height: 100%;
    margin-top: 25px;
    display: flex;
    flex-wrap: wrap;
    align-items: stretch;
    margin-left: -7px;
    margin-right: -7px;

    .title {
        font-size: 20px;
        color: var(--text-color1);
        width: 100%;
        height: 50px;
        text-align: center;
        // background-color:red;
    }

    .user-item {
        height: 230px;
        cursor: pointer;
        width: 25%;
        padding-left: 10px;
        padding-right: 10px;
        box-sizing: border-box;
        position: relative;
        margin-bottom: 30px;
        color: var(--text-color1);

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
</style>