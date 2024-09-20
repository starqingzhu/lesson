package myinterface

func PrintInterfaceType() {
	DumpMethodSet((*Writer)(nil))
	DumpMethodSet((*Reader)(nil))
	DumpMethodSet((*ReadWriter)(nil))
	DumpMethodSet((*ReadCloser)(nil))
	DumpMethodSet((*ReadWriteCloser)(nil))
}

type (
	Reader interface {
		Read(p []byte) (n int, err error)
	}

	Writer interface {
		Write(p []byte) (n int, err error)
	}

	Closer interface {
		Close() error
	}
	//组合接口类型
	ReadWriter interface {
		Reader
		Writer
	}
	ReadCloser interface {
		Reader
		Closer
	}
	ReadWriteCloser interface {
		Reader
		Writer
		Closer
	}
)
