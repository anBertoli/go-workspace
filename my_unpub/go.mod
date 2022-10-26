// The module is listed in the go.work file. It has NO published
// versions. Inside the workspace the local copy is imported and can
// be used even if it is not present in the go.mod of the dependant
// (because my_unpub is present in go.work). Note however that using
// a local module without specifiyng it in the go.mod is a bad practice,
// since the dependent cannot be built outside a non-workspace env.
module github.com/anBertoli/go-workspace/my_unpub

go 1.19
