import client from "./client";

export function LIST_BLOG(params) {
    return client({
        url: "/vblog/api/v1/blog/",
        method: "get",
        params: params,
    });
}

export function GET_BLOG(id, params) {
    return client({
        url: `/vblog/api/v1/blog/${id}`,
        method: "get",
        params: params,
    });
}

export function DELETE_BLOG(id, params) {
    return client({
        url: `/vblog/api/v1/blog/${id}`,
        method: "delete",
        params: params,
    });
}

export function CREATE_BLOG(data) {
    return client({
        url: `/vblog/api/v1/blog`,
        method: "post",
        data: data,
    });
}
export function UPDATE_BLOG(id,data) {
    return client({
        url: `/vblog/api/v1/blog/${id}`,
        method: "put",
        data: data,
    });
}

export function UPDATE_BLOG_STATUS(id,data) {
    return client({
        url: `/vblog/api/v1/blog/${id}/status`,
        method: "post",
        data:data,
    });
}

