;; Incrementor w/ closure
(define inc (
  (lambda ()
    (define i 0)
    (lambda ()
      (set! i (+ i 1))
    )
  )
))
(display (quote Incrementor:) (inc))
(display (quote Incrementor:) (inc))
(display (quote Incrementor:) (inc))
(display (quote Incrementor:) (inc))
(display (quote Incrementor:) (inc))

;; Recursive counter
(define loop (lambda (i)
  (display (quote Loop:) i)
  (if (> i 1) (loop (- i 1)))
))
(loop 5)
