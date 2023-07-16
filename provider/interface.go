package provider

type Builder interface { // all cloud providers would belong to this
	Deploy() ([]byte, error)
}
