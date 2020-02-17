package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := -1
	var err error

	pathParams := mux.Vars(r)
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for userID"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for commentID"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func startSending1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	timeInSeconds := -1
	var err error

	pathParams := mux.Vars(r)
	if val, ok := pathParams["timeInSeconds"]; ok {
		timeInSeconds, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "timeInSeconds must be a number"}`))
			return
		}
	}

	concurrentThreads := -1
	if val, ok := pathParams["concurrentThreads"]; ok {
		concurrentThreads, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for commentID"}`))
			return
		}
	}

	timeInSeconds = timeInSeconds + 1
	concurrentThreads = concurrentThreads + 1

	w.Write([]byte(fmt.Sprintf(`{"timeInSeconds": %d, "concurrentThreads": %d, " OK" }`, timeInSeconds, concurrentThreads)))
}

func startSending2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//timeInSeconds = 60
	//concurrentThreads := 1
	query := r.URL.Query()
	timeInSeconds, ok1 := strconv.Atoi(query.Get("timeInSeconds"))
	concurrentThreads, ok2 := strconv.Atoi(query.Get("concurrentThreads"))

	println(ok1)
	println(ok2)
	println(timeInSeconds)
	println(concurrentThreads)

	go startSendingRequests(timeInSeconds, concurrentThreads)

	w.Write([]byte(fmt.Sprintf(`{"timeInSeconds": %d, "concurrentThreads": %d, " OK" }`, timeInSeconds, concurrentThreads)))

}

func startSendingRequests(timeInSeconds int, concurrentThreads int) {
	println("Starting ...")
	busyWait()
	println("completed!")
}

func busyWait() {
	for a := 0; a <= 10000-1; a++ {
		for c := 0; c <= 1000000-1; c++ {
			b := 0
			b = b + 1
		}
	}
}

/*******************************
Clear is Better than Clever
Rob Pike
******************************** */
func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	api.HandleFunc("/start/{timeInSeconds}/{concurrentThreads}", startSending1).Methods(http.MethodGet)
	api.HandleFunc("/start", startSending2).Methods(http.MethodGet) // ?timeInSeconds=3&concurrentThreads=4

	log.Fatal(http.ListenAndServe(":8080", r))
}
