package main

import (
	"testing"
)

func TestCircularType(t *testing.T)   { checkOutputEqual(t, "circulartype.go") }
func TestEmbeddedStruct(t *testing.T) { checkOutputEqual(t, "structs/embed.go") }
