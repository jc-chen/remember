# remember
command line todo list written in go

## Installation
### Compile from source
```
$ go get github.com/happyyi008/remember
$ ./setup.sh
```

or

### Run the release

If you don't want to setup Go you can download the binary and copy it into your `/usr/local/bin/`.
For Windows: release/remember.exe
For MacOS: release/rmb

The current build is only for MacOS and Windows

## Usage:
```
$ rmb -help | -h		# prints usage
$ rmb ls                        # print your list of todos
$ rmb rm <index>                # removes the todo at <index> from your list
$ rmb Get this to compile       # adds a new todo to your list
```
## Contribute
Please send in pull requests, or create an issue for feature requests.

Or email me at happyyi008@gmail.com
