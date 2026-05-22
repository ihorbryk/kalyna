package apps

import (
	"path/filepath"
	"runtime"

	"github.com/ihorbryk/kalyna/settings"
	kalynatemplate "github.com/ihorbryk/kalyna/template"
)

type App struct {
	sttfn  func(stt *settings.Settings)
	appDir string
}

func New() *App {
	_, file, _, _ := runtime.Caller(1)
	app := &App{
		appDir: filepath.Dir(file),
	}

	return app
}

func (app *App) AppDir() string {
	return app.appDir
}

func (app *App) AppTemplatesDir() string {
	return filepath.Join(app.appDir, kalynatemplate.BaseTemplateFolder)
}

func (app *App) SetSetup(sttfn func(stt *settings.Settings)) {
	app.sttfn = sttfn
}

func (app *App) Setup(stt *settings.Settings) {
	if app.sttfn != nil {
		app.sttfn(stt)
	}
}

func Apps(apps ...*App) []*App {
	return apps
}
