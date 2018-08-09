package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func AllCRMEndPoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "not implemented yet !")
}

func FindCRMEndpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "not implemented yet !")
}

func CreateCRMEndPoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "not implemented yet !")
}

func UpdateCRMEndPoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "not implemented yet !")
}

func DeleteCRMEndPoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "not implemented yet !")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/CRM", AllCRMEndPoint).Methods("GET")
    r.HandleFunc("/CRM", CreateCRMEndPoint).Methods("POST")
    r.HandleFunc("/CRM", UpdateCRMEndPoint).Methods("PUT")
    r.HandleFunc("/CRM", DeleteCRMEndPoint).Methods("DELETE")
    r.HandleFunc("/CRM/{id}", FindCRMEndpoint).Methods("GET")
    if err := http.ListenAndServe(":3999", r); err != nil {
        log.Fatal(err)
    }
}
