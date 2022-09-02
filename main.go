package main

import (
	"fmt"
	"os"
	"strings"
)

var globalStore = make(map[string]string)

type Map = map[string]string

type Transaction struct {
	store Map
	next *Transaction
}


