# ASCII Art Color
Ascii Art Color is a CLI program written in Go Language that draws the ASCII art of the ASCII text you pass to it as an argument and colors the substring you can optionally provide.

It uses only the standard libraries of Go language. 

It uses banner files that have the art for each character arranged in the order of the ASCII table and separated by a newline.

It only works for ASCII characters. Unicode characters beyond the [ASCII table](https://www.ascii-code.com/) will cause errors.

Command characters apart from newlines will cause panics.

Making it iterative, character by character drawing, with ANSI Control Sequences was a novel experience.

I hope to be able to expand the project to more than just ASCII characters.

## Installation
- Ask to be a collaborator
- `git clone https://acad.learn2earn.ng/git/obolarinw/ascii-art-color.git`

## Usage
- Change directory to the `ascii-art-color` folder
- Run `go run . --color=<color>OPTIONAL <substring to be colored>OPTIONAL <text to draw>` in the terminal
- To use a style other than the standard style. 
- - Ensure that the banner file of the style is in the `banners` folder as a `.txt` file.
- - Ensure the banner file is formatted properly as described in the project description.
- - Run `go run . --color=<color>OPTIONAL <substring to be colored>OPTIONAL <text to draw> <style banner file name without .txt>OPTIONAL`

- The program gives an error if there is no text or if the color is invalid.
- The valid colors are: black, red, green, yellow, blue, magenta, cyan, white, default

## Demo
### Basic Usage
![Demonstration of how the program runs with one argument](./demo_images/image.png)

### With two arguments (text and banner style)
![Demonstration of how the program runs with two arguments. "Text" text first and "shadow" banner style ](./demo_images/image-1.png)

### With color flag and text argument
![Demonstration of how the program runs with  color flag = red and a single text argument](./demo_images/image-2.png)

### With color flag, substring and text containing the substring
![Demonstration of how the program runs with color flag = red, a substring and a text containing the substring](./demo_images/image-3.png)

### With color flag, substring and text that doesn't contain the substring
![Demonstration of how the program runs with color flag = red, a substring and a text that doesn't contain the substring](./demo_images/image-4.png)

### With color flag, substring, text containing the substring and banner style, shadow
![Demonstration of how the program runs with color flag = yellow, a substring, a text that contains the substring and a banner style, shadow](./demo_images/image-5.png)

## Credits
- [Olamide Ifarajimi](https://acad.learn2earn.ng/git/oifaraji) (Developer)
- [Opeoluwa Adekunle](https://acad.learn2earn.ng/git/oadekunle) (Tester)
- [Oladimeji Bolarinwa](https://acad.learn2earn.ng/git/obolarinw) (Team Lead)

## License
Copyright © 2026

This Project is [GPL](https://www.gnu.org/licenses/gpl-3.0.en.html) Licensed