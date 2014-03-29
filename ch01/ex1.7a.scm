(define (sqrt x)
  (sqrt-iter 1.0 x))

(define (sqrt-iter guess x)
  (if (good-enough? guess x)
	guess
	(sqrt-iter (improve-guess guess x) x)))

(define (good-enough? guess x)
  (< (abs (- (square guess) x))
	 1e-15))

(define (improve-guess guess x)
  (average guess (/ x guess)))

(define (average a b)
  (/ (+ a b) 2))

(sqrt 2)
(sqrt 2e-50)
(sqrt 2e50)
