# Stellar Go 
[![Build Status](https://travis-ci.org/stellar/go.svg?branch=master)](https://travis-ci.org/stellar/go) 
[![GoDoc](https://godoc.org/github.com/stellar/go?status.svg)](https://godoc.org/github.com/stellar/go)

This repo is the home for all of the public go code produced by SDF.  In addition to various tools and services, this repository is the SDK from which you may develop your own applications that integrate with the stellar network.

## Dependencies

This repository depends upon a [number of external dependencies](./glide.yaml), and we use [Glide](https://glide.sh/) to manage them.  Glide is used to populate the [vendor directory](http://glide.readthedocs.io/en/latest/vendor/), ensuring that builds are reproducible even as upstream dependencies are changed. Please see the [Glide](http://glide.sh/) website for installation instructions.

When creating this project, we had to decide whether or not we committed our external dependencies to the repo.  We decided that we would not, by default, do so.  This lets us avoid the diff churn associated with updating dependencies while allowing an acceptable path to get reproducible builds.  To do so, simply install glide and run `glide install` in your checkout of the code.  We realize this is a judgement call; Please feel free to open an issue if you would like to make a case that we change this policy.


## Directory Layout

In addition to the other top-level packages, there are a few special directories that contain specific types of packages:

* **clients** contains packages that provide client packages to the various Stellar services.
* **exp** contains experimental packages.  Use at your own risk.
* **handlers** contains packages that provide pluggable implementors of `http.Handler` that make it easier to incorporate portions of the Stellar protocol into your own http server. 
* **support** contains packages that are not intended for consumption outside of Stellar's other packages.  Packages that provide common infrastructure for use in our services and tools should go here, such as `db` or `log`. 
* **support/scripts** contains single-file go programs and bash scripts used to support the development of this repo. 
* **services** contains packages that compile to applications that are long-running processes (such as API servers).
* **tools** contains packages that compile to command line applications.

Each of these directories have their own README file that explain further the nature of their contents.

### Other packages

In addition to the packages described above, this repository contains various packages related to working with the Stellar network from a go program.  It's recommended that you use [godoc](https://godoc.org/github.com/stellar/go#pkg-subdirectories) to browse the documentation for each.


## Package source layout

While much of the code in individual packages is organized based upon different developers' personal preferences, many of the packages follow a simple convention for organizing the declarations inside of a package that aim to aid in your ability to find code.

Every package should have a `main.go` file.  This file contains the package documentation (unless a separate `doc.go` file is used), _all_ of the exported vars, consts, types and funcs for the package.  It may also contain any unexported declarations that are not tied to any particular type.  In addition to `main.go`, a package often has a single go source file for each type that has method declarations.  This file uses the snake case form of the type name (for example `loggly_hook.go` should contain methods for the type `LogglyHook`).

Each non-test file can have a test counterpart like normal, whose name ends with `_test.go`.  Additionally, an `assertions_test.go` file should contain any custom assertions that are related to the package in some way.  This allows a developer to include the package in their tests to gain access to assertions that make writing tests that involve the package more simple.  Finally, a `helpers_test.go` file can contain test utility functions that are not necessarily custom assertions.

Generally, file contents are sorted by exported/unexported, then declaration type  (ordered as consts, vars, types, then funcs), then finally alphabetically.


## Coding conventions

- Always document exported package elements: vars, consts, funcs, types, etc.
- Tests are better than no tests.