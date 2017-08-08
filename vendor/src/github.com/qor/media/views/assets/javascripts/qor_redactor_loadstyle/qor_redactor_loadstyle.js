// Load style plugin for redactor editor
// By Jason weng @theplant
//
//
// Plugin name: loadstyle
// required options:
// loadStyleNamespace, loadStyleLink
//
/* ********************************* Usage in qor-example:
product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager,
    Plugins: []admin.RedactorPlugin{
        {Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
        {Name: "loadstyle", Source: "/admin/assets/javascripts/qor_redactor_loadstyle.js"},
    },
    Settings: map[string]interface{}{
        "medialibraryUrl": "/admin/product_images",
        "loadStyleNamespace": "yourNamespace",
        "loadStyleLink": "http://your_stylesheets_file_path",
    }
}})
********************************* */

$.Redactor.prototype.loadstyle = function() {
    return {
        init: function () {
            this.loadstyle.loadStyle();
        },

        loadStyle: function () {
            // loadStyleNamespace, loadStyleLink are required options
            if (typeof this.opts.loadStyleLink === 'undefined' || typeof this.opts.loadStyleNamespace === 'undefined') {
                return;
            }

            var ss = document.createElement('link');

            // insert stylesheet
            ss.type = 'text/css';
            ss.rel = 'stylesheet';
            ss.href = this.opts.loadStyleLink;
            document.getElementsByTagName('head')[0].appendChild(ss);

            // add namespace class into editor
            this.core.editor().addClass(this.opts.loadStyleNamespace);

        }
    };
};