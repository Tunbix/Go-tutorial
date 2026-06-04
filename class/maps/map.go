package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "google.com",
		"Facebook": "facebook.com",
		"Twitter":  "twitter.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Facebook"])

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Twitter")
	fmt.Println(websites)
}
