<script setup lang="ts">
import { ref } from 'vue'
import { useAdd } from '@/hooks/useAdd';
const a = useAdd()
import FlodNav from '@/components/FlodNav.vue'
// import Test from './test.vue'
</script>

<template>
  <div class="about">
    <h1>This is an about page </h1>
    <h3>关于用户协议和隐私政策</h3>
    <div>{{ a }}</div>
    <FlodNav />
    <!-- <Test /> -->
  </div>
</template>

<style lang="scss" scoped>
// 页面居中显示
.about {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}


@media (min-width: 1024px) {
  .about {
    // min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }
}
</style>
