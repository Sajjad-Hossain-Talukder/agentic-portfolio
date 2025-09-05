package main

import (
	"flag"
	"log"

	chatagent "github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/agent"
	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/api"
	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/config"
)

func main() {

	env := flag.String("env", "production", "Environment to run (staging, production)")
	flag.Parse()

	// // Load configuration
	cfg, err := config.LoadConfig(*env, false)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	chatAgent := chatagent.NewChatAgent(cfg)
	server := api.NewServer(cfg, chatAgent)
	err = server.RunWith(cfg.ServerAddress)
	if err != nil {
		log.Fatal("Failed to run server at %s", cfg.ServerAddress )
	}
	log.Printf("Running in %s environment on port %s", cfg.AppEnv, cfg.APIPort )
}