# PesPalDeveloperChallenge
 PesaPal Developer Challenge Solution
 
## The Pesapal Developer Challenge


We are looking for the most talented, young developers in the region. If you are a skilled programmer with a love for computing, Pesapal might be the place for you.

But beware and be-warned, our work is difficult. You will be expected to work at our pace, which is terrifying at times. You will be asked to come up with unorthodox solutions and solve problems beyond your ability. We will expect you to stretch and apply yourself, and above all, to think for yourself.



The Ideal Candidate.

Formal training in a relevant field (computer science, engineering, etc.) is not required. We do require, however, that

You love programming and computing

You are willing to work with others

You deliver within predictable timelines

You can write well.

The Problems.

We are doing a problem-oriented assessment to make our analysis as objective as possible. Choose (at least) one of the below problems and attempt to solve it. 

It is better to provide a thorough, elegant solution to one of the problems than it is to provide four or five poor solutions. Our aim in providing five is not to burden you with them, but to give you the flexibility to show your skills as best as you can. Also, we are not as interested in complete solutions (although these definitely help!) as we are in clear thinking and evidence of determination. For this reason we are willing to accept a well-documented attempt if you are unable to complete your implementation before the deadline.

You may work with any programming language or framework according to your taste and skill.

Use of libraries and existing programs is permitted, and even recommended. However, do not use a library that directly solves the stated problem: we will not be impressed with a build of Jekyll together with some shell scripts for the static-site generator problem, for instance.

Finally, do not underestimate the importance of documentation. Showing us that you can write well is perhaps the best way to show us that you can think well.


We define plagiarism as the attempt to pass off another's work as one's own. If you copy code or rely on something you find online (or offline, for that matter), simply attribute it and you will be safe. NB: If we judge that you have engaged in plagiarism you will be disqualified.

Use of ChatGPT and such tools is by no means prohibited, but you should indicate where and how you have relied on them. Also, keep in mind that everyone has access to these tools, so don't expect to impress us with five minutes of work!

## Problem 0: A computer.
Below is the instruction set for a machine:
0x00   halt -- Terminate program

0x01   nop  -- Do nothing

0x02   li   -- Load Immediate: li R1 0x00000000
               Load 0x00000000 into R1

0x03   lw   -- Load Word: lw R1 R2
               Load the contents of the memory location
               pointed to by R2 into R1

0x04   sw   -- Store Word: sw R1 R2
               Store the contents of R2 in the memory
               location pointed to by R1

0x05   add  -- Add: add R3 R1 R2
               Add R1 to R2 and store the result in R3

0x06   sub  -- Subtract: sub R3 R1 R2
               Subtract R2 from R1 and store the result in R3

0x07   mult -- Multiply: mult R3 R1 R2
               Multiply R1 by R2 and store the result in R3

0x08   div  -- Divide: div R3 R1 R2
               Divide R1 by R2 and store the result in R3

0x09   j    -- Unconditional Jump: j 0x00000000
               Jump to memory location 0x00000000

0x0A   jr   -- Unconditional Jump (Register): jr R1
               Jump to memory location stored in R1

0x0B   beq  -- Branch if Equal: bne R1 R2 R3
               Branch to memory location stored in R3
               if R1 and R2 are equal

0x0C   bne  -- Branch if Not Equal: beq R1 R2 R3
               Branch to memory location stored in R3
               if R1 and R2 are not equal

0x0D   inc  -- Increment Register: inc R1
               Increment R1

0x0E   dec  -- Decrement Register: dec R1
               Decrement R1
It has five registers and 64K of memory in a 32-bit address space, that is 0x00000000–0x0000FFFF. The five registers consist of three general purpose (R1, R2, R3); and two special purpose, a program counter (instruction pointer) register (PC), and a conditional register (COND) that stores conditional flags which provide information about the most recently executed calculation allowing programs to check logical conditions.

Each instruction is encoded in a half word (16 bits) in little endian. The first 4 bits (half byte) contain the instruction number, which can be anything from 0x0 to 0xE, while the second, third and fourth half-bytes (4 bit sections) contain register numbers. (Note: For some instructions like li the remainder of the 8 bits after the instruction number and the register number contain an 8 bit immediate value.)
Write an assembler for the instruction set that takes a text assembly program written for the above instruction set and produces the program as a set of 16-bit numbers. Basically, the task is to encode the text of the assembly into the bytecode format.

Write a simulator for the machine that will take the output of the assembler and execute it, correctly. For example, you should be able to run the following program:
; a simple counter program.
li R1 0x00000000
; end
li R2 0x0000FFFF
; memory location of loop start
li R3 loop
loop:
  ; store the contents of R1 at the memory location pointed by R1
  sw R1 R1
  ; increment the counter
  inc R1
  ; loop if the counter hasn't yet reached the end
  bne R1 R2 R3
  ; end program
  halt
Programs should be loaded from 0x0000CFFF to the end of memory so that any memory before that is usable by the programmer.
The simulator should log the register values after every cycle (upon execution of every instruction).

## Problem 1: A static-site generator.

Design and implement a simple static-site generator. 

It should be able to take a folder containing Markdown (or another non-HTML markup-type format) pages and produce a website. There should be support for a homepage, articles and supporting pages (e.g. an about page and some error pages).

## Problem 2: A diff and a patch.
The Unix tools diff and patch work in such a way that one can run a diff between file A and file B, and then use patch with the output of the diff and file A to produce file B.

Write a pair of programs, a diff and and a patch, which allow one to do this same operation, to compare two files and use the output and one of the files to produce the other file. Write them to work on the shell similarly to the POSIX manual descriptions (linked above), but you have freedom in terms of the algorithms used and the nature of the actual diff output. 

When your diff application is run on two files, it should be possible to use either file together with the diff output to produce the other. (Don't write a silly concatenating diff which simply concatenates the two files. We should be able to see that the diff output is actually the differences between the files.)

## Problem 3: A distributed system.

Build a TCP server that can accept and hold a maximum of N clients (where N is configurable).
These clients are assigned ranks based on first-come-first-serve, i.e whoever connects first receives the next available high rank. Ranks are from 0–N, 0 being the highest rank.

Clients can send to the server commands that the server distributes among the clients. Only a client with a lower rank can execute a command of a higher rank client. Higher rank clients cannot execute commands by lower rank clients, so these commands are rejected. The command execution can be as simple as the client printing to console that command has been executed.

If a client disconnects the server should re-adjust the ranks and promote any client that needs to be promoted not to leave any gaps in the ranks.

## Problem 4: A Boolean logic interpreter.

Write a Boolean logic interpreter that can evaluate simple expressions, for example:
λ> T ∨ F
T

λ> T ∧ F
F

λ> (T ∧ F) = F
T
There should also be support for variables, such as the following:
λ> let X = F
X: F

λ> let Y = ¬X
Y: T

λ> ¬X ∧ Y
T
(Here the precedence of ¬ is higher than that of ∧.)

The exact syntax and scope is yours to decide on, but be sure to include support for arbitrary sequences of values ("true" and "false" and variables) combined using the operators AND, OR and NOT (respectively ∧, ∨, and ¬ in our example syntax) and parentheses.

Describe the syntax (and operator precedence rules) in your documentation with some examples.

## Submitting.

To submit your application place it in a public repository (GitHub, GitLab, BitBucket, SourceHut etc.) and link to it in the application form on this page.
Be sure to document in README.md, as usual. For more information, see the following articles:

## About READMEs — GitHub Docs

Writing on GitHub — GitHub Docs

How to Write a Good README File for Your GitHub Project — Hillary Nyakundi

If any of the problem statements, or indeed anything in this job prompt, is unclear, feel free to reach us at jdev@pesapal.com.

Submit your application with corresponding solutions by 23:59:59 UTC+03:00 on 17th February 2023.
