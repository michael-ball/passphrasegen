# passphrasegen
A passphrase generator written in Go

## Usage
Install with by running `go install`, or build and run by running `go build`.

Generate a passphrase by running `passphrasegen`.

Optional arguments:
* -w number: number of words to use in the passphrase. Defaults to 4.
* Path to words file. Defaults to `/usr/share/dict/words`.

### Example

`passwordgen -w 5 /usr/share/dict/american-english`
