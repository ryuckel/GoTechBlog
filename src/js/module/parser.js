'use strict';

// Remarkableを利用するための処理

const { Remarkable } = window.remarkable;

const md = new Remarkable({
  html: true,
  breaks: true,
  langPrefix: 'language-',
  highlight: function(str, lang) {
    return Prism.highlight(str, Prism.languages[lang], lang);
  }
});
