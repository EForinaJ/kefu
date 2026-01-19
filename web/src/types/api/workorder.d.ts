declare namespace WorkOrder {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            code?: string;
            status?: number;
        }
        interface Model {
            title: string | null,
            content: string | null,
            priority: number| null,
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            code: string;
            title: string;
            content: string;
            priority: number;
            status: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            avatar: string;
            sex: number;
            address: string[];
            birthday: string;
            name: string;
            game: string;
            title: string;
            rate: number;
            album: string[];
            status: number;
            createTime: string;
        }
    }
}