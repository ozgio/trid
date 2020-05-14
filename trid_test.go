package trid

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		id  string
		err error
	}{
		{"43848743180", nil},
		{"88733835038", nil},
		{"17786671430", nil},
		{"95111173064", nil},
		{"83579056890", nil},
		{"", errors.New("invalid length")},
		{"03579056890", errors.New("id cannot start with 0")},
		{"83579056890 ", errors.New("invalid length")},
		{" 83579056890", errors.New("invalid length")},
		{"83579056891", errors.New("invalid checksum, id should end with 90")},
		{"83579056A90", errors.New("'A' at digit 9 is not a number")},
		{"835790568A0", errors.New("invalid checksum, id should end with 90")},
		{"8357905689A", errors.New("invalid checksum, id should end with 90")},
		{"11111111111", errors.New("invalid checksum, id should end with 10")},
	}

	for i, tt := range tests {
		err := Validate(tt.id)
		switch {
		case err == nil && tt.err == nil:
			continue
		case err == nil || tt.err == nil:
			fallthrough
		case err.Error() != tt.err.Error():
			t.Errorf("%d: expecting '%v' as error but it returned '%v'", i, tt.err, err)
		}
	}
}
