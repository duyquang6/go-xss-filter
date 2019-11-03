package xss

import (
	"html"
	"log"
	"reflect"
)

func StructEscapeXSS(myStruct interface{}) {
	value := reflect.ValueOf(myStruct)
	if value.Kind() == reflect.Ptr {
		value = reflect.Indirect(value)
	}

	if value.Kind() != reflect.Struct {
		log.Fatalln("This is not the struct we want")
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.Struct {
			StructEscapeXSS(field.Addr().Interface())
		}
		if field.Type() != reflect.TypeOf("") {
			continue
		}
		str := field.Interface().(string)
		field.SetString(html.EscapeString(str))
	}
}

// Type safe
func MapEscapeCSS(myMap map[string]interface{}) {
	for key, value := range myMap {
		switch value.(type) {
		case string:
			myMap[key] = html.EscapeString(value.(string))
		case map[string]interface{}:
			MapEscapeCSS(value.(map[string]interface{}))
		}
	}
}
