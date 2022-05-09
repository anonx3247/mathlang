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

## Examples

Here are a few examples of the benefits of this system:

### No more backslashes

```
forall x in A, exists y in B
```

```latex
\forall x \in A, \exists y \in B
````

```
sin^2 theta + cos^2 theta = 1
```

```latex
\sin^2 \theta + \cos^2 \theta = 1
```
### Intelligent, elegant fractions

```
5/3 , 2 + {2 - 6t}/{3+8m}, {-3 + 5}/2c+6 a+b/c
```
  

```latex
\frac{5}{3} , 2 + \frac{2 - 6t}{3+8m}, \frac{-3 + 5}{2c+6} \frac{a+b}{c}
```

They even handle fractions of fractions:

```
varphi = 1 + 1/{1 + 1/{1 + 1/{1 + ...}}}
```

```latex
\varphi = 1 + \frac{1}{1 + \frac{1}{1 + \frac{1}{1 + \cdots}}}
```

### Smart annotations

Wether its vectors, conjugates, or hats, mathlang has better syntax!

```
a^{^}, b^{_}, c^{->}, d^{~}
```

```latex
\hat{a}, \overline{b}, \overrightarrow{c}, \tilde{d}
```

### Automatically takes the best choices

Parenthesis are already chosen to be the size-adjusting type, wether you need them or not:

```
(a+b)
```
  
```latex
\left(a+b\right)
```

### Matrices, made easy

```
A = &{a,b,c;d,e,f;g,h,i}
```

```latex
A = \begin{pmatrix} a & b & c \\ d & e & f \\ g & h & i \end{pmatrix}
```

It can also be good for other formatting:

```
(a+b)^n = sum_{k=0}^n &{n;k} a^k b^{n-k}
```

```latex
\left(a+b\right)^n = \sum_{k=0}^n \begin{pmatrix} n \\ k \end{pmatrix} a^k b^{n-k}
```