declare namespace Product {
    namespace Params {
        /** 角色搜索参数 */
        type Query = {
            page: number;
            limit: number;
            name?: string;
            code?: string;
            gameId?: number | null;
        }
        interface Model {
            id?: number;
            userId: number | null;
            quantity: number | null;
            requirements: string | null;
        }
    }
    namespace Response {
        type Info = {
            id: number;
            code: string;
            name: string;
            pic: string;
            game: string;
            category: string;
            price: number;
            rate: number;
            salesCount:number;
            status:number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        
        type Detail = {
            id?: number;
            name: string;
            code: string;
            game: string |null;
            price: number | null;
            rate: number;
            witkeyCount: number | null;
            salesCount: number | null;
            category: string |null;
            unit: string | null;
            purchaseLimit: number | null;
            pic: string;
            description: string;
            content: string;
            status: number;
        }
    }
}