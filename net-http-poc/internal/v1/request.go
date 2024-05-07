package handler

// Compilation represents a compilation request
// @Description Compilation represents a compilation request
type Compilation struct {
	Fqbn string `json:"fqbn" example:"arduino:avr:uno"`
	OTA  bool   `json:"ota" example:"false"`
	// Path	string `json:"path"`
	Sketch  Sketch `json:"sketch"`
	Verbose bool   `json:"verbose" example:"false"`
}

// Sketch represents a sketch to be compiled
// @Description Sketch represents a sketch to be compiled
type Sketch struct {
	Files    []File            `json:"files"`
	Ino      File              `json:"ino"`
	Metadata IncludedLibraries `json:"metadata"`
	Name     string            `json:"name" example:"test"`
}

// TODO probably we can remove this :point_down:
// IncludedLibraries represents the libraries to be included in the compilation
// @Description IncludedLibraries represents the libraries to be included in the compilation
type IncludedLibraries struct {
	IncludedLibrary []Library `json:"included_libs"`
}

// File represents a file to be compiled
// @Description File represents a file to be compiled
type File struct {
	Name string `json:"name" example:"test.ino"`
	Data string `json:"data" example:"#include <Arduino.h>\nvoid setup() {}\nvoid loop() {}"`
}

// Library represents a library to be included in the compilation
// @Description Library represents a library to be included in the compilation
type Library struct {
	Name    string `json:"name" example:"ArduinoJson"`
	Version string `json:"version" example:"6.17.2"`
}
