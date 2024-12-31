/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

type (
	dataError struct {
		Error     string `json:"error" xml:"error" yaml:"error"`
		RequestID string `json:"request_id" xml:"request_id" yaml:"request_id"`
	}
)

/*
####### END ############################################################################################################
*/
