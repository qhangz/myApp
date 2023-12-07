<template>
    <div id="accordian">
        <ul class="show-dropdown main-navbar">
            <div class="selector-active"></div>
            <li :class="{ 'active': $route.path === '/' }">
                <a href="javascript:;"><i class="fas fa-tachometer-alt"></i>Dashboard</a>
            </li>
            <li :class="{ 'active': $route.path === '/address-book' }">
                <a href="javascript:;"><i class="far fa-address-book"></i>Address Book</a>
            </li>
            <!-- 其他菜单项 -->
        </ul>
    </div>
</template>
  
<script>
export default {
    // 组件的逻辑代码
}
</script>
  
<style lang="scss">
@import url("https://fonts.googleapis.com/css?family=Roboto");

body {
    font-family: "Roboto", sans-serif;
    margin: 0;
    padding: 0;
}

i {
    margin-right: 10px;
}

h2,
h6 {
    margin: 0;
}

#accordian {
    background: #5161ce;
    width: 250px;
    padding: 10px;
    float: left;
    height: 100vh;
    overflow-x: hidden;
    position: relative;
    padding-right: 0px;

    .main-navbar {
        position: relative;

        li {
            list-style-type: none;

            a {
                color: rgba(255, 255, 255, 0.5);
                text-decoration: none;
                font-size: 15px;
                line-height: 45px;
                display: block;
                padding: 0px 20px;
                transition-duration: 0.6s;
                transition-timing-function: cubic-bezier(0.68, -0.55, 0.265, 1.55);
                position: relative;

                &:not(:only-child):after {
                    content: "\f105";
                    position: absolute;
                    right: 20px;
                    top: 10%;
                    font-size: 14px;
                    font-family: "Font Awesome 5 Free";
                    display: inline-block;
                    padding-right: 3px;
                    vertical-align: middle;
                    font-weight: 900;
                    transition: 0.5s;
                }
            }

            &.active {
                a {
                    color: #5161ce;
                    background-color: transparent;
                    transition: all 0.7s;

                    &:not(:only-child):after {
                        transform: rotate(90deg);
                    }
                }
            }
        }

        .selector-active {
            width: 100%;
            display: inline-block;
            position: absolute;
            height: 37px;
            top: 0px;
            left: 0px;
            transition-duration: 0.6s;
            transition-timing-function: cubic-bezier(0.68, -0.55, 0.265, 1.55);
            background-color: #fff;
            border-top-left-radius: 50px;
            border-bottom-left-radius: 50px;

            .top,
            .bottom {
                position: absolute;
                width: 25px;
                height: 25px;
                background-color: #fff;
                right: 0;

                &.top {
                    top: -25px;
                }

                &.bottom {
                    bottom: -25px;
                }

                &.top:before,
                &.bottom:before {
                    content: "";
                    position: absolute;
                    width: 50px;
                    height: 50px;
                    border-radius: 50%;
                    background-color: #5161ce;
                }

                &.top:before {
                    left: -25px;
                    top: -25px;
                }

                &.bottom:before {
                    bottom: -25px;
                    left: -25px;
                }
            }
        }
    }
}
</style>
  