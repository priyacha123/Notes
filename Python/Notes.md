Python is an Interpreted language (line by line)
random.random() -> here () is used becuz random (2nd one) is function

# Python's inner working
- When python is installed, then along with it a 'Python Virtual Machine' is also installed.
- Create a script or program -> abc.py file
1. Compile to byte-code (mostly hidden -> visible when import is used (__pycache__)) 
    - byte-code is low-level code and platform-independent
    - byte-code runs faster
    - '.pyc' -> compiled python (also called Frozen Binaries)
    - Works only for imported files and not the top level (main) files
    - '__pycache__' is the folder generated automatically when a file is imported -> it holds .pyc files which shows changes in the original files just like git 
        - It uses different Finding Algo (DiffFN) which compares the prev and new changes in the source code
        - Source chamge and Python vrerions -> eg: abc.cpython-312.pyc
            - here, 'cpython' is the standard python 
            - '312' is the 3.12 version  

# Python Virtual Machine (PVM)
- Code loop to iterate byte-code
- It is a runtime engine
- Also known as python interpreter

## Company related ques
- Byte code is NOT machine code
- Python specific interpretation
- Standard implementation ->  cpython 
- Regular implementation -> jython, IronPython, Stackless, PyPy

# Python Shell
- IDLE is shell in window OR direclty in cmd write 'python' to open python shell
- ctrl + D -> Terminate powershell
- ctril + C -> Suspend powershell
- if import xyz.py in shell and then updated code, so in shell updated code will not work, so make it work reload the file -> syntax code is:
> from importlib import reload

> reload(xyz)

# Immutable and Mutable
- List, Set Dict, bytearray, array are mutable data types
- Strings, int, float, boolean, tuples, frozen set, bytes are immutable i.e., no change
- The variables get the reference block which holds the values and if the variables values are changed then this variables get another reference block to hold the changed values and thus the prev value is not changed. The prev value then becomes garbage value and is deleted.

# Data types OR Object types
- Every data types in python is object
- Number: 124, 3.215, 3+4j, 0b111, Decimal(), Fraction()
- String: "", '', b'a/chd4c'
- List: [1,2,3, [7,5,6]], list(range(10))
- Tuple: (1,54, 'U', 'spma'), tuple('spam), namedtuple
- Dictionary: {'food': 'spam'}, dict(hours=10)
- Set: set('abc), {'a', 'b', 'f'}
- File: open('abc.txt'), open(path)
- Boolean: True, False
- None: None
- Functions, modules, classes
- Advance: Decorators, Generators, Iterators, MetaProgramming

    - The variables don't have any datatype but their value have.

- If l2 is a copy of l1 then lists will have same reference so changes in l1 list will show in l2 list too -> but is l2 is decalred again with same values as l1, then changes in l1 will not show in l2 becuz they will have diff reference no.

- if slicing is used (eg: l2 = l1[ : ]), then it makes copy of l1 and reference value is diff

- '==' checks values -> true
- 'is' checks memory reference 
    - m = n -> m is n -> true
    - m = n -> n = [1,2] -> m is n -> false

# String
- repr('chai') -> "'cahi'"
    - It is a built-in function that takes an object as an argument and returns its 'official' string representation
- str('chai') -> 'cahi'
    - It is a type conversion, that converts other datatypes to string datatype
- print('chai') -> cahi
    - it prints output
# Numbers
- Octal -> starts with '0o'
- Hexa -> starts with '0x'
- Binary -> starts with '0b'

# Sets and Dict
- An empty set is represented by 'set()'
- An empty dict is represented by '{}'

# Strings
- to print ( "" ) inside ( "" ) means -> ( " "" " ) -> then use \" -> this will print " as a string 
- to print rawstring use -> syntax: r"string or sentence"

# List
- pop() -> give o/p -> pop (means last) value
- remove() -> Gives o/p None
- list is in sequential order -> order matters

# Dict
- Order doesn't matter

# Tuple
- Immutable

# Decorators 
- Like a tollbooth
- It is a toll where if functions are executed then it'll have to go through decorators
