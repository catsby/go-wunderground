go-wunderground
==============

A simple library/command line tool for working with [Wunderground API][1].

## Usage

This is packaged as a library (`go-wunderground`) and CLI tool (`wunderground`). 

### Using the Library

The library is fledgling and in flux. It currently supports the following:

- [Current conditions][2] (but barely)
- [Forecast][3](a little bit more than the above)

See [cmd/wunderground/main.go][4] for some usage.

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
$ wunderground -postal 65203
	Getting weather for 65203

Forcast Days:
    Wednesday 0
      Clear. Lows overnight in the low 70s.
      [...]
```

Awesome (っˆーˆ)っ 

[1]: http://www.wunderground.com/api
[2]: http://www.wunderground.com/weather/api/d/docs?d=data/conditions
[3]: http://www.wunderground.com/weather/api/d/docs?d=data/forecast
[4]: https://github.com/catsby/go-wunderground/blob/master/cmd/wunderground/main.go