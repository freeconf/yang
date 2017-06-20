package launch

type App struct {
	Id      string
	Type    string
	Startup map[string]interface{}
}

type Pad struct {
	Apps     map[string]*App
	Launcher Launcher
}

func New() *Pad {
	return &Pad{
		Apps: make(map[string]*App),
	}
}

type Launcher interface {
	Launch(app *App) error
}
