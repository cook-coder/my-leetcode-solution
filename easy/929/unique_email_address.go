package uniqueEmailAddress

import (
	"strings"
)

func numUniqueEmails(emails []string) int {

	var uniqueEmails = map[string]bool{}
	for _, email := range emails {
		emailComponents := strings.Split(email, "@")
		localName := emailComponents[0]
		if plusIndex := strings.Index(localName, "+"); plusIndex != -1 {
			localName = localName[0:plusIndex]
		}
		localName = strings.Replace(localName, ".", "", -1)
		emailComponents[0] = localName
		uniqueEmails[strings.Join(emailComponents, "@")] = true
	}
	return len(uniqueEmails)
}

func numUniqueEmailsV2(emails []string) int {

	var uniqueEmails = map[string]bool{}
	for _, email := range emails {
		localName := email[0:strings.Index(email, "@")]
		domainName := email[strings.Index(email, "@")+1:]
		if plusIndex := strings.Index(localName, "+"); plusIndex > -1 {
			localName = localName[0:plusIndex]
		}
		if dotIndex := strings.Index(localName, "."); dotIndex > -1 {
			localName = strings.Replace(localName, ".", "", -1)
		}
		uniqueEmails[localName+"@"+domainName] = true
	}
	return len(uniqueEmails)
}

func numUniqueEmailsPro(emails []string) int {
	var uniqueEmails = map[string]bool{}
	var indexOfAt int8
	var localName string
	var plusIndex int8 = -1
	var dotIndex int8 = -1
	for i := 0; i < len(emails); i++ {
		indexOfAt = int8(strings.Index(emails[i], "@"))
		localName = emails[i][0:indexOfAt]
		if plusIndex = int8(strings.Index(localName, "+")); plusIndex > -1 {
			localName = localName[0:plusIndex]
		}
		if dotIndex = int8(strings.Index(localName, ".")); dotIndex > -1 {
			localName = strings.Replace(localName, ".", "", -1)
		}
		uniqueEmails[localName+"@"+emails[i][indexOfAt+1:]] = true
	}
	return len(uniqueEmails)
}
