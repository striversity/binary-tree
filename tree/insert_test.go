package tree

import (
	"testing"
)

func TestNode_Insert(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args []args
	}{
		{name: "insert face", args: []args{{"a"}, {"c"}, {"e"}}},
	}
	for _, tt := range tests {
		tree := &Node{Key: "f"}
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tree.Insert(arg.key)
			}
		})
	}
}
