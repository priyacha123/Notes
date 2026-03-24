# deriving a new class from an existing class so that new class inherits all members (attributes + methods) of existing class is called as inheritance

#  old class = praent class, base class, existing class, super class

# new class = child class, sub class, derived class

### OBJECT CLASS = all classes in python are derived from built-in 'object' class


# Need of inheritance
- for code reusability
- when you have relations among classes 

# How constructor works in inheritance
- By default, constructor of parent class available to child class

# Constructor Overriding
- if child class has constructor then child's priority is high

# Types of inheritance (depending on number of child and parent classes involved)
- Single inheritance 
    - one parent and one classe
    - object -> parent -> child
- Multi-level inheritance 
    - parent class and child class further inherited into new class forming multiple levels
    - object -> parent -> child -> grand_child
- Hierarchical inheritance
    - one parent and multiple classes
    - object -> parent -> obj-1, obj-2 _ _ _ obj-n
    - the childs cannot access other classes props 
    - can only access parents props
- Multiple inheritance
    - Class is derived from multiple base classes
    - object -> parent1 , parent2 ->(carries props of both parent1 and parent2) child
    - parents can't access eact other props
    - when super() is used in parent class then it will move to that parent which is left to the present parent in child class parameter

- Hybrid inheritance
    - it conatins multiple type of inheritance
- Cyclic inheritance
    - 


# MRO
- MRO represents how properties (attr + methods) are searched in inheritance
- Applicable in all types of inheritance
- Rules
    1. Python First search in child class and then goes to parent class
    - priority is to child class
    2. MRO follows 'Depth irst Left to right approach'
    3. mro() method can be used for knowing mro of any class objects
    - mro() method gives correct order always 
    - manual cal can give wrong ans sometime

# DocString
- written before any statements
- can be printed using ( __doc__() )
- eg: fun add() {
      """ This is docstring """
      ''' This is docstring another way '''
      } 
      print(add.__doc__())
- syntax : """ xyz """ OR ''' xyz '''
- triple quotes 
- purposes
    - function: arguments and return values
    - class: attributes and methods
    - modules: list of classes and functions
    - package: list of modules with functionality

# Diff b/w docstrings and comments
- comments are ignorned but not docstrings
- mmemory is allocated for docstrings
- comment: description of code
  docstring: purpose of code


# Magic methods
- Proper definition: "a method that is called implicitly by python to execute a certain operation on a type, suvh as addition. Such methods have names starting and ending with double underscores"
- they are special methods that havae double underscores (dunder) on both sides of the method name
- naming convention: __methodname__
- no need to call explicitely. Python automatically calls them as a response to certain operations, such as instantiation (cretaing an object)
-  eg: __init__(), __add__(), __new__(), etc. 