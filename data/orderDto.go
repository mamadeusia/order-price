package data

import "encoding/json"

type Order struct {
	ID         string
	UserId     string
	Price      int64
	StartPrice int64
	Amount     int64
	Pair       string
}

type OrderCreate struct {
	UserId string
	Price  int64
	Amount int64
	Pair   string
}

type OrderUpdate struct {
	ID    string
	Price int64
}

func (o *Order) CacheHashKey() string {
	return "order:id:" + o.ID
}

func (o *Order) CacheHashField() string {
	return "order"
}
func (o *Order) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, o)
}
func (o *Order) MarshalBinary() (data []byte, err error) {
	return json.Marshal(o)
}
