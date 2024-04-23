package handler

var compilationsQueue map[string]CompilationStatus

type CompilationStatus struct {
	Compilation Compilation `json:"compilation"`
	Status      string      `json:"status"`
}

// Init initializes the compilations queue
func Init() {
	compilationsQueue = make(map[string]CompilationStatus)
}
