<template>
    <footer class="app-footer">

        <div class="contact">
            <div class="container">
                <dl class="box1">
                    <dt>客户服务</dt>
                    <dd><i class="iconfont icon-kefu"></i> 在线客服</dd>
                    <dd><i class="iconfont icon-question"></i> 问题反馈</dd>
                </dl>
                <dl class="box2">
                    <dt>关注我们</dt>
                    <dd><i class="iconfont icon-weixin"></i> 公众号</dd>
                    <dd><i class="iconfont icon-weibo"></i> 微博</dd>
                </dl>
                <dl class="box3">
                    <dt>下载APP</dt>
                    <dd class="qrcode"><img src="@/assets/images/qrcode.jpg" /></dd>
                    <dd class="download">
                        <span>扫描二维码</span>
                        <span>立马下载APP</span>
                        <a href="javascript:;">下载页面</a>
                    </dd>
                </dl>
                <dl class="box4">
                    <dt>服务热线</dt>
                    <dd class="hotline">100-8610-086 <small>周一至周日 9:00-21:00</small></dd>
                </dl>
            </div>
        </div>
        <!-- 其它 -->
        <div class="extra">
            <div class="container">
                <!-- 版权信息 -->
                <div class="copyright">
                    <p>
                        <a href="javascript:;">关于我们</a>
                        <a href="javascript:;">帮助中心</a>
                        <a href="javascript:;">开发工具</a>
                        <a href="javascript:;">代码人生</a>
                        <a href="javascript:;">阅读</a>
                        <a href="javascript:;">搜索推荐</a>
                        <a href="javascript:;">友情链接</a>
                    </p>
                    <p>CopyRight © CodeX</p>
                </div>
            </div>
        </div>
    </footer>
</template>

<style scoped lang='scss'>
.app-footer {
    overflow: hidden;
    background-color: #f5f5f5;
    padding-top: 20px;

    .contact {
        background: #fff;
        display: flex;
        justify-content: center;

        .container {
            padding: 15px 0 0px 25px;
            display: flex;
        }

        dl {
            height: 135px;
            text-align: center;
            padding: 0 72px;
            border-right: 1px solid #f2f2f2;
            color: #999;

            &:first-child {
                padding-left: 0;
            }

            &:last-child {
                border-right: none;
                padding-right: 0;
            }
        }

        dt {
            line-height: 1;
            font-size: 18px;
        }

        dd {
            margin: 15px 12px 0 0;
            float: left;
            width: 92px;
            height: 92px;
            padding-top: 10px;
            border: 1px solid #ededed;

            .iconfont {
                font-size: 36px;
                display: block;
                color: #666;
            }

            &:hover {
                .iconfont {
                    color: $xtxColor;
                }
            }

            &:last-child {
                margin-right: 0;
            }
        }

        .qrcode {
            width: 92px;
            height: 92px;
            padding: 7px;
            border: 1px solid #ededed;
        }

        .download {
            padding-top: 5px;
            font-size: 14px;
            width: auto;
            height: auto;
            border: none;

            span {
                display: block;
            }

            a {
                display: block;
                line-height: 1;
                padding: 10px 25px;
                margin-top: 5px;
                color: #fff;
                border-radius: 2px;
                background-color: $xtxColor;
            }
        }

        .hotline {
            padding-top: 20px;
            font-size: 22px;
            color: #666;
            width: auto;
            height: auto;
            border: none;

            small {
                display: block;
                font-size: 15px;
                color: #999;
            }
        }
    }

    .extra {
        background-color: #333;
    }

    .copyright {
        height: 90px;
        padding-top: 20px;
        text-align: center;
        color: #999;
        font-size: 15px;

        p {
            line-height: 1;
            margin-bottom: 20px;
        }

        a {
            color: #999;
            line-height: 1;
            padding: 0 10px;
            border-right: 1px solid #999;

            &:last-child {
                border-right: none;
            }
        }
    }
}

@media(max-width: 1250px) {
    .box3 {
        display: none;
    }
}

@media(max-width: 911px) {
    .box4 {
        display: none;
    }
}

@media(max-width: 650px) {
    .box1 {
        display: none;
    }
    dl{
        padding: 0 0px !important;
    }
}

@media(max-width:580px) {
    a {
        font-size: small;
    }

}

@media(max-width: 520px) {
    a {
        font-size: 10px;
    }
}
@media(max-width: 430px) {
    a {
        font-size: 8px;
    }
}
</style>