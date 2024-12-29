/*
####### sdk.application (c) 2024 Archivage Numérique ############################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package component

import (
	"net/http"

	"github.com/archnum/sdk.application/container"
)

func Handler(c container.Container) http.Handler {
	return container.Value[http.Handler](c, "http.handler")
}

/*
####### END ############################################################################################################
*/
