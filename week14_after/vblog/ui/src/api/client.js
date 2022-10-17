import axios from "axios";
import { useStorage } from "@vueuse/core";

// 初始化一个http客户端实例
// create an axios instance
const client = axios.create({
    // 由于前端配置代理, 所以不用使用baseURL
    // baseURL: 'http://localhost:8050', // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 5000, // request timeout
});

// 添加请求连接器, 用于任务 都需要携带Token, 或者携带base auth
client.interceptors.request.use(
    (request) => {
        // request 的对象
        // const login = useStorage("login", {
        //     username: "",
        //     password: "",
        //     expiredAt: 0,
        // });
        //
        // console.log(login.value.username, login.value.password);
        // request.headers["X-OAUTH-TOKEN"] = "xxxxx";
        // console.log(request);
        //获取当前用户的登录状态（username,password)
        const username = localStorage.getItem("username")|| "";
        const password = localStorage.getItem("password")|| "";
        request.auth = { username,password };

        return request;
    },
    (error) => {
        console.log(error);
    }
);

// 添加响应连接器, 用于处理返回异常处理, code != 0, 报错提示
client.interceptors.response.use(
    (response) => {
        // console.log(response);
        return response;
    },
    (error) => {
        // 这个异常如何处理？
        if (error.response.data){
            return Promise.reject(error);
        }

    }
);

// 导出该client 对象
export default client;
