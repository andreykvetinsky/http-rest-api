package model_test

import (
	"testing"

	"github.com/andreykvetinsky/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestNote_ValidateBeforeCreate(t *testing.T) {
	testCases := []struct {
		name string
		n    *model.Note
		res  string
	}{
		{
			name: "first_test",
			n:    model.TestNote(t),
			res:  "conquer Everest",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.n.ValidateBeforeCreate())
			assert.Equal(t, tc.res, tc.n.Note)
		})
	}
}
