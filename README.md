# pls
The PLS tool (Pretty LS) was made in order give help in finding and anlyzing the file structure for large files.
In order to implement it, I was inspired by the `ls`, `df`, `du` and `tree`.

# Installation

1. Make sure to install [Go](https://go.dev/doc/install).
2. Check your `GOBIN` env and if not set, configure it as 
    ```sh
    go env -w GOBIN="$GOPATH/bin"
    ```
3. Install the CLI tool using  
    ```sh
    go install github.com/PuloV/pls@latest
    ```
4. Test the installation using 
    ```sh
    pls -v
    ```

# Usage

In order to use it properly run the script from any directory on your machine using the following syntax
```sh
pls <OPTIONS> <FILES_OR_DIRS>
```

