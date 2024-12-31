/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import "strings"

func splitPath(path string) (string, string) {
	n := strings.Index(path, "/")
	if n == -1 {
		return path, ""
	}

	s := path[:n]
	if s == "" {
		return splitPath(path[n+1:])
	}

	return s, path[n+1:]
}

/*
####### END ############################################################################################################
*/
