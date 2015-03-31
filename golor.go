
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
  L_CYAN = "\033[1;36m"
  CYAN = "\033[0;36m"
  L_BLUE = "\033[1;34m"
  BLUE = "\033[0;34m"
  L_PURPLE = "\033[1;35m"
  PURPLE = "\033[0;35m"
  BROWN = "\033[0;33m"
  L_GREY = "\033[0;37m"
  GREY = "\033[1;30m"
  BLACK = "\033[0;30m"
  WHITE = "\033[1;37m"
  DEFAULT = "\033[0;00m"
)

func build(c string, s string) string {
  return c + s + DEFAULT
}

// ======
// Fun :)
// ======

// Formats a string to be FAAAAABULOUS
func Rainbow(s string) string {
  rb := ""
  for i, ch := range s {
    c := string(ch)
    switch i % 11 {
      case 0:
        rb += RED + c
        break
      case 1:
        rb += L_RED + c
        break
      case 2:
        rb += YELLOW + c
        break
      case 3:
        rb += L_GREEN + c
        break
      case 4:
        rb += GREEN + c
        break
      case 5:
        rb += L_CYAN + c
        break
      case 6:
        rb += CYAN + c
        break
      case 7:
        rb += L_BLUE + c
        break
      case 8:
        rb += BLUE + c
        break
      case 9:
        rb += L_PURPLE + c
        break
      case 10:
        rb += PURPLE + c
        break
    }
  }
  rb += DEFAULT
  return rb
}

// ===
// Red
// ===

func Red(s string) string {
  return build(RED, s)
}

func PrintRed(s string) {
  fmt.Print(Red(s))
}

func PrintlnRed(s string) {
  fmt.Println(Red(s))
}

func PrintfRed(s string, a ...interface{}) {
  fmt.Printf(Red(s), a...)
}

// =========
// Light Red
// =========

func LightRed(s string) string {
  return build(L_RED, s)
}

func PrintLightRed(s string) {
  fmt.Print(LightRed(s))
}

func PrintlnLightRed(s string) {
  fmt.Println(LightRed(s))
}

func PrintfLightRed(s string, a ...interface{}) {
  fmt.Printf(LightRed(s), a...)
}

// ======
// Yellow
// ======

func Yellow(s string) string {
  return build(YELLOW, s)
}

func PrintYellow(s string) {
  fmt.Print(Yellow(s))
}

func PrintlnYellow(s string) {
  fmt.Println(Yellow(s))
}

func PrintfYellow(s string, a ...interface{}) {
  fmt.Printf(Yellow(s), a...)
}

// ===========
// Light Green
// ===========

func LightGreen(s string) string {
  return build(L_GREEN, s)
}

func PrintLightGreen(s string) {
  fmt.Print(LightGreen(s))
}

func PrintlnLightGreen(s string) {
  fmt.Println(LightGreen(s))
}

func PrintfLightGreen(s string, a ...interface{}) {
  fmt.Printf(LightGreen(s), a...)
}

// =====
// Green
// =====

func Green(s string) string {
  return build(GREEN, s)
}

func PrintGreen(s string) {
  fmt.Print(Green(s))
}

func PrintlnGreen(s string) {
  fmt.Println(Green(s))
}

func PrintfGreen(s string, a ...interface{}) {
  fmt.Printf(Green(s), a...)
}

// ==========
// Light Cyan
// ==========

func LightCyan(s string) string {
  return build(L_CYAN, s)
}

func PrintLightCyan(s string) {
  fmt.Print(LightCyan(s))
}

func PrintlnLightCyan(s string) {
  fmt.Println(LightCyan(s))
}

func PrintfLightCyan(s string, a ...interface{}) {
  fmt.Printf(LightCyan(s), a...)
}

// ====
// Cyan
// ====

func Cyan(s string) string {
  return build(CYAN, s)
}

func PrintCyan(s string) {
  fmt.Print(Cyan(s))
}

func PrintlnCyan(s string) {
  fmt.Println(Cyan(s))
}

func PrintfCyan(s string, a ...interface{}) {
  fmt.Printf(Cyan(s), a...)
}

// ==========
// Light Blue
// ==========

func LightBlue(s string) string {
  return build(L_BLUE, s)
}

func PrintLightBlue(s string) {
  fmt.Print(LightBlue(s))
}

func PrintlnLightBlue(s string) {
  fmt.Println(LightBlue(s))
}

func PrintfLightBlue(s string, a ...interface{}) {
  fmt.Printf(LightBlue(s), a...)
}

// ====
// Blue
// ====

func Blue(s string) string {
  return build(BLUE, s)
}

func PrintBlue(s string) {
  fmt.Print(Blue(s))
}

func PrintlnBlue(s string) {
  fmt.Println(Blue(s))
}

func PrintfBlue(s string, a ...interface{}) {
  fmt.Printf(Blue(s), a...)
}

// ============
// Light Purple
// ============

func LightPurple(s string) string {
  return build(L_PURPLE, s)
}

func PrintLightPurple(s string) {
  fmt.Print(LightPurple(s))
}

func PrintlnLightPurple(s string) {
  fmt.Println(LightPurple(s))
}

func PrintfLightPurple(s string, a ...interface{}) {
  fmt.Printf(LightPurple(s), a...)
}

// ======
// Purple
// ======

func Purple(s string) string {
  return build(PURPLE, s)
}

func PrintPurple(s string) {
  fmt.Print(Purple(s))
}

func PrintlnPurple(s string) {
  fmt.Println(Purple(s))
}

func PrintfPurple(s string, a ...interface{}) {
  fmt.Printf(Purple(s), a...)
}

// ====
// Brown
// ====

func Brown(s string) string {
  return build(BROWN, s)
}

func PrintBrown(s string) {
  fmt.Print(Brown(s))
}

func PrintlnBrown(s string) {
  fmt.Println(Brown(s))
}

func PrintfBrown(s string, a ...interface{}) {
  fmt.Printf(Brown(s), a...)
}

// =========
// LightGrey
// =========

func LightGrey(s string) string {
  return build(L_GREY, s)
}

func PrintLightGrey(s string) {
  fmt.Print(LightGrey(s))
}

func PrintlnLightGrey(s string) {
  fmt.Println(LightGrey(s))
}

func PrintfLightGrey(s string, a ...interface{}) {
  fmt.Printf(LightGrey(s), a...)
}

// ====
// Grey
// ====

func Grey(s string) string {
  return build(GREY, s)
}

func PrintGrey(s string) {
  fmt.Print(Grey(s))
}

func PrintlnGrey(s string) {
  fmt.Println(Grey(s))
}

func PrintfGrey(s string, a ...interface{}) {
  fmt.Printf(Grey(s), a...)
}

// =====
// Black
// =====

func Black(s string) string {
  return build(BLACK, s)
}

func PrintBlack(s string) {
  fmt.Print(Black(s))
}

func PrintlnBlack(s string) {
  fmt.Println(Black(s))
}

func PrintfBlack(s string, a ...interface{}) {
  fmt.Printf(Black(s), a...)
}

// =====
// White
// =====

func White(s string) string {
  return build(WHITE, s)
}

func PrintWhite(s string) {
  fmt.Print(White(s))
}

func PrintlnWhite(s string) {
  fmt.Println(White(s))
}

func PrintfWhite(s string, a ...interface{}) {
  fmt.Printf(White(s), a...)
}
