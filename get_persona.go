package ews

type GetPersonaRequest struct {
	XMLName   struct{}  `xml:"m:GetPersona"`
	PersonaId PersonaId `xml:"m:PersonaId"`
}

type getPersonaResponseEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    getPersonaResponseBody `xml:"Body"`
}
type getPersonaResponseBody struct {
	FindPeopleResponse GetPersonaResponse `xml:"GetPersonaResponseMessage"`
}

type GetPersonaResponse struct {
	Response
	Persona Persona `xml:"Persona"`
}
