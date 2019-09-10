package tagjoin

import (
	"reflect"
)

// JoinRuleFunc rule for field joining
// count is the number of elements aprooved to be joined to string
// field is the field being parsed
// the returning string will be placed between fields joining
// the returning bool representes if the joining will happen or not
type JoinRuleFunc func(count int, field interface{}, tagValue string) (string, bool)

// TagJoin join all tags from struct into a single string
func TagJoin(m interface{}, tagName string, joinRule JoinRuleFunc) string {
	whereStr := ""

	reflModelType := reflect.TypeOf(m).Elem()
	reflModelValue := reflect.ValueOf(m).Elem()
	fieldCount := 0

	for i := 0; i < reflModelType.NumField(); i++ {
		fieldTag := reflModelType.Field(i).Tag.Get(tagName)
		fieldValue := reflModelValue.Field(i).Interface()

		if joinStr, shouldJoin := joinRule(fieldCount, fieldValue, fieldTag); shouldJoin {
			whereStr += joinStr
			fieldCount++
		}
	}

	return whereStr
}
