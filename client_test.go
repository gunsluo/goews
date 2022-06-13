package goews

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	c, err := NewClient(
		Config{
			Address:  "https://outlook.office365.com/EWS/Exchange.asmx",
			Username: "daziplqa@daziplqa.onmicrosoft.com",
			Password: "systemsystem@123",
			Dump:     true,
			NTLM:     true,
			SkipTLS:  true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	/*
		err = testSendEmail(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testListUsersEvents(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testCreateEvent(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testGetRoomLists(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testFindPeople(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testGetUserPhoto(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testEWSUtilFindPeople(c)
		if err != nil {
			t.Fatal(err)
		}

		err = testGetPersona(c)
		if err != nil {
			t.Fatal(err)
		}
	*/

	_ = c

	fmt.Println("--- success ---")
}

func testSendEmail(c Client) error {
	return c.SendEmail(
		"daziplqa@daziplqa.onmicrosoft.com",
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)
}

func testListUsersEvents(c Client) error {
	eventUsers := []EventUser{
		{
			Email:        "mhewedy@mhewedy.onmicrosoft.com",
			AttendeeType: AttendeeTypeRequired,
		},
		{
			Email:        "example@mhewedy.onmicrosoft.com",
			AttendeeType: AttendeeTypeRequired,
		},
		{
			Email:        "room001@mhewedy.onmicrosoft.com",
			AttendeeType: AttendeeTypeResource,
		},
	}
	start, _ := time.Parse(time.RFC3339, "2019-12-10T11:00:00+03:00")

	events, err := c.ListUsersEvents(eventUsers, start, 48*time.Hour)
	if err != nil {
		return err
	}

	fmt.Println(events)

	return nil
}

func testCreateEvent(c Client) error {
	return c.CreateEvent(
		[]string{"mhewedy@mhewedy.onmicrosoft.com", "example2@mhewedy.onmicrosoft.com"},
		[]string{},
		"An Event subject",
		"An Event body, as plain text",
		"Room 55",
		time.Now().Add(24*time.Hour),
		30*time.Minute,
	)
}

func testGetRoomLists(c Client) error {
	response, err := c.GetRoomLists()
	if err != nil {
		return err
	}
	fmt.Println(response)

	return nil
}

func testFindPeople(c Client) error {
	req := &FindPeopleRequest{IndexedPageItemView: IndexedPageItemView{
		MaxEntriesReturned: math.MaxInt32,
		Offset:             0,
		BasePoint:          BasePointBeginning,
	}, ParentFolderId: ParentFolderId{
		DistinguishedFolderId: DistinguishedFolderId{Id: "directory"}},
		PersonaShape: &PersonaShape{BaseShape: BaseShapeIdOnly,
			AdditionalProperties: AdditionalProperties{
				FieldURI: []FieldURI{
					{FieldURI: "persona:DisplayName"},
					{FieldURI: "persona:Title"},
					{FieldURI: "persona:EmailAddress"},
				},
			}},
		QueryString: "ex",
	}

	resp, err := c.FindPeople(req)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func testGetUserPhoto(c Client) error {
	bytes, err := c.GetDecodingUserPhoto("mhewedy@mhewedy.onmicrosoft.com")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("/tmp/file.png", bytes, os.ModePerm)
	fmt.Println("written to: /tmp/file.png")

	return err
}

func testEWSUtilFindPeople(c Client) error {
	resp, err := c.FindPeopleByCondition("test")
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func testGetPersona(c Client) error {
	personas, _ := c.FindPeopleByCondition("hewedy")

	resp, err := c.GetPersona(&GetPersonaRequest{
		PersonaId: personas[0].PersonaId,
	})

	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
