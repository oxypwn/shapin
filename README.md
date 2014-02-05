# Shapin

Take input on stdin and print a sha512 to stdout.

[![Build Status](https://travis-ci.org/pandrew/shapin.png?branch=master)](https://travis-ci.org/pandrew/shapin)

# Notice!

There is an upcoming feature in metasploits meterpreter watching clipboard and being able to retreieve data from it. This means that any software using the clipboard will be compromised when using the module against it. I have not seen this module being tested against osx yet.


## Features
+   Leaves no traces of input in your shells history.
+   Leaves no traces of input in ps
+   Clears pastebuffer and shell after 10 seconds.

## system requirements
OSX 10.9 for (pbcopy to clear clipboard)
