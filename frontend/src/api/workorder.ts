import request from '@/utils/http'

export function fetchGetWorkOrderList(params: WorkOrder.Params.Query) {
    return request.get<WorkOrder.Response.List>({
      url: '/workorder/list',
      params
    })
}

export function fetchPostWorkOrderCreate(data: WorkOrder.Params.Model) {
  return request.post({
    url: '/workorder/create',
    data
  })
}
export function fetchPostWorkOrderCancel(data: {id:number}) {
  return request.post({
    url: '/workorder/cancel',
    data
  })
}
export function fetchPostWorkOrderSolve(data: {id:number}) {
  return request.post({
    url: '/workorder/solve',
    data
  })
}