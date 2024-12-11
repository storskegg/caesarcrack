# CaesarCrack

This is my poor implementation of a dictionary attack on Caesar ciphers.

Originally, I wrote this to guess at the best shift (most dictionary hits) of the three longest words. I decided to
perform the attack on all 26 shifts, and pick the highest confidence based on number of dictionary hits as a means to
better handle dictionary misses, misspellings, etc.

CaesarCrack is vulnerable to simply misspelling every ciphered word, causing all confidences to be very low.

### TODO:

- ~~Replace test string with file input~~
- Support stdin for cli piping (e.g. `pbpaste | caesarcrack`)
- consider displaying confidence distribution chart in cli
- ~~Support punctuation, etc.~~
- ~~Handle case~~
- ~~Support Internal and Declared Dictionaries~~
- Support misspellings (being drawn up in https://github.com/storskegg/autocorrect)
