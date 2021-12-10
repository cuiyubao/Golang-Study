package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]interface{})
	m["id"]=2671514
	i := m["id"].(int)
	strconv.Itoa(i)
	fmt.Println(m)

}
