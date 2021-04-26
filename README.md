# go-server-embed-demo

This small snippet demonstrates how to serve embed files in a simple server.  Files from the `html/` subdirectory are served as is.  The `html/hello.html` file is also served as a template.  This "README" file is also served to demonstrate embedding a single file's contents.

## Background

Embedding data in Golang bundles resource files and directories in the compiled binary.  This  makes it easy to move Golang binaries without accidentally splitting them from necessary runtime resources.  Deployments are easier and sharing binaries with others is more simple.

## Programming notes

* Embedded data must be assigned to variables using the var directive.  Embedded variables are populated when the executable starts.
* To embed a file or file system, a `//go:embed <path/file>` compiler directive comment must appear before the declaration of each embedded variable.
* Embedded files and directories must be referenced relative to directory the source code is in without leading or trailing path separators.  The files are not allowed be in a higher directory or be referenced with a `..` path.  Soft (symbolic) links (shortcuts) are not allowed to reference the files.  Hard links are allowed.
* Directories (folders) must be embedded as `embed.FS`.  This structure implements the `io/fs.FS` interface, which other standard libraries accept.
* Files are be embedded as strings or byte arrays.  If no directories are imported in a source code file, the embed package must be imported with the blank identifier to ensure the compiler directive comment works correctly.  This is done with `import _ "embed"`.

## Dependencies

* [Go](https://golang.org/dl/) 1.16+ - used to by the source code
* [Make](https://www.gnu.org/software/make/) - used to build and run the demo server

## Build/Run

* to run the tests, run: `make build/main.test`
* to build the server, run: `make build/main`
* to run the server, run: `make serve`

The server runs on port 8000.  To stop it, press ctrl+C.

When the server is running, files in the html/ folder are at <http://localhost:8000/html/>.  The dynamic HTTP template is located at <http://localhost:8000/hello>.  To evaluate it with a name, add a `name` query parameter, like <http://localhost:8000/hello?name=Joe>.  This documentation is at <http://localhost:8000/about>

