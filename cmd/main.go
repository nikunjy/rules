package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/conekta/Conekta-Golang-Rules-Engine/parser"
)

func main() {
	jsonFile, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(fmt.Errorf("error reading obj from file %v", err))
	}
	rulesBytes, err := ioutil.ReadFile("rules.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("error reading rule from file %v", err))
	}
	var rulesString = string(rulesBytes)
	var info map[string]interface{}
	json.Unmarshal([]byte(jsonFile), &info)
	ev, err := parser.NewEvaluator(rulesString)
	if err != nil {
		log.Fatal(fmt.Errorf("error making evaluator from the rule %v, %v", rulesString, err))
	}
	ans, err := ev.Process(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(`
		Rule: %v
		Object: %v
		Answer: %v
	`, rulesString, info, ans,
	)
	if err := ev.LastDebugErr(); err != nil {
		fmt.Println("Last debug error", ev.LastDebugErr())
	}
}
