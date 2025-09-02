package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

  type CV struct {
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	Contact     ContactInfo       `json:"contact"`
	Education   []EducationEntry  `json:"education"`
	Experience  []ExperienceEntry `json:"experience"`
	Projects    []ProjectEntry    `json:"projects"`
	Skills      Skills            `json:"skills"`
	Achievements Achievements     `json:"achievements"`
	Publication Publication       `json:"publication"`
	Organizations []Organization  `json:"organizations"`
	OnlineJudges OnlineJudges     `json:"online_judges"`
	PersonalInfo PersonalInfo     `json:"personal_info"`
}

type ContactInfo struct {
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	LinkedIn string `json:"linkedin"`
	GitHub   string `json:"github"`
	Website  string `json:"website,omitempty"`
}

type EducationEntry struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	Location    string `json:"location"`
	Period      string `json:"period"`
}

type ExperienceEntry struct {
	Role          string   `json:"role"`
	Company       string   `json:"company"`
	Location      string   `json:"location"`
	Period        string   `json:"period"`
	Responsibilities []string `json:"responsibilities"`
}

type ProjectEntry struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Tech        []string `json:"tech"`
	Image       string   `json:"image,omitempty"`
}

type Skills struct {
	LanguagesFrameworks []string `json:"languages_frameworks"`
	Tools               []string `json:"tools"`
	Interests           []string `json:"interests"`
}

type Achievements struct {
	ICPC_NCPC []string `json:"icpc_ncpc"`
	IUPC      []string `json:"iupc"`
	Scholarship []string `json:"scholarship"`
}

type Publication struct {
	Title      string `json:"title"`
	Conference string `json:"conference"`
}

type Organization struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	Period string `json:"period"`
}

type OnlineJudges struct {
	Codeforces JudgeProfile `json:"codeforces"`
	Codechef   JudgeProfile `json:"codechef"`
	Leetcode   JudgeProfile `json:"leetcode"`
	LightOJ    JudgeProfile `json:"lightoj"`
	TOPH      JudgeProfile `json:"toph"`
}

type JudgeProfile struct {
	Handle         string `json:"handle"`
	Profile        string `json:"profile"`
	MaxRating      int    `json:"max_rating,omitempty"`
	Status         string `json:"status,omitempty"`
	ProblemsSolved int    `json:"problems_solved,omitempty"`
}

type PersonalInfo struct {
	DOB          string   `json:"dob"`
	Gender       string   `json:"gender"`
	MaritalStatus string   `json:"marital_status"`
	Address      string   `json:"address"`
	Nationality  string   `json:"nationality"`
	Religion     string   `json:"religion"`
	Languages    []string `json:"languages"`
	Hobbies      []string `json:"hobbies"`
}


func NewGeminiClient() *openai.Client {
	config := openai.DefaultConfig(os.Getenv("GEMINI_API_KEY")) 
	config.BaseURL = "https://generativelanguage.googleapis.com/v1beta/"
	client := openai.NewClientWithConfig(config)
	return client
}

func GeneratePortfolioHTML(cvPath, images string) (string, error) {
	data, err := os.ReadFile(cvPath)
	if err != nil {
		return "", fmt.Errorf("error reading CV file: %w", err)
	}

	var cv CV
	if err := json.Unmarshal(data, &cv); err != nil {
		return "", fmt.Errorf("error parsing CV JSON: %w", err)
	}

	fmt.Println("Data >> ", string (data) )

	client := NewGeminiClient()
	ctx := context.Background()

	prompt := fmt.Sprintf(`
		You are an expert AI web designer and developer. 
		Using the following JSON data, generate a **single-page HTML portfolio** for a software engineer. 
		Use TailwindCSS for styling. The HTML should be modern, fully responsive, and mobile-friendly. 

		**Data (JSON)**: %s

		**Requirements:**
		1. **Single-page layout** with smooth scroll for each section.
		2. Sections to include, in order: 
		- Header: Name, title, profile picture, contact links.
		- About: Short description or summary.
		- Education: Degree, institution, location, period.
		- Experience: Role, company, location, period, responsibilities.
		- Projects: Name, description, tech stack, link, optional image.
		- Skills: Grouped by categories (languages/frameworks, tools, interests).
		- Achievements: Include ICPC/NCPC, IUPC, scholarships.
		- Online Judges: Codeforces, Codechef, Leetcode, LightOJ, TOPH.
		- Publications: Title, conference.
		- Organizations: Role, period.
		- Personal Info: DOB, gender, address, nationality, languages, hobbies.
		3. Use "/portfolio/images/profile.png" as profile picture .
		4. Keep **clean, minimal styling** using TailwindCSS classes.
		5. Use **semantic HTML tags** (<header>, <section>, <article>, <footer>).
		6. Include a **navigation bar** at the top linking to each section for smooth scrolling.
		7. Output must be **complete HTML** with "<html>", "<head>", "<body>" tags.
		8. Make sure **all fields from JSON are included**. Include the name prominently in the header.

		**Constraints:** 
		- No external JS frameworks; only TailwindCSS is allowed.
		- The page must fit in a single scrollable HTML page.
		- Ensure text content from JSON is fully visible; do not truncate.
		`, string(data))


	fmt.Println("Prompt >> ", string (prompt) )


	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gemini-2.5-flash",
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		MaxTokens: 3000,
	})
	if err != nil {
		return "", fmt.Errorf("openai request failed: %w", err)
	}

	// 5. Extract HTML
	html := resp.Choices[0].Message.Content

	fmt.Println("Resp >> ", html)
	return html, nil
}
