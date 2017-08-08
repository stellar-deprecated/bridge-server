package admin

import (
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
)

type RichEditorConfig struct {
	AssetManager *Resource
	Plugins      []RedactorPlugin
	Settings     map[string]interface{}
	metaConfig
}

type RedactorPlugin struct {
	Name   string
	Source string
}

// ConfigureQorMeta configure rich editor meta
func (richEditorConfig *RichEditorConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*Meta); ok {
		meta.Type = "rich_editor"

		// Compatible with old rich editor setting
		if meta.Resource != nil {
			richEditorConfig.AssetManager = meta.Resource
			meta.Resource = nil
		}

		setter := meta.GetSetter()
		meta.SetSetter(func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			metaValue.Value = utils.HTMLSanitizer.Sanitize(utils.ToString(metaValue.Value))
			setter(resource, metaValue, context)
		})

		if richEditorConfig.Settings == nil {
			richEditorConfig.Settings = map[string]interface{}{}
		}

		plugins := []string{"source"}
		for _, plugin := range richEditorConfig.Plugins {
			plugins = append(plugins, plugin.Name)
		}
		richEditorConfig.Settings["plugins"] = plugins
	}
}
