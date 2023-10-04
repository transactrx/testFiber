package appHelper

type App struct {
	Name    string
	Version string
}

func (app *App) GetName() string {
	return app.Name
}
