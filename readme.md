# goctl customized

### use antlr 4.13

#### install antlr 4.13

***notice: make sure that's same version for antlr software and antlr go package, 4.13***

install wget
```shell
$ brew install wget
```

downlaod antlr
```shell
$ cd ~/dev
$ wget https://www.antlr.org/download/antlr-4.13.1-complete.jar
```

config antlr
```shell
$ vim ~/.zshrc
CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:/usr/local/lib/antlr-4.11.1-complete.jar
alias antlr4='java -jar ~/dev/antlr-4.13.1-complete.jar'

$ source ~/.zshrc
```

test antlr
```shell
$ antlr4
ANTLR Parser Generator  Version 4.13.1
 -o ___              specify output directory where all output is generated
 -lib ___            specify location of grammars, tokens files
 -atn                generate rule augmented transition network diagrams
 -encoding ___       specify grammar file encoding; e.g., euc-jp
 -message-format ___ specify output style for messages in antlr, gnu, vs2005
 -long-messages      show exception details when available for errors and warnings
 -listener           generate parse tree listener (default)
 -no-listener        don't generate parse tree listener
 -visitor            generate parse tree visitor
 -no-visitor         don't generate parse tree visitor (default)
 -package ___        specify a package/namespace for the generated code
 -depend             generate file dependencies
 -D<option>=value    set/override a grammar-level option
 -Werror             treat warnings as errors
 -XdbgST             launch StringTemplate visualizer on generated code
 -XdbgSTWait         wait for STViz to close before continuing
 -Xforce-atn         use the ATN simulator for all predictions
 -Xlog               dump lots of logging info to antlr-timestamp.log
 -Xexact-output-dir  all output goes into -o dir regardless of paths/package
```

#### regenerate api by using antlr
enter this project folder, and go to the g4 folder
```shell
$ cd api/parser/g4

$ antlr4 -visitor -Dlanguage=Go -o parser -package parser ApiParser.g4
```

generate parser
```shell

```

```shell
$ go get github.com/antlr4-go/antlr/v4@v4.13.0

```

Read document at https://go-zero.dev/docs/tutorials/cli/overview
