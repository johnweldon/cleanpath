# cleanpath

Simple tool to clean up PATH like strings by removing redundant and
incorrect entries but keeping the relative order of the paths that
remain.

For example:

```sh
$ export FAKEPATH=/usr/bin:/usr/local/bin:/usr/bin:/usr/sbin:.
$ export FAKEPATH=$(cleanpath $FAKEPATH)
$ echo $FAKEPATH
/usr/bin:/usr/local/bin:/usr/sbin
```

Notice:
1) Redundant paths removed, but relative order of first occurence is maintained.
2) Relative (unsafe) paths removed.


[![Build Status](https://travis-ci.org/johnweldon/cleanpath.svg)](https://travis-ci.org/johnweldon/cleanpath)
[![GoDoc](https://godoc.org/github.com/johnweldon/cleanpath?status.svg)](https://godoc.org/github.com/johnweldon/cleanpath)
