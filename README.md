# golor
A go library for formatting the color of text output to a terminal

Status:

This library does work. However, I primarily use it just as a reference for all the ASCII color codes in other projects. I'm sure there are better terminal coloring libraries for Go available now, so I'd suggest using those over this one.

# Usage

`go get github.com/mhoc/golor`

```
import github.com/mhoc/golor
package main
func main() {
  golor.PrintlnPurple("This text is absolutely royal!")
}
```

# Functions

```
func <ColorName>(string) string`
```

Returns a string which is formatted the given color, suitable for printing. Example: `golor.Red("ERROR!")`

```
func Print<ColorName>(string)`
```

Prints out a string in the given color. Example: `golor.PrintYellow("WARNING!")`

```
func Println<ColorName>(string)`
```

Printlns a string in the given color. Example: `golor.PrintlnGreen("GREAT!")`

```
func Printf<ColorName>(string, ...interface{})`
```

Printfs a string and arguments, just like you'd expect. Example: `golor.PrintfCyan("Is this %s\n", "Blue?")`

# Colors

You can pick from any of these fabulous colors!

* Red
* LightRed
* Yellow
* LightGreen
* Green
* LightCyan
* Cyan
* LightBue 
* Blue
* LightPurple
* Purple
* Brown
* LightGrey
* Grey
* Black
* White

# Extras

```
func Rainbow(string) string
```

You can guess what this does :)
