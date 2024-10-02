package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TagsToSqlQueryString(t *testing.T) {
	tags := []string{"a", "b", "c"}
	res := TagsToSqlQueryString(tags)

	exp := `{"a","b","c"}`
	assert.Equal(t, exp, res)
}
