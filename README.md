# go-atlas

## 平台简介

后端 使用go-kratos编写, 完成了基础的api

前端在RuoYi-Vue上简单修改

+ Vue2版本： [atlas-ui](https://github.com/ut-cloud/atlas-ui)
+ Vue3版本：[atlas-vue-ui](https://github.com/ut-cloud/atlas-vue3-ui)

预览地址：https://admin.go-atlas.cn

## 完成的功能

- 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
- 部门管理：配置系统组织机构，树结构展现支持数据权限。
- 岗位管理：配置系统用户所属担任职务。
- 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
- 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
- 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
- 角色切换，不同角色可以数据权限不一致。
- 数据权限：分全部权限，本部门权限，本部门及以下权限，自定义权限，本人权限 五种权限
- 部门切换：可以设置用户可以存在的多个部门，但是只能激活一个部门，可以切换；
- 权限管理: 由后端返回路由动态生成路由；前端按钮级权限统一由后端返回权限标志控制

## 代码生成

kratos模版代码做了部分修改，节省集成的时间

```shell
kratos new demo -r https://github.com/ut-cloud/kratos-layout.git
```