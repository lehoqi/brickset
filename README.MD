## Abstract
brickset is a Go client library for accessing the brickset.com [API v3](https://brickset.com/article/52664/api-version-3-documentation).

If using this API, you may need to complete the following:
- Register for the brickset.com,click [here](https://brickset.com/login)
- All methods require a valid API key to be passed, which can be obtained [here](https://brickset.com/tools/webservices/requestkey).

## Interfaces
```go
type IBrickSet interface {
	GetSets(ctx context.Context, params *GetSetRequest) (int, []*Sets, error)
	GetThemes(ctx context.Context) (int, []*Themes, error)
	GetReviews(ctx context.Context, setID int) (int, []*Review, error)
	GetSubthemes(ctx context.Context, theme string) (int, []*Subthemes, error)
	GetInstructions(ctx context.Context, setID int) (int, []*Instruction, error)
	GetInstructions2(ctx context.Context, setNumber string) (int, []*Instruction, error)
	GetAdditionalImages(ctx context.Context, setID int) (int, []*Image, error)
	GetYears(ctx context.Context, theme string) (int, []*Years, error)
}
```
## Installation

```bash
# Go Modules
require github.com/lehoqi/brickset
```

## Usage

The following samples will assist you to become as comfortable as possible with brickset library.

```go
package main

import (
	"context"
	"github.com/lehoqi/brickset"
	"log"
)

func main() {
	ctx := context.Background()
	svc := brickset.New("api-key", "username", "password", brickset.WithDebug(true))
	_, themes, err := svc.GetThemes(ctx)
	if err != nil {
		panic(err)
	}
	for _, theme := range themes {
		_, sets, err := svc.GetSets(ctx, &brickset.GetSetRequest{Theme: theme.Theme, PageSize: 2})
		if err != nil {
			panic(err)
		}
		for _, set := range sets {
			log.Println(set.Name)
		}
	}
}
```
## Tests

```shell
go test -v
```