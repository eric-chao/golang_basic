package helper

import (
	"bytes"
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

type OrganizationSchoolGrade struct {
	Id              int64
	SchoolId        int64 `gorm: "type: bigint; not null"`
	SchoolLevel     int8
	SchoolType      int32
	BaseGradeValue  int32
	SchoolGradeName string
	Sort            int32
	SchoolYear      int32 `gorm: "type: smallint; not null"`
	Status          int8
}

func updateAttr(grades []*OrganizationSchoolGrade) {
	funk.Filter(grades, func(g *OrganizationSchoolGrade) bool {
		g.SchoolGradeName = "yyy"
		return true
	})
}

func assign(grade *OrganizationSchoolGrade) {
	fmt.Println(grade)
	grade = &OrganizationSchoolGrade{
		Id: 1003, SchoolId: 3, SchoolGradeName: "zzz",
	}
	fmt.Println(grade)

	//grade.SchoolGradeName = "xxx_yyy_zzz"
}

func TestSliceDefaultLenAndCap(t *testing.T) {
	var intSlice []int64
	var gradeSlice []*OrganizationSchoolGrade
	t.Logf("intSlice len = %d, cap = %d", len(intSlice), cap(intSlice))
	t.Logf("gradeSlice len = %d, cap = %d", len(gradeSlice), cap(gradeSlice))
}

func TestSliceUpdate(t *testing.T) {
	var grade = &OrganizationSchoolGrade{Id: 1003, SchoolId: 3}
	var grades = []*OrganizationSchoolGrade{
		{Id: 1001, SchoolId: 3, SchoolGradeName: "xxx"},
		{Id: 1002, SchoolId: 3, SchoolGradeName: "xxx"},
	}

	assign(grade)
	t.Logf("School Id = %d, SchoolId = %d, SchoolGradeName = %s", grade.Id, grade.SchoolId, grade.SchoolGradeName)

	updateAttr(grades)

	for _, g := range grades {
		t.Logf("School Id = %d, SchoolId = %d, SchoolGradeName = %s", g.Id, g.SchoolId, g.SchoolGradeName)
	}
}

func TestLengthAndCapacity(t *testing.T) {
	var names = make([]string, 0, 8)
	t.Logf("length = %d, capacity = %d", len(names), cap(names))

	names = append(names, "zhang san")
	names = append(names, "li si")
	names = append(names, "wang wu")

	t.Logf("length = %d, capacity = %d", len(names), cap(names))
}

func TestString(t *testing.T) {
	str := "我爱"
	str2 := "golang，"
	// 该方案每次合并会创建一个新的字符串
	var s = str + str2

	var length = len([]rune(s))
	t.Logf("s length = %d", length)
	t.Logf("s' = %s", string([]rune(s)[:length-1]))

	// 该方案更更快，直接连接底层的 []byte
	var buffer bytes.Buffer
	buffer.WriteString(str)
	buffer.WriteString(str2)
	fmt.Println(buffer.String())
}
