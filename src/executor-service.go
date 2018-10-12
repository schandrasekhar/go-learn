package main

import (
    "log"
)

type ExecutorService interface {
    execute()
    submit() chan
}

type Runnable interface {
    run()
}

type CustomJobExecutor struct {
    
}

type MyJob struct {
    value string
}

func (executor CustomJobExecutor) execute(r Runnable) {
    r.run()
}

func (job MyJob) run() {
    log.Println(job.value)
}


func main() {
    var executor CustomJobExecutor = CustomJobExecutor{}
    var myJob MyJob = MyJob{value: "testing"}
    executor.execute(myJob)
}