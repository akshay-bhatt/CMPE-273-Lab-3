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
	hi.GET("/keys/:key_id", get_Key1)
	hi.PUT("/keys/:key/:val", new_Key_Val1)
	hi.GET("/keys", all_Key1)

	http.ListenAndServe("localhost:3000", hi)
}

func new_Key_Val1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
	var temp1 ReqKeyVal
	temp1.Key = p.ByName("key")
	temp1.Value = p.ByName("val")
	RKVserver1[temp1.Key] = temp1.Value
	rw.WriteHeader(200)
	fmt.Fprint(rw)
}

func get_Key1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {

	kId := p.ByName("key_id")
	var temp1 ReqKeyVal
	val := RKVserver1[kId]
	temp1.Key = kId
	temp1.Value = val
	x1, _ := json.Marshal(&temp1)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", x1)
}

func all_Key1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
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
