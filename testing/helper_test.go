package testing

import (
	"fmt"
	"github.com/qingcc/yi/utils"
	"testing"
)

func Test_Struct2Map(t *testing.T) {
	eg := Eg{
		Age:   23,
		Name:  "example",
		Email: "example@email.com",
		Edu: []struct {
			EduName string
			Year    int
		}{{EduName: "eduName1", Year: 2014}, {EduName: "eduName2", Year: 2007}},
		S: Stru{A: nil, B: 2, C: "3"},
	}
	fmt.Println(utils.ToJson(utils.Struct2Map(eg)))
}

type Eg struct {
	Age   int    `db:"age" json:"_age"`
	Name  string `json:"_name"`
	Email string `json:"-" db:"-"`
	Edu   []struct {
		EduName string
		Year    int
	}
	Inter interface{}
	S     Stru
}
type Stru struct {
	A interface{}
	B int
	C string
}
