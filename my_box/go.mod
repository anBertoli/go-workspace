// The module is listed in the go.work file. It has also some published
// versions (e.g. my_box/v0.0.12). Inside the workspace (and GOWORK=on)
// the module is imported from the local directory/module, in a non-workspace
// environment (or GOWORK=off) the module is searched in the corresponding
// repo of the remote VCS.
module github.com/anBertoli/go-workspace/my_box

go 1.19
