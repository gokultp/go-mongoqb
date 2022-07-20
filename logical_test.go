package mongoqb

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestLogicalQuery_Build(t *testing.T) {

	tests := []struct {
		name string
		q    *LogicalQuery
		want bson.M
	}{
		{
			name: "test or query",
			q:    NewLogicalQuery(operationOr, NewComparisonQuery(operationEq, "a", 1), NewComparisonQuery(operationGt, "b", 0)),
			want: bson.M{
				"$or": []bson.M{{"a": 1}, {"b": bson.M{"$gt": 0}}},
			},
		},
		{
			name: "test simple and  query",
			q:    NewLogicalQuery(operationAnd, NewComparisonQuery(operationEq, "a", 1), NewComparisonQuery(operationGt, "b", 0)),
			want: bson.M{
				"a": 1,
				"b": bson.M{"$gt": 0},
			},
		},
		{
			name: "test complex and  query",
			q:    NewLogicalQuery(operationAnd, NewComparisonQuery(operationEq, "a", 1), NewComparisonQuery(operationGt, "a", 0)),
			want: bson.M{
				"$and": []bson.M{{"a": 1}, {"a": bson.M{"$gt": 0}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.q.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogicalQuery.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
