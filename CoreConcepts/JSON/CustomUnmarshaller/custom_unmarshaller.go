package main
import (
	"fmt"
	"encoding/json"
	"strconv"
)

type Student struct{
	Name string
	Age Flexint
	Gender *Gender
}

type Gender int

const (
	MALE Gender = 0
	FEMALE Gender = 1 
)

type Flexint int
func (fi *Flexint) UnmarshalJSON(b []byte) error{
	if b[0] != '"'{
		return json.Unmarshal(b, (*int)(fi))
	}
	var str string
        if err := json.Unmarshal(b, &str); err != nil {
               return err
        }
        i, err := strconv.Atoi(str)
        if err != nil {
               return err
        }
        *fi = Flexint(i)
        return nil
}

func main(){
	fmt.Println("Custom JSON unmarshaller.")
	s := &Student{}
	if err := json.Unmarshal([]byte(`{ "Name": "AK", "Age": "20", "Gender": "FEMALE"}`), &s); err != nil {
		fmt.Println("Error: %v",err)
        }
	fmt.Println("Student: %v",*s)
}
