<script setup lang="ts">
// import HomeNav from './components/HomeNav.vue'
// import HomeAside from './components/HomeAside.vue'
import BackTop from '@/components/BackTop.vue'
</script>

<template>
    <div class="home">
        this is home page
        <div class="test" style="height:5000px;width:auto;background-color: antiquewhite;">
            d
        </div>
        <!-- <HomeNav /> -->
        <!-- <router-view /> -->
        <!-- <HomeAside /> -->
        <BackTop />
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    padding-top: 20px;
    gap: 20px;

}
</style>