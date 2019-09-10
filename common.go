package tagjoin

import (
	"fmt"
	"reflect"
)

// WhereJoinNonNil tag joining commonly used on selects, only non nil values will be included
func WhereJoinNonNil(count int, field interface{}, tagValue string) (string, bool) {
	if field == nil || (reflect.ValueOf(field).Kind() == reflect.Ptr && reflect.ValueOf(field).IsNil()) {
		return "", false
	}

	if count != 0 {
		return fmt.Sprintf(" and %s = :%s", tagValue, tagValue), true
	}

	return fmt.Sprintf("%s = :%s", tagValue, tagValue), true
}

// WhereJoin tag joining commonly used on selects
func WhereJoin(count int, field interface{}, tagValue string) (string, bool) {
	if count != 0 {
		return fmt.Sprintf(" and %s = :%s", tagValue, tagValue), true
	}

	return fmt.Sprintf("%s = :%s", tagValue, tagValue), true
}

// CommaJoinNonNil tag joining with commas, general purpose, only non nil values will be included
func CommaJoinNonNil(count int, field interface{}, tagValue string) (string, bool) {
	if field == nil || (reflect.ValueOf(field).Kind() == reflect.Ptr && reflect.ValueOf(field).IsNil()) {
		return "", false
	}

	if count != 0 {
		return fmt.Sprintf(", %s", tagValue), true
	}

	return fmt.Sprintf("%s", tagValue), true
}

// CommaJoin tag joining with commas, general purpose
func CommaJoin(count int, field interface{}, tagValue string) (string, bool) {
	if count != 0 {
		return fmt.Sprintf(", %s", tagValue), true
	}

	return fmt.Sprintf("%s", tagValue), true
}
