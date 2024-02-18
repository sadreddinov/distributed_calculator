package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sadreddinov/distributed_calculator/computing_daemon/pkg/models"
	"github.com/spf13/viper"
)

func Ping(agent models.Agent) {
	newBuf := new(bytes.Buffer)
	//если горутины занятя, то work_state = in_progress иначе free
	json.NewEncoder(newBuf).Encode(agent)

	ping_url := viper.GetString("orchestrator.register_url")

	ping_req, _ := http.NewRequest("PATCH", ping_url, newBuf)
	client := &http.Client{}
	_, err := client.Do(ping_req)
	if err != nil {
		log.Print(err.Error())
	}
}

func Register() models.Agent {
	id := os.Getenv("UUID")
	uuid, _ := uuid.Parse(id)

	agent := &models.Agent{Id: uuid, Work_state: "in_progress"}
	newBuf := new(bytes.Buffer)
	json.NewEncoder(newBuf).Encode(agent)

	register_url := viper.GetString("orchestrator.register_url")

	register_req, _ := http.NewRequest("POST", register_url, newBuf)
	client := &http.Client{}
	resp, err := client.Do(register_req)
	if err != nil {
		log.Print(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		time.Sleep(5 * time.Second)
		Register()
	}
	return *agent
}

func Shutdown(agent models.Agent) {
	newBuf := new(bytes.Buffer)
	agent.Work_state = "lost_connection"
	json.NewEncoder(newBuf).Encode(agent)

	shutdown_url := viper.GetString("orchestrator.shutdown_url")

	Shutdown_req, _ := http.NewRequest("PATCH", shutdown_url, newBuf)
	client := &http.Client{}
	_, err := client.Do(Shutdown_req)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Shutdown")
}
