package examples

import (
	"encoding/json"
	"fmt"
	"log"
	xss "github.com/duyquang6/go-xss-filter"
)

type SomeStruct struct {
	IntField                int                `json:"intField"`
	StringField             string             `json:"stringField"`
	NestedStructLevel1      NestedStructLevel1 `json:"level1"`
	OtherNestedStructLevel1 NestedStructLevel1 `json:"other_level1"`
}

type NestedStructLevel1 struct {
	StringField        string             `json:"stringField"`
	NestedStructLevel2 NestedStructLevel2 `json:"level2"`
}

type NestedStructLevel2 struct {
	StringField string `json:"stringField"`
}

func main() {
	myStruct := SomeStruct{
		IntField:    42,
		StringField: "<script>alert('foo');</script>",
		NestedStructLevel1: NestedStructLevel1{
			StringField: "<script>alert('foo');</script>",
			NestedStructLevel2: NestedStructLevel2{
				StringField: "<script>alert('foo');</script>",
			},
		},
		OtherNestedStructLevel1: NestedStructLevel1{
			StringField: "<script>alert('foo');</script>",
			NestedStructLevel2: NestedStructLevel2{
				StringField: "<script>alert('foo');</script>",
			},
		},
	}
	payload, err := json.Marshal(&myStruct)
	if err != nil {
		log.Fatalln("Cannot marshal")
	}

	xss.StructEscapeXSS(&myStruct)

	var data map[string]interface{}

	err = json.Unmarshal(payload, &data)
	if err != nil {
		log.Fatalln("Cannot marshal")
	}
	xss.MapEscapeCSS(data)

	fmt.Printf("My struct: %v\n", myStruct.StringField)
	fmt.Printf("My map: %v", data["stringField"])
}
