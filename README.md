## Golang
This is a repository of miscellaneous apps I am developing using the Go programming language (otherwise known as "Golang"). This README.md file also serves as my personal Golang cheat-sheet.  
![gopher.png](/gopher.png)  

**How to Install Go on a Debian Linux-based System**
```bash
sudo apt update
sudo apt install golang -y
```

**If/Else Statement**
```go
package main // The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.
import "fmt"
func main() {
    if 1 + 1 == 2 {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
```

## Tooling
```bash
go list ... # list all installed Go packages

go mod init demo
go mod tidy

go get github.com/mattn/go-sqlite3 # add Go package to Go module
```

## Imports
```go
import (
    _ "github.com/mattn/go-sqlite3"
)
    /* 
        To import a package solely for its side-effects (initialization), use the blank identifier as explicit package name.
        In the case of "go-sqlite3," the side-effect is being able to register the sqlite3 driver as a database driver within the init() function without importing any other functions.
    */
```

**Importing sqlite3**  
![mingw-64](/mingw-64.png)

## Functions
```go
import (
    "crypto/md5"
    "encoding/hex"
)

func HashPassword(text string) string {
    hash := md5.Sum([]byte(text))
    return hex.EncodeToString(hash[:])
}
```

## Resources
Go for Windows  
[https://go.dev/dl/](https://go.dev/dl/)

gVim (a text-editor) for Windows  
[https://ftp.nluug.nl/pub/vim/pc/gvim82.exe](https://ftp.nluug.nl/pub/vim/pc/gvim82.exe)

Git for Windows  
[https://github.com/git-for-windows/git/releases/download/v2.34.1.windows.1/Git-2.34.1-64-bit.exe](https://github.com/git-for-windows/git/releases/download/v2.34.1.windows.1/Git-2.34.1-64-bit.exe)

MySQL for Windows  
[https://dev.mysql.com/downloads/installer/](https://dev.mysql.com/downloads/installer/)
