/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"strings"
	"testing"
)

const (
	_path = "a/b/c//d/e/f//"
)

func Benchmark_StringsSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Split(_path, "/")
	}
}

func Benchmark_splitPath(b *testing.B) {
	var (
		s string
		r string
	)

	for i := 0; i < b.N; i++ {
		s, r = splitPath(_path)

		for s != "" {
			s, r = splitPath(r)
		}
	}
}

/*
####### END ############################################################################################################
*/
