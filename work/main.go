package main

import (
    "github.com/fahedouch/work/work"
    "fmt"
    "sync"
    "time"
)

var books =  []string{"book1", "book2", "book3", "book4", "book5"}

type Printer struct{
    id int
}

func (p *Printer) Task() {
    fmt.Println("printer %s is now printing", p.id)
    time.Sleep(time.Second)
}

func main() {
    //create new Pool
    pool := work.New(5)

    var wg sync.WaitGroup
    wg.Add(5)
    for _,i := range books {
         fmt.Println("Preparing book %s for printing", i)
         machine := Printer{
             id: 1,
         }
        go func(){
            pool.Run(&machine)
            wg.Done()
        }()
    }

    wg.Wait()

    pool.Shutdown()
}