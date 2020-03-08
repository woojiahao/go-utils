package main

import (
  "fmt"
  "github.com/woojiahao/go-utils/src/slice"
)

func main() {
  data := slice.ConcurrentSlice{Data: make([]interface{}, 0)}
  data.Add("Value")
  data.Add("Foo")
  data.Add("Bar")
  fmt.Println(data)
  fmt.Println(data.Contains("Value"))
  fmt.Println(data.Contains("Something"))
  data.Remove("Value")
  fmt.Println(data)
}
