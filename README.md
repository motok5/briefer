![gopher](./gopher.gif)

 

# briefer

Go language implementation of sentence summarization by automatic summarization algorithm LexRank.

 

## Dependency

* [github.com/ikawaha/kagome](https://github.com/ikawaha/kagome)

* [github.com/JesusIslam/tldr](https://github.com/JesusIslam/tldr)

  

## Setup

It can be installed by "go get".

```shell
$ go get github.com/motok5/briefer
```



## Usage

```shell
$ briefer -l=ja -c=10 input.txt
```

### -l option (Optional)

 Select the language of the sentences to be summarized. "briefer" supports English and Japanese.

The default is English.

### -c option (Optional)

Specify the number of output lines of the summary. The default is 3.

### input file (Required)

Specify the file to be summarized.



## Licence

This software is released under the MIT License, see LICENSE.

  