# go-purge-fastly

This is a utility that purges fastly cache using
https://github.com/sethvargo/go-fastly for backend work and
https://github.com/spf13/cobra for command line parameters.

## Building

```
# using go lang 1.7
go build -ldflags "-X=main.VERSION=$(VERSION)-$(BUILD)"" -o $(TARGET)
```

## Command Options

```
Usage:
  go-purge-fastly [command]

Available Commands:
  help        Help about any command
  purge       Purge instantly an individual URL.
  purgeall    Purges everything from a service
  purgekey    A brief description of your command
  version     Show go-purge-fastly client version

  Flags: (purgeall only)
      --service string   Service ID to purge

  Flags: (purgekey only)
      --surrkey string   Key to purge

Flags: (Global)
      --apikey string    Fastly API key, if not set uses env FASTLY_API_KEY) (default "")
      --config string    Config file (default is $HOME/.go-purge-fastly.yaml)
      --file string      Input file with url list to purge from fastly. (default "purge.txt")
      --sleep int        Amount of time to wait between purge requests (ms). (default 500)
  -s, --soft             Sends a soft purge request (default true)
```

## Running

```
# run the go-purge-fastly
go-purge-fastly version

# purge list of urls based on input file, wait 500 ms each request
go-purge-fastly purge --sleep 500 --file purge.txt

# purge a service and check freshness based on list of urls
go-purge-fastly purgeall --sleep 5000 --file purge.txt  --service string

# purge urls with key
go-purge-fastly purgekey --sleep 1000 --file purge.txt

```
