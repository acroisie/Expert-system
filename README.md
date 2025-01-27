# Expert-system

Priority : ! > ∧ > ∨ > ⊕ > → > ↔.

                ↔
              /   \
            →       F
          /   \
        ⊕       E
       /  \
      |    D
     / \
   AND  C
  /   \
!A     B


Expression   ::= Term ( ('+' | '|') Term )*
Term         ::= Factor ( '^' Factor )*
Factor       ::= '!'? Variable | '(' Expression ')'
Variable     ::= [A-Z]
