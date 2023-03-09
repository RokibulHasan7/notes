# Makefile

- Makefiles are used to help decide which parts of a large program need to be recompiled.
- The goal of Makefiles is to compile whatever files need to be compiled, based on what files have changed.
- When files in interpreted languages change, nothing needs to get recompiled. When the program runs, the most recent 
  version of the file is used.
- ```$@``` is an ```automatic variable``` that contains the target name.
- ```*``` searches filesystem for matching filenames. It should be wrapped in ```wildcard``` function, otherwise it may
  fall into a common pitfall.
- ```*``` may be used in the target, prerequisites, or in the ```wildcard``` function.
- ```*``` may not be directly used in a variable definitions.
- When ```*``` matches no files, it is left as it is (unless run in the wildcard function).
- ```
  thing_wrong := *.o # Don't do this! '*' will not get expanded
  thing_right := $(wildcard *.o)
  
  all: one two three four
  
  # Fails, because $(thing_wrong) is the string "*.o"
  one: $(thing_wrong)
  
  # Stays as *.o if there are no files that match this pattern :(
  two: *.o
  
  # Works as you would expect! In this case, it does nothing.
  three: $(thing_right)
  
  # Same as rule three
  four: $(wildcard *.o)
  ```
- When ```%``` used in "matching" mode, it matches one or more characters in a string. This match is called the ```stem```.
- When ```%``` used in "replacing" mode, it takes the stem that was matched and replaces that in a string.
- ```%``` is most often used in rule definitions and in some specific functions.
- Automatic variables:
  - ```$@``` prints the target name.
  - ```$?``` output all prerequisites newer than the target.
  - ```$^``` outputs all prerequisites.
- By default, Makefile targets are "file targets" - they are used to build files from other files. Make assumes its target
  is a file, and this makes writing Makefiles relatively easy:
  ```
  foo: bar
    create_one_from_the_other foo bar
  ```
  However, sometimes you want your Makefile to run commands that do not represent physical files in the file system. Good
  examples for this are the common targets "clean" and "all". Chances are this isn't the case, but you may potentially have
  a file named ```clean``` in your main directory. In such a case Make will be confused because by default the ```clean```
  target would be associated with this file and Make will only run it when the file doesn't appear to be up-to-date with
  regards to its dependencies. <br>
  These special targets are called ```phony``` and you can explicitly tell Make they are not associated with files, e.g.:
  ```
  .PHONY: clean
  clean:
    rm -rf *.o
  ```
  Now ```make clean``` will run as expected even if you do have a file named ```clean```.<br>
  In terms of Make, a phony target is simply a target that is always out-of-date, so whenever you ask ```make <phony_target>```,
  it will run, independent from the state of the file system. Some common ```make``` targets that are often phony are:
  ```all```,```install```, ```clean```,```distclean```,```TAGS```,```info```,```check```.
- 