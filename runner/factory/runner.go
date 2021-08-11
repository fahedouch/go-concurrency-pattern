package factory

import (
    "errors"
    "os"
    "os/signal"
    "time"
)

type Runner struct {
    interrupt chan os.Signal
    complete  chan error
    timeout   <-chan time.Time
    tasks     []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interupt")

func (r *Runner) Add(f ...func(int)) {
     r.tasks = append (r.tasks, f...)
}

func New(d time.Duration) *Runner{
    return &Runner{
                     interrupt: make (chan os.Signal, 1),
                     complete: make(chan error),
                     timeout:  time.After(d),
       }
}
func (r *Runner) Start() error {
            signal.Notify(r.interrupt, os.Interrupt)
            go func() {
                 r.complete <- r.run()
            }()

       select{
          case err := <- r.complete:
                return err
          case <- r.timeout:
                return errors.New("I am in timeout")
       }

}

func (r *Runner) run() error {
    for id, task := range r.tasks {
        if r.getInterrupt() {
            return ErrInterrupt
        }
        task(id)
    }
    return nil
}

func (r *Runner) getInterrupt() bool {
     select {
        case <-r.interrupt:
          return true
        default:
          return false
     }
}