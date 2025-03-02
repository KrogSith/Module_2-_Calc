package application

import (
	"calculator/pkg/calculation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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
		fmt.Fprintf(w, "    \"error\": \"%s\"", err.Error())
	} else {
		fmt.Fprintf(w, "    \"result\": %f", result)
	}
	fmt.Fprintln(w, "\n{")
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
