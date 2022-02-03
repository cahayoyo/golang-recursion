# Recursion in Golang

## Recursion Uses Stack
```
All recursion uses stack
It will call function, and returning function 1 by 1 through the stack
```

## Types Of Recursion
**1. Tail Recursion**
  - All operation will performed at calling time only, no operation at returning time
  - Position in last statement
 
**2. Head Recursion**
  - Operation will performed at returning time
  - Position in first statement
 
**3. Tree Recursion**
  - Recursive Function calling itself more than one time
  - Time Complexity is O(n<sup>2</sup>)
  - Space Complexity is O(n)
 
**4. Indirect Recursion**
  - Calling another function that will lead into cycle

**5. Nested Recursion**
  - Parameter in function recursion will be recursion also
