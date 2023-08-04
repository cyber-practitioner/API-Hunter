package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	// Replace this with the URL of the website you want to check
	websiteURL := "https://chat.openai.com/c/98c8750b-fe5c-45f5-a02b-544d661ff85a"

	// Regular expression to find potential API endpoints in JavaScript code
	apiRegex := regexp.MustCompile(`(?i)fetch\s*\(\s*['"]([^'"]+)['"]\s*\)|\$.ajax\s*\(\s*{\s*url:\s*['"]([^'"]+)['"]`)

	// Create a new headless Chrome browser context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Intercept network events
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			// Analyze JavaScript files for API calls
			if strings.HasSuffix(e.Request.URL, ".js") {
				if resp, err := http.Get(e.Request.URL); err == nil {
					defer resp.Body.Close()
					if jsCode, err := parseJavaScriptCode(resp.Body, apiRegex); err == nil {
						fmt.Printf("API calls in %s:\n", e.Request.URL)
						for _, match := range apiRegex.FindAllStringSubmatch(jsCode, -1) {
							for _, api := range match[1:] {
								fmt.Println(api)
							}
						}
					}
				}
			}
		}
	})

	// Navigate to the website
	err := chromedp.Run(ctx,
		chromedp.Navigate(websiteURL),
	)
	if err != nil {
		fmt.Println("Error navigating the website:", err)
		return
	}
}

func parseJavaScriptCode(body io.Reader, apiRegex *regexp.Regexp) (string, error) {
	jsCode, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	// Remove comments from JavaScript code
	jsCode = regexp.MustCompile(`(?s)/\*.*?\*/`).ReplaceAll(jsCode, []byte{})
	jsCode = regexp.MustCompile(`//.*`).ReplaceAll(jsCode, []byte{})

	return string(jsCode), nil
}
