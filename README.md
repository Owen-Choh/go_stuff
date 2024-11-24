# Introduction
This repository is a collection of Go code I wrote while learning the language. It contains simple programs, experiments, and notes. 

It's mainly for my personal learning, but I am making the repo public in case it may be helpful to others. 

> [!NOTE]  
> Since this is just for personal learning / experimenting, the code probably won't adhere to best practices, be the most efficient or work in every edge case.


# Contents
`tutorial/` Basic Go tutorials

`sorting/` Sorting algorithms with tests and benchmarks

# Golang commands
To initialise the go.mod file (i.e. start writing golang)

`go mod init github.com/Owen-Choh/<module>`

<br>

To build an executable named after the module and run it

`go build; ./<module>`

<br>

To get some remote package, will also update the go mod file

`go get github.com/Owen-Choh/<module>`
- also updates add .sum to keep track of dependancies

<br>

Hack for local development, better not to do in production (put in the go.mod file)

`replace github.com/Owen-Choh/go_stuff/mystrings v0.0.0 => ../mystrings`

### packages
- "main" means you can build and run this file
- "something else" means this file is imported into something else
  - first letter of function name needs to be caps to export

<br>

### running tests and benchmark

`go test <optional test file path>`

some useful flags
- `-v` verbose message to see what test pass and fail, and logs even when test passes
- `-bench=.` or `-bench .` to run all benchmarks in this file
  - `-benchmem` if you want to see memory allocation
  - example `go test -bench . -benchmem ./` run all benchmarks found in this directory and show the memory allocations too

# Future plans
Stuff I plan to work on in the future:

- build a simple API and probably put it in a container
- experiment more on concurrency
- experiment with using a db
