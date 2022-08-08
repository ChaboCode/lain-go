package main

type (
	IdentifierEntry uint64
	Identifier      string
)

type SymbolHashTable map[IdentifierEntry]Identifier

var symbolTable SymbolHashTable

func hash(identifier Identifier) IdentifierEntry {
	var h IdentifierEntry = 0
	for i := 0; i < len(identifier); i++ {
		h = IdentifierEntry(31)*h + IdentifierEntry(identifier[i])
	}
	return h
}

func InsertIdentifier(identifier Identifier) {
	key := hash(identifier)
	if symbolTable == nil {
		symbolTable = make(map[IdentifierEntry]Identifier)
	}
	symbolTable[key] = identifier
}

func GetIdentifier(key IdentifierEntry) Identifier {
	return symbolTable[key]
}
