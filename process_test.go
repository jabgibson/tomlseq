package tomlseq

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	type args struct {
		identifier string
		data       []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "no frills add sequence",
			args: args{
				identifier: "seq",
				data:       []byte(tomlExample1),
			},
			want: []byte(tomlExpected1),
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Process(tt.args.identifier, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

const tomlExample1 = `
[[item]]
value="boom"
count=7

[[item]]
value="bam"
count=3
`

const tomlExpected1 = `
[[item]]
seq = 0
value="boom"
count=7

[[item]]
seq = 1
value="bam"
count=3
`