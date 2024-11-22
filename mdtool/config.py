import os
from dataclasses import dataclass


@dataclass
class Config:
    groq_api_key: str = os.getenv("GROQ_API_KEY", "")
    github_token: str = os.getenv("GITHUB_TOKEN", "")
    default_model: str = "llama-3.2-90b-vision-preview"


config = Config()
