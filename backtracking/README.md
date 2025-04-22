# Backtracking

- A recursive technique
- Solution incrementally built
- Iteratively checks if base condition reached
- If exhaustively ran, it terminates with failure

## Implementation
### Steps:
- Prune invalid approaches
    - check early if path violates certain constraints, if so eliminate it
- Generate a partial solution, by iterating through available alternatives
- Check validity of selected alternative
- Check base cases for solution completion
- Backtrack if:
    - Pruning
    - Exhausted all options from a decision point

### Functions
- After determining backtracking strategy implementation usually starts with two main functions
    - driver func
    - recursive func
#### Recursive Function
- receives current state as input
- checks base cases
- performs a recursive call
- communicates solution back to drive via returns, shared vars

#### Driver function
- initializes the necessary data structure or inputs
- makes initial call to the recursive function
- interprets the outcome
- returns or prints final result

## Complexity
- negative effect on performance
- does not guarantee an optimal solution

## Application
- Board games to select next move
- Graphs and trees through the use of a Depth First Search
- Object detection in image processing