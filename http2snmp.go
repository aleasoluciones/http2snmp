// Copyright 2014 The http2snmp Authors. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aleasoluciones/gosnmpquerier"
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
	address := flag.String("address", "0.0.0.0", "listen address")
	port := flag.String("port", "8080", "listen port")
	contention := flag.Int("contention", 4, "max concurrent queries per destination")
	circuitBreakErrors := flag.Int("maxerrors", 4, "consecutive errors to mark destination as faulty")
	circuitResetSeconds := flag.Int("resettime", 30, "time reset faulty state for a destination (seconds)")
	flag.Parse()

	querier := gosnmpquerier.NewSyncQuerier(
		*contention,
		*circuitBreakErrors,
		time.Duration(*circuitResetSeconds)*time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootHandler(querier, w, r)
	})
	addressAndPort := fmt.Sprintf("%s:%s", *address, *port)
	log.Println("Server running in", addressAndPort, " ...")
	log.Fatal(http.ListenAndServe(addressAndPort, nil))
}
