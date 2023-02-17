## Problem 4: A Boolean logic interpreter.

Write a Boolean logic interpreter that can evaluate simple expressions, for example:
λ> T ∨ F
T

λ> T ∧ F
F

λ> (T ∧ F) = F
T
There should also be support for variables, such as the following:
λ> let X = F
X: F

λ> let Y = ¬X
Y: T

λ> ¬X ∧ Y
T
(Here the precedence of ¬ is higher than that of ∧.)

The exact syntax and scope is yours to decide on, but be sure to include support for arbitrary sequences of values ("true" and "false" and variables) combined using the operators AND, OR and NOT (respectively ∧, ∨, and ¬ in our example syntax) and parentheses.

Describe the syntax (and operator precedence rules) in your documentation with some examples.
