goroutine
=========

GoRoutine is a test to make sure you're running on all cylinders of your 
multi-core CPU.

Installation
------------

	Edit BIG_NUMBER	in main() - to finish faster reduce it
	Edit sress in cpu_hog() - true will push CPU, false it loafs along
	Edit loops_to_do in main() - set to more than NUM_CPUS to evoke throttle
	
	go build

Features
--------

*	Figures out how many CPUs you have

*	[optionally] Pushes them hard with cpu_hog().  Watch on your CPU load graph.

*   About 70 lines of code.


Output
------

	What to expect from a 6 Core system with stress=true: 
	$ time ./goroutine
	goroutine running, this may take a while...
	starting 6 goroutines in total
	goroutine[1] started
	goroutine[2] started
	cpu_hog called with n = 1, maxj = 81
	cpu_hog called with n = 2, maxj = 87
	goroutine[6] started
	cpu_hog called with n = 6, maxj = 47
	goroutine[3] started
	goroutine[4] started
	goroutine[5] started
	cpu_hog called with n = 4, maxj = 81
	cpu_hog called with n = 5, maxj = 18
	cpu_hog called with n = 3, maxj = 59
	cpu_hog exited with n = 5
	received(50)
	cpu_hog exited with n = 6
	received(60)
	cpu_hog exited with n = 3
	received(30)
	cpu_hog exited with n = 1
	received(10)
	cpu_hog exited with n = 4
	received(40)
	cpu_hog exited with n = 2
	received(20)
	goroutine finish
	
	real	0m28.340s
	user	2m2.620s
	sys	0m0.010s

	If you use $ xtime ./goroutine the last lines may be similar to :
	122.62u 0.00s 28.34r 4752kB ./goroutine

Discussion
----------
runtime.NumCPU() tells us how many cores.

We use a channel with one slot per core as a "throttle" mechanism.  We can call
the cpu_hog more times than that, but it will block waiting for the throttle
channel when all CPUs are in use.  

The goroutines are started by anonymous function so we inject the loop counter (i)
so that we can track them. 
Once the function compeletes we increment the throttle channel to indicate a CPU
is available again.

When the process winds down we have to insure we wait till all the goroutines 
we created in the loop have time to finish.  The wg.Wait() blocks until all the
goroutines still running have completed.

If you want to test synchronization but don't want to stress-test your CPU change
the stress variable in cpu_hog() to false.  This will substitute time.Sleep() with
the result that CPU time may be near zero :
    goroutine finish
    0.01u 0.00s 40.01r 5232kB ./goroutine

INFO: user time is the total of all cores.  Real time is wallclock.

License
-------

goroutine.go is distributed under the Simplified BSD License:

> Copyright (c) 2012 David Rook. All rights reserved.
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
