package globalkey

import "fmt"

const (
	InitDatabaseLock     = "init_database_lock"
	InitDatabaseState    = "init_database_state"
	InitDatabaseErrorMsg = "init_database_error_msg"
)

const (
	RoleList           = "role_list"
	RoleListItemID     = "%d_rl_id"
	RoleListItemValue  = "%d_rl_value"
	RoleListItemStatus = "%d_rl_status"
)

func GetRoleListID(id uint64) string {
	return fmt.Sprintf(RoleListItemID, id)
}
func GetRoleListValue(id uint64) string {
	return fmt.Sprintf(RoleListItemValue, id)
}
func GetRoleListStatus(id uint64) string {
	return fmt.Sprintf(RoleListItemStatus, id)
}

const (
	TokenBlackList = "b_token_%s"
)

func GetBlackListToken(token string) string {
	return fmt.Sprintf(TokenBlackList, token)
}
