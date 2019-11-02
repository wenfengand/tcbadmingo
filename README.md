# tcb-admin-go

使用类似reset api的方式实现go语言的tcb admin。

目前只实现了简单的数据库增删改查，有用到时再解决文件、云函数等。
# 使用方法

## 新建`config.yaml`

```yaml
env: *******************
secretId: *********************
secretKey: ***********************
```

## 初始化
```go
yml_conf := tcb.Conf{}
getConf(filePath, &yml_conf)
TcbAdmin := tcb.New(yml_conf)
TcbAdmin.Debug(true)
```
## 增
```go
option := tcb.Option{}
option.Path = "/db/test"
option.Method = "post"

user := User{}
user.Name = "testetet"
user.Age  = 1010
option.Body  = tcb.Struct2Map(user)
ret := TcbAdmin.Post(option)
fmt.Println(ret)
```
返回id，字符串类型。

## 删
```go
option := tcb.Option{}
option.Path = "/db/test"
option.Method = "delete"
option.Params = map[string]interface{}{"_id":"4a741dc95dbd51d5006ea699315a850c"}
ret := TcbAdmin.Delete(option)
fmt.Println(ret)
```
返回是否成功，1为成功，0为失败
## 改

```go
option := tcb.Option{}
option.Path = "/db/test"
option.Method = "patch"
option.Params = map[string]interface{}{"_id":"9888d322-fe11-4aaa-af7f-bc91ea870f4c"}
option.Body  = map[string]interface{}{"$set":map[string]interface{}{"name":"gotest2"}}
ret := TcbAdmin.Patch(option)
fmt.Println(ret)
```
返回是否修改成功，1为成功，0为失败

## 查
```go
option := tcb.Option{}
option.Path = "/db/test"
option.Method = "get"
option.Params = map[string]interface{}{}

var users []User
option.Dst = &users
total := TcbAdmin.Get(option)
fmt.Printf("Total number %d", total)
fmt.Println(users)
fmt.Println(users[0].Age)
```

返回查询结果条数，int；同时可以通过`option.Dst`设置解析结果，直接映射为go对象类型。

## 开发
在`log`中是tcb_admin_node调试信息；`tcb_test`中是所有测试。
