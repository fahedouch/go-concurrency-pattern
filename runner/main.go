package main

import (
    "github.com/fahedouch/go-concurrency-pattern/runner/factory"
    "time"
    "fmt"
)
func main()  {
  r := factory.New(1 * time.Nanosecond)
  r.Add(createTask(), createTask(), createTask(), createTask(), createTask())
  err := r.Start()
  if err != nil {
   fmt.Println("here we have error")
  }
}

func createTask() func(int){
    return func(id int){
        fmt.Println("kaka")
        fmt.Println(id)
    }

}