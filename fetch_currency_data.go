package main
import (
"fmt"
"net/http"
"io/ioutil"
)

func main() {
url := "https://westinpay.com/currency/fiat_api?api_key=YOUR-API-KEY&base=USD&output=JSON"
resp, err := http.Get(url)
if err != nil {
panic(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
panic(err)
}
fmt.Println(string(body))
}
