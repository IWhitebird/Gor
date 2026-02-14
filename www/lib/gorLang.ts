import { StreamLanguage } from '@codemirror/language'

export const gorLanguage = StreamLanguage.define({
  token(stream) {
    if (stream.match(/#.*/)) return 'comment'
    if (stream.match(/"(?:[^"\\]|\\.)*"/)) return 'string'
    if (stream.match(/\b(fn|let|const|if|else|for|return)\b/)) return 'keyword'
    if (stream.match(/\b(true|false|null)\b/)) return 'atom'
    if (stream.match(/\b(print|swap)\b/)) return 'variableName.special'
    if (stream.match(/\b\d+\b/)) return 'number'
    if (stream.match(/[+\-*/%]=?|&&|\|\||[!=<>]=?|[&|]/)) return 'operator'
    if (stream.match(/[a-zA-Z_]\w*/)) return 'variableName'
    if (stream.match(/[{}()\[\],.:;]/)) return 'punctuation'
    stream.next()
    return null
  },
  languageData: {
    commentTokens: { line: '#' },
  },
})
