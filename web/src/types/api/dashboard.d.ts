declare namespace Dashboard {

    namespace Response {
        type Detail = {
            todaySales:number;
            salesComparison:number;
            hotProductList:Product.Response.Info[];
        }
    }
}