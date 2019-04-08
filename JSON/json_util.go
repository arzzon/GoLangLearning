package main

import (
	"encoding/json"
	"os"
	"time"
	//"fmt"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "AK", time.Now()},
	)

	//h := json.RawMessage(`{"ID": 2, "Name": Arbaaz, "LastSeen": time.Time}`)
	//fmt.Println(h)

}
