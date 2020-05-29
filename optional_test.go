package optional

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

type Test struct {
	Int    Int
	String String
}

func TestStruct(t *testing.T) {
	fmt.Print("Start Test")
	o := Test{}
	j, _ := json.Marshal(o)
	fmt.Println(string(j))
	var d Test
	json.Unmarshal(j, &d)

	if o.Int.IsPresent() {
		t.Error("dismatch Int")
	}
	if o.String.IsPresent() {
		t.Error("dismatch string")
	}
}

func TestEmptyVal(t *testing.T) {

	str := OfString("")
	if str.IsPresent() {
		t.Fatal("is blank")
	}
	if str.IsNil() {
		t.Fatal("not OT is blank")
	}
	if !str.IsBlank() {
		t.Fatal("is blank")
	}
}

func TestNilVal(t *testing.T) {
	var str String
	if !str.IsNil() {
		t.Fatal("is nil")
	}
	if !str.IsBlank() {
		t.Fatal("is blank")
	}
	if str.IsPresent() {
		t.Fatal("is not present")
	}
}

func TestNormalVal(t *testing.T) {
	str := OfString("T")
	if str.IsNil() {
		t.Fatal("not is nil")
	}
	if str.IsBlank() {
		t.Fatal("not is blank")
	}
	if !str.IsPresent() {
		t.Fatal("is present")
	}
}

func TestEmptyIntVal(t *testing.T) {
	str := OfInt(0)
	if str.IsNil() {
		t.Fatal("not is nil")
	}
	if str.IsBlank() {
		t.Fatal("not is blank")
	}
	if !str.IsPresent() {
		t.Fatal("is present")
	}
}

func TestFloatVal(t *testing.T) {
	type A struct {
		A Float64
	}
	type B struct {
		A String
	}
	a := A{}
	b := B{}

	if err := copier.Copy(&b, a); err != nil {
		t.Fatal("error:" + err.Error())
	}
	s, _ := json.Marshal(b)
	fmt.Println("tttt", b, string(s))
}


func TestNumberToBool(t *testing.T) {
	a := "1"
	var b Bool
	if err := json.Unmarshal([]byte(a),&b) ; err != nil {
		t.Fatal(err)
	}
}