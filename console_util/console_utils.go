package console_util

import (
"os"
"fmt"
"os/exec"
"runtime"
)

var clear map[string]func() //create a map for storing clear funcs



func init() {
    clear = make(map[string]func()) // Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") // Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cls") // Windows example it is untested, but I think its working 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

var COLOR_BLACK = "\033[30m"
var COLOR_RED = "\033[31m"
var COLOR_GREEN = "\033[32m"
var COLOR_YELLOW = "\033[33m"
var COLOR_BLUE = "\033[34m"
var COLOR_PINK = "\033[35m"
var COLOR_CYAN = "\033[36m"
var COLOR_WHITE = "\033[37m"
var COLOR_NORMAL = "\033[0;39m"

func PrintColored(text string, color string){
    fmt.Printf(color + text)
}

func PrintlnColored(text string, color string){
    PrintColored(text + "\n", color)
}

func PrintNormal(text string){
    PrintColored(text, COLOR_NORMAL)
}

func PrintlnNormal(text string){
    PrintNormal(text + "\n")
}


func ExitProgram() {
    os.Exit(0)
}