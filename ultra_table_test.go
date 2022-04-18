package ultra_table

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/longbridgeapp/ultra-table/test_data/pb"
	. "github.com/smartystreets/goconvey/convey"
)

type Person struct {
	Name     string
	Phone    string
	Age      int32
	BirthDay int32
	Gender   uint8
}

func (p *Person) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
func (p *Person) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

type PersonWithNormal struct {
	Name     string `json:"name,omitempty" idx:"normal"`
	Phone    string `json:"phone,omitempty" idx:"normal"`
	Age      int32  `json:"age,omitempty" idx:"normal"`
	BirthDay int32  `json:"birthDay,omitempty"`
	Gender   uint8  `json:"gender,omitempty"`
}

func (p *PersonWithNormal) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
func (p *PersonWithNormal) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

func TestQuery(t *testing.T) {
	Convey("query", t, func() {
		Convey("get", func() {
			table, err := NewWithInitializeData([]*Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   0,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   1,
			}})
			So(err, ShouldBeNil)
			results := table.Get(func(person *Person) bool {
				return person.Name == "jacky"
			})
			So(len(results), ShouldEqual, 1)
			_, err = table.GetWithIdx("Name", "kevin")
			So(err, ShouldEqual, RecordNotFound)
			So(table.Len(), ShouldEqual, 2)
			i := table.Remove(func(person *Person) bool {
				return person.Age == 31
			})
			So(i, ShouldEqual, 2)
			So(table.Len(), ShouldEqual, 0)
		})
		Convey("get with index", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			results, err := table.GetWithIdx("Age", int32(31))
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 2)

			results, err = table.GetWithIdx("Name", "kevin")
			So(err, ShouldNotBeNil)
			So(len(results), ShouldEqual, 0)

			results, err = table.GetWithIdx("Phone", "+8613575468007")
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results, err = table.GetWithIdx("Phone", "+8613575468009")
			So(err, ShouldNotBeNil)
			So(len(results), ShouldEqual, 0)
		})
		Convey("get with intersection", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)
			results, err := table.GetWithIdxIntersection(map[string]interface{}{
				"Age":  int32(31),
				"Name": "jacky",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results, err = table.GetWithIdxIntersection(map[string]interface{}{
				"Age": int32(31),
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 2)

			results, err = table.GetWithIdxIntersection(map[string]interface{}{
				"Age":  int32(31),
				"Name": "anthony",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 0)

			results, err = table.GetWithIdxIntersection(map[string]interface{}{
				"Phone": "+8613575468002",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)
		})
		Convey("get with aggregate", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)
			results, err := table.GetWithIdxAggregate(map[string]interface{}{
				"Age":  int32(31),
				"Name": "jacky",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 2)

			results, err = table.GetWithIdxAggregate(map[string]interface{}{
				"Age": int32(31),
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 2)

			results, err = table.GetWithIdxAggregate(map[string]interface{}{
				"Age":  int32(31),
				"Name": "anthony",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 3)

			results, err = table.GetWithIdxAggregate(map[string]interface{}{
				"Phone": "+8613575468002",
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)

			results, err = table.GetWithIdxAggregate(map[string]interface{}{
				"Phone": "+8613575468019",
				"Age":   int32(30),
			})
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 1)
		})
	})
}

func TestDelete(t *testing.T) {
	Convey("delete", t, func() {
		Convey("clear", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)
			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)
			table.Clear()
			So(table.Len(), ShouldEqual, 0)
			So(table.Cap(), ShouldEqual, 0)

			table.Add(&pb.Person{
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			})
			table.Add(&pb.Person{
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			})
			So(table.Len(), ShouldEqual, 2)
			So(table.Cap(), ShouldEqual, 2)

		})
		Convey("remove with index", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			count := table.Remove(func(person *pb.Person) bool {
				return person.Age == 31
			})
			So(count, ShouldEqual, 2)
			count = table.Remove(func(person *pb.Person) bool {
				return person.Age == 31
			})
			So(count, ShouldEqual, 0)
			So(table.Len(), ShouldEqual, 3)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdx("Phone", "+8613575468003")
			So(count, ShouldEqual, 1)
			So(table.Len(), ShouldEqual, 2)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdx("Phone", "+8613575468007")
			So(count, ShouldEqual, 0)

			count = table.Remove(func(person *pb.Person) bool {
				return person.Age < 30
			})
			So(count, ShouldEqual, 1)
			So(table.Len(), ShouldEqual, 1)
			So(table.Cap(), ShouldEqual, 5)
		})
		Convey("remove with index intersection", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			count := table.RemoveWithIdxIntersection(map[string]interface{}{
				"Age":  int32(31),
				"Name": "jack",
			})
			So(count, ShouldEqual, 0)

			count = table.RemoveWithIdxIntersection(map[string]interface{}{
				"Age": int32(31),
			})
			So(count, ShouldEqual, 2)
			So(table.Len(), ShouldEqual, 3)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdxIntersection(map[string]interface{}{
				"Age": int32(31),
			})
			So(count, ShouldEqual, 0)

			count = table.RemoveWithIdxIntersection(map[string]interface{}{
				"Phone": "+8613575468003",
			})
			So(count, ShouldEqual, 1)
			So(table.Len(), ShouldEqual, 2)
			So(table.Cap(), ShouldEqual, 5)
		})

		Convey("remove with index aggregate", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			count := table.RemoveWithIdxAggregate(map[string]interface{}{
				"Age":  int32(31),
				"Name": "jack",
			})
			So(count, ShouldEqual, 2)
			So(table.Len(), ShouldEqual, 3)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdxAggregate(map[string]interface{}{
				"Age": int32(31),
			})
			So(count, ShouldEqual, 0)
			So(table.Len(), ShouldEqual, 3)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdxAggregate(map[string]interface{}{
				"Phone": "+8613575468003",
			})
			So(count, ShouldEqual, 1)
			So(table.Len(), ShouldEqual, 2)
			So(table.Cap(), ShouldEqual, 5)

			count = table.RemoveWithIdxAggregate(map[string]interface{}{
				"Name": "alex",
				"Age":  int32(32),
			})
			So(count, ShouldEqual, 2)
			So(table.Len(), ShouldEqual, 0)
			So(table.Cap(), ShouldEqual, 5)
		})
	})
}

func Test_Update(t *testing.T) {
	Convey("Update", t, func() {
		Convey("UpdateWithUniqueIdx", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			err = table.UpdateWithUniqueIdx("Name", "jacky", &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901112,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldNotBeNil)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			err = table.UpdateWithUniqueIdx("Age", int32(31), &pb.Person{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      32,
				BirthDay: 19891016,
				Gender:   pb.Gender_women,
			})
			So(err, ShouldNotBeNil)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			err = table.UpdateWithUniqueIdx("Phone", "+8613575468007", &pb.Person{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      32,
				BirthDay: 19891016,
				Gender:   pb.Gender_women,
			})
			So(err, ShouldNotBeNil)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			err = table.UpdateWithUniqueIdx("Phone", "+8613575468009", &pb.Person{
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldBeNil)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			result, err := table.GetWithIdx("Name", "anthony")
			So(err, ShouldBeNil)
			So(result[0].Name, ShouldEqual, "anthony")
			So(result[0].Phone, ShouldEqual, "+8613575468009")
			So(result[0].Age, ShouldEqual, 31)
			So(result[0].BirthDay, ShouldEqual, 19901016)
		})

		Convey("UpdateWithNormalIdx", func() {
			table, err := NewWithInitializeData([]*pb.Person{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   pb.Gender_women,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   pb.Gender_men,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   pb.Gender_women,
			}})
			So(err, ShouldBeNil)

			count, err := table.UpdateWithNormalIdx("Name", "jacky", &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901112,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldNotBeNil)
			So(count, ShouldEqual, 0)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			tableWithNormal, err := NewWithInitializeData([]*PersonWithNormal{{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   0,
			}, {
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901016,
				Gender:   1,
			}, {
				Name:     "anthony",
				Phone:    "+8613575468009",
				Age:      32,
				BirthDay: 19890812,
				Gender:   0,
			}, {
				Name:     "alex",
				Phone:    "+8613575468002",
				Age:      29,
				BirthDay: 19920808,
				Gender:   0,
			}, {
				Name:     "alice",
				Phone:    "+8613575468003",
				Age:      30,
				BirthDay: 19910608,
				Gender:   1,
			}})
			So(err, ShouldBeNil)

			count, err = tableWithNormal.UpdateWithNormalIdx("Age", int32(31), &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901112,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)

			So(table.Len(), ShouldEqual, 5)
			So(table.Cap(), ShouldEqual, 5)

			count, err = tableWithNormal.UpdateWithNormalIdx("Phone", "+8613575468001", &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      31,
				BirthDay: 19901112,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 0)

			count, err = tableWithNormal.UpdateWithNormalIdx("Phone", "+8613575468008", &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901112,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)
		})

		Convey("SaveWithUniqueIdx", func() {
			table := New[*pb.Person]()
			err := table.SaveWithUniqueIdx("Phone", "+8613575468007", &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldBeNil)
			So(table.Len(), ShouldEqual, 1)
			So(table.Cap(), ShouldEqual, 1)

			count := table.RemoveWithIdx("Phone", "+8613575468007")
			So(count, ShouldEqual, 1)
			So(table.Len(), ShouldEqual, 0)
			So(table.Cap(), ShouldEqual, 1)

			err = table.SaveWithUniqueIdx("Phone", "+8613575468007", &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldBeNil)
			So(table.Len(), ShouldEqual, 1)
			So(table.Cap(), ShouldEqual, 1)

			err = table.SaveWithUniqueIdx("Phone", "+8613575468008", &pb.Person{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      32,
				BirthDay: 19891016,
				Gender:   pb.Gender_women,
			})
			So(err, ShouldBeNil)
			So(table.Len(), ShouldEqual, 2)
			So(table.Cap(), ShouldEqual, 2)

			err = table.SaveWithUniqueIdx("Name", "rose", &pb.Person{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      33,
				BirthDay: 19881016,
				Gender:   pb.Gender_women,
			})
			So(err, ShouldNotBeNil)

			err = table.SaveWithUniqueIdx("Phone", "+8613575468007", &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      30,
				BirthDay: 19911111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldNotBeNil)

			table = New[*pb.Person]()

			//if first init add, dose not check unique idx
			err = table.SaveWithUniqueIdx("Name", "rose", &pb.Person{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      33,
				BirthDay: 19881016,
				Gender:   pb.Gender_women,
			})
			So(err, ShouldBeNil)

		})

		Convey("SaveWithNormalIdxIntersection", func() {
			table := New[*pb.Person]()

			count, err := table.SaveWithNormalIdxIntersection(map[string]interface{}{}, &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)

			count, err = table.SaveWithNormalIdxIntersection(map[string]interface{}{}, &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldNotBeNil)
			So(count, ShouldEqual, 0)

			tableWithNormal := New[*PersonWithNormal]()

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 1)

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{
				"Name": "jacky",
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      30,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 1)

			results, err := tableWithNormal.GetWithIdx("Phone", "+8613575468008")
			So(err, ShouldBeNil)
			So(results[0].Age, ShouldEqual, int32(30))
			So(results[0].BirthDay, ShouldEqual, int32(19911111))

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{
				"Name": "rose",
			}, &PersonWithNormal{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19921111,
				Gender:   1,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 2)

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{
				"Phone": "+8613575468008",
			}, &PersonWithNormal{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19921111,
				Gender:   1,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)
			So(tableWithNormal.Len(), ShouldEqual, 2)

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{
				"Name": "jacky",
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      30,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 3)

			count, err = tableWithNormal.SaveWithNormalIdxIntersection(map[string]interface{}{
				"Name": "jacky",
				"Age":  int32(29),
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 4)
		})

		Convey("SaveWithNormalIdxAggregate", func() {
			table := New[*pb.Person]()

			count, err := table.SaveWithNormalIdxAggregate(map[string]interface{}{}, &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)

			count, err = table.SaveWithNormalIdxAggregate(map[string]interface{}{}, &pb.Person{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   pb.Gender_men,
			})
			So(err, ShouldNotBeNil)
			So(count, ShouldEqual, 0)

			tableWithNormal := New[*PersonWithNormal]()

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468007",
				Age:      31,
				BirthDay: 19901111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 1)

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{
				"Name": "jacky",
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      30,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 1)

			results, err := tableWithNormal.GetWithIdx("Phone", "+8613575468008")
			So(err, ShouldBeNil)
			So(results[0].Age, ShouldEqual, int32(30))
			So(results[0].BirthDay, ShouldEqual, int32(19911111))

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{
				"Name": "rose",
			}, &PersonWithNormal{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19921111,
				Gender:   1,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 2)

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{
				"Phone": "+8613575468008",
			}, &PersonWithNormal{
				Name:     "rose",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19921111,
				Gender:   1,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)
			So(tableWithNormal.Len(), ShouldEqual, 2)

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{
				"Name": "jacky",
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      30,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
			So(tableWithNormal.Len(), ShouldEqual, 3)

			count, err = tableWithNormal.SaveWithNormalIdxAggregate(map[string]interface{}{
				"Name": "jacky",
				"Age":  int32(29),
			}, &PersonWithNormal{
				Name:     "jacky",
				Phone:    "+8613575468008",
				Age:      29,
				BirthDay: 19911111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 3)
			So(tableWithNormal.Len(), ShouldEqual, 3)
		})
	})
}

func Test_Concurrent(t *testing.T) {
	Convey("Concurrent", t, func() {
		Convey("Concurrent-1", func() {
			waitGroup := sync.WaitGroup{}
			len := 10
			waitGroup.Add(len * 3)

			ultraTable := New[*PersonWithNormal]()
			for i := 0; i < len; i++ {
				go func(j int) {
					ultraTable.Add(&PersonWithNormal{
						Name:     fmt.Sprintf("anthony-%v", j),
						Phone:    fmt.Sprintf("+86135754680%v", j),
						Age:      int32(20 + j),
						BirthDay: 19890812,
						Gender:   0,
					})
					waitGroup.Done()
				}(i)
			}
			for i := 0; i < len; i++ {
				go func(i int) {
					ultraTable.UpdateWithNormalIdx("Age", int32(20+i), &PersonWithNormal{
						Name:     "anthony",
						Phone:    "+8613575468009",
						Age:      int32(20 + i),
						BirthDay: int32(19890812 + i),
						Gender:   0,
					})
					waitGroup.Done()
				}(i)
			}

			for i := 0; i < len; i++ {
				go func() {
					ultraTable.GetWithIdx("Name", "anthony")
					waitGroup.Done()
				}()
			}
			waitGroup.Wait()
			So(ultraTable.Len(), ShouldEqual, len)
		})

		Convey("Concurrent-Read-Write", func() {
			waitGroup := sync.WaitGroup{}
			len := 100
			waitGroup.Add(len * 2)

			ultraTable := New[*PersonWithNormal]()
			for i := 0; i < len; i++ {
				go func(i int) {
					ultraTable.Add(&PersonWithNormal{
						Name:     "anthony",
						Phone:    "+8613575468009",
						Age:      int32(20 + i),
						BirthDay: int32(19890812 + i),
						Gender:   0,
					})
					waitGroup.Done()
				}(i)
			}

			for i := 0; i < len; i++ {
				go func(i int) {
					ultraTable.GetWithIdx("Age", int32(20+i))
					waitGroup.Done()
				}(i)
			}
			waitGroup.Wait()
			So(ultraTable.Len(), ShouldEqual, len)
		})
	})
}

func Test_SimulationWorkFlow(t *testing.T) {
	Convey("case-1", t, func() {
		table := New[*pb.Person]()

		for i := 1; i <= 100; i++ {
			err := table.Add(&pb.Person{
				Name:     fmt.Sprintf("jack.%v", i),
				Phone:    fmt.Sprintf("+8613575468%v", i),
				Age:      int32(i),
				BirthDay: 19901111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
		}
		So(table.Len(), ShouldEqual, 100)

		count := table.Remove(func(person *pb.Person) bool {
			return person.Age > 50
		})
		So(count, ShouldEqual, 50)
		So(table.Len(), ShouldEqual, 50)

		for i := 51; i <= 100; i++ {
			err := table.Add(&pb.Person{
				Name:     fmt.Sprintf("jack.%v", i),
				Phone:    fmt.Sprintf("+8613575468%v", i),
				Age:      int32(i),
				BirthDay: 19901111,
				Gender:   0,
			})
			So(err, ShouldBeNil)
		}
		So(table.Len(), ShouldEqual, 100)

		count = table.Remove(func(person *pb.Person) bool {
			return person.Age > 20
		})
		So(count, ShouldEqual, 80)
		So(table.Len(), ShouldEqual, 20)

		count = table.Remove(func(person *pb.Person) bool {
			return person.Age <= 20
		})
		So(count, ShouldEqual, 20)
		So(table.Len(), ShouldEqual, 0)
	})
}
