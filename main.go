package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
    "strconv"
)

type message struct {
    Text string `json:"text"`
}

type inputNums struct {
    Var1 int `json:"a"`
    Var2 int `json:"b"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input inputNums
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
		// If there is an error in decoding, return an error response
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}


    msgMap := map[string]string{"result": strconv.Itoa(input.Var1 + input.Var2)}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msgMap)
}


func squareHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    num := r.URL.Query().Get("num")
    if num == "" {
		// If there is an error in decoding, return an error response
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
    }
	intNum, err := strconv.Atoi(num)
	if err != nil {
		// If there is an error in decoding, return an error response
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
        return
	}

    msgMap := map[string]string{"result": strconv.Itoa(intNum * intNum)}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msgMap)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    msg := r.URL.Query().Get("msg")
    if msg == "" {
		// If there is an error in decoding, return an error response
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
    }
    msgMap := map[string]string{"echo": msg}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msgMap)
}

func statusOkHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    msg := map[string]string{"status": "ok"}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    msg := map[string]string{"msg": "welcome varun"}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    msg := message{Text: "Hello, API!"}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    msg := map[string]string{"time" : time.Now().Format(time.RFC3339)}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}

func main() {
    http.HandleFunc("/welcome", welcomeHandler)
    http.HandleFunc("/status", statusOkHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/time", timeHandler)
    http.HandleFunc("/echo", echoHandler)
    http.HandleFunc("/square", squareHandler)
    http.HandleFunc("/add", addHandler)
    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}
