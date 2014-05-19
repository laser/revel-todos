revel-todos
===========

Todo management application built with Revel

install steps
-------------

1. download and install go: http://golang.org/doc/install
2. follow this path to get Revel configured: http://revel.github.io/tutorial/gettingstarted.html
3. clone this repo (or move it from where you cloned it) to $GOPATH/src/github.com/laser/revel-todos
4. run: 'go get ./...'
5. run: 'revel run github.com/laser/revel-todos'
6. open the browser and hit http://localhost:9000

interesting things to note
--------------------------

1. demonstrates using gorp and transactions (ORM)
2. lots of view-helper stuff ("url" etc.)
