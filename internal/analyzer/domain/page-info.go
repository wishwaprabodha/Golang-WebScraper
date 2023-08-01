package domain

type PageInfo struct {
	HTMLVersion       string
	PageTitle         string
	HeadingCounts     map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
	Error             error
}
