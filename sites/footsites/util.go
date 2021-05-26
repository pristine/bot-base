package footsites

import "fmt"

func HandleStatusCodes(statusCode int) string {
	switch statusCode {
	case 400:
		return "Bad Request"
	case 403:
		return "Banned"
	case 429:
		return "Rate Limited"
	case 529:
		return "Queue"
	case 531:
		return "Out of Stock"
	default:
		return fmt.Sprintf("%d", statusCode)
	}
}
