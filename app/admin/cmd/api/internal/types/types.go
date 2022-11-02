// Code generated by goctl. DO NOT EDIT.
package types

// The basic response without data | 基础不带数据信息
// swagger:response BaseMsgResp
type BaseMsgResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// The SimpleMsgResp | 简单信息Resp
// swagger:response SimpleMsgResp
type SimpleMsgResp struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageReq
type PageReq struct {
	// Page number | 第几页
	// Minimum: 1
	PageNo int `json:"pageNo,optional,default=1" validate:"number,min=1"`
	// Page size | 单页数据行数
	// Maximum: 100
	// Default: 10
	PageSize int `json:"pageSize,optional,default=10," validate:"number,max=100"`
}

// Basic id request | 基础id参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate:"number"`
}

// Basic Get id request | 基础id参数请求
// swagger:parameters IDParamReq ID
type IDParamReq struct {
	// ID
	// Required: true
	ID uint64 `form:"id" validate:"number"`
}

// Page information return ｜ 分页信息返回
// swagger:model Pagination
type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// Basic id request | 基础id参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	ID uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// UUID
	// Required: true
	// Max length: 36
	UUID string `json:"UUID" validate:"len=36"`
}

// The base response data | 基础信息
// swagger:model BaseInfo
type BaseInfo struct {
	// ID
	ID uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt"`
	// Delete date | 删除日期
	DeletedAt int64 `json:"deletedAt"`
}

// The request params of setting boolean status | 设置状态参数
// swagger:model SetBooleanStatusReq
type SetBooleanStatusReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate:"number"`
	// Status code | 状态码
	// Required: true
	Status uint8 `json:"status" validate:"number"`
}

// Role meta info | 角色meta数据
// swagger:model RoleMetaInfo
type RoleMetaInfo struct {
	// Role name | 角色名
	RoleName string `json:"roleName"`
	// Role value | 角色值
	Value string `json:"value"`
}

// swagger:model RoleInfo
type RoleInfo struct {
	// Role ID | 角色 ID
	ID uint64 `json:"id,optional,default=0"`
	// Role Name | 角色名
	Name string `json:"name"`
	// Role value | 角色值
	Value string `json:"value"`
	// Role's default page | 角色默认管理页面
	DefaultRouter string `json:"defaultRouter"`
	// Role status | 角色状态
	Status uint8 `json:"status,default=0"`
	// Role remark | 角色备注
	Remark string `json:"remark,default=''"`
	// Role's sorting number | 角色排序
	OrderNo uint32 `json:"orderNo"`
}

// swagger:model CreateRoleReq
type CreateRoleReq struct {
	// Role Name | 角色名
	// Required : true
	// Max length: 20
	// Example: "admin"
	Name string `json:"name" validate:"max=20"`
	// Role value | 角色值
	// Required : true
	// Max length: 10
	// Example: "admin"
	Value string `json:"value" validate:"max=10"`
	// Role's default page | 角色默认管理页面
	// Max length: 20
	// Example: "/dashboard"
	DefaultRouter string `json:"defaultRouter,optional,default='dashboard'" validate:"max=50"`
	// Role status | 角色状态
	// Maximum: 1
	// Example: 0
	Status uint8 `json:"status,optional,default=0" validate:"number,oneof=0 1"`
	// Role remark | 角色备注
	// Max length: 200
	// Example: "admin remark"
	Remark string `json:"remark,optional,default=''" validate:"omitempty,max=200"`
	// Role's sorting number | 角色排序
	// Maximum: 1000
	// Example: 0
	OrderNo uint32 `json:"orderNo,optional,default=0" validate:"number,max=1000"`
}

// swagger:model UpdateRoleReq
type UpdateRoleReq struct {
	// Role ID | 角色 ID
	// Required : true
	// Default: 1
	ID uint64 `json:"id" validate:"number,max=1000"`
	// Role Name | 角色名
	// Max length: 20
	Name *string `json:"name,optional" validate:"omitempty,max=20"`
	// Role value | 角色值
	// Max length: 10
	Value *string `json:"value,optional" validate:"omitempty,max=10"`
	// Role's default page | 角色默认管理页面
	// Max length: 20
	DefaultRouter *string `json:"defaultRouter,optional" validate:"omitempty,max=50"`
	// Role status | 角色状态
	// Maximum: 10
	Status *uint8 `json:"status,optional" validate:"omitempty,number,max=10"`
	// Role remark | 角色备注
	// Max length: 200
	Remark *string `json:"remark,default=''" validate:"omitempty,max=200"`
	// Role's sorting number | 角色排序
	// Required : true
	// Maximum: 1000
	OrderNo *uint32 `json:"orderNo" validate:"omitempty,number,max=1000"`
}

// The response data of role list | 角色列表数据
// swagger:model RoleListResp
type RoleListResp struct {
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The role list data | 角色列表数据
	// in: body
	List []*RoleInfo `json:"data"`
}

// The request params of setting role status | 设置角色状态参数
// swagger:model SetStatusReq
type SetStatusReq struct {
	// ID
	// Required: true
	// Maximum: 1000
	ID uint64 `json:"id" validate:"number,max=1000"`
	// Status code | 状态码
	// Required: true
	// Maximum: 1
	Status uint8 `json:"status" validate:"number,oneof=0 1"`
}

// The response data of captcha | 验证码返回数据
// swagger:response CaptchaInfoResp
type CaptchaInfoResp struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

// login request | 登录参数
// swagger:model LoginReq
type LoginReq struct {
	// UserName | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha Id | 验证码编号
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId"  validate:"len=20"`
	// The Captcha | 验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"len=5"`
}

// The login response data | 登录返回数据
// swagger:model LoginResp
type LoginResp struct {
	// 用户信息
	//in: body
	User *UserInfo `json:"user"`
	// User's role information| 用户的角色信息
	// in: body
	Role *RoleMetaInfo `json:"role"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	ExpiredAt int64 `json:"expiredAt"`
}

// The profile response data | 个人信息返回数据
// swagger:response ProfileResp
type ProfileResp struct {
	User *UserInfo `json:"user"`
}

// The profile update request data | 个人信息update请求参数
// swagger:model UpdateProfileReq
type UpdateProfileReq struct {
	// user's nickname | 用户的昵称
	// Max length: 10
	Nickname *string `json:"nickname,optional" validate:"alphanumunicode,max=10"`
	// The user's avatar path | 用户的头像路径
	// Max length: 512
	Avatar *string `json:"avatar,optional" validate:"max=512"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile *string `json:"mobile,optional" validate:"numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email *string `json:"email" validate:"email,max=100"`
}

// register request | 注册参数
// swagger:model RegisterReq
type RegisterReq struct {
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha Id which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId" validate:"len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	// Default: 00000
	Captcha string `json:"captcha" validate:"len=5"`
	// The user's email address | 用户的邮箱
	// Required: true
	// Max length: 100
	Email string `json:"email" validate:"email,max=100"`
}

// change user's password request | 修改密码请求参数
// swagger:model ChangePasswordReq
type ChangePasswordReq struct {
	// User's old password | 用户旧密码
	// Required: true
	// Max length: 30
	OldPassword string `json:"oldPassword" validate:"max=30"`
	// User's new password | 用户新密码
	// Required: true
	// Max length: 30
	NewPassword string `json:"newPassword" validate:"max=30"`
}

// The response data of user's information | 用户信息返回数据
// swagger:response UserInfoResp
type UserInfoResp struct {
	// in: body
	User *UserInfo `json:"user"`
}

//	用户信息
//
// swagger:model UserInfo
type UserInfo struct {
	// User's id | 用户Id
	ID uint64 `json:"id"`
	// User's UUID | 用户的UUID
	UUID string `json:"UUID"`
	// User Name | 用户名
	Username string `json:"username"`
	// User's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// User's mobile phone number | 用户的手机号码
	Mobile string `json:"mobile"`
	// User's role id | 用户的角色Id
	RoleID uint64 `json:"roleId"`
	// The user's email address | 用户的邮箱
	Email string `json:"email"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// The user's layout mode | 用户的布局
	SideMode string `json:"sideMode"`
	// The user's status | 用户状态
	// 0 正常 1 拉黑
	Status    uint8 `json:"status"`
	CreatedAt int64 `json:"createAt"`
	UpdatedAt int64 `json:"updateAt"`
}

// The response data of user's basic information | 用户基本信息返回数据
// swagger:model GetUserInfoResp
type GetUserInfoResp struct {
	// user
	// in: body
	User *UserInfo `json:"user"`
	// User's role information| 用户的角色Meta信息
	// in: body
	Roles *RoleMetaInfo `json:"roles"`
}

// The response data of user list | 用户列表数据
// swagger:model UserListResp
type UserListResp struct {
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The user list data | 用户列表数据
	// in: body
	List []*UserInfo `json:"list"`
}

// The permission code for front end permission control | 权限码： 用于前端权限控制
// swagger:response PermCodeResp
type PermCodeResp struct {
	// Permission code data | 权限码数据
	Data []string `json:"data"`
}

// Create user information request | 创建新用户信息
// swagger:model CreateUserReq
type CreateUserReq struct {
	// User Name | 用户名
	// Max length: 20
	// Required: true
	// Example: admin
	Username string `json:"username" validate:"alphanum,max=20"`
	// User's nickname | 用户的昵称
	// Max length: 20
	// Required: true
	// Example: admin nickname
	Nickname string `json:"nickname" validate:"alphanumunicode,max=20"`
	// Password | 密码
	// Max length: 30
	// Min length: 6
	// Required: true
	// Example: 123456
	Password string `json:"password" validate:"max=30,min=6"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	// Required: true
	// Example: admin@gmail.com
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	// Example: 13888888888
	Mobile *string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// User's role id | 用户的角色Id
	// Minimum: 1
	// Maximum: 10
	// Example: 1
	RoleID *uint64 `json:"roleId,optional" validate:"omitempty,number,max=10"`
	// The user's avatar path | 用户的头像路径
	// Example: https://localhost/static/1.png
	Avatar *string `json:"avatar,optional" validate:"omitempty,url"`
	// The user's status | 用户状态
	// 0 normal, 1 ban | 0 正常 1 拉黑
	// Maximum: 0
	Status *uint8 `json:"status,optional" validate:"omitempty,number,oneof=0 1"`
}

// Update user information request | 更新用户信息
// swagger:model UpdateUserReq
type UpdateUserReq struct {
	// User's id | 用户Id
	// Minimum: 1
	// Required: true
	// Example: 1
	ID uint64 `json:"id,optional" validate:"omitempty,number"`
	// User Name | 用户名
	// Max length: 20
	// Example: admin
	Username *string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// User's nickname | 用户的昵称
	// Max length: 10
	// Example: admin nickname
	Nickname *string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=20"`
	// Password | 密码
	// Max length: 30
	// Min length: 6
	// Example: 123456
	Password *string `json:"password,optional" validate:"omitempty,max=30,min=6"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	// Example: 13888888888
	Mobile *string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// User's role id | 用户的角色Id
	// Maximum: 1000
	// Example: 1
	RoleID *uint64 `json:"roleId,optional" validate:"omitempty,number,max=1000"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	// Example: admin@gmail.com
	Email *string `json:"email,optional" validate:"omitempty,email,max=100"`
	// The user's avatar path | 用户的头像路径
	// Example: https://localhost/static/1.png
	Avatar *string `json:"avatar,optional" validate:"omitempty,url"`
	// The user's status | 用户状态
	// 0 normal, 1 ban | 0 正常 1 拉黑
	// Maximum: 2
	Status *uint8 `json:"status,optional" validate:"omitempty,number,oneof=0 1"`
}

// Get user list request | 获取用户列表请求参数
// swagger:model GetUserListReq
type GetUserListReq struct {
	PageReq
	// User Name | 用户名
	// Max length: 20
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// User's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
	// User's role ID | 用户的角色Id
	// Maximum: 1000
	RoleId uint64 `json:"roleId,optional" validate:"omitempty,number,max=1000"`
}

// The response data of API information | API信息
// swagger:model ApiInfo
type ApiInfo struct {
	// ID
	ID uint64 `json:"id"`
	// API path | API路径
	Path string `json:"path"`
	// API Description | API 描述
	Description string `json:"description"`
	// API group | API分组
	Group string `json:"group"`
	// API request method e.g. POST | API请求类型 如POST
	Method string `json:"method"`
	// created time | 创建时间
	CreatedAt int64 `json:"createdAt"`
}

// Create API information request | 创建或更新API信息
// swagger:model CreateApiReq
type CreateApiReq struct {
	// API path | API路径
	// Required: true
	// Min length: 1
	// Max length: 50
	Path string `json:"path" validate:"min=1,max=50"`
	// API Description | API 描述
	// Required: true
	// Max length: 50
	Description string `json:"description" validate:"max=50"`
	// API group | API分组
	// Require: true
	// Min length: 1
	// Max length: 10
	Group string `json:"group" validate:"alphanum,min=1,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Required: true
	// Example: POST
	// Min length: 3
	// Max length: 4
	Method string `json:"method" validate:"uppercase,min=3,max=4"`
}

// Update API information request | 创建或更新API信息
// swagger:model UpdateApiReq
type UpdateApiReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate:"number"`
	// API path | API路径
	// Min length: 1
	// Max length: 50
	Path *string `json:"path,optional" validate:"omitempty,min=1,max=50"`
	// API Description | API 描述
	// Max length: 50
	Description *string `json:"description,optional" validate:"omitempty,max=50"`
	// API group | API分组
	// Min length: 1
	// Max length: 10
	Group *string `json:"group,optional" validate:"omitempty,alphanum,min=1,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Min length: 3
	// Max length: 4
	Method *string `json:"method,optional" validate:"omitempty,uppercase,min=3,max=4"`
}

// The response data of API list | API列表数据
// swagger:model ApiListResp
type ApiListResp struct {
	// Page information | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The API list data | API列表数据
	// in: body
	List []*ApiInfo `json:"data"`
}

// Get API list request params | API列表请求参数
// swagger:model ApiListReq
type ApiListReq struct {
	PageReq
	// API path | API路径
	// Max length: 100
	Path *string `json:"path,optional" validate:"omitempty,max=100"`
	// API Description | API 描述
	// Max length: 50
	Description *string `json:"description,optional" validate:"omitempty,max=50"`
	// API group | API分组
	// Max length: 10
	Group *string `json:"group,optional" validate:"omitempty,alphanum,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Max length: 4
	Method *string `json:"method,optional" validate:"omitempty,uppercase,max=4"`
}

// The response data of api authorization | API授权数据
// swagger:model ApiAuthorityInfo
type ApiAuthorityInfo struct {
	// API path | API 路径
	Path string `json:"path"`
	// API method | API请求方法
	Method string `json:"method"`
}

// Create api authorization information request | 创建API授权信息
// swagger:model CreateApiAuthorityReq
type CreateApiAuthorityReq struct {
	// Role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// API authorization list | API授权列表数据
	// Required: true
	Data []*ApiAuthorityInfo `json:"data"`
}

// Update api authorization information request | 更新API授权信息
// swagger:model UpdateApiAuthorityReq
type UpdateApiAuthorityReq struct {
	// Role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// API authorization list | API授权列表数据
	// Required: true
	Data []*ApiAuthorityInfo `json:"data"`
}

// The response data of api authorization list | API授权列表数据
// swagger:model ApiAuthorityListResp
type ApiAuthorityListResp struct {
	// Pageinfo | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The api authorization list data | API授权列表数据
	// in: body
	List []*ApiAuthorityInfo `json:"list"`
}

// Create menu authorization information request params | 创建或更新菜单授权信息参数
// swagger:model CreateMenuAuthorityReq
type CreateMenuAuthorityReq struct {
	// role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// menu ID array | 菜单ID数组
	// Required: true
	MenuIds []uint64 `json:"menuIds"`
}

// Update menu authorization information request params | 创建或更新菜单授权信息参数
// swagger:model UpdateMenuAuthorityReq
type UpdateMenuAuthorityReq struct {
	// role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// menu ID array | 菜单ID数组
	// Required: true
	MenuIds []uint64 `json:"menuIds"`
}

// Menu authorization information request params | 菜单授权信息参数
// swagger:response MenuAuthorityInfoResp
type MenuAuthorityInfoResp struct {
	// role ID | 角色ID
	RoleId uint64 `json:"roleId"`
	// menu ID array | 菜单ID数组
	MenuIds []uint64 `json:"menuIds"`
}

// The response data of dictionary information | 字典信息
// swagger:model DictionaryInfo
type DictionaryInfo struct {
	// ID
	ID uint64 `json:"id"`
	// Dictionary title | 字典显示名称
	Title string `json:"title"`
	// Dictionary name | 字典名称
	Name string `json:"name"`
	// Dictionary status | 字典状态
	Status bool `json:"status"`
	// Dictionary description | 字典描述
	Description string `json:"description"`
	// Create time | 创建时间
	CreatedAt int64 `json:"createdAt"`
}

// Create dictionary information request | 创建或更新字典信息请求
// swagger:model CreateDictionaryReq
type CreateDictionaryReq struct {
	// Dictionary title | 字典显示名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Title string `json:"title" validate:"min=1,max=50"`
	// Dictionary name | 字典名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Name string `json:"name" validate:"min=1,max=50"`
	// Dictionary status | 字典状态
	// Required: true
	// Example: 0
	Status uint8 `json:"status" validator:"number,oneof=0 1"`
	// Dictionary description | 字典描述
	// Required: true
	// Max length: 50
	// Example: dictionary description
	Description string `json:"description" validate:"max=50"`
}

// Update dictionary information request | 创建或更新字典信息请求
// swagger:model UpdateDictionaryReq
type UpdateDictionaryReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Dictionary title | 字典显示名称
	// Min length: 1
	// Max length: 50
	Title *string `json:"title,optional" validate:"omitempty,min=1,max=50"`
	// Dictionary name | 字典名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Name *string `json:"name,optional" validate:"omitempty,min=1,max=50"`
	// Dictionary status | 字典状态
	Status *uint8 `json:"status,optional" validator:"omitempty,number,oneof=0 1"`
	// Dictionary description | 字典描述
	// Max length: 50
	Description *string `json:"description,optional" validate:"omitempty,max=50"`
}

// The response data of dictionary list | 字典列表数据
// swagger:model DictionaryListResp
type DictionaryListResp struct {
	// PageInfo | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The dictionary list data | 字典列表数据
	// in: body
	List []*DictionaryInfo `json:"list"`
}

// Get dictionary list request params | 字典列表请求参数
// swagger:model DictionaryListReq
type DictionaryListReq struct {
	PageReq
	// Dictionary title | 字典显示名称
	// Max length: 50
	Title *string `json:"title,optional" validate:"omitempty,max=50"`
	// Dictionary name | 字典名称
	// Max length: 50
	Name *string `json:"name,optional" validate:"omitempty,max=50"`
}

// The response data of dictionary information | 字典信息
// swagger:model DictionaryDetailInfo
// swagger:model DictionaryDetailInfo
type DictionaryDetailInfo struct {
	// ID
	ID uint64 `json:"id"`
	// Dictionary title | 字典显示名称
	Title string `json:"title"`
	// Key name | 键
	Key string `json:"key"`
	// Value | 值
	Value string `json:"value"`
	// Status | 状态
	Status uint8 `json:"status"`
	// Create time | 创建时间
	CreatedAt int64 `json:"createdAt"`
}

// The response data of dictionary KV list | 字典值的列表数据
// swagger:model DictionaryDetailListResp
type DictionaryDetailListResp struct {
	// PageInfo | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The dictionary list data | 字典列表数据
	// in: body
	List []*DictionaryDetailInfo `json:"list"`
}

// Create KV information request | 创建或更新字典键值信息请求
// swagger:model CreateDictionaryDetailReq
type CreateDictionaryDetailReq struct {
	// Detail title | 字典值显示名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Title string `json:"title" validate:"min=1,max=50"`
	// Detail key | 键
	// Required: true
	// Min length: 1
	// Max length: 50
	Key string `json:"key" validate:"min=1,max=50"`
	// Detail value | 值
	// Required: true
	Value string `json:"value"`
	// Status | 状态
	// Required: true
	Status uint8 `json:"status" validate:"number,oneof=0 1"`
	// Parent ID | 所属字典ID
	// Required: true
	ParentID uint64 `json:"parentId" validate:"number"`
}

// Update dictionary KV information request | 创建或更新字典键值信息请求
// swagger:model UpdateDictionaryDetailReq
type UpdateDictionaryDetailReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate:"number"`
	// Detail title | 字典值显示名称
	// Min length: 1
	// Max length: 50
	Title *string `json:"title,optional" validate:"omitempty,min=1,max=50"`
	// Detail key | 键
	// Min length: 1
	// Max length: 50
	Key *string `json:"key,optional" validate:"omitempty,min=1,max=50"`
	// Detail value | 值
	Value *string `json:"value,optional"`
	// Status | 状态
	Status *uint8 `json:"status,optional" validate:"omitempty,number,oneof=0 1"`
	// Parent ID | 所属字典ID
	ParentID *uint64 `json:"parentId,optional" validate:"omitempty,number"`
}

// Get dictionary detail list by dictionary name request | 根据字典名称获取对应键值请求
// swagger:model DictionaryDetailReq
type DictionaryDetailReq struct {
	// Dictionary name | 字典名
	Name string `json:"name"`
}

// The response data of oauth provider information | 提供者信息
// swagger:model OauthProviderInfo
type OauthProviderInfo struct {
	// ID
	ID uint64 `json:"id"`
	// Provider name | 提供商名字
	Name string `json:"name"`
	// Client ID | 客户端ID
	ClientID string `json:"clientID"`
	// Client secret | 客户端密码
	ClientSecret string `json:"clientSecret"`
	// Redirect URL | 跳转URL
	RedirectURL string `json:"redirectURL"`
	// Scopes | 范围
	Scopes string `json:"scopes"`
	// Auth URL | 鉴权URL
	AuthURL string `json:"authURL"`
	// Token URL | 获取 Token 的网址
	TokenURL string `json:"tokenURL"`
	// Auth Style | 鉴权方式, 0 表示自动检测
	AuthStyle uint8 `json:"authStyle"`
	// Get User information URL | 获取用户信息地址
	InfoURL string `json:"infoURL"`
	// Created time | 创建时间
	CreatedAt int64 `json:"createdAt"`
}

// Create provider information request | 创建提供商信息
// swagger:model CreateProviderReq
type CreateProviderReq struct {
	// Provider name | 提供商名字
	// Required: true
	// Min length: 1
	// Max length: 50
	Name string `json:"name" validate:"min=1,max=50"`
	// Client ID | 客户端ID
	// Required: true
	// Max length: 100
	ClientID string `json:"clientID" validate:"max=100"`
	// Client secret | 客户端密码
	// Require: true
	// Min length: 1
	// Max length: 100
	ClientSecret string `json:"clientSecret" validate:"min=1,max=100"`
	// Redirect URL | 跳转URL
	// Required: true
	// Max length: 200
	RedirectURL string `json:"redirectURL" validate:"max=200"`
	// Scopes | 范围
	// Required: true
	// Max length: 200
	Scopes string `json:"scopes" validate:"max=200"`
	// Auth URL | 鉴权URL
	// Required: true
	// Max length: 200
	AuthURL string `json:"authURL" validate:"max=200"`
	// Token URL | 获取 Token 的网址
	// Required: true
	// Max length: 200
	TokenURL string `json:"tokenURL" validate:"max=200"`
	// Auth Style is specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect. | 鉴权方式, 0 表示自动检测
	// 0 auto detect 1 client ID and client secret 2 username and password
	// Required: true
	// Example: 0
	AuthStyle uint8 `json:"authStyle" validate:"number,oneof=0 1 2"`
	// Get User information URL | 获取用户信息地址
	// Required: true
	// Max length: 200
	InfoURL string `json:"infoURL" validate:"max=200"`
}

// Update provider information request | 更新提供商信息
// swagger:model UpdateProviderReq
type UpdateProviderReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate:"number"`
	// Provider name | 提供商名字
	// Required: true
	// Min length: 1
	// Max length: 50
	Name *string `json:"name,optional" validate:"omitempty,min=1,max=50"`
	// Client ID | 客户端ID
	// Required: true
	// Max length: 100
	ClientID *string `json:"clientID,optional" validate:"omitempty,max=100"`
	// Client secret | 客户端密码
	// Require: true
	// Min length: 1
	// Max length: 100
	ClientSecret *string `json:"clientSecret,optional" validate:"omitempty,min=1,max=100"`
	// Redirect URL | 跳转URL
	// Required: true
	// Max length: 200
	RedirectURL *string `json:"redirectURL,optional" validate:"omitempty,max=200"`
	// Scopes | 范围
	// Required: true
	// Max length: 200
	Scopes *string `json:"scopes,optional" validate:"omitempty,max=200"`
	// Auth URL | 鉴权URL
	// Required: true
	// Max length: 200
	AuthURL *string `json:"authURL,optional" validate:"omitempty,max=200"`
	// Token URL | 获取 Token 的网址
	// Required: true
	// Max length: 200
	TokenURL *string `json:"tokenURL,optional" validate:"omitempty,max=200"`
	// Auth Style is specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect. | 鉴权方式, 0 表示自动检测
	// 0 auto detect 1 client ID and client secret 2 username and password
	// Required: true
	// Example: 0
	AuthStyle *uint8 `json:"authStyle,optional" validate:"omitempty,number,oneof=0 1 2"`
	// Get User information URL | 获取用户信息地址
	// Required: true
	// Max length: 200
	InfoURL *string `json:"infoURL,optional" validate:"omitempty,max=200"`
}

// The response data of provider list | 提供商列表数据
// swagger:model ProviderListResp
type ProviderListResp struct {
	//Page information | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The provider list data | 提供商列表数据
	// in: body
	List []*OauthProviderInfo `json:"list"`
}

// Oauth log in request | Oauth 登录请求
// swagger:model OauthLoginReq
type OauthLoginReq struct {
	// State code to avoid hack | 状态码，请求前后相同避免安全问题
	// Required: true
	// Max Length: 30
	State string `json:"state" validate:"max=30"`
	// Provider name | 提供商名字
	// Required: true
	// Max Length: 40
	// Example: google
	Provider string `json:"provider" validate:"max=40"`
}

// Oauth Callback in Req | Oauth 回调请求
// swagger:parameters OauthCallbackParamReq OauthCallback
type OauthCallbackParamReq struct {
	// Required: true
	// Max Length: 30
	State string `form:"state" validate:"max=30"`
	// Required: true
	// Max Length: 40
	Code string `form:"code" validate:"max=100"`
}

// Redirect response | 跳转网址
// swagger:response RedirectResp
type RedirectResp struct {
	URL string `json:"URL"`
}

// The oauth callback response data | Oauth回调数据
// swagger:response CallbackResp
type CallbackResp struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// User's role information| 用户的角色信息
	// in: body
	Role *RoleMetaInfo `json:"role"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	ExpiredAt uint64 `json:"expiredAt"`
}

// The response data of Token information | Token信息
// swagger:model TokenInfo
type TokenInfo struct {
	// ID
	ID uint64 `json:"id"`
	// User's UUID | 用户的UUID
	UUID string `json:"UUID"`
	// Token string | Token 字符串
	Token string `json:"token"`
	// Log in source such as github | Token 来源 （本地为core, 第三方如github等）
	Source string `json:"source"`
	// JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	Status uint8 `json:"status"`
	// Create time | 创建时间
	CreatedAt int64 `json:"createdAt"`
	// Expire time | 过期时间
	ExpiredAt int64 `json:"expiredAt"`
}

// Create token information request | 创建token信息
// swagger:model CreateTokenReq
type CreateTokenReq struct {
	// User's UUID | 用户的UUID
	// Required: true
	// Max Length: 36
	UUID string `json:"UUID" validate:"max=36"`
	// Token string | Token 字符串
	// Required: true
	Token string `json:"token"`
	// Log in source such as github | Token 来源 （本地为core, 第三方如github等）
	// Required: true
	// Max Length: 50
	Source string `json:"source" validate:"max=50"`
	// JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	// Required: true
	// Example: 1
	Status uint8 `json:"status,optional,default=0" validate:"number,oneof=0 1"`
	// Expire time | 过期时间
	// Required: true
	ExpireAt int64 `json:"expireAt" validate:"number"`
}

// Update token information request | 创建或更新token信息
// swagger:model UpdateTokenReq
type UpdateTokenReq struct {
	// ID
	// Required: true
	ID uint64 `json:"id" validate="number"`
	// User's UUID | 用户的UUID
	// Max Length: 36
	UUID *string `json:"UUID,optional" validate:"omitempty,max=36"`
	// Token string | Token 字符串
	Token *string `json:"token,optional" validate:"omitempty,max=50"`
	// Log in source such as github | Token 来源 （本地为core, 第三方如github等）
	// Max Length: 50
	Source *string `json:"source,optional" validate:"omitempty,max=50"`
	// JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	Status *uint8 `json:"status,optional" validate:"omitempty,number"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional" validate:"omitempty,number"`
	// Expire time | 过期时间
	ExpiredAt *int64 `json:"expiredAt,optional" validate:"omitempty,number"`
}

// The response data of Token list | Token列表数据
// swagger:model TokenListResp
type TokenListResp struct {
	// The Page information | 分页信息
	// in: body
	Pagination *Pagination `json:"pagination"`
	// The token list data | Token列表数据
	// in: body
	List []*TokenInfo `json:"list"`
}

// Get token list request params | token列表请求参数
// swagger:model TokenListReq
type TokenListReq struct {
	PageReq
	// User's UUID | 用户的UUID
	// Required: true
	// Max Length: 36
	UUID *string `json:"UUID,optional" validate:"omitempty,max=36"`
	// user's nickname | 用户的昵称
	// Required: true
	// Max length: 10
	Nickname *string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username *string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// The user's email address | 用户的邮箱
	// Required: true
	// Max length: 100
	Email *string `json:"email,optional" validate:"omitempty,email,max=100"`
}
