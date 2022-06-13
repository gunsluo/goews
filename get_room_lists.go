package ews

type GetRoomListsRequest struct {
	XMLName struct{} `xml:"m:GetRoomLists"`
}

type GetRoomListsResponse struct {
	Response
	RoomLists RoomLists `xml:"RoomLists"`
}

type RoomLists struct {
	Address []EmailAddress `xml:"Address"`
}

type ItemId struct {
	Id        string `xml:"Id,attr"`
	ChangeKey string `xml:"ChangeKey,attr"`
}

type getRoomListsResponseEnvelop struct {
	XMLName struct{}                 `xml:"Envelope"`
	Body    getRoomListsResponseBody `xml:"Body"`
}
type getRoomListsResponseBody struct {
	GetRoomListsResponse GetRoomListsResponse `xml:"GetRoomListsResponse"`
}
