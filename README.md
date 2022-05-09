# Mathlang

#### By: Neosapien

This project aims to create a much lighter and more orthogonal syntax for mathematical writing on the computer
instead of trying to replace LaTeX as a *typesetting* system, mathlang aims only to change they *syntax*.

Thus mathlang converts its own syntax to pure math latex code, allowing you to write in something readable and comfortable
while always allowing you to change to a syntax understood by every math typesetting program wether it be Latex, Markdown, HTML,
or other.

## Installation

To test it out run:

```
make
```

To install run:

```
sudo make install
```

(Note: while testing, the binary has to be in the same directory as `syntax_regexp.json` to find it)

## Usage

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
mathlang -d
```

This allows you to specify delimiters for mathlang code,
thus allowing you to pipe an entire **file** into the converter and only change things between
the delimiters.

This allows `mathlang` to be used as a preprocessor, e.g:
```
mathlang -df "myfile.md" | pandoc -o myfile.pdf
```