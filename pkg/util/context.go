package util

import (
	"context"

	"github.com/ilhamosaurus/HRIS/pkg/types"
)

const (
	UsernameKey = "username"
	RoleKey     = "role"
	IDKey       = "id"
)

func GetUserIDFromCtx(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(IDKey).(int64)
	return id, ok
}

func GetUsernameFromCtx(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)
	return username, ok
}

func GetRoleFromCtx(ctx context.Context) (types.Role, bool) {
	role, ok := ctx.Value(RoleKey).(types.Role)
	return role, ok
}
