/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

func (impl *implRenderer) writeXML(status int, data any) error {
	buf := new(bytes.Buffer)

	encoder := xml.NewEncoder(buf)

	if err := encoder.Encode(data); err != nil {
		http.Error(impl.ResponseWriter(), err.Error(), http.StatusInternalServerError) /////////////////////////////////
		_ = encoder.Close()

		return err
	}

	if err := encoder.Close(); err != nil {
		http.Error(impl.ResponseWriter(), err.Error(), http.StatusInternalServerError) /////////////////////////////////
		return err
	}

	impl.setContentType()
	impl.ResponseWriter().WriteHeader(status)

	_, err := impl.ResponseWriter().Write(buf.Bytes())
	return err
}

/*
####### END ############################################################################################################
*/
