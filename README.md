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
$ wunderground 65203
Getting weather for 65203...

Forcast for Columbia, MO

    Thursday
      Plentiful sunshine. High 68F. Winds NNE at 5 to 10 mph.

    Thursday Night
      Clear to partly cloudy. Low 51F. Winds NE at 5 to 10 mph.

    Friday
      Cloudy skies. High 64F. Winds NNE at 10 to 15 mph.

    Friday Night
      Cloudy. Low 43F. Winds N at 5 to 10 mph.

    Saturday
      Abundant sunshine. High 66F. Winds NE at 5 to 10 mph.

    Saturday Night
      Clear skies. Low around 45F. Winds light and variable.

    Sunday
      Partly cloudy. High 72F. Winds light and variable.

    Sunday Night
      A few clouds. Low 57F. Winds SE at 5 to 10 mph.
```

Awesome (っˆーˆ)っ 

[1]: http://www.wunderground.com/api
[2]: http://www.wunderground.com/weather/api/d/docs?d=data/conditions
[3]: http://www.wunderground.com/weather/api/d/docs?d=data/forecast
[4]: https://github.com/catsby/go-wunderground/blob/master/cmd/wunderground/main.go
