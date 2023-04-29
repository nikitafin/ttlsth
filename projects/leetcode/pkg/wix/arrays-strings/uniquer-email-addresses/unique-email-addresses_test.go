package arrays_strings

import "testing"

func TestUniqueEmailAddressEmpty(t *testing.T) {
	if numUniqueEmails([]string{}) != 0 {
		t.Error("Error: []string{\"\"} != 0")
	}
}

func TestUniqueEmailAddressOne(t *testing.T) {
	if numUniqueEmails([]string{"test@mail.com"}) != 1 {
		t.Error("Error: []string{\"test@mail.com\"} != 1")
	}
}

func TestUniqueEmailAddressTwo(t *testing.T) {
	i := numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	if i != 2 {
		t.Errorf(`Error: []string{"test.email+alex@leetcode.com",
			 "test.e.mail+bob.cathy@leetcode.com",
			 "testemail+david@lee.tcode.com"} != 2 (%d)`, i)
	}
}

func TestUniqueEmailAddressThree(t *testing.T) {
	i := numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"})
	if i != 3 {
		t.Errorf(`Error: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"} != 3 (%d)`, i)
	}
}
