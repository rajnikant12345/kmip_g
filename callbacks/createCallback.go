package callbacks

import "context"


type CreateCallBackType func(context.Context  ,map[string]interface{}) ( string, error)