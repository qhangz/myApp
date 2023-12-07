<script setup lang="ts">
import BackTop from '@/components/BackTop.vue'
</script>

<template>
    <div class="home">
        this is home page
        
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