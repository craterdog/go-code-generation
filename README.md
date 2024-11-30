<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Syntax Generation

### Overview
This project provides template-based code generation capabilities that allow
the generation of a Go module containing a `grammar` package defining a parser,
validator and formatter for a specified `Syntax.cdsn` file.  An `ast` package
containing a class for each node in the corresponding _abstract syntax tree_ is
also generated.

### Quick Links
For more information on this project click on the following links:
 * [project documentation](https://github.com/craterdog/go-code-generation/wiki)
 * [coding conventions](https://github.com/craterdog/go-class-model/wiki)
 * [command-line tools](https://github.com/craterdog/go-syntax-tools/wiki)

### Getting Started
To include the Go packages for this module use the following import statement:
```go
import (
	gen "github.com/craterdog/go-code-generation/v5"
)
```

### Contributing
Project contributors are always welcome. Check out the contributing guidelines
[here](https://github.com/craterdog/go-code-generation/blob/main/.github/CONTRIBUTING.md).

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
