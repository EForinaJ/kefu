declare namespace Witkey {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name' | 'phone'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            name: string | null;
            titleId: number | null;
            gameId: number| null;
            userId: number| null;
        }
        interface ChangeTitle {
            id: number | null;
            titleId: number | null;
            gameId: number| null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            user: {
                name:string;
                avatar:string;
            };
            game: string;
            title: string;
            status: number;
            rate: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            name:string;
            user: {
                name:string;
                avatar:string;
            };
            game: string;
            title: string;
            rate: number;
            status: number;
            album: string[];
            createTime: string;
        }


    }
}