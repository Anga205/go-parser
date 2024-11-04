# AFLL Miniproject, Go Parser

This Go project contains a parser that checks the syntax of Go code for specific elements such as braces, `if` statements, `while` statements, and `for` statements. The parser reads code from a file named `sample.txt` and validates the syntax based on predefined rules.

## Files

- `parser.go`: The main Go file that contains the logic for parsing and validating the Go code.
- `sample.txt`: A sample Go code file that is used as input for the parser.

## Functions

### `CheckBraces(code string) bool`
Checks if the braces `{}` in the code are balanced.

### `CheckIfStatements(code string) bool`
Checks if the `if` statements in the code are correctly formatted.

### `CheckWhileStatements(code string) bool`
Checks if the `while` statements in the code are correctly formatted.

### `CheckForStatements(code string) bool`
Checks if the `for` statements in the code are correctly formatted.

## Usage

1. Place your Go code in the `sample.txt` file.
2. Run the `parser.go` file using the command:
    ```sh
    go run parser.go
    ```
3. The parser will read the code from `sample.txt` and validate it. The results will be printed in the console.

## Example

### `sample.txt`
```go
package main

import "fmt"

func main() {
     if true {
          fmt.Println("Hello, World!")
}
     for i := 0; i < 10; i++ {
          fmt.Println(i)
     }
     while true {
          log.Println()
     }
}
```

### Output
```
Braces do not match, exiting code.
'If' statements are correct.
'For' statements are correct.
'While' statements are incorrect, exiting code.
Finished checks in Xms
--------------------------------
Code is invalid.
```
