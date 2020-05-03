package main

// Termograph ...
type Termograph struct {
	ID     string
	Probes []Probe
}

// Probe ...
type Probe struct {
	ID          string
	Temperature float32
}
