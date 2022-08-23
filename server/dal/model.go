package dal

/*
	User (returned to client) is separate from the UserDocument (stored in DB).

	This is for several reasons:
	- Decouple interface and DB, so changes can be made to the JSON returned without affecting the DB (or vice-versa)
	- CreateDate can depend on client timezone (not implemented), so DB should store the Unix timestamp instead of date string
	- CreateTs can be easily validated and sorted
*/

type User struct {
	Id          *string `json:"id"`
	Name        *string `json:"name,omitempty"`
	DateOfBirth *string `json:"dob,omitempty"`
	Address     *string `json:"address,omitempty"`
	Desc        *string `json:"description,omitempty"`
	CreateDate  *string `json:"createdAt,omitempty"`
}

type UserDocument struct {
	Id          *string `bson:"id"`
	Name        *string `bson:"name"`
	DateOfBirth *string `bson:"dob"`
	Address     *string `bson:"address"`
	Desc        *string `bson:"description"`
	CreateTs    *int64  `bson:"createTs"`
}
