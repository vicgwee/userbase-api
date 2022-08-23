package dal

/*
	User (returned to client) is separate from the UserDocument (stored in DB).

	This is for several reasons:
	- Decouple interface and DB, so changes can be made to the JSON returned without affecting the DB (or vice-versa)
	- CreateDate can depend on client timezone (not implemented), so DB should store the Unix timestamp instead of date string
	- CreateTs can be easily validated and sorted
*/

type User struct {
	Id          *string `json:"id" validate:"required" minLength:"1" maxLength:"16" example:"1"`
	Name        *string `json:"name,omitempty" validate:"required" minLength:"1" maxLength:"100" example:"test"`
	DateOfBirth *string `json:"dob,omitempty" validate:"required" minLength:"8" maxLength:"8" example:"20060102"`
	Address     *string `json:"address,omitempty" validate:"required" minLength:"1" maxLength:"100" example:"10 Anson Road, #17-06, International Plaza, 097903"`
	Desc        *string `json:"description,omitempty" validate:"required" minLength:"1" maxLength:"1000" example:"testDescription"`
	CreateDate  *string `json:"createdAt,omitempty" validate:"optional" minLength:"8" maxLength:"8" example:"20220801"`
}

type UserDocument struct {
	Id          *string `bson:"id"`
	Name        *string `bson:"name"`
	DateOfBirth *string `bson:"dob"`
	Address     *string `bson:"address"`
	Desc        *string `bson:"description"`
	CreateTs    *int64  `bson:"createTs"`
}
