package settings

type TemplateRenderer interface {
	Render(filePath string, context map[string]any) (string, error)
}

type Settings struct {
	Addr             string
	TemplatesDirs    []string
	TemplateRenderer TemplateRenderer
}

type Option func(s *Settings)

func WithAddr(addr string) Option {
	return func(s *Settings) {
		s.Addr = addr
	}
}

func WithTemplateDirs(td []string) Option {
	return func(s *Settings) {
		s.TemplatesDirs = td
	}
}
