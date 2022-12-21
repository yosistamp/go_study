package main

import "fmt"

type UserResponse = []struct {
	// Address address
	Address *string `json:"address,omitempty"`

	// Age age
	Age *int32 `json:"age,omitempty"`

	// Id ID
	Id *string `json:"id,omitempty"`

	// Name ID
	Name *string `json:"name,omitempty"`

	// Tel tel
	Tel *string `json:"tel,omitempty"`
}

func main() {

	var users UserResponse
	for i := 0; i < 5; i++ {
		address := "osaka"
		age := int32(i + 10)
		id := "1"
		name := "usagisan"
		tel := "00-0000-0000"
		u := struct {
			Address *string "json:\"address,omitempty\""
			Age     *int32  "json:\"age,omitempty\""
			Id      *string "json:\"id,omitempty\""
			Name    *string "json:\"name,omitempty\""
			Tel     *string "json:\"tel,omitempty\""
		}{
			Address: &address,
			Age:     &age,
			Id:      &id,
			Name:    &name,
			Tel:     &tel,
		}
		users = append(users, u)
	}
	for _, u := range users {
		fmt.Println(*u.Address)
		fmt.Println(*u.Age)
		fmt.Println(*u.Id)
		fmt.Println(*u.Name)
		fmt.Println(*u.Tel)
	}
}
