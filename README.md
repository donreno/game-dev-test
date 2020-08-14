# Game Dev Test
This is a test to develop a game in golang using [libsdl](http://www.libsdl.org/) and inspired by https://github.com/velovix/gaming-in-go examples.
The purpose of this is for my own learning and hoping that it brings some examples of golang development with TDD and fancy refactors.

## How to build
First of all you will need to install libsdl on your machine.

Here is how to do it on ubuntu

```bash
$ sudo apt-get install libsdl2-dev
```

After this you should be able to build (considering you have everything set up to develop on go 1.14)

First run tests

```bash
$ make test coverage
```

Then build

```bash
$ make build
```

This will build the binary and move it to build folder along with the assets 🍻!!!
