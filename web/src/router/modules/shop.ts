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
      path: 'shop_comment',
      name: 'ShopComment',
      component: '/shop/comment',
      meta: {
        title: '评价审核',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
  ]
}
