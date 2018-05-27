GOCMD=go
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

CABOCHAFILE=neko.txt.cabocha
RESULTFILE=result.txt

GETALLRANK=sort $(RESULTFILE) | uniq --count | sort --numeric-sort --reverse
GETDORANK=grep "する" $(RESULTFILE) | sort | uniq --count | sort --numeric-sort --reverse
GETSEERANK=grep "見る" $(RESULTFILE) | sort | uniq --count | sort --numeric-sort --reverse
GETGIVERANK=grep "与える" $(RESULTFILE) | sort | uniq --count | sort --numeric-sort --reverse

RM=rm -f

.PHONY:
	all test benchmark clean run sort
all: test run sort
test:
	$(GOTEST) -v
benchmark:
	$(GOTEST) -bench . --behchmem
clean:
	$(GOCLEAN)
	-$(RM) $(RESULTFILE)
run:
	$(GORUN) ./main.go -f $(CABOCHAFILE) -d $(RESULTFILE)
sort:
	$(GETALLRANK)
	$(GETDORANK)
	$(GETSEERANK)
	$(GETGIVERANK)
