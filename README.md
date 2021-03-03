# SPM - Go
Software Package Metrics tool for Go

Supported metrics:
- Number of files
- Dependencies
- Instability
- Abstractness
- Distance from main sequence

# Description of supported metrics

## Number of files
The Go source files (go extension) in each package.

## Dependencies
Displays the dependencies of each package. SPM-Go classifies the dependencies in three groups:
- **Standard dependencies:** Dependencies contained in go standard library (e.g. _math_, _fmt_, _net/http_).
- **External dependencies:** Dependencies imported from external sources (e.g. _github.com/spf13/cobra_, _golang.org/x/tools_).
- **Internal dependencies:** Dependencies contained in the current module.

## Instability
> The ratio of efferent coupling (Ce) to total coupling (Ce + Ca) such that I = Ce / (Ce + Ca). 
> This metric is an indicator of the package's resilience to change.
> The range for this metric is 0 to 1, with I=0 indicating a completely stable package and I=1 indicating a 
> completely unstable package. (See: https://en.wikipedia.org/wiki/Software_package_metrics)

SPM-Go uses the following criteria for efferent and afferent coupling:
- **Efferent Coupling:** Counts only the internal dependencies (packages inside the module)
- **Afferent Coupling:** Packages inside the module  that depends on the current package.

## Abstractness
> The ratio of the number of abstract classes (and interfaces) in the analyzed package to the 
> total number of classes in the analyzed package. The range for this metric is 0 to 1, 
> with A=0 indicating a completely concrete package and A=1 indicating a completely abstract package. (See: https://en.wikipedia.org/wiki/Software_package_metrics)

As Go doesn't have classes, SPM-Go uses the following criteria for determining the required params to calculate the abstractness:
- **Abstractions:** Counts structs and interfaces definitions
- **Implementations:** Counts functions and methods

So the metric formula is: _Abstractness = Abstractions/Implementations_ 

## Distance from main sequence
> The perpendicular distance of a package from the idealized line A + I = 1. D is calculated as D = | A + I - 1 |. 
> This metric is an indicator of the package's balance between abstractness and stability. A package squarely on the 
> main sequence is optimally balanced with respect to its abstractness and stability. Ideal packages are either completely 
> abstract and stable (I=0, A=1) or completely concrete and unstable (I=1, A=0). The range for this metric is 0 to 1, 
> with D=0 indicating a package that is coincident with the main sequence and D=1 indicating a package that is as far 
> from the main sequence as possible. (See: https://en.wikipedia.org/wiki/Software_package_metrics)

# Usage
To install spm-go, run
```bash
$ go get github.com/fdaines/spm-go
```

To execute this tool you have to be in the module path
```bash
$ cd [path-to-your-module]
```

Now you can execute Spm-Go tool
```bash
$ spm-go command [flags]
```

## Supported commands

| Command      | Description                              |
| ------------ |:----------------------------------------:|
| packages     | Lists packages and number of files       |
| dependencies | Lists dependencies of each package       |
| instability  | Analyzes instability of packages         |
| abstractness | Analyzes abstractness of packages        |
| distance     | Analyzes distance from the main sequence |
| all          | Displays all metrics for each package    |

## Supported flags

| Flag      | Description                                                                                     |
| --------- |:-----------------------------------------------------------------------------------------------:|
| --format  | Specifies the output format for the command. Supported values are: _json_, _console_ and _csv_. |
| --verbose | Includes detailed information while the command is running                                      |


## Examples
```bash
$ spm-go packages
$ spm-go dependencies -v
$ spm-go instability --verbose
$ spm-go abstractness -f csv
$ spm-go distance --format json
$ spm-go all -v -f json
```

## Output Formats

### console (Default)
This format uses `github.com/jedib0t/go-pretty/v6/table` package to print out a formatted table into the console.
```
+----+-----------------+-------------+
|  # | PACKAGE         | FILES COUNT |
+----+-----------------+-------------+
|  1 | module          |           1 |
|  2 | module/foo      |           9 |
|  3 | module/foo/bar  |           2 |
|  4 | module/utils    |           2 |
+----+-----------------+-------------+
```

### json
```json
{
  "packages": [
    {
      "name": "main",
      "path": "module",
      "files": [
        "main.go"
      ],
      "files_count": 1,
      "dependencies": {
        "standard": [
          "errors",
          "fmt"
        ],
        "internals": [
          "module/utils"
        ],
        "externals": [],
        "standard_count": 2,
        "internals_count": 1,
        "externals_count": 0,
        "count": 3
      },
      "dependants": [
        "module"
      ],
      "afferent_coupling": 0,
      "efferent_coupling": 1,
      "instability": 1,
      "abstractness_details": {
        "functions": 0,
        "methods": 0,
        "interfaces": 0,
        "structs": 0
      },
      "abstractions_count": 0,
      "implementations_count": 0,
      "abstractness": 0,
      "distance": 0
    },
    {
      "name": "utils",
      "path": "module/utils",
      "files": [
        "temporal.go",
        "numeric.go",
        "other.go"
      ],
      "files_count": 3,
      "afferent_coupling": 1,
      "efferent_coupling": 0,
      "instability": 0,
      "abstractness_details": {
        "functions": 0,
        "methods": 0,
        "interfaces": 0,
        "structs": 0
      },
      "dependencies": {
        "standard": [
          "errors",
          "fmt",
          "math"
        ],
        "internals": [],
        "externals": [
          "golang.org/x/tools/go/packages"
        ],
        "standard_count": 3,
        "internals_count": 0, 
        "externals_count": 1,
        "count": 4
      },
      "dependants": [
        "module"
      ],
      "abstractions_count": 0,
      "implementations_count": 0,
      "abstractness": 0,
      "distance": 1
    }
  ]
}
```

### csv
This format uses a semicolon separated format
```
Package;Files
module;1
module/foo;2
module/foo/bar;3
module/utils;6
```

# Contributions
Feel free to contribute.

