/*
  warden http proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package mode

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"southwinds.dev/warden/lib"
)

func Basic(address string, verbose bool) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose
	fmt.Println(lib.Banner())
	// lib.InfoLogger.Printf("warden starting @ %s\n", address)
	// if verbose {
	//     lib.InfoLogger.Printf("verbose output enabled\n")
	// }
	log.Fatal(http.ListenAndServe(address, proxy))
}
