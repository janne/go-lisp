go-lisp.go
==========

A simple Lisp interpreter developed in Go.

Go is a perfect fit for building a Lisp interpreter, since it's portable,
garbage collected and fast. The result is a one file distribution, which makes
it easy to deploy. The garbage collection unloads that work from the
interpreter, using the Go already tuned GC.

Inspired by: [http://norvig.com/lispy.html]

Setup
-----

#### Install Go

[http://golang.org/doc/install]()

#### Make sure your GOPATH is setup correctly, something like

    $ export GOPATH=$HOME/go
    $ mkdir -p $GOPATH

#### Clone the repo

    $ go get github.com/janne/go-lisp

#### Run the tests

    $ go test github.com/janne/go-lisp/lisp

#### Run the example

    $ go-lisp -file $HOME/go/src/github.com/janne/go-lisp/example.lsp

#### Run it interactively

    $ go-lisp -i

Syntax
------

#### (quote exp)

Return the exp literally; do not evaluate it.

    (quote (a b c))

Returns:

    (a b c)

#### (if test conseq alt)

Evaluate test; if true, evaluate and return conseq; otherwise evaluate and
return alt.

    (if (< 10 20) (+ 1 1) (+ 3 3))

Returns:

    2

#### (set! var exp)

Evaluate exp and assign that value to var, which must have been previously
defined (with a define or as a parameter to an enclosing procedure).

    (set! x2 (* x x))

#### (define var exp)

Define a new variable and give it the value of evaluating the expression exp.

    (define r 3)
    (define square (lambda (x) (* x x)))

#### (lambda (var...) exp)

Create a procedure with parameter(s) named var... and the expression as the
body.

    (lambda (r) (* 3.141592653 (* r r)))

#### (begin exp...)

Evaluate each of the expressions in left-to-right order, and return the final
value.

    (begin (define x 1) (set! x (+ x 1)) (* x 2))

Returns:

    4

#### (proc exp...)

If proc is anything other than one of the symbols if, set!, define, lambda,
begin, or quote then it is treated as a procedure. It is evaluated using the
same rules defined here. All the expressions are evaluated as well, and then
the procedure is called with the list of expressions as arguments.

    (square 12)

Returns:

    144

License
-------
Licensed under [MIT license](https://github.com/janne/go-lisp/blob/master/LICENSE).
