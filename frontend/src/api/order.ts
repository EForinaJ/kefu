
import request from '@/utils/http'

export function fetchGetOrderList(params: Order.Params.Query) {
    return request.get<Order.Response.List>({
      url: '/order/list',
      params
    })
}

export function fetchGetOrderDetail(params: {id:number}) {
  return request.get<Order.Response.Detail>({
    url: '/order/detail',
    params
  })
}

export function fetchPostOrderAddDiscount(data: Order.Params.AddDiscount) {
  return request.post({
    url: '/order/add/discount',
    data
  })
}

export function fetchPostOrderPaid(data:Order.Params.Paid) {
  return request.post({
    url: '/order/paid',
    data
  })
}
export function fetchPostOrderCancel(data:{id:number}) {
  return request.post({
    url: '/order/cancel',
    data
  })
}
export function fetchPostOrderStart(data:{id:number}) {
  return request.post({
    url: '/order/start',
    data
  })
}
export function fetchPostOrderComplete(data:{id:number}) {
  return request.post({
    url: '/order/complete',
    data
  })
}
export function fetchPostOrderAftersales(data:Order.Params.Aftersales) {
  return request.post({
    url: '/order/aftersales',
    data
  })
}
export function fetchPostOrderDistribute(data:Order.Params.Distribute) {
  return request.post({
    url: '/order/distribute',
    data
  })
}


// export function fetchGetOrderDistributeDetail(params: {id:number}) {
//   return request.get<Order.Response.DistributeList>({
//     url: '/order/distribute/detail',
//     params
//   })
// }


