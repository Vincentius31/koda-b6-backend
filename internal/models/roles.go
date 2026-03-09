package models

type Role struct {
	IDRoles   int    `json:"id_roles"`
	NameRoles string `json:"name_roles"`
}

type CreateRoleRequest struct{
	NameRoles string `json:"name_roles" binding:"required"`
}

type UpdateRoleRequest struct {
	NameRoles *string `json:"name_roles"`
}
