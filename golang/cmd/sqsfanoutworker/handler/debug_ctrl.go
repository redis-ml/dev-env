package handler

var (
	isDebugMode    = false
	isFakeSinkMode = false
)

func SetDebugMode(debug bool) {
	isDebugMode = debug
}

func SetUseFakeSink(ctrl bool) {
	isFakeSinkMode = ctrl
}
