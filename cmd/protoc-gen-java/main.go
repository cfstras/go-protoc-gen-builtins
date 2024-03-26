package main

import (
	"github.com/cfstras/go-protoc-gen-builtins/internal/runner"
	"github.com/cfstras/go-protoc-gen-builtins/internal/wasm"
)

func main() {
	runner.Run("protoc-gen-java", wasm.ProtocGenJava)
}
