include $(GOROOT)/src/Make.inc

TARG=MyBot
GOFILES=\
	game.go\
	map.go\
	MyBot.go\
	io.go \
	debug.go \
	path.go \
	orders.go \
	point.go \
	assign.go \
	combat.go \
	pattern_set.go \
	math.go \
	intercept.go \
	path2.go \

include $(GOROOT)/src/Make.cmd

zip:
	rm -f MyBot.zip
	zip MyBot.zip $(GOFILES) Makefile
