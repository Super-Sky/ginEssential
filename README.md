# ginEssential

`一个基于gin和gorm的开发示例。 初步完成其中用户模块及文章管理模块。`

###用户模块
#####注册：
`POST:/api/auth/register`
#####登录：
`POST:/api/auth/login`
#####用户信息
`POST:/api/auth/info`
###目录
#####新增目录
`POST:/categories`
#####修改目录
`PUT:/categories:id`
#####查看目录
`GET:/categories:id`
#####删除目录
`DELETE:/categories:id`
###文章
#####新增文章
`POST:/posts`
#####更新文章
`PUT:/posts:id`
#####查看文章
`GET:/posts:id`
#####删除文章
`DELETE:/posts:id`
#####分页
`POST:/posts/page/list`