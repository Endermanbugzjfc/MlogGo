package editor

import (
	"encoding/json"
	"flag"
	"os"
)

type OnLaunch string

const (
	OpenLastProject OnLaunch = "open-last-project"
	NewEmptyProject          = "new-empty-project"
	ListProjects             = "list-projects"
)

type Config struct {
	ProjectPaths []string
	OnLaunch     OnLaunch
	DebugMode    bool
	CompactView  bool
	Keys         struct {
		Read        KeySet
		Write       KeySet
		Draw        KeySet
		Print       KeySet
		DrawFlush   KeySet
		PrintFlush  KeySet
		GetLink     KeySet
		Control     KeySet
		Radar       KeySet
		Sensor      KeySet
		Set         KeySet
		Operation   KeySet
		End         KeySet
		Jump        KeySet
		UnitBind    KeySet
		UnitControl KeySet
		UnitRadar   KeySet
		UnitLocate  KeySet
	}
}

const (
	dataPath = "~/.mloggo/"
)

func DefaultConfig() Config {
	return Config{
		ProjectPaths: []string{
			dataPath + "projects",
		},
		OnLaunch:    OpenLastProject,
		DebugMode:   false,
		CompactView: true,
		Keys: struct {
			Read        KeySet
			Write       KeySet
			Draw        KeySet
			Print       KeySet
			DrawFlush   KeySet
			PrintFlush  KeySet
			GetLink     KeySet
			Control     KeySet
			Radar       KeySet
			Sensor      KeySet
			Set         KeySet
			Operation   KeySet
			End         KeySet
			Jump        KeySet
			UnitBind    KeySet
			UnitControl KeySet
			UnitRadar   KeySet
			UnitLocate  KeySet
		}{
			Read:        "r",
			Write:       "w",
			Draw:        "w",
			Print:       "t",
			DrawFlush:   "f",
			PrintFlush:  "f",
			GetLink:     "g",
			Control:     "c",
			Radar:       "a",
			Sensor:      "s",
			Set:         "s", // Var.
			Operation:   "a",
			End:         "e",
			Jump:        "g", // If.
			UnitBind:    "b",
			UnitControl: "r", // F is the only set that requires two keys.
			UnitRadar:   "d",
			UnitLocate:  "e",
		},
	}
}

// MustLoadConfig should only be used on desktop.
// Logs error and return default config if there is one.
func MustLoadConfig(logger Logger) (config Config) {
	var err error
	config, err = LoadConfig()

	if err != nil {
		logger.Errorf("Failed to load config, default config will be used until manual reload: %s", err)
	}

	return
}

// LoadConfig should only be used on desktop.
func LoadConfig() (config Config, err error) {
	const (
		usage        = "Raw JSON data or a file path."
		defaultValue = dataPath + "config.json"
	)

	argument := flag.String("config", defaultValue, usage)
	config = DefaultConfig()

	var data []byte
	if data, err = os.ReadFile(*argument); err != nil {
		data = []byte(*argument)
	}

	err = json.Unmarshal(data, &config)

	return
}
