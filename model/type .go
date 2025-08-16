package model

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
