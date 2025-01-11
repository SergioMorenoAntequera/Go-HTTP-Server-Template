package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Model struct {
	id         string
	_modelName string
	updatedAt  time.Time
	createdAt  time.Time
}

func (m *Model) Init(modelName string) {
	m._modelName = modelName
	m.updatedAt = time.Now()
	if (m.createdAt == time.Time{}) {
		m.createdAt = time.Now()
	}
}

type ModelBreakdown struct {
	typeGo string
	name   string
	value  any
}

func extractAttributes(model interface{}) []ModelBreakdown {
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
			keys = append(keys, ModelBreakdown{
				typeGo: attribute.Type.Name(),
				name:   attribute.Name,
				value:  attributeValue,
			})
		}

	}

	return keys
}

func Create(m interface{}) {

	insertQueryHeader := "INSERT INTO " + " ("
	insertQueryFooter := "VALUES ("

	attributes := extractAttributes(m)

	for i, attribute := range attributes {

		if strings.HasPrefix(attribute.name, "_") {
			continue
		}

		insertQueryHeader += attribute.name
		if i != len(attributes)-1 {
			insertQueryHeader += ", "
		}
		if i == len(attributes)-1 {
			insertQueryHeader += ")"
		}
	}

	fmt.Println(insertQueryHeader, insertQueryFooter)

}
