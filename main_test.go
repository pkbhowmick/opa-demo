package main_test

import (
	"io/ioutil"
	"os"
	"testing"

	opa "github.com/pkbhowmick/opa-demo"
	"github.com/stretchr/testify/require"
)

type Options struct {
	JsonInput      []byte
	RegoQuery      string
	RegoFilePath   string
	ExpectedResult interface{}
}

func TestCheckPolicy(t *testing.T) {
	t.Parallel()

	opts := []Options{
		{
			JsonInput: []byte(`{"statements" : [
			{
				"qid" : "q1",
				"max_time" : 5.123
			},
			{
				"qid" : "q2",
				"max_time" : 9.999
			}
			]
			}`),
			RegoQuery:      "data.example.no_slow_queries",
			RegoFilePath:   "./example.rego",
			ExpectedResult: true,
		},
		{
			JsonInput: []byte(`{"statements" : [
			{
				"qid" : "q1",
				"max_time" : 5.123
			},
			{
				"qid" : "q2",
				"max_time" : 10.001
			}
			]
			}`),
			RegoQuery:      "data.example.no_slow_queries",
			RegoFilePath:   "./example.rego",
			ExpectedResult: false,
		},
		{
			JsonInput: []byte(`{"networks": [
			{
				"id" : "net1",
				"public": false
			},
			{
				"id": "net2",
				"public": true
			}]}`),
			RegoQuery:      "data.example.no_public_network",
			RegoFilePath:   "./example.rego",
			ExpectedResult: false,
		},
		{
			JsonInput: []byte(`{"networks": [
			{
      			"id" : "net1",
      			"public": false
    		},
    		{
      			"id": "net2",
      			"public": false
    		}]}`),
			RegoQuery:      "data.example.no_public_network",
			RegoFilePath:   "./example.rego",
			ExpectedResult: true,
		},
	}

	tmpFile := "test.json"

	for _, o := range opts {
		err := ioutil.WriteFile(tmpFile, o.JsonInput, 0644)
		require.NoError(t, err)

		rs, err := opa.CheckPolicy(o.RegoQuery, o.RegoFilePath, tmpFile)
		require.NoError(t, err)

		require.Equal(t, o.ExpectedResult, rs)
	}
	err := os.Remove(tmpFile)
	require.NoError(t, err)
}
