package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	// Register handlers
	r := mux.NewRouter()
	s := r.Headers("Accept", "application/json").Subrouter()

	s.Methods("GET").Path("/product/{id:[0-9]+}").
		HandlerFunc(GetProduct)

	s.Methods("GET").Path("/debug").
		HandlerFunc(LogDebug)

	s.Methods("GET").Path("/warning").
		HandlerFunc(LogWarning)

	s.Methods("GET").Path("/error").
		HandlerFunc(LogError)

	s.Methods("GET").Path("/fatal").
		HandlerFunc(LogFatal)

	s.Methods("GET").Path("/").
		HandlerFunc(IndexHandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))

	log.Info(http.ListenAndServe(":8080", nil))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]
	if len(productId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("id not included: %v\n", productId)
	} else {
		p := `{"productId":` + productId + `}`
		log.Infof("productId: %v", productId)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, p)
	}
}

func LogDebug(w http.ResponseWriter, r *http.Request) {
	log.Debug("Log /debug URL invoked. Test debug log")
	c := `{"DebugPage": "This is a debug"}`
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, c)
}

func LogWarning(w http.ResponseWriter, r *http.Request) {
	log.Warn("Log /warning URL invoked. Test warning log")
	c := `{"WarningPage": "This is a warning"}`
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, c)
}

func LogError(w http.ResponseWriter, r *http.Request) {
	log.Error("Log /error URL invoked. Test error log")
	c := `{"ErrorPage": "This is an error. Plus a minor change."}`
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, c)
}

func LogFatal(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Log /fatal URL invoked. Test fatal log")
	c := `{"FatalPage": "This is a fatal"}`
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, c)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Inside IndexHandler")
	c := `{"HomePage": "hello"}`
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, c)
}
