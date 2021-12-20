// Create a network speed test cli

package main

// Import packages
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Define variables
var (
	urls = []string{
		"https://www.google.com",
		"https://www.cox.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.baidu.com",
		"https://www.qq.com",
		"https://www.taobao.com",
		"https://www.live.com",
		"https://www.tmall.com",
		"https://www.yahoo.com",
		"https://www.sohu.com",
		"https://www.sina.com.cn",
		"https://www.pinterest.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.ebay.com",
		"https://www.paypal.com",
		"https://www.microsoft.com/en-us/store/apps",
		"https://www.microsoft.com/en-us/windows/microsoft-edge",
		"https://www.groupon.com",
	}

	// Define a function to run the test
	runTest = func(url string) (time.Duration, error) {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		return time.Since(start), nil
	}
)

// Define a function to print the results
func printResults(url string, duration time.Duration) {
	fmt.Printf("%s: %s\n", url, duration)
}

// Define a function to run the test
func main() {
	// Run the test for each url
	for _, url := range urls {
		duration, err := runTest(url)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		printResults(url, duration)
	}
}

// End of file
