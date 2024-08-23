package main

import (
	"github.com/Gabrieltrinidad0101/Make-Language/src"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

type Javascript struct{}

type Dom struct{}

func New(element *parser.Element) {
	src.NewMakeLanguage("./conf.json", "./main.mkl")
}
