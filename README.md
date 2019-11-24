Quick source management tool

# Usage

- `qs`: show status of current folder
- `qs d`: download sources (git pull)
- `qs u -m [message]`: upload sources (git add + commit + push)
- `qs r`: Increment version

# Read version from go file
Create file `version.go` in src folder with body
````go
package main

var Major = 1
var Minor = 0
var Patch = 0
````
Now you can get version from `go` file