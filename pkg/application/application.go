package application

import (
	"calculator/pkg/calculation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var id int = 0
var expressions = make([]expression, 0, 1024)

type expression struct {
	id         int
	status     string
	result     float64
	expression string
	task       string
}

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) Run() error {
	// for {
	// 	log.Println("Input expression:")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	text, err := reader.ReadString('\n')
	// 	text = text[:len(text)-2]
	// 	if err != nil {
	// 		fmt.Println("Failed to read application from console")
	// 	}
	// 	result, err := calculation.Calc(text)
	// 	if text == "exit" {
	// 		log.Println(text, " calculation failed with error: ", err)
	// 	} else {
	// 		log.Println(text, "=", result)
	// 	}
	// }
	text := "(1-(2+3))+2-1"
	result, _ := calculation.Calc(text)
	log.Println(text, "=", result)
	return nil
}

type Request struct {
	Expression string `json:"expression"`
	ID         int    `json:"id"`
	Task       string `json:"task"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	fmt.Fprintln(w, "{")
	if err != nil {
		expr := expression{id: id, status: err.Error(), result: result, expression: request.Expression}
		expressions = append(expressions, expr)
		fmt.Fprintf(w, "    \"error\": \"%s\"", err.Error())
		id += 1
	} else {
		expr := expression{id: id, status: "OK", result: result}
		expressions = append(expressions, expr)
		fmt.Fprintf(w, "    \"id\": %v", id)
		id += 1
	}
	fmt.Fprintln(w, "\n{")
}

func ExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "{")
	fmt.Fprintf(w, "    \"expressions\": [\n")
	for i := 0; i < len(expressions); i++ {
		fmt.Fprintf(w, "        {\n")
		fmt.Fprintf(w, "            \"id\": %v\n", expressions[i].id)
		fmt.Fprintf(w, "            \"status\": %v\n", expressions[i].status)
		fmt.Fprintf(w, "            \"result\": %v\n", expressions[i].result)
		fmt.Fprintf(w, "        },\n")
	}
	fmt.Fprintf(w, "    ]\n")
	fmt.Fprintln(w, "}")
}

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "{")
	fmt.Fprintf(w, "    \"expression\":\n")
	fmt.Fprintf(w, "        {\n")
	for i := 0; i < len(expressions); i++ {
		if request.ID == expressions[i].id {
			fmt.Fprintf(w, "            \"id\": %v\n", request.ID)
			fmt.Fprintf(w, "            \"status\": %v\n", expressions[i].status)
			fmt.Fprintf(w, "            \"result\": %v\n", expressions[i].result)
			break
		} else if i == len(expressions)-1 {
			fmt.Fprintf(w, "            \"id\": %v\n", request.ID)
			fmt.Fprintf(w, "            \"status\": \"Expression not found\"\n")
			fmt.Fprintf(w, "            \"result\": \"Result not found\"\n")
			w.WriteHeader(http.StatusNotFound)
		}
	}
	if len(expressions) == 0 {
		fmt.Fprintf(w, "            \"id\": %v\n", request.ID)
		fmt.Fprintf(w, "            \"status\": \"Expression not found\"\n")
		fmt.Fprintf(w, "            \"result\": \"Result not found\"\n")
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "        }\n")
	fmt.Fprintln(w, "}")
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	http.HandleFunc("/api/v1/expressions", ExpressionsHandler)
	http.HandleFunc("/api/v1/expressions/:id", ExpressionHandler)
	http.HandleFunc("/internal/task", TaskHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
