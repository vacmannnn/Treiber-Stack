# Treiber stack
The [Treiber stack](https://en.wikipedia.org/wiki/Treiber_stack) algorithm is a scalable
lock-free stack utilizing the fine-grained concurrency primitive compare-and-swap.

Module implements this data structes with corresponding methods:

- `Push` value at the top of stack
- Remove (`Pop`) last value
- Get value from `top` of stack
  
All methods are concurrency-safety. Stack also implements the `Stringer()` interface, 
which shows how many elements are on stack.
