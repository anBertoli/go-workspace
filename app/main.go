package main

import (
	"fmt"

	// They are both submodules of this repository, listed here, but not in go.work.
	// These are fetched from the corresponding repos in the remote VCS. Tags for this
	// module are in the form 'path/to/module/vX.Y.Z' (e.g. 'my_math/v0.0.2') since
	// submodules tagging is done with a path prefix.
	"github.com/anBertoli/go-workspace/my_math"
	"github.com/anBertoli/go-workspace/my_strings"

	// This is a submodule of the repository, listed both in this go.mod and as a
	// module in the go.work file. In a GOWORK=off environment the version v0.0.12
	// is searched in the corresponding repo of the remote VCS. In a GOWORK=on env
	// the module is imported from the local directory/module (so box.Zero() works
	// even if not included in the v0.0.12 tag).
	"github.com/anBertoli/go-workspace/my_box"

	// The my_unpub module is not included in the REQUIRE directives of the go.mod
	// file, but listed in the go.work file. Using items from my_unpub works only in
	// a GOWORK=on environment (the go command knows it is part of the workspace).
	// In a GOWORK=off environment the build will fail since this source file
	// (app/main.go) tries to import a module not listed in the go.mod.
	"github.com/anBertoli/go-workspace/my_unpub"

	// External dependency, searched in the remote repo in a VCS as usual.
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(my_math.Add(2, 3))
	fmt.Println(my_math.Mul(2, 3))

	str := my_strings.NewStr("Ehy")
	str.ToUpper()
	fmt.Println(str)

	box := my_box.NewBox(45)
	fmt.Println(box)
	// box.Zero() is a new method not included in my_box/v0.0.12. It works
	// because the workspace includes the local my_box module, where
	// the box.Zero() method is present.
	box.Zero()
	fmt.Println(box)

	// my_unpub.HelloUnpublished() is not included in a REQUIREd module. It
	// works because the module is in the workspace.
	my_unpub.HelloUnpublished()
}
