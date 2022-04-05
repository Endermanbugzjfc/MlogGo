package editor

type Editor interface {
	GetLogger() Logger
	OpenProject(project any) // TODO: Mlog project.
	CloseProject()
	ListProjects(projects []any) // TODO: Mlog project manifests.
}
