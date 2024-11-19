// Code generated by makefile. DO NOT EDIT.

package v1

func GetAllOperations() []string {
	return []string{
		OperationAuthCaptcha,                //"/api.core.v1.Auth/Captcha"
		OperationAuthLogin,                  //"/api.core.v1.Auth/Login"
		OperationAuthLogout,                 //"/api.core.v1.Auth/Logout"
		OperationAuthRegister,               //"/api.core.v1.Auth/Register"
		OperationAuthRouters,                //"/api.core.v1.Auth/Routers"
		OperationAuthUserInfo,               //"/api.core.v1.Auth/UserInfo"
		OperationSysDictCreateSysDictData,   //"/api.core.v1.SysDict/CreateSysDictData"
		OperationSysDictCreateSysDictType,   //"/api.core.v1.SysDict/CreateSysDictType"
		OperationSysDictDeleteSysDictData,   //"/api.core.v1.SysDict/DeleteSysDictData"
		OperationSysDictDeleteSysDictType,   //"/api.core.v1.SysDict/DeleteSysDictType"
		OperationSysDictGetSysDictData,      //"/api.core.v1.SysDict/GetSysDictData"
		OperationSysDictGetSysDictType,      //"/api.core.v1.SysDict/GetSysDictType"
		OperationSysDictListSysDictData,     //"/api.core.v1.SysDict/ListSysDictData"
		OperationSysDictListSysDictType,     //"/api.core.v1.SysDict/ListSysDictType"
		OperationSysDictOptionSelectType,    //"/api.core.v1.SysDict/OptionSelectType"
		OperationSysDictRefreshCacheSysDict, //"/api.core.v1.SysDict/RefreshCacheSysDict"
		OperationSysDictSysDictDataType,     //"/api.core.v1.SysDict/SysDictDataType"
		OperationSysDictUpdateSysDictData,   //"/api.core.v1.SysDict/UpdateSysDictData"
		OperationSysDictUpdateSysDictType,   //"/api.core.v1.SysDict/UpdateSysDictType"
		OperationSysRoleAllocatedList,       //"/api.core.v1.SysRole/AllocatedList"
		OperationSysRoleAuthUserCancel,      //"/api.core.v1.SysRole/AuthUserCancel"
		OperationSysRoleAuthUserCancelAll,   //"/api.core.v1.SysRole/AuthUserCancelAll"
		OperationSysRoleAuthUserSelectAll,   //"/api.core.v1.SysRole/AuthUserSelectAll"
		OperationSysRoleChangeStatusSysRole, //"/api.core.v1.SysRole/ChangeStatusSysRole"
		OperationSysRoleCreateSysRole,       //"/api.core.v1.SysRole/CreateSysRole"
		OperationSysRoleDataScopeSysRole,    //"/api.core.v1.SysRole/DataScopeSysRole"
		OperationSysRoleDeleteSysRole,       //"/api.core.v1.SysRole/DeleteSysRole"
		OperationSysRoleGetSysRole,          //"/api.core.v1.SysRole/GetSysRole"
		OperationSysRoleListSysRole,         //"/api.core.v1.SysRole/ListSysRole"
		OperationSysRoleUnAllocatedList,     //"/api.core.v1.SysRole/UnAllocatedList"
		OperationSysRoleUpdateSysRole,       //"/api.core.v1.SysRole/UpdateSysRole"
		OperationSysMenuCreateSysMenu,       //"/api.core.v1.SysMenu/CreateSysMenu"
		OperationSysMenuDeleteSysMenu,       //"/api.core.v1.SysMenu/DeleteSysMenu"
		OperationSysMenuGetSysMenu,          //"/api.core.v1.SysMenu/GetSysMenu"
		OperationSysMenuGetSysRoleMenu,      //"/api.core.v1.SysMenu/GetSysRoleMenu"
		OperationSysMenuGetTreeSelect,       //"/api.core.v1.SysMenu/GetTreeSelect"
		OperationSysMenuListSysMenu,         //"/api.core.v1.SysMenu/ListSysMenu"
		OperationSysMenuUpdateSysMenu,       //"/api.core.v1.SysMenu/UpdateSysMenu"
		OperationSysUserAuthRoleSysUser,     //"/api.core.v1.SysUser/AuthRoleSysUser"
		OperationSysUserCreateSysUser,       //"/api.core.v1.SysUser/CreateSysUser"
		OperationSysUserDeleteSysUser,       //"/api.core.v1.SysUser/DeleteSysUser"
		OperationSysUserGetAuthRoleSysUser,  //"/api.core.v1.SysUser/GetAuthRoleSysUser"
		OperationSysUserGetOtherInfo,        //"/api.core.v1.SysUser/GetOtherInfo"
		OperationSysUserGetSysUser,          //"/api.core.v1.SysUser/GetSysUser"
		OperationSysUserListSysUser,         //"/api.core.v1.SysUser/ListSysUser"
		OperationSysUserProfile,             //"/api.core.v1.SysUser/Profile"
		OperationSysUserResetPwd,            //"/api.core.v1.SysUser/ResetPwd"
		OperationSysUserSaveSysUser,         //"/api.core.v1.SysUser/SaveSysUser"
		OperationSysUserUpdatePassword,      //"/api.core.v1.SysUser/UpdatePassword"
		OperationSysUserUpdateProfile,       //"/api.core.v1.SysUser/UpdateProfile"
		OperationSysUserUpdateSysUser,       //"/api.core.v1.SysUser/UpdateSysUser"
		OperationSysDeptCreateSysDept,       //"/api.core.v1.SysDept/CreateSysDept"
		OperationSysDeptDeleteSysDept,       //"/api.core.v1.SysDept/DeleteSysDept"
		OperationSysDeptDeptTree,            //"/api.core.v1.SysDept/DeptTree"
		OperationSysDeptExcludeDept,         //"/api.core.v1.SysDept/ExcludeDept"
		OperationSysDeptGetSysDept,          //"/api.core.v1.SysDept/GetSysDept"
		OperationSysDeptGetSysRoleDept,      //"/api.core.v1.SysDept/GetSysRoleDept"
		OperationSysDeptListSysDept,         //"/api.core.v1.SysDept/ListSysDept"
		OperationSysDeptUpdateSysDept,       //"/api.core.v1.SysDept/UpdateSysDept"
		OperationSysPostCreateSysPost,       //"/api.core.v1.SysPost/CreateSysPost"
		OperationSysPostDeleteSysPost,       //"/api.core.v1.SysPost/DeleteSysPost"
		OperationSysPostGetSysPost,          //"/api.core.v1.SysPost/GetSysPost"
		OperationSysPostListSysPost,         //"/api.core.v1.SysPost/ListSysPost"
		OperationSysPostUpdateSysPost,       //"/api.core.v1.SysPost/UpdateSysPost"
		OperationSysConfigConfigByKey,       //"/api.core.v1.SysConfig/ConfigByKey"
		OperationSysConfigCreateSysConfig,   //"/api.core.v1.SysConfig/CreateSysConfig"
		OperationSysConfigListSysConfig,     //"/api.core.v1.SysConfig/ListSysConfig"
		OperationMonitorMonitorServer,       //"/api.core.v1.Monitor/MonitorServer"
	}
}
