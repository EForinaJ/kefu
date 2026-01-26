import request from '@/utils/http'

export function fetchGetWitkeyList(params: Witkey.Params.Query) {
    return request.get<Witkey.Response.List>({
      url: '/witkey/list',
      params
    })
}

export function fetchGetWitkeyDetail(params: {id:number}) {
  return request.get<Witkey.Response.Detail>({
    url: '/witkey/detail',
    params
  })
}
export function fetchPostWitkeyCreate(data: Witkey.Params.Model) {
  return request.post({
    url: '/witkey/create',
    data
  })
}
export function fetchPostWitkeyChangeTitle(data: Witkey.Params.ChangeTitle) {
  return request.post({
    url: '/witkey/change/title',
    data
  })
}