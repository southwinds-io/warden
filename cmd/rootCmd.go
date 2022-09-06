/*
  warden http proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package cmd

import (
	_ "embed"
	"github.com/spf13/cobra"
	"southwinds.dev/warden/lib"
)

type RootCmd struct {
	Cmd *cobra.Command
}

// NewRootCmd
// https://textkool.com/en/ascii-art-generator?hl=default&vl=default&font=Red%20Phoenix&text=Warden
func NewRootCmd() *RootCmd {
	c := &RootCmd{
		Cmd: &cobra.Command{
			Use:     "warden",
			Short:   "warden: internet traffic proxying and inspection",
			Long:    lib.Banner(),
			Version: lib.Version,
		},
	}
	c.Cmd.SetVersionTemplate("version: {{.Version}}\n")
	cobra.OnInitialize(c.initConfig)
	return c
}

func (c *RootCmd) initConfig() {
}
