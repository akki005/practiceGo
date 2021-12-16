package types

//install all go tools globally with `go install golang.org/x/tools/...@latest`
//https://medium.com/a-journey-with-go/go-stringer-command-efficiency-through-code-generation-df49f97f3954
//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Pcm
	Aspirin
)
