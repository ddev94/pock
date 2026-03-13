package helpers

import (
	"fmt"
	"pock/internal/utils"
)

func PrintSuccessLine(message string) {
	fmt.Printf("%s %s\n", utils.Green(SymbolSuccess), message)
}

func PrintInfoLine(message string) {
	fmt.Printf("%s %s\n", utils.Blue(SymbolInfo), message)
}
