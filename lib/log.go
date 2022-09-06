/*
  warden http proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package lib

import (
	"log"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
)

// func init() {
//     InfoLogger = log.New(os.Stdout, "W-INFO: ", log.Ldate|log.Ltime|log.Lmsgprefix|log.LUTC|log.Lmicroseconds)
//     WarningLogger = log.New(os.Stdout, "W-WARNING: ", log.Ldate|log.Ltime|log.Lmsgprefix|log.LUTC|log.Lmicroseconds)
//     ErrorLogger = log.New(os.Stderr, "W-ERROR: ", log.Ldate|log.Ltime|log.Lmsgprefix|log.LUTC|log.Lmicroseconds)
//     DebugLogger = log.New(os.Stdout, "W-DEBUG: ", log.Ldate|log.Ltime|log.Lmsgprefix|log.LUTC|log.Lmicroseconds)
// }
