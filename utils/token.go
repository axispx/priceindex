package utils

import "strings"

const ANTIaddress = "HB8KrN7Bb3iLWUPsozp67kS4gxtbA4W5QJX4wKPvpump"
const PROaddress = "CWFa2nxUMf5d1WwKtG9FS9kjUKGwKXWSjH8hFdWspump"

func GetTokenAddresses(tokens ...string) []string {
	addresses := []string{}
	for _, token := range tokens {
		address := GetTokenAddress(token)
		addresses = append(addresses, address)
	}

	return addresses
}

func GetTokenAddress(token string) string {
	if strings.ToLower(token) == "anti" {
		return ANTIaddress
	} else if strings.ToLower(token) == "pro" {
		return PROaddress
	} else {
		return token
	}
}
