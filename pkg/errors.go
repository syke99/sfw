package pkg

import "errors"

var (
	OpeningWebFileError   = errors.New("error opening web file")
	ReadWebFileError      = errors.New("error reading web file contents")
	WebUnmarshalError     = errors.New("error while unmarshalling web")
	OpenLinesDirError     = errors.New("error opening lines directory")
	ReadLinesEntriesError = errors.New("error reading lines entries")
	ReadLineError         = errors.New("error reading line")
	LineUnmarshalError    = errors.New("error while unmarshalling line")
	BuildLinesError       = errors.New("error building lines")
	OpenKnotsDirError     = errors.New("error opening knots directory")
	ReadKnotsEntriesError = errors.New("error reading knots entries")
	ReadKnotError         = errors.New("error reading knot file contents")
	KnotUnmarshalError    = errors.New("error while unmarshalling knot")
	BuildKnotsError       = errors.New("error building knots")
)
