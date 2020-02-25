package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Address struct {
	Flat     string
	Locality string
	City     string
	Country  string
}

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Addr     Address   `json:"address"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	// Encoding primitive data types into JSON
	// func Marshal(v interface{}) ([]byte, error)
	p1, _ := json.Marshal('a')
	fmt.Println("p1 = ", string(p1))
	p2, _ := json.Marshal(5)
	fmt.Println("p2 = ", string(p2))
	p3, _ := json.Marshal(false)
	fmt.Println("p3 = ", string(p3))

	// Encoding Arrays and slices into JSON
	arr1 := [3]string{"abc", "def", "ghi"}
	jArr1, _ := json.Marshal(arr1)
	fmt.Println(string(jArr1))

	slc1 := []int{1, 2, 3, 4}
	jSlc1, _ := json.Marshal(slc1)
	fmt.Println(string(jSlc1))

	// Encoding Maps into JSON
	map1 := map[string]string{"name": "ak", "Address": "India"}
	jMap1, _ := json.Marshal(map1)
	fmt.Println(string(jMap1))

	// Encoding user defined structs to JSON
	user1 := &User{
		ID:   1,
		Name: "AK",
		Addr: Address{
			Flat:     "A-10",
			Locality: "KP",
			City:     "Pune",
			Country:  "India",
		},
		LastSeen: time.Now(),
	}
	jUser1, _ := json.Marshal(user1)
	fmt.Println(string(jUser1))

	// Decoding JSON
	// func Unmarshal(data []byte, v interface{}) error
	// Using the encoded json from the above example
	var user2 User
	if err := json.Unmarshal(jUser1, &user2); err != nil {
		panic(err)
	}
	fmt.Println(user2.Addr.City)

	// Creating json data and unmarshaling it
	jUser2 := []byte(`{"id":1,"name":"AK","address":{"Flat":"A-10","Locality":"KP","City":"Pune","Country":"India"},"lastSeen":"2019-04-12T15:22:41.367931572+05:30"}`)
	var user3 User
	if err := json.Unmarshal(jUser2, &user3); err != nil {
		panic(err)
	}
	fmt.Println(user3.Name)

}
