package editor

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/df-mc/atomic"
)

type OnLaunch string

const (
	OpenLastProject OnLaunch = "open-last-project"
	NewEmptyProject          = "new-empty-project"
	ListProjects             = "list-projects"
)

var (
	ConfigAtomic atomic.Value[Config]
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

		SwitchCodeBlockList KeySet
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

			SwitchCodeBlockList KeySet
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

			SwitchCodeBlockList: "W",
		},
	}
}

// MustLoadConfig should only be used on desktop.
// Stores the loaded config to variable ConfigAtomic.
// Or stores the default config if there is an error.
// So it does not have to be stored in an external package.
// Logs error and return default config if there is one.
func MustLoadConfig(
	logger Logger,
	configArgument string,
) (config Config) {
	var err error
	config, err = LoadConfig(configArgument)

	if err != nil {
		logger.Errorf("Failed to load config, default config will be used until manual reload: %s", err)
	}

	return
}

// RegisterConfigArgument must be called and only before flag.Parse().
func RegisterConfigArgument() *string {
	const (
		usage        = "Raw JSON data or a file path."
		defaultValue = dataPath + "config.json"
	)

	return flag.String("config", defaultValue, usage)
}

// LoadConfig should only be used on desktop.
// Stores the loaded config to variable ConfigAtomic.
// Or stores the default config if there is an error.
// So it does not have to be stored in an external package.
func LoadConfig(configArgument string) (config Config, err error) {
	config = DefaultConfig()

	var data []byte
	if data, err = os.ReadFile(configArgument); err != nil {
		data = []byte(configArgument)
	}

	err = json.Unmarshal(data, &config)

	ConfigAtomic.Store(config)

	return
}
