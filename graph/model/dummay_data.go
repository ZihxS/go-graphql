package model

var DummyMeetups = []*Meetup{
	{
		ID:          "1",
		Name:        "A first meetup for Bob",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A first meetup for Jon",
		Description: "A description",
		UserID:      "2",
	},
	{
		ID:          "3",
		Name:        "A second meetup for Jon",
		Description: "A description",
		UserID:      "2",
	},
}

var DummyUsers = []*User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "jon@gmail.com",
	},
}
