;;;; ale core: basics

(declare *env* *args*)

(def *pos-inf* (/ 1.0 0.0))
(def *neg-inf* (/ -1.0 0.0))

;; syntax-quoting requires it
(def concat!
  (lambda colls
    (letrec [concat-inner
             (lambda [colls head]
               (if (is-empty colls)
                   (apply list head)
                   (let [f (first colls)
                         r (rest colls)]
                     (if (is-empty f)
                         (concat-inner r head)
                         (concat-inner (cons (rest f) r)
                                       (append head (first f)))))))]
      (concat-inner colls []))))

(def defmacro
  (letrec [defmacro
           (macro [name . forms]
             `(def ,name
                (letrec [,name (macro ,@forms)]
                  ,name)))]
    defmacro))

(defmacro assert-args
  ([] '())
  ([clause]
    (raise "assert-args clauses must be paired"))
  (clauses
    `(if ,(clauses 0)
         (assert-args ,@(rest (rest clauses)))
         (raise ,(clauses 1)))))

(defmacro fn
  [name . forms]
  (if (is-local name)
    `(letrec [,name (lambda ,@forms)] ,name)
    `(lambda ,name ,@forms)))

(defmacro defn
  [name . forms]
  `(def ,name (fn ,name ,@forms)))

(defmacro define body
  (let [f (first body)
        r (rest body)]
    (if (is-list f)
        (let [name (first f) args (rest f)]
          `(defn ,name ,(apply vector args) ,@r))
        `(def ,@body))))

(defmacro !eq
  [value . comps]
  `(not (eq ,value ,@comps)))

(define (is-even value)
  (= (mod value 2) 0))

(define (is-odd value)
  (= (mod value 2) 1))

(define (is-true value)
  (if value #t #f))

(define (is-false value)
  (if value #f #t))

(define (inc value)
  (+ value 1))

(define (dec value)
  (- value 1))

(define (no-op . _))

(define (identity value) value)

(define (constantly value)
  (lambda _ value))

(defmacro .
  [target method . args]
  `((get ,target ,method) ,@args))