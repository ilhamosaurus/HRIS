package types

import "strings"

type Role uint32

const (
	Unknown_Role Role = 0
	Admin        Role = 1 << iota
	Employee
)

func StringToRole(s string) Role {
	switch strings.ToUpper(s) {
	case "ADMIN":
		return Admin
	case "EMPLOYEE":
		return Employee
	default:
		return Unknown_Role
	}
}

func (r Role) String() string {
	switch r {
	case Admin:
		return "ADMIN"
	case Employee:
		return "EMPLOYEE"
	default:
		return "UNKNOWN"
	}
}

type Status uint32

const (
	Unknown_Status Status = 0
	Pending        Status = 1 << iota
	Approved
	Rejected
)

func StringToStatus(s string) Status {
	switch strings.ToUpper(s) {
	case "PENDING":
		return Pending
	case "APPROVED":
		return Approved
	case "REJECTED":
		return Rejected
	default:
		return Unknown_Status
	}
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "PENDING"
	case Approved:
		return "APPROVED"
	case Rejected:
		return "REJECTED"
	default:
		return "UNKNOWN"
	}
}

type User struct {
	ID       *int32   `json:"id"`
	Username string   `json:"username" validate:"required,gte=6"`
	Password *string  `json:"password" validate:"gte=8"`
	Role     string   `json:"role" validate:"required,role"`
	Salary   *float64 `json:"salary"`
	Active   *bool    `json:"active"`
}

func (u User) GetSalary() float64 {
	if u.Salary == nil {
		return 0
	}
	return *u.Salary
}

func (u User) IsActive() bool {
	if u.Active == nil {
		return false
	}
	return *u.Active
}

func (u User) GetID() int32 {
	if u.ID == nil {
		return 0
	}
	return *u.ID
}

func (u User) GetPassword() string {
	if u.Password == nil {
		return ""
	}
	return *u.Password
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,gte=3"`
	Password string `json:"password" validate:"required,gte=8"`
}

type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GenerateReponse(code int, message string, data any) (int, HTTPResponse) {
	return code, HTTPResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
