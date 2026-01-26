
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
  return request.post<{
    code:string
  }>({
    url: '/product/purchase',
    data
  })
}