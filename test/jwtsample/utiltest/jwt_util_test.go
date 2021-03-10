package utiltest

import (
	"Golang-JWT/internal/jwtsample/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleSignAndValidate(t *testing.T) {

	var data = util.Data{
		Name:       "sammidev",
		Identifier: "2003113948",
		Email:      "sammidev@gmail.com",
	}

	var token = util.JwtBuildAndSignJSON(data)
	var validatedData, _ = util.JwtValidate(token)
	assert.Equal(t, data, validatedData)
}