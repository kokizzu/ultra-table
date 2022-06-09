package main

import (
	"fmt"
	"log"

	ultra_table "github.com/longbridgeapp/ultra-table"
	"github.com/longbridgeapp/ultra-table/test_data/easyjson"
	"github.com/longbridgeapp/ultra-table/test_data/pb"
)

func main() {
	baseEasyjson() //serialization based easyjson
	basegogo()     //serialization based gogo protobuf
}

func baseEasyjson() {
	table := ultra_table.New[*easyjson.Person](&easyjson.Person{})

	err := table.Add(&easyjson.Person{
		Name:     "jacky",
		Phone:    "+8613575468007",
		Age:      31,
		BirthDay: 19901111,
		Gender:   0,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = table.Add(&easyjson.Person{
		Name:     "rose",
		Phone:    "+8613575468008",
		Age:      31,
		BirthDay: 19901016,
		Gender:   1,
	})
	if err != nil {
		log.Fatalln("easyjson", err)
	}

	infos, err := table.GetWithIdx("Phone", "+8613575468007")
	if err != nil {
		log.Fatalln("easyjson", err)
	}
	for i := 0; i < len(infos); i++ {
		log.Printf("easyjson %+v \n", infos[i])
	}

	_, err = table.GetWithIdxIntersection(map[string]interface{}{
		"Age":  31,
		"Name": "rose",
	})
	log.Println("easyjson", err)

	m := map[string]string{
		"HKD": "0",
		"USD": "10",
		"SGD": "0",
	}
	clear(m)
	fmt.Println("after", m)

}

func clear(cashBooks map[string]string) {
	fmt.Println("before", cashBooks)
	for currency, detail := range cashBooks {
		if detail == "0" {
			delete(cashBooks, currency)
		}
	}
}

func basegogo() {
	table := ultra_table.New[*pb.Person](&pb.Person{})

	err := table.Add(&pb.Person{
		Name:     "jacky",
		Phone:    "+8613575468007",
		Age:      31,
		BirthDay: 19901111,
		Gender:   pb.Gender_men,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = table.Add(&pb.Person{
		Name:     "rose",
		Phone:    "+8613575468008",
		Age:      31,
		BirthDay: 19901016,
		Gender:   pb.Gender_women,
	})
	if err != nil {
		log.Fatalln("gogo", err)
	}

	infos, err := table.GetWithIdx("Phone", "+8613575468007")
	if err != nil {
		log.Fatalln("gogo", err)
	}
	for i := 0; i < len(infos); i++ {
		log.Printf("gogo %+v \n", infos[i])
	}

	_, err = table.GetWithIdxIntersection(map[string]interface{}{
		"Age":  31,
		"Name": "rose",
	})
	log.Println("gogo", err)
}
