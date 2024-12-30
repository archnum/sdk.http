/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package context

type (
	ErrorWithStatus interface {
		error
		Status() int
	}
)

/*
####### END ############################################################################################################
*/
