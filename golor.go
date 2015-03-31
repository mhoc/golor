
// Golor is a library for printing colored text to a terminal
// github.com/mhoc/golor
package golor

import (
  "fmt"
)

const (
  RED = "\033[0;31m"
  L_RED = "\033[1;31m"
  YELLOW = "\033[1;33m"
  L_GREEN = "\033[1;32m"
  GREEN = "\033[0;32m"
  L_CYAN = = "\033[1;36m"
  CYAN = "\033[0;36m"
  L_BLUE = "\033[1;34m"
  BLUE = "\033[0;34m"
  L_PURPLE = "\033[1;35m"
  PURPLE = "\033[0;35m"
  BROWN = "\033[0;33m"
  L_GREY = "\033[0;37m"
  D_GREY = "\033[1;30m"
  BLACK = "\033[0;30m"
  WHITE = "\033[1;37m"
  DEFAULT = "\033[0;00m"
)

func build(c string, s string) string {
  return c + s + DEFAULT
}

// ===
// Red
// ===

// Formats and returns a string in the "red" color
func Red(s string) string {
  return build(RED, s)
}

// Prints a string in the red color
func PrintRed(s string) {
  fmt.Print(Red(s))
}

// Prints a string and appends a newline in the red color
func PrintRedln(s string) {
  fmt.Println(Red(s))
}

// Printfs a string in the red color
func PrintRedf(s string, a ...interface{}) {
  fmt.Printf(Red(s), a...)
}

// =========
// Light Red
// =========

// Formats and returns a string in the "light red" color
func LightRed(s string) string {
  return build(L_RED, s)
}

// Prints a string in the light red color
func PrintLightRed(s string) {
  fmt.Print(LightRed(s))
}

// Prints a string and appends a newline in the light red color
func PrintLightRedln(s string) {
  fmt.Println(LightRed(s))
}

// Printfs a string in the light red color
func PrintLightRedf(s string, a ...interface{}) {
  fmt.Printf(LightRed(s), a...)
}

// ======
// Yellow
// ======

// Formats and returns a string in the "yellow" color
func Yellow(s string) string {
  return build(YELLOW, s)
}

// Prints a string in the yellow color
func PrintYellow(s string) {
  fmt.Print(Yellow(s))
}

// Prints a string and appends a newline in the yellow color
func PrintYellowln(s string) {
  fmt.Println(Yellow(s))
}

// Printfs a string in the yellow color
func PrintYellowf(s string, a ...interface{}) {
  fmt.Printf(Yellow(s), a...)
}
