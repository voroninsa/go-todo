package common

import "testing"

func Test_TagsToSqlQueryString(t *testing.T) {
	tags := []string{"a", "b", "c"}
	res := TagsToSqlQueryString(tags)
	t.Log(res)
}
