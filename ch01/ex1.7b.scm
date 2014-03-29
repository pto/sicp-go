(define (sqrt x)
  (sqrt-iter 1.0 0.0 x))

(define (sqrt-iter guess prev-guess x)
  (if (good-enough? guess prev-guess)
	guess
	(sqrt-iter (improve-guess guess x) guess x)))

(define (good-enough? guess prev-guess)
  (< (/ (abs (- guess prev-guess))
		guess)
	 1e-15))

(define (improve-guess guess x)
  (average guess (/ x guess)))

(define (average a b)
  (/ (+ a b) 2))

(sqrt 2)
(sqrt 2e-50)
(sqrt 2e50)
