# GoPraf

## Description

GoPraf is a tool for formatting Go code that supports configuration through a configuration file, logging actions, and ignoring certain files and directories. This tool helps keep code clean and structured, improving readability and maintainability.

## Main Features

- **Code Formatting**: Automatic formatting of Go files.
- **Configuration**: Ability to configure ignored directories and files.
- **Logging**: All actions are logged for later analysis.
- **Configuration Editing**: User-friendly interface for changing settings.
- **Multilingual Support**: Supports English and Russian languages.

## Installation

1. Clone the repository:
   
```bash
   git clone https://github.com/yourusername/my-go-formatter.git
   cd my-go-formatter
```

2. Navigate to the project directory and install dependencies:
   
```bash
   go mod tidy
```

## Usage

### Formatting a File or Directory

```bash
go run cmd/formatter.go <filename_or_directory> [config.json]

- `<filename_or_directory>`: The path to the Go file or directory to format.
- `[config.json]`: (optional) The path to the configuration file.
```


### Editing Configuration

To edit the configuration file, use the command:

bash
go run cmd/formatter.go edit config.json

This will open an interface for changing ignored directories and files.

### Language Selection

You can specify the interface language by adding `ru` or `en` at the beginning of the command:

```bash
go run cmd/formatter.go ru <filename_or_directory> [config.json]
```

or

```bash
go run cmd/formatter.go en <filename_or_directory> [config.json]
```

## Notes

- Ignored directories and files can be specified in `config.json`.
- Formatted files will be created with `.formatted` added to the original file name.
- Logs will be saved in `logs/formatter.log`.

## Example Configuration File `config.json`

```json
{
    "ignore_dirs": ["vendor", "node_modules"],
    "ignore_files": ["*_test.go"]
}
```
### Example of disorganized and incorrect code (Before)

```go
package main
import "fmt"
func main(){x:=3.14; y:=x*x; if y>10{fmt.Println("x is large")}else{fmt.Println("x is small")}}
```

### ProblLack of Indentationack of Indentation**: The code is hard to read due to the absence of spaces Package Importing*Package Importing**: Packages are imported in one line, which does not conforInefficient Formattingficient Formatting**: All code elements are placed in one line, making it difficulNo Empty Lines Between Logical Blockseen Logical Blocks**: There is no separation between different sections of the code.

### Code after using the Formatter (After)

```go
package main

import "fmt"

func main() {
    x := 3.14
    y := x * x

    if y > 10 {
        fmt.Println("x is large")
    } else {
        fmt.Println("x is small")
    }
}
```

### Changes aImproved Formattingmproved Formatting**: The code is now well-formatted, making it more readable. Indentation and spaces are uCorrect Package Importing Package Importing**: Packages are imported in a single block, which is standardLogical Structure*Logical Structure**: The code is broken down into logical blocks, making it easieEmpty Lines.
4. **Empty Lines**: Empty lines have been added between logical parts of the code, enhancing readability.

### Conclusion

Using the code formatter significantly improves the structure and formatting of the code, making it more readable and maintainable, which is especially important when working on collaborative projects.
