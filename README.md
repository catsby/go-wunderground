go-wunderground
==============

A simple library/command line tool for working with [Wunderground API][1].

## Usage

This is packaged as a library (`go-wunderground`) and CLI tool (`wunderground`). 

### Using the Library

*Don't*

### Using the CLI

Install the CLI with standard `go get`

```console
$ go get github.com/catsby/go-wunderground/cmd/wunderground
```

You need to [register at Wunderground.com/api and get you're own API key][1].
Once you have that, it's probably easiest to put in your `~/.bashrc` or
~/.zshrc` or whatever you're using... 

```console
$ echo "export WUNDERGROUND_API_KEY=youruniqueandawesomekey" >> .zshrc
```

Now you can use the CLI:

```console
$ wunderground 65203
	Getting weather for 65203

Forcast Days:
    Wednesday 0
      Clear. Lows overnight in the low 70s.
      [...]
```

Awesome (っˆーˆ)っ 

[1]: http://www.wunderground.com/api
