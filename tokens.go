package main

type Token int

const EOF = 0

const (
	tokenNumber     = NUMBER     // integer literal
	tokenIdentifier = IDENTIFIER // programmer defined names
	tokenSeparator  = SEPARATOR  // separator e.g. semicolon
	tokenLet        = LET        // let keyword
)
