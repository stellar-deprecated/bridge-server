package admin

import (
	"errors"

	"github.com/qor/qor/resource"
)

// SingleEditConfig meta configuration used for single edit
type SingleEditConfig struct {
	Template string
	metaConfig
}

// GetTemplate get template for single edit
func (singleEditConfig SingleEditConfig) GetTemplate(context *Context, metaType string) ([]byte, error) {
	if metaType == "form" && singleEditConfig.Template != "" {
		return context.Asset(singleEditConfig.Template)
	}
	return nil, errors.New("not implemented")
}

// ConfigureQorMeta configure single edit meta
func (singleEditConfig *SingleEditConfig) ConfigureQorMeta(metaor resource.Metaor) {
}
