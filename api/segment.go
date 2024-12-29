/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

type (
	segment struct {
		childs map[string]*segment
		fns    map[string]HandlerFunc
	}
)

func newSegment() *segment {
	return &segment{
		childs: make(map[string]*segment),
	}
}

func (seg *segment) addHandlerFunc(method string, fn HandlerFunc) {
	if seg.fns == nil {
		seg.fns = make(map[string]HandlerFunc)
	}

	seg.fns[method] = fn
}

/*
####### END ############################################################################################################
*/
