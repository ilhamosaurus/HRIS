package dto

type User struct {
	ID       *int32   `json:"id"`
	Username string   `json:"username" validate:"required,gte=6"`
	Password *string  `json:"password" validate:"gte=8"`
	Role     string   `json:"role" validate:"required,role"`
	Salary   *float64 `json:"salary"`
	Active   *bool    `json:"active"`
}

func (u *User) GetSalary() float64 {
	if u.Salary == nil {
		return 0
	}
	return *u.Salary
}

func (u *User) IsActive() bool {
	if u.Active == nil {
		return false
	}
	return *u.Active
}

func (u *User) GetID() int32 {
	if u.ID == nil {
		return 0
	}
	return *u.ID
}

func (u *User) GetPassword() string {
	if u.Password == nil {
		return ""
	}
	return *u.Password
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,gte=3"`
	Password string `json:"password" validate:"required,gte=8"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required,gte=8"`
	NewPassword string `json:"newPassword" validate:"required,password"`
}

type CreateUserRequest struct {
	Username string  `json:"username" validate:"required,gte=6"`
	Password string  `json:"password" validate:"required,gte=8"`
	Email    string  `json:"email" validate:"required,email"`
	Role     string  `json:"role" validate:"required,oneof=admin manager employee"`
	Salary   float64 `json:"salary,omitempty" validate:"omitempty,min=1"`
	Active   bool    `json:"active"`
}

type UpdateUserRequest struct {
	ID       int64   `json:"id" validate:"required,min=1"`
	Username string  `json:"username" validate:"omitempty,gte=6"`
	Password string  `json:"password" validate:"omitempty,gte=8"`
	Email    string  `json:"email" validate:"omitempty,email"`
	Role     string  `json:"role" validate:"omitempty,oneof=admin manager employee"`
	Salary   float64 `json:"salary,omitempty" validate:"omitempty,min=1"`
	Active   bool    `json:"active"`
}

type UserQuery struct {
	Role     string `query:"role,omitempty" validate:"omitempty,oneof=admin manager employee"`
	Active   string `query:"active,omitempty" validate:"omitempty,oneof=true false"`
	Page     int    `query:"page" validate:"required,min=1"`
	PageSize int    `query:"page_size" validate:"required,min=1,max=100"`
}

type UserResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Salary   float64 `json:"salary,omitempty"`
	Active   bool    `json:"active"`
}

type UserListResponse struct {
	Users      []*UserResponse `json:"list"`
	TotalCount int64           `json:"total"`
	PageSize   int             `json:"page_size"`
	Page       int             `json:"page"`
}
