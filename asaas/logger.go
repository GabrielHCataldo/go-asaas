package asaas

import (
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"log"
	"os"
)

var lInfo = log.New(os.Stdout, "[ASAAS \u001b[34mINFO: \u001B[0m", log.LstdFlags)
var lWarning = log.New(os.Stdout, "[ASAAS \u001b[33mWARNING: \u001B[0m", log.LstdFlags)
var lError = log.New(os.Stdout, "[ASAAS \u001b[31mERROR: \u001b[0m", log.LstdFlags)
var lDebug = log.New(os.Stdout, "[ASAAS \u001b[36mDEBUG: \u001B[0m", log.LstdFlags)

func logInfo(env Env, v ...any) {
	if env == EnvProduction {
		return
	}
	lInfo.Print(getSystemMessageDefault(3), fmt.Sprintln(v...))
}

func logInfoSkipCaller(skipCaller int, env Env, v ...any) {
	if env == EnvProduction {
		return
	}
	lInfo.Print(getSystemMessageDefault(skipCaller), fmt.Sprintln(v...))
}

func logWarning(v ...any) {
	lWarning.Print(getSystemMessageDefault(3), fmt.Sprintln(v...))
}

func logDebugSkipCaller(skipCaller int, v ...any) {
	lDebug.Print(getSystemMessageDefault(skipCaller), fmt.Sprintln(v...))
}

func logError(v ...any) {
	lError.Print(getSystemMessageDefault(3), fmt.Sprintln(v...))
}

func logErrorSkipCaller(skipCaller int, v ...any) {
	lError.Print(getSystemMessageDefault(skipCaller), fmt.Sprintln(v...))
}

func getSystemMessageDefault(skipCaller int) string {
	_, line, funcName := util.GetSystemInfo(skipCaller + 1)
	return funcName + ":" + line + "] "
}
