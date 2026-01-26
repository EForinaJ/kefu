
import request from '@/utils/http'
/**
 * 获取用户信息
 * @returns 用户信息
 */
export function fetchGetAccountInfo() {
    return request.get<Account.Response.Info>({
      url: '/account'
    })
}

export function fetchPostAccountEdit(data: Account.Params.Model) {
  return request.post({
    url: '/account/edit',
    data
  })
}

export function fetchPostAccountChangePass(data: Account.Params.ChangePass) {
  return request.post({
    url: '/account/change/pass',
    data
  })
}