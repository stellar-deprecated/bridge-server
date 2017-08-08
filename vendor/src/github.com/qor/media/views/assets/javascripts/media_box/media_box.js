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

    let $body = $('body'),
        $document = $(document),
        NAMESPACE = 'qor.medialibrary.select',
        PARENT_NAMESPACE = 'qor.bottomsheets',
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_RELOAD = 'reload.' + PARENT_NAMESPACE,
        CLASS_SELECT_ICON = '.qor-select__select-icon',
        CLASS_SELECT_HINT = '.qor-selectmany__hint',
        CLASS_PARENT = '.qor-field__mediabox',
        CLASS_LISTS = '.qor-field__mediabox-list',
        CLASS_ITEM = '.qor-field__mediabox-item',
        CLASS_LISTS_DATA = '.qor-field__mediabox-data',
        CLASS_SELECTED = 'is_selected',
        CLASS_DELETE = 'is_deleted',
        CLASS_CROPPER_OPTIONS = 'textarea.qor-file__options',
        CLASS_CROPPER_DELETE = '.qor-cropper__toggle--delete',
        CLASS_CROPPER_UNDO = '.qor-cropper__toggle--undo',
        CLASS_MEDIABOX = 'qor-bottomsheets__mediabox';

    function getYoutubeID(url) {
        var regExp = /^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*/;
        var match = url.match(regExp);
        if (match && match[7].length == 11) {
            return match[7];
        } else {
            return false;
        }
    }

    function QorMediaBox(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorMediaBox.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorMediaBox.prototype = {
        constructor: QorMediaBox,

        init: function() {
            var $element = this.$element;
            this.SELECT_MEDIABOX_UNDO_TEMPLATE = $element.find('[name="media-box-undo-delete"]').html();
            this.bind();
            this.initSelectedMedia();
        },

        bind: function() {
            $document.on(EVENT_RELOAD, '.' + CLASS_MEDIABOX, this.reloadData.bind(this));

            this.$element
                .on(EVENT_CLICK, '[data-mediabox-url]', this.openBottomSheets.bind(this))
                .on(EVENT_CLICK, CLASS_CROPPER_DELETE, this.deleteSelected.bind(this))
                .on(EVENT_CLICK, CLASS_CROPPER_UNDO, this.undoDeleteSelected.bind(this))
                .on('change.qor.cropper', CLASS_CROPPER_OPTIONS, this.imageCrop.bind(this));
        },

        deleteSelected: function(e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_ITEM);

            $selectFeild.addClass(CLASS_DELETE).append(this.SELECT_MEDIABOX_UNDO_TEMPLATE).find('.qor-file__list').hide();
            this.updateMediaLibraryData($target.closest(CLASS_LISTS));
            this.$element.find(CLASS_LISTS_DATA).data('isDeleted', true);

            return false;
        },

        undoDeleteSelected: function(e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_ITEM);

            $selectFeild.removeClass(CLASS_DELETE).find('.qor-file__list').show();
            this.updateMediaLibraryData($target.closest(CLASS_LISTS));
            $target.closest('.qor-fieldset__alert').remove();
            this.$element.find(CLASS_LISTS_DATA).data('isDeleted', false);

            return false;
        },

        imageCrop: function(e) {
            var $parent = $(e.target).closest(CLASS_ITEM);
            this.syncImageCrop($parent);
        },

        openBottomSheets: function(e) {
            var $ele = $(e.target).closest('[data-mediabox-url]'),
                data = $ele.data(),
                $parent;

            if (data.isDisabled) {
                return;
            }

            this.BottomSheets = $body.data('qor.bottomsheets');
            this.bottomsheetsData = data;
            this.$parent = $parent = $ele.closest(CLASS_PARENT);
            this.$selectFeild = $parent.find(CLASS_LISTS);

            data.url = data.mediaboxUrl;

            // select many templates
            this.SELECT_MANY_SELECTED_ICON = $('[name="select-many-selected-icon"]').html();
            this.SELECT_MANY_HINT = $('[name="select-many-hint"]').html();

            this.TEMPLATE_IMAGE = $parent.find('[name="media-box-template"]').html();
            this.TEMPLATE_FILE = $parent.find('[name="media-box-file-template"]').html();
            this.TEMPLATE_UPLOADEDVIDEO = $parent.find('[name="media-box-uploadedvideo-template"]').html();
            this.TEMPLATE_VIDEOLINK = $parent.find('[name="media-box-videolink-template"]').html();
            this.SELECT_MEDIABOX_UNDO_TEMPLATE = $parent.find('[name="media-box-undo-delete"]').html();

            this.BottomSheets.open(data, this.handleSelectMany.bind(this));
        },

        initSelectedMedia: function() {
            var $element = this.$element,
                $selectedMedias = $element.find(CLASS_ITEM),
                $selectedMedia,
                selectedMediaData,
                mediaData = JSON.parse($element.find(CLASS_LISTS_DATA).val());

            if (mediaData) {
                for (var i = 0; i < mediaData.length; i++) {
                    $selectedMedia = $selectedMedias.filter('[data-primary-key="' + mediaData[i].ID + '"]');
                    selectedMediaData = $selectedMedia.data().description;
                    if (!selectedMediaData) {
                        $selectedMedia.data('description', mediaData[i].Description);
                    }
                }
            }
        },

        initMedia: function() {
            var $selectFeild = this.$selectFeild,
                $items = $selectFeild.find(CLASS_ITEM).not('.' + CLASS_DELETE),
                $trs = this.$bottomsheets.find('tbody tr'),
                _this = this,
                $tr,
                $img,
                key;

            $items.each(function() {
                key = $(this).data().primaryKey;
                $tr = $trs.filter('[data-primary-key="' + key + '"]').addClass(CLASS_SELECTED);
                _this.changeIcon($tr, true);
            });

            $trs.each(function() {
                $tr = $(this);
                $img = $tr.find('.qor-table--ml-slideout p img').first();
                $tr.find('.qor-table__actions').remove();
                if ($img.length) {
                    $tr.find('.qor-table--medialibrary-item').css('background-image', 'url(' + $img.prop('src') + ')');
                    $img.parent().remove();
                }
            });
            if (this.bottomsheetsData.maxItem != '1') {
                this.updateHint(this.getSelectedItemData());
            }
        },

        reloadData: function() {
            this.$selectFeild && this.initMedia();
        },

        renderHint: function(data) {
            return window.Mustache.render(this.SELECT_MANY_HINT, data);
        },

        getSelectedItemData: function($ele) {
            var $selectFeild = $ele ? $ele : this.$selectFeild,
                $items = $selectFeild.find(CLASS_ITEM).not('.' + CLASS_DELETE),
                files = [],
                item;

            if ($items.size()) {
                $items.each(function() {
                    item = $(this).data();

                    files.push({
                        ID: item.primaryKey,
                        Url: item.originalUrl.replace(/.original.(\w+)$/, '.$1'),
                        Description: item.description,
                        FileName: item.fileName,
                        VideoLink: item.videolink
                    });
                });
            }

            return {
                files: files,
                selectedNum: files.length
            };
        },

        updateHint: function(data) {
            var template;

            $.extend(data, this.bottomsheetsData);
            template = this.renderHint(data);

            $(CLASS_SELECT_HINT).remove();
            this.$bottomsheets.find('.qor-page__body').before(template);
        },

        updateMediaLibraryData: function($ele, data) {
            var $dataInput = $ele ? $ele.find(CLASS_LISTS_DATA) : this.$selectFeild.find(CLASS_LISTS_DATA),
                fileData = this.getSelectedItemData($ele);

            $dataInput.val(JSON.stringify(fileData.files)).data('mediaData', data).trigger('changed.medialibrary', [data]);
        },

        changeIcon: function($ele, isNew) {
            var $item = $ele.find('.qor-table--medialibrary-item'),
                $target = $item.size() ? $item : $ele.find('td:first');

            $ele.find(CLASS_SELECT_ICON).remove();

            if (isNew) {
                if (isNew == 'one') {
                    $('.' + CLASS_MEDIABOX).find(CLASS_SELECT_ICON).remove();
                }
                $target.prepend(this.SELECT_MANY_SELECTED_ICON);
            }
        },

        syncImageCrop: function($ele, callback) {
            let item = JSON.parse($ele.find(CLASS_CROPPER_OPTIONS).val()),
                url = $ele.data().mediaLibraryUrl,
                syncData = {},
                sizes = ['Width', 'Height'],
                sizeResolutionData,
                sizeData,
                $imgs = $ele.find('img[data-size-name]');

            delete item.ID;
            delete item.Url;

            item.Sizes = {};

            $imgs.each(function() {
                sizeData = $(this).data();

                if (sizeData['data-size-resolution-width'] || sizeData['data-size-resolution']) {
                    item['Sizes'][sizeData.sizeName] = {};
                    for (let i = 0; i < sizes.length; i++) {
                        sizeResolutionData = sizeData['sizeResolution' + sizes[i]];
                        if (!sizeResolutionData) {
                            sizeResolutionData = sizeData['sizeResolution'][sizes[i]];
                        }
                        item['Sizes'][sizeData.sizeName][sizes[i]] = sizeResolutionData;
                    }
                }
            });

            syncData.MediaOption = JSON.stringify(item);

            $.ajax({
                type: 'PUT',
                url: url,
                data: JSON.stringify(syncData),
                contentType: 'application/json',
                dataType: 'json',
                success: function(data) {
                    syncData.MediaOption = JSON.parse(data.MediaOption);

                    if (callback && $.isFunction(callback)) {
                        callback(syncData, $ele);
                    }
                }
            });
        },

        showHiddenItem: function($hiddenItem) {
            $hiddenItem.removeClass(CLASS_DELETE).find('.qor-file__list').show();
            $hiddenItem.find('.qor-fieldset__alert').remove();
        },

        removeItem: function(data) {
            let primaryKey = data.primaryKey;

            this.$selectFeild.find('[data-primary-key="' + primaryKey + '"]').remove();
            this.changeIcon(data.$clickElement);
        },

        compareCropSizes: function(data) {
            let cropOptions = data.MediaOption.CropOptions,
                needCropSizes = this.bottomsheetsData.cropSizes,
                needCropSizesSize,
                cropOptionsKeys;

            if (!needCropSizes || data.SelectedType != 'image') {
                return false;
            }

            needCropSizes = needCropSizes.split(',');
            needCropSizesSize = needCropSizes.length - 1;

            if (window._.isObject(cropOptions)) {
                cropOptionsKeys = Object.keys(cropOptions);
            } else {
                return false;
            }

            if (cropOptionsKeys.length) {
                for (let i = 0; i < needCropSizesSize; i++) {
                    if (cropOptionsKeys.indexOf(needCropSizes[i]) == -1) {
                        return true;
                    }
                }
            }

            return false;
        },

        addItem: function(data, isNewData) {
            let $template = $(window.Mustache.render(this.TEMPLATE_IMAGE, data)),
                $input = $template.find('.qor-file__input'),
                $item = $input.closest(CLASS_ITEM),
                $hiddenItem = this.$selectFeild.find('[data-primary-key="' + data.primaryKey + '"]'),
                maxItem = this.bottomsheetsData.maxItem,
                selectedItem = this.getSelectedItemData().selectedNum,
                cropOptions = data.MediaOption.CropOptions,
                needCropSize = this.compareCropSizes(data),
                selectedType = data.SelectedType,
                isSVG = /.svg$/.test(data.MediaOption.FileName),
                _this = this;

            if (!isNewData) {
                if (maxItem == 1) {
                    this.changeIcon(data.$clickElement, 'one');
                } else {
                    this.changeIcon(data.$clickElement, true);
                }
            }

            if (maxItem && selectedItem >= maxItem) {
                if (maxItem == 1) {
                    this.$selectFeild.find(CLASS_ITEM).remove();
                } else {
                    window.alert(this.bottomsheetsData.maxItemHint);
                    return;
                }
            }

            if ($hiddenItem.size()) {
                this.showHiddenItem($hiddenItem);
                if (maxItem == 1) {
                    setTimeout(function() {
                        _this.$bottomsheets.remove();
                    }, 1000);
                }
                return;
            }

            if (maxItem == 1) {
                this.$selectFeild.find(CLASS_ITEM).filter('.is_deleted').remove();
            }

            if (!isSVG) {
                if (selectedType === 'video') {
                    $template = $(window.Mustache.render(this.TEMPLATE_UPLOADEDVIDEO, data));
                } else if (selectedType === 'video_link') {
                    data.VideoLink = `//www.youtube.com/embed/${getYoutubeID(data.MediaOption.Video)}?rel=0&fs=0&modestbranding=1&disablekb=1`;
                    $template = $(window.Mustache.render(this.TEMPLATE_VIDEOLINK, data));
                } else if (selectedType === 'file') {
                    $template = $(window.Mustache.render(this.TEMPLATE_FILE, data));
                }
            }

            $template.data({
                description: data.MediaOption.Description,
                mediaData: data,
                videolink: selectedType === 'video_link' ? data.MediaOption.Video : ''
            });

            if (isSVG) {
                $template.addClass('is-svg').find('.qor-file__input').remove();
            }
            $template.appendTo(this.$selectFeild);

            // if image alread have CropOptions, replace original images as [big,middle, small] images.
            if (cropOptions && selectedType === 'image') {
                this.resetImages(data, $template);
            }

            // trigger cropper function for new item
            if (selectedType === 'image') {
                $template.find(CLASS_CROPPER_OPTIONS).val(JSON.stringify(data.MediaOption));
            }

            $template.trigger('enable');

            // if not have crop options or have crop options but have anothre size name to crop
            if ((!cropOptions || needCropSize) && $input.data('qor.cropper') && !isSVG) {
                $input.data('qor.cropper').load(data.MediaOption.URL, function() {
                    _this.syncImageCrop($item, _this.resetImages);
                });
            }

            if (isNewData || maxItem == 1) {
                setTimeout(function() {
                    _this.$bottomsheets.remove();
                }, 150);
            }
        },

        resetImages: function(data, $template) {
            var cropOptions = data.MediaOption.CropOptions,
                keys = Object.keys(cropOptions),
                url = data.MediaOption.OriginalURL;

            if (!/original/.test(url)) {
                return;
            }

            for (var i = keys.length - 1; i >= 0; i--) {
                cropOptions[keys[i]]['URL'] = url.replace(/original/, keys[i]);
            }

            $template.find('img').each(function() {
                var $this = $(this),
                    sizeName = $this.data().sizeName;

                if (sizeName && sizeName != 'original' && cropOptions[sizeName]) {
                    $this.prop('src', cropOptions[sizeName]['URL']);
                }
            });
        },

        handleSelectMany: function($bottomsheets) {
            let options = {
                onSelect: this.onSelectResults.bind(this), // render selected item after click item lists
                onSubmit: this.onSubmitResults.bind(this) // render new items after new item form submitted
            };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MEDIABOX);
            this.$bottomsheets = $bottomsheets;
            this.initMedia();
        },

        onSelectResults: function(data) {
            this.handleResults(data);
        },

        onSubmitResults: function(data) {
            this.handleResults(data, true);
        },

        handleResults: function(data, isNewData) {
            if (isNewData) {
                data.MediaOption = JSON.parse(data.MediaOption);
                this.handleResultsData(data, isNewData);
            } else {
                this.handleResultsData(data);
            }
        },

        handleResultsData: function(data, isNewData) {
            var $element = data.$clickElement,
                isSelected;

            if (!data.mediaLibraryUrl && !isNewData) {
                data.mediaLibraryUrl = data.url;
            }

            if (isNewData) {
                data.mediaLibraryUrl = this.bottomsheetsData.mediaboxUrl + '/' + data.primaryKey;
                this.addItem(data, isNewData);
                this.updateDatas(data);
                return;
            }

            $element.toggleClass(CLASS_SELECTED);
            isSelected = $element.hasClass(CLASS_SELECTED);

            if (isSelected) {
                this.addItem(data);
            } else {
                this.removeItem(data);
            }
            this.updateDatas(data);
        },

        updateDatas: function(data) {
            if (this.bottomsheetsData.maxItem != '1') {
                this.updateHint(this.getSelectedItemData());
            }
            this.updateMediaLibraryData(null, data);
        }
    };

    QorMediaBox.plugin = function(options) {
        return this.each(function() {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorMediaBox(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function() {
        var selector = '[data-toggle="qor.mediabox"]';
        $(document)
            .on(EVENT_DISABLE, function(e) {
                QorMediaBox.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function(e) {
                QorMediaBox.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorMediaBox;
});
