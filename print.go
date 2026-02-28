package req

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	C "github.com/fatih/color"
)

var (
	enabled      bool
	filter       = C.New(C.BgYellow, C.Bold)
	forceDisable bool
)

func EnabledDebug() bool {
	return enabled
}

func DisableDebug() {
	enabled = false
}

func ForceDisableDebug(disable bool) {
	forceDisable = disable
}

func EnableDebug() {
	enabled = true
}

// func PrintBlack(format string, a ...interface{}) {
// 	base(false, "black", format, a...)
// }

// func PrintRed(format string, a ...interface{}) {
// 	base(false, "red", format, a...)
// }

// func PrintGreen(format string, a ...interface{}) {
// 	base(false, "green", format, a...)
// }

// func PrintYellow(format string, a ...interface{}) {
// 	base(false, "yellow", format, a...)
// }

// func PrintBlue(format string, a ...interface{}) {
// 	base(false, "blue", format, a...)
// }

// func PrintMagenta(format string, a ...interface{}) {
// 	base(false, "magenta", format, a...)
// }

// func PrintCyan(format string, a ...interface{}) {
// 	base(false, "cyan", format, a...)
// }

// func PrintWhite(format string, a ...interface{}) {
// 	base(false, "white", format, a...)
// }

// func SprintBlack(format string, a ...interface{}) string {
// 	return baseString("black", format, a...)
// }

// func SprintRed(format string, a ...interface{}) string {
// 	return baseString("red", format, a...)
// }

// func SprintGreen(format string, a ...interface{}) string {
// 	return baseString("green", format, a...)
// }

// func SprintYellow(format string, a ...interface{}) string {
// 	return baseString("yellow", format, a...)
// }

// func SprintBlue(format string, a ...interface{}) string {
// 	return baseString("blue", format, a...)
// }

// func SprintMagenta(format string, a ...interface{}) string {
// 	return baseString("magenta", format, a...)
// }

// func SprintCyan(format string, a ...interface{}) string {
// 	return baseString("cyan", format, a...)
// }

// func SprintWhite(format string, a ...interface{}) string {
// 	return baseString("white", format, a...)
// }

func base(force bool, kind string, format string, a ...interface{}) {
	if !force && !enabled {
		return
	}

	if force {
		if forceDisable {
			return
		}
	}

	switch kind {
	case "black":
		C.Black(format, a...)
	case "red":
		C.Red(format, a...)
	case "green":
		C.Green(format, a...)
	case "yellow":
		C.Yellow(format, a...)
	case "blue":
		C.Blue(format, a...)
	case "magenta":
		C.Magenta(format, a...)
	case "cyan":
		C.Cyan(format, a...)
	case "white":
		C.White(format, a...)
	}

	C.Unset()
}

// func Black(format string, a ...interface{}) {
// 	base(false, "black", format, a...)
// }

func printRed(format string, a ...interface{}) {
	base(false, "red", format, a...)
}

func printGreen(format string, a ...interface{}) {
	base(false, "green", format, a...)
}

func printYellow(format string, a ...interface{}) {
	base(false, "yellow", format, a...)
}

// func Blue(format string, a ...interface{}) {
// 	base(false, "blue", format, a...)
// }

func printMagenta(format string, a ...interface{}) {
	base(false, "magenta", format, a...)
}

func printCyan(format string, a ...interface{}) {
	base(false, "cyan", format, a...)
}

// func White(format string, a ...interface{}) {
// 	base(false, "white", format, a...)
// }

func forceBlack(format string, a ...interface{}) {
	base(true, "black", format, a...)
}

func forceRed(format string, a ...interface{}) {
	base(true, "red", format, a...)
}

func forceGreen(format string, a ...interface{}) {
	base(true, "green", format, a...)
}

func forceYellow(format string, a ...interface{}) {
	base(true, "yellow", format, a...)
}

func forceBlue(format string, a ...interface{}) {
	base(true, "blue", format, a...)
}

func forceMagenta(format string, a ...interface{}) {
	base(true, "magenta", format, a...)
}

func forceCyan(format string, a ...interface{}) {
	base(true, "cyan", format, a...)
}

func forceWhite(format string, a ...interface{}) {
	base(true, "white", format, a...)
}

// func SBlack(format string, a ...interface{}) string {
// 	return baseString("black", format, a...)
// }

// func SRed(format string, a ...interface{}) string {
// 	return baseString("red", format, a...)
// }

// func SGreen(format string, a ...interface{}) string {
// 	return baseString("green", format, a...)
// }

// func SYellow(format string, a ...interface{}) string {
// 	return baseString("yellow", format, a...)
// }

// func SBlue(format string, a ...interface{}) string {
// 	return baseString("blue", format, a...)
// }

// func SMagenta(format string, a ...interface{}) string {
// 	return baseString("magenta", format, a...)
// }

// func SCyan(format string, a ...interface{}) string {
// 	return baseString("cyan", format, a...)
// }

// func SWhite(format string, a ...interface{}) string {
// 	return baseString("white", format, a...)
// }

func baseString(kind string, format string, a ...interface{}) string {
	if !enabled {
		return ""
	}

	switch kind {
	case "black":
		return C.BlackString(format, a...)
	case "red":
		return C.RedString(format, a...)
	case "green":
		return C.GreenString(format, a...)
	case "yellow":
		return C.YellowString(format, a...)
	case "blue":
		return C.BlueString(format, a...)
	case "magenta":
		return C.MagentaString(format, a...)
	case "cyan":
		return C.CyanString(format, a...)
	case "white":
		return C.WhiteString(format, a...)
	}

	return ""
}

func printPanic(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	panic(s)
}

// func RedPanic(format string, a ...interface{}) {
// 	s := SRed(format, a...)
// 	panic(s)
// }

// 系统类: Cyan
// 错误类: Red
// 提示类: Magenta
// 强提示: Yellow
// 成功: Green

func prefix(mark, format string) string {
	ts := time.Now().Format("2006-01-02 15:04:05.000")
	mark = fmt.Sprintf("[%s]", mark)
	return fmt.Sprintf("%23s %9s: %s", ts, mark, format)

}

func Mark(format string, a ...interface{}) {
	format = prefix("mark", format)
	forceRed(format, a...)
	os.Exit(0)
}

func printSystem(format string, a ...interface{}) {
	format = prefix("system", format)
	forceCyan(format, a...)
}

func printError(format string, a ...interface{}) {
	format = prefix("error", format)
	forceRed(format, a...)
}

func printNotice(format string, a ...interface{}) {
	format = prefix("notice", format)
	printMagenta(format, a...)
}

func printWarning(format string, a ...interface{}) {
	format = prefix("warning", format)
	printYellow(format, a...)
}

func printInfo(format string, a ...interface{}) {
	format = prefix("info", format)
	printGreen(format, a...)
}

func printJson(obj any) {
	if !enabled {
		return
	}
	js, _ := json.Marshal(obj)
	base(false, "red", string(js))
}

// func JsonExit(obj any) {
// 	if !enabled {
// 		return
// 	}
// 	js, _ := json.Marshal(obj)
// 	base(false, "red", string(js))
// 	os.Exit(0)
// }

// func Filter(format string, a ...interface{}) {
// 	if !enabled {
// 		return
// 	}

// 	format = strings.TrimSuffix(format, "\n")

// 	filter = C.New(C.BgYellow, C.Bold)
// 	format = prefix("filter", format)
// 	_, _ = filter.Printf(format, a...)
// 	fmt.Println()
// }

// func Request(format string, a ...interface{}) {
// 	if !enabled {
// 		return
// 	}

// 	format = strings.TrimSuffix(format, "\n")

// 	filter = C.New(C.BgCyan, C.Bold)
// 	format = prefix("request", format)
// 	_, _ = filter.Printf(format, a...)
// 	fmt.Println()
// }

// func SRequest(format string, a ...interface{}) string {
// 	if !enabled {
// 		return ""
// 	}

// 	format = strings.TrimSuffix(format, "\n")

// 	filter = C.New(C.BgCyan, C.Bold)
// 	format = prefix("request", format)
// 	return filter.Sprintf(format, a...)
// }

// func RequestError(format string, a ...interface{}) {
// 	if !enabled {
// 		return
// 	}

// 	format = strings.TrimSuffix(format, "\n")

// 	filter = C.New(C.BgRed, C.Bold)
// 	format = prefix("error", format)
// 	_, _ = filter.Printf(format, a...)
// 	fmt.Println()
// }
