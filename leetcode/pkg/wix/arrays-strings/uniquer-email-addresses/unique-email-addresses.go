package arrays_strings

import "strings"

type void struct{}

func cleanEmail(email string) string {
	sp := strings.Split(email, "@")
	local := strings.ReplaceAll(strings.Split(sp[0], "+")[0], ".", "")
	return local + "@" + sp[1]
}

func numUniqueEmails(emails []string) int {
	localParts := make(map[string]void)
	for _, email := range emails {
		localParts[cleanEmail(email)] = void{}
	}

	return len(localParts)
}
