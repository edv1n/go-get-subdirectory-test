# go-get-subdirectory-test

This repo is for testing Go's 1.25 new `subdirectory` field in `go-import` meta tags.
It use GitHub Pages to serve the meta tags and as the module path.

## Go Packages / Modules

All Go modules are under the `gopkg` directory.

There are two modules:

- `edv1n.github.io/go-get-subdirectory-test`
- `edv1n.github.io/go-get-subdirectory-test/sub`

### Modules Repository

The repository for both modules is `github.com/edv1n/go-get-subdirectory-test`.

### Modules Location in Repository

- `edv1n.github.io/go-get-subdirectory-test` is located at `gopkg/` directory.
- `edv1n.github.io/go-get-subdirectory-test/sub` is located at `gopkg/sub/` directory.

### Mapping Module to Repository with go-import Meta Tags

The `go-import` meta tags are served by GitHub Pages at `https://edv1n.github.io/go-get-subdirectory-test`.

The `go-import` meta tags are as follows:

For `edv1n.github.io/go-get-subdirectory-test` module:

```html
<meta name="go-import" content="edv1n.github.io/go-get-subdirectory-test git https://github.com/edv1n/go-get-subdirectory-test gopkg">
```

For `edv1n.github.io/go-get-subdirectory-test/sub` module:

```html
<meta name="go-import" content="edv1n.github.io/go-get-subdirectory-test/sub git https://github.com/edv1n/go-get-subdirectory-test gopkg/sub">
```

The meta tags HTML files are located in `gh-pages` branch. The structure is formed according to GitHub Pages requirements.


## Mapping Versions to commits

As specified by [Go Modules Reference: Mapping versions to commitsÂ¶](https://go.dev/ref/mod#vcs-version):

> If a module is defined in a subdirectory within the repository, that is, the module subdirectory portion of the module path is not empty, then each tag name must be prefixed with the module subdirectory, followed by a slash. For example, the module golang.org/x/tools/gopls is defined in the gopls subdirectory of the repository with root path golang.org/x/tools. The version v0.4.0 of that module must have the tag named gopls/v0.4.0 in that repository.

### Tags for `edv1n.github.io/go-get-subdirectory-test` module

Since `edv1n.github.io/go-get-subdirectory-test` is in a subdirectory of this repository (`gopkg/`), its tags must be prefixed with `gopkg/`.

Thus, the tags for `edv1n.github.io/go-get-subdirectory-test` module is `gopkg/vX.Y.Z`. For example, `gopkg/v0.0.2`.

### Tags for `edv1n.github.io/go-get-subdirectory-test/sub` module

Since `edv1n.github.io/go-get-subdirectory-test/sub` is in a subdirectory of this repository (`gopkg/sub/`), its tags must be prefixed with `gopkg/sub/`.

Thus, the tags for `edv1n.github.io/go-get-subdirectory-test/sub` module is `gopkg/sub/vX.Y.Z`. For example, `gopkg/sub/v0.0.2`.

## Importing the Modules

You can import the modules in your Go code as follows:

### For `edv1n.github.io/go-get-subdirectory-test` module

`go get` command:

```shell
# For edv1n.github.io/go-get-subdirectory-test module with latest version
go get edv1n.github.io/go-get-subdirectory-test@latest
```

In your Go code:

```go
import "edv1n.github.io/go-get-subdirectory-test"
```

### For `edv1n.github.io/go-get-subdirectory-test/sub` module

`go get` command:

```shell
# For edv1n.github.io/go-get-subdirectory-test/sub module with latest version
go get edv1n.github.io/go-get-subdirectory-test/sub@latest
```

In your Go code:

```go
import "edv1n.github.io/go-get-subdirectory-test/sub"
```

### Debugging `go get` Issues

If you encounter issues while using `go get`, you can enable verbose output to help diagnose the problem:

For `edv1n.github.io/go-get-subdirectory-test` module:

```shell
GODEBUG=gocacheverify=1 go get -x -v edv1n.github.io/go-get-subdirectory-test@latest

# get https://proxy.golang.org/edv1n.github.io/@v/list
# get https://proxy.golang.org/edv1n.github.io/go-get-subdirectory-test/@v/list
# get https://proxy.golang.org/edv1n.github.io/go-get-subdirectory-test/@v/list: 200 OK (0.089s)
# get https://proxy.golang.org/edv1n.github.io/@v/list: 404 Not Found (1.789s)
go: added edv1n.github.io/go-get-subdirectory-test v0.0.2
``` 

For `edv1n.github.io/go-get-subdirectory-test/sub` module:

```bash
GODEBUG=gocacheverify=1 go get -x -v edv1n.github.io/go-get-subdirectory-test/sub@latest
```

### Checking `go-import` Meta Tags

You can check the `go-import` meta tags served by GitHub Pages using the following commands:

For `edv1n.github.io/go-get-subdirectory-test` module:

```shell
$ curl -L https://edv1n.github.io/go-get-subdirectory-test?go-get=1

<head><meta name="go-import" content="edv1n.github.io/go-get-subdirectory-test git https://github.com/edv1n/go-get-subdirectory-test gopkg">

$ curl -L edv1n.github.io/go-get-subdirectory-test/sub\?=go-get=1

<head><meta name="go-import" content="edv1n.github.io/go-get-subdirectory-test/sub git https://github.com/edv1n/go-get-subdirectory-test gopkg/sub">
```

The output should include the meta tag mentioned above.
