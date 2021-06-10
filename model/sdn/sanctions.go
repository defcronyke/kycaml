package sdn

import "encoding/xml"

/** HTTP Route: GET /sdn */
type Sanctions struct {
	XMLName              xml.Name
	Version              string                 `xml:",attr"`
	SchemaLocation       string                 `xml:"schemaLocation,attr"`
	DateOfIssue          DateOfIssueSA          `json:",omitempty"`
	ReferenceValueSets   ReferenceValueSetsSA   `json:",omitempty"`
	Locations            LocationsSA            `json:",omitempty"`
	IDRegDocuments       IDRegDocumentsSA       `json:",omitempty"`
	DistinctParties      DistinctPartiesSA      `json:",omitempty"`
	ProfileRelationships ProfileRelationshipsSA `json:",omitempty"`
	SanctionsEntries     SanctionsEntriesSA     `json:",omitempty"`
}

type DateOfIssueSA struct {
	XMLName        xml.Name
	CalendarTypeID string `xml:",attr"`
	Year           string `json:",omitempty"`
	Month          string `json:",omitempty"`
	Day            string `json:",omitempty"`
}

type ReferenceValueSetsSA struct {
	XMLName                       xml.Name
	AliasTypeValues               AliasTypeValuesSA               `json:",omitempty"`
	AreaCodeValues                AreaCodeValuesSA                `json:",omitempty"`
	AreaCodeTypeValues            AreaCodeTypeValuesSA            `json:",omitempty"`
	CalendarTypeValues            CalendarTypeValuesSA            `json:",omitempty"`
	CountryValues                 CountryValuesSA                 `json:",omitempty"`
	CountryRelevanceValues        CountryRelevanceValuesSA        `json:",omitempty"`
	DecisionMakingBodyValues      DecisionMakingBodyValuesSA      `json:",omitempty"`
	DetailReferenceValues         DetailReferenceValuesSA         `json:",omitempty"`
	DetailTypeValues              DetailTypeValuesSA              `json:",omitempty"`
	DocNameStatusValues           DocNameStatusValuesSA           `json:",omitempty"`
	EntryEventTypeValues          EntryEventTypeValuesSA          `json:",omitempty"`
	EntryLinkTypeValues           string                          `json:",omitempty"`
	ExRefTypeValues               string                          `json:",omitempty"`
	FeatureTypeValues             FeatureTypeValuesSA             `json:",omitempty"`
	FeatureTypeGroupValues        FeatureTypeGroupValuesSA        `json:",omitempty"`
	IDRegDocDateTypeValues        IDRegDocDateTypeValuesSA        `json:",omitempty"`
	IDRegDocTypeValues            IDRegDocTypeValuesSA            `json:",omitempty"`
	IdentityFeatureLinkTypeValues IdentityFeatureLinkTypeValuesSA `json:",omitempty"`
	LegalBasisValues              LegalBasisValuesSA              `json:",omitempty"`
	LegalBasisTypeValues          LegalBasisTypeValuesSA          `json:",omitempty"`
	ListValues                    ListValuesSA                    `json:",omitempty"`
	LocPartTypeValues             LocPartTypeValuesSA             `json:",omitempty"`
	LocPartValueStatusValues      LocPartValueStatusValuesSA      `json:",omitempty"`
	LocPartValueTypeValues        LocPartValueTypeValuesSA        `json:",omitempty"`
	NamePartTypeValues            NamePartTypeValuesSA            `json:",omitempty"`
	OrganisationValues            OrganisationValuesSA            `json:",omitempty"`
	PartySubTypeValues            PartySubTypeValuesSA            `json:",omitempty"`
	PartyTypeValues               PartyTypeValuesSA               `json:",omitempty"`
	RelationQualityValues         RelationQualityValuesSA         `json:",omitempty"`
	RelationTypeValues            RelationTypeValuesSA            `json:",omitempty"`
	ReliabilityValues             ReliabilityValuesSA             `json:",omitempty"`
	SanctionsProgramValues        SanctionsProgramValuesSA        `json:",omitempty"`
	SanctionsTypeValues           SanctionsTypeValuesSA           `json:",omitempty"`
	ScriptValues                  ScriptValuesSA                  `json:",omitempty"`
	ScriptStatusValues            ScriptStatusValuesSA            `json:",omitempty"`
	SubsidiaryBodyValues          SubsidiaryBodyValuesSA          `json:",omitempty"`
	SupInfoTypeValues             string                          `json:",omitempty"`
	TargetTypeValues              string                          `json:",omitempty"`
	ValidityValues                ValidityValuesSA                `json:",omitempty"`
}

type AliasTypeValuesSA struct {
	XMLName   xml.Name
	AliasType []AliasTypeSA
}

type AliasTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type AreaCodeValuesSA struct {
	XMLName  xml.Name
	AreaCode []AreaCodeSA `json:",omitempty"`
}

type AreaCodeSA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	CountryID      string `xml:",attr"`
	Description    string `xml:",attr"`
	AreaCodeTypeID string `xml:",attr"`
	Value          string `xml:",innerxml"`
}

type AreaCodeTypeValuesSA struct {
	XMLName      xml.Name
	AreaCodeType []AreaCodeTypeSA `json:",omitempty"`
}

type AreaCodeTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type CalendarTypeValuesSA struct {
	XMLName      xml.Name
	CalendarType []CalendarTypeSA `json:",omitempty"`
}

type CalendarTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type CountryValuesSA struct {
	XMLName xml.Name
	Country []CountrySA `json:",omitempty"`
}

type CountrySA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	ISO2    string `xml:",attr"`
	Value   string `xml:",innerxml"`
}
type CountryRelevanceValuesSA struct {
	XMLName          xml.Name
	CountryRelevance []CountryRelevanceSA `json:",omitempty"`
}

type CountryRelevanceSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DecisionMakingBodyValuesSA struct {
	XMLName            xml.Name
	DecisionMakingBody []DecisionMakingBodySA `json:",omitempty"`
}

type DecisionMakingBodySA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	OrganisationID string `xml:",attr"`
	Value          string `xml:",innerxml"`
}

type DetailReferenceValuesSA struct {
	XMLName         xml.Name
	DetailReference []DetailReferenceSA `json:",omitempty"`
}

type DetailReferenceSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DetailTypeValuesSA struct {
	XMLName    xml.Name
	DetailType []DetailTypeSA `json:",omitempty"`
}

type DetailTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type DocNameStatusValuesSA struct {
	XMLName       xml.Name
	DocNameStatus []DocNameStatusSA `json:",omitempty"`
}

type DocNameStatusSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type EntryEventTypeValuesSA struct {
	XMLName        xml.Name
	EntryEventType []EntryEventTypeSA `json:",omitempty"`
}

type EntryEventTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type FeatureTypeValuesSA struct {
	XMLName     xml.Name
	FeatureType []FeatureTypeSA `json:",omitempty"`
}

type FeatureTypeSA struct {
	XMLName            xml.Name
	ID                 string `xml:",attr"`
	FeatureTypeGroupID string `xml:",attr"`
	Value              string `xml:",innerxml"`
}

type FeatureTypeGroupValuesSA struct {
	XMLName          xml.Name
	FeatureTypeGroup []FeatureTypeGroupSA `json:",omitempty"`
}

type FeatureTypeGroupSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IDRegDocDateTypeValuesSA struct {
	XMLName          xml.Name
	IDRegDocDateType []IDRegDocDateTypeSA `json:",omitempty"`
}

type IDRegDocDateTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IDRegDocTypeValuesSA struct {
	XMLName      xml.Name
	IDRegDocType []IDRegDocTypeSA `json:",omitempty"`
}

type IDRegDocTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type IdentityFeatureLinkTypeValuesSA struct {
	XMLName                 xml.Name
	IdentityFeatureLinkType []IdentityFeatureLinkTypeSA `json:",omitempty"`
}

type IdentityFeatureLinkTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LegalBasisValuesSA struct {
	XMLName    xml.Name
	LegalBasis []LegalBasisSA `json:",omitempty"`
}

type LegalBasisSA struct {
	XMLName            xml.Name
	ID                 string `xml:",attr"`
	LegalBasisShortRef string `xml:",attr"`
	LegalBasisTypeID   string `xml:",attr"`
	SanctionsProgramID string `xml:",attr"`
	Value              string `xml:",innerxml"`
}

type LegalBasisTypeValuesSA struct {
	XMLName        xml.Name
	LegalBasisType []LegalBasisTypeSA `json:",omitempty"`
}

type LegalBasisTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type ListValuesSA struct {
	XMLName xml.Name
	List    []ListSA `json:",omitempty"`
}

type ListSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartTypeValuesSA struct {
	XMLName     xml.Name
	LocPartType []LocPartTypeSA `json:",omitempty"`
}

type LocPartTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartValueStatusValuesSA struct {
	XMLName            xml.Name
	LocPartValueStatus []LocPartValueStatusSA `json:",omitempty"`
}

type LocPartValueStatusSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocPartValueTypeValuesSA struct {
	XMLName          xml.Name
	LocPartValueType []LocPartValueTypeSA `json:",omitempty"`
}

type LocPartValueTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type NamePartTypeValuesSA struct {
	XMLName      xml.Name
	NamePartType []NamePartTypeSA `json:",omitempty"`
}

type NamePartTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type OrganisationValuesSA struct {
	XMLName      xml.Name
	Organisation []OrganisationSA `json:",omitempty"`
}

type OrganisationSA struct {
	XMLName   xml.Name
	ID        string `xml:",attr"`
	CountryID string `xml:",attr"`
	Value     string `xml:",innerxml"`
}

type PartySubTypeValuesSA struct {
	XMLName      xml.Name
	PartySubType []PartySubTypeSA `json:",omitempty"`
}

type PartySubTypeSA struct {
	XMLName     xml.Name
	ID          string `xml:",attr"`
	PartyTypeID string `xml:",attr"`
	Value       string `xml:",innerxml"`
}

type PartyTypeValuesSA struct {
	XMLName   xml.Name
	PartyType []PartyTypeSA `json:",omitempty"`
}

type PartyTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type RelationQualityValuesSA struct {
	XMLName         xml.Name
	RelationQuality []RelationQualitySA `json:",omitempty"`
}

type RelationQualitySA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type RelationTypeValuesSA struct {
	XMLName      xml.Name
	RelationType []RelationTypeSA `json:",omitempty"`
}

type RelationTypeSA struct {
	XMLName     xml.Name
	ID          string `xml:",attr"`
	Symmetrical string `xml:",attr"`
	Value       string `xml:",innerxml"`
}

type ReliabilityValuesSA struct {
	XMLName     xml.Name
	Reliability []ReliabilitySA `json:",omitempty"`
}

type ReliabilitySA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type SanctionsProgramValuesSA struct {
	XMLName          xml.Name
	SanctionsProgram []SanctionsProgramSA `json:",omitempty"`
}

type SanctionsProgramSA struct {
	XMLName          xml.Name
	ID               string `xml:",attr"`
	SubsidiaryBodyID string `xml:",attr"`
	Value            string `xml:",innerxml"`
}

type SanctionsTypeValuesSA struct {
	XMLName       xml.Name
	SanctionsType []SanctionsTypeSA `json:",omitempty"`
}

type SanctionsTypeSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type ScriptValuesSA struct {
	XMLName xml.Name
	Script  []ScriptSA `json:",omitempty"`
}

type ScriptSA struct {
	XMLName    xml.Name
	ID         string `xml:",attr"`
	ScriptCode string `xml:",attr"`
	Value      string `xml:",innerxml"`
}

type ScriptStatusValuesSA struct {
	XMLName      xml.Name
	ScriptStatus []ScriptStatusSA `json:",omitempty"`
}

type ScriptStatusSA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type SubsidiaryBodyValuesSA struct {
	XMLName        xml.Name
	SubsidiaryBody []SubsidiaryBodySA `json:",omitempty"`
}

type SubsidiaryBodySA struct {
	XMLName              xml.Name
	ID                   string `xml:",attr"`
	Notional             string `xml:",attr"`
	DecisionMakingBodyID string `xml:",attr"`
	Value                string `xml:",innerxml"`
}

type ValidityValuesSA struct {
	XMLName  xml.Name
	Validity []ValiditySA `json:",omitempty"`
}

type ValiditySA struct {
	XMLName xml.Name
	ID      string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type LocationsSA struct {
	XMLName  xml.Name
	Location []LocationSA `json:",omitempty"`
}

type LocationSA struct {
	XMLName                 xml.Name
	ID                      string                    `xml:",attr"`
	LocationAreaCode        LocationAreaCodeSA        `json:",omitempty"`
	FeatureVersionReference FeatureVersionReferenceSA `json:",omitempty"`
	LocationCountry         LocationCountrySA         `json:",omitempty"`
	LocationPart            []LocationPartSA          `json:",omitempty"`
}

type LocationAreaCodeSA struct {
	XMLName    xml.Name
	AreaCodeID string `xml:",attr"`
}

type FeatureVersionReferenceSA struct {
	XMLName          xml.Name
	FeatureVersionID string `xml:",attr"`
}

type LocationCountrySA struct {
	XMLName            xml.Name
	CountryID          string `xml:",attr"`
	CountryRelevanceID string `xml:",attr"`
}

type LocationPartSA struct {
	XMLName           xml.Name
	LocPartTypeID     string                `xml:",attr"`
	LocationPartValue []LocationPartValueSA `json:",omitempty"`
}

type LocationPartValueSA struct {
	XMLName              xml.Name
	Primary              string `xml:",attr"`
	LocPartValueTypeID   string `xml:",attr"`
	LocPartValueStatusID string `xml:",attr"`
	Comment              string `json:",omitempty"`
	Value                string `json:",omitempty"`
}

type IDRegDocumentsSA struct {
	XMLName       xml.Name
	IDRegDocument []IDRegDocumentSA `json:",omitempty"`
}

type IDRegDocumentSA struct {
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
	DocumentedNameReference DocumentedNameReferenceSA `json:",omitempty"`
}

type DocumentedNameReferenceSA struct {
	XMLName          xml.Name
	DocumentedNameID string `xml:",attr"`
}

type DistinctPartiesSA struct {
	XMLName       xml.Name
	DistinctParty []DistinctPartySA `json:",omitempty"`
}

type DistinctPartySA struct {
	XMLName  xml.Name
	FixedRef string      `xml:",attr"`
	Comment  string      `json:",omitempty"`
	Profile  []ProfileSA `json:",omitempty"`
}

type ProfileSA struct {
	XMLName        xml.Name
	ID             string       `xml:",attr"`
	PartySubTypeID string       `xml:",attr"`
	Identity       []IdentitySA `json:",omitempty"`
	Feature        []FeatureSA  `json:",omitempty"`
}

type IdentitySA struct {
	XMLName        xml.Name
	ID             string             `xml:",attr"`
	FixedRef       string             `xml:",attr"`
	Primary        string             `xml:",attr"`
	False          string             `xml:",attr"`
	Alias          []AliasSA          `json:",omitempty"`
	NamePartGroups []NamePartGroupsSA `json:",omitempty"`
}

type AliasSA struct {
	XMLName        xml.Name
	FixedRef       string             `xml:",attr"`
	AliasTypeID    string             `xml:",attr"`
	Primary        string             `xml:",attr"`
	LowQuality     string             `xml:",attr"`
	DocumentedName []DocumentedNameSA `json:",omitempty"`
}

type DocumentedNameSA struct {
	XMLName            xml.Name
	ID                 string                 `xml:",attr"`
	FixedRef           string                 `xml:",attr"`
	DocNameStatusID    string                 `xml:",attr"`
	DocumentedNamePart []DocumentedNamePartSA `json:",omitempty"`
}

type DocumentedNamePartSA struct {
	XMLName       xml.Name
	NamePartValue []NamePartValueSA `json:",omitempty"`
}

type NamePartValueSA struct {
	XMLName         xml.Name
	NamePartGroupID string `xml:",attr"`
	ScriptID        string `xml:",attr"`
	ScriptStatusID  string `xml:",attr"`
	Acronym         string `xml:",attr"`
	DocNameStatusID string `xml:",attr"`
	Value           string `xml:",innerxml"`
}

type NamePartGroupsSA struct {
	XMLName             xml.Name
	MasterNamePartGroup []MasterNamePartGroupSA `json:",omitempty"`
}

type MasterNamePartGroupSA struct {
	XMLName       xml.Name
	NamePartGroup []NamePartGroupSA `json:",omitempty"`
}

type NamePartGroupSA struct {
	XMLName        xml.Name
	ID             string `xml:",attr"`
	NamePartTypeID string `xml:",attr"`
}

type FeatureSA struct {
	XMLName           xml.Name
	ID                string             `xml:",attr"`
	FeatureTypeID     string             `xml:",attr"`
	FeatureVersion    []FeatureVersionSA `json:",omitempty"`
	IdentityReference string             `json:",omitempty"`
}

type FeatureVersionSA struct {
	XMLName       xml.Name
	ID            string          `xml:",attr"`
	ReliabilityID string          `xml:",attr"`
	Comment       string          `json:",omitempty"`
	DatePeriod    []DatePeriodSA  `json:",omitempty"`
	VersionDetail VersionDetailSA `json:",omitempty"`
}

type DatePeriodSA struct {
	XMLName        xml.Name
	CalendarTypeID string    `xml:",attr"`
	YearFixed      string    `xml:",attr"`
	MonthFixed     string    `xml:",attr"`
	DayFixed       string    `xml:",attr"`
	Start          []StartSA `json:",omitempty"`
	End            []EndSA   `json:",omitempty"`
}

type StartSA struct {
	XMLName     xml.Name
	Approximate string   `xml:",attr"`
	YearFixed   string   `xml:",attr"`
	MonthFixed  string   `xml:",attr"`
	DayFixed    string   `xml:",attr"`
	From        []FromSA `json:",omitempty"`
	To          []ToSA   `json:",omitempty"`
}

type FromSA struct {
	XMLName xml.Name
	Year    string `json:",omitempty"`
	Month   string `json:",omitempty"`
	Day     string `json:",omitempty"`
}

type ToSA struct {
	XMLName xml.Name
	Year    string `json:",omitempty"`
	Month   string `json:",omitempty"`
	Day     string `json:",omitempty"`
}

type EndSA struct {
	XMLName xml.Name
	From    FromSA `json:",omitempty"`
	To      ToSA   `json:",omitempty"`
}

type VersionDetailSA struct {
	XMLName      xml.Name
	DetailTypeID string `xml:",attr"`
}

type ProfileRelationshipsSA struct {
	XMLName             xml.Name
	ProfileRelationship []ProfileRelationshipSA `json:",omitempty"`
}

type ProfileRelationshipSA struct {
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

type SanctionsEntriesSA struct {
	XMLName        xml.Name
	SanctionsEntry []SanctionsEntrySA `json:",omitempty"`
}

type SanctionsEntrySA struct {
	XMLName          xml.Name
	ID               string               `xml:",attr"`
	ProfileID        string               `xml:",attr"`
	ListID           string               `xml:",attr"`
	EntryEvent       []EntryEventSA       `json:",omitempty"`
	SanctionsMeasure []SanctionsMeasureSA `json:",omitempty"`
}

type EntryEventSA struct {
	XMLName          xml.Name
	ID               string   `xml:",attr"`
	EntryEventTypeID string   `xml:",attr"`
	LegalBasisID     string   `xml:",attr"`
	Comment          string   `json:",omitempty"`
	Date             []DateSA `json:",omitempty"`
}

type DateSA struct {
	XMLName        xml.Name
	CalendarTypeID string `xml:",attr"`
	Year           string `json:",omitempty"`
	Month          string `json:",omitempty"`
	Day            string `json:",omitempty"`
}

type SanctionsMeasureSA struct {
	XMLName         xml.Name
	ID              string         `xml:",attr"`
	SanctionsTypeID string         `xml:",attr"`
	Comment         string         `json:",omitempty"`
	DatePeriod      []DatePeriodSA `json:",omitempty"`
}
