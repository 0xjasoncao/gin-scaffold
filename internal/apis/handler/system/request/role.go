package request

type RoleCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment"`
}

type RoleDeleteRequest struct {
	IDs []uint64 `json:"ids" binding:"required,min=1"`
}
