package main

type Token int

const EOF = 0

const (
	tokenNumber     = NUMBER     // integer literal
	tokenIdentifier = IDENTIFIER // programmer defined names
)
