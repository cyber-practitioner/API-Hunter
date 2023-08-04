// future updates to main.go feel free to contribute
package main

import (
	"errors"
	"sync"

	"github.com/projectdiscovery/gologger"
)

var errNoScripts = errors.New("no script tags found")

func main() {

	// Wrap in goroutine to allow concurrent fetching
	scripts, err := fetchScriptsAsync("https://nineatcollegepark2.residentportal.com/auth")
	if err != nil {
		gologger.Error().Msgf("Could not fetch scripts: %v\n", err)
		return
	}

	if len(scripts) == 0 {
		gologger.Error().Msg(errNoScripts)
		return
	}

	endpoints := extractEndpoints(scripts)

	// Print endpoints
	// etc...

}

// fetcher.go

func fetchScriptsAsync(url string) ([]string, error) {

	var scripts []string
	var wg sync.WaitGroup

	// Spin up 20 concurrent fetchers
	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			script, err := fetchScript(url)
			if err != nil {
				gologger.Error().Msgf("Fetch error: %v\n", err)
				return
			}

			scripts = append(scripts, script)
		}()
	}

	wg.Wait()

	if len(scripts) == 0 {
		return nil, errNoScripts
	}

	return scripts, nil

}

// extractor.go

func extractEndpoints(scripts []string) []string {

	// Wrap in goroutine to allow concurrent parsing
	var endpoints []string
	var wg sync.WaitGroup

	for _, script := range scripts {
		wg.Add(1)

		go func(content string) {
			defer wg.Done()

			eps := parseScript(content)
			endpoints = append(endpoints, eps...)

		}(script)
	}

	wg.Wait()

	return endpoints
}
