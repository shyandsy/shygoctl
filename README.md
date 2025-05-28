# shyandsy/shygoctl

This is a customized goctl (v1.8.3-based) that supports the following features:
- [x] support for external types in DTO definitions
- [x] code gen plugin mechanism. shygoctl focus on api file parse and construction of api specification object, and the code generation wa moved to plugin   

> original package: https://github.com/zeromicro/go-zero 

## 1. Install
install shygoctl
```
go install -u github.com/shyandsy/shygoctl
```

install gozero code template
```
go install -u github.com/shyandsy/shygoctl_gozero
```

generate gozero code based on api file
```
shygoctl api gen -tpn shygoctl_gozero -api test.api -dir ./ --style=goZero
```

## 2. Features

### use external type in for DTOs
it can specify the source package in import statement like below:
(it also support complex type using in Pointer)

```go
import (
	"time"
    "github.com/shyandsy/shygoctl/common"   // specific the source package path
)

type (
    Book {
        Id []int64 `json:"id"`
		common.BaseBook         // use types defined in common package
        Created  time.Time  `json:"created"`
        Modified *time.Time `json:"modified"`
    }
)
```

command to generate code
```shell
$ shygoctl api go -api test.api -dir ./ --style=goZero
```

it wille generate type.go like below
```go
package types

import (
	"github.com/shyandsy/shygoctl/common"
	"time"
)

Book {
    Id []int64 `json:"id"`
    common.BaseBook
    Created  time.Time  `json:"created"`
    Modified *time.Time `json:"modified"`
}
```

### Generate code by plugin

generate gozero code based on api file
```
shygoctl api gen -tpn shygoctl_gozero -api test.api -dir ./ --style=goZero
```

the parameter `-tpn` specify the name of code gen plugin  

## 3. Example
```shell
# this project folder
$ cd shygoctl/demo   

# generate code base on test.api   
$ shygoctl api go -api test.api -dir ./ --style=goZero       
```

## 4. Customize code gen for any language or framework
The idea is simple. The core workflow operates as follows:
1. Shygoctl generates API specification objects
2. These specifications are persisted to a file
3. The file is passed to the code-generation plugin
4. The plugin retains full control over code generation logic

how to develop your own plugin for your language and framework
please refer to github.com/shyandsy/shygoctl_gozero