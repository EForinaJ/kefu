import { AppRouteRecord } from '@/types/router'
import { dashboardRoutes } from './dashboard'

import { witkeyRoutes } from './witkey'
import { shopRoutes } from './shop'
import { orderRoutes } from './order'
/**
 * 导出所有模块化路由
 */
export const routeModules: AppRouteRecord[] = [
  dashboardRoutes,
  shopRoutes,
  orderRoutes,
  witkeyRoutes,
  // systemRoutes,
  // resultRoutes,
  // exceptionRoutes
]
