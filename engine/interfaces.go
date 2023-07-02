package engine

type Director interface {
}

type Engine interface {
	Run() EngineResponse
}
