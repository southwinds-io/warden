/*
  warden http proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package lib

import (
	"fmt"
)

func Banner() string {
	return fmt.Sprintf(`
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++
|   __      __                    .___                  |
|  /  \    /  \_____  _______   __| _/ ____    ____     |
|  \   \/\/   /\__  \ \_  __ \ / __ |_/ __ \  /    \    |
|   \        /  / __ \_|  | \// /_/ |\  ___/ |   |  \   |
|    \__/\  /  (____  /|__|   \____ | \___  >|___|  /   |
|         \/        \/             \/     \/      \/    |
|            Traffic Proxying and Inspection            |
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++

version: %s
`, Version)
}
