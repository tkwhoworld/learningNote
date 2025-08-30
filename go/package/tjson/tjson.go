package tjson

import (
	"fmt"
	"encoding/json"
)

type tjson struct {
	
}

func Begin(){
	fmt.Println("t json is begin")
}

func (t tjson)MarshalJSON() (a2 []byte, err error){
	var myArr [6]byte = [6]byte{1,2,3,4,5,6}
	a2 = myArr[1:]
	err = nil
	return 
}

func GetTypeNumberInfo(){
	var nu json.Number = "12131211213123123131231"
	in6, err := nu.Int64()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("the int64 number is:",in6)

	fl6, err := nu.Float64()
	if err != nil {
		fmt.Println("the float64 method error is:",err)
	}
	fmt.Println("the float64 number is:",fl6)

	str := nu.String()
	fmt.Println("the string return is:",str)
}