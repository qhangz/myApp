# api说明
```text
user:
login：http://localhost/api/user/login
register：/api/user/register
userLisr: /api/user/list

```
# 接口返回参数说明
```json
// register
register sucess
{
    "code": "200",
    "msg": "register success"
}
register failed
{
    "code": "400",
    "error": "error message"
}
// login
login success
{
    "code": "200",
    "data": {
        "token": "1",
        "userInfo": {
            "age": 0,
            "avatar_image": "",
            "created_at": "2023-12-01T22:49:05+08:00",
            "summary": "",
            "username": "HANG"
        }
    },
    "msg": "登录成功"
}
login failed
{
    "code": "400",
    "error": "error message"
}

```
返回的msg
```text
register fail:
register success ;StatusOK
invalid registration information:数据验证失败
user already exists
email already exists
email format is wrong 

login fail:
invalid login information
user not exit
```