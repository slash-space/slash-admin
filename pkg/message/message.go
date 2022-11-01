package message

const (

	// DatabaseError
	// normal database error
	DatabaseError string = "database error occur"

	// RedisError
	// normal redis error
	RedisError string = "redis error occur"

	// request error

	// ApiRequestFailed
	// EN: The interface request failed, please try again later!
	// ZH_CN: 请求出错，请稍候重试
	ApiRequestFailed string = "sys.api.apiRequestFailed"

	// CreateSuccess
	// EN: Create successfully
	// ZH_CN: 新建成功
	CreateSuccess string = "common.createSuccess"

	// CreateFailed
	// EN: Create failed
	// ZH_CN: 新建失败
	CreateFailed string = "common.createFailed"

	// UpdateSuccess
	// EN: Update successfully
	// ZH_CN: 更新成功
	UpdateSuccess string = "common.updateSuccess"

	// UpdateFailed
	// EN: Update failed
	// ZH_CN: 更新失败
	UpdateFailed string = "common.updateFailed"

	// DeleteSuccess
	// EN: Delete successfully
	// ZH_CN: 删除成功
	DeleteSuccess string = "common.deleteSuccess"

	// DeleteFailed
	// EN: Delete failed
	// ZH_CN: 删除失败
	DeleteFailed string = "common.deleteFailed"

	// GetInfoSuccess
	// EN: Get information Successfully
	// ZH_CN: 获取信息成功
	GetInfoSuccess string = "common.getInfoSuccess"

	// GetInfoFailed
	// EN: Get information Failed
	// ZH_CN: 获取信息失败
	GetInfoFailed string = "common.getInfoFailed"

	// TargetNotExist
	// EN: Target does not exist
	// ZH_CN: 目标不存在
	TargetNotExist string = "common.targetNotExist"

	// Success
	// EN: Successful
	// ZH_CN: 成功
	Success string = "common.successful"

	// Failed
	// EN: Failed
	// ZH_CN: 失败
	Failed string = "common.failed"

	// InitRunning
	// EN: The initialization is running...
	// ZH_CN: 正在初始化...
	InitRunning string = "sys.init.initializeIsRunning"

	// AlreadyInit
	// EN: The database had been initialized
	// ZH_CN: 数据库已被初始化
	AlreadyInit string = "sys.init.alreadyInit"

	// UserAlreadyExists
	// EN: Username or email address had been registered
	// ZH_CN: 用户名或者邮箱已被注册
	UserAlreadyExists string = "sys.login.signupUserExist"

	// UserNotExists
	// EN: User is not registered
	// ZH_CN: 用户不存在
	UserNotExists string = "sys.login.userNotExist"

	// CaptchaCreateFail
	// EN: Failed to create captcha
	// ZH_CN: 生成验证码失败
	CaptchaCreateFail string = "sys.captcha.createFail"

	// WrongCaptcha
	// EN: Wrong captcha
	// ZH_CN: 验证码错误
	WrongCaptcha string = "sys.login.wrongCaptcha"

	// WrongUsernameOrPassword
	// EN: Wrong username or password
	// ZH_CN: 用户名或密码错误
	WrongUsernameOrPassword string = "sys.login.wrongUsernameOrPassword"

	// menu service messages

	// ChildrenExistError
	// EN: Please delete menu's children first
	// ZH_CN: 请先删除子菜单
	ChildrenExistError string = "sys.menu.deleteChildrenDesc"

	// ParentNotExist
	// EN: The parent does not exist
	// ZH_CN: 父级不存在
	ParentNotExist string = "sys.menu.parentNotExist"

	// MenuNotExists
	// EN: Menu does not exist
	// ZH_CN: 菜单不存在
	MenuNotExists string = "sys.menu.menuNotExists"

	// MenuAlreadyExists
	// EN: Menu already exists
	// ZH_CN: 菜单已存在
	MenuAlreadyExists string = "sys.menu.menuAlreadyExists"

	// role service messages

	// DuplicateRoleValue
	// EN: Duplicate role value
	// ZH_CN: 角色值重复
	DuplicateRoleValue string = "sys.role.duplicateRoleValue"

	// UserExists
	// EN: Please delete users who belong to this role
	// ZH_CN: 请先删除该角色下的用户
	UserExists string = "sys.role.userExists"

	// RoleForbidden
	// EN: Your role is forbidden
	// ZH_CN: 您的角色已停用
	RoleForbidden string = "sys.role.roleForbidden"

	// RoleNotExists
	// EN: Role does not exist
	// ZH_CN: 角色不存在
	RoleNotExists string = "sys.role.roleNotExists"

	// WrongPassword
	// EN: Wrong Password
	// ZH_CN: 密码错误
	WrongPassword string = "sys.user.wrongPassword"

	// UserParamsError
	// EN: User params error
	// ZH_CN: 用户参数错误
	UserParamsError string = "sys.user.userParamsError"

	// GenerateTokenError
	// EN: Generate token error
	// ZH_CN: 生成token错误
	GenerateTokenError string = "sys.user.generateTokenError"

	// ApiAlreadyExists
	// En: API already exists
	// ZH_CN: API已存在
	ApiAlreadyExists string = "sys.apis.apiExists"

	// ApiNotExists
	// En: API does not exist
	// ZH_CN: API不存在
	ApiNotExists string = "sys.apis.apiNotExists"

	// DictionaryAlreadyExists
	// En: Dictionary already exists
	// ZH_CN: 字典已存在
	DictionaryAlreadyExists string = "sys.dict.dictExists"

	// CreateAccountFirst
	// En: Please register an account with this email or bind the email to an account
	// ZH_CH: 请创建一个该邮箱的账号或绑定该邮箱到一个账号
	CreateAccountFirst string = "sys.oauth.createAccount"

	// ProviderNotExists
	// En: Provider does not exist
	// ZH_CH: 第三方登录不存在
	ProviderNotExists string = "sys.oauth.providerNotExists"

	// OAuthUserUnmarshalError
	// En: OAuth user unmarshal error
	// ZH_CH: 第三方用户解析错误
	OAuthUserUnmarshalError string = "sys.oauth.userUnmarshalError"

	// OAuthUserEmailBindError
	// En: OAuth user email bind error
	// ZH_CH: 第三方用户邮箱绑定错误
	OAuthUserEmailBindError string = "sys.oauth.userEmailBindError"

	// OAuthUserBindError
	// En: OAuth user bind error
	// ZH_CH: 第三方用户绑定错误
	OAuthUserBindError string = "sys.oauth.userBindError"
)
