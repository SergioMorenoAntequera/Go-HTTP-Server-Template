package models

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Model struct {
	_modelName string
	Id         string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

func (m *Model) Init(modelName string) {
	m._modelName = modelName
	m.UpdatedAt = time.Now()
	if (m.CreatedAt == time.Time{}) {
		m.CreatedAt = time.Now()
	}

}

type ModelBreakdown struct {
	typeGo string
	name   string
	value  string
}

var TYPE_TO_STRING_CONVERTION = map[string]func(value reflect.Value) string{
	"int":  func(value reflect.Value) string { return strconv.FormatInt(value.Int(), 10) },
	"Time": func(value reflect.Value) string { return value.Interface().(time.Time).Format("2006-01-02") },
}

var TYPES_WITH_QUOTES = []string{"string", "Time"}

func extractAttributes(model any) []ModelBreakdown {
	keys := []ModelBreakdown{}

	baseType := reflect.Indirect(reflect.ValueOf(model))
	for i := range baseType.NumField() {

		var attributeValue reflect.Value
		var attribute reflect.StructField

		if reflect.ValueOf(model).Kind() == reflect.Pointer {
			attributeValue = reflect.ValueOf(model).Elem().Field(i)
			attribute = reflect.TypeOf(model).Elem().Field(i)
		}

		if reflect.ValueOf(model).Kind() == reflect.Struct {
			attributeValue = reflect.ValueOf(model).Field(i)
			attribute = reflect.TypeOf(model).Field(i)
		}

		if attribute.Type.String() == "models.Model" {
			keys = append(keys, extractAttributes(attributeValue.Interface().(Model))...)
		}

		if attribute.Type.String() != "models.Model" {

			var convertedType string
			converter, found := TYPE_TO_STRING_CONVERTION[attribute.Type.Name()]
			if found {
				convertedType = converter(attributeValue)
			}
			if !found {
				convertedType = attributeValue.String()
			}

			keys = append(keys, ModelBreakdown{
				typeGo: attribute.Type.Name(),
				name:   attribute.Name,
				value:  convertedType,
			})
		}

	}

	return keys
}

func Create(m any) {

	attributes := extractAttributes(m)
	modelName := slices.IndexFunc(attributes, func(a ModelBreakdown) bool { return a.name == "_modelName" })

	insertQueryHeader := "INSERT INTO " + attributes[modelName].value + " ("
	insertQueryFooter := "VALUES ("

	for i, attribute := range attributes {

		if strings.HasPrefix(attribute.name, "_") {
			continue
		}

		var usesQuotes = slices.IndexFunc(TYPES_WITH_QUOTES, func(a string) bool { return a == attribute.typeGo }) != -1
		if !usesQuotes {
			insertQueryFooter += attribute.value
		} else {
			insertQueryFooter += "'" + attribute.value + "'"
		}
		insertQueryHeader += attribute.name

		if i != len(attributes)-1 {
			insertQueryHeader += ", "
			insertQueryFooter += ", "
		}
		if i == len(attributes)-1 {
			insertQueryHeader += ")"
			insertQueryFooter += ");"
		}
	}

	fmt.Println(insertQueryHeader, insertQueryFooter)

}
