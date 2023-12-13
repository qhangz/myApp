<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { getUserByUsername, usernameUpdate, passwordUpdate, emailUpdate, ageUpdate, summaryUpdate } from '@/api/user';

// message function
const updateSuccessMsg = () => {
    ElMessage({
        message: '更新成功',
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
    password: '',
    newusername: '',
    newpassword: '',
    newemail: '',
    newage: '',
    newsummary: '',
})
// change function
const changeUsername = () => {
    console.log("changeUsername:");
    console.log(form.username, form.newusername);
    if (form.username != '' && form.newusername != '') {
        usernameUpdate(form.username, form.newusername).then(res => {
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
const changePassword = () => {
    console.log("changePassword:");
    console.log(form.username, form.password, form.newpassword);
    if (form.username != '' && form.password != '' && form.newpassword != '') {
        passwordUpdate(form.username, form.password, form.newpassword).then(res => {
            console.log("res of passwordUpdate:", res);
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
const changeEmail = () => {
    console.log("changeEmail:");
    console.log(form.username, form.newemail);
    if (form.username != '' && form.newemail != '') {
        emailUpdate(form.username, form.newemail).then(res => {
            console.log("res of emailUpdate:", res);
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
const changeAge = () => {
    console.log("changeAge:");
    console.log(form.username, form.newage);
    if (form.username != '' && form.newage != '') {
        ageUpdate(form.username, form.newage).then(res => {
            console.log("res of ageUpdate:", res);
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
const changeSummary = () => {
    console.log("changeSummary:");
    console.log(form.username, form.newsummary);
    if (form.username != '' && form.newsummary != '') {
        summaryUpdate(form.username, form.newsummary).then(res => {
            console.log("res of summaryUpdate:", res);
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
    <div class="user-update">
        <div class="title">
            user update form
        </div>
        <div class="update-box">
            <div class="form origin-msg">
                <input type="text" placeholder="原用户名" v-model="form.username">
                <input type="password" placeholder="原密码" v-model="form.password">
            </div>
            <div class="form update-username">
                <input type="text" placeholder="新的用户名" v-model="form.newusername">
                <button class="changebtn" @click="changeUsername">修改</button>
            </div>
            <div class="form update-password">
                <input type="password" placeholder="新密码" v-model="form.newpassword">
                <button class="changebtn" @click="changePassword">修改</button>
            </div>
            <div class="form update-email">
                <input type="text" placeholder="新的邮箱" v-model="form.newemail">
                <button class="changebtn" @click="changeEmail">修改</button>
            </div>
            <div class="form update-age">
                <input type="text" placeholder="新的年龄" v-model="form.newage">
                <button class="changebtn" @click="changeAge">修改</button>
            </div>
            <div class="form update-summary">
                <input type="text" placeholder="新的简介" v-model="form.newsummary">
                <button class="changebtn" @click="changeSummary">修改</button>
            </div>
        </div>
        <div class="about">
            <h5>更新需要输入：</h5>
            <h5>更新用户名：原用户名，新的用户名</h5>
            <h5>更新密码：原用户名，原密码，新密码</h5>
            <h5>更新邮箱：原用户名，新的邮箱</h5>
            <h5>更新年龄：原用户名，新的年龄</h5>
            <h5>更新简介：原用户名，新的简介</h5>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.user-update {
    display: flex;
    align-items: center;
    flex-direction: column;
    // min-height: calc(100vh - 188px);
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