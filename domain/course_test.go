package domain_test

import (
	"oddshub/domain"
	"testing"
	"time"
)

func Test_course_make_schedule_should_return_class_with_scheduled_date(t *testing.T) {
	start := time.Date(2023, 2, 13, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 2, 17, 17, 0, 0, 0, time.Local)
	course := domain.Course{}
	class := course.MakeSchedule(domain.Schedule{
		Start: start,
		End:   end,
	})

	if class.Name != course.Name {
		t.Errorf("want class name to be %s but got %s", course.Name, class.Name)
	}
	if class.Schedule.Start != start {
		t.Errorf("want class start to be %v but got %s", start, class.Schedule.Start)
	}
	if class.Schedule.End != end {
		t.Errorf("want class end to be %v but got %s", end, class.Schedule.End)
	}
}

func Test_class_add_attendee_should_increase_amount_of_attendee_in_class(t *testing.T) {
	start := time.Date(2023, 2, 13, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 2, 17, 17, 0, 0, 0, time.Local)
	course := domain.Course{}
	class := course.MakeSchedule(domain.Schedule{
		Start: start,
		End:   end,
	})

	if class.AttendeeCount() != 0 {
		t.Errorf("want attendee count to be %d but got %d", 0, class.AttendeeCount())
	}

	class.AddAtendee(domain.Attendee{
		FirstName: "Peerawat",
		LastName:  "Poombua",
		Email:     "peerawat@odds.team",
	})

	if class.AttendeeCount() != 1 {
		t.Errorf("want attendee cnount to be %d but got %d", 1, class.AttendeeCount())
	}
}

func Test_class_send_welcome_email_should_call_send_email(t *testing.T) {
	start := time.Date(2023, 2, 13, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 2, 17, 17, 0, 0, 0, time.Local)
	course := domain.Course{}
	class := course.MakeSchedule(domain.Schedule{
		Start: start,
		End:   end,
	})
	class.AddAtendee(domain.Attendee{
		FirstName: "Peerawat",
		LastName:  "Poombua",
		Email:     "peerawat@odds.team",
	})
	class.AddAtendee(domain.Attendee{
		FirstName: "Benya",
		LastName:  "Poombua",
		Email:     "gap@odds.team",
	})

	sendEmailFunc := func(emails []domain.Email) {
		if len(emails) != 2 {
			t.Errorf("want two email to be sent")
		}
		expectTo := "peerawat@odds.team"
		actualTo := emails[0].To
		if actualTo != expectTo {
			t.Errorf("want email to be sent to %s but got %s", expectTo, actualTo)
		}

		expectFrom := "welcome@mail.com"
		actualFrom := emails[0].From
		if actualFrom != expectFrom {
			t.Errorf("want email be sent from %s but got %s", expectFrom, actualFrom)
		}
	}
	class.SetEmailSender(sendEmailFunc)
	class.SendWelcomeEmail()
}
