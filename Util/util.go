package Util

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key          string
	responseChan chan<- result
}

type Mem struct {
	requestChan chan request
}

func New(f Func) *Mem {
	mem := &Mem{requestChan: make(chan request)}
	go mem.serve(f)
	return mem
}

func (mem *Mem) Get(key string) (interface{}, error) {
	response := make(chan result)
	mem.requestChan <- request{
		key:          key,
		responseChan: response,
	}
	res := <-response
	return res.value, res.err
}

func (mem *Mem) Close() {
	close(mem.requestChan)
}

func (mem *Mem) serve(f Func) {
	cache := make(map[string]*entry)
	for req := range mem.requestChan {
		e, ok := cache[req.key]
		if !ok {
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.delive(req.responseChan)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) delive(responseChan chan<- result) {
	<-e.ready
	responseChan <- e.res
}
