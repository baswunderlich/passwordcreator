package bapoPass

import "crypto/sha512"

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'="

func GeneratePassword(seed string, plattform string) string {
	string_to_hash := seed + plattform
	hasing_result := sha512.Sum512([]byte(string_to_hash))
	result_string := ""
	for _, e := range hasing_result {
		result_string += string(chars[int64(e)%int64(len(chars))])
	}
	return result_string
}
