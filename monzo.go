package gomonzo

import (
	"github.com/0xdeafcafe/gomonzo/helpers"
	"github.com/0xdeafcafe/gomonzo/models"
)

// GoMonzo is a pretty swell go client for Monzo
type GoMonzo struct {
	options    *models.MonzoOptions
	httpHelper *helpers.HTTPHelper
}

// New creates a new implementation of the MonzoGo client
func New(options *models.MonzoOptions) *GoMonzo {
	goMonzo := &GoMonzo{
		options:    options,
		httpHelper: helpers.NewHTTPHelper(),
	}

	return goMonzo
}
