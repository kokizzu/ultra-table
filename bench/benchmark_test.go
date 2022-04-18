package bench

import (
	"fmt"
	"testing"

	ultra_table "github.com/longbridgeapp/ultra-table"
	"github.com/longbridgeapp/ultra-table/test_data/easyjson"
	"github.com/longbridgeapp/ultra-table/test_data/pb"
)

func BenchmarkAddWithGoGo(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAddWithEasyjson(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*easyjson.Person](new(easyjson.Person))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := ultraTable.Add(&easyjson.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   0,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetWithUniqueIndexWithGoGo(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		results, err := ultraTable.GetWithIdx("Phone", "+8613575468006")
		if err != nil {
			b.Fatal(err)
		}
		if results[0].Name != "jacky" || results[0].Phone != "+8613575468006" {
			b.Fail()
		}
	}
}

func BenchmarkGetWithUniqueIndexWithEasyjson(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*easyjson.Person](new(easyjson.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&easyjson.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   0,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		results, err := ultraTable.GetWithIdx("Phone", "+8613575468006")
		if err != nil {
			b.Fatal(err)
		}
		if results[0].Name != "jacky" || results[0].Phone != "+8613575468006" {
			b.Fail()
		}
	}
}
func BenchmarkGetWithNormalIndex(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		results, err := ultraTable.GetWithIdx("Age", int32(30))
		if err != nil {
			b.Fatal(err)
		}
		if results[0].Name != "jacky" || results[0].Phone != "+86135754630" {
			b.Fail()
		}
	}
}

func BenchmarkGetWithIdxIntersectionNotFound(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		results, err := ultraTable.GetWithIdxIntersection(map[string]interface{}{
			"Phone": "+86135754630",
			"Age":   int32(31),
		})
		if err != nil {
			b.Fatal(err)
		}
		if len(results) > 0 {
			b.Fail()
		}
	}
}

func BenchmarkGetWithIdxIntersection(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		results, err := ultraTable.GetWithIdxIntersection(map[string]interface{}{
			"Phone": "+86135754630",
			"Age":   int32(30),
		})
		if err != nil {
			b.Fatal(err)
		}
		if len(results) != 1 {
			b.Fail()
		}
		if results[0].Name != "jacky" || results[0].Phone != "+86135754630" {
			b.Fail()
		}
	}
}

func BenchmarkRemoveWithIndex(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		count := ultraTable.RemoveWithIdx("Phone", fmt.Sprintf("+861357546%d", i))
		if count != 1 && i < 100000 {
			b.Fail()
		}
	}
}

func BenchmarkUpdateWithIndex(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	for i := 0; i < 100000; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := ultraTable.UpdateWithUniqueIdx("Phone", fmt.Sprintf("+861357546%d", i), &pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+871357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_women,
		})
		if err != nil && i < 100000 {
			b.Fatal(err)
		}
	}
}

func BenchmarkAddAndRemove(b *testing.B) {
	b.StopTimer()
	ultraTable := ultra_table.New[*pb.Person](new(pb.Person))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := ultraTable.Add(&pb.Person{
			Name:     "jacky",
			Phone:    fmt.Sprintf("+861357546%d", i),
			Age:      int32(i),
			BirthDay: 19901111,
			Gender:   pb.Gender_men,
		})
		if err != nil {
			b.Fatal(err)
		}
		count := ultraTable.RemoveWithIdxIntersection(map[string]interface{}{
			"Phone": fmt.Sprintf("+861357546%d", i),
		})
		if count != 1 && ultraTable.Len() != 0 {
			b.Fail()
		}
	}
}

// func BenchmarkConcurrent(b *testing.B) {
// 	b.StopTimer()

// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {

// 		waitGroup := sync.WaitGroup{}
// 		waitGroup.Add(30)

// 		ultraTable := ultra_table.New[*pb.Person]()
// 		for i := 0; i < 10; i++ {
// 			go func(id int) {
// 				ultraTable.Add(&pb.Person{
// 					Name:     "jacky",
// 					Phone:    fmt.Sprintf("+861357546%d", id),
// 					Age:      int32(id),
// 					BirthDay: 19901111,
// 					Gender:   pb.Gender_men,
// 				})
// 				waitGroup.Done()
// 			}(i)
// 		}
// 		for i := 0; i < 10; i++ {
// 			go func(id int) {
// 				ultraTable.UpdateWithIdx(`ID`, id, &pb.Person{
// 					Name:     "jacky",
// 					Phone:    fmt.Sprintf("+861357546%d", id),
// 					Age:      int32(id),
// 					BirthDay: 19901111,
// 					Gender:   pb.Gender_women,
// 				})
// 				waitGroup.Done()
// 			}(i)
// 		}

// 		for i := 0; i < 10; i++ {
// 			go func(id int) {
// 				ultraTable.GetWithIdx("ID", id)
// 				waitGroup.Done()
// 			}(i)
// 		}
// 		waitGroup.Wait()
// 		if ultraTable.Len() != 10 {
// 			b.Fatal(ultraTable.Len(), "!=10")
// 		}
// 	}
// }
