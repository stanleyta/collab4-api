import argparse
import sys
from dotenv import load_dotenv
from llm.provider import LLMProvider
from service.graph import LangGraphService

def main():
    load_dotenv()
    
    parser = argparse.ArgumentParser(description="LangGraph Python Example")
    parser.add_argument("--provider", type=str, default="ollama", choices=["gemini", "anthropic", "openai", "ollama"],
                        help="LLM provider to use")
    parser.add_argument("--model", type=str, help="Specific model name")
    parser.add_argument("--temperature", type=float, default=1.2, help="Temperature for the LLM")
    parser.add_argument("--prompt", type=str, default="Hello, tell me a short joke.", help="The prompt to send")
    
    args = parser.parse_args()

    try:
        provider = LLMProvider(args.provider, args.model, temperature=args.temperature)
        print(f"Service initialized. Provider: {args.provider}, Model: {provider.get_model_name()}, Temperature: {args.temperature}")
        service = LangGraphService(provider.get_llm())
        
        print(f"Processing prompt: {args.prompt}")
        response = service.run(args.prompt)
        
        print("\n--- Response ---")
        print(response)
        
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
