import { syntaxHighlighting, HighlightStyle } from "@codemirror/language"
import { EditorView } from "@codemirror/view"
import * as highlight from "@lezer/highlight"

const chalky = "#e5c07b", coral = "#e06c75", cyan = "#56b6c2", invalid = "#ffffff", ivory = "#abb2bf", stone = "#7d8799", // Brightened compared to original to increase contrast
malibu = "#61afef", sage = "#98c379", whiskey = "#d19a66", violet = "#c678dd", darkBackground = "#21252b", highlightBackground = "#2c313a", background = "#282c34", tooltipBackground = "#353a42", selection = "#3E4451", cursor = "#528bff";

const oneDarkTheme = EditorView.theme({
    "&": {
        color: ivory,
        backgroundColor: background
    },
    ".cm-content": {
        caretColor: cursor
    },
    ".cm-cursor, .cm-dropCursor": { borderLeftColor: cursor },
    "&.cm-focused > .cm-scroller > .cm-selectionLayer .cm-selectionBackground, .cm-selectionBackground, .cm-content ::selection": { backgroundColor: selection },
    ".cm-panels": { backgroundColor: darkBackground, color: ivory },
    ".cm-panels.cm-panels-top": { borderBottom: "2px solid black" },
    ".cm-panels.cm-panels-bottom": { borderTop: "2px solid black" },
    ".cm-searchMatch": {
        backgroundColor: "#72a1ff59",
        outline: "1px solid #457dff"
    },
    ".cm-searchMatch.cm-searchMatch-selected": {
        backgroundColor: "#6199ff2f"
    },
    ".cm-activeLine": { backgroundColor: "#6699ff0b" },
    ".cm-selectionMatch": { backgroundColor: "#aafe661a" },
    "&.cm-focused .cm-matchingBracket, &.cm-focused .cm-nonmatchingBracket": {
        backgroundColor: "#bad0f847"
    },
    ".cm-gutters": {
        backgroundColor: background,
        color: stone,
        border: "none"
    },
    ".cm-activeLineGutter": {
        backgroundColor: highlightBackground
    },
    ".cm-foldPlaceholder": {
        backgroundColor: "transparent",
        border: "none",
        color: "#ddd"
    },
    ".cm-tooltip": {
        border: "none",
        backgroundColor: tooltipBackground
    },
    ".cm-tooltip .cm-tooltip-arrow:before": {
        borderTopColor: "transparent",
        borderBottomColor: "transparent"
    },
    ".cm-tooltip .cm-tooltip-arrow:after": {
        borderTopColor: tooltipBackground,
        borderBottomColor: tooltipBackground
    },
    ".cm-tooltip-autocomplete": {
        "& > ul > li[aria-selected]": {
            backgroundColor: highlightBackground,
            color: sage
        }
    }
}, { dark: true });

const oneDarkHighlightStyle = HighlightStyle.define([
    { tag: highlight.tags.keyword,
        color: violet, fontWeight: "bold" },
    { tag: [highlight.tags.name, highlight.tags.deleted, highlight.tags.character, highlight.tags.propertyName, highlight.tags.macroName],
        color: coral },
    { tag: [highlight.tags.function(highlight.tags.variableName), highlight.tags.labelName],
        color: malibu },
    { tag: [highlight.tags.color, highlight.tags.constant(highlight.tags.name), highlight.tags.standard(highlight.tags.name)],
        color: whiskey },
    { tag: [highlight.tags.definition(highlight.tags.name), highlight.tags.separator],
        color: ivory },
    { tag: [highlight.tags.typeName], color: chalky, fontWeight: "bold" },
    { tag: [highlight.tags.className, highlight.tags.number, highlight.tags.changed, highlight.tags.annotation, highlight.tags.modifier, highlight.tags.self, highlight.tags.namespace],
        color: chalky },
    { tag: [highlight.tags.operator, highlight.tags.operatorKeyword, highlight.tags.url, highlight.tags.escape, highlight.tags.regexp, highlight.tags.link, highlight.tags.special(highlight.tags.string)],
        color: cyan },
    { tag: [highlight.tags.meta, highlight.tags.comment],
        color: stone },
    { tag: highlight.tags.strong,
        fontWeight: "bold" },
    { tag: highlight.tags.emphasis,
        fontStyle: "italic" },
    { tag: highlight.tags.strikethrough,
        textDecoration: "line-through" },
    { tag: highlight.tags.link,
        color: stone,
        textDecoration: "underline" },
    { tag: highlight.tags.heading,
        fontWeight: "bold",
        color: coral },
    { tag: [highlight.tags.atom, highlight.tags.bool, highlight.tags.special(highlight.tags.variableName)],
        color: whiskey },
    { tag: [highlight.tags.processingInstruction, highlight.tags.string, highlight.tags.inserted],
        color: sage },
    { tag: highlight.tags.invalid,
        color: invalid },
]);

export const oneDark = [oneDarkTheme, syntaxHighlighting(oneDarkHighlightStyle)];

const lightTheme = EditorView.theme({
    "&": {
        color: "#24292e",
        backgroundColor: "#fff"
    },
    "&.cm-editor .cm-scroller":{
        
    },
    ".cm-gutters":{
        backgroundColor: "#fff",
        color: "#6e7781",
        borderColor: "transparent"
    },
    "&.cm-focused .cm-selectionBackground, &.cm-focused .cm-line::selection, & .cm-selectionLayer .cm-selectionBackground, .cm-content ::selection":{
        backgroundColor: '#BBDFFF !important'
    },
    "& .cm-selectionMatch":{
        backgroundColor: '#BBDFFF'
    }
}, {dark : false})

const lightHighlightStyle = HighlightStyle.define([
    { tag: [highlight.tags.standard(highlight.tags.tagName), highlight.tags.tagName], color: '#116329' },
    { tag: [highlight.tags.bracket], color: '#6a737d' },
    { tag: [highlight.tags.comment], color: 'rgb(76, 200, 107)' },
    { tag: [highlight.tags.className, highlight.tags.propertyName], color: '#6f42c1' },
    { tag: [highlight.tags.variableName, highlight.tags.attributeName, highlight.tags.number], color: '#005cc5' },
    { tag: [highlight.tags.function(highlight.tags.variableName)], color: 'black' },
    { tag: [highlight.tags.typeOperator], color: '#ff8900' },
    { tag: [highlight.tags.typeName], color: '#ff8900', fontWeight: "bold" },
    { tag: [highlight.tags.keyword], color: '#d73a49',fontWeight: "bold" },
    { tag: [highlight.tags.operator], color: 'rgb(255, 0, 0)' },
    { tag: [highlight.tags.string, highlight.tags.meta, highlight.tags.regexp], color: 'blue' },
    { tag: [highlight.tags.special(highlight.tags.string)], color: 'rgb(197, 6, 11)'},
    { tag: [highlight.tags.name, highlight.tags.quote], color: '#22863a' },
    { tag: [highlight.tags.heading], color: '#24292e', fontWeight: 'bold' },
    { tag: [highlight.tags.emphasis], color: '#24292e', fontStyle: 'italic' },
    { tag: [highlight.tags.deleted], color: '#b31d28', backgroundColor: 'ffeef0' },
    { tag: [highlight.tags.atom, highlight.tags.bool, highlight.tags.special(highlight.tags.variableName)], color: '#e36209' },
    { tag: [highlight.tags.url, highlight.tags.escape, highlight.tags.regexp, highlight.tags.link], color: '#032f62' },
    { tag: highlight.tags.link, textDecoration: 'underline' },
    { tag: highlight.tags.strikethrough, textDecoration: 'line-through' },
    { tag: highlight.tags.invalid, color: '#cb2431' },
])

export const light = [lightTheme, syntaxHighlighting(lightHighlightStyle)];