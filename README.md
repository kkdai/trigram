Trigram Indexing
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/trigram/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/trigram?status.svg)](https://godoc.org/github.com/kkdai/trigram)  [![Build Status](https://travis-ci.org/kkdai/trigram.svg?branch=master)](https://travis-ci.org/kkdai/trigram)


Trigram Phrase Matching is a method of identifying phrases that have a high probability of being synonyms. It is based on representing each phrase by a set of character trigrams that are extracted from that phrase. The character trigrams are used as key terms in a representation of the phrase much as words are used as key terms to represent a document. The similarity of phrases is then computed using the vector cosine similarity measure.  (cited from [here](http://ii.nlm.nih.gov/MTI/Details/trigram.shtml))



Install
---------------
`go get github.com/kkdai/trigram`


Usage
---------------

```go


```


Inspired
---------------

- [Google Code Search (using Go)](https://github.com/google/codesearch)
- [Trigram Algorithm](http://ii.nlm.nih.gov/MTI/Details/trigram.shtml)
- [https://github.com/dgryski/go-trigram](https://github.com/dgryski/go-trigram)
- [Regular Expression Matching with a Trigram Index](https://swtch.com/~rsc/regexp/regexp4.html)
- [Approximate string-matching algorithms, part 2](http://www.morfoedro.it/doc.php?n=223&lang=en#SimilarityMetric)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

