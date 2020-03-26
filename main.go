package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	schedulerv1 "k8s.io/kube-scheduler/extender/v1"
)

func filter(args schedulerv1.ExtenderArgs) *schedulerv1.ExtenderFilterResult {
	for _, node := range args.Nodes.Items {

	}
	return nil
}

func Filter(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	body := io.TeeReader(r.Body, &buf)

	var extenderArgs schedulerv1.ExtenderArgs
	var extenderFilterResult *schedulerv1.ExtenderFilterResult
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		extenderFilterResult = &schedulerv1.ExtenderFilterResult{
			Error: err.Error(),
		}
	} else {
		extenderFilterResult = filter(extenderArgs)
	}

	if response, err := json.Marshal(extenderFilterResult); err != nil {
		log.Fatalln(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/nodepool-filter", Filter)
	if err := http.ListenAndServe(":8090", r); err != nil {
		log.Fatal(err)
	}
}
