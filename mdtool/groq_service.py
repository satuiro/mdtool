from groq import Groq
from typing import List, Dict, Optional
from .config import config
import json


class GroqService:
    def __init__(self):
        self.client = Groq(api_key=config.groq_api_key)

    async def generate_readme(self, files: Dict[str, Dict], metadata: Dict) -> str:
        """Generate a comprehensive README based on project structure and metadata."""

        # Create a structured project summary
        file_summaries = []
        for path, info in files.items():
            content_preview = (
                info["content"][:200] + "..."
                if len(info["content"]) > 200
                else info["content"]
            )
            file_summaries.append(f"### {path}\n```\n{content_preview}\n```")

        project_summary = "\n\n".join(file_summaries)

        # Format repository metadata
        metadata_summary = "\n".join(
            [
                f"- **{key}**: {value}"
                for key, value in metadata.items()
                if value is not None
            ]
        )

        prompt = f"""Generate a comprehensive README.md for the following project:

Repository Metadata:
{metadata_summary}

Project Files:
{project_summary}

Create a detailed README that includes:
1. Project title and description
2. Key features
3. Installation instructions
4. Usage examples
5. Project structure overview
6. Dependencies and requirements
7. Contributing guidelines (if applicable)
8. License information (if available)

Use clear markdown formatting with appropriate sections. Focus on creating a helpful
and informative README that would help users understand and use the project.
If the project has a specific focus or unique features, highlight those prominently.
"""

        try:
            response = self.client.chat.completions.create(
                messages=[{"role": "user", "content": prompt}],
                model=config.default_model,
                temperature=0.7,
                max_tokens=2000,
            )
            return response.choices[0].message.content.strip()
        except Exception as e:
            print(f"Error generating README: {str(e)}")
            return ""
