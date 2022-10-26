// The module is NOT listed in the go.work file. It has some
// published versions (e.g. my_strings/v0.0.1). Inside the workspace
// the local copy is NOT imported (because my_strings is NOT present
// in go.work).
module github.com/anBertoli/go-workspace/my_strings

go 1.19
