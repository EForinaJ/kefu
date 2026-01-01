import { AppRouteRecord } from '@/types/router'

export const dashboardRoutes: AppRouteRecord = {
  name: 'Dashboard',
  path: '/dashboard',
  component: '/index/index',
  meta: {
    title: '仪表盘',
    icon: 'solar:pie-chart-2-bold',
  },
  children: [
    {
      path: 'console',
      name: 'Console',
      component: '/dashboard/console',
      meta: {
        title: '工作台',
        keepAlive: false
      }
    },
    {
      path: 'account',
      name: 'Account',
      component: '/account/user-center',
      meta: {
        title: '账户中心',
        keepAlive: true,
        isHide: true,
        isHideTab: true,
        activePath: '/dashboard/console'
      }
    },
  ]
}
