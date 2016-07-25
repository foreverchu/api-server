logger格式
====

说明: 
----
packageName.structName.MethodName空格冒号空格其它信息
eg: models.User.FindBy : userId = 1 is not found

如果没有struct,则:
packageName.FunctionName空格冒号空格其它信息
eg: partySrv.PartyCreate : invalid reg time 2013-12-01 13:04:05


目的:
----
1. 日志分析可以较为方便地分析出哪些模块容易出错
2. 可以方便地做统计


方法:
----
通过在method/function的defer

```
func SomeMethod() err error{
    defer func(){
        if err != nil {
            utils.Logger.Error("packageName.StructName.MethodName : %s", error.Error())
        }else{
            utils.Logger.Debug("packageName.StructName.MethodName : success")
        }
    }()
}
```
