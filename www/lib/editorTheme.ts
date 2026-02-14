import { EditorView } from '@codemirror/view'
import { HighlightStyle, syntaxHighlighting } from '@codemirror/language'
import { tags } from '@lezer/highlight'

const bg = '#0a0a0a'
const gutterBg = '#0a0a0a'
const gutterFg = '#525252'
const fg = '#d4d4d8'
const selection = '#dc262630'
const activeLine = '#ffffff06'
const activeLineGutter = '#171717'
const cursor = '#ef4444'
const border = '#262626'
const tooltip = '#171717'
const highlight = '#dc262620'

const gorThemeBase = EditorView.theme({
  '&': {
    color: fg,
    backgroundColor: bg,
  },
  '.cm-content': {
    caretColor: cursor,
  },
  '.cm-cursor, .cm-dropCursor': { borderLeftColor: cursor },
  '&.cm-focused > .cm-scroller > .cm-selectionLayer .cm-selectionBackground, .cm-selectionBackground, .cm-content ::selection': {
    backgroundColor: selection,
  },
  '.cm-panels': { backgroundColor: bg, color: fg },
  '.cm-panels.cm-panels-top': { borderBottom: `1px solid ${border}` },
  '.cm-panels.cm-panels-bottom': { borderTop: `1px solid ${border}` },
  '.cm-searchMatch': {
    backgroundColor: '#dc262640',
    outline: '1px solid #dc2626',
  },
  '.cm-searchMatch.cm-searchMatch-selected': {
    backgroundColor: '#dc262625',
  },
  '.cm-activeLine': { backgroundColor: activeLine },
  '.cm-selectionMatch': { backgroundColor: highlight },
  '&.cm-focused .cm-matchingBracket, &.cm-focused .cm-nonmatchingBracket': {
    backgroundColor: '#dc262635',
  },
  '.cm-gutters': {
    backgroundColor: gutterBg,
    color: gutterFg,
    border: 'none',
    borderRight: `1px solid ${border}`,
  },
  '.cm-activeLineGutter': {
    backgroundColor: activeLineGutter,
  },
  '.cm-foldPlaceholder': {
    backgroundColor: 'transparent',
    border: 'none',
    color: '#737373',
  },
  '.cm-tooltip': {
    border: `1px solid ${border}`,
    backgroundColor: tooltip,
  },
  '.cm-tooltip .cm-tooltip-arrow:before': {
    borderTopColor: 'transparent',
    borderBottomColor: 'transparent',
  },
  '.cm-tooltip .cm-tooltip-arrow:after': {
    borderTopColor: tooltip,
    borderBottomColor: tooltip,
  },
  '.cm-tooltip-autocomplete': {
    '& > ul > li[aria-selected]': {
      backgroundColor: '#262626',
      color: fg,
    },
  },
}, { dark: true })

const gorHighlightStyle = HighlightStyle.define([
  { tag: tags.keyword,
    color: '#f87171' },                    // red — keywords (fn, let, const, if, return)
  { tag: [tags.name, tags.deleted, tags.character, tags.propertyName, tags.macroName],
    color: '#e5e5e5' },                    // near-white — names
  { tag: [tags.function(tags.variableName), tags.labelName],
    color: '#60a5fa' },                    // blue — function names
  { tag: [tags.color, tags.constant(tags.name), tags.standard(tags.name)],
    color: '#fb923c' },                    // orange — constants
  { tag: [tags.definition(tags.name), tags.separator],
    color: '#d4d4d8' },                    // light gray — definitions
  { tag: [tags.typeName, tags.className, tags.number, tags.changed, tags.annotation, tags.modifier, tags.self, tags.namespace],
    color: '#fbbf24' },                    // amber — numbers, types
  { tag: [tags.operator, tags.operatorKeyword, tags.url, tags.escape, tags.regexp, tags.link, tags.special(tags.string)],
    color: '#a78bfa' },                    // violet — operators
  { tag: [tags.meta, tags.comment],
    color: '#525252' },                    // dim gray — comments
  { tag: tags.strong,
    fontWeight: 'bold' },
  { tag: tags.emphasis,
    fontStyle: 'italic' },
  { tag: tags.strikethrough,
    textDecoration: 'line-through' },
  { tag: tags.link,
    color: '#525252',
    textDecoration: 'underline' },
  { tag: tags.heading,
    fontWeight: 'bold',
    color: '#f87171' },
  { tag: [tags.atom, tags.bool, tags.special(tags.variableName)],
    color: '#fb923c' },                    // orange — booleans, builtins
  { tag: [tags.processingInstruction, tags.string, tags.inserted],
    color: '#4ade80' },                    // green — strings
  { tag: tags.invalid,
    color: '#ef4444' },
])

export const gorTheme = [gorThemeBase, syntaxHighlighting(gorHighlightStyle)]
