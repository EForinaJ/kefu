declare namespace Witkey {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            name?: string;
            phone?: string;
        }
        interface Model {
            name: string | null,
            phone: string | null,
            password: string | null,
            address: string[],
            birthday: number,
            sex: number | null,
            titleId: number | null;
            gameId: number| null;
        }
        interface ChangeTitle {
            id: number | null;
            titleId: number | null;
            gameId: number| null;
            description: string| null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            game: string;
            title: string;
            avatar: string;
            sex: number;
            rate: number;
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