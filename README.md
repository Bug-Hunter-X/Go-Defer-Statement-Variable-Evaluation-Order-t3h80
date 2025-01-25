# Go Defer Statement Variable Evaluation Order

This repository demonstrates a common, yet subtle, error in Go related to the evaluation order of variables within `defer` statements.

The `bug.go` file shows a simple program that unexpectedly prints 15 instead of 15, due to how `defer` statements are evaluated in Go. The evaluation of the expression `x + 10` is deferred until after the `defer` function call, but the value of `x` is captured at the time the `defer` statement is executed, not when it's finally invoked.

The `bugSolution.go` file provides a corrected version and explanation of how to avoid such behavior.