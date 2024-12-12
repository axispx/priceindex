package utils

import "strings"

const ANTIaddress = "HB8KrN7Bb3iLWUPsozp67kS4gxtbA4W5QJX4wKPvpump"
const PROaddress = "CWFa2nxUMf5d1WwKtG9FS9kjUKGwKXWSjH8hFdWspump"

func GetTokenAddress(token string) string {
	token = strings.TrimSpace(strings.ToLower(token))

	if token == "anti" {
		return ANTIaddress
	}

	return PROaddress
}
