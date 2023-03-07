package pkgmodb3

import (
	"fmt"

	"github.com/EnricoPicci/modc/pkgmodc2"
)

func ApiThatCallsModuleCApiThatCallsModuleDApi() {
	fmt.Println("I am an API of Module B in pkgmodb3 that calls an API of Module C that calls an API of Module D ")
	pkgmodc2.ApiThatCallsModuleDApi()
}
