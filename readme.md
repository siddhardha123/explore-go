### Go

 - static-ly typed
 - compiler throws an error if there is any unused variable.
 - fmt package is used for output 
 - how do we init a variable
    - var {nameOfVariable} {type} = {value}
    - we can skip the type when we initialize it with a value 
      i.e  var a = "hey", a is by default considered as string
    - we can even skip var by using shorthand
      a := "hey"
    - default values if not initialized 
       - int,float = 0
       - string = ''
       - bool  = false
    - const can be used for constants
- how do we write functions 
  - func name(){}
  - if func returns anything , type should be mentioned like
    func name() int {}
  - we can return multiple items
    func name() (int,int) {} 
   


