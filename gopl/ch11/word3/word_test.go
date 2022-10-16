package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	IsPalindrome("A man, a plan, a canal: Panama")
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
