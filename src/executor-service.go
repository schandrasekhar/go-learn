package main

import (
    "log"
)

type ExecutorService interface {
    execute()
    submit() chan string
}

type Runnable interface {
    run()
}

type Callable interface {
    call()
    getChan() chan string
}

type CustomJobExecutor struct {
    
}

type MyJob struct {
    c chan string
    value string
}

func (executor CustomJobExecutor) execute(r Runnable) {
    go r.run()
}

func (executor CustomJobExecutor) submit(c Callable) (chan string) {
    go c.call()
    return c.getChan()
}

func (job MyJob) run() {
    log.Println(job.value)
}

func (job MyJob) call() {
    job.c <- job.value
}

func (job MyJob) getChan() (chan string){
    return job.c
}


func main() {
    var executor CustomJobExecutor = CustomJobExecutor{}
    var myJob MyJob = MyJob{value: "testing", c: make(chan string)}
    executor.execute(myJob)
    c := executor.submit(myJob)
    result := <- c
    log.Println(">>>>>>")
    log.Println(result)
    //infinite loop
    for {}
}