# Mathlang

#### By: Neosapien

This project aims to create a much lighter and more orthogonal syntax for mathematical writing on the computer
instead of trying to replace LaTeX as a *typesetting* system, mathlang aims only to change they *syntax*.

Thus mathlang converts its own syntax to pure math latex code, allowing you to write in something readable and comfortable
while always allowing you to change to a syntax understood by every math typesetting program wether it be Latex, Markdown, HTML,
or other.


## Brief Presentation

Here you can see a few examples of what mathlang is capable of:

### Vectors, hats, and column vectors
##### Mathlang
```
u^{->} : [{ i^{^}; j^{^}}] 
```
##### LaTeX
```latex
\overrightarrow{u} : \left[\begin{matrix} \hat{i}\\ \hat{j}\end{matrix}\right] 
```
### Sets, / alternate symbols
##### Mathlang
```
RR^{_} = RR cup \{+- inf\}
```
##### LaTeX
```latex
\overline{\mathbb{R}} = \mathbb{R} \cup \{\pm \infty\}
```
### Sums, fractions, matrices (`\binom`)
##### Mathlang
```
sum_{k=0}^n ({n;k}) ( alpha/beta )^{n-k} ( gamma/delta )^k = ( alpha/beta + gamma/delta )^n
```
##### LaTeX
```latex
\sum_{k=0}^n \left(\begin{matrix}n\\k\end{matrix}\right) \left( \frac{\alpha}{\beta}
\right)^{n-k} \left( \frac{\gamma}{\delta} \right)^k =
\left( \frac{\alpha}{\beta} + \frac{\gamma}{\delta} \right)^n
```
### Matrices
##### Mathlang
```
A = ({a,b,c;d,e,f;g,h,i})
```
##### LaTeX
```latex
A = \left(\begin{matrix}a & b & c\\d & e & f\\g & h & i\end{matrix}\right)
```

## Dependencies
This uses standard go libraries so it doesn't need particular dependencies, simply install go.

## Installation
```bash
git clone https://github.com/neosapien3247/mathlang
```
```bash
cd mathlang
```
Then to build run:
```bash
make
```
and to install to `/usr/local/bin` run
```bash
sudo make install
```

## Syntax
See the [wiki](https://github.com/neosapien3247/mathlang/wiki) for syntax

## Projected Usage (currently only piping to stdin is working)

```bash
mathlang -e "expression"
```
This will output the "translation" of the expression into \LaTeX

```bash
... | mathlang
```
This will output the `STDIN` translated into \LaTeX
This can be very useful in certain editors for example, `kakoune`
by piping the content of a line and binding it, you can translate
mathlang to \LaTeX on the fly in one keystroke.

```bash
mathlang -rD "\[" -lD "\]"
```

This allows you to specify delimiters for mathlang code,
thus allowing you to pipe an entire **file** into the converter and only change things between
the delimiters.

This allows `mathlang` to be used as a preprocessor, e.g:
```bash
mathlang -D "$$" -d "$" -f "myfile.md" | pandoc -o myfile.pdf
```

Note here the options "-D" and "-d" without "r" or "l" are for delimiters which aren't different depending on orientation
with "-D" for display equation delimiters and "-d" for inline equation delimiters.

