# PLS
The PLS (Pretty LS) tool was made in order to give help in finding and analyzing the file structure for large files.
In order to implement it, I was inspired by the `ls`, `df`, `du`, and `tree`.

## Installation

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

## Usage

In order to use it properly run the script from any directory on your machine using the following syntax
```sh
pls <OPTIONS> <FILES_OR_DIRS>
```
and it will give you a tree-like structure of files ordered based on the file sizes.

Possible options: 
- `-v` will show the current version of the CLI tool
- `-l <VALUE>` will limit the displayed files to the largest `<VALUE>` in the given directory. By default shows all files.
- `-d <VALUE>` will limit the displayed files in depth to `<VALUE>` levels in the given directory. By default displays files only on the 1st level.
- `-s` will show human-readable file sizes of the listed files. By default it's disabled.
- `-t` will show the types of the listed files. By default it's disabled.
- `-sp` will show the percentage of the displayed file size according to its parent directory. By default it's disabled.
- `-fc` will show the count of the files inside the listed files. By default it's disabled.
- `-fp` will show the percentage of the count of files inside the listed files according to its parent directory. By default it's disabled.

## TODO

A TODO list to be implemented:
- [ ] Add goroutine pool in order to control the load
    - [ ] Add a flag to configure processed, by default to be the GOMAXPROC
    - [ ] Add `pool` package to handle the load
- [ ] Add `logger`
    - [ ] Implement a `log_level` flag to control it
    - [ ] Replace all `fmt.Print*` with `logger`
- [ ] Add presentable logic
    - [ ] Implement a `output` flag to control it. 
    - [ ] Implement `stdout` presentable logic 
    - [ ] Implement `json` presentable logic 
    - [ ] Implement `xml` presentable logic 
- [ ] Add tests