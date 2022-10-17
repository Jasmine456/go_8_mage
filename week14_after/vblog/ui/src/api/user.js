import client from "./client"
export function LOGIN(data){
    return client({
        url:"/vblog/api/v1/user/auth",
        method:"post",
        data:data,
    });
}