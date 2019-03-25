package uniqueEmailAddress

import "testing"

func TestUniqueEmailAddress(t *testing.T) {

	emails := []string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}
	t.Log(numUniqueEmails(emails))
}

func TestUniqueEmailAddressPro(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}
	t.Log(numUniqueEmailsPro(emails))
}
