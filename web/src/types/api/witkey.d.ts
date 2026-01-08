declare namespace Witkey {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name' | 'phone'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            name: string | null,
            phone: string | null,
            password: string | null,
            address: string[],
            birthday: number,
            description: string | null,
            sex: number | null,
            avatar: string,
            status: number,
            titleId: number | null;
            gameId: number| null;
            album: string[];
            rate: number;
        }
        interface ChangeTitle {
            id: number | null;
            titleId: number | null;
            gameId: number| null;
            // type: number| null;
            // description: string| null;
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