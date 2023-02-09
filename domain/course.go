package domain

import "time"

type Course struct {
	Name        string
	Description string
	Capacity    float64
	Price       float64
	Trainer     Trainer
}

type Trainer struct {
	FirstName string
	LastName  string
	Email     string
}

func (c Course) MakeSchedule(s Schedule) Class {
	return Class{
		Name:     c.Name,
		Schedule: s,
	}
}

type Class struct {
	Name      string
	Attendees []Attendee
	Schedule  Schedule
	SendEmail func([]Email)
}

type Schedule struct {
	Start time.Time
	End   time.Time
}

type Attendee struct {
	FirstName string
	LastName  string
	Email     string
}

func (c *Class) AddAtendee(attendee Attendee) {
	c.Attendees = append(c.Attendees, attendee)
}

func (c Class) AttendeeCount() int {
	return len(c.Attendees)
}

func (c Class) PrepareWelcomeEmail() []Email {
	emails := []Email{}
	for _, attendee := range c.Attendees {
		emails = append(emails, Email{
			To:   attendee.Email,
			From: "welcome@mail.com",
		})
	}
	return emails
}

func (c *Class) SetEmailSender(sendEmailFunc func([]Email)) {
	c.SendEmail = sendEmailFunc
}

type Email struct {
	To   string
	From string
}

func (c Class) SendWelcomeEmail() {
	c.SendEmail(c.makeEmailList())
}

func (c Class) makeEmailList() []Email {
	emails := []Email{}
	for _, attendee := range c.Attendees {
		emails = append(emails, Email{
			To:   attendee.Email,
			From: "welcome@mail.com",
		})
	}
	return emails
}
