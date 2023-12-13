<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { userDelete } from '@/api/user';

// message function
const updateSuccessMsg = () => {
    ElMessage({
        message: '删除成功',
        type: 'success',
    })
}
const updateWarningMsg = () => {
    ElMessage({
        message: '信息有误',
        type: 'warning',
    })
}
const errorMsg = () => {
    ElMessage.error('Oops. 出错了')
}
// reload function
const reload = () => {
    window.location.reload()
}
// form info
const form = reactive({
    username: '',
})
// change function
const deteleUsername = () => {
    console.log("changeUsername:");
    console.log(form.username, form.username);
    if (form.username != '') {
        userDelete(form.username).then(res => {
            console.log("res of usernameUpdate:", res);
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else {
        updateWarningMsg()
    }
}


</script>
<template>
    <div class="user-delete">
        <div class="title">
            user detele form
        </div>
        <div class="delete-box">
            <div class="form delete-username">
                <input type="text" placeholder="用户名" v-model="form.username">
                <button class="changebtn" @click="deteleUsername">删除</button>
            </div>
        </div>
        <div class="about">
            <h5>更新需要输入：用户名</h5>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.user-delete {
    display: flex;
    align-items: center;
    flex-direction: column;
    min-height: calc(200px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 1280px;
    // margin-top:20px;

    .title {
        font-size: 20px;
        color: var(--text-color1);
        width: 100%;
        height: 50px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .delete-box {
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

    .about {
        h5 {
            margin: 5px auto;
        }
    }
}
</style>