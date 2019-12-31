package main

import (
	"fmt"
	"encoding/json"
)

type RequestRecord struct {
	Name             string `json:"name"`
}

func MakeRequestTyped(r interface{}) {
	fmt.Printf("type: %T value:%v address:%p\n", r, r, r)
	json.Unmarshal([]byte(`{"name": "xxx"}`), r)
}

func main() {
	//1. working
	//var r *RequestRecord
	//MakeRequestTyped(&r)

	//2. working
	//r:=&RequestRecord{}
	//MakeRequestTyped(r)

	//3. not working, pointer will copied at the MakeRequestTyped call
	var r *RequestRecord
	fmt.Printf("type: %T value:%v address:%p\n", r, r, r)
	MakeRequestTyped(r)

	if r == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	fmt.Printf("%#v", r)
}


//package main
//
//import (
//"fmt"
//"encoding/json"
//)
//type RequestRecord struct {
//	Name             string `json:"name"`
//}
//
//func MakeRequestTyped(r interface{}) {
//	json.Unmarshal([]byte(`[{"name": "xxx"}, {"name": "xxx"}]`), r)
//}
//
//func main() {
// This is okay though!!!
//	var r *[]RequestRecord
//	MakeRequestTyped(&r)
//	fmt.Printf("%#v", r)
//}