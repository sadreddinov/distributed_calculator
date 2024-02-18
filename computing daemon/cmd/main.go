package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sadreddinov/distributed_calculator/computing_daemon/pkg/models"
	"github.com/sadreddinov/distributed_calculator/computing_daemon/pkg/task"
	"github.com/sadreddinov/distributed_calculator/computing_daemon/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup

func main() {
	if err := initconfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	agent := utils.Register()
	go func(models.Agent) {
		for {
			select {
			case <-time.Tick(5 * time.Minute):
				utils.Ping(agent)
			}
		}
	}(agent)

	wg.Add(task.Num_of_goroutines)
	for i := 0; i < task.Num_of_goroutines; i++ {
		go Calculate()
	}
	for i := 0; i < task.Num_of_goroutines; i++ {
		task.WorkersDoneChan <- i
	}
	go task.GetTasks()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	go func(models.Agent) {
		<-quit
		utils.Shutdown(agent)
		for i := 0; i < task.Num_of_goroutines; i++ {
			wg.Done()
		}
	}(agent)
	wg.Wait()
}

func initconfig() error {
	viper.AddConfigPath("..\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func Calculate() {
	for {
		select {
		case calc := <-task.InputChan:
			result := 0.0
			switch calc.Operation {
			case "+":
				result = calc.Num1 + calc.Num2
			case "*":
				result = calc.Num1 * calc.Num2
			case "/":
				result = calc.Num1 / calc.Num2
			case "-":
				result = calc.Num1 - calc.Num2
			}
			dur, _ := time.ParseDuration(calc.OperationTime + "s")
			time.Sleep(dur)

			task.Out.Chans[calc.TaskNum] <- result
		}
	}
}
