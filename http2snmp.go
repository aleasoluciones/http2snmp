// Copyright 2014 The http2snmp Authors. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aleasoluciones/gosnmpquerier"
)

const (
	CONTENTION = 4
)

func rootHandler(querier gosnmpquerier.SyncQuerier, w http.ResponseWriter, r *http.Request) {
	cmd, _ := gosnmpquerier.ConvertCommand(r.FormValue("cmd"))
	query := gosnmpquerier.Query{
		Cmd:         cmd,
		Community:   r.FormValue("community"),
		Oids:        []string{r.FormValue("oid")},
		Timeout:     time.Duration(10) * time.Second,
		Retries:     1,
		Destination: r.FormValue("destination"),
	}
	processed := querier.ExecuteQuery(query)
	jsonProcessed, err := gosnmpquerier.ToJson(&processed)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, jsonProcessed)

}

func main() {

	querier := gosnmpquerier.NewSyncQuerier(CONTENTION, 3, 3*time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootHandler(querier, w, r)
	})
	address := "0.0.0.0:8080"
	log.Println("Server running in", address, " ...")
	log.Fatal(http.ListenAndServe(address, nil))
}
