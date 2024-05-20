import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import Order from '../views/Order.vue'
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/order/:regionId/:corporationId/:offerId',
    name: 'Order',
    component: Order,
    props: (route) => ({
      ...route.params,
      ...route.query
    })
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
