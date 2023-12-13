# api说明
```text
1. user:
login：http://localhost/api/user/login
register：/api/user/register
userLisr: /api/user/list
userInfo: /api/user/info/:username
userInfo updata: /api/user/update/username
userInfo updata: /api/user/update/password

2. discuss
publish discuss: /api/discuss/publish
get discuss list: /api/discuss/list
get discuss info by discuss id from request: /api/discussinfo

3. comment
publish comment: /api/comment/publish
```
# 接口返回参数说明
1. user:
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

// user list
success
{
    "code": "200",
	"data": "userList",
}
failed
{
    "code": "400",
    "error": "error message"
}

// get user info by username
success
{
    "code": "200",
	"data": "userInfo",
}
failed
{
    "code": "400",
    "error": "error message"
}

// user info update (username,password,email,age,summary,avatar_image)
success
{
    "code": "200",
    "msg": "update success",
}

// user delete
success
{
    "code": "200",
    "msg": "delete success",
}

返回的msg
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

2. discuss:
```json
/api/discuss/publish (post)
{
    "code": "400",
    "msg": "success",
}

/api/discuss/list (get)

/api/discuss/discussinfo (get query)
{
    "code": "200",
    "data": {
        "ID": 1,
        "CreatedAt": "2023-12-12T16:38:17.2+08:00",
        "UpdatedAt": "2023-12-12T16:38:17.2+08:00",
        "DeletedAt": null,
        "author": "HANG",
        "title": "discuss",
        "summaty": "test discuss",
        "content": "just test discuss api",
        "like_num": 0,
        "view_num": 0,
        "Comment": [
            {
                "ID": 1,
                "CreatedAt": "2023-12-12T17:33:50.192+08:00",
                "UpdatedAt": "2023-12-12T17:33:50.192+08:00",
                "DeletedAt": null,
                "DiscussID": 1,
                "author": "hang",
                "content": "this is comment",
                "like_num": 0,
                "view_num": 0
            },
            {
                "ID": 3,
                "CreatedAt": "2023-12-12T17:35:00.281+08:00",
                "UpdatedAt": "2023-12-12T17:35:00.281+08:00",
                "DeletedAt": null,
                "DiscussID": 1,
                "author": "hang",
                "content": "this is comment2",
                "like_num": 0,
                "view_num": 0
            }
        ]
    }
}

```