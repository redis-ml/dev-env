package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spyzhov/ajson"
)

func main() {
	json := []byte(`...`)

	root, _ := ajson.Unmarshal(json)
	nodes, _ := root.JSONPath("$..price")
	for _, node := range nodes {
		node.SetNumeric(node.MustNumeric() * 1.25)
		node.Parent().AppendObject("currency", ajson.StringNode("", "EUR"))
	}
	result, _ := ajson.Marshal(root)

	fmt.Printf("%s", result)

	fmt.Println("vim-go")
	glog.Info("Prepare to repel boarders")
}
