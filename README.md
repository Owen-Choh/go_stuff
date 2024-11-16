To initialise the go.mod file

`go mod init github.com/Owen-Choh/<module>`

To build an executable named after the module and run it

`go build; ./<module>`
- package main means can build and run this file
- package something else means this file is imported into something else
  - first letter of function name needs to be caps

To get some remote package, will also update the go mod file

`go get github.com/Owen-Choh/<module>`
- also updates add .sum to keep track of dependancies

Hack for local development, dont do in production
`replace github.com/Owen-Choh/mystrings v0.0.0 => ../mystrings`