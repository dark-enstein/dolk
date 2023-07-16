package awspile

type Serve interface { // all aws resources would belong to this
	Deploy() ([]byte, error)
}
