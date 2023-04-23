package goltcd

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case01",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "case02",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "case03",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
