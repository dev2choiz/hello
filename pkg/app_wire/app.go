package app_wire

// InitApp initialize app dependencies
func InitApp() {
	InitializeLogger()
}

// InitAppTest initialize app dependencies for test
func InitAppTest() {
	InitializeLoggerMock()
}
