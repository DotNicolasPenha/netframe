package logger

import "log"

func LogError(err error) {
	log.Printf("[ERROR] - %s.\n", err.Error())
}
func LogInfo(info string) {
	log.Printf("[INFO] - %s.\n", info)
}
func LogDebug(debugInfo string) {
	log.Printf("[DEBUG] - %s.\n", debugInfo)
}
func LogWarn(warnInfo string) {
	log.Printf("[WARN] - %s.\n", warnInfo)
}
