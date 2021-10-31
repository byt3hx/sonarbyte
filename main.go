package main

import(
	"fmt"
	"flag"
	"sync"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"bufio"
	"time"
)

func main() {
	var domainfile string
	var domain string
	flag.StringVar(&domain , "d" , "" , "Set the domain name")
	flag.StringVar(&domainfile, "df" , "" , "Set the file")
	flag.Parse()

	if domain != "" {
		sub , err := scan(domain)
		if err != nil {
			fmt.Println(err)
		}
		sort := unique(sub)
		for _, finaldomain := range sort {
			fmt.Println(finaldomain)
		}
	}

	if domainfile != "" {
		file, err := os.Open(domainfile)
		if err != nil {
			fmt.Println("File could not be read!")
		}
		defer file.Close()
		time.Sleep(time.Millisecond * 10)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			var domains []string
			domains = append(domains, scanner.Text())
			var wg sync.WaitGroup
			for i := 0 ; i <= 0; i++ {
				wg.Add(1)
				go func() {
					for _, finaldomain := range domains {
						sorting(finaldomain)
					}
					wg.Done()
				}()
				wg.Wait()
			} 
		}
	}
	
}

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

func sorting(domain string) {
	sub , err := scan(domain)
	if err != nil {
		fmt.Println(err)
	}
	sort := unique(sub)
	for _, finaldomain := range sort {
		fmt.Println(finaldomain)
	}

}
