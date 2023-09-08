import { RouteRecordRaw } from 'vue-router'

const productRoutes: Array<RouteRecordRaw> = [
    {
      path: '/productorder',
      name: 'productorder',
      component: () => import('@/views/product/ProductOrder.vue')
    },
    {
      path: '/productlists',
      name: 'productlists',
      component: () => import('@/views/product/ProductList.vue')
    }
]

export default productRoutes