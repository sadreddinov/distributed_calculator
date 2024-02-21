package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/overseven/go-math-expression-parser/parser"
	"github.com/sadreddinov/distributed_calculator/computing_daemon/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var working_goroutines = 0
var EnvNum = ""
var Num_of_goroutines = 0
var WorkersDoneChan = make(chan int, 5)
var InputChan = make(chan models.Calculation, Num_of_goroutines)
var Out *OutputChan

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	EnvNum = os.Getenv("NumOfGoroutine")
	Num_of_goroutines, _ = strconv.Atoi(EnvNum)
	Out = NewOutputChan()
	
}



type OutputChan struct {
	Chans []chan float64
}

func GetTask(num int) bool {
	id := os.Getenv("UUID")
	task_url := viper.GetString("orchestrator.get_task_url")+"/" + id

	register_req, _ := http.NewRequest("GET", task_url, nil)
	client := &http.Client{}
	resp, err := client.Do(register_req)
	if err != nil {
		log.Print(err.Error())
		time.Sleep(5*time.Second)
		GetTask(num)
		return true
	}
	var expression models.Expression
	err = json.NewDecoder(resp.Body).Decode(&expression)
	if _, ok := resp.Header["Plus"]; !ok {
		time.Sleep(5*time.Second)
		GetTask(num)
		return ok
	}
	var operations models.Operation
	operations.Plus = resp.Header["Plus"][0]
	operations.Divide = resp.Header["Divide"][0]
	operations.Minus = resp.Header["Minus"][0]
	operations.Multiply = resp.Header["Multiply"][0]
	go Solve(expression, operations, num)
	return true
}

func Solve(expression models.Expression, operations models.Operation, num int) {
	working_goroutines++
	parser := parser.NewParser()
	expr, _ := parser.Parse(expression.Expr)
	strExpr := expr.String()
	strExpr = strings.ReplaceAll(strExpr, "( ", "")
	strExpr = strings.ReplaceAll(strExpr, " )", "")
	fmt.Println(strExpr)
	result, err := parsePrefixExpression(strExpr, operations, num)
	resString := strconv.FormatFloat(result, 'f', 5, 64)
	expression.Result = resString
	expression.Work_state = "solved"
	fmt.Println(result)
	PostResult(expression, err)
	WorkersDoneChan <- num
	working_goroutines--
}

func GetTasks() {
	for {
		for working_goroutines < Num_of_goroutines {
			i := <- WorkersDoneChan 
			GetTask(i)
		}
	}
}

func parsePrefixExpression(expression string, operations models.Operation, num int) (float64, error) {
	fmt.Println(num)
	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, fmt.Errorf("empty expression")
	}
	stack := make([]float64, 0)
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("expression parsing error")
			}
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64
			switch token {
			case "+":
				fmt.Println(operand1, "+", operand2)
				calc := &models.Calculation{Operation: "+", Num1: operand1, Num2: operand2, OperationTime: operations.Plus, TaskNum: num}
				InputChan <- *calc
				result = <-Out.Chans[num]
			case "-":
				fmt.Println(operand1, "-", operand2)
				calc := &models.Calculation{Operation: "-", Num1: operand1, Num2: operand2, OperationTime: operations.Minus, TaskNum: num}
				InputChan <- *calc
				result = <-Out.Chans[num]
			case "*":
				fmt.Println(operand1, "*", operand2)
				calc := &models.Calculation{Operation: "*", Num1: operand1, Num2: operand2, OperationTime: operations.Multiply, TaskNum: num}
				InputChan <- *calc
				result = <-Out.Chans[num]
			case "/":
				if operand2 == 0 {
					return 0, fmt.Errorf("cannot divide by zero")
				}
				fmt.Println(operand1, "/", operand2)
				calc := &models.Calculation{Operation: "/", Num1: operand1, Num2: operand2, OperationTime: operations.Divide, TaskNum: num}
				InputChan <- *calc
				result = <-Out.Chans[num]
			}
			stack = append(stack, result)
		default:
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("expression parsing error", token)
			}
			stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("некорректное выражение")
	}

	return stack[0], nil
}

func NewOutputChan() *OutputChan {
	chans := make([]chan float64, 0)
	for i := 0; i < Num_of_goroutines; i++ {
		ch := make(chan float64)
		chans = append(chans, ch)
	}
	return &OutputChan{Chans: chans}
}

func PostResult(result models.Expression, err error) {
	fmt.Println(err.Error())
	if err != nil {
		newBuf := new(bytes.Buffer)
		result.Result = err.Error()
		json.NewEncoder(newBuf).Encode(result)

		post_result_url := viper.GetString("orchestrator.post_result_url")

		post_result_req, _ := http.NewRequest("POST", post_result_url, newBuf)
		client := &http.Client{}
		_, ReqErr := client.Do(post_result_req)
		if ReqErr != nil {
			time.Sleep(5 * time.Second)
			PostResult(result, err)
		}
	}
	newBuf := new(bytes.Buffer)
	json.NewEncoder(newBuf).Encode(result)

	post_result_url := viper.GetString("orchestrator.post_result_url")

	post_result_req, _ := http.NewRequest("POST", post_result_url, newBuf)
	client := &http.Client{}
	_, ReqErr := client.Do(post_result_req)
	if ReqErr != nil {
		time.Sleep(5 * time.Second)
		PostResult(result, err)
	}
}
