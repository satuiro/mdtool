import typer
import asyncio
from typing import Optional, Dict, List
import os
from rich.console import Console
from rich.markdown import Markdown
from rich.panel import Panel
from rich.progress import Progress
from github import Github, GithubException
import base64
from .groq_service import GroqService
from .config import config

app = typer.Typer()
console = Console()


def print_warning(message: str) -> None:
    """Print a warning message in yellow"""
    console.print(f"[yellow]WARNING:[/yellow] {message}")


def print_error(message: str) -> None:
    """Print an error message in red"""
    console.print(f"[red]ERROR:[/red] {message}")


def print_success(message: str) -> None:
    """Print a success message in green"""
    console.print(f"[green]SUCCESS:[/green] {message}")


def print_info(message: str) -> None:
    """Print an info message in blue"""
    console.print(f"[blue]INFO:[/blue] {message}")


class ReadmeGenerator:
    def __init__(
        self,
        github_token: Optional[str] = None,
        repo_name: Optional[str] = None,
        max_file_size: int = 100000,  # 100KB default max file size
        max_files_per_batch: int = 5,  # Process 5 files per batch
        exclude_patterns: List[str] = None,  # Patterns to exclude
    ):
        self.github_token = (
            github_token or os.getenv("GITHUB_TOKEN") or config.github_token
        )
        if not self.github_token:
            raise ValueError(
                "GitHub token not found. Please provide it via --github-token or set GITHUB_TOKEN environment variable"
            )

        self.github = Github(self.github_token)
        self.repo_name = repo_name
        self.max_file_size = max_file_size
        self.max_files_per_batch = max_files_per_batch
        self.exclude_patterns = exclude_patterns or [
            "node_modules/",
            "venv/",
            ".git/",
            "__pycache__/",
            "*.pyc",
            "*.pyo",
            "*.pyd",
            "*.so",
            "*.dylib",
            "*.dll",
        ]

    def should_include_file(self, file_path: str, file_size: int) -> bool:
        """Determine if file should be included in README analysis"""
        if file_size > self.max_file_size:
            return False

        for pattern in self.exclude_patterns:
            if pattern.startswith("*."):
                ext = pattern[1:]
                if file_path.endswith(ext):
                    return False
            elif pattern in file_path:
                return False

        return True

    async def get_repo_files(self) -> Dict:
        """Gather repository files and metadata"""
        try:
            repo = self.github.get_repo(self.repo_name)
            files_data = {}
            repo_metadata = {
                "name": repo.name,
                "description": repo.description,
                "language": repo.language,
                "license": repo.license.name if repo.license else None,
                "stars": repo.stargazers_count,
                "forks": repo.forks_count,
                "open_issues": repo.open_issues_count,
            }

            def process_contents(path=""):
                try:
                    contents = repo.get_contents(path)

                    # Handle single file case
                    if not isinstance(contents, list):
                        contents = [contents]

                    for content in contents:
                        if content.type == "dir":
                            # Recursively process directory
                            process_contents(content.path)
                        elif content.type == "file":
                            if not self.should_include_file(content.path, content.size):
                                continue

                            try:
                                file_data = base64.b64decode(content.content).decode(
                                    "utf-8"
                                )
                                files_data[content.path] = {
                                    "content": file_data,
                                    "size": content.size,
                                }
                                print_info(f"Added file: {content.path}")
                            except (GithubException, UnicodeDecodeError) as e:
                                print_warning(f"Skipping file {content.path}: {str(e)}")
                                continue
                except GithubException as e:
                    print_warning(f"Error accessing {path}: {str(e)}")

            with Progress() as progress:
                scan_task = progress.add_task(
                    "[cyan]Scanning repository...", total=None
                )
                process_contents()
                progress.update(scan_task, completed=100)

            print_success(f"Found {len(files_data)} relevant files")
            return {"files": files_data, "metadata": repo_metadata}

        except Exception as e:
            raise Exception(f"Failed to gather repository files: {str(e)}")

    def chunk_files(self, files: Dict) -> List[Dict]:
        """Split files into smaller batches"""
        file_items = list(files.items())
        return [
            {k: v for k, v in file_items[i : i + self.max_files_per_batch]}
            for i in range(0, len(file_items), self.max_files_per_batch)
        ]

    async def generate_readme_sections(self, repo_data: Dict) -> str:
        """Generate README content from repo data"""
        try:
            groq_service = GroqService()

            # Process files in batches
            file_batches = self.chunk_files(repo_data["files"])
            all_sections = []

            with Progress() as progress:
                analyze_task = progress.add_task(
                    "[green]Analyzing files...", total=len(file_batches)
                )

                for i, batch in enumerate(file_batches, 1):
                    print_info(f"Processing batch {i}/{len(file_batches)}")

                    # Create context for this batch
                    batch_context = {
                        "files": {
                            path: {"content": info["content"], "size": info["size"]}
                            for path, info in batch.items()
                        },
                        "metadata": repo_data["metadata"],
                    }

                    # Generate section content
                    section_content = await groq_service.generate_readme(
                        batch_context["files"], batch_context["metadata"]
                    )
                    all_sections.append(section_content)
                    progress.update(analyze_task, advance=1)

            # Combine all sections
            final_readme = "\n\n".join(filter(None, all_sections))
            return final_readme

        except Exception as e:
            print_error(f"README generation failed: {str(e)}")
            return ""

    def display_readme(self, content: str, format: str = "preview") -> None:
        """Display generated README"""
        if format == "raw":
            console.print(content)
        else:
            console.print("\n")
            console.print(
                Panel(Markdown(content), title="Generated README", border_style="green")
            )

    async def generate(self) -> str:
        """Main method to generate README"""
        try:
            print_info("Starting README generation...")
            repo_data = await self.get_repo_files()

            if not repo_data["files"]:
                print_warning("No suitable files found for analysis")
                return ""

            readme_content = await self.generate_readme_sections(repo_data)
            return readme_content

        except Exception as e:
            print_error(f"Generation failed: {str(e)}")
            return ""


@app.command()
def generate(
    github_token: Optional[str] = typer.Option(
        None,
        "--github-token",
        "-t",
        help="GitHub personal access token",
        envvar="GITHUB_TOKEN",
    ),
    repo_name: str = typer.Option(
        ..., "--repo", "-r", help="Repository name in format 'owner/repo'"
    ),
    output: str = typer.Option(
        "preview", "--output", "-o", help="Output format (preview/raw)"
    ),
    save: bool = typer.Option(False, "--save", "-s", help="Save README to file"),
) -> None:
    """Generate a comprehensive README for a GitHub repository"""
    try:
        generator = ReadmeGenerator(
            github_token=github_token,
            repo_name=repo_name,
        )

        readme_content = asyncio.run(generator.generate())

        if readme_content:
            # Display the generated README
            generator.display_readme(readme_content, output)

            # Save if requested
            if save:
                with open("README.md", "w") as f:
                    f.write(readme_content)
                print_success("\nREADME.md saved successfully!")
        else:
            print_error("Failed to generate README content")
            raise typer.Exit(1)

    except Exception as e:
        print_error(f"Generation failed: {str(e)}")
        raise typer.Exit(1)


if __name__ == "__main__":
    app()
