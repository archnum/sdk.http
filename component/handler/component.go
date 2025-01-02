/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package handler

import (
	"net/http"

	"github.com/archnum/sdk.application/container"
)

const (
	_name = "http.handler"
)

type (
	Builder func() (http.Handler, error)

	implComponent struct {
		*container.Component
		builder Builder
	}
)

func New(c container.Container, builder Builder) *implComponent {
	return &implComponent{
		Component: container.NewComponent(_name, c),
		builder:   builder,
	}
}

func Value(c container.Container) http.Handler {
	return container.Value[http.Handler](c, _name)
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	handler, err := impl.builder()
	if err != nil {
		return err
	}

	impl.SetValue(handler)

	return nil
}

/*
####### END ############################################################################################################
*/
