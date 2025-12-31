
import request from '@/utils/http'

export function fetchGetProductList(params: Product.Params.Query) {
    return request.get<Product.Response.List>({
      url: '/product/list',
      params
    })
}

export function fetchGetProductDetail(params: {id:number}) {
  return request.get<Product.Response.Detail>({
    url: '/product/detail',
    params
  })
}


export function fetchGetProductPurchase(data: Product.Params.Model) {
  return request.post({
    url: '/product/purchase',
    data
  })
}