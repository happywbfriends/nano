package config

import (
	"reflect"
	"testing"
)

func TestParseConfigLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		// errors
		{"", args{"abc"}, nil, true},
		{"", args{"abc='"}, nil, true},
		{"", args{"=1"}, nil, true},
		// good
		{"", args{"a=1 b="}, map[string]string{"a": "1", "b": ""}, false},
		{"", args{"a= b="}, map[string]string{"a": "", "b": ""}, false},
		{"", args{"a=1 b=2"}, map[string]string{"a": "1", "b": "2"}, false},
		{"", args{"a=' b=3'"}, map[string]string{"a": " b=3"}, false},
		{"", args{`a=" b=3"`}, map[string]string{"a": " b=3"}, false},
		{"", args{`a=777   b="  "    c='""'`}, map[string]string{"a": "777", "b": "  ", "c": `""`}, false},
		{"", args{`   a=777b=9   `}, map[string]string{"a": "777b=9"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfigLine(tt.args.s, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfigLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfigLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
