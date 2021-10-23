# Mathlang

#### By: Neosapien

This project aims to create a much lighter and more orthogonal syntax for mathematical writing on the computer
instead of trying to replace LaTeX as a *typesetting* system, mathlang aims only to change they *syntax*.

Thus mathlang converts its own syntax to pure math latex code, allowing you to write in something readable and comfortable
while always allowing you to change to a syntax understood by every math typesetting program wether it be Latex, Markdown, HTML,
or other.

## Projected Usage (currently only piping to stdin is working)

```
mathlang -e "expression"
```
This will output the "translation" of the expression into \LaTeX

```
... | mathlang
```
This will output the `STDIN` translated into \LaTeX
This can be very useful in certain editors for example, `kakoune`
by piping the content of a line and binding it, you can translate
mathlang to \LaTeX on the fly in one keystroke.

```
mathlang -rD "\[" -lD "\]"
```

This allows you to specify delimiters for mathlang code,
thus allowing you to pipe an entire **file** into the converter and only change things between
the delimiters.

This allows `mathlang` to be used as a preprocessor, e.g:
```
mathlang -D "$$" -d "$" -f "myfile.md" | pandoc -o myfile.pdf
```

Note here the options "-D" and "-d" without "r" or "l" are for delimiters which aren't different depending on orientation
with "-D" for display equation delimiters and "-d" for inline equation delimiters.

