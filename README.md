# Finding API Calls in JavaScript Files (Go)

This Go script is designed to find potential API calls made in JavaScript files of a website. It uses the `chromedp` package to navigate to the website and intercept network events, then fetches the content of JavaScript files and searches for API-related patterns.

## How It Works

1. The script starts by defining the URL of the website you want to analyze for API calls.

2. A regular expression `apiRegex` is used to identify potential API calls in the JavaScript code. The script looks for two common patterns of API calls: `fetch` and `$.ajax`.

3. A headless Chrome browser context is created using `chromedp.NewContext`. This context is used to control the headless browser.

4. The script intercepts network events, such as `network.EventRequestWillBeSent`, to analyze JavaScript files for potential API calls.

5. When a JavaScript file is requested, the script downloads its content using `http.Get` and removes comments to simplify parsing.

6. The downloaded JavaScript code is then searched for API calls using the regular expression `apiRegex`.

7. If any matches are found, the API endpoints are printed to the console.

## How to Use

1. Replace `https://example.com` with the URL of the website you want to analyze.

2. Make sure you have the necessary Go packages installed (`chromedp`, `net/http`, `regexp`, etc.).

3. Run the script using the Go compiler, and the API calls in the JavaScript files will be displayed in the console.

4. usage: <code>go run main.go https://exmaple.com/auth/</code>
   


## Disclaimer

Please note that this is a simplified version of a more complex task. The provided script may not handle all scenarios, such as minified or obfuscated JavaScript code. It's important to perform automated scanning or analysis responsibly, with proper authorization from website owners and adherence to ethical standards.

Always respect website owners' terms and ensure you have explicit permission for any automated assessments or security testing.
