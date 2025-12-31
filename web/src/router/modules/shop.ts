import { AppRouteRecord } from '@/types/router'

export const shopRoutes: AppRouteRecord = {
  path: '/shop',
  name: 'Shop',
  component: '/index/index',
  meta: {
    title: '商城中心',
    icon: 'solar:shop-bold',
    roles: ['R_SUPER']
  },
  children: [
    {
      path: 'shop_product',
      name: 'ShopProduct',
      component: '/shop/product',
      meta: {
        title: '商品列表',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'shop_order',
      name: 'ShopOrder',
      component: '/shop/order',
      meta: {
        title: '订单列表',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'shop_aftersales',
      name: 'ShopAftersales',
      component: '/order/aftersales',
      meta: {
        title: '售后中心',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'shop_settlement',
      name: 'ShopSettlement',
      component: '/order/settlement',
      meta: {
        title: '报单结算',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
  ]
}
