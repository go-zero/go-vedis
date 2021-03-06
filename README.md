go-vedis [![Build Status](https://travis-ci.org/go-zero/go-vedis.svg?branch=master)](https://travis-ci.org/go-zero/go-vedis) [![GoDoc](https://godoc.org/github.com/go-zero/go-vedis?status.png)](https://godoc.org/github.com/go-zero/go-vedis) [![Coverage Status](https://coveralls.io/repos/go-zero/go-vedis/badge.svg?branch=master&service=github)](https://coveralls.io/github/go-zero/go-vedis?branch=master) [![Go Report Card](http://goreportcard.com/badge/go-zero/go-vedis)](http://goreportcard.com/report/go-zero/go-vedis)
========

go-vedis it's a Go bind to the [Vedis](http://vedis.symisc.net) datastore.

About Vedis
-----------

Vedis is an embeddable datastore C library built with over 70 [commands](http://vedis.symisc.net/commands.html) similar in concept to [Redis](http://redis.io) but without the networking layer since Vedis run in the same process of the host application.

Unlike most other datastores (i.e. memcache, Redis), Vedis does not have a separate server process. Vedis reads and writes directly to ordinary disk files. A complete database with multiple collections, is contained in a [single disk file](http://vedis.symisc.net/features.html#single_file). The database file format is cross-platform, you can freely copy a database between 32-bit and 64-bit systems or between [big-endian](http://en.wikipedia.org/wiki/Endianness) and [little-endian](http://en.wikipedia.org/wiki/Endianness) architectures.

Installation
------------

This package can be installed with the go get command:

    go get github.com/go-zero/go-vedis

Documentation
-------------

API documentation can be found here: http://godoc.org/github.com/go-zero/go-vedis

License
-------

MIT: https://github.com/go-zero/go-vedis/blob/master/LICENSE

Vedis License: https://github.com/symisc/vedis/blob/master/license.txt

Author
------

Jairo Luiz (a.k.a [TangZero](https://github.com/tangzero), a.k.a [go-zero](https://github.com/go-zero))
