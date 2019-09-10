package tagjoin

import "testing"

func TestTagJoin(t *testing.T) {
	type modeltest struct {
		Name     *string `db:"name"`
		LastName string  `db:"last_name"`
		Age      *int64  `db:"age"`
	}

	mt := &modeltest{
		LastName: "Manuel",
	}
	age := int64(19)
	mt.Age = &age

	type args struct {
		m        interface{}
		tagName  string
		joinRule JoinRuleFunc
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				m:        mt,
				tagName:  "db",
				joinRule: WhereJoinNonNil,
			},
			want: "last_name = :last_name and age = :age",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagJoin(tt.args.m, tt.args.tagName, tt.args.joinRule); got != tt.want {
				t.Errorf("TagJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
