package schema

// The ConversationAction element contains a single action to be applied to a single conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationaction
import "encoding/xml"

type ConversationAction struct {
	XMLName xml.Name

	// The Action element contains the action to perform on the conversation specified by the ConversationId element.
	Action *ActionConversationActionTypeType `xml:"Action"`
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"Categories"`
	// The ContextFolderId element indicates the folder that is targeted for actions that use folders. This element must be present when copying, deleting, moving, and setting read state on conversation items in a target folder.
	ContextFolderId *ContextFolderId `xml:"ContextFolderId"`
	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"ConversationId"`
	// The ConversationLastSyncTime element contains the date and time that a conversation was last synchronized. This element must be present when trying to delete all items in a conversation that were received up to the specified time.
	ConversationLastSyncTime *ConversationLastSyncTime `xml:"ConversationLastSyncTime"`
	// The DeleteType element indicates how items in a conversation are deleted.
	DeleteType *DeleteType `xml:"DeleteType"`
	// The DestinationFolderId element indicates the destination folder for copy and move actions.
	DestinationFolderId *DestinationFolderId `xml:"DestinationFolderId"`
	// The EnableAlwaysDelete element specifies a flag that enables deletion for all new items in a conversation.
	EnableAlwaysDelete *EnableAlwaysDelete `xml:"EnableAlwaysDelete"`
	// The IsRead element indicates whether a message has been read.
	IsRead *IsRead `xml:"IsRead"`
	// The ProcessRightAway element indicates whether the response is sent as soon as the action starts processing on the server or whether the response is sent after the action has completed. This element must be present for the response to be sent asynchronous to the requested action.
	ProcessRightAway *ProcessRightAway `xml:"ProcessRightAway"`
}

const (
	// Always categorize the conversation.
	ConversationActionAlwaysCategorize = `AlwaysCategorize`
	// Always delete the conversation.
	ConversationActionAlwaysDelete = `AlwaysDelete`
	// Always move the conversation.
	ConversationActionAlwaysMove = `AlwaysMove`
	// Copy the conversation.
	ConversationActionCopy = `Copy`
	// Delete the conversation.
	ConversationActionDelete = `Delete`
	// Move the conversation.
	ConversationActionMove = `Move`
	// Set the read state of the conversation.
	ConversationActionSetReadState = `SetReadState`
	// Set the retention policy for the conversation.
	ConversationActionSetRetentionPolicy = `SetRetentionPolicy`
)

func (C *ConversationAction) SetForMarshal() {
	C.XMLName.Local = "t:ConversationAction"
}

func (C *ConversationAction) GetSchema() *Schema {
	return &SchemaTypes
}
