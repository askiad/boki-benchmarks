package main

import (
	"flag"
)

var FLAGS_faas_gateway string
var FLAGS_fn_prefix string
var FLAGS_concurrency int

func init() {
	flag.StringVar(&FLAGS_faas_gateway, "faas_gateway", "127.0.0.1:8081", "")
	flag.StringVar(&FLAGS_fn_prefix, "fn_prefix", "", "")
	flag.IntVar(&FLAGS_concurrency, "concurrency", 1, "")
}

func main() {
	flag.Parse()
}
