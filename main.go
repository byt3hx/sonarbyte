package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bufio"
	"flag"
	"os"
	"sync"
)

func scan(domain string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://sonar.omnisint.io/subdomains/%s", domain))
	if err != nil {
		fmt.Println(nil)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err !=nil {
		fmt.Println(err)
	}

	var values []string
	
	if err := json.Unmarshal([]byte(body), &values); err != nil {

	panic(err) }
	f  := make([]string, 0)
	for _, item := range values {
		f  = append(f,item)
	}
	return f,nil

}

func unique(strSlice []string) []string {

	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	var domains []string
	if flag.NArg() > 0 {
		domains = []string{flag.Arg(0)}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			domains = append(domains, sc.Text())
		}
		if err := sc.Err(); err !=nil {
			fmt.Fprintf(os.Stderr, "failed to read input: %s\n" , err)
		}
	}

	results := make(chan string)

	var wg sync.WaitGroup 
	for _, domain := range domains {
		wg.Add(1)
		go func (domain string) {
			sub , err := scan(domain)
			if err != nil {
				fmt.Println(err)
			}
			sort := unique(sub)
			for _, finaldomain := range sort {
				fmt.Println(finaldomain)
			}
			defer wg.Done()
		}(domain)
	}
	wg.Wait()
	close(results)
}