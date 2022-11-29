package common

const (
	// PortEnvVarName is the name of the environment variable that defines the
	// server port
	PortEnvVarName = "PORT"

	// GoogleCloudProjectIDEnvVarName is used to determine the ID of the GCP project e.g for setting up StackDriver client
	GoogleCloudProjectIDEnvVarName = "GOOGLE_CLOUD_PROJECT_ID"

	// AppName is the name of "this server"
	AppName = "api-gateway"

	// AppVersion is the app version (used for StackDriver error reporting)
	AppVersion = "0.0.1"
)
