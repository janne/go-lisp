go-lisp.go
==========

A simple, embeddable LISP interpreter in Go.

Syntax
------

### (quote exp)

Return the exp literally; do not evaluate it.

    (quote (a b c))

Returns:

    (a b c)

### (if test conseq alt)

Evaluate test; if true, evaluate and return conseq; otherwise evaluate and
return alt.

    (if (< 10 20) (+ 1 1) (+ 3 3))

Returns:

    2

### (set! var exp)

Evaluate exp and assign that value to var, which must have been previously
defined (with a define or as a parameter to an enclosing procedure).

    (set! x2 (* x x))

### (define var exp)

Define a new variable and give it the value of evaluating the expression exp.

    (define r 3)
    (define square (lambda (x) (* x x)))

### (lambda (var...) exp)

Create a procedure with parameter(s) named var... and the expression as the
body.

    (lambda (r) (* 3.141592653 (* r r)))

### (begin exp...)

Evaluate each of the expressions in left-to-right order, and return the final
value.

    (begin (define x 1) (set! x (+ x 1)) (* x 2))

Returns:

    4

### (proc exp...)

If proc is anything other than one of the symbols if, set!, define, lambda,
begin, or quote then it is treated as a procedure. It is evaluated using the
same rules defined here. All the expressions are evaluated as well, and then
the procedure is called with the list of expressions as arguments.

    (square 12)

Returns:

    144

Inspired by: http://norvig.com/lispy.html
