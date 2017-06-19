package orchestrator

type App struct {
	Id      string
	Type    string
	Startup map[string]interface{}
}
type Orchestrator struct {
	Apps    map[string]*App
	Factory Factory
}

func New(f Factory) *Orchestrator {
	return &Orchestrator{
		Factory: f,
		Apps:    make(map[string]*App),
	}
}

type Factory interface {
	NewApp(app *App) error
}

type CommandFactory struct {
}
