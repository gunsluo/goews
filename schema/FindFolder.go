package schema

// The FindFolder element defines a request to find folders in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolder
import "encoding/xml"

type FindFolder struct {
	XMLName xml.Name

	// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
	FolderShape *FolderShape `xml:"FolderShape"`
	// The FractionalPageFolderView element describes where the paged view starts and the maximum number of folders returned in a FindFolder request.
	FractionalPageFolderView *FractionalPageFolderView `xml:"FractionalPageFolderView"`
	// The IndexedPageFolderView element describes how paged item information is returned in a FindFolder response.
	IndexedPageFolderView *IndexedPageFolderView `xml:"IndexedPageFolderView"`
	// The ParentFolderIds element identifies folders for the FindItem and FindFolder operations to search.
	ParentFolderIds *ParentFolderIds `xml:"ParentFolderIds"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"Restriction"`
	// Defines how a search is performed. This attribute is required.
	Traversal *string `xml:"Traversal,attr"`
}

const (
	// Instructs the FindFolder operation to search only the identified folder and to return only the folder IDs for items that have not been deleted. This is called a shallow traversal.
	FindFolderShallow = `Shallow`
	// Instructs the FindFolder operation to search in all child folders of the identified parent folder and to return only the folder IDs for items that have not been deleted. This is called a deep traversal.
	FindFolderDeep = `Deep`
	// Instructs the FindFolder operation to perform a shallow traversal search for deleted items.
	FindFolderSoftDeleted = `SoftDeleted`
)

func (F *FindFolder) SetForMarshal() {
	F.XMLName.Local = "m:FindFolder"
}

func (F *FindFolder) GetSchema() *Schema {
	return &SchemaMessages
}
