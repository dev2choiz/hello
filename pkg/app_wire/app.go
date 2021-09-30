package app_wire

// InitApp initialize app dependencies
func InitApp() {
	InitializeLogger()
	InitializePostgres()
}

// InitAppTest initialize app dependencies for test
func InitAppTest() {
	InitializeLoggerMock()
}
