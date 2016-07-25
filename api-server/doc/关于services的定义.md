关于service
====

概述:
----

1. 一个service层, 是指业务逻辑层, 之所以有这一层, 是因为, 放在controller层与model都有一些问题
2. 放在controller层的话, 从定义上说, controller层应该做调用, 它根据HTTP请求调用资源
3. 放在model层的话, 一早一个业务逻辑关联到其它model的话, 那么这个业务归属就较难定义了
4. 因此有了service层, 让controller做应该做的事情, 让model层只关心自己的model

补充:
----
1. controller层只关心http相关功能, controller只能调用service, 不能调用model的模块
2. service是业务逻辑层, 用于完成实际业务逻辑, service可以调用其它service, 也能调用model层, service对controller一无所知
3. model层db层, 只提供一些数据逻辑层面功能, 不关心具体的业务, 只封装简单的功能, 当一个model需要关心多个表的时候, 将其提升到service层






