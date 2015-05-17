<center>
# simple
</center>

 (c) David Rook 2012 - released under Simplified BSD 2-clause License

## OVERVIEW

This project is a collection of simple go programs.  Each one
should compile and run with Go 1.  If you're still using r60.3 or earlier 
then it might or
might not work without conversion.  They were useful to me when learning go and
I still look at them when I need to refresh my memory on how certain things work.
If you have go installed they should build without problems.  If you don't then
they can still be copied into the go "playground" at golang.org for experimentation.

The programs in alphabetic order:

* after.go
* april15c.go
	* allocating very large blocks of RAM
* bigInt.go rho.go
	* calculate common divisors of large numbers
* binaryEncode.go
	* empty interface array
* bugBounce.go
	* interfaces
* byteOrder go snip by Rob Pike
* call.go, call_test.go
	* build & run vs test
* cgo.go
	* currently failing - what changed? or was it wrong to start?
	* demonstrate how to call a C program from go code
* closure.go, closure2.go
	* functions that return functions
* debug.go
	* variadic parameters
	* attaching Printf methods to boolean variables
* env.go, env_test.go
* eval.go
	* variable is a function
* file.go
    * read and write to text files
* fileinfo.go
	* build flags 
	* syscall to get the "other" info about a file
* filter.go, check.go
	* gather args from keyboard 
* freeport.go
	* get OS to give you a free port number for server
	* String() method
* goroutine.go
	* channel as throttle to limit resource consumption (CPU here) 
	* waitgroup
* treewalk.go


### Installation

Since there are multiple subdirectories the easiest installation is 

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/simple.git```

### Features

* Simple, usually just a single topic being demonstrated
* The typing has already been done :-)
 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for program] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/simple "github.com/hotei/simple"

Comments can be sent to <hotei1352@gmail.com> or to "hotei" at github.com.


License
-------
The programs contained within the "hotei/simple" github project are released with the 
following Simplified BSD License:

> Copyright (c) 2012-2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2015 David Rook 

// EOF README-simple.md  (this is a markdown document and tested OK with blackfriday)