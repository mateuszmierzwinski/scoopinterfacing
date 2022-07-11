package scoopinterfacing

import "flag"

/*
Version describes version details, that each plugin can produce
*/
type Version struct {
	VersionStr string
	CPUArch    string
	BuildDate  string
}

/*
ScoopService defines interfacing from the service shared as library, to SCOOP
*/
type ScoopService interface {
	// GetServiceVersion returns version details
	GetServiceVersion() Version

	// GetServiceFlags returns configuration flags
	GetServiceFlags() []flag.Flag

	// Execute executes service with passed SCOOP flags
	Execute([]string) error

	// Returns TXT readme as string, that will be executed on scoop <module> -h
	GetHelp() string
}
