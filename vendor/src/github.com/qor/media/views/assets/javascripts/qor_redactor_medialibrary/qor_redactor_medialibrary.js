// Add media library button for redactor editor
// By Jason weng @theplant
//
$(function() {
    $.Redactor.prototype.medialibrary = function() {
        return {
            init: function() {
                var button = this.button.add('medialibrary', 'MediaLibrary');
                this.button.addCallback(button, this.medialibrary.addMedialibrary);
                this.button.setIcon(button, '<i class="material-icons">photo_library</i>');
                $(document).on('reload.qor.bottomsheets', '.qor-bottomsheets__mediabox', this.medialibrary.initItem);
            },

            addMedialibrary: function() {
                var $element = this.$element,
                    data = { selectModal: 'mediabox', maxItem: '1' },
                    mediaboxUrl = $element.data().redactorSettings.medialibraryUrl,
                    BottomSheets;

                this.medialibrary.BottomSheets = BottomSheets = $('body').data('qor.bottomsheets');
                data.url = mediaboxUrl;
                BottomSheets.open(data, this.medialibrary.handleMediaLibrary);
            },

            handleMediaLibrary: function($bottomsheets) {
                var options = {
                    onSelect: this.medialibrary.selectResults, // render selected item after click item lists
                    onSubmit: this.medialibrary.submitResults // render new items after new item form submitted
                };

                this.medialibrary.$bottomsheets = $bottomsheets;
                $bottomsheets.qorSelectCore(options).addClass('qor-bottomsheets__mediabox');
                this.medialibrary.initItem();
            },

            initItem: function() {
                var $trs = $('.qor-bottomsheets').find('tbody tr'),
                    $tr,
                    $img;

                $trs.each(function() {
                    $tr = $(this);
                    $img = $tr.find('.qor-table--ml-slideout p img').first();
                    $tr.find('.qor-table__actions').remove();
                    if ($img.length) {
                        $tr.find('.qor-table--medialibrary-item').css('background-image', 'url(' + $img.prop('src') + ')');
                        $img.parent().remove();
                    }
                });
            },

            selectResults: function(data) {
                this.medialibrary.handleResults(data);
            },

            submitResults: function(data) {
                this.medialibrary.handleResults(data, true);
            },

            handleResults: function(data, isNew) {
                isNew && (data.MediaOption = JSON.parse(data.MediaOption));
                var reVideo = /\.mp4$|\.m4p$|\.m4v$|\.m4v$|\.mov$|\.mpeg$|\.webm$|\.avi$|\.ogg$|\.ogv$/,
                    mediaOption = data.MediaOption;

                if (data.SelectedType == 'video_link' || mediaOption.Video || mediaOption.URL.match(reVideo)) {
                    this.medialibrary.insertVideo(data);
                } else {
                    this.medialibrary.insertImage(data);
                }

                this.medialibrary.$bottomsheets.remove();
            },

            insertVideo: function(data) {
                this.opts.mediaContainerClass = typeof this.opts.mediaContainerClass === 'undefined' ? 'qor-video-container' : this.opts.mediaContainerClass;

                var htmlCode,
                    videoLink,
                    iframeStart,
                    iframeEnd,
                    videoType,
                    $html,
                    youkuID,
                    callbackData = {},
                    mediaContainerClass = this.opts.mediaContainerClass,
                    reUrlYoutube = this.opts.regexps.linkyoutube,
                    reUrlVimeo = this.opts.regexps.linkvimeo,
                    reUrlYouku = /http?:\/\/(www\.)|(v\.)youku.com/,
                    reUrlYoukuID = /(\/id_)(\w+)/,
                    $currentTag = this.selection.$currentTag,
                    reVideo = /\.mp4$|\.m4p$|\.m4v$|\.m4v$|\.mov$|\.mpeg$|\.webm$|\.avi$|\.ogg$|\.ogv$/,
                    randomString = (Math.random() + 1).toString(36).substring(7),
                    videoIdentification = 'qor-video-' + randomString,
                    mediaOption = data.MediaOption,
                    description = mediaOption.Description;

                iframeStart = '<figure class="' + mediaContainerClass + '"><iframe title="' + description + '" width="100%" height="380px" src="';
                iframeEnd = '" frameborder="0" allowfullscreen></iframe><figcaption>' + description + '</figcaption></figure>';

                if (data.SelectedType == 'video_link') {
                    videoLink = mediaOption.Video;

                    if (videoLink.match(reUrlYoutube)) {
                        videoType = 'youtube';
                        htmlCode = videoLink.replace(reUrlYoutube, iframeStart + '//www.youtube.com/embed/$1' + iframeEnd);
                    }

                    if (videoLink.match(reUrlVimeo)) {
                        videoType = 'vimeo';
                        htmlCode = videoLink.replace(reUrlVimeo, iframeStart + '//player.vimeo.com/video/$2' + iframeEnd);
                    }

                    if (videoLink.match(reUrlYouku) && reUrlYoukuID.test(videoLink)) {
                        videoType = 'youku';
                        youkuID = videoLink.match(reUrlYoukuID)[2];
                        htmlCode = '<iframe width=100% height=400 src="http://player.youku.com/embed/' + youkuID + '" frameborder=0 "allowfullscreen"></iframe>';
                    }
                } else if (mediaOption.URL.match(reVideo)) {
                    videoType = 'uploadedVideo';
                    htmlCode =
                        '<figure class="' +
                        mediaContainerClass +
                        '"><div role="application"><video width="100%" title="' +
                        description +
                        '" aria-label="' +
                        description +
                        '" height="380px" controls="controls" aria-describedby="' +
                        videoIdentification +
                        '" tabindex="0"><source src="' +
                        mediaOption.URL +
                        '"></video></div><figcaption id="' +
                        videoIdentification +
                        '">' +
                        description +
                        '</figcaption></figure>';
                }

                if (!htmlCode) {
                    return;
                }

                $html = $(htmlCode).addClass(videoIdentification);

                if ($currentTag) {
                    $currentTag.after($html);
                } else {
                    this.$editor.prepend($html);
                }

                this.code.sync();

                // trigger insertedVideo.redactor event after inserted videos
                callbackData.type = videoType;
                callbackData.videoLink = videoLink || mediaOption.URL;
                callbackData.videoIdentification = videoIdentification;
                callbackData.description = description;
                callbackData.$editor = this.$editor;
                callbackData.$element = this.$element;

                this.$element.trigger('insertedVideo.redactor', [callbackData]);
            },

            insertImage: function(data) {
                var src,
                    $currentTag = this.selection.$currentTag,
                    $img = $('<img>'),
                    $figure = $('<' + this.opts.imageTag + '>'),
                    mediaOption = data.MediaOption,
                    callbackData = {};

                src = mediaOption.URL.replace(/image\..+\./, 'image.');

                $img.attr({
                    src: src,
                    alt: mediaOption.Description
                });
                $figure.append($img);

                if ($currentTag) {
                    $currentTag.after($figure);
                } else {
                    this.$editor.prepend($figure);
                }

                this.image.setEditable($img);
                this.code.sync();

                // trigger insertedVideo.redactor event after inserted images
                callbackData.description = mediaOption.Description;
                callbackData.$img = $figure;
                callbackData.$editor = this.$editor;
                callbackData.$element = this.$element;

                this.$element.trigger('insertedImage.redactor', [callbackData]);
            }
        };
    };
});
