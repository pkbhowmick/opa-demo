package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/open-policy-agent/opa/rego"
)

var (
	regoQuery     string
	regoFilePath  string
	inputFilePath string
)

func CheckPolicy(regoQuery, regoFilePath, inputFilePath string) (interface{}, error) {
	ctx := context.Background()

	r := rego.New(
		rego.Query(fmt.Sprintf("x = %s", regoQuery)),
		rego.Load([]string{regoFilePath}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return nil, err
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		return nil, err
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return nil, err
	}

	//fmt.Println(rs)

	if len(rs) > 0 {
		return rs[0].Bindings["x"], nil
	}
	return false, nil
}

func main() {
	flag.StringVar(&regoQuery, "query", "data.example.hello", "Rego query for evaluation")
	flag.StringVar(&regoFilePath, "regofile", "./example.rego", "Rego file path")
	flag.StringVar(&inputFilePath, "input", "./input.json", "Input file path")
	flag.Parse()

	fmt.Println(regoQuery, regoFilePath, inputFilePath)

	rs, err := CheckPolicy(regoQuery, regoFilePath, inputFilePath)
	if err != nil {
		fmt.Println(err)
	}

	if rs.(bool) {
		fmt.Println("Policy is maintained")
	} else {
		fmt.Println("Alert!!! Policy is not maintained")
	}
}
