$(function() {
    let html = `<div id="dialog" style="display: none;">
                  <div class="mdl-dialog-bg"></div>
                  <div class="mdl-dialog">
                      <div class="mdl-dialog__content">
                        <p><i class="material-icons">warning</i></p>
                        <p class="mdl-dialog__message dialog-message">
                        </p>
                      </div>
                      <div class="mdl-dialog__actions">
                        <button type="button" class="mdl-button mdl-button--raised mdl-button--colored dialog-ok dialog-button" data-type="confirm">
                          ok
                        </button>
                        <button type="button" class="mdl-button dialog-cancel dialog-button" data-type="">
                          cancel
                        </button>
                      </div>
                    </div>
                </div>`,
        _ = window._,
        QOR = window.QOR,
        $dialog = $(html).appendTo('body');

    // ************************************ Refactor window.confirm ************************************
    $(document)
        .on('keyup.qor.confirm', function(e) {
            if (e.which === 27) {
                if ($dialog.is(':visible')) {
                    setTimeout(function() {
                        $dialog.hide();
                    }, 100);
                }
            }
        })
        .on('click.qor.confirm', '.dialog-button', function() {
            let value = $(this).data('type'),
                callback = QOR.qorConfirmCallback;

            $.isFunction(callback) && callback(value);
            $dialog.hide();
            QOR.qorConfirmCallback = undefined;
            return false;
        });

    QOR.qorConfirm = function(data, callback) {
        let okBtn = $dialog.find('.dialog-ok'),
            cancelBtn = $dialog.find('.dialog-cancel');

        if (_.isString(data)) {
            $dialog.find('.dialog-message').text(data);
            okBtn.text('ok');
            cancelBtn.text('cancel');
        } else if (_.isObject(data)) {
            if (data.confirmOk && data.confirmCancel) {
                okBtn.text(data.confirmOk);
                cancelBtn.text(data.confirmCancel);
            } else {
                okBtn.text('ok');
                cancelBtn.text('cancel');
            }

            $dialog.find('.dialog-message').text(data.confirm);
        }

        $dialog.show();
        QOR.qorConfirmCallback = callback;
        return false;
    };

    // *******************************************************************************

    // ****************Handle download file from AJAX POST****************************
    let objectToFormData = function(obj, form) {
        let formdata = form || new FormData(),
            key;

        for (var variable in obj) {
            if (obj.hasOwnProperty(variable) && obj[variable]) {
                key = variable;
            }

            if (obj[variable] instanceof Date) {
                formdata.append(key, obj[variable].toISOString());
            } else if (typeof obj[variable] === 'object' && !(obj[variable] instanceof File)) {
                objectToFormData(obj[variable], formdata);
            } else {
                formdata.append(key, obj[variable]);
            }
        }

        return formdata;
    };

    QOR.qorAjaxHandleFile = function(url, contentType, fileName, data) {
        let request = new XMLHttpRequest();

        request.responseType = 'arraybuffer';
        request.open('POST', url, true);
        request.onload = function() {
            if (this.status === 200) {
                let blob = new Blob([this.response], {
                        type: contentType
                    }),
                    url = window.URL.createObjectURL(blob),
                    a = document.createElement('a');

                document.body.appendChild(a);
                a.href = url;
                a.download = fileName || 'download-' + $.now();
                a.click();
            } else {
                window.alert('server error, please try again!');
            }
        };

        if (_.isObject(data)) {
            if (Object.prototype.toString.call(data) != '[object FormData]') {
                data = objectToFormData(data);
            }

            request.send(data);
        }
    };

    // ********************************convert video link********************
    // linkyoutube: /https?:\/\/(?:[0-9A-Z-]+\.)?(?:youtu\.be\/|youtube\.com\S*[^\w\-\s])([\w\-]{11})(?=[^\w\-]|$)(?![?=&+%\w.\-]*(?:['"][^<>]*>|<\/a>))[?=&+%\w.-]*/ig,
    // linkvimeo: /https?:\/\/(www\.)?vimeo.com\/(\d+)($|\/)/,

    let converVideoLinks = function() {
        let $ele = $('.qor-linkify-object'),
            linkyoutube = /https?:\/\/(?:[0-9A-Z-]+\.)?(?:youtu\.be\/|youtube\.com\S*[^\w\-\s])([\w\-]{11})(?=[^\w\-]|$)(?![?=&+%\w.\-]*(?:['"][^<>]*>|<\/a>))[?=&+%\w.-]*/gi;

        if (!$ele.length) {
            return;
        }

        $ele.each(function() {
            let url = $(this).data('video-link');
            if (url.match(linkyoutube)) {
                $(this).html(`<iframe width="100%" height="100%" src="//www.youtube.com/embed/${url.replace(linkyoutube, '$1')}" frameborder="0" allowfullscreen></iframe>`);
            }
        });
    };

    $.fn.qorSliderAfterShow.converVideoLinks = converVideoLinks;
    converVideoLinks();
});
