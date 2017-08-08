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

    var $window = $(window);
    var _ = window._;
    var NAMESPACE = 'qor.fixer';
    var EVENT_ENABLE = 'enable.' + NAMESPACE;
    var EVENT_DISABLE = 'disable.' + NAMESPACE;
    var EVENT_CLICK = 'click.' + NAMESPACE;
    var EVENT_RESIZE = 'resize.' + NAMESPACE;
    var EVENT_SCROLL = 'scroll.' + NAMESPACE;
    var CLASS_IS_HIDDEN = 'is-hidden';
    var CLASS_IS_FIXED = 'is-fixed';
    var CLASS_HEADER = '.qor-page__header';

    function QorFixer(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorFixer.DEFAULTS, $.isPlainObject(options) && options);
        this.$clone = null;
        this.init();
    }

    QorFixer.prototype = {
        constructor: QorFixer,

        init: function() {
            var options = this.options;
            var $this = this.$element;
            if (this.buildCheck()) {
                return;
            }
            this.$thead = $this.find('thead:first');
            this.$tbody = $this.find('tbody:first');
            this.$header = $(options.header);
            this.$subHeader = $(options.subHeader);
            this.$content = $(options.content);
            this.marginBottomPX = parseInt(this.$subHeader.css('marginBottom'));
            this.paddingHeight = options.paddingHeight;
            this.fixedHeaderWidth = [];
            this.isEqualed = false;

            this.resize();
            this.bind();
        },

        bind: function() {
            this.$element.on(EVENT_CLICK, $.proxy(this.check, this));

            this.$content.on(EVENT_SCROLL, $.proxy(this.toggle, this));
            $window.on(EVENT_RESIZE, $.proxy(this.resize, this));
        },

        unbind: function() {
            this.$element.off(EVENT_CLICK, this.check);

            this.$content.
            off(EVENT_SCROLL, this.toggle).
            off(EVENT_RESIZE, this.resize);
        },

        build: function() {
            if (!this.$content.length) {
                return;
            }

            var $this = this.$element;
            var $thead = this.$thead;
            var $clone = this.$clone;
            var self = this;
            var $items = $thead.find('> tr').children();
            var pageBodyTop = this.$content.offset().top + $(CLASS_HEADER).height();

            if (!$clone) {
                this.$clone = $clone = $thead.clone().prependTo($this).css({ top: pageBodyTop });
            }

            $clone.
            addClass([CLASS_IS_FIXED, CLASS_IS_HIDDEN].join(' ')).
            find('> tr').
            children().
            each(function(i) {
                $(this).outerWidth($items.eq(i).outerWidth());
                self.fixedHeaderWidth.push($(this).outerWidth());
            });
        },

        unbuild: function() {
            this.$clone.remove();
        },

        buildCheck: function() {
            var $this = this.$element;
            // disable fixer if have multiple tables or in search page or in media library list page
            if ($('.qor-page__body .qor-js-table').length > 1 || $('.qor-global-search--container').length > 0 || $this.hasClass('qor-table--medialibrary') || $this.is(':hidden') || $this.find('tbody > tr:visible').length <= 1) {
                return true;
            }
            return false;
        },

        check: function(e) {
            var $target = $(e.target);
            var checked;

            if ($target.is('.qor-js-check-all')) {
                checked = $target.prop('checked');

                $target.
                closest('thead').
                siblings('thead').
                find('.qor-js-check-all').prop('checked', checked).
                closest('.mdl-checkbox').toggleClass('is-checked', checked);
            }
        },

        toggle: function() {
            if (!this.$content.length) {
                return;
            }
            var self = this;
            var $clone = this.$clone;
            var $thead = this.$thead;
            var scrollTop = this.$content.scrollTop();
            var scrollLeft = this.$content.scrollLeft();
            var offsetTop = this.$subHeader.outerHeight() + this.paddingHeight + this.marginBottomPX;
            var headerHeight = $('.qor-page__header').outerHeight();

            if (!this.isEqualed) {
                this.headerWidth = [];
                var $items = $thead.find('> tr').children();
                $items.each(function() {
                    self.headerWidth.push($(this).outerWidth());
                });
                var notEqualWidth = _.difference(self.fixedHeaderWidth, self.headerWidth);
                if (notEqualWidth.length) {
                    $('thead.is-fixed').find('>tr').children().each(function(i) {
                        $(this).outerWidth(self.headerWidth[i]);
                    });
                    this.isEqualed = true;
                }
            }
            if (scrollTop > offsetTop - headerHeight) {
                $clone.css({ 'margin-left': -scrollLeft }).removeClass(CLASS_IS_HIDDEN);
            } else {
                $clone.css({ 'margin-left': '0' }).addClass(CLASS_IS_HIDDEN);
            }
        },

        resize: function() {
            this.build();
            this.toggle();
        },

        destroy: function() {
            if (this.buildCheck()) {
                return;
            }
            this.unbind();
            this.unbuild();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorFixer.DEFAULTS = {
        header: false,
        content: false
    };

    QorFixer.plugin = function(options) {
        return this.each(function() {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                $this.data(NAMESPACE, (data = new QorFixer(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.call(data);
            }
        });
    };

    $(function() {
        var selector = '.qor-js-table';
        var options = {
            header: '.mdl-layout__header',
            subHeader: '.qor-page__header',
            content: '.mdl-layout__content',
            paddingHeight: 2 // Fix sub header height bug
        };

        $(document).
        on(EVENT_DISABLE, function(e) {
            QorFixer.plugin.call($(selector, e.target), 'destroy');
        }).
        on(EVENT_ENABLE, function(e) {
            QorFixer.plugin.call($(selector, e.target), options);
        }).
        triggerHandler(EVENT_ENABLE);
    });

    return QorFixer;

});