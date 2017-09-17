# imgrep

`imgrep` is a command-line utility (and website) written in Go to search
for keywords found within images.
![](.product/web_no_query.png)

## Installation

`imgrep` depends on
[Tesseract](https://github.com/tesseract-ocr/tesseract).
  * On Fedora: `sudo dnf install tesseract-devel`
  * On Debian: `sudo apt-get install libtesseract-dev`

To install:
```
# install dependencies
sudo dnf install tesseract-devel leptonica-devel golang # Fedora
#sudo apt-get install libtesseract-dev libleptonica-dev golang # Debian

# fetch src and install binary
go get github.com/keeferrourke/imgrep
go install github.com/keeferrourke/imgrep

```

## Usage
`imgrep` like `grep`, searches file contents for text. Unlike `grep`
however, `imgrep` searches image files only using Tesseract OCR.

`imgrep` comes with two interfaces; a CLI as one might expect, and a
web-UI graphical front-end.

### CLI
To speed up the process on some machines, `imgrep` can pre-process and
index files by keywords found within them. To preindex an entire
directory (including subdirectories):
```
imgrep init # pre-process and index image files in working directory
```

By default, `imgrep` uses this database of pre-indexed files to perform
simple queries. To check preindexed directories:
```
imgrep search QUERY
```

To use `imgrep` without checking against the database of preindexed
files, simply call
```
imgrep search -n QUERY
```

Like the `grep` family of functions, `imgrep` is useful with Unix-pipes:
```
# Example: Count the number of images that contain the first line of a
#          plain-text file
head -n1 myfile | imgrep search -n - | wc -l
```

### Web UI
`imgrep` comes with a familiar search-based web UI that interacts with
the pre-indexed sqlite database. This web-UI was added to make demo-ing
this utitility easier in a hackathon setting, and is a good POC for how
this `imgrep` may be used.

To start a server on localhost:1337:

```
imgrep server
```
Then just visit 'localhost:1337' in your favourite web brower ;)

#### Preview
![](.product/web_node_query.png)
![](.product/web_me_query.png)


## License
`imgrep` is free software licensed under the MIT license.

Copyright (c) 2017 Keefer Rourke, Ivan Zhang, and Thomas Dedinsky.

See LICENSE for details.
