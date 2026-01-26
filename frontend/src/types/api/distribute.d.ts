declare namespace Distribute {
    namespace Params {
        interface Query  {
            page: number;
            limit: number;
            code?: string;
            orderCode?: string;
        }

        interface Cancel {
            id: number;
            reason: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            order:string;
            manage: string;
            witkey:string;
            game:string;
            title:string;
            type:number;
            status:number;
            createTime:string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            order:{
                id: number;
                code:string;
                createTime:string
                user: {
                    id:number
                    name:string;
                    avatar:string;
                };
                product: {
                    id: number;
                    name: string;
                    pic: string;
                    quantity: number;
                    unitPrice: number;
                    unit: string;
                    game: string;
                    category: string;
                };
                totalAmount: number;
                discountAmount: number;
                actualAmount:number;
                payAmount: number;
                status: number;
                payStatus: number;
                payMode: number;
                payTime:string | null;
                startTime:string | null;
                finishTime:string | null;
                requirements:string | null;
            };
            manage: string;
            witkey:string;
            game:string;
            title:string;
            type:number;
            status:number;
            reason:string;
            createTime:string;
        }

    }
}