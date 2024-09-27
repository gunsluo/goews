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
	BodyType string `xml:"t:BodyType,omitempty"`
	// The ConvertHtmlCodePageToUTF8 element indicates whether the item HTML body is converted to UTF8.
	ConvertHtmlCodePageToUTF8 ConvertHtmlCodePageToUTF8 `xml:"t:ConvertHtmlCodePageToUTF8,omitempty"`
	// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
	FilterHtmlContent FilterHtmlContent `xml:"t:FilterHtmlContent,omitempty"`
	// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
	IncludeMimeContent IncludeMimeContent `xml:"t:IncludeMimeContent,omitempty"`
}

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
	Constant *Constant `xml:"t:Constant,omitempty"`
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
	Constant *Constant `xml:"t:Constant,omitempty"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI,omitempty"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI,omitempty"`
}

type Constant struct {
	// Specifies the value to compare in the restriction.
	Value string `xml:"Value,attr"`
}

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
	ResponseMessages *GetFindItemResponseMessages `xml:"ResponseMessages"`
}

type GetFindItemResponseMessages struct {
	FindItemResponseMessage *FindItemResponseMessage `xml:"FindItemResponseMessage"`
}

type FindItemResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey,omitempty"`
	// The MessageText element provides a text description of the status of the response.
	MessageText MessageText `xml:"MessageText,omitempty"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml,omitempty"`
	// The ResponseCode element provides status information about the request.
	ResponseCode ResponseCode `xml:"ResponseCode"`
	// The RootFolder element contains the results of a search of a single root folder during a FindItem operation.
	RootFolder *RootFolderFindItemResponseMessage `xml:"RootFolder,omitempty"`
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
	Groups *Groups `xml:"Groups,omitempty"`
	// The Items element contains an array of items.
	Items *Items `xml:"Items,omitempty"`
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

type Items struct {
	XMLName xml.Name

	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem []*CalendarItem `xml:"CalendarItem"`
	// // The Contact element represents a contact item in the Exchange store.
	// Contact []*Contact `xml:"Contact"`
	// // The DistributionList element represents a distribution list.
	// DistributionList []*DistributionList `xml:"DistributionList"`
	// // The Item element represents a generic item in the Exchange store.
	// Item []*Item `xml:"Item"`
	// // The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	// MeetingCancellation []*MeetingCancellation `xml:"MeetingCancellation"`
	// // The MeetingMessage element represents a meeting in the Exchange store.
	// MeetingMessage []*MeetingMessage `xml:"MeetingMessage"`
	// // The MeetingRequest element represents a meeting request in the Exchange store.
	// MeetingRequest []*MeetingRequest `xml:"MeetingRequest"`
	// // The MeetingResponse element represents a meeting response in the Exchange store.
	// MeetingResponse []*MeetingResponse `xml:"MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message []*Message `xml:"Message"`
	// The PostItem element represents a post item in the Exchange store.
	// PostItem []*PostItem `xml:"PostItem"`
	// // The Task element represents a task in the Exchange store.
	// Task []*Task `xml:"Task"`
}

// type TMessage struct {
// 	ItemClass     string        `xml:"t:ItemClass,omitempty"`
// 	ItemId        *ItemId       `xml:"t:ItemId,omitempty"`
// 	Subject       string        `xml:"t:Subject"`
// 	Body          TBody         `xml:"t:Body"`
// 	Sender        OneMailbox    `xml:"t:Sender"`
// 	ToRecipients  XMailbox      `xml:"t:ToRecipients"`
// 	CcRecipients  *XMailbox     `xml:"t:CcRecipients,omitempty"`
// 	BccRecipients *XMailbox     `xml:"t:BccRecipients,omitempty"`
// 	Attachments   *TAttachments `xml:"t:Attachments,omitempty"`
// }

type Message struct {
	// // The Attachments element contains the items or files that are attached to an item in the Exchange store.
	// Attachments *Attachments `xml:"Attachments"`
	// // The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	// BccRecipients *BccRecipients `xml:"BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"Body"`
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	// Categories *Categories `xml:"Categories"`
	// // The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	// CcRecipients *CcRecipients `xml:"CcRecipients"`
	// // The ConversationId element contains the identifier of an item or conversation.
	// ConversationId *ConversationId `xml:"ConversationId"`
	// // The ConversationIndex element contains a binary ID that represents the thread to which this message belongs.
	// ConversationIndex *ConversationIndex `xml:"ConversationIndex"`
	// // The ConversationTopic element represents the conversation topic.
	// ConversationTopic *ConversationTopic `xml:"ConversationTopic"`
	// // The Culture element represents the culture for a given item in a mailbox.
	// Culture *Culture `xml:"Culture"`
	// // The DateTimeCreated element represents the date and time that an item in the mailbox was created.
	// DateTimeCreated *DateTimeCreated `xml:"DateTimeCreated"`
	// // The DateTimeReceived element represents the date and time that an item in a mailbox was received.
	// DateTimeReceived *DateTimeReceived `xml:"DateTimeReceived"`
	// // The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
	// DateTimeSent *DateTimeSent `xml:"DateTimeSent"`
	// // The DisplayCc element represents the display string that is used for the contents of the Cc box. This is the concatenated string of all Cc recipient display names.
	// DisplayCc *DisplayCc `xml:"DisplayCc"`
	// // The DisplayTo element represents the display string that is used for the contents of the To box. This is the concatenated string of all To recipient display names.
	// DisplayTo *DisplayTo `xml:"DisplayTo"`
	// // The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	// EffectiveRights *EffectiveRights `xml:"EffectiveRights"`
	// // The ExtendedProperty element identifies extended MAPI properties on folders and items.
	// ExtendedProperty *ExtendedProperty `xml:"ExtendedProperty"`
	// // The From element represents the address from which the message was sent.
	// From *From `xml:"From"`
	// // The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	// HasAttachments *HasAttachments `xml:"HasAttachments"`
	// // The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	// Importance *Importance `xml:"Importance"`
	// // The InReplyTo element represents the identifier of the item to which this item is a reply.
	// InReplyTo *InReplyTo `xml:"InReplyTo"`
	// // The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	// InternetMessageHeaders *InternetMessageHeaders `xml:"InternetMessageHeaders"`
	// // The InternetMessageId element represents the Internet message identifier of an item.
	// InternetMessageId *InternetMessageId `xml:"InternetMessageId"`
	// // The IsAssociated element indicates whether the item is associated with a folder.
	// IsAssociated *IsAssociated `xml:"IsAssociated"`
	// // The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	// IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"IsDeliveryReceiptRequested"`
	// // The IsDraft element indicates whether an item has not yet been sent.
	// IsDraft *IsDraft `xml:"IsDraft"`
	// // The IsFromMe element indicates whether a user sent an item to him or herself.
	// IsFromMe *IsFromMe `xml:"IsFromMe"`
	// // The IsRead element indicates whether a message has been read.
	// IsRead *IsRead `xml:"IsRead"`
	// // The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	// IsReadReceiptRequested *IsReadReceiptRequested `xml:"IsReadReceiptRequested"`
	// // The IsResend element indicates whether the item had previously been sent.
	// IsResend *IsResend `xml:"IsResend"`
	// // The IsResponseRequested element indicates whether a response to an item is requested.
	// IsResponseRequested *IsResponseRequested `xml:"IsResponseRequested"`
	// // The IsSubmitted element indicates whether an item has been submitted to the Outbox default folder.
	// IsSubmitted *IsSubmitted `xml:"IsSubmitted"`
	// // The IsUnmodified element indicates whether the item has been modified.
	// IsUnmodified *IsUnmodified `xml:"IsUnmodified"`
	// // The ItemClass element represents the message class of an item.
	// ItemClass *ItemClass `xml:"ItemClass"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The LastModifiedName element contains the display name of the last user to modify an item. This element is read-only. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// LastModifiedName *LastModifiedName `xml:"LastModifiedName"`
	// // The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	// LastModifiedTime *LastModifiedTime `xml:"LastModifiedTime"`
	// // The MimeContent element contains the ASCII MIME stream of an object that is represented in base64Binary format and supports [RFC2045].
	// MimeContent *MimeContent `xml:"MimeContent"`
	// // The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	// ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// // The ReceivedBy element identifies the delegate in a delegate access scenario.
	// ReceivedBy *ReceivedBy `xml:"ReceivedBy"`
	// // The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	// ReceivedRepresenting *ReceivedRepresenting `xml:"ReceivedRepresenting"`
	// // The References element represents the Usenet header that is used to associate replies with the original messages.
	// References *References `xml:"References"`
	// // The ReminderDueBy element represents the date and time when the event occurs. This is used by the ReminderMinutesBeforeStart element to determine when the reminder is displayed.
	// ReminderDueBy *ReminderDueBy `xml:"ReminderDueBy"`
	// // The ReminderIsSet element indicates whether a reminder has been set for an item in the Exchange store.
	// ReminderIsSet *ReminderIsSet `xml:"ReminderIsSet"`
	// // The ReminderMessageData element specifies the data in a reminder message.
	// ReminderMessageData *ReminderMessageData `xml:"ReminderMessageData"`
	// // The ReminderMinutesBeforeStart element represents the number of minutes before an event occurs when a reminder is displayed.
	// ReminderMinutesBeforeStart *ReminderMinutesBeforeStart `xml:"ReminderMinutesBeforeStart"`
	// // The ReplyTo element identifies an array of addresses to which replies should be sent.
	// ReplyTo *ReplyTo `xml:"ReplyTo"`
	// // The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
	// ResponseObjects *ResponseObjects `xml:"ResponseObjects"`
	// The Sender element identifies the sender of an item.
	Sender *Sender `xml:"Sender"`
	// // The Sensitivity element indicates the sensitivity level of an item.
	// Sensitivity *Sensitivity `xml:"Sensitivity"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"Size"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// // The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	// ToRecipients []*ToRecipients `xml:"ToRecipients"`
	// // The UniqueBody element represents an HTML fragment or plain text which represents the unique body of this conversation.
	// UniqueBody *UniqueBody `xml:"UniqueBody"`
	// // The WebClientEditFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to edit an item in Outlook Web App.
	// WebClientEditFormQueryString *WebClientEditFormQueryString `xml:"WebClientEditFormQueryString"`
	// // The WebClientReadFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to read an item in Outlook Web App.
	// WebClientReadFormQueryString *WebClientReadFormQueryString `xml:"WebClientReadFormQueryString"`
}

type Body struct {
	// Specifies the type of the body.
	BodyType *string `xml:"BodyType,attr"`
	// Boolean value that indicates whether the body is truncated.
	IsTruncated *string `xml:"IsTruncated,attr"`
	TEXT        string  `xml:",chardata"`
}

type Sender struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

type Size struct {
	TEXT int64 `xml:",chardata"`
}

type Subject struct {
	TEXT string `xml:",chardata"`
}
