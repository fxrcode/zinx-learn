package ziface

// IServer interface
type IServer interface {
	// start server
	Start()
	// stop server
	Stop()
	// start serving
	Serve()
}
