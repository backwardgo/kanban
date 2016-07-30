The first commit is "just enough code" to provide a minimal command line
interface with decent help and logging configuration.

```
levi@needle:~/backwardgo/src/github.com/backwardgo/kanban (master *%=)
$ kanban
2016/07/30 12:03:33
Usage:

	kanban [command]

Available commands are:

  httpd - starts the http server that provides access to our UI and internal API


levi@needle:~/backwardgo/src/github.com/backwardgo/kanban (master *%=)
$ kanban httpd
[kanban::httpd] 2016/07/30 12:03:37 listen and serve on port :
^C

levi@needle:~/backwardgo/src/github.com/backwardgo/kanban (master *%=)
$ PORT=8000 kanban httpd
[kanban::httpd] 2016/07/30 12:03:46 listen and serve on port :8000
^C

```
