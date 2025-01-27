package app

const (
	// Name is the application name. Also used to derive names for various application-related objects.
	Name = "gardener-custom-metrics"
	// Uri is an all-purpose identifier of the application, in URI format.
	Uri = "custom-metrics.gardener.cloud"
)

// Log verbosity
const (
	VerbosityError   = 0
	VerbosityWarning = 25
	VerbosityInfo    = 50
	VerbosityVerbose = 75
	VerbosityDebug   = 100
)
