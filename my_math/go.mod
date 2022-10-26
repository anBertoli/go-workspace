// The module is NOT listed in the go.work file. It has some
// published versions (e.g. my_math/v0.0.2). Inside the workspace
// the local copy is NOT imported (because my_math is NOT present
// in go.work).
module github.com/anBertoli/go-workspace/my_math

go 1.19
