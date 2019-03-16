package common

type MetaData map[string][]string

// callInfo contains all related configuration and information .
type callInfo struct {
	metaData MetaData
}

func defaultCallInfo() *callInfo {
	return &callInfo{}
}

// CallOption configures a Call before it starts or extracts information from
// a Call after it completes.
type CallOption interface {
	// before is called before the call is sent to any server.  If before
	// returns a non-nil error, the RPC fails with that error.
	before(*callInfo) error

	// after is called after the call has completed.  after cannot return an
	// error, so any failures should be reported via output parameters.
	after(*callInfo)
}
