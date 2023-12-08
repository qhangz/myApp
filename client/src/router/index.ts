import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import Layout from '@/views/Layout/index.vue'
import { useUserStore } from '@/stores/userStore'
import { createPinia } from 'pinia';

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'layout',
      component: Layout,
      children: [
        {
          // home page，path: ' ' is the default child path
          path: '/',
          name: 'home',
          component: () => import('../views/Home/index.vue'),
          meta: { 
            title: '主页' 
          },
        },
        {
          // user page
          path: '/user',
          name: 'user',
          component: () => import('../views/User/index.vue'),
          meta: {
            isAuth: true,
            title: '用户页', 
          }
        }
      ]
    },

    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/index.vue'),
      meta: { 
        title: '登录' 
      },
    },

    // about page
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/About/AboutView.vue'),
      meta: { 
        title: '关于' 
      },
    }
  ]
})

const loginMsg = () => {
  ElMessage.error('请先登录！')
}

// 通过useUserStore()获取userStore，进行登录状态判断
// let userStore: any = null
// router.beforeEach(async (to, from, next) => {
//   if (userStore === null) {
//     userStore = await useUserStore()
//   }
//   if (to.meta.isAuth) { // 对需要对登录的页面进行判断
//     if (userStore.userState.isLogin === true) {
//       next()  // 路由页面防止
//     } else {
//       next('/login')
//       loginMsg()
//     }
//   } else {
//     next()
//   }
// })

// 通过localStorage获取登录状态
router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title as string ? to.meta.title : '加载中';
  }
  next();
  if (to.meta.isAuth) { // 判断该路由是否需要登录权限
    if (localStorage.isLogin === 'true') {
      next();
    } else {
      next('/login');
      loginMsg();
    }
  } else {
    next();
  }
})

// router.beforeEach((to, from, next) => {
//   if (to.matched.length === 0) {
//     // 路由不存在
//     next('/404');
//   } else {
//     next();
//   }
// });

export default router
