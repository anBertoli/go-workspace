### Module `my_unpub`

The module is listed in the `go.work` file. It has NO published versions. Inside the workspace
the local copy is imported and can be used even if it is not present in the `go.mod` of the
dependant (because _my_unpub_ is present in `go.work`). Note however that using a local module
without specifying it in the go.mod is a bad practice, since the dependent cannot be built/ran
outside a non-workspace env.