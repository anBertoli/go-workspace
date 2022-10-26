# Go workspaces and multi-module repos

This document contains some general information about workspaces and multi-module repos. Browse
the files in the (commented) repository to look at how things works in practice. It is highly
recommended to clone the repo and experiment with it.

Please note that the workspace feature and multi-module repositories are distinct and independent
concepts and they **do not** necessarily have to be coupled, although they might work well together.
For a more general and throughout discussion about Go modules have a look at the references at the 
end of this document. 

There are also offical docs and tutorials about workspaces. I tried to both condense and enrich them
in the present document.
- https://go.dev/blog/get-familiar-with-workspaces (general discussion)
- https://go.dev/doc/tutorial/workspaces (simple tutorial)
- https://go.dev/ref/mod#workspaces (reference)

## Go workspaces

Go 1.18 adds workspace mode to Go, which lets you work on multiple modules simultaneously. With
workspaces different modules in the same workspace can refer to the local version of each other.

Previously, to add a feature to one module (dependency) and use it in another module (dependent), 
you needed to either publish the changes to the dependency, or edit the `go.mod` file of the dependent 
module with a `REPLACE` directive for your local, unpublished dependency changes. In order to publish 
the main module (the dependent) without errors, then you had to remove the `REPLACE` directive from 
its `go.mod` file after you published the local changes in the first module (the dependency). With 
Go workspaces you control all your dependencies using a `go.work` file in the root of your workspace 
directory and modules in the workspace can refer to the local versions of each other (even if 
unpublished). Versioning and publishing of single modules can be done as usual, but modules changes 
and integrations can be easily tested before publishing new module versions.

Workspaces are flexible and support a variety of workflows (https://go.dev/blog/get-familiar-with-workspaces#workflows).
Workspaces can be used to manage multiple modules in different ways, for example:
- _permanently_: the modules in the workspace are permanently added to the workspace with a `USE` directive 
  in the `go.work` file, so they can refer and use the local version of each other. Modules can still be 
  versioned and published independently,
- _temporarily_: similarly to the first option, but when you are done making changes and testing integrations, 
  you remove `USE` directives from `go.work`. In this way, your modules will always refer to published versions.

About the usage of your modules outside the workspace: external modules and users continue to use your modules 
referring to and fetching specific (tagged) versions of them.

### Mechanism

In workspace mode, given a module B (dependency) used by a module A (dependent), both included in the 
workspace, is searched by module A in the local workspace. If module B is not present, it is searched 
using to the standard Go rules (in a repo in a remote VCS). Even if the module B is not listed in the 
`go.mod` file of module A, module A can use module B if it is present in the workspace.

Note that for module A to work in a non-workspace mode its `go.mod` file must list the module B requirement
as usual and the module B must be downloadable from a remote VCS as usual.

If a `go.work` file is found in the working directory or a parent directory, or one is specified using
the `GOWORK` environment variable, the `go` command is in workspace mode. To determine which `go.work` 
file is being used, run `go env GOWORK`. The output is empty if the go command is not in workspace mode.
You can manually disable the workspace mode setting `GOWORK=off`.

### Go work command

- `go work init`: initialize a workspace in the current directory, creates the go.work file.
- `go work use [moddir]`: add modules to the workspace (can be done manually editing the go.work file).
- `go work edit`
- `go work sync`

```shell
$ go work init
$ go work use ./app
$ go work use ./my_box
$ go work use ./my_unpub
$ cat go.work
```

```go.mod
go 1.19

use (
	// The 'use' directive tells Go that the modules in the listed
	// directories are modules managed by this workspace and should
	// be considered when searching for modules in workspace mode.
	./app
	./my_box
	./my_unpub
)
```

## Multi-module repositories

Even if not usual, a repository can contain multiple Go modules. Each module can be separately
versioned and published (see chapter below for tags conventions). 

In a non-workspace context, each module can refer to other **local** modules only via `REPLACE` 
directives. Alternatively the old versioning mechanism must be used: a module refers to nearby 
modules via published versions in a remote VCS. With this second option, to develop and test 
changes in one dependency it must be pushed to the VCS at each change, and the requirements of
the dependent must be updated each time. Using `REPLACE` is easier but still awkward. 

When using workspaces the workflow can be simplified since we can directly import from local modules.
In this way we can make changes, test everything locally and tag & publish the dependencies only 
when the changes work fine.

### Versioning submodules

Tags for submodules are in the form `path/to/module/vX.Y.Z` (e.g. `my_math/v0.0.2`) since
submodules tagging is done with a path prefix. Each submodule version can be referred and 
fetched individually from other modules using its dedicated tags. Note however that if a 
module is listed in the `go.work` file, dependents from the same workspace will continue 
to use a local version (but you can remove the `USE` directive at any time).

Here's an example in a workspace context, where `my_math` is not included in the workspace
while `my_box` is:

```shell
# generate new tag for module my_math
$ cd ./my_math
$ git tag my_math/v0.0.2
$ git push --tags

# in app module, use my_box and newly published version of my_math
$ cd ../app
$ go get github.com/anBertoli/go-workspace/my_box@latest
$ go get github.com/anBertoli/go-workspace/my_math@v0.0.2
$ cat go.mod
```

Here's the resulting `go.mod`.

```go.mod
module github.com/anBertoli/go-workspace/app

go 1.19

require (
    // Added this specific version to the requirements of the go.mod. 
    // The my_math module is not listed in go.work and it is always 
    // fetched from the remote repo in the VCS (as usually happens
    // in the go deps system).
    github.com/anBertoli/go-workspace/my_math v0.0.2
    
    // This module is listed as a usual dependency in go.mod and also 
    // in the go.work file. In a GOWORK context it will be imported from
    // the local directory. In a GOWORK=off context it is fetched from 
    // the remote repo (as usual).
    github.com/anBertoli/go-workspace/my_box v0.0.12
    
    // ...
)
```

Here's an example using the new version of the submodule from another repo:

```shell
# in another project/repo use my_math
$ cd /path/to/another/project
$ go get github.com/anBertoli/go-workspace/my_math@v0.0.2
$ go get github.com/anBertoli/go-workspace/my_box@latest
$ cat go.mod
```

```go.mod
module github.com/anBertoli/another-repo

go 1.19

require (
    // Both fetched from the other remote repo in the VCS (as usually happens
    // in the go deps system).
    github.com/anBertoli/go-workspace/my_math v0.0.2
    github.com/anBertoli/go-workspace/my_box v0.0.12
      
    // ...
)
```

### List of tags
- `my_box/v0.0.10`
- `my_box/v0.0.11`
- `my_box/v0.0.12`
- `my_box/v0.0.4`
- `my_math/v0.0.1`
- `my_math/v0.0.2`
- `my_strings/v0.0.1`
- `v0.0.2`
- `v0.0.25`
- `v0.0.26`
- `v0.0.3`
- `v0.0.4`

## Docker builds in Go workspaces

When building Docker images for our modules in a workspace we should keep attention if we want to 
use local modules or published versions. In the first case the Docker build must consider not 
only the module itself but also the workspace environment, so the build must include the `go.work`
file and local modules we want to use. The Dockerfile in the app module follows this mode.

```shell
# docker build from upper dir
$ cd /path/to/go-workspace/repo
$ docker build -f ./app/Dockerfile . 
```

To build an image independently we must use a dockerfile that doesn't refer to local modules,
where all dependencies are specified in the `go.mod` file, the specified versions are all 
downloadable from the indicated remote VCS location and all items used from imported modules 
are in published versions. In our case we should add the `my_unpub` module in the `go.mod` file
(and also publish some changes of the `my_box` module, the Zero() method is not present in a 
published version yet). Note that it's a bad practice not to include a dependency in the `go.mod`
file, since the dependent cannot be built outside a non-workspace environment. 

The following steps are only showed here and were not performed on the project.

```shell
# Add a pseudo version requirement of the unpublished module.
$ cd ./app
$ go get github.com/anBertoli/go-workspace/my_unpub@latest
$ cat go.mod
```

```go.mod
module github.com/anBertoli/another-repo
go 1.19

require (
    github.com/anBertoli/go-workspace/my_math v0.0.2
    github.com/anBertoli/go-workspace/my_box v0.0.12
    github.com/anBertoli/go-workspace/my_unpub v0.0.0-20221025151004-9ef1df64f7ae     
    // ...
)
```

Then we can add a dedicated dockerfile used to build the app using only the published 
modules, as showed below (currently not present in the repo).

```dockerfile
FROM golang:1.19-alpine AS BUILDER
WORKDIR /app
COPY ./ ./
RUN go build -o ./app_bin

FROM alpine as RUNNER
WORKDIR /app
COPY --from=BUILDER /app/app_bin ./app_bin
EXPOSE 8080
CMD ["./app_bin"]
```

## References
- Offical docs and tutorials about workspaces (some are basic):
  - https://go.dev/blog/get-familiar-with-workspaces (general discussion)
  - https://go.dev/doc/tutorial/workspaces (simple tutorial)
  - https://go.dev/ref/mod#workspaces (reference)
- More general and throughout discussion about Go modules: 
  - https://go.dev/doc/modules/managing-dependencies 
  - https://go.dev/doc/#developing-modules
- Versioning multiple modules in a repo:
  - https://go.dev/doc/modules/managing-source#multiple-module-source
  - https://github.com/golang/go/wiki/Modules#publishing-a-release
