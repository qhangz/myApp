<template>
    <div class="home-nav">
        <div v-for="(item, index) in items" :key="index"
            :style="{ backgroundColor: selectedIndex === index ? 'var(--bg5)' : 'white' }" @click="selectItem(index)">
            <svg class="icon" aria-hidden="true">
                <use xlink:href="#icon-menu-dataMining"></use>
            </svg>
            {{ item }}
        </div>
    </div>
</template>

<script lang="ts">
export default {
    data() {
        return {
            items: ['关注', '前端', '后端','Android','IOS','人工智能','代码人生','阅读'],
            selectedIndex: null
        };
    },
    methods: {
        selectItem(index: any) {
            this.selectedIndex = index;
            console.log(this.selectedIndex);
        }
    }
};
</script>

<style lang="scss" scoped>
.home-nav {
    background: var(--bg1);
    border-radius: 5px;
    padding: 10px 0;

    div {
        width: 130px;
        height: 45px;
        display: flex;
        // justify-content: center;
        align-items: center;
        // flex-direction: row;
        color: var(--text-color1);
        margin: 0 10px;
        padding-left: 15px;
        border-radius: 5px;
        .icon{
            font-size:20px;
            margin-right:10px;
        }
    }
}
</style>