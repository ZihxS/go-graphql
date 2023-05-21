package model

var DummyMeetups = []*Meetup{
	{
		ID:          1,
		Name:        "A first meetup for Bob",
		Description: "A description first meetup for Bob",
		UserID:      1,
	},
	{
		ID:          2,
		Name:        "A first meetup for Jon",
		Description: "A description first meetup for Jon",
		UserID:      2,
	},
	{
		ID:          3,
		Name:        "A second meetup for Jon",
		Description: "A description second meetup for Jon",
		UserID:      2,
	},
}

var DummyUsers = []*User{
	{
		ID:       1,
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       2,
		Username: "Jon",
		Email:    "jon@gmail.com",
	},
	{
		ID:       3,
		Username: "Jane",
		Email:    "jane@gmail.com",
	},
	{
		ID:       4,
		Username: "John",
		Email:    "john@gmail.com",
	},
}
