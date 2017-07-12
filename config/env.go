package config

import (
	"os"
	"reflect"
	"strconv"
)

func fromEnv(config interface{}) error {
	configPtrType := reflect.TypeOf(config)
	configPtrValue := reflect.ValueOf(config)

	if configPtrType.Kind() != reflect.Ptr {
		return &EnvError{
			message: "config must be a pointer",
		}
	}

	configType := configPtrType.Elem()
	configValue := configPtrValue.Elem()

	if configType.Kind() != reflect.Struct {
		return &EnvError{
			message: "config must be a pointer to a struct",
		}
	}

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		target := configValue.Field(i)

		err := fieldFromEnv(target, field)
		if err != nil {
			return err
		}
	}

	return nil
}

func fieldFromEnv(target reflect.Value, field reflect.StructField) error {
	tag, ok := field.Tag.Lookup("env")

	if !ok {
		return &EnvError{
			message: "Missing 'env' tag",
			data: map[string]interface{}{
				"field": field.Name,
			},
		}
	}

	raw := os.Getenv(tag)
	value, err := parseField(raw, field, field.Type)
	if err != nil {
		return err
	}

	target.Set(value)

	return nil
}

func parseField(raw string, field reflect.StructField, typ reflect.Type) (reflect.Value, error) {
	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(raw), nil
	case reflect.Int:
		val, err := strconv.Atoi(raw)
		return reflect.ValueOf(val), err
	}

	return reflect.ValueOf(nil), &EnvError{
		message: "Unknown type on field",
		data: map[string]interface{}{
			"field": field.Name,
			"type":  typ.Name(),
		},
	}
}
