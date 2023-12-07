import { createRouter, createWebHistory,createWebHashHistory } from 'vue-router'
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
          
        },
        {
          // events page
          path: '/events',
          name: 'events',
          component: () => import('../views/Events/index.vue')
        },
        {
          // chanllenge page
          path: '/chanllenge',
          name: 'chanllenge',
          component: () => import('../views/Chanllenge/index.vue')
        },
        {
          // course page
          path: '/course',
          name: 'course',
          component: () => import('../views/Course/index.vue')
        },
        {
          // games page
          path: '/games',
          name: 'games',
          component: () => import('../views/Games/index.vue')
        },
        {
          // user page
          path: '/user',
          name: 'user',
          component: () => import('../views/User/index.vue')
        }
      ]
    },

    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/index.vue')
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
        isAuth: true
      }
    }
  ]
})

const loginMsg = () => {
  ElMessage.error('请先登录！')
}

let userStore: any = null
router.beforeEach(async (to, from, next) => {
  // userStore = useUserStore()
  if (userStore === null) {
    userStore = await useUserStore()
  }
  console.log("router.ts:", userStore.userState.isLogin);
  if (to.meta.isAuth) { // 对需要对登录的页面进行判断
    if (userStore.userState.isLogin === true) {
      next()  // 路由页面防止
    } else {
      next('/login')
      loginMsg()
    }
  } else {
    next()
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
