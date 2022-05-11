package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	result := ""
// 	for i := len(word) - 1; i >= 0; i-- {
// 		result += string(word[i])
// 	}
// 	return result
// }

func Reverse(word string) string {
	result := ""
	for _, r := range word {
		result = string(r) + result
	}
	return result
}
