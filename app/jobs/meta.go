package jobs

import (
	"log"
)

type CreateMeta struct{}

func (c CreateMeta) Run() {
	log.Print("CreateMeta")
	log.Print("Run")
}

func init() {
}
