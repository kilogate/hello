package word

import "testing"

func TestIsPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试Case1",
			args: args{
				s: "kfk",
			},
			want: true,
		},
		{
			name: "测试Case2",
			args: args{
				s: "AKL",
			},
			want: false,
		},
		{
			name: "测试Case3",
			args: args{
				s: "卡夫卡",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试Case1",
			args: args{
				s: "kfk",
			},
			want: true,
		},
		{
			name: "测试Case2",
			args: args{
				s: "AKL",
			},
			want: false,
		},
		{
			name: "测试Case3",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
