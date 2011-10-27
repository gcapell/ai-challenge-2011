include $(GOROOT)/src/Make.inc

TARG=MyBot
GOFILES=\
	ants.go\
	map.go\
	MyBot.go\
	io.go \
	debug.go \
	path.go \
	orders.go \
	point.go \

include $(GOROOT)/src/Make.cmd
