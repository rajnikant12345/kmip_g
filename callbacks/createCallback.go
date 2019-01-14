package callbacks

import (
	"context"
	"github.com/rajnikant12345/kmip_g/objects"
)


type CreateCallBackType func(context.Context  ,map[string]interface{} ,  []*objects.Name) ( string, error)