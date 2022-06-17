package main

type Token int

const EOF = 0

const (
	tokenInt        = NUMBER     // integer literal
	tokenIdentifier = IDENTIFIER // programmer defined names
)
