package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Esse é o programa do exercicio de compilação cruzada. foi compilado num linus/amd64 e agora esta rodando em um sistema:", runtime.GOARCH, runtime.GOOS)
	time.Sleep(10 * time.Second)
}
