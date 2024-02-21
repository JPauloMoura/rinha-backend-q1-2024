package entities

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_Validate(t *testing.T) {
	type fields struct {
		UserId      int
		Value       int
		Type        string
		Description string
		CreatedAt   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "should return err when value < 0",
			fields:  fields{Value: -1},
			wantErr: true,
		},
		{
			name:    "should return err when value is igual 0",
			fields:  fields{Value: 0},
			wantErr: true,
		},
		{
			name:    "should return err when type is empty",
			fields:  fields{Value: 1, Type: ""},
			wantErr: true,
		},
		{
			name:    "should return err when type is not 'c' or 'd'",
			fields:  fields{Value: 1, Type: "x"},
			wantErr: true,
		},
		{
			name:    "should return err when description is empty",
			fields:  fields{Value: 1, Type: "c", Description: ""},
			wantErr: true,
		},
		{
			name:    "should return err when description lenght > 10",
			fields:  fields{Value: 1, Type: "d", Description: strings.Repeat("B", 11)},
			wantErr: true,
		},
		{
			name:    "sucess",
			fields:  fields{Value: 1, Type: "d", Description: strings.Repeat("B", 10)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				Value:       tt.fields.Value,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
			}
			assert.Equal(t, tt.wantErr, tr.Validate() != nil)
		})
	}
}
