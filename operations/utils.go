package operations

func GetAttribute(name string, m map[string]interface{}) interface{} {
	if val,ok :=  m[name];ok {
		return val
	}
	return nil
}
