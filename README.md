# Bonus.ly Multi Points

Give out Bonus.ly bonuses, 1 point at a time.

It is a bit cheeky to do it, but you get more points as a "giver" this way :)

## Installation

    go get github.com/scottjbarr/bonusly-mp

## Releases

If you don't have [Go](http://golang.org) installed, you can download binaries
from the [releases](../../releases) page.

## Example Usage

You will need an access token to run this.

This example will send 1 point to a@b.com, 10 times.

    bonusly-mp -token abc -points 10 -email a@b.com -reason "for the lulz #wat"

User a@b.com receives 10 bonus points, and you gave 10 times. Everybody wins!

## References

- [The Bonus.ly API](https://bonus.ly/api)

## Licence

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENCE.md](LICENCE.md)
