package main

import (
	"context"
	"test/machinerydemo/worker"
)

func main() {
	worker.SendHelloWorldTask(context.Background())
}
