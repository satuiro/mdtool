# MDTool
================

> Your automated README wizard for GitHub repositories

MDTool is a CLI tool designed to simplify the process of generating high-quality README files for your GitHub repositories. With its user-friendly interface and automatic generation capabilities, MDTool is the perfect solution for developers looking to create engaging and informative READMEs.

## Key Features
---------------

*   **Automatic README Generation**: Analyze your repository and generate a comprehensive README file with essential information.
*   **Customization Options**: Tailor your README to fit your project's unique needs with various customization options.
*   **Rich Text Formatting**: Use Markdown formatting to create visually appealing and easy-to-read READMEs.
*   **Support for Multiple Repository Types**: Generate READMEs for various types of repositories, including open-source projects, personal projects, and more.

## Installation
------------

To install MDTool, follow these steps:

### Using pip

```bash
pip install mdtool
```

### From Source

1.  Clone the repository:

    ```bash
git clone https://github.com/your-username/mdtool.git
```

2.  Navigate to the project directory:

    ```bash
cd mdtool
```

3.  Install the dependencies:

    ```bash
pip install -r requirements.txt
```

4.  Install the package:

    ```bash
pip install .
```

## Usage
-----

To generate a README file using MDTool, follow these steps:

1.  Navigate to your repository directory:

    ```bash
cd your-repo
```

2.  Run the MDTool command:

    ```bash
mdtool generate
```

This will prompt you to provide some basic information about your repository. Once you've entered the required details, MDTool will generate a comprehensive README file for your repository.

### Example Usage

Here's an example of how to use MDTool to generate a README file:

```bash
$ mdtool generate
? What is the name of your repository? My Awesome Project
? What is the description of your repository? A brief description of my project
? What are the keywords for your repository? (comma-separated) project, awesome, github
? What is the license for your repository? MIT
? What are the contributors for your repository? (comma-separated) John Doe, Jane Doe

README file generated successfully!
```

## Project Structure
------------------

The MDTool project is organized into the following directories and files:

*   `mdtool/`: The main package directory containing the CLI tool and supporting modules.
*   `mdtool/cli.py`: The CLI tool entry point.
*   `mdtool/readme.py`: The module responsible for generating README files.
*   `mdtool/__init__.py`: The package initialization file.
*   `LICENSE`: The project license file.
*   `README.md`: This README file.

## Dependencies and Requirements
-------------------------------

MDTool depends on the following packages:

*   `typer`: A Python library for building CLI tools.
*   `rich`: A Python library for rich text formatting.

These dependencies are specified in the `requirements.txt` file.

## Contributing
------------

Contributions to MDTool are welcome! If you'd like to contribute to the project, please follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and commit them.
4.  Open a pull request to the main repository.

Please ensure that your contributions adhere to the project's code style and conventions.

## License
-------

MDTool is released under the Apache License 2.0. See the `LICENSE` file for more information.

By using MDTool, you agree to the terms and conditions outlined in the license.

## Acknowledgments
---------------

MDTool is built on top of several open-source libraries and tools. We'd like to extend our gratitude to the developers and maintainers of these projects for their hard work and dedication.

*   `typer`: A Python library for building CLI tools.
*   `rich`: A Python library for rich text formatting.

We're grateful for the opportunity to build upon their work and create a useful tool for the developer community.

# mdtool
================

**A Markdown Tool for Developers**
--------------------------------

**Description**
---------------

mdtool is a Python-based command-line tool designed to assist developers in managing and generating Markdown content. With its Groq and GitHub integrations, mdtool streamlines the process of creating, reading, and updating Markdown files.

**Key Features**
-------------

*   **Groq Integration**: Seamlessly interact with Groq's API to generate and manage Markdown content.
*   **GitHub Integration**: Utilize GitHub tokens to access and manage Markdown files within your repositories.
*   **Configurable**: Easily customize the tool to suit your needs through environment variables and a dataclass-based configuration system.
*   **Rich CLI Output**: Enjoy a visually appealing command-line interface with rich formatting and progress bars.

**Installation**
------------

To install mdtool, follow these steps:

1.  Clone the repository:
    ```bash
git clone https://github.com/your-username/mdtool.git
```
2.  Navigate to the project directory:
    ```bash
cd mdtool
```
3.  Install the required dependencies using pip:
    ```bash
pip install -r requirements.txt
```
4.  Install the mdtool package:
    ```bash
pip install .
```

**Usage Examples**
-----------------

To use mdtool, run the following commands:

*   **Generate a new Markdown file**:
    ```bash
mdtool generate --title "My Markdown File" --content "This is a sample Markdown file."
```
*   **Read a Markdown file**:
    ```bash
mdtool read --path "/path/to/your/markdown/file.md"
```
*   **Update a Markdown file**:
    ```bash
mdtool update --path "/path/to/your/markdown/file.md" --title "My Updated Markdown File" --content "This is an updated sample Markdown file."
```

**Project Structure Overview**
-----------------------------

The mdtool project is organized into the following directories and files:

*   `mdtool/`: The main package directory.
    *   `config.py`: Defines the configuration dataclass and loads environment variables.
    *   `groq_service.py`: Provides a service class for interacting with Groq's API.
    *   `readme.py`: Contains the main command-line interface and functionality.
*   `pyproject.toml`: Specifies the project's build system and metadata.

**Dependencies and Requirements**
-------------------------------

mdtool relies on the following dependencies:

*   `groq`: A Python client for Groq's API.
*   `typer`: A Python library for building command-line interfaces.
*   `rich`: A Python library for rich text formatting.
*   `setuptools`: A build system for Python packages.

**Contributing Guidelines**
---------------------------

Contributions to mdtool are welcome! To contribute, follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes and test them thoroughly.
4.  Open a pull request to merge your branch into the main repository.

**License Information**
----------------------

mdtool is released under the Apache License 2.0. See the `LICENSE` file for more information.

**Acknowledgments**
-----------------

Special thanks to the developers of Groq, typer, rich, and setuptools for their excellent libraries and tools.