package Util

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	orr   error
}

type entry struct {
	res   result
	ready chan struct{}
}
