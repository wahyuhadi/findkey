package services

import "fmt"

// RegexCheckCommit function to find malicious in commit  file
// like : key, password stored, credential db , host, port, etc
// if data find it will return string args
type Found struct {
	keyword string
}

func RegexCheckCommit(change string) (string, bool, error) {

	var arr []Found

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
		data := Found{"keyword": keyword}
		arr = append(arr, data)
		// fmt.Println(keyword, change)
		// return keyword, true, nil
	}
	fmt.Println(arr)
	return "nil", true, nil
}
