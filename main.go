package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Thanh",
	}

	p2 := person{
		First: "Xuan",
	}

	pn := []person{p1, p2}

	pns, err := json.Marshal(pn)

	if err != nil {
		panic(err)
	}

	fmt.Println("convert Go struct to json", string(pns))

	ppn := []person{}

	err = json.Unmarshal(pns, &ppn)
	if err != nil {
		panic(err)
	}

	fmt.Println("convert json into go struct", ppn)

}
