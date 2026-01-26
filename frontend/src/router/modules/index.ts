import { AppRouteRecord } from '@/types/router'
import { dashboardRoutes } from './dashboard'

import { witkeyRoutes } from './witkey'
import { shopRoutes } from './shop'
import { orderRoutes } from './order'
import { accountRoutes } from './account'
/**
 * 导出所有模块化路由
 */
export const routeModules: AppRouteRecord[] = [
  dashboardRoutes,
  shopRoutes,
  orderRoutes,
  witkeyRoutes,
  accountRoutes,
  // resultRoutes,
  // exceptionRoutes
]
