package vec2

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		vec      F
		expected string
	}{
		{
			name:     "zero vector",
			vec:      F{X: 0, Y: 0},
			expected: "[0,0]",
		},
		{
			name:     "positive values",
			vec:      F{X: 1.5, Y: 2.5},
			expected: "[1.5,2.5]",
		},
		{
			name:     "negative values",
			vec:      F{X: -1.5, Y: -2.5},
			expected: "[-1.5,-2.5]",
		},
		{
			name:     "mixed values",
			vec:      F{X: -3.14, Y: 2.71},
			expected: "[-3.14,2.71]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.vec)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(result))
		})
	}
}

func TestFUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected F
		wantErr  bool
	}{
		{
			name:     "zero vector",
			input:    "[0,0]",
			expected: F{X: 0, Y: 0},
			wantErr:  false,
		},
		{
			name:     "positive values",
			input:    "[1.5,2.5]",
			expected: F{X: 1.5, Y: 2.5},
			wantErr:  false,
		},
		{
			name:     "negative values",
			input:    "[-1.5,-2.5]",
			expected: F{X: -1.5, Y: -2.5},
			wantErr:  false,
		},
		{
			name:     "mixed values",
			input:    "[-3.14,2.71]",
			expected: F{X: -3.14, Y: 2.71},
			wantErr:  false,
		},
		{
			name:     "fallback to object notation",
			input:    `{"X":10,"Y":20}`,
			expected: F{X: 10, Y: 20},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result F
			err := json.Unmarshal([]byte(tt.input), &result)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.X, result.X)
				assert.Equal(t, tt.expected.Y, result.Y)
			}
		})
	}
}

func TestIMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		vec      I
		expected string
	}{
		{
			name:     "zero vector",
			vec:      I{X: 0, Y: 0},
			expected: "[0,0]",
		},
		{
			name:     "positive values",
			vec:      I{X: 10, Y: 20},
			expected: "[10,20]",
		},
		{
			name:     "negative values",
			vec:      I{X: -10, Y: -20},
			expected: "[-10,-20]",
		},
		{
			name:     "mixed values",
			vec:      I{X: -15, Y: 30},
			expected: "[-15,30]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.vec)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(result))
		})
	}
}

func TestIUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected I
		wantErr  bool
	}{
		{
			name:     "zero vector",
			input:    "[0,0]",
			expected: I{X: 0, Y: 0},
			wantErr:  false,
		},
		{
			name:     "positive values",
			input:    "[10,20]",
			expected: I{X: 10, Y: 20},
			wantErr:  false,
		},
		{
			name:     "negative values",
			input:    "[-10,-20]",
			expected: I{X: -10, Y: -20},
			wantErr:  false,
		},
		{
			name:     "mixed values",
			input:    "[-15,30]",
			expected: I{X: -15, Y: 30},
			wantErr:  false,
		},
		{
			name:     "fallback to object notation",
			input:    `{"X":100,"Y":200}`,
			expected: I{X: 100, Y: 200},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result I
			err := json.Unmarshal([]byte(tt.input), &result)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.X, result.X)
				assert.Equal(t, tt.expected.Y, result.Y)
			}
		})
	}
}

func TestDMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		vec      D
		expected string
	}{
		{
			name:     "zero vector",
			vec:      D{X: 0, Y: 0},
			expected: "[0,0]",
		},
		{
			name:     "positive values",
			vec:      D{X: 1.5, Y: 2.5},
			expected: "[1.5,2.5]",
		},
		{
			name:     "negative values",
			vec:      D{X: -1.5, Y: -2.5},
			expected: "[-1.5,-2.5]",
		},
		{
			name:     "mixed values",
			vec:      D{X: -3.14159265359, Y: 2.71828182846},
			expected: "[-3.14159265359,2.71828182846]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.vec)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(result))
		})
	}
}

func TestDUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected D
		wantErr  bool
	}{
		{
			name:     "zero vector",
			input:    "[0,0]",
			expected: D{X: 0, Y: 0},
			wantErr:  false,
		},
		{
			name:     "positive values",
			input:    "[1.5,2.5]",
			expected: D{X: 1.5, Y: 2.5},
			wantErr:  false,
		},
		{
			name:     "negative values",
			input:    "[-1.5,-2.5]",
			expected: D{X: -1.5, Y: -2.5},
			wantErr:  false,
		},
		{
			name:     "mixed values",
			input:    "[-3.14159265359,2.71828182846]",
			expected: D{X: -3.14159265359, Y: 2.71828182846},
			wantErr:  false,
		},
		{
			name:     "fallback to object notation",
			input:    `{"X":10.5,"Y":20.5}`,
			expected: D{X: 10.5, Y: 20.5},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result D
			err := json.Unmarshal([]byte(tt.input), &result)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.X, result.X)
				assert.Equal(t, tt.expected.Y, result.Y)
			}
		})
	}
}

// TestRoundTripF tests marshaling and then unmarshaling F type
func TestRoundTripF(t *testing.T) {
	original := F{X: 3.14, Y: -2.71}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	var result F
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, original.X, result.X)
	assert.Equal(t, original.Y, result.Y)
}

// TestRoundTripI tests marshaling and then unmarshaling I type
func TestRoundTripI(t *testing.T) {
	original := I{X: 42, Y: -99}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	var result I
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, original.X, result.X)
	assert.Equal(t, original.Y, result.Y)
}

// TestRoundTripD tests marshaling and then unmarshaling D type
func TestRoundTripD(t *testing.T) {
	original := D{X: 3.141592653589793, Y: -2.718281828459045}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	var result D
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, original.X, result.X)
	assert.Equal(t, original.Y, result.Y)
}

// TestMarshalInStruct tests marshaling vectors as fields in a struct
func TestMarshalInStruct(t *testing.T) {
	type TestStruct struct {
		FloatVec  F `json:"float"`
		IntVec    I `json:"int"`
		DoubleVec D `json:"double"`
	}

	original := TestStruct{
		FloatVec:  F{X: 1.5, Y: 2.5},
		IntVec:    I{X: 10, Y: 20},
		DoubleVec: D{X: 3.14, Y: 2.71},
	}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	expected := `{"float":[1.5,2.5],"int":[10,20],"double":[3.14,2.71]}`
	assert.JSONEq(t, expected, string(data))

	var result TestStruct
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, original.FloatVec, result.FloatVec)
	assert.Equal(t, original.IntVec, result.IntVec)
	assert.Equal(t, original.DoubleVec, result.DoubleVec)
}

// TestMarshalSlice tests marshaling slices of vectors
func TestMarshalSlice(t *testing.T) {
	floatVecs := []F{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
		{X: 5, Y: 6},
	}

	data, err := json.Marshal(floatVecs)
	assert.NoError(t, err)

	expected := `[[1,2],[3,4],[5,6]]`
	assert.JSONEq(t, expected, string(data))

	var result []F
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Len(t, result, 3)
	assert.Equal(t, floatVecs, result)
}
