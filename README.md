# Introduction

- Process Data, any piece of code that does something, they utilise data structures to accomplish their end goal
- Data Structures store the data
- Behind data structures there is specific types of memory allocation.
    -   Example: a list could be stored as a dynamic array (Python) or linked list (Java)

## Big O Notation
- Measures performance for different data structures and algorithms
- Example:
    - We have two cars driving next to each other, by time one car takes a parallel road which is slightly bent outwards. Over a period of time the two cars will be very far from each other. When it comes to algorithms, when working on small sets of data the difference will be negligible. However as the data grows that difference will be clearer.

### Time Complexity
- Array of 5 elements vs linked list of 5 elements
    - To fetch the last element in the array is easy given we have a fixed memory block
    - For linked list it is iterative 1->2->3->4->5 
- Array -> O(1) = constant time
- Linked List -> where n is the size of list O(n)
- Other classifications
    - n! is very slow!
    - n^2^ 
    - log n
#### Optimisations
- optimise fetching last element of a linked list
    - linked list has a pointer to the start
    - one optimisation is caching that element  