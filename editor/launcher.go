package editor

func LaunchByConfig(config Config, i Interface) {
	log := i.GetLogger()
	debug := config.DebugMode
	log.SetDebugMode(debug)

	if debug {
		log.Debugf("You should now be able to see debug and trace logs.")
	}

	// TODO: Run launch actions after projects are implemented.
}
