package lee_code

import "testing"

func TestIsPalindromeNumber(t *testing.T)  {
	t.Log(isPalindrome(1))
	t.Log(isPalindrome(10))
	t.Log(isPalindrome(121))
	t.Log(isPalindrome(1221))
	t.Log(isPalindrome(12321))

	t.Log(isPalindrome_Better(1))
	t.Log(isPalindrome_Better(10))
	t.Log(isPalindrome_Better(121))
	t.Log(isPalindrome_Better(1221))
	t.Log(isPalindrome_Better(12321))
}