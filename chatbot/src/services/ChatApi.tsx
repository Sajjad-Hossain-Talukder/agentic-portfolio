// src/services/chatApi.ts
import { post } from "./ApiService";


export interface ChatMessage {
  role: "user" | "system";
  content: string;
}

interface ChatRequest {
  message: string;
  history: ChatMessage[];
}

interface ChatResponse {
  reply: string;
}

export async function ChatWithAgent(message: string, history: ChatMessage[] = []): Promise<ChatResponse> {
  const payload: ChatRequest = { message, history };
  return post<ChatResponse>("chat", payload);
}
