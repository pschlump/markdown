
# Markdown Specification 

This is a full grammar for what is implemented in this tool.
It is done as a **formal** grammar.

Markdown is a **line based** tool.  The special marker `^` is the beginning of a line, the `$` is the end of line.
Most of the tokens in markdown are context sensitive.

```
markdown ::= document

document ...




heading ::= ^ ['blank'|'tab']{0,4} #{1,6} ['blank'|'tab']? ['text']* $      // % header
    | ^ ['text']* $ ^ [=]{2,n}                                              // text =========       h1 header
    | ^ ['text']* $ ^ [-]{2,n}                                              // text ---------       h2 header
    | ^ ['text']* $ ^ [_]{2,n}                                              // text _________       h3 header
    | ^ ['text']* $ ^ [^]{2,n}                                              // text ^^^^^^^^^       h4 header
    | ^ ['text']* $ ^ [~]{2,n}                                              // text ~~~~~~~~~       h5 header
    ;
    
```
