import os
from typing import Optional
from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_anthropic import ChatAnthropic
from langchain_openai import ChatOpenAI
from langchain_core.language_models.chat_models import BaseChatModel

class LLMProvider:
    def __init__(self, provider_type: str, model_name: Optional[str] = None):
        self.provider_type = provider_type.lower()
        self.model_name = model_name
        self.llm = self._init_llm()

    def _init_llm(self) -> BaseChatModel:
        if self.provider_type == "gemini":
            model = self.model_name or "gemini-2.5-flash"
            return ChatGoogleGenerativeAI(model=model)
        elif self.provider_type == "anthropic":
            model = self.model_name or "claude-3-5-sonnet-20240620"
            return ChatAnthropic(model=model)
        elif self.provider_type == "openai":
            model = self.model_name or "gpt-4o"
            return ChatOpenAI(model=model)
        else:
            raise ValueError(f"Unsupported provider: {self.provider_type}")

    def get_llm(self) -> BaseChatModel:
        return self.llm
