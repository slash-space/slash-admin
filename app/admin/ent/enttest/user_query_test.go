package enttest

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserById(t *testing.T) {
	client := NewTestClient(t)

	user, err := client.SysUser.Get(context.Background(), 1)

	assert.Equal(t, nil, err)
	assert.Equal(t, "admin", user.Username)
}
