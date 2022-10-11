# rlat

This program compares pieces of writing to identify passages that may have influenced one another---it finds passages, that is to say, that might be "rlat"-ed 🤦.

Rlat has (or has in development) some features that make it suitable for studies using large corpora that contain many spelling variations:
- a word stemmer using the [Snowball language](https://snowballstem.org), 
- support for filtering using [stop word](https://en.wikipedia.org/wiki/Stop_words) lists, 
- matching of words within a specified [Damarau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau–Levenshtein_distance).

## Project Status

This program is not currently in a usable state.  This notice will be removed once the following objectives are met.

- [ ] Select all words when comparing to self.
- [ ] Implement an edit distance (`distance.DamLev("word1", "word2")`) function.
- [ ] Implement a flag to return comparator density using a function like that used by [Neidorf 2019](https://www.nature.com/articles/s41562-019-0570-1).
- [ ] Implement tests.

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

Binaries for various platforms will eventually be provided on a releases page.

Manual installation requires [Go](https://golang.org), and should only require the command `go get github.com/shushcat/rlat`.

## Contributing

Please submit pull requests and bug reports at https://github.com/shushcat/rlat.

## License, Credit & History

rlat is available free of restrictions under the terms of the [MIT License](https://opensource.org/licenses/MIT).

The basic algorithm `rlat` follows in comparing texts was inspired by the plagiarism-detection program [WCopyfind](http://plagiarism.bloomfieldmedia.com/wordpress/software/wcopyfind/) by Lou Bloomfield.

An earlier version of this program is archived at [the Rcopyfind repository](https://github.com/shushcat/rcopyfind_prototype).
