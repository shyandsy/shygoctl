# shyandsy/shygoctl

This is an customized goctl version based on v1.8.3, which support external types in dto
- [x] generate golang code for gozero

> original package: https://github.com/zeromicro/go-zero 

## 1. Install
```
go install -u github.com/shyandsy/shygoctl
```

## 2. Getting Started

support external type in dto, we can specific the source package in import statement like below:
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

## 3. Example
```shell
# this project folder
$ cd shygoctl/demo   

# generate code base on test.api   
$ shygoctl api go -api test.api -dir ./ --style=goZero       
```
