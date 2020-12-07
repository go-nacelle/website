package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
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

	return invokeParallel(paths, bustCacheForFile)
}

func gatherPaths() (paths []string, _ error) {
	err := filepath.Walk(publicDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, strings.TrimPrefix(strings.TrimSuffix(path, "/index.html"), publicDir))
		}

		return err
	})

	return paths, err
}

func bustCacheForFile(path string) error {
	u, err := url.Parse(baseURL + "/" + path)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-No-Cache", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return nil
}

func invokeParallel(values []string, f func(arg string) error) error {
	ch := make(chan string, len(values))
	for _, value := range values {
		ch <- value
	}
	close(ch)

	n := runtime.GOMAXPROCS(0)
	errs := make(chan error, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for value := range ch {
				if err := f(value); err != nil {
					errs <- err
				}
			}
		}()
	}

	wg.Wait()
	close(errs)
	return <-errs
}
