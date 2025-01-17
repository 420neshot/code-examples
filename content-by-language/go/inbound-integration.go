package main

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
)

type Payload struct {
  Account string `json:"account_number"`
}

func main() {
  data := Payload{
    Account: "ACC00000000000000000",
  }
  payloadBytes, err := json.Marshal(data)
  if err != nil {
    fmt.Println(err)
  }

  body := bytes.NewReader(payloadBytes)

  req, err := http.NewRequest("POST", "{VAULT_URL}/post", body)
  if err != nil {
    fmt.Println(err)
  }
  req.Header.Set("Content-Type", "application/json")

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    fmt.Println(err)
  }

  defer resp.Body.Close()

  respB, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(respB))
}
