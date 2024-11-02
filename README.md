# Go Projects

This is the repo to contain my learnings of [The Go Programming Language](https://go.dev/). They will mostly be from the [Go Docs](https://go.dev/doc/). I will try to put the notes of the topics I learn in this README.

## [Getting started)](https://go.dev/doc/tutorial/getting-started)

### Hello, world

Create a [hello.go](./01-getting-started-with-go/hello.go) file, and write the following code.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

This creates a `main` **package** inside the containing directory. Then inside the `main` **function**, it prints "Hello, world" to the STDOUT, using the `Println` **function** from the `fmt` **package** imported earlier.

To add dependency tracking, run the following commnad:

```sh
go mod init example/hello
```

This is still not fully clear to me. I might check their [docs on this](https://go.dev/doc/modules/managing-dependencies) at some point.

#### Running the code directly

To run the code, either of the following commands will work:

```sh
go run .
# or:
go run hello.go
# or
go run hello
```

because the `main` function is inside the hello.go file.

### Running by building an executable

To create an executable, first run _any one_ the following commands:

```sh
go build .
# or:
go build hello.go
# or:
go build hello
```

This will create an executable in the same directory as the source code. On Windows, it will be called "hello.exe".

Now run the executable:

```sh
./hello
# or:
./hello.exe
```

**Note:** I'm using the Zsh terminal for my development. If you're using the standard CMD prompt on Windows, the execution command might vary (most likely, it will be the same commands without the `./`, so either `hello` or `hello.exe`). On Powershell, the Zsh commands I've used above will work.

### Using external package

Inside the hello.go file, import the `rsc.quote` command. Then call `Go` from `quote`. I've done this in a separate function called `phrase`, and called the it in the `main` function.

```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func phrase() {
	fmt.Println(quote.Go())
}

func main() {
	fmt.Println("Hello, world")
	phrase()
}
```

I'm not entirly sure of the terminology just yet. Coming from the JS world, I would guess `quote` would be an "object", and `Go` invoked from `quote` would be a "method", but I would need to look into this more to confirm this.

Now to actually make sure the package works, install the `tidy` module like:

```sh
go mod tidy
```

And then run the code the same way shown [here](#running-the-code-directly) or [here](#running-by-building-an-executable).
