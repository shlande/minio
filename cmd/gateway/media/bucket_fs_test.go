package media

import "testing"

func TestFileBucket_isMapped(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "other", args: "/other", want: ""},
		{name: "match", args: "/mapped/test", want: "test"},
		{name: "match/sub", args: "/mapped/test/sub", want: "test/sub"},
		{name: "other-bucket", args: "/ofsef/test/sub", want: ""},
	}
	f := FileBucket{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.isMapped(tt.args); got != tt.want {
				t.Errorf("isMapped() = %v, want %v", got, tt.want)
			}
		})
	}
}
