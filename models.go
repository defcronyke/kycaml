package kycaml

/** GET /static/xml/cons_advanced.xml */
type ConsAdvanced struct {
	Sanctions SanctionsCA
}

type SanctionsCA struct {
	DateOfIssue DateOfIssueCA
}

type DateOfIssueCA struct {
	Year  uint64
	Month uint8
	Day   uint8
	Eggs  string
}
