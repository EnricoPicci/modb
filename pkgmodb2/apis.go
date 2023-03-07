package pkgmodb2

import (
	"fmt"

	"github.com/EnricoPicci/modc/pkgmodc1"
)

func ApiThatCallsModuleCApiThatCallRemainsInteral() {
	fmt.Println("I am an API of Module B in pkgmodb2 that calls an API of Module C that remains internal to Module C ")
	pkgmodc1.ApiThatCallsInternalApi()
}
