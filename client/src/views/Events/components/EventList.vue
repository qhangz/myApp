<script setup lang="ts">

</script>

<template>
    <div class="event-list">
        <a href="javascript:;" class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </a>
        <div class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </div>
        <div class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </div>
        <div class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </div>
        <div class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </div>
        <div class="event-item">
            <div class="inner">
                <div class="banner"></div>
                <div class="message">event</div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.event-list {
    width: 100%;
    height: 100%;
    margin-top: 35px;
    display: flex;
    flex-wrap: wrap;
    align-items: stretch;
    margin-left: -7px;
    margin-right: -7px;

    .event-item {
        // width: 135px;
        height: 230px;
        cursor: pointer;
        width: 25%;
        padding-left: 7px;
        padding-right: 7px;
        box-sizing: border-box;
        position: relative;
        margin-bottom: 30px;
        color: var(--text-color1);

        .inner {
            background-color: var(--bg1);
            box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
            transition: all .2s linear;

            .banner {
                background-image: url("https://i.imgtg.com/2023/07/02/Okushr.jpg");
                background-size: cover;
                padding-top: 58.82%;
                background-color: #ccc;
                background-repeat: no-repeat;
            }

            .message {
                padding: 20px 12px 15px;
                background-color: var(--bg1);
            }

            &:hover {
                content: "";
                transform: translate3d(0, -3.5px, 0);
                box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
            }
        }
    }



    // .event-item:nth-child(4n+1) {
    //     background-color: red;
    // }
    // .event-item::before {
    //     content: "";
    //     position: absolute;
    //     left: 7px;
    //     right: 7px;
    //     top: 0;
    //     bottom: 0;
    //     z-index: -1;
    //     border-radius: 2px;
    //     overflow: hidden;
    //     background-color: red;
    //     box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
    //     transition: all .2s linear;
    // }
}
</style>
```