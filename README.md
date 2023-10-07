## pyingo
pyingo is a Golang library that allows you to use Python-like syntax in your Golang code.
It provides a set of data types and built-in functions commonly found in Python, making it easier for developers familiar with Python to work with Golang.

### Features
- **Python-like syntax**: With pyingo, you can write Golang code using Python-like syntax, making it more intuitive for Python developers to transition to Golang.
- **List data type**: pyingo includes a list data type that behaves similarly to Python lists. You can perform operations like appending, slicing, and iterating over elements.
- **Tuple data type**: Similar to Python tuples, pyingo offers a tuple data type that allows you to group multiple values together.
- **Dictionary data type**: pyingo provides a dictionary data type, allowing you to store key-value pairs and perform operations such as accessing, updating, and deleting elements.
- **Built-in functions**: pyingo includes a range of built-in functions commonly used in Python, such as print(), open(), max(), min(), etc. These functions work seamlessly with the Python-like syntax.

### Installation
To use pyingo in your Golang project,
1. Get it into your project:
```bash
go get github.com/dev-bittu/pyingo
```
2. Simply import it as a package:
```go
import (
    py "github.com/dev-bittu/pyingo"
)
```

### Usage
Here's an example of how you can use pyingo in your Golang code:
```go
package main

import (
    py "github.com/dev-bittu/pyingo"
)

func main() {
    py.Println("Hello pyingo")
}
```

In the above example,
We import pyingo and run println (similar to print) function to execute

### Contributing
Contributions to pyingo are welcome!
If you find any issues or have suggestions for improvements,
please feel free to open an issue or submit a pull request.

### License
pyingo is released under the MIT License.
See the [LICENSE](LICENSE) file for more details.

That's just a basic example,
but you can customize it further based on your project's specific features and requirements.
