package chatagent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/config"
	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/utils"

	"github.com/sashabaranov/go-openai"
)

const (
	summaryFilePath 	= "src/resources/summary.txt"
	cvFilePath      	= "src/resources/cv.pdf"	
	linkedinFilePath 	= "src/resources/linkedin.pdf"
	modelName 			= "gemini-2.5-flash" 
	toolType 			= "function"
)

// Roles 
const (
	UserRole 		= "user"
	SystemRole 		= "system"
	ToolRole 		= "tool"
	AssistantRole 	= "assistant"
)



type Tool struct {
	Type     openai.ToolType           	`json:"type"`
	Function *openai.FunctionDefinition `json:"function,omitempty"`
	Func     ToolFunc          			 `json:"-"` 
}

type ToolFunc func(args map[string]string, pushUser, pushToken string) (map[string]string, error)

func NewGeminiClient(geminiApiKey string) *openai.Client {
	config := openai.DefaultConfig(geminiApiKey) 
	config.BaseURL = "https://generativelanguage.googleapis.com/v1beta/"
	client := openai.NewClientWithConfig(config)
	return client
}

type ChatMessage struct {
	Role    string `json:"role"` 
	Content string `json:"content"`
}

type ChatAgent struct {
    Name     		string
    Summary  		string
    LinkedInData 	string
	CVData 			string
    Client  	 	*openai.Client
	ModelName 		string
	Tools      		[]Tool
	PushUser		string 
	PushToken 		string 
}

func NewChatAgent(config *config.Config) *ChatAgent {
    client := NewGeminiClient(config.GeminiApiKey)

    summaryData, err := os.ReadFile(summaryFilePath)
    if err != nil {
        log.Println("Summary file error:", err)
    }

	linkedInText := utils.ReadPDFText(linkedinFilePath)
	cvText := utils.ReadPDFText(cvFilePath)

    return &ChatAgent{
        Name:    		"Sajjad Hossain Talukder",
        Summary: 		string(summaryData),
        LinkedInData: 	linkedInText,
		CVData: 		cvText,
        Client:  		client,
		ModelName:		modelName,
		PushUser: 		config.PushoverUser,
		PushToken:      config.PushoverToken,			
		Tools: []Tool{
			{
				Type: toolType,
				Function: &openai.FunctionDefinition{
					Name:        "record_user_details",
					Description: "Record that a user is interested in being in touch and provided an email",
					Strict:      true,
					Parameters: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"email": map[string]string{"type": "string", "description": "Email address"},
							"name":  map[string]string{"type": "string", "description": "User name"},
							"notes": map[string]string{"type": "string", "description": "Additional notes"},
						},
						"required": []string{"email"},
					},
				},
				Func: RecordUserDetails,
			},
			{
				Type: toolType,
				Function: &openai.FunctionDefinition{
					Name:        "record_unknown_question",
					Description: "Record any question that couldn't be answered",
					Strict:      true,
					Parameters: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"question": map[string]string{"type": "string", "description": "Question text"},
						},
						"required": []string{"question"},
					},
				},
				Func: RecordUnknownQuestion,
			}, 
		},
    }
}

func (ca *ChatAgent) SystemPrompt() string {
    prompt := fmt.Sprintf(`You are acting as %s. Answer questions professionally and engagingly. 
		Use the summary, CV data and LinkedIn profile data below. 
		Record any unknown questions using record_unknown_question, 
		record emails using record_user_details.  

		## Summary:
		%s

		## CV Data:
		%s

		## LinkedIn Profile:
		%s

		Always stay in character.
		If you find a conversation history, respond carefully without repeating greetings like “Hi” or “Hello” or any unnecessary introductions.
		Check the system prompt thoroughly.
		If there is a message from the system role, it means the user has already been greeted.
		Base your reply on the context of the conversation only. `, ca.Name, ca.Summary, ca.CVData, ca.LinkedInData)
	return prompt
}

func (ca *ChatAgent) Chat(message string, history []ChatMessage) (string, error) {
    messages := []openai.ChatCompletionMessage{
        {Role: UserRole, Content: ca.SystemPrompt()},
    }

    for _, h := range history {
        messages = append(messages, openai.ChatCompletionMessage{
            Role:    h.Role,
            Content: h.Content,
        })
    }

    messages = append(messages, openai.ChatCompletionMessage{
        Role:    UserRole,
        Content: message,
    })

	done := false
    var finalContent string

	toolsForRequest := make([]openai.Tool, len(ca.Tools))
	for i, t := range ca.Tools {
		toolsForRequest[i] = openai.Tool{
			Type:  openai.ToolType(t.Type),
			Function: t.Function,
		}
	}

	 for !done {
        resp, err := ca.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
            Model:    ca.ModelName,
            Messages: messages,
			Tools: toolsForRequest,
        })
        if err != nil {
            return "", err
        }

        choice := resp.Choices[0]

		if choice.FinishReason == openai.FinishReasonToolCalls {
			for _, tool := range choice.Message.ToolCalls {
				argsJSON := tool.Function.Arguments
				toolName := tool.Function.Name
			
				args := make(map[string]string)
				if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
					log.Printf("Failed to parse tool arguments for %s: %v", toolName, err)
					continue
				}

				result, err := ca.HandleToolCall(toolName, args)
				if err != nil {
					log.Printf("Error executing tool %s: %v", toolName, err)
					continue 
				}
				content := fmt.Sprintf("%v tool call %v", toolName, result[Recorder])
		
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    AssistantRole,
					Content: content,
				})
			}
		} else {
			finalContent = choice.Message.Content
			done = true
		}
    }

    return finalContent, nil
}

func (ca *ChatAgent) HandleToolCall(toolName string, args map[string]string) (map[string]string, error) {
    for _, t := range ca.Tools {
        if t.Function.Name == toolName {
            return t.Func(args, ca.PushUser, ca.PushToken)
        }
    }
    return nil, fmt.Errorf("tool %s not found", toolName)
}

