package data

import (
	"encoding/json"
	"fmt"
)

// Type denotes a data type
type Type int

const (
	STRING Type = iota + 1
	INTEGER
	NUMBER
	BOOLEAN
	OBJECT
	ARRAY
	PARAMS
	ANY
)

var types = [...]string{
	"string",
	"integer",
	"number",
	"boolean",
	"object",
	"array",
	"params",
	"any",
}

var typeMap = map[string]Type{
	"string":  STRING,
	"integer": INTEGER,
	"number":  NUMBER,
	"boolean": BOOLEAN,
	"object":  OBJECT,
	"array":   ARRAY,
	"params":  PARAMS,
	"any":     ANY,
}

func (t Type) String() string {
	return types[t-1]
}

// ToTypeEnum get the data type that corresponds to the specified name
func ToTypeEnum(typeStr string) (Type, bool) {
	dataType, found := typeMap[typeStr]

	return dataType, found
}

// GetType get the Type of the supplied value
func GetType(val interface{}) (Type, error) {

	switch t := val.(type) {
	case string:
		return STRING, nil
	case int:
		return INTEGER, nil
	case float64:
		return NUMBER, nil
	case json.Number:
		return NUMBER, nil
	case bool:
		return BOOLEAN, nil
	case map[string]interface{}:
		return OBJECT, nil
	case []interface{}:
		return ARRAY, nil
	default:
		return 0, fmt.Errorf("Unable to determine type of %#v", t)
	}
}
