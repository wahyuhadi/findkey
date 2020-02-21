package services

import "fmt"

// RegexCheckCommit function to find malicious in commit  file
// like : key, password stored, credential db , host, port, etc
// if data find it will return string args

func RegexCheckCommit(change string) (string, bool, error) {

	keywords := []string{
		"password",
		"AKIA",
		"ASIA",
		"DB",
		"password",
		"host",
	}
	change = ""
	for _, keyword := range keywords {
		fmt.Println(keyword, change)
		return keyword, true, nil
	}

	return "nil", true, nil
}
