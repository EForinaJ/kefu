import request from '@/utils/http'

export function fetchGetWitkeyList(params: Witkey.Params.Query) {
    return request.get<Witkey.Response.List>({
      url: '/witkey/list',
      params
    })
}
