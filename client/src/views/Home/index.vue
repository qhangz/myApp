<script setup lang="ts">
import BackTop from '@/components/BackTop.vue'
import UserList from './components/UserList.vue'
import UserUpdate from './components/UserUpdate.vue'
import UserDelete from './components/UserDelete.vue'
</script>

<template>
    <div class="home">
        <UserList />
        <UserUpdate />
        <UserDelete />
        <BackTop />
    </div>
</template>

<style lang="scss" scoped>
.home {
    display: flex;
    // justify-content: center;
    align-items: center;
    flex-direction: column;
    min-height: calc(100vh - 188px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 1280px;
}
</style>