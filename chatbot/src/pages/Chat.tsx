import React, { useState, useRef, useEffect } from "react";
import { FiSend } from "react-icons/fi"; 
import botImage from "../assets/bot.jpg";
import { Helmet } from "react-helmet";

type Message = {
  id: number;
  sender: "user" | "bot";
  text: string;
};

const Chat: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState("");
  const [loading, setLoading] = useState(false);
  const messagesContainerRef = useRef<HTMLDivElement>(null);
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  // Auto-scroll to bottom whenever messages or loading changes
  const scrollToBottom = () => {
    if (messagesContainerRef.current) {
      messagesContainerRef.current.scrollTop =
        messagesContainerRef.current.scrollHeight;
    }
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages, loading]);

  // Auto-expand textarea up to max height
  useEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.style.height = "auto";
      const maxHeight = 144; // px, ~8 lines
      textareaRef.current.style.height = Math.min(
        textareaRef.current.scrollHeight,
        maxHeight
      ) + "px";
    }
  }, [input]);

  const handleSend = () => {
    if (!input.trim()) return;

    const userMessage: Message = {
      id: Date.now(),
      sender: "user",
      text: input,
    };
    setMessages((prev) => [...prev, userMessage]);
    setInput("");
    setLoading(true);

    // Simulate bot response
    setTimeout(() => {
      const botMessage: Message = {
        id: Date.now() + 1,
        sender: "bot",
        text: `You said: "${userMessage.text}"`,
      };
      setMessages((prev) => [...prev, botMessage]);
      setLoading(false);
    }, 1500);
  };

  return (
    <> 
         <Helmet>
        <title>Let's Chat</title>
        <link rel="icon" type="image/svg+xml" href="/favicon.svg" />

      </Helmet>
   
    <div className="min-h-screen w-full bg-gray-900 flex items-end justify-center p-4">
      <div className="flex flex-col w-full max-w-3xl rounded-2xl shadow-lg overflow-hidden h-[90vh]">
        {/* Messages */}
        <div
          ref={messagesContainerRef}
          className="flex-1 p-6 flex flex-col space-y-4 overflow-y-auto scrollbar-hide"
        >
        
        {messages.length === 0 && !loading && (
        <div className="flex flex-col items-center justify-center mt-16 space-y-4 text-center">
            {/* Bot Image */}
            <img
            src={botImage} // Make sure the path is correct
            alt="Bot"
            className="w-24 h-24 rounded-full object-cover shadow-xl animate-bounce-slow"
            />

            {/* Title */}
            <p className="text-2xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-500 animate-pulse">
            Hey there! I'm Sajjad's AI assistant 🤖
            </p>

            {/* Subtitle */}
            <p className="text-gray-400 italic">
            Ask me anything, and let's have a chat!
            </p>
        </div>
        )}



          {messages.map((msg) => (
            <div
              key={msg.id}
              className={`flex ${
                msg.sender === "user" ? "justify-start" : "justify-end"
              }`}
            >
              <div
                className={`px-4 py-2 max-w-[80%] break-words rounded-xl shadow ${
                  msg.sender === "user"
                    ? "bg-blue-500 text-white rounded-tl-none"
                    : "bg-gray-700 text-white rounded-tr-none"
                }`}
              >
                {msg.text}
              </div>
            </div>
          ))}

          {loading && (
            <div className="flex justify-end">
              <div className="bg-gray-700 text-white px-4 py-2 rounded-xl animate-pulse rounded-tr-none">
                Typing...
              </div>
            </div>
          )}
        </div>

        <div className="border-t border-gray-700 p-4 flex gap-2 bg-gray-800 items-end">
        <textarea
            ref={textareaRef}
            rows={1}
            className="flex-1 rounded-md px-4 py-3 bg-gray-900 text-white border border-gray-700 focus:outline-none focus:border-blue-400 resize-none overflow-y-auto max-h-24 scrollbar-hide transition-all duration-150"
            placeholder="Type your message..."
            value={input}
            onChange={(e) => {
            setInput(e.target.value);

            // Auto-expand
            if (textareaRef.current) {
                textareaRef.current.style.height = "auto";
                const maxHeight = 48; // 24 * 2px lines
                textareaRef.current.style.height = Math.min(
                textareaRef.current.scrollHeight,
                maxHeight
                ) + "px";

                // Scroll last line into view
                textareaRef.current.scrollTop = textareaRef.current.scrollHeight;
            }
            }}
            onKeyDown={(e) => {
            if (e.key === "Enter" && !e.shiftKey) {
                e.preventDefault();
                handleSend();
            }
            }}
        />
        <button
        className="bg-blue-600 w-12 h-12 flex items-center justify-center rounded-full text-white hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed flex-shrink-0"
        onClick={handleSend}
        disabled={!input.trim() || loading}
        >
        <FiSend size={20} />
        </button>

        </div>

      </div>
    </div>
     </>
  );
};

export default Chat;
