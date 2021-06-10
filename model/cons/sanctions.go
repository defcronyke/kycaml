package cons

import "encoding/xml"

/** HTTP Route: GET /cons */
type Sanctions struct {
	XMLName              xml.Name
	Version              string                 `xml:",attr"`
	SchemaLocation       string                 `xml:"schemaLocation,attr"`
	DateOfIssue          DateOfIssueCA          `json:",omitempty"`
	ReferenceValueSets   ReferenceValueSetsCA   `json:",omitempty"`
	Locations            LocationsCA            `json:",omitempty"`
	IDRegDocuments       IDRegDocumentsCA       `json:",omitempty"`
	DistinctParties      DistinctPartiesCA      `json:",omitempty"`
	ProfileRelationships ProfileRelationshipsCA `json:",omitempty"`
	SanctionsEntries     SanctionsEntriesCA     `json:",omitempty"`
}

type DateOfIssueCA struct {
	XMLName        xml.Name
	CalendarTypeID string `xml:",attr"`
	Year           string `json:",omitempty"`
	Month          string `json:",omitempty"`
	Day            string `json:",omitempty"`
}

type ReferenceValueSetsCA struct {
	XMLName                       xml.Name
	AliasTypeValues               AliasTypeValuesCA               `json:",omitempty"`
	AreaCodeValues                AreaCodeValuesCA                `json:",omitempty"`
	AreaCodeTypeValues            AreaCodeTypeValuesCA            `json:",omitempty"`
	CalendarTypeValues            CalendarTypeValuesCA            `json:",omitempty"`
	CountryValues                 CountryValuesCA                 `json:",omitempty"`
	CountryRelevanceValues        CountryRelevanceValuesCA        `json:",omitempty"`
	DecisionMakingBodyValues      DecisionMakingBodyValuesCA      `json:",omitempty"`
	DetailReferenceValues         DetailReferenceValuesCA         `json:",omitempty"`
	DetailTypeValues              DetailTypeValuesCA              `json:",omitempty"`
	DocNameStatusValues           DocNameStatusValuesCA           `json:",omitempty"`
	EntryEventTypeValues          EntryEventTypeValuesCA          `json:",omitempty"`
	EntryLinkTypeValues           string                          `json:",omitempty"`
	ExRefTypeValues               string                          `json:",omitempty"`
	FeatureTypeValues             FeatureTypeValuesCA             `json:",omitempty"`
	FeatureTypeGroupValues        FeatureTypeGroupValuesCA        `json:",omitempty"`
	IDRegDocDateTypeValues        IDRegDocDateTypeValuesCA        `json:",omitempty"`
	IDRegDocTypeValues            IDRegDocTypeValuesCA            `json:",omitempty"`
	IdentityFeatureLinkTypeValues IdentityFeatureLinkTypeValuesCA `json:",omitempty"`
	LegalBasisValues              LegalBasisValuesCA              `json:",omitempty"`
	LegalBasisTypeValues          LegalBasisTypeValuesCA          `json:",omitempty"`
	ListValues                    ListValuesCA                    `json:",omitempty"`
	LocPartTypeValues             LocPartTypeValuesCA             `json:",omitempty"`
	LocPartValueStatusValues      LocPartValueStatusValuesCA      `json:",omitempty"`
	LocPartValueTypeValues        LocPartValueTypeValuesCA        `json:",omitempty"`
	NamePartTypeValues            NamePartTypeValuesCA            `json:",omitempty"`
	OrganisationValues            OrganisationValuesCA            `json:",omitempty"`
	PartySubTypeValues            PartySubTypeValuesCA            `json:",omitempty"`
	PartyTypeValues               PartyTypeValuesCA               `json:",omitempty"`
	RelationQualityValues         RelationQualityValuesCA         `json:",omitempty"`
	RelationTypeValues            RelationTypeValuesCA            `json:",omitempty"`
	ReliabilityValues             ReliabilityValuesCA             `json:",omitempty"`
	SanctionsProgramValues        SanctionsProgramValuesCA        `json:",omitempty"`
	SanctionsTypeValues           SanctionsTypeValuesCA           `json:",omitempty"`
	ScriptValues                  ScriptValuesCA                  `json:",omitempty"`
	ScriptStatusValues            ScriptStatusValuesCA            `json:",omitempty"`
	SubsidiaryBodyValues          SubsidiaryBodyValuesCA          `json:",omitempty"`
	SupInfoTypeValues             string                          `json:",omitempty"`
	TargetTypeValues              string                          `json:",omitempty"`
	ValidityValues                ValidityValuesCA                `json:",omitempty"`
}

type AliasTypeValuesCA struct {
	XMLName   xml.Name
	AliasType []AliasTypeCA
}

type AliasTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type AreaCodeValuesCA struct {
	XMLName  xml.Name
	AreaCode []AreaCodeCA `json:",omitempty"`
}

type AreaCodeCA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	CountryID      string `xml:",attr"`
	Description    string `xml:",attr"`
	AreaCodeTypeID string `xml:",attr"`
	Value          string `xml:",innerxml"`
}

type AreaCodeTypeValuesCA struct {
	XMLName      xml.Name
	AreaCodeType []AreaCodeTypeCA `json:",omitempty"`
}

type AreaCodeTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type CalendarTypeValuesCA struct {
	XMLName      xml.Name
	CalendarType []CalendarTypeCA `json:",omitempty"`
}

type CalendarTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type CountryValuesCA struct {
	XMLName xml.Name
	Country []CountryCA `json:",omitempty"`
}

type CountryCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	ISO2    string `xml:",attr"`
	Value   string `xml:",innerxml"`
}
type CountryRelevanceValuesCA struct {
	XMLName          xml.Name
	CountryRelevance []CountryRelevanceCA `json:",omitempty"`
}

type CountryRelevanceCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DecisionMakingBodyValuesCA struct {
	XMLName            xml.Name
	DecisionMakingBody []DecisionMakingBodyCA `json:",omitempty"`
}

type DecisionMakingBodyCA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	OrganisationID string `xml:",attr"`
	Value          string `xml:",innerxml"`
}

type DetailReferenceValuesCA struct {
	XMLName         xml.Name
	DetailReference []DetailReferenceCA `json:",omitempty"`
}

type DetailReferenceCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DetailTypeValuesCA struct {
	XMLName    xml.Name
	DetailType []DetailTypeCA `json:",omitempty"`
}

type DetailTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DocNameStatusValuesCA struct {
	XMLName       xml.Name
	DocNameStatus []DocNameStatusCA `json:",omitempty"`
}

type DocNameStatusCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type EntryEventTypeValuesCA struct {
	XMLName        xml.Name
	EntryEventType []EntryEventTypeCA `json:",omitempty"`
}

type EntryEventTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type FeatureTypeValuesCA struct {
	XMLName     xml.Name
	FeatureType []FeatureTypeCA `json:",omitempty"`
}

type FeatureTypeCA struct {
	XMLName            xml.Name
	ID                 string `xml:",attr"`
	FeatureTypeGroupID string `xml:",attr"`
	Value              string `xml:",innerxml"`
}

type FeatureTypeGroupValuesCA struct {
	XMLName          xml.Name
	FeatureTypeGroup []FeatureTypeGroupCA `json:",omitempty"`
}

type FeatureTypeGroupCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IDRegDocDateTypeValuesCA struct {
	XMLName          xml.Name
	IDRegDocDateType []IDRegDocDateTypeCA `json:",omitempty"`
}

type IDRegDocDateTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IDRegDocTypeValuesCA struct {
	XMLName      xml.Name
	IDRegDocType []IDRegDocTypeCA `json:",omitempty"`
}

type IDRegDocTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IdentityFeatureLinkTypeValuesCA struct {
	XMLName                 xml.Name
	IdentityFeatureLinkType []IdentityFeatureLinkTypeCA `json:",omitempty"`
}

type IdentityFeatureLinkTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LegalBasisValuesCA struct {
	XMLName    xml.Name
	LegalBasis []LegalBasisCA `json:",omitempty"`
}

type LegalBasisCA struct {
	XMLName            xml.Name
	ID                 string `xml:",attr"`
	LegalBasisShortRef string `xml:",attr"`
	LegalBasisTypeID   string `xml:",attr"`
	SanctionsProgramID string `xml:",attr"`
	Value              string `xml:",innerxml"`
}

type LegalBasisTypeValuesCA struct {
	XMLName        xml.Name
	LegalBasisType []LegalBasisTypeCA `json:",omitempty"`
}

type LegalBasisTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type ListValuesCA struct {
	XMLName xml.Name
	List    []ListCA `json:",omitempty"`
}

type ListCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartTypeValuesCA struct {
	XMLName     xml.Name
	LocPartType []LocPartTypeCA `json:",omitempty"`
}

type LocPartTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartValueStatusValuesCA struct {
	XMLName            xml.Name
	LocPartValueStatus []LocPartValueStatusCA `json:",omitempty"`
}

type LocPartValueStatusCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartValueTypeValuesCA struct {
	XMLName          xml.Name
	LocPartValueType []LocPartValueTypeCA `json:",omitempty"`
}

type LocPartValueTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type NamePartTypeValuesCA struct {
	XMLName      xml.Name
	NamePartType []NamePartTypeCA `json:",omitempty"`
}

type NamePartTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type OrganisationValuesCA struct {
	XMLName      xml.Name
	Organisation []OrganisationCA `json:",omitempty"`
}

type OrganisationCA struct {
	XMLName   xml.Name
	ID        string `xml:",attr"`
	CountryID string `xml:",attr"`
	Value     string `xml:",innerxml"`
}

type PartySubTypeValuesCA struct {
	XMLName      xml.Name
	PartySubType []PartySubTypeCA `json:",omitempty"`
}

type PartySubTypeCA struct {
	XMLName     xml.Name
	ID          string `xml:",attr"`
	PartyTypeID string `xml:",attr"`
	Value       string `xml:",innerxml"`
}

type PartyTypeValuesCA struct {
	XMLName   xml.Name
	PartyType []PartyTypeCA `json:",omitempty"`
}

type PartyTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type RelationQualityValuesCA struct {
	XMLName         xml.Name
	RelationQuality []RelationQualityCA `json:",omitempty"`
}

type RelationQualityCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type RelationTypeValuesCA struct {
	XMLName      xml.Name
	RelationType []RelationTypeCA `json:",omitempty"`
}

type RelationTypeCA struct {
	XMLName     xml.Name
	ID          string `xml:",attr"`
	Symmetrical string `xml:",attr"`
	Value       string `xml:",innerxml"`
}

type ReliabilityValuesCA struct {
	XMLName     xml.Name
	Reliability []ReliabilityCA `json:",omitempty"`
}

type ReliabilityCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type SanctionsProgramValuesCA struct {
	XMLName          xml.Name
	SanctionsProgram []SanctionsProgramCA `json:",omitempty"`
}

type SanctionsProgramCA struct {
	XMLName          xml.Name
	ID               string `xml:",attr"`
	SubsidiaryBodyID string `xml:",attr"`
	Value            string `xml:",innerxml"`
}

type SanctionsTypeValuesCA struct {
	XMLName       xml.Name
	SanctionsType []SanctionsTypeCA `json:",omitempty"`
}

type SanctionsTypeCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type ScriptValuesCA struct {
	XMLName xml.Name
	Script  []ScriptCA `json:",omitempty"`
}

type ScriptCA struct {
	XMLName    xml.Name
	ID         string `xml:",attr"`
	ScriptCode string `xml:",attr"`
	Value      string `xml:",innerxml"`
}

type ScriptStatusValuesCA struct {
	XMLName      xml.Name
	ScriptStatus []ScriptStatusCA `json:",omitempty"`
}

type ScriptStatusCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type SubsidiaryBodyValuesCA struct {
	XMLName        xml.Name
	SubsidiaryBody []SubsidiaryBodyCA `json:",omitempty"`
}

type SubsidiaryBodyCA struct {
	XMLName              xml.Name
	ID                   string `xml:",attr"`
	Notional             string `xml:",attr"`
	DecisionMakingBodyID string `xml:",attr"`
	Value                string `xml:",innerxml"`
}

type ValidityValuesCA struct {
	XMLName  xml.Name
	Validity []ValidityCA `json:",omitempty"`
}

type ValidityCA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocationsCA struct {
	XMLName  xml.Name
	Location []LocationCA `json:",omitempty"`
}

type LocationCA struct {
	XMLName                 xml.Name
	ID                      string                    `xml:",attr"`
	LocationAreaCode        LocationAreaCodeCA        `json:",omitempty"`
	FeatureVersionReference FeatureVersionReferenceCA `json:",omitempty"`
	LocationCountry         LocationCountryCA         `json:",omitempty"`
	LocationPart            []LocationPartCA          `json:",omitempty"`
}

type LocationAreaCodeCA struct {
	XMLName    xml.Name
	AreaCodeID string `xml:",attr"`
}

type FeatureVersionReferenceCA struct {
	XMLName          xml.Name
	FeatureVersionID string `xml:",attr"`
}

type LocationCountryCA struct {
	XMLName            xml.Name
	CountryID          string `xml:",attr"`
	CountryRelevanceID string `xml:",attr"`
}

type LocationPartCA struct {
	XMLName           xml.Name
	LocPartTypeID     string                `xml:",attr"`
	LocationPartValue []LocationPartValueCA `json:",omitempty"`
}

type LocationPartValueCA struct {
	XMLName              xml.Name
	Primary              string `xml:",attr"`
	LocPartValueTypeID   string `xml:",attr"`
	LocPartValueStatusID string `xml:",attr"`
	Comment              string `json:",omitempty"`
	Value                string `json:",omitempty"`
}

type IDRegDocumentsCA struct {
	XMLName       xml.Name
	IDRegDocument []IDRegDocumentCA `json:",omitempty"`
}

type IDRegDocumentCA struct {
	XMLName                 xml.Name
	ID                      string                    `xml:",attr"`
	IDRegDocTypeID          string                    `xml:",attr"`
	IdentityID              string                    `xml:",attr"`
	IssuedByCountryID       string                    `xml:"IssuedBy-CountryID,attr"`
	IssuedInLocationID      string                    `xml:"IssuedIn-LocationID,attr"`
	ValidityID              string                    `xml:",attr"`
	Comment                 string                    `json:",omitempty"`
	IDRegistrationNo        string                    `json:",omitempty"`
	IssuingAuthority        string                    `json:",omitempty"`
	DocumentedNameReference DocumentedNameReferenceCA `json:",omitempty"`
}

type DocumentedNameReferenceCA struct {
	XMLName          xml.Name
	DocumentedNameID string `xml:",attr"`
}

type DistinctPartiesCA struct {
	XMLName       xml.Name
	DistinctParty []DistinctPartyCA `json:",omitempty"`
}

type DistinctPartyCA struct {
	XMLName  xml.Name
	FixedRef string      `xml:",attr"`
	Comment  string      `json:",omitempty"`
	Profile  []ProfileCA `json:",omitempty"`
}

type ProfileCA struct {
	XMLName        xml.Name
	ID             string       `xml:",attr"`
	PartySubTypeID string       `xml:",attr"`
	Identity       []IdentityCA `json:",omitempty"`
	Feature        []FeatureCA  `json:",omitempty"`
}

type IdentityCA struct {
	XMLName        xml.Name
	ID             string             `xml:",attr"`
	FixedRef       string             `xml:",attr"`
	Primary        string             `xml:",attr"`
	False          string             `xml:",attr"`
	Alias          []AliasCA          `json:",omitempty"`
	NamePartGroups []NamePartGroupsCA `json:",omitempty"`
}

type AliasCA struct {
	XMLName        xml.Name
	FixedRef       string             `xml:",attr"`
	AliasTypeID    string             `xml:",attr"`
	Primary        string             `xml:",attr"`
	LowQuality     string             `xml:",attr"`
	DocumentedName []DocumentedNameCA `json:",omitempty"`
}

type DocumentedNameCA struct {
	XMLName            xml.Name
	ID                 string                 `xml:",attr"`
	FixedRef           string                 `xml:",attr"`
	DocNameStatusID    string                 `xml:",attr"`
	DocumentedNamePart []DocumentedNamePartCA `json:",omitempty"`
}

type DocumentedNamePartCA struct {
	XMLName       xml.Name
	NamePartValue []NamePartValueCA `json:",omitempty"`
}

type NamePartValueCA struct {
	XMLName         xml.Name
	NamePartGroupID string `xml:",attr"`
	ScriptID        string `xml:",attr"`
	ScriptStatusID  string `xml:",attr"`
	Acronym         string `xml:",attr"`
	DocNameStatusID string `xml:",attr"`
	Value           string `xml:",innerxml"`
}

type NamePartGroupsCA struct {
	XMLName             xml.Name
	MasterNamePartGroup []MasterNamePartGroupCA `json:",omitempty"`
}

type MasterNamePartGroupCA struct {
	XMLName       xml.Name
	NamePartGroup []NamePartGroupCA `json:",omitempty"`
}

type NamePartGroupCA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	NamePartTypeID string `xml:",attr"`
}

type FeatureCA struct {
	XMLName           xml.Name
	ID                string             `xml:",attr"`
	FeatureTypeID     string             `xml:",attr"`
	FeatureVersion    []FeatureVersionCA `json:",omitempty"`
	IdentityReference string             `json:",omitempty"`
}

type FeatureVersionCA struct {
	XMLName       xml.Name
	ID            string          `xml:",attr"`
	ReliabilityID string          `xml:",attr"`
	Comment       string          `json:",omitempty"`
	DatePeriod    []DatePeriodCA  `json:",omitempty"`
	VersionDetail VersionDetailCA `json:",omitempty"`
}

type DatePeriodCA struct {
	XMLName        xml.Name
	CalendarTypeID string    `xml:",attr"`
	YearFixed      string    `xml:",attr"`
	MonthFixed     string    `xml:",attr"`
	DayFixed       string    `xml:",attr"`
	Start          []StartCA `json:",omitempty"`
	End            []EndCA   `json:",omitempty"`
}

type StartCA struct {
	XMLName     xml.Name
	Approximate string   `xml:",attr"`
	YearFixed   string   `xml:",attr"`
	MonthFixed  string   `xml:",attr"`
	DayFixed    string   `xml:",attr"`
	From        []FromCA `json:",omitempty"`
	To          []ToCA   `json:",omitempty"`
}

type FromCA struct {
	XMLName xml.Name
	Year    string `json:",omitempty"`
	Month   string `json:",omitempty"`
	Day     string `json:",omitempty"`
}

type ToCA struct {
	XMLName xml.Name
	Year    string `json:",omitempty"`
	Month   string `json:",omitempty"`
	Day     string `json:",omitempty"`
}

type EndCA struct {
	XMLName xml.Name
	From    FromCA `json:",omitempty"`
	To      ToCA   `json:",omitempty"`
}

type VersionDetailCA struct {
	XMLName      xml.Name
	DetailTypeID string `xml:",attr"`
}

type ProfileRelationshipsCA struct {
	XMLName             xml.Name
	ProfileRelationship []ProfileRelationshipCA `json:",omitempty"`
}

type ProfileRelationshipCA struct {
	XMLName           xml.Name
	ID                string `xml:",attr"`
	FromProfileID     string `xml:"From-ProfileID,attr"`
	ToProfileID       string `xml:"To-ProfileID,attr"`
	RelationTypeID    string `xml:",attr"`
	RelationQualityID string `xml:",attr"`
	Former            string `xml:",attr"`
	SanctionsEntryID  string `xml:",attr"`
	Comment           string `json:",omitempty"`
}

type SanctionsEntriesCA struct {
	XMLName        xml.Name
	SanctionsEntry []SanctionsEntryCA `json:",omitempty"`
}

type SanctionsEntryCA struct {
	XMLName          xml.Name
	ID               string               `xml:",attr"`
	ProfileID        string               `xml:",attr"`
	ListID           string               `xml:",attr"`
	EntryEvent       []EntryEventCA       `json:",omitempty"`
	SanctionsMeasure []SanctionsMeasureCA `json:",omitempty"`
}

type EntryEventCA struct {
	XMLName          xml.Name
	ID               string   `xml:",attr"`
	EntryEventTypeID string   `xml:",attr"`
	LegalBasisID     string   `xml:",attr"`
	Comment          string   `json:",omitempty"`
	Date             []DateCA `json:",omitempty"`
}

type DateCA struct {
	XMLName        xml.Name
	CalendarTypeID string `xml:",attr"`
	Year           string `json:",omitempty"`
	Month          string `json:",omitempty"`
	Day            string `json:",omitempty"`
}

type SanctionsMeasureCA struct {
	XMLName         xml.Name
	ID              string         `xml:",attr"`
	SanctionsTypeID string         `xml:",attr"`
	Comment         string         `json:",omitempty"`
	DatePeriod      []DatePeriodCA `json:",omitempty"`
}
