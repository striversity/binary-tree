package tree

import (
	"testing"
)

func TestNode_Search(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args []args
	}{
		{name: "instert ace", args: []args{{"a"}, {"c"}, {"e"}}},
	}
	tree := &Node{Key: "f"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tree.Insert(arg.key)
			}
		})
	}
	for _, key := range []string{"f", "c"} {
		node, err := tree.Search(key)
		if err != nil || node.Key != key {
			t.Failed()
		}
	}
	_, err := tree.Search("a")
	if err == nil {
		t.Failed()
	}
}
