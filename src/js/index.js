'use strict';

// DOM Treeの構築完了後に処理開始
document.addEventListener('DOMContentLoaded', () => {
    // HTML要素を取得する
    const saveBtn = document.querySelector('.submit');
    const form    = document.forms.namedItem('submit-form');

    // CSRFトークンを取得する
    const csrfToken = document.getElementsByName('csrf')[0].content;

    // TODO: 送信方法はあとで変える
    // TODO: とりあえず簡単に送信できるようにする。
    saveBtn.addEventListener('click', event => {
        event.preventDefault();

        // フォームに入力された内容を取得する
        const formData = new FormData(form)
        let status;

        fetch('/post', {
            // とりあえずpostだけ
            method: 'POST',
            headers: { 'X-CSRF-Token': csrfToken },
            body: formData
        })
          .then(res => {
              status = res.status;
          })
          .then(body => {
              console.log(JSON.stringify(body))

              if (status === 200) {
                  window.location.href = '/';
              }
          })
          .catch(err => console.log(err))
    });
});