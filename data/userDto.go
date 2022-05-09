package data

import "encoding/json"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Ordersid  []string
}
type UserCreate struct {
	FirstName string
	LastName  string
}

func (u *User) CacheHashKey() string {
	return "user:id:" + u.ID
}

func (u *User) CacheHashField() string {
	return "user"
}
func (u *User) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, u)
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
