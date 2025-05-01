package intermediate

import (
	"fmt"
	"net/url"
)

func URLParsingIntro() {
	fmt.Println("this is url parsing")
	basicParsing()
	buildUrl()
}

func basicParsing() {
	rawURL := "https://username:password@www.example.com:8080/path/to/resource?query1=val1&query2=val2#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Hostname:", parsedURL.Hostname())
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	fmt.Println("User:", parsedURL.User)

	// Extract username and password
	if parsedURL.User != nil {
		username := parsedURL.User.Username()
		password, _ := parsedURL.User.Password()
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)
	}
}

func buildUrl() {

	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/search",
		Fragment: "section1",
		User:     url.UserPassword("shivam", "secret"),
	}

	// Add query params
	params := url.Values{}
	params.Add("q", "golang")
	params.Add("lang", "en")
	u.RawQuery = params.Encode()

	fmt.Println("Built URL:", u.String())

}
