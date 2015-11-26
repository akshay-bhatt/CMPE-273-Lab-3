package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Response struct {
	Key   string
	Value string
}

type Node struct {
	Key   string
	Value int
}

type NodeList []Node

func (nl NodeList) Len() int {
	return len(nl)
}
func (nl NodeList) Less(i, j int) bool {
	return nl[i].Value < nl[j].Value
}
func (nl NodeList) Swap(i, j int) {
	nl[i], nl[j] = nl[j], nl[i]
}

var serverMapNo map[int]string
var sortans NodeList
var serverMapNohash map[string]int

func hashfunc() NodeList {
	serverMapNohash = make(map[string]int)
	for _, val := range serverMapNo {
		serverHash := int(crc32.ChecksumIEEE([]byte(val)))
		serverMapNohash[val] = serverHash

	}
	ans := sortfunc(serverMapNohash)
	return ans
}

func main() {

	serverMapNo = make(map[int]string)
	serverMapNo[0] = "3002"
	//	var input int
	//	fmt.Scanln(&input)
	serverMapNo[1] = "3000"

	//	keyvalarr := []string{"1,a", "2,b", "3,c", "4,d", "5,e", "6,f", "7,g", "8,h", "9,i", "10,j"}
	serverMapNo[2] = "3001"

	//	var input int
	//	fmt.Scanln(&input)
	sortans = hashfunc()
	keyvalarr := []string{"1,a", "2,b", "3,c", "4,d", "5,e", "6,f", "7,g", "8,h", "9,i", "10,j"}

	client := &http.Client{}
	for i := 0; i < len(keyvalarr); i++ {
		arrsplit := strings.Split(keyvalarr[i], ",")
		key := arrsplit[0]
		val := arrsplit[1]

		hash := int(crc32.ChecksumIEEE([]byte(key)))
		var sPort string

		big := sortans[0].Value
		small := sortans[1].Value
		cen := sortans[2].Value

		if hash > big {
			sPort = sortans[2].Key
		}
		if hash > cen && hash < big || hash == big {
			/*if hash && hash < cen || hash ==
			,< cen && hash < big || hash == big {
				sPort = sortans[1].Key
			}*/
			sPort = sortans[0].Key
		}
		if hash > small && hash < cen || hash == cen {
			sPort = sortans[1].Key
		}
		/*if hash
		,< cen && hash < big || hash == big {
			sPort = sortans[0].Key
		}*/
		if hash < small || hash == small {
			sPort = sortans[2].Key
		}

		serverport := sPort

		url := "http://localhost:" + serverport + "/keys/" + key + "/" + val
		aaa := []byte(`{}`)
		xx, _ := http.NewRequest("PUT", url, bytes.NewBuffer(aaa))

		xx.Header.Set("Content-Type", "application/json")
		xx.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response, err := client.Do(xx)
		if err != nil {
			panic(err)
		}
		//		fmt.Println(response)
		defer response.Body.Close()
	}

	keyarr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for i := 0; i < len(keyarr); i++ {
		key := keyarr[i]

		hash := int(crc32.ChecksumIEEE([]byte(key)))
		var sPort string

		big := sortans[0].Value
		small := sortans[1].Value
		cen := sortans[2].Value

		if hash > big {
			sPort = sortans[2].Key
		}
		if hash > cen && hash < big || hash == big {
			sPort = sortans[0].Key
		}
		// nl[i].Value < nl[j].Value
		if hash > small && hash < cen || hash == cen {
			sPort = sortans[1].Key
		}
		if hash < small || hash == small {
			sPort = sortans[2].Key
		}

		serverport := sPort

		url := "http://localhost:" + serverport + "/keys/" + key
		xx, _ := http.Get(url)
		/*
			Url := "https://" + url1

						res, _ := http.Get(Url)
						data, _ := ioutil.ReadAll(res.Body)
						res.Body.Close()
						_ = json.Unmarshal(data, &Ubp)
		*/
		defer xx.Body.Close()
		contents, _ := ioutil.ReadAll(xx.Body)
		fmt.Printf("%s\n", string(contents))
		fmt.Println("Port :", serverport)
	}

}

func sortfunc(mapToBeSorted map[string]int) NodeList {
	tempss := make(NodeList, len(mapToBeSorted))
	i := 0
	for keyy, valls := range mapToBeSorted {
		/*for i := 0; i < len(keyvalarr); i++ {
		arrsplit := strings.Split(keyvalarr[i], ",")
		*/
		tempss[i] = Node{keyy, valls}
		i++
	}

	sort.Sort(sort.Reverse(tempss))
	return tempss
}
