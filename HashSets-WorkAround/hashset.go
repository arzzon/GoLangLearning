package main
import "strconv"
type HashSet struct{
	hs map[interface{}]struct{}
}

func NewHashSet() *HashSet{
	return &HashSet{
		hs: map[interface{}]struct{}{},
	}
}

func (hs *HashSet) Add(val interface{}){
	hs.hs[val] = struct{}{}
}

func (hs *HashSet) Contains(ele interface{}) bool{
	_, found := hs.hs[ele]
	return found
}

func (hs *HashSet) Remove(ele interface{}){
	delete(hs.hs, ele)
}

func (hs *HashSet) String() string{
	res := "{ "
	for k, _ := range hs.hs {
		switch k.(type) {
		case string:
			res += k.(string) + " "
		case int:
			res += strconv.Itoa(k.(int)) + " "

		}
	}
	res += "}"
	return res
}

func main(){
	//hashset = map[string]struct{}{} // value occupies zero bytes of space
	hs := NewHashSet()
	hs.Add("arbaaz")
	hs.Add("KHAN")
	//hs.Add(5)
	print(hs.Contains("arbaaz"))
	//hs.Remove("KHAN")
	print(hs.String())
}

/*
The size of a struct is the sum of the size of the types of its fields, since there are no fields: no size!
*/