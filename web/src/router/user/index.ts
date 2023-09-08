import { RouteRecordRaw } from 'vue-router'

const userRoutes: Array<RouteRecordRaw> = [
    {
      path: '/userinfo',
      name: 'userinfo',
      component: () => import('@/views/user/UserInfo.vue')
    },
]

export default userRoutes