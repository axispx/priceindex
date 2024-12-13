package utils

import "strings"

const ANTIaddress = "HB8KrN7Bb3iLWUPsozp67kS4gxtbA4W5QJX4wKPvpump"
const PROaddress = "CWFa2nxUMf5d1WwKtG9FS9kjUKGwKXWSjH8hFdWspump"

func GetTokenAddress(tokens ...string) []string {
	addresses := []string{}
	for _, token := range tokens {
		token = strings.TrimSpace(strings.ToLower(token))

		if token == "anti" {
			addresses = append(addresses, ANTIaddress)
		} else if token == "pro" {
			addresses = append(addresses, PROaddress)
		} else {
			addresses = append(addresses, token)
		}
	}

	return addresses
}
