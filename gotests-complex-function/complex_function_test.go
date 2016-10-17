package monkey

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_getversions(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			"test1",
			[]string{"123", "3de", "456", "777"},
		},
	}
	for _, tt := range tests {
		if got := getversions(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getversions() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		input  InputJSON
		logger *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    checkOutputJSON
		wantErr bool
	}{
		{
			"test1",
			args{
				input: InputJSON{
					Params: map[string]string{"param1": "Hello Jeff", "param2": "How are you?"},
					Source: map[string]string{"source1": "sourcefoo1", "source2": "sourcefoo2"},
					Version: version{
						Ref: "456",
					},
				},
				logger: log.New(os.Stderr, "resource:", log.Lshortfile),
			},
			checkOutputJSON{
				version{Ref: "123"},
				version{Ref: "3de"},
				version{Ref: "456"},
				version{Ref: "777"},
			},
			false,
		},
	}
	for _, tt := range tests {
		got, err := Check(tt.args.input, tt.args.logger)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Check() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Check() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
