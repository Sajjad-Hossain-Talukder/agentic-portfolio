package chatagent

import (
	"fmt"

	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/services"
)

const (
	Recorder		= "recorder"
	RecorderOk 		= "successful"
	RecorderFail 	= "failed"
	EmailKey		= "email"
	NameKey			= "name"
	NotesKey 		= "notes"
	QuestionKey		= "question"
)



func RecordUserDetails(args map[string]string, pushUser, pushToken string) (map[string]string, error) {
    email := args[EmailKey]
    name := args[NameKey]
    notes := args[NotesKey]
    message := fmt.Sprintf("Recording user: %s (%s), notes: %s\n", name, email, notes)
	err := services.PushService(message, pushUser, pushToken)
	if err != nil {
		 return map[string]string{Recorder: RecorderFail}, nil
	}
    return map[string]string{Recorder: RecorderOk}, nil
}

func RecordUnknownQuestion(args map[string]string, pushUser, pushToken string) (map[string]string, error) {
    question := args[QuestionKey]
	message := fmt.Sprintf("Recording unknown question: %s\n", question)
	err := services.PushService(message, pushUser, pushToken)
	if err != nil {
		 return map[string]string{Recorder: RecorderFail}, nil
	}
    return map[string]string{Recorder: RecorderOk}, nil
}