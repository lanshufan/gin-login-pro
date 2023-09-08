import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw } from 'vue-router'
import productRoutes from './product'
import userRoutes from './user'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/login/Login.vue')
  },
  ...productRoutes,
  ...userRoutes
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  // 路由守护，判断是否有token
  const token = localStorage.getItem('token')
  if (to.name !== 'Login' && !token) next({ name: 'Login' })
  else next()
})

export default router
