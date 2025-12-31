import { AppRouteRecord } from '@/types/router'
import { dashboardRoutes } from './dashboard'

import { witkeyRoutes } from './witkey'
import { shopRoutes } from './shop'
/**
 * 导出所有模块化路由
 */
export const routeModules: AppRouteRecord[] = [
  dashboardRoutes,
  shopRoutes,
  witkeyRoutes,
  // systemRoutes,
  // resultRoutes,
  // exceptionRoutes
]
