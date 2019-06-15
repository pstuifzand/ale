;;;; ale core: exceptions

(letfn [(fn is-call [sym clause]
          (and (local? sym)
               (list? clause)
               (eq sym (first clause))))

        (fn is-catch-binding [form]
          (and (vector? form)
               (= 2 (length form))
               (local? (form 1))))

        (fn is-catch [clause parsed]
          (and (is-call 'catch clause)
               (is-catch-binding (nth clause 1))
               (!seq? (:block parsed))))

        (fn is-finally [clause parsed]
          (and (is-call 'finally clause)
               (!seq? (:catch parsed))
               (!seq? (:block parsed))))

        (fn is-expr [clause parsed]
          (!or (is-call 'catch clause)
               (is-call 'finally clause)))

        (fn try-append [parsed keyword clause]
          (conj parsed [keyword (conj (keyword parsed) clause)]))

        (fn try-prepend [parsed keyword clause]
          (conj parsed [keyword (cons clause (keyword parsed))]))

        (fn try-parse [clauses]
          (unless (seq? clauses)
                  {:block '() :catch '() :finally []}
                  (let* [f (first clauses)
                         r (rest clauses)
                         p (try-parse r)]
                    (cond
                      (is-catch f p)   (try-prepend p :catch f)
                      (is-finally f p) (try-append p :finally f)
                      (is-expr f p)    (try-prepend p :block f)
                      :else            (raise "malformed try-catch-finally")))))

        (fn try-catch-predicate [pred err-sym]
          (let* [l (thread-to-list pred)
                 f (first l)
                 r (rest l)]
            (cons f (cons err-sym r))))

        (fn try-catch-branch [clauses err-sym]
          (assert-args
           (seq? clauses) "catch branch not paired")
          (lazy-seq
           (let* [clause (first clauses)
                  var    ((clause 1) 0)
                  expr   (rest (rest clause))]
             (cons (list 'ale/let
                         [var err-sym]
                         [#f (cons 'ale/do expr)])
                   (try-catch-clauses (rest clauses) err-sym)))))

        (fn try-catch-clauses [clauses err-sym]
          (lazy-seq
           (when (seq? clauses)
             (let* [clause (first clauses)
                    pred   ((clause 1) 1)]
               (cons (try-catch-predicate pred err-sym)
                     (try-catch-branch clauses err-sym))))))

        (fn try-body [clauses]
          `(lambda [] [#f (do ,@clauses)]))

        (fn try-catch [clauses]
          (let [err (gensym "err")]
            `(lambda [,err]
               (cond
                 ,@(apply list (try-catch-clauses clauses err))
                 :else [#t ,err]))))

        (fn try-catch-finally [parsed]
          (let [block   (:block parsed)
                recover (:catch parsed)
                cleanup (:finally parsed)]
            (cond
              (seq? cleanup)
              (let [first# (rest (first cleanup))
                    rest#  (conj parsed [:finally (rest cleanup)])]
                `(defer
                   (lambda [] ,(try-catch-finally rest#))
                   (lambda [] ,@first#)))

              (seq? recover)
              `(let [rec# (recover ,(try-body block) ,(try-catch recover))
                     err# (rec# 0)
                     res# (rec# 1)]
                 (if err# (raise res#) res#))

              (seq? block) `(do ,@block)

              :else        '())))]

  (defmacro try clauses
    (try-catch-finally (try-parse clauses))))