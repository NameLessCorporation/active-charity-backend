package role

import (
	"context"
	"github.com/NameLessCorporation/active-charity-backend/extra/role"
)

func (r *RoleEndpoint) UpdateRole(ctx context.Context, req *role.RoleRequest) (*role.RoleMessageResponse, error) {
	return &role.RoleMessageResponse{}, nil
}
