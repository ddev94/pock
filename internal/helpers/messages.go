package helpers

import (
	"fmt"
	"pock/internal/utils"
)

func PrintCommandNotFound(name string) {
	fmt.Printf("%s Command \"%s\" not found!\n", utils.Red(SymbolError), name)
	fmt.Printf("%s\n", utils.Blue("Use \"pock list\" to see all saved commands."))
}

func PrintFeatureNotImplemented(nextStep string) {
	fmt.Printf("%s\n", utils.Yellow("Marketplace integration not yet implemented."))
	fmt.Printf("%s %s\n", utils.Blue(SymbolInfo), nextStep)
}
