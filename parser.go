package cudnn_log_parser

import parsec "github.com/prataprc/goparsec"

var (

	// terminal parsers.
	tagobrk  = parsec.Atom("<", "OT")
	tagcbrk  = parsec.Atom(">", "CT")
	tagcend  = parsec.Atom("/>", "CT")
	tagcopen = parsec.Atom("</", "CT")
	equal    = parsec.Atom(`=`, "EQ")
	single   = parsec.Atom("'", "SQUOTE")
	double   = parsec.Atom(`"`, "DQUOTE")
	tagname  = parsec.Token(`[a-zA-Z0-9]+`, "TAGNAME")
	attrname = parsec.Token(`[a-zA-Z0-9_-]+`, "ATTRNAME")
	attrval1 = parsec.Token(`[^\s"'=<>`+"`]+", "ATTRVAL1")
	attrval2 = parsec.Token(`[^']*`, "ATTRVAL2")
	attrval3 = parsec.Token(`[^"]*`, "ATTRVAL3")
	entity   = parsec.Token(`&#?[a-bA-Z0-9]+;`, "ENTITY")
	text     = parsec.Token(`[^<]+`, "TEXT")
	doctype  = parsec.Token(`<!doctype[^>]+>`, "DOCTYPE")
)
