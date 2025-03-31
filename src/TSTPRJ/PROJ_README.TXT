
$ mkdir -p mypkg/nstpkg

$ touch main.go

// top level main

$ touch mypkg/mypkg.go

// nested pkg

$ touch mypkg/nstpkg/nstpkg.go

// init the module 

$ go mod init TSTPRJ
go: creating new go.mod: module TSTPRJ
go: to add module requirements and sums:
        go mod tidy

// above cmd create below go.mod file

$ more go.mod
module TSTPRJ

go 1.24.1

// then exec

$ go run main.go 
(public) mypkg -> Hello  world!
(public) nstpkg -> Hello  world!

