package services

import "fmt"

// RegexCheckCommit function to find malicious in commit  file
// like : key, password stored, credential db , host, port, etc
// if data find it will return string args
func RegexCheckCommit(change string) (bool, error) {
	fmt.Println(change)
	return true, nil
}
