# MDTool

> Your automated README wizard for GitHub repositories

MDTool is a CLI Tool to generate README files for your Github repositories.

## Key Features

- **Automatic README Generation**: Analyzes your repository structure and generates detailed documentation
- **Smart Analysis**: Identifies key components, dependencies, and project structure
- **Customizable Output**: Multiple output formats and styling options
- **Batch Processing**: Efficiently handles repositories of any size
- **Rich Terminal Output**: Beautiful, colored terminal output for better visibility

## 🚀 Installation

1. Clone the repository:

```bash
git clone https://github.com/satuiro/mdtool.git
cd mdtool
```

2. Create and activate a virtual environment:

```bash
# Linux/macOS
python -m venv venv
source venv/bin/activate

# Windows
python -m venv venv
.\venv\Scripts\activate
```

3. Install the package in editable mode:

```bash
pip install -e .
```

4. Set up environment variables:

```bash
export GROQ_API_KEY="your-groq-api-key"
export GITHUB_TOKEN="your-github-token"
```

## 📖 Usage

### Generate README

```bash
# Preview README for a repository
mdtool readme generate --repo username/repo

# Generate and save README
mdtool readme generate --repo username/repo --save

# Generate README in raw format
mdtool readme generate --repo username/repo --output raw
```

### Options

```
--repo          Repository path or GitHub URL
--save          Save the generated README to file
--output        Output format (default: pretty, options: raw, pretty)
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

Distributed under the Apache License. See `LICENSE` for more information.

## ⭐ Show your support

Give a ⭐️ if this project helped you!
