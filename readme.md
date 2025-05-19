# cutomized goctl

## version 0.1 feature
support external type in dto, we can specific the source package in import statement like below:
(it also support complex type using in Pointer)

```go
import (
	"time"
    "github.com/shyandsy/shy-goctl/common"
)

type (
    Book {
        Id []int64 `json:"id"`
        BaseBook
        Created  time.Time  `json:"created"`
        Modified *time.Time `json:"modified"`
    }
)
```

it wille generate type.go like below
```go
package types

import (
	"github.com/shyandsy/shy-goctl/common"
	"time"
)

......


```


2. suppoei