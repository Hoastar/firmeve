package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueInt(t *testing.T) {
	s := []int{
		1, 2, 3, 1, 2, 4,
	}

	expected := []int{
		1, 2, 3, 4,
	}
	assert.Equal(t, expected, UniqueInt(s))
}

func TestUniqueString(t *testing.T) {
	s := []string{
		"a", "b", "a", "string",
	}

	expected := []string{
		"a", "b", "string",
	}
	assert.Equal(t, expected, UniqueString(s))
}

func TestUniqueInterface(t *testing.T) {
	s := []interface{}{
		1, 2, 3, 1, 2, 4, "a", "b", "a", false, true, false, "string",
	}

	expected := []interface{}{
		1, 2, 3, 4, "a", "b", false, true, "string",
	}
	assert.Equal(t, expected, UniqueInterface(s))
}

func TestInInt(t *testing.T) {
	s := []int{
		1, 2, 3, 1, 2, 4,
	}
	assert.Equal(t, true, InInt(s, 1))
	assert.Equal(t, false, InInt(s, 5))
}

func TestInString(t *testing.T) {
	s := []string{
		"a", "b", "a", "string",
	}
	assert.Equal(t, true, InString(s, "a"))
	assert.Equal(t, false, InString(s, "s"))
}

func TestInInterface(t *testing.T) {
	s := []interface{}{
		1, 2, 3, 1, 2, 4, "a", "b", "a", false, true, false, "string",
	}

	assert.Equal(t, true, InInterface(s, "a"))
	assert.Equal(t, false, InInterface(s, "d"))
}
