module github.com/anBertoli/go-workspace/app

go 1.19

require (
	// This is a submodule of the repository, listed both in this go.mod and as a
	// module in the go.work file. In a GOWORK=on environment the module is imported
	// from the local directory/module, in a GOWORK=off environment the version v0.0.12
	// is searched in the corresponding repo of the remote VCS.
	github.com/anBertoli/go-workspace/my_box v0.0.12

	// Note that the module github.com/anBertoli/go-workspace/my_unpub is not listed
	// here, but is used anyway in source code of the app. This is possible because
	// we are in a workspace that includes the my_unpub module. Anyway it's a bad
	// practice.
	//
	// github.com/anBertoli/go-workspace/my_unpub <some-pseudo-version>

	// These are both submodules of this repository, listed here, but not in go.work.
	// These are fetched from the corresponding repos in the remote VCS. Tags for this
	// module are in the form 'path/to/module/vX.Y.Z' (e.g. 'my_math/v0.0.2') since
	// submodules tagging is done with a path prefix.
	//
	// Note: if we need to modify these dependiencies, it is very useful to temporarily
	// add them in the go.work file, make the necessary changes, test the expected new
	// behaviour (from this app), tag the dependency with a new version & push, remove
	// it from go.work and bump the version requirements in the dependant (this go.mod).
	github.com/anBertoli/go-workspace/my_math v0.0.2
	github.com/anBertoli/go-workspace/my_strings v0.0.1

	// External dependency, searched in the remote repo in a VCS as usual.
	golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
)
