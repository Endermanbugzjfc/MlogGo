package editor

type OnLaunch string

const (
	OpenLastProject OnLaunch = "open-last-project"
	NewEmptyProject          = "new-empty-project"
	ListProjects             = "list-projects"
)

type Config struct {
	ProjectPaths []string
	OnLaunch     OnLaunch
}

func DefaultConfig() Config {
	return Config{
		ProjectPaths: []string{
			"~/.mloggo/projects",
		},
		OnLaunch: OpenLastProject,
	}
}
