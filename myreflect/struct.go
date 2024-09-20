package myreflect

import "time"

type (
	Producet struct {
		ID        uint32
		Name      string
		Price     uint32
		LeftCount uint32 `orm:"left_count"`
		Batch     string `orm:"batch_number"`
		Updateed  time.Time
	}

	Person struct {
		ID      string
		Name    string
		Age     uint32
		Gender  string
		Addr    string `orm:"address"`
		Updated time.Time
	}
)
