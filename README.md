# Findall

Golang tool using fastwalk to recursively search through the working directory for an occurance of a string. File extension needs to be explicitly included

## Usage
```findall go fmt```
Will recursively find all .go files, search through each gofile for an occurance of "fmt" within the file. Extremely good for searching for variable occurances in a directory.

## Installation
To use this program, you can either download the binary in the bin directory, or build the file via
```go build findall.go```

## Speed

Leveraging fastwalk, this program is competitive with fd, and faster than powershell's ```ls -r```
