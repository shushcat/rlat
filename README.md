# rlat

This program compares pieces of writing to identify passages that may have influenced one another---it finds passages, that is to say, that might be "rlat"-ed 🤦.

`rlat` has (or has in development) some features that make it suitable for studies using large corpora that contain many spelling variations:

- a word stemmer using the [Snowball language](https://snowballstem.org), 
- support for filtering using [stop word](https://en.wikipedia.org/wiki/Stop_words) lists, 
- matching of words within a specified [Damarau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau–Levenshtein_distance).

## Usage

To check for similar passages in a target and source text with the default parameters, run the command `rlat -t <target file> -s <source file>`.

Allowable parameters are viewable with the `-h` flag, and are as follows.

| Flag          | Description                                             | 
| --------      | --------                                                | 
| -h            | display **h**elp message                                | 
| -s <file>     | path to **s**ource file                                 | 
| -t <file>     | path to **t**arget file                                 | 
| -w            | **w**indow in which to check for shared words           | 
| -msw          | **m**inimum **s**hared **w**ords in a window            | 
| -mwl          | **m**inimum **w**ord **l**ength for matched words       | 
| -no           | allow for window matches which are **n**ot **o**rdered  | 
| -stem         | **stem** all words in target and source texts           | 
| -stop <file>  | filter all **stop**words in a specified file            | 

## Installation

To install manually, first setup [Go](https://golang.org), then run the command `go install github.com/shushcat/rlat@latest`.

Binaries for various platforms will eventually be provided on a releases page.

## Contributing

Please submit pull requests and bug reports at https://github.com/shushcat/rlat.

## Todo

- [x] Make sure all words are selected when a text is compared to itself with a maximum word length of 1.
- [ ] Import all old tests.
- [ ] Implement an edit distance (`distance.DamLev("word1", "word2")`) function.
- [ ] Implement a flag to return comparator density (see [Neidorf 2019](https://www.nature.com/articles/s41562-019-0570-1)).
- [ ] Use the [Snowball stemmer](https://snowballstem.org/) to get word stems.
- [ ] Process multiple texts in parallel.

## Resources

Martin Porter; "Lovins revisited"; December 2001 (revised November 2008); http://snowball.tartarus.org/algorithms/lovins/festschrift.html; _Charting a new course: progress in natural language processing and information retrieval: a festschrift for professor Karen Sparck Jones_, edited by John Tait; 2005.

## License, Credit & History

`rlat` is available free of restrictions under the terms of the [MIT License](https://opensource.org/licenses/MIT).

The basic algorithm `rlat` follows in comparing texts was inspired by the plagiarism-detection program [WCopyfind](http://plagiarism.bloomfieldmedia.com/wordpress/software/wcopyfind/) by Lou Bloomfield.

An earlier version of this program is archived at [the Rcopyfind repository](https://github.com/shushcat/rcopyfind_prototype).
