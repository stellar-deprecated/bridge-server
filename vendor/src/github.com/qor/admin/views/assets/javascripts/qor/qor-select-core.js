(function(factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as anonymous module.
        define(['jquery'], factory);
    } else if (typeof exports === 'object') {
        // Node / CommonJS
        factory(require('jquery'));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function($) {

    'use strict';

    let FormData = window.FormData,
        NAMESPACE = 'qor.selectcore',
        EVENT_SELECTCORE_BEFORESEND = 'selectcoreBeforeSend.' + NAMESPACE,
        EVENT_ONSELECT = 'afterSelected.' + NAMESPACE,
        EVENT_ONSUBMIT = 'afterSubmitted.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_SUBMIT = 'submit.' + NAMESPACE,
        CLASS_CLICK_TABLE = '.qor-table-container tbody tr';

    function QorSelectCore(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSelectCore.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorSelectCore.prototype = {
        constructor: QorSelectCore,

        init: function() {
            this.bind();
        },

        bind: function() {
            this.$element.
                on(EVENT_CLICK, CLASS_CLICK_TABLE, this.processingData.bind(this)).
                on(EVENT_SUBMIT, 'form', this.submit.bind(this));
        },

        unbind: function() {
            this.$element.
                off(EVENT_CLICK, '.qor-table tbody tr', this.processingData.bind(this)).
                off(EVENT_SUBMIT, 'form', this.submit.bind(this));
        },

        processingData: function(e) {
            let $this = $(e.target).closest('tr'),
                data = {},
                url,
                options = this.options,
                onSelect = options.onSelect;

            data = $.extend({}, data, $this.data());
            data.$clickElement = $this;

            url = data.mediaLibraryUrl || data.url;

            if (url) {
                $.getJSON(url, function(json) {
                    json.MediaOption && (json.MediaOption = JSON.parse(json.MediaOption));
                    data = $.extend({}, json, data);
                    if (onSelect && $.isFunction(onSelect)) {
                        onSelect(data, e);
                        $(document).trigger(EVENT_ONSELECT);
                    }
                });
            } else {
                if (onSelect && $.isFunction(onSelect)) {
                    onSelect(data, e);
                    $(document).trigger(EVENT_ONSELECT);
                }
            }
            return false;
        },

        submit: function(e) {
            let form = e.target,
                $form = $(form),
                _this = this,
                $submit = $form.find(':submit'),
                data,
                onSubmit = this.options.onSubmit;

            $(document).trigger(EVENT_SELECTCORE_BEFORESEND);

            if (FormData) {
                e.preventDefault();

                $.ajax($form.prop('action'), {
                    method: $form.prop('method'),
                    data: new FormData(form),
                    dataType: 'json',
                    processData: false,
                    contentType: false,
                    beforeSend: function() {
                        $form.parent().find('.qor-error').remove();
                        $submit.prop('disabled', true);
                    },
                    success: function(json) {
                        data = json;
                        data.primaryKey = data.ID;

                        $('.qor-error').remove();

                        if (onSubmit && $.isFunction(onSubmit)) {
                            onSubmit(data, e);
                            $(document).trigger(EVENT_ONSUBMIT);
                        } else {
                            _this.refresh();
                        }

                    },
                    error: function(xhr, textStatus, errorThrown) {
                        let error;

                        if (xhr.responseJSON) {
                            error = `<ul class="qor-error"><li><label><i class="material-icons">error</i><span>${xhr.responseJSON.errors[0]}</span></label></li></ul>`;
                        } else {
                            error = `<ul class="qor-error">${$(xhr.responseText).find('#errors').html()}</ul>`;
                        }

                        $('.qor-bottomsheets .qor-page__body').scrollTop(0);

                        if (xhr.status === 422 && error) {
                            $form.before(error);
                        } else {
                            window.alert([textStatus, errorThrown].join(': '));
                        }
                    },
                    complete: function() {
                        $submit.prop('disabled', false);
                    }
                });
            }
        },

        refresh: function() {
            setTimeout(function() {
                window.location.reload();
            }, 350);
        },

        destroy: function() {
            this.unbind();
        }

    };


    QorSelectCore.plugin = function(options) {
        return this.each(function() {
            let $this = $(this),
                data = $this.data(NAMESPACE),
                fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }
                $this.data(NAMESPACE, (data = new QorSelectCore(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.apply(data);
            }
        });
    };

    $.fn.qorSelectCore = QorSelectCore.plugin;

    return QorSelectCore;

});
