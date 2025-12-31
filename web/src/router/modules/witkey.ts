import { AppRouteRecord } from '@/types/router'

export const witkeyRoutes: AppRouteRecord = {
  path: '/witkey',
  name: 'Witkey',
  component: '/index/index',
  meta: {
    title: '威客中心',
    icon: 'solar:user-id-bold',
    roles: ['R_SUPER']
  },
  children: [
    {
      path: 'witkey_list',
      name: 'WitkeyList',
      component: '/witkey/list',
      meta: {
        title: '威客列表',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
  ]
}
