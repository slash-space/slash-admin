package globalkey

import "fmt"

const (
	InitDatabaseLock     = "init_database_lock"
	InitDatabaseState    = "init_database_state"
	InitDatabaseErrorMsg = "init_database_error_msg"
)

const (
	RoleList = "role_list"
)

const (
	TokenBlackList = "b_token_%s"
)

func GetBlackListToken(token string) string {
	return fmt.Sprintf(TokenBlackList, token)
}
