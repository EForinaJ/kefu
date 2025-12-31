import { AppRouteRecord } from '@/types/router'

export const orderRoutes: AppRouteRecord = {
  path: '/order',
  name: 'Order',
  component: '/index/index',
  meta: {
    title: '订单管理',
    icon: 'solar:clipboard-heart-bold',
    roles: ['R_SUPER']
  },
  children: [
    {
      path: 'order_list',
      name: 'OrderList',
      component: '/order/list',
      meta: {
        title: '订单中心',
        keepAlive: true,
        roles: ['R_SUPER']
      },
    },
    {
      path: 'order_aftersales',
      name: 'OrderAftersales',
      component: '/order/aftersales',
      meta: {
        title: '售后中心',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'order_settlement',
      name: 'OrderSettlement',
      component: '/order/settlement',
      meta: {
        title: '报单结算',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
  ]
}
