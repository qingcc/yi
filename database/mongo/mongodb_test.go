package mongo

import (
	"github.com/siddontang/go/bson"
	"log"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Id   int
}

func TestGetMongoConn(t *testing.T) {
	sess := GetMongoConn()

	//a := Person{
	//	Name: "a",
	//	Age:  10,
	//	Id:   2,
	//}
	//b := Person{
	//	Name: "b",
	//	Age:  50,
	//	Id:   1,
	//}
	//c := Person{
	//	Name: "c",
	//	Age:  25,
	//	Id:   3,
	//}
	//err := sess.DB("test").C("test_sort").Insert(&a)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = sess.DB("test").C("test_sort").Insert(&b)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = sess.DB("test").C("test_sort").Insert(&c)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	ps := make([]Person, 0, 5)
	ps1 := make([]Person, 0, 5)
	ps2 := make([]Person, 0, 5)
	err := sess.DB("test").C("test_sort").Find(bson.M{"id": bson.M{"$gt": 0}}).All(&ps)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("len:", len(ps))
	for _, p := range ps {
		log.Printf("p:%+v", p)
	}
	log.Println()
	log.Println()
	log.Println("sort id")
	err = sess.DB("test").C("test_sort").Find(bson.M{"id": bson.M{"$gt": 0}}).Sort("id").All(&ps1)
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range ps1 {
		log.Printf("p:%+v", p)
	}
	log.Println()
	log.Println()
	log.Println()
	err = sess.DB("test").C("test_sort").Find(bson.M{"id": bson.M{"$gt": 0}}).Sort("-id").All(&ps2)
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range ps2 {
		log.Printf("p:%+v", p)
	}
}
