package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ReqKeyVal struct {
	Key   string
	Value string
}

type Response struct {
	Response []ReqKeyVal
}

var RKVserver1 map[string]string = make(map[string]string)

func main() {

	hi := httprouter.New()
	hi.GET("/keys/:key_id", get_Key2)
	hi.PUT("/keys/:key/:val", new_Key_Val2)
	hi.GET("/keys", all_Key2)

	http.ListenAndServe("localhost:3001", hi)
}

func new_Key_Val2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
	var temp2 ReqKeyVal
	temp2.Key = p.ByName("key")
	temp2.Value = p.ByName("val")
	RKVserver1[temp2.Key] = temp2.Value
	rw.WriteHeader(200)
	fmt.Fprint(rw)
}

func get_Key2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {

	kId := p.ByName("key_id")
	var temp2 ReqKeyVal
	val := RKVserver1[kId]
	temp2.Key = kId
	temp2.Value = val
	x1, _ := json.Marshal(&temp2)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", x1)
}

func all_Key2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
	var Res1 Response
	for key, val := range RKVserver1 {
		temp := ReqKeyVal{}
		temp.Key = key
		temp.Value = val
		Res1.Response = append(Res1.Response, temp)
	}
	x1, _ := json.MarshalIndent(Res1, "", "\t")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", x1)

}
