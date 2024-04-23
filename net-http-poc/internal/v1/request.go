package handler

type Compilation struct {
	Fqbn string `json:"fqbn"`
	OTA  bool   `json:"ota"`
	// Path	string `json:"path"`
	Sketch  Sketch `json:"sketch"`
	Verbose bool   `json:"verbose"`
}

type Sketch struct {
	Files    []File    `json:"files"`
	Ino      File      `json:"ino"`
	Metadata []Library `json:"metadata"`
	Name     string    `json:"name"`
}

type File struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Library struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
