/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package server

import (
	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.http/server"
)

type (
	configProvider interface {
		ConfigServer() *server.Config
	}
)

func config(c container.Container) *server.Config {
	return container.Value[configProvider](c, "config").ConfigServer()
}

/*
####### END ############################################################################################################
*/
