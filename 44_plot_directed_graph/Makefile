GOCMD=go
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

DOT=dot
DOTGIF=$(DOT) -Kdot -Tgif
CABOCHAFILE=neko.txt.cabocha
DOTFILE=digraph.dot
IMAGEFILE=sample.gif

RM=rm -f
WINRM=cmd.exe /C del

ifeq ($(OS),Windows_NT)
	RM=$(WINRM)
endif

.PHONY:
	all test benchmark clean run gif
all: test run gif
test:
	$(GOTEST) -v
benchmark:
	$(GOTEST) -bench . --behchmem
clean:
	$(GOCLEAN)
	-$(RM) $(DOTFILE)
	-$(RM) $(IMAGEFILE)
run:
	$(GORUN) ./main.go -f $(CABOCHAFILE) -n 25 --dot $(DOTFILE)
gif:
	$(DOTGIF) $(DOTFILE) -o $(IMAGEFILE)
