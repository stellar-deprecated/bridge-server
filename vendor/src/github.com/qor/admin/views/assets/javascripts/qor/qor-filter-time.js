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

    var location = window.location;
    var $document = $(document);
    var NAMESPACE = 'qor.filter';
    var EVENT_FILTER_CHANGE = 'filterChanged.' + NAMESPACE;
    var EVENT_ENABLE = 'enable.' + NAMESPACE;
    var EVENT_DISABLE = 'disable.' + NAMESPACE;
    var EVENT_CLICK = 'click.' + NAMESPACE;
    var CLASS_BOTTOMSHEETS = '.qor-bottomsheets';
    var CLASS_DATE_START = '.qor-filter__start';
    var CLASS_DATE_END = '.qor-filter__end';
    var CLASS_SEARCH_PARAM = '[data-search-param]';
    var CLASS_FILTER_SELECTOR = '.qor-filter__dropdown';
    var CLASS_FILTER_TOGGLE = '.qor-filter-toggle';
    var CLASS_IS_SELECTED = 'is-selected';

    function QorFilterTime(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorFilterTime.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorFilterTime.prototype = {
        constructor: QorFilterTime,

        init: function() {
            this.bind();
            var $element = this.$element;

            this.$timeStart = $element.find(CLASS_DATE_START);
            this.$timeEnd = $element.find(CLASS_DATE_END);
            this.$searchParam = $element.find(CLASS_SEARCH_PARAM);
            this.$searchButton = $element.find(this.options.button);

            this.startWeekDate = window.moment().startOf('isoweek').toDate();
            this.endWeekDate = window.moment().endOf('isoweek').toDate();

            this.startMonthDate = window.moment().startOf('month').toDate();
            this.endMonthDate = window.moment().endOf('month').toDate();
            this.initActionTemplate();

        },

        bind: function() {
            var options = this.options;

            this.$element.
            on(EVENT_CLICK, options.trigger, this.show.bind(this)).
            on(EVENT_CLICK, options.label, this.setFilterTime.bind(this)).
            on(EVENT_CLICK, options.clear, this.clear.bind(this)).
            on(EVENT_CLICK, options.button, this.search.bind(this));

            $document.on(EVENT_CLICK, this.close);
        },

        unbind: function() {
            var options = this.options;
            this.$element.
            off(EVENT_CLICK, options.trigger, this.show.bind(this)).
            off(EVENT_CLICK, options.label, this.setFilterTime.bind(this)).
            off(EVENT_CLICK, options.clear, this.clear.bind(this)).
            off(EVENT_CLICK, options.button, this.search.bind(this));
        },

        initActionTemplate: function() {
            var scheduleStartAt = this.getUrlParameter('schedule_start_at'),
                scheduleEndAt = this.getUrlParameter('schedule_end_at'),
                $filterToggle = $(this.options.trigger);

            if (scheduleStartAt || scheduleEndAt) {
                this.$timeStart.val(scheduleStartAt);
                this.$timeEnd.val(scheduleEndAt);

                scheduleEndAt = !scheduleEndAt ? '' : ' - ' + scheduleEndAt;
                $filterToggle.addClass('active clearable').find('.qor-selector-label').html(scheduleStartAt + scheduleEndAt);
                $filterToggle.append('<i class="material-icons qor-selector-clear">clear</i>');
            }
        },

        show: function() {
            this.$element.find(CLASS_FILTER_SELECTOR).toggle();
        },

        close: function(e) {
            var $target = $(e.target),
                $filter = $(CLASS_FILTER_SELECTOR),
                filterVisible = $filter.is(':visible'),
                isInFilter = $target.closest(CLASS_FILTER_SELECTOR).length,
                isInToggle = $target.closest(CLASS_FILTER_TOGGLE).length,
                isInModal = $target.closest('.qor-modal').length,
                isInTimePicker = $target.closest('.ui-timepicker-wrapper').length;

            if (filterVisible && (isInFilter || isInToggle || isInModal || isInTimePicker)) {
                return;
            }
            $filter.hide();
        },

        setFilterTime: function(e) {
            var $target = $(e.target),
                data = $target.data(),
                range = data.filterRange,
                startTime, endTime, startDate, endDate;

            if (!range) {
                return false;
            }

            $(this.options.label).removeClass(CLASS_IS_SELECTED);
            $target.addClass(CLASS_IS_SELECTED);

            if (range == 'events') {
                this.$timeStart.val(data.scheduleStartAt || '');
                this.$timeEnd.val(data.scheduleEndAt || '');
                this.$searchButton.click();
                return false;
            }

            switch (range) {
                case 'today':
                    startDate = endDate = new Date();
                    break;
                case 'week':
                    startDate = this.startWeekDate;
                    endDate = this.endWeekDate;
                    break;
                case 'month':
                    startDate = this.startMonthDate;
                    endDate = this.endMonthDate;
                    break;
            }

            if (!startDate || !endDate) {
                return false;
            }

            startTime = this.getTime(startDate) + ' 00:00';
            endTime = this.getTime(endDate) + ' 23:59';

            this.$timeStart.val(startTime);
            this.$timeEnd.val(endTime);
            this.$searchButton.click();
        },

        getTime: function(dateNow) {
            var month = dateNow.getMonth() + 1,
                date = dateNow.getDate();

            month = (month < 8) ? ('0' + month) : month;
            date = (date < 10) ? ('0' + date) : date;

            return (dateNow.getFullYear() + '-' + month + '-' + date);
        },

        clear: function() {
            var $trigger = $(this.options.trigger),
                $label = $trigger.find('.qor-selector-label');

            $trigger.removeClass('active clearable');
            $label.html($label.data('label'));
            this.$timeStart.val('');
            this.$timeEnd.val('');

            this.$searchButton.click();
            return false;

        },

        getUrlParameter: function(name) {
            var search = location.search;
            name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
            var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
            var results = regex.exec(search);
            return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
        },

        updateQueryStringParameter: function(key, value, uri) {
            var href = uri || location.href,
                escapedkey = String(key).replace(/[\\^$*+?.()|[\]{}]/g, '\\$&'),
                re = new RegExp('([?&])' + escapedkey + '=.*?(&|$)', 'i'),
                separator = href.indexOf('?') !== -1 ? '&' : '?';

            if (href.match(re)) {
                if (value) {
                    return href.replace(re, '$1' + key + '=' + value + '$2');
                } else {
                    if (RegExp.$1 === '?' || RegExp.$1 === RegExp.$2) {
                        return href.replace(re, '$1');
                    } else {
                        return href.replace(re, '');
                    }
                }
            } else if (value) {
                return href + separator + key + '=' + value;
            }

            return href;
        },

        search: function() {
            var $searchParam = this.$searchParam,
                uri,
                _this = this,
                type = 'qor.filter.time';


            if (!$searchParam.length) {
                return;
            }

            $searchParam.each(function() {
                var $this = $(this),
                    searchParam = $this.data().searchParam,
                    val = $this.val();
                uri = _this.updateQueryStringParameter(searchParam, val, uri);
            });

            if (this.$element.closest(CLASS_BOTTOMSHEETS).length) {
                $(CLASS_BOTTOMSHEETS).trigger(EVENT_FILTER_CHANGE, [uri, type]);
            } else {
                location.href = uri;
            }
        },

        destroy: function() {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorFilterTime.DEFAULTS = {
        label: false,
        trigger: false,
        button: false,
        clear: false
    };

    QorFilterTime.plugin = function(options) {
        return this.each(function() {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorFilterTime(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.apply(data);
            }
        });
    };

    $(function() {
        var selector = '[data-toggle="qor.filter.time"]';
        var options = {
            label: '.qor-filter__block-buttons button',
            trigger: 'a.qor-filter-toggle',
            button: '.qor-filter__button-search',
            clear: '.qor-selector-clear'
        };

        $(document).
        on(EVENT_DISABLE, function(e) {
            QorFilterTime.plugin.call($(selector, e.target), 'destroy');
        }).
        on(EVENT_ENABLE, function(e) {
            QorFilterTime.plugin.call($(selector, e.target), options);
        }).
        triggerHandler(EVENT_ENABLE);
    });

    return QorFilterTime;

});