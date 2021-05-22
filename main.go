package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"io/ioutil"
	"log"
)

var (
	regoQuery     string
	regoFilePath  string
	inputFilePath string
)

func main() {
	flag.StringVar(&regoQuery, "query", "data.example.hello", "Rego query for evaluation")
	flag.StringVar(&regoFilePath, "regofile", "./example.rego", "Rego file path")
	flag.StringVar(&inputFilePath, "input", "./input.json", "Input file path")
	flag.Parse()

	fmt.Println(regoQuery, regoFilePath, inputFilePath)

	ctx := context.Background()

	r := rego.New(
		rego.Query(fmt.Sprintf("x = %s", regoQuery)),
		rego.Load([]string{regoFilePath}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		log.Fatalln(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalln(err)
	}

	if rs[0].Bindings["x"].(bool) {
		fmt.Println("Policy is maintained")
	} else {
		fmt.Println("Alert!!! Policy is not maintained")
	}
}
