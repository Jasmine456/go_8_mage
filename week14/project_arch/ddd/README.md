# 基于领域驱动的代码组织架构

## 划分业务领域

把一个大的业务系统（复杂的问题），把他拆解成多个小的业务领域（小的问题），各个击破（分而治之的思想
+ blog：文章管理
+ tag：标签管理

## 领域建模

定义数据结构和接口，这个领域的持有该数据的所有权

其他服务模块比如blog需要获取该blog的所有tag，需要通过tag这个领域模块提供的领域方法来操作。

## 面向对象的调用模型

低耦合高内聚，面向对象的编程思维

HTTP Server(Gin) --> Blog HTTP Handler(Object) --> Blog Servce impl(Object,
可以有多种实现 mock，mysql，...) [Controller --> DAO] --> Tag Server Impl[Controler --> DAO]
