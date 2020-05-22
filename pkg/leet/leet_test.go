package leet_test

import (
	"testing"

	"github.com/hsmtkk/c016_leet/pkg/leet"
	"github.com/stretchr/testify/assert"
)

func TestLeet(t *testing.T) {
	want := "P4124"
	got, err := leet.Leet("PAIZA")
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")

	want = "4L4NTUR1N6"
	got, err = leet.Leet("ALANTURING")
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}
