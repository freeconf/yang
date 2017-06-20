package app

type App struct {
	Id      string
	Type    string
	Startup map[string]interface{}
}

type Orchestrator struct {
	Apps    map[string]*App
	Builder Builder
}

func New() *Orchestrator {
	return &Orchestrator{
		Apps: make(map[string]*App),
	}
}

type Builder interface {
	NewApp(app *App) error
}
