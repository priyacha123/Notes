# GOLANG

### (Features similar to Java and C++)
- Strong typing (once a variable is initialized to be an integer then it can;t be changes later on)
- Statically typing (all the variables has to be defined at compile time)
- Go does have fetaures that allow user to get around its type system, but 99% NO.

-  Excellent community

### Key Features
- Simplicity
- Fast compile times
- Garbage colletor
- Built-in concurrency
- Compile to standalone binaries


<details close>
<h1 align="center">Hello in go-lang</h1>

```
package main
import "fmt"

func main() {
    fmt.Println("Hello, Priya")
}
```

</details>


# Variables
## Variable declaration
In 3 ways:
1. Simple reading formt  

```bash
# initialize variables
 var i int 

# assigning value
i = 42

# OR initialize + assigning value
var i int = 42
 ```

 ## GO Files Type
 1. Executable File
 - Every exc file has "main" pkg (inside main pkg, it has to have main fun)
 - Mostly used for backend

 2. Library/Package File
 - Used for importing in other files

## go.mod file
It has 3 main benefits:
1. Constains prpject name or code hosting provider name
2. Minimum golang version required
3. External dependencies

## Packages
- Every go file starts with pkg declaration.
- Folder/File name often match the pkg name
- Inside "main" pkg, it has to have main fun which is the entry point of the go file i.e., the "main" fun is where the execution of the file and calling of all other fun happens

### GO supports UTF-8 which supports 120000 characters from different languages unlike ASCII

## GO build command
- Go provide a fetaure "cross platform single binary executables" which means go can generate binary executable file which can be rum in any OS.

``` 
go build filename.go
```

- To run a file in window from linux, run this command:
```
GOOS=windows GOARCH=amd64 go build
```

- go build command compiles all the packages and generate binary executable file.
- It throws out all the result of the compilation and does not store anything  
- So, when again the build command is runned, it start from scratch and thus <strong> is slower tha "go run" command </strong>

## GO run command
``` 
go run filename.go
```
- go run command compiles all the packages and directly runs the program without generating binary file.

<strong> It is relatabily faster </strong> becuz it store the result in cache so that when run command is runned, the output comes faster.

## Imports
- When there are more than one imports, then imports are grouped into a parenthesized, "factored" import statement.

```
import (
	"fmt"
	"math"
)
```

- It can also be written as:

``` 
import "fmt"
import "math"
```

<strong> But it is good style to use the factored import statement. </strong>

## Name Convention
- Package name cannot be "Capital letter"
- Package name should be as small as possible (eg: math/rand -> for random no.)
- Package name should match the folder name
- Those datatypes whch start with capital letter are accessiblle from outside the packages (eg: math.Pi)

## Functions
```
func add(x int, y int) int { // the "int" after ")" is return's type
	return x + y
}

// OR 

func add(x, y int) int { // the "int" after ")" is return's type
	return x + y
}
```


