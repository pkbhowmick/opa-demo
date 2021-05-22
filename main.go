package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"io/ioutil"
	"log"
)

func main() {
	ctx := context.Background()

	r := rego.New(
		rego.Query("x = data.example"),
		rego.Load([]string{"./demo.rego"}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatalln(err)
	}


	bs, err := ioutil.ReadFile("./input.json")
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

	fmt.Println(rs[0].Bindings["x"].(map[string]interface{})["hello"])
}
