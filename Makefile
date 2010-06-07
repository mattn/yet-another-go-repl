include $(GOROOT)/src/Make.$(GOARCH)

TARG = go-repl
GOFILES= go-repl.go

include $(GOROOT)/src/Make.cmd
