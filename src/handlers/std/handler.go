package std

import (
	"../file"
	"bufio"
	"os"
)

type StdHandler struct {
	file.FileHandler
}

func NewHandler() *StdHandler {
	writer := bufio.NewWriter(os.Stdout)
	return &StdHandler{file.FileHandler{writer}}
}
