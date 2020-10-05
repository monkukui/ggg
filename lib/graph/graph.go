package graph

import (
	"io"
	"os"
)

type Graph struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func New() *Graph {
	return &Graph{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
