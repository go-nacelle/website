package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const baseURL = "https://nacelle.dev"
const publicDir = "public"

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
		os.Exit(1)
	}
}

func mainErr() error {
	paths, err := gatherPaths()
	if err != nil {
		return err
	}

	allPaths := map[string]struct{}{}
	for _, path := range paths {
		if strings.HasSuffix(path, "/index.html") {
			allPaths[strings.TrimSuffix(path, "index.html")] = struct{}{}
			allPaths[strings.TrimSuffix(path, "/index.html")] = struct{}{}
		}

		allPaths[path] = struct{}{}
	}

	hasDeadLinks := false
	for _, path := range paths {
		links, err := getLinks(path)
		if err != nil {
			return err
		}

		var deadLinks []string
		for _, link := range links {
			if _, ok := allPaths[link]; !ok {
				deadLinks = append(deadLinks, link)
			}
		}

		if len(deadLinks) > 0 {
			fmt.Printf("%s contains dead links to:\n", path)

			for _, link := range deadLinks {
				fmt.Printf("\t- %s\n", link)
			}

			fmt.Printf("\n")
			hasDeadLinks = true
		}
	}

	if hasDeadLinks {
		os.Exit(1)
	}

	return nil
}

func gatherPaths() (paths []string, _ error) {
	err := filepath.Walk(publicDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, strings.TrimPrefix(path, publicDir))
		}

		return err
	})

	return paths, err
}

var hrefPattern = regexp.MustCompile(fmt.Sprintf(`href="([^"]+)"`))

func getLinks(path string) ([]string, error) {
	contents, err := ioutil.ReadFile(filepath.Join(publicDir, path))
	if err != nil {
		return nil, err
	}

	var links []string
	for _, m := range hrefPattern.FindAllSubmatch(contents, -1) {
		url := string(m[1])
		switch {
		case strings.HasPrefix(url, baseURL):
			links = append(links, strings.TrimPrefix(url, baseURL))
		case !strings.HasPrefix(url, "http"):
			links = append(links, filepath.Join(path, url))
		default:
		}
	}

	return links, nil
}

// func invokeParallel(values []string, f func(arg string) error) error {
// 	ch := make(chan string, len(values))
// 	for _, value := range values {
// 		ch <- value
// 	}
// 	close(ch)

// 	n := runtime.GOMAXPROCS(0)
// 	errs := make(chan error, n)
// 	var wg sync.WaitGroup

// 	for i := 0; i < n; i++ {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()

// 			for value := range ch {
// 				if err := f(value); err != nil {
// 					errs <- err
// 				}
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	close(errs)
// 	return <-errs
// }
