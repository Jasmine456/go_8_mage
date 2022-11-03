# RESTful 客户端

对接用户中心 RESTful接口的客户端

```
http.Post("http://127.0.0.1:9010/maudit/api/v1/book", body)
```

k8s RESTful API

// RESTful 链式调用客户端
Reousrce("deploy").Namespace("system").LIST().Do().Error

封装一个链式的RESTful客户端:
```
client.POST("/xxxx").Header(k, v).Header(k,v).Body(payload).Do().To(resp).Error
```