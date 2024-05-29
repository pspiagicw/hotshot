# Bugs

- Write `(car { { "hotdogs" } "and" "pickle" "relish")` into the REPL and watch the program hang.
- Write `((lambda (x y) (+ x y)) 2 3)` and find out that we don't support inline lambda's.
- [x] Try import something (even invalid) and then invoke a identifier that is not declared. 
