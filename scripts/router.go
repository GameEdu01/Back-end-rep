package scripts

import (
	"fmt"
)

// RouterUrls executes script requested by request
func RouterUrls(getUrl string, postUrl string, time string) error{
	switch reqUrl := getUrl; reqUrl {
	case "test_get1":
		fmt.Println("test_get1 is working properly")
	case "test_get2":
		fmt.Println("test_get2 is working properly")
	default:
		return fmt.Errorf("incorect data in get url")
	}

	switch reqUrl := postUrl; reqUrl {
	case "test_post1":
		fmt.Println("test_post1 is working properly")
	case "test_post2":
		fmt.Println("test_post2 is working properly")
	default:
		return fmt.Errorf("incorect data in post url")
	}

	fmt.Printf(time)
	return nil
}