import os
from typing import Optional
from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_anthropic import ChatAnthropic
from langchain_openai import ChatOpenAI
from langchain_ollama import ChatOllama
from langchain_core.language_models.chat_models import BaseChatModel

class LLMProvider:
    def __init__(self, provider_type: str, model_name: Optional[str] = None, temperature: float = 1.2):
        self.provider_type = provider_type.lower()
        self.temperature = temperature
        self.effective_model = model_name
        self.llm = self._init_llm()

    def _init_llm(self) -> BaseChatModel:
        if self.provider_type == "gemini":
            self.effective_model = self.effective_model or "gemini-2.5-flash"
            return ChatGoogleGenerativeAI(model=self.effective_model, temperature=self.temperature)
        elif self.provider_type == "anthropic":
            self.effective_model = self.effective_model or "claude-3-5-sonnet-20240620"
            return ChatAnthropic(model=self.effective_model, temperature=self.temperature)
        elif self.provider_type == "openai":
            self.effective_model = self.effective_model or "gpt-4o"
            return ChatOpenAI(model=self.effective_model, temperature=self.temperature)
        elif self.provider_type == "ollama":
            self.effective_model = self.effective_model or "llama3"
            return ChatOllama(model=self.effective_model, temperature=self.temperature)
        else:
            raise ValueError(f"Unsupported provider: {self.provider_type}")

    def get_llm(self) -> BaseChatModel:
        return self.llm

    def get_model_name(self) -> str:
        return self.effective_model
