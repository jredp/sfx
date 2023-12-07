package main

const modOUTDATED = -1
const modEQUAL = 0
const modUPDATE = 1
const modINSTALL = 2

func getInstallMode(config *peConfig, destConfig *peConfig) int8 {
	return modINSTALL
}
