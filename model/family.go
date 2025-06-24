package model

type Family struct {
	ID   uint
	Name string
}

type FamilyUser struct {
	FamilyID uint
	UserID   uint
}
