package models

// Model for ORM
type Invite struct {
	Email     string
	Code      string
	Activated bool
}
