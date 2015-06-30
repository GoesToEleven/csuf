package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("TERM"))
	fmt.Println(os.Getenv("PS1"))
	fmt.Println(os.Getenv("MAIL"))
	fmt.Println(os.Getenv("TEMP"))
	fmt.Println(os.Getenv("JAVA_HOME"))
	fmt.Println(os.Getenv("ORACLE_HOME"))
	fmt.Println(os.Getenv("TZ"))
	fmt.Println(os.Getenv("PWD"))
	fmt.Println(os.Getenv("HISTFILE"))
	fmt.Println(os.Getenv("HISTFILESIZE"))
	fmt.Println(os.Getenv("HOSTNAME"))
	fmt.Println(os.Getenv("LD_LIBRARY_PATH"))
	fmt.Println(os.Getenv("USER"))
	fmt.Println(os.Getenv("DISPLAY"))
	fmt.Println(os.Getenv("SHELL"))
	fmt.Println(os.Getenv("TERMCAP"))
	fmt.Println(os.Getenv("OSTYPE"))
	fmt.Println(os.Getenv("MACHTYPE"))
	fmt.Println(os.Getenv("EDITOR"))
	fmt.Println(os.Getenv("PAGER"))
	fmt.Println(os.Getenv("MANPATH"))
	fmt.Println(os.Getenv("ENV"))
}
