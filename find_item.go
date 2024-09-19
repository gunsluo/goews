package goews

import "encoding/xml"

type FindItem struct {
	XMLName xml.Name `xml:"m:FindItem"`

	// The CalendarView element defines a FindItem operation as returning calendar items in a set as they appear in a calendar.
	CalendarView *CalendarView `xml:"m:CalendarView,omitempty"`
	// The ContactsView element defines a search for contact items based on alphabetical display names.
	ContactsView *ContactsView `xml:"m:ContactsView,omitempty"`
	// The DistinguishedGroupBy element provides standard groupings for FindItem queries.
	DistinguishedGroupBy *DistinguishedGroupBy `xml:"m:DistinguishedGroupBy,omitempty"`
	// The FractionalPageItemView element describes where the paged view starts and the maximum number of items returned in a FindItem request.
	FractionalPageItemView *FractionalPageItemView `xml:"m:FractionalPageItemView,omitempty"`
	// The GroupBy element specifies an arbitrary grouping for FindItem queries.
	GroupBy *GroupBy `xml:"m:GroupBy,omitempty"`
	// The IndexedPageItemView element describes how paged conversation or item information is returned for a FindItem operation or FindConversation operation request.
	IndexedPageItemView *IndexedPageItemView `xml:"m:IndexedPageItemView,omitempty"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"m:ItemShape,omitempty"`
	// The ParentFolderIds element identifies folders for the FindItem and FindFolder operations to search.
	ParentFolderIds *ParentFolderIds `xml:"m:ParentFolderIds,omitempty"`
	// The QueryString element contains a mailbox query string based on Advanced Query Syntax (AQS).
	QueryString *QueryStringQueryStringType `xml:"m:QueryString,omitempty"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"m:Restriction,omitempty"`
	// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
	SortOrder *SortOrder `xml:"m:SortOrder,omitempty"`
	// Defines whether the search finds items in folders or the folders' dumpsters. This attribute is required.
	Traversal string `xml:"Traversal,attr"`
}

const (
	// Returns only the identities of items in the folder.
	FindItemShallow = `Shallow`
	// Returns only the identities of items that are in a folder's dumpster. Note that a soft-deleted traversal combined with a search restriction will result in zero items returned even if there are items that match the search criteria.
	FindItemSoftDeleted = `SoftDeleted`
	// Returns only the identities of associated items in the folder.
	FindItemAssociated = `Associated`
)

type CalendarView struct {
	// Identifies the end of a time span queried for calendar items. All calendar items that have a start time that is on or after EndDate will not be returned. The value of EndDate can be specified in UTC format, as in 2006-02-02T12:00:00Z, or in a format where local time and time zone offset is specified, as in 2006-02-02T04:00:00-08:00.  EndDate must be greater than or equal to StartDate; otherwise an error is returned. This attribute is required.
	EndDate *string `xml:"EndDate,attr"`
	// Describes the maximum number of results to return in the FindItem response.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Identifies the start of a time span queried for calendar items. All calendar items that have an end time that is before StartDate will not be returned. The value of StartDate can be specified in coordinated universal time (UTC) format, as in 2006-01-02T12:00:00Z, or in a format where local time and time zone offset is specified, as in 2006-01-02T04:00:00-08:00.  This attribute is required.
	StartDate *string `xml:"StartDate,attr"`
}

type ContactsView struct {
	// Defines the last name in the contacts list to return in the response. If the FinalName attribute is omitted, the response will contain all subsequent contacts in the specified sort order. If the specified final name is not in the contacts list, the next alphabetical name as defined by the cultural context will be excluded.  For example, if FinalName="Name", but Name is not in the contacts list, contacts that have display names of Name1 or NAME will not be included.  This attribute is optional.
	FinalName *string `xml:"FinalName,attr"`
	// Defines the first name in the contacts list to return in the response. If the specified initial name is not in the contacts list, the next alphabetical name as defined by the cultural context will be returned, except if the next name comes after FinalName. If the InitialName attribute is omitted, the response will contain a list of contacts that starts with the first name in the contact list. This attribute is optional.
	InitialName *string `xml:"InitialName,attr"`
	// Describes the maximum number of results to return in the FindItem response.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
}

type DistinguishedGroupBy struct {
	// The StandardGroupBy element represents the standard grouping and aggregating mechanisms for a grouped FindItem operation.
	StandardGroupBy *StandardGroupBy `xml:"t:StandardGroupBy,omitempty"`
}

type StandardGroupBy string

type FractionalPageItemView struct {
	// Represents the denominator of the fractional offset from the start of the total number of items in the result set. This attribute is required. This attribute must represent an integral value that is greater than one.   For more information, see Remarks later in this topic.
	Denominator *string `xml:"Denominator,attr"`
	// Identifies the maximum number of results to return in the FindItem response. This attribute is optional. If this attribute is not specified, the call will return all available items.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Represents the numerator of the fractional offset from the start of the result set. This attribute is required. The numerator must be equal to or less than the denominator. This attribute must represent an integral value that is equal to or greater than zero.   For more information, see Remarks later in this topic.
	Numerator *string `xml:"Numerator,attr"`
}

type GroupBy struct {
	// The AggregateOn element represents the property that is used to determine the order of grouped items for a grouped FindItem result set.
	AggregateOn *AggregateOn `xml:"t:AggregateOn,omitempty"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
	// Determines the order of the groups in the grouped item array that is returned in the response. This attribute is of type SortDirectionType.
	Order string `xml:"Order,attr"`
}

const (
	// The groups are ordered in ascending order.
	GroupByAscending = `Ascending`
	// The groups are ordered in descending order.
	GroupByDescending = `Descending`
)

type AggregateOn struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
	// Indicates the maximum or minimum value of the property identified by the FieldURI element that is used for ordering the groups of items.The following are the possible values:  - Minimum  - Maximum
	Aggregate string `xml:"Aggregate,attr"`
}

type IndexedFieldURI struct {
	// Identifies the member of the dictionary to return. This attribute is required.
	FieldIndex *string `xml:"FieldIndex,attr"`
	// Identifies the dictionary that contains the member to return. This attribute is required.
	FieldURI *string `xml:"FieldURI,attr"`
}

type ItemShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties,omitempty"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape BaseShape `xml:"t:BaseShape,omitempty"`
	// The BodyType element identifies how the body text is formatted in the response.
	BodyType BodyType `xml:"t:BodyType,omitempty"`
	// The ConvertHtmlCodePageToUTF8 element indicates whether the item HTML body is converted to UTF8.
	ConvertHtmlCodePageToUTF8 ConvertHtmlCodePageToUTF8 `xml:"t:ConvertHtmlCodePageToUTF8,omitempty"`
	// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
	FilterHtmlContent FilterHtmlContent `xml:"t:FilterHtmlContent,omitempty"`
	// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
	IncludeMimeContent IncludeMimeContent `xml:"t:IncludeMimeContent,omitempty"`
}

type BodyType string

const (
	// The response will return the richest available content of body text. This is useful if it is unknown whether the content is text or HTML. The returned body will be text if the stored body is plain text. Otherwise, the response will return HTML if the stored body is in either HTML or RTF format. This is the default value.
	BodyTypeBest string = `Best`
	// The response will return an item body as HTML.
	BodyTypeHTML string = `HTML`
	// The response will return an item body as plain text.
	BodyTypeText string = `Text`
)

type ConvertHtmlCodePageToUTF8 bool

type FilterHtmlContent bool

const (
	// false
	FilterHtmlContentfalse bool = false
	// true
	FilterHtmlContenttrue bool = true
)

type IncludeMimeContent bool

const (
	// false
	IncludeMimeContentfalse bool = false
	// true
	IncludeMimeContenttrue bool = true
)

type ParentFolderIds struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
}

type QueryStringQueryStringType struct {
	// Indicates that the cache should be reset.
	ResetCache *string `xml:"ResetCache,attr"`
	// Indicates that deleted items should be returned.
	ReturnDeletedItems *string `xml:"ReturnDeletedItems,attr"`
	// Indicates that highlighted terms should be returned.
	ReturnHighlightTerms *string `xml:"ReturnHighlightTerms,attr"`
	TEXT                 string  `xml:",chardata"`
}

type Restriction struct {
	// The And element represents a search expression that allows you to perform a Boolean AND operation between two or more search expressions. The result of the AND operation is true if all the search expressions contained within the And element are true.
	And *And `xml:"t:And,omitempty"`
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"t:Contains,omitempty"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes,omitempty"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"t:Exists,omitempty"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"t:IsEqualTo,omitempty"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"t:IsGreaterThan,omitempty"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"t:IsGreaterThanOrEqualTo,omitempty"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"t:IsLessThan,omitempty"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"t:IsLessThanOrEqualTo,omitempty"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"t:IsNotEqualTo,omitempty"`
	// The Not element represents a search expression that negates the Boolean value of the search expression that it contains.
	Not *Not `xml:"t:Not,omitempty"`
	// The Or element represents a search expression that performs a logical OR on the search expression that it contains. Or will return true if any of its children return true. Or must have two or more children.
	Or *Or `xml:"t:Or,omitempty"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression,omitempty"`
}

type SortOrder struct {
	// The FieldOrder element represents a single field by which to sort results and indicates the direction for the sort.
	FieldOrder *FieldOrder `xml:"t:FieldOrder,omitempty"`
}

type FieldOrder struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
	// Describes the sort order direction. The following are the possible values:  - Ascending  - Descending
	Order string `xml:"Order,attr"`
}

type And struct {
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"t:Contains,omitempty"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes,omitempty"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"t:Exists,omitempty"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"t:IsEqualTo,omitempty"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"t:IsGreaterThan,omitempty"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"t:IsGreaterThanOrEqualTo,omitempty"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"t:IsLessThan,omitempty"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"t:IsLessThanOrEqualTo,omitempty"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"t:IsNotEqualTo,omitempty"`
	// The Not element represents a search expression that negates the Boolean value of the search expression that it contains.
	Not *Not `xml:"t:Not,omitempty"`
	// The Or element represents a search expression that performs a logical OR on the search expression that it contains. Or will return true if any of its children return true. Or must have two or more children.
	Or *Or `xml:"t:Or,omitempty"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression,omitempty"`
}

type Contains struct {
	// The Constant element identifies a constant value in a restriction.
	Constant *Constant `xml:"Constant"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
	// Determines whether the search ignores cases and spaces.
	ContainmentComparison *string `xml:"ContainmentComparison,attr"`
	// Identifies the boundaries of a search.
	ContainmentMode *string `xml:"ContainmentMode,attr"`
}

const (
	// The comparison must be exact.
	ContainsExact = `Exact`
	// The comparison ignores casing.
	ContainsIgnoreCase = `IgnoreCase`
	// The comparison ignores non-spacing characters.
	ContainsIgnoreNonSpacingCharacters = `IgnoreNonSpacingCharacters`
	// To be removed.
	ContainsLoose = `Loose`
	// The comparison ignores casing and non-spacing characters.
	ContainsIgnoreCaseAndNonSpacingCharacters = `IgnoreCaseAndNonSpacingCharacters`
	// To be removed.
	ContainsLooseAndIgnoreCase = `LooseAndIgnoreCase`
	// To be removed.
	ContainsLooseAndIgnoreNonSpace = `LooseAndIgnoreNonSpace`
	// To be removed.
	ContainsLooseAndIgnoreCaseAndIgnoreNonSpace = `LooseAndIgnoreCaseAndIgnoreNonSpace`
	// The comparison is between the full string and the constant. The property value and the supplied constant are precisely the same.
	ContainsFullString = `FullString`
	// The comparison is between the string prefix and the constant.
	ContainsPrefixed = `Prefixed`
	// The comparison is between a substring of the string and the constant.
	ContainsSubstring = `Substring`
	// The comparison is between a prefix on individual words in the string and the constant.
	ContainsPrefixOnWords = `PrefixOnWords`
	// The comparison is between an exact phrase in the string and the constant.
	ContainsExactPhrase = `ExactPhrase`
)

type Excludes struct {
	// The Bitmask element represents a hexadecimal or decimal mask to be used during an Excludes restriction operation.
	Bitmask Bitmask `xml:"t:Bitmask,omitempty"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes,omitempty"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type Bitmask string

type Exists struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type IsEqualTo struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type FieldURIOrConstant struct {
	// The Constant element identifies a constant value in a restriction.
	Constant Constant `xml:"t:Constant,omitempty"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type Constant string

type IsGreaterThan struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type IsGreaterThanOrEqualTo struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type IsLessThan struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type IsLessThanOrEqualTo struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type IsNotEqualTo struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type Not struct {
	// The And element represents a search expression that allows you to perform a Boolean AND operation between two or more search expressions. The result of the AND operation is true if all the search expressions contained within the And element are true.
	And *And `xml:"t:And,omitempty"`
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"t:Contains,omitempty"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes,omitempty"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"t:Exists,omitempty"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"t:IsEqualTo,omitempty"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"t:IsGreaterThan,omitempty"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"t:IsGreaterThanOrEqualTo,omitempty"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"t:IsLessThan,omitempty"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"t:IsLessThanOrEqualTo,omitempty"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"t:IsNotEqualTo,omitempty"`
	// The Or element represents a search expression that performs a logical OR on the search expression that it contains. Or will return true if any of its children return true. Or must have two or more children.
	Or *Or `xml:"t:Or,omitempty"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression,omitempty"`
}

type Or struct {
	// The And element represents a search expression that allows you to perform a Boolean AND operation between two or more search expressions. The result of the AND operation is true if all the search expressions contained within the And element are true.
	And *And `xml:"t:And,omitempty"`
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"t:Contains,omitempty"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes,omitempty"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"t:Exists,omitempty"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"t:IsEqualTo,omitempty"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"t:IsGreaterThan,omitempty"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"t:IsGreaterThanOrEqualTo,omitempty"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"t:IsLessThan,omitempty"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"t:IsLessThanOrEqualTo,omitempty"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"t:IsNotEqualTo,omitempty"`
	// The Not element represents a search expression that negates the Boolean value of the search expression that it contains.
	Not *Not `xml:"t:Not,omitempty"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression,omitempty"`
}

type SearchExpression struct {
	TEXT interface{} `xml:",chardata"`
}

type getFindItemResponseEnvelop struct {
	XMLName struct{}                `xml:"Envelope"`
	Body    getFindItemResponseBody `xml:"Body"`
}
type getFindItemResponseBody struct {
	FindItemResponse FindItemResponse `xml:"FindItemResponse"`
}

type FindItemResponse struct {
	ResponseMessages *GetFindItemResponseMessages `xml:"m:ResponseMessages"`
}

type GetFindItemResponseMessages struct {
	FindItemResponseMessage *FindItemResponseMessage `xml:"m:FindItemResponseMessage"`
}

type FindItemResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey,omitempty"`
	// The MessageText element provides a text description of the status of the response.
	MessageText MessageText `xml:"m:MessageText,omitempty"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml,omitempty"`
	// The ResponseCode element provides status information about the request.
	ResponseCode ResponseCode `xml:"m:ResponseCode"`
	// The RootFolder element contains the results of a search of a single root folder during a FindItem operation.
	RootFolder *RootFolderFindItemResponseMessage `xml:"m:RootFolder,omitempty"`
	// Describes the status of a FindItem operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	FindItemResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while processing an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store goes offline during the batch.  - Active Directory Domain Services (AD DS) goes offline.  - Mailboxes are moved.  - The message database (MDB) goes offline.  - A password is expired.  - A quota was exceeded.
	FindItemResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	FindItemResponseMessageError = `Error`
)

type RootFolderFindItemResponseMessage struct {
	// The Groups element contains a collection of groups that are found with the search and aggregation criteria that is identified in the FindItem operation request.
	Groups *Groups `xml:"t:Groups,omitempty"`
	// The Items element contains an array of items.
	Items *Items `xml:"m:Items,omitempty"`
	// Represents the next denominator to use for the next request when doing fractional paging.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// Indicates whether the current results contain the last item in the query, such that further paging is not needed.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when using an indexed paging view.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when using fraction page views.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Represents the total number of items that pass the restriction. In a grouped FindItem operation, the TotalItemsInView attribute returns the total number of items in the view plus the total number of groups.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}

type Groups struct {
	// The GroupedItems element represents a collection of items that are the result of a grouped FindItem operation call.
	GroupedItems *GroupedItems `xml:"t:GroupedItems,omitempty"`
}

type GroupedItems struct {
	// The GroupIndex element represents the property value that is used to group items for the current group of items in a FindItem operation call.
	GroupIndex GroupIndex `xml:"t:GroupIndex,omitempty"`
	// The Items element contains an array of items.
	Items *Items `xml:"t:Items,omitempty"`
}

type GroupIndex string
