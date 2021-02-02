# A Filesystem for GO HTTP FileServer

This module when used with http.FileServer type will return the content of a file or index.html if the requested path is a directory, but it will not return a directory listing if index.html is not present.

In case of a request received by the filesystem to list a certain directory (which has no index.html) it will return a permission denied error which will be converted by http.FileServer to 403 (Forbidden) http error.

You can also define masks or files that can be queried on the filesystem and all other file requests will return permission denied.

## Example
```go
func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(&HttpFileSystem{
		RootDir:    "ui/static/",
		TrimPrefix: "/static",
		Masks:      nil,
	})

	mux.Handle("/static/", fileServer)

    s := &http.Server{
    	Addr: ":8080",
    	Handler: mux,
    }
    
    log.Println("starting server on address:", s.Addr)
    
    log.Fatal(s.ListenAndServe())
}
```