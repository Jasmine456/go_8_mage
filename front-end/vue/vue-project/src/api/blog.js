import client from './client'

export function LIST_BLOG(params){
   return client({url:"/vblog/api/v1/blog/",method:"get",params:params});
}

export function GET_BLOG(id,params){
    return client({
        url:`/vblog/api/v1/blog/${id}`,
        method:"get",
        params:params,
    });
}