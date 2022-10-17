import NProgress from "nprogress";
import "nprogress/nprogress.css"
export function beforeEachHandler(to, from, next) {
    NProgress.start()
//哪些页面需要认证：/backend/xxxx 都需要登陆后才能访问
    if (to.fullPath.indexOf('/backend') === 0) {
        // 正在访问backend的页面
        // 判断用户是否登录
        const username = localStorage.getItem("username")
        const password = localStorage.getItem("password")
        if (
            username === null || password === null || username === "" || password === ""
        ) {
            // 未登录
            console.log("not login")
            //    这里的Next 就相当于router.push()
            next({
                //携带上用户之前访问页面的所有参数：页面的名称，页面参数（&name=xxx&env=test)
                query: {
                    redirect: to.name, ...to.query,
                },
            });
        } else {
            console.log("login")
            // 已登录，直接放行
            next();
        }
    } else {
        // 访问的非后台页面，直接放行
        next();
    }

}

export function afterEachHandler() {
    NProgress.done();
}
