# ASCII Art
Ascii Art is a CLI program written in Go Language that draws the ASCII art of the ASCII text you pass to it as an argument.

It uses only the standard libraries of Go language. 

It uses banner files that have the art for each character arranged in the order of the ASCII table and separated by a newline.

It only works for ASCII characters. Unicode characters beyond the [ASCII table](https://www.ascii-code.com/) will cause errors.

Command characters apart from newlines will cause panics.

Making it concurrent was a challenging and enlightening experience.

## Installation
- Ask to be a collaborator
- git clone (this repo)

## Usage
- Change directory to the `ascii-art` folder
- Run `go run . <sample text>` in the terminal
- To use a style other than the standard style. 
- - Ensure that the banner file of the style is in the `banners` folder as a `.txt` file.
- - Ensure the banner file is formatted properly as described in the project description.
- - Run `go run . <sample text> <style banner file name without .txt`

## Examples
### Basic Usage
![Example of how the program runs without newlines and without special sytle](test(1).png)

### With Newline
![Example of how the program runs with newlines without special style](test(2).png)

### With Special Style
![Example of how the program runs with newlines and with special style](test(3).png)

## Credits
- Me (Developer)
- Myself (Tester)
- I (Contributor)

## License
Copyright © 2026 [Olamide Ifarajimi](https://acad.learn2earn.ng/git/oifaraji)

This Project is [GPL](https://www.gnu.org/licenses/gpl-3.0.en.html) Licensed