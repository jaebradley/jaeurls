package keyhandler

import "testing"

func TestEncode(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"should return empty string", args{value: 0}, ""},
		{"should return 1", args{value: 1}, "1"},
		{"should return j1", args{value: 81}, "1j"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.value); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"should return first character", args{value: "0"}, 0},
		{"should return 81", args{value: "1j"}, 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.value); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateIndex(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"zero index", args{value: 0}, 0},
		{"another zero index", args{value: uint64(len(alphabet))}, 0},
		{ "non-zero index", args{value: 108}, 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateIndex(tt.args.value); got != tt.want {
				t.Errorf("calculateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
