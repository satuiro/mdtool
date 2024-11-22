import typer
from rich.console import Console
from . import __version__
from .readme import app as readme_app

app = typer.Typer(
    help="mdtool - A CLI tool to generate and render README for a github project",
    no_args_is_help=True,
)

console = Console()


@app.callback(invoke_without_command=True)
def main(
    version: bool = typer.Option(
        False, "--version", "-V", help="Print version information"
    ),
) -> None:
    """
    mdtool - A CLI tool to generate and render README for a github project
    """
    if version:
        console.print(f"mdtool version: {__version__}")
        raise typer.Exit()


# Add commands directly to the main app
app.add_typer(readme_app, name="readme", help="Generate README for the repo")

if __name__ == "__main__":
    app()
