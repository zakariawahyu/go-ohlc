package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFiles(t *testing.T) {
	data, err := LoadFiles()

	fmt.Println(data)
	assert.NotNil(t, data, "data should not be nil")
	assert.NoError(t, err, "data should not be error")
}
