package main


import (
 "bufio"
 "flag"
 "fmt"
 "net/http"
 "os"
 "strings"
 "sync"
)


// Color codes
const (
 ColorGreen  = "\033[32m"
 ColorYellow = "\033[33m"
 ColorRed    = "\033[31m"
 ColorGray   = "\033[90m"
 ColorReset  = "\033[0m"
)


func main() {
 url := flag.String("u", "", "Target URL")
 wordlist := flag.String("w", "", "Path to wordlist")
 threads := flag.Int("t", 10, "Number of threads")
 flag.Parse()


 if *url == "" || *wordlist == "" {
 fmt.Println("Usage: Rfuzz -u <url> -w <wordlist> -t <threads>")
 os.Exit(1)
 }


 queue := make(chan string)
 var wg sync.WaitGroup


 for i := 0; i < *threads; i++ {
 wg.Add(1)
 go func() {
 defer wg.Done()
 for word := range queue {
 targetURL := *url + "/" + word
 resp, err := http.Get(targetURL)
 if err != nil {
 fmt.Printf("%sError: %s%s\n", ColorRed, err, ColorReset)
 continue
 }
 defer resp.Body.Close()


 switch resp.StatusCode {
 case http.StatusOK:
 fmt.Printf("%s[200 OK] %s%s\n", ColorGreen, targetURL, ColorReset)
 case http.StatusMovedPermanently:
 fmt.Printf("%s[301 Redirect] %s%s\n", ColorYellow, targetURL, ColorReset)
 case http.StatusForbidden:
 fmt.Printf("%s[403 Forbidden] %s%s\n", ColorRed, targetURL, ColorReset)
 case http.StatusNotFound:
 fmt.Printf("%s[404 Not Found] %s%s\n", ColorGray, targetURL, ColorReset)
 default:
 fmt.Printf("[%d] %s\n", resp.StatusCode, targetURL)
 }
 }
 }()
 }


 file, err := os.Open(*wordlist)
 if err != nil {
 fmt.Printf("%sError opening wordlist: %s%s\n", ColorRed, err, ColorReset)
 os.Exit(1)
 }
 defer file.Close()


 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
 queue <- strings.TrimSpace(scanner.Text())
 }
 close(queue)


 wg.Wait()
}
