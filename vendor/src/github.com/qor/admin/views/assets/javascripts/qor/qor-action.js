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
    let Mustache = window.Mustache,
        NAMESPACE = 'qor.action',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_UNDO = 'undo.' + NAMESPACE,
        ACTION_FORMS = '.qor-action-forms',
        ACTION_HEADER = '.qor-page__header',
        ACTION_BODY = '.qor-page__body',
        ACTION_BUTTON = '.qor-action-button',
        MDL_BODY = '.mdl-layout__content',
        ACTION_SELECTORS = '.qor-actions',
        ACTION_LINK = 'a.qor-action--button',
        MENU_ACTIONS = '.qor-table__actions a[data-url],[data-url][data-method=POST],[data-url][data-method=PUT],[data-url][data-method=DELETE]',
        BUTTON_BULKS = '.qor-action-bulk-buttons',
        QOR_TABLE = '.qor-table-container',
        QOR_TABLE_BULK = '.qor-table--bulking',
        QOR_SEARCH = '.qor-search-container',
        CLASS_IS_UNDO = 'is_undo',
        QOR_SLIDEOUT = '.qor-slideout',

        ACTION_FORM_DATA = 'primary_values[]';

    function QorAction(element, options) {
        this.$element = $(element);
        this.$wrap = $(ACTION_FORMS);
        this.options = $.extend({}, QorAction.DEFAULTS, $.isPlainObject(options) && options);
        this.ajaxForm = {};
        this.init();
    }

    QorAction.prototype = {
        constructor: QorAction,

        init: function() {
            this.bind();
            this.initActions();
        },

        bind: function() {
            this.$element.on(EVENT_CLICK, $.proxy(this.click, this));
            $(document)
                .on(EVENT_CLICK, '.qor-table--bulking tr', this.click)
                .on(EVENT_CLICK, ACTION_LINK, this.actionLink);
        },

        unbind: function() {
            this.$element.off(EVENT_CLICK, this.click);

            $(document)
                .off(EVENT_CLICK, '.qor-table--bulking tr', this.click)
                .off(EVENT_CLICK, ACTION_LINK, this.actionLink);

        },

        initActions: function() {
            this.tables = $(QOR_TABLE).find('table').length;

            if (!this.tables) {
                $(BUTTON_BULKS).find('button').attr('disabled', true);
                $(ACTION_LINK).attr('disabled', true);
            }
        },

        collectFormData: function() {
            let checkedInputs = $(QOR_TABLE_BULK).find('.mdl-checkbox__input:checked'),
                formData = [],
                normalFormData = [],
                tempObj;

            if (checkedInputs.length) {
                checkedInputs.each(function() {
                    let id = $(this).closest('tr').data('primary-key');

                    tempObj = {};
                    if (id) {

                        formData.push({
                            name: ACTION_FORM_DATA,
                            value: id.toString()
                        });

                        tempObj[ACTION_FORM_DATA] = id.toString();
                        normalFormData.push(tempObj);
                    }
                });
            }
            this.ajaxForm.formData = formData;
            this.ajaxForm.normalFormData = normalFormData;
            return this.ajaxForm;
        },

        actionLink: function() {
            // if not in index page
            if (!$(QOR_TABLE).find('table').length) {
                return false;
            }
        },

        actionSubmit: function($action) {
            let $target = $($action);
            this.$actionButton = $target;
            if ($target.data().method) {
                this.submit();
                return false;
            }
        },

        click: function(e) {
            let $target = $(e.target);
            this.$actionButton = $target;

            if ($target.data().ajaxForm) {
                this.collectFormData();
                this.ajaxForm.properties = $target.data();
                this.submit();
                return false;
            }


            if ($target.is('.qor-action--bulk')) {
                this.$wrap.removeClass('hidden');
                $(BUTTON_BULKS).find('button').toggleClass('hidden');
                this.appendTableCheckbox();
                $(QOR_TABLE).addClass('qor-table--bulking');
                $(ACTION_HEADER).find(ACTION_SELECTORS).addClass('hidden');
                $(ACTION_HEADER).find(QOR_SEARCH).addClass('hidden');
            }

            if ($target.is('.qor-action--exit-bulk')) {
                this.$wrap.addClass('hidden');
                $(BUTTON_BULKS).find('button').toggleClass('hidden');
                this.removeTableCheckbox();
                $(QOR_TABLE).removeClass('qor-table--bulking');
                $(ACTION_HEADER).find(ACTION_SELECTORS).removeClass('hidden');
                $(ACTION_HEADER).find(QOR_SEARCH).removeClass('hidden');
            }


            if ($(this).is('tr') && !$target.is('a')) {

                let $firstTd = $(this).find('td').first();

                // Manual make checkbox checked or not
                if ($firstTd.find('.mdl-checkbox__input').get(0)) {
                    let hasPopoverForm = $('body').hasClass('qor-bottomsheets-open') || $('body').hasClass('qor-slideout-open'),
                        $checkbox = $firstTd.find('.mdl-js-checkbox'),
                        slideroutActionForm = $('[data-toggle="qor-action-slideout"]').find('form'),
                        formValueInput = slideroutActionForm.find('.js-primary-value'),
                        primaryValue = $(this).data('primary-key'),
                        $alreadyHaveValue = formValueInput.filter('[value="' + primaryValue + '"]'),
                        isChecked;

                    $checkbox.toggleClass('is-checked');
                    $firstTd.parents('tr').toggleClass('is-selected');

                    isChecked = $checkbox.hasClass('is-checked');

                    $firstTd.find('input').prop('checked', isChecked);

                    if (slideroutActionForm.length && hasPopoverForm) {

                        if (isChecked && !$alreadyHaveValue.length) {
                            slideroutActionForm.prepend('<input class="js-primary-value" type="hidden" name="primary_values[]" value="' + primaryValue + '" />');
                        }

                        if (!isChecked && $alreadyHaveValue.length) {
                            $alreadyHaveValue.remove();
                        }

                    }

                    return false;
                }

            }
        },

        renderFlashMessage: function(data) {
            let flashMessageTmpl = QorAction.FLASHMESSAGETMPL;
            Mustache.parse(flashMessageTmpl);
            return Mustache.render(flashMessageTmpl, data);
        },

        submit: function() {
            let _this = this,
                $parent,
                $element = this.$element,
                $actionButton = this.$actionButton,
                ajaxForm = this.ajaxForm || {},
                properties = ajaxForm.properties || $actionButton.data(),
                url = properties.url,
                undoUrl = properties.undoUrl,
                isUndo = $actionButton.hasClass(CLASS_IS_UNDO),
                isInSlideout = $actionButton.closest(QOR_SLIDEOUT).length,
                needDisableButtons = $element && !isInSlideout;

            if (properties.fromIndex && (!ajaxForm.formData || !ajaxForm.formData.length)) {
                window.alert(ajaxForm.properties.errorNoItem);
                return;
            }

            if (properties.confirm && properties.ajaxForm && !properties.fromIndex) {
                window.QOR.qorConfirm(properties, function(confirm) {
                    if (confirm) {
                        properties = $.extend({}, properties, {
                            _method: properties.method
                        });

                        $.post(properties.url, properties, function() {
                            window.location.reload();
                        });
                    } else {
                        return;
                    }
                });

            } else {

                if (isUndo) {
                    url = properties.undoUrl;
                }

                $.ajax(url, {
                    method: properties.method,
                    data: ajaxForm.formData,
                    dataType: properties.datatype,
                    beforeSend: function() {
                        if (undoUrl) {
                            $actionButton.prop('disabled', true);
                        } else if (needDisableButtons) {
                            _this.switchButtons($element, 1);
                        }
                    },
                    success: function(data, status, response) {
                        let contentType = response.getResponseHeader("content-type"),
                            // handle file download from form submit
                            disposition = response.getResponseHeader('Content-Disposition');

                        if (disposition && disposition.indexOf('attachment') !== -1) {
                            var fileNameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/,
                                matches = fileNameRegex.exec(disposition),
                                fileData = {},
                                fileName = '';

                            if (matches != null && matches[1]) {
                                fileName = matches[1].replace(/['"]/g, '');
                            }

                            if (properties.method) {
                                fileData = $.extend({}, ajaxForm.normalFormData, {
                                    _method: properties.method
                                });
                            }

                            window.QOR.qorAjaxHandleFile(url, contentType, fileName, fileData);

                            if (undoUrl) {
                                $actionButton.prop('disabled', false);
                            } else {
                                _this.switchButtons($element);
                            }

                            return;
                        }

                        // has undo action
                        if (undoUrl) {
                            $element.triggerHandler(EVENT_UNDO, [$actionButton, isUndo, data]);
                            isUndo ? $actionButton.removeClass(CLASS_IS_UNDO) : $actionButton.addClass(CLASS_IS_UNDO);
                            $actionButton.prop('disabled', false);
                            return;
                        }

                        if (contentType.indexOf('json') > -1) {
                            // render notification
                            $('.qor-alert').remove();
                            needDisableButtons && _this.switchButtons($element);
                            isInSlideout ? $parent = $(QOR_SLIDEOUT) : $parent = $(MDL_BODY);
                            $parent.find(ACTION_BODY).prepend(_this.renderFlashMessage(data));
                        } else {
                            // properties.fromIndex || properties.fromMenu
                            window.location.reload();
                        }
                    },
                    error: function(xhr, textStatus, errorThrown) {
                        if (undoUrl) {
                            $actionButton.prop('disabled', false);
                        } else if (needDisableButtons) {
                            _this.switchButtons($element);
                        }
                        window.alert([textStatus, errorThrown].join(': '));
                    }
                });

            }
        },

        switchButtons: function($element, disbale) {
            let needDisbale = disbale ? true : false;
            $element.find(ACTION_BUTTON).prop('disabled', needDisbale);
        },

        destroy: function() {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        },

        // Helper
        removeTableCheckbox: function() {
            $('.qor-page__body .mdl-data-table__select').each(function(i, e) {
                $(e).parents('td').remove();
            });
            $('.qor-page__body .mdl-data-table__select').each(function(i, e) {
                $(e).parents('th').remove();
            });
            $('.qor-table-container tr.is-selected').removeClass('is-selected');
            $('.qor-page__body table.mdl-data-table--selectable').removeClass('mdl-data-table--selectable');
            $('.qor-page__body tr.is-selected').removeClass('is-selected');
        },

        appendTableCheckbox: function() {
            // Only value change and the table isn't selectable will add checkboxes
            $('.qor-page__body .mdl-data-table__select').each(function(i, e) {
                $(e).parents('td').remove();
            });
            $('.qor-page__body .mdl-data-table__select').each(function(i, e) {
                $(e).parents('th').remove();
            });
            $('.qor-table-container tr.is-selected').removeClass('is-selected');
            $('.qor-page__body table').addClass('mdl-data-table--selectable');

            // init google material
            new window.MaterialDataTable($('.qor-page__body table').get(0));
            $('thead.is-hidden tr th:not(".mdl-data-table__cell--non-numeric")').clone().prependTo($('thead:not(".is-hidden") tr'));

            let $fixedHeadCheckBox = $('thead:not(".is-fixed") .mdl-checkbox__input'),
                isMediaLibrary = $('.qor-table--medialibrary').length,
                hasPopoverForm = $('body').hasClass('qor-bottomsheets-open') || $('body').hasClass('qor-slideout-open');

            isMediaLibrary && ($fixedHeadCheckBox = $('thead .mdl-checkbox__input'));

            $fixedHeadCheckBox.on('click', function() {

                if (!isMediaLibrary) {
                    $('thead.is-fixed tr th').eq(0).find('label').click();
                    $(this).closest('label').toggleClass('is-checked');
                }

                let slideroutActionForm = $('[data-toggle="qor-action-slideout"]').find('form'),
                    slideroutActionFormPrimaryValues = slideroutActionForm.find('.js-primary-value');

                if (slideroutActionForm.length && hasPopoverForm) {

                    if ($(this).is(':checked')) {
                        let allPrimaryValues = $('.qor-table--bulking tbody tr');
                        allPrimaryValues.each(function() {
                            let primaryValue = $(this).data('primary-key');
                            if (primaryValue) {
                                slideroutActionForm.prepend('<input class="js-primary-value" type="hidden" name="primary_values[]" value="' + primaryValue + '" />');
                            }
                        });
                    } else {
                        slideroutActionFormPrimaryValues.remove();
                    }
                }

            });
        }

    };
    QorAction.FLASHMESSAGETMPL = (
        `<div class="qor-alert qor-action-alert qor-alert--success [[#error]]qor-alert--error[[/error]]" [[#message]]data-dismissible="true"[[/message]] role="alert">
          <button type="button" class="mdl-button mdl-button--icon" data-dismiss="alert">
            <i class="material-icons">close</i>
          </button>
          <span class="qor-alert-message">
            [[#message]]
              [[message]]
            [[/message]]
            [[#error]]
              [[error]]
            [[/error]]
          </span>
        </div>`
    );

    QorAction.DEFAULTS = {};

    $.fn.qorSliderAfterShow.qorActionInit = function(url, html) {
        let hasAction = $(html).find('[data-toggle="qor-action-slideout"]').length,
            $actionForm = $('[data-toggle="qor-action-slideout"]').find('form'),
            $checkedItem = $('.qor-page__body .mdl-checkbox__input:checked');

        if (hasAction && $checkedItem.length) {
            // insert checked value into sliderout form
            $checkedItem.each(function(i, e) {
                let id = $(e).parents('tbody tr').data('primary-key');
                if (id) {
                    $actionForm.prepend('<input class="js-primary-value" type="hidden" name="primary_values[]" value="' + id + '" />');
                }
            });
        }

    };

    QorAction.plugin = function(options) {
        return this.each(function() {
            let $this = $(this),
                data = $this.data(NAMESPACE),
                fn;

            if (!data) {
                $this.data(NAMESPACE, (data = new QorAction(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.call(data);
            }
        });
    };

    $(function() {
        let options = {},
            selector = '[data-toggle="qor.action.bulk"]';

        $(document).
        on(EVENT_DISABLE, function(e) {
            QorAction.plugin.call($(selector, e.target), 'destroy');
        }).
        on(EVENT_ENABLE, function(e) {
            QorAction.plugin.call($(selector, e.target), options);
        }).
        on(EVENT_CLICK, MENU_ACTIONS, function() {
            (new QorAction()).actionSubmit(this);
            return false;
        }).
        triggerHandler(EVENT_ENABLE);
    });

    return QorAction;

});
