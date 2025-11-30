# Git Developer Profile Analyzer

An intelligent tool for analyzing developer profiles and project experience based on Git commit history, powered by Google Gemini AI for deep analysis.

[中文文档](README.md) | English

> 💡 **Quick Start**: Check out [Quick Start Guide](QUICKSTART_EN.md) for a 5-minute guide

## Features

- 🌍 **Multi-language Support**: Complete Chinese and English bilingual interface
- 🤖 Automatically analyze Git commit history to generate developer profiles
- 🧠 Intelligent analysis using Google Gemini AI
- 📦 Multiple analysis modes:
  - **Single Repository**: Analyze a specific Git repository
  - **Multiple Repositories**: Automatically discover and analyze all Git repositories in a directory
- 👤 Multi-dimensional developer profile:
  - Tech stack analysis (programming languages, frameworks, tools)
  - Work habits (commit frequency, time distribution, code style)
  - Professional domain (frontend/backend/fullstack/DevOps, etc.)
  - Contribution statistics (code volume, activity, project participation)
- 📝 Project experience summary:
  - Types and scale of projects participated in
  - Technical challenges solved
  - Technical growth trajectory
  - Core competencies
- 📄 Multiple output formats (Text, Markdown, JSON)
- ⏰ Flexible time range selection (last 3 months, 6 months, 1 year, or custom)

## Installation

```bash
go install github.com/kway-teow/git-work-profile/cmd/git-work-profile@latest
```

## Usage

### Interactive Mode (Recommended)

Run the command directly to configure through an interactive interface:

```bash
# Start interactive configuration
git-work-profile
```

Interactive mode will guide you through:
1. Select analysis type (developer profile/project experience/tech stack)
2. Select time range (3 months/6 months/1 year/2 years)
3. Select repository mode (current directory/single repository/multiple repositories)
4. Select output format (Markdown/JSON/Text)
5. Set output file and author filter

If API key is not set, you will be prompted to enter it.

### Multi-language Support

The tool supports Chinese and English bilingual interface, switch via environment variable:

```bash
# Use Chinese interface (default)
git-work-profile

# Use English interface
export GIT_PROFILE_LANG=en
git-work-profile

# Or temporarily use English
GIT_PROFILE_LANG=en git-work-profile --help
```

### Command Line Mode

```bash
# Set Gemini API key
export GEMINI_API_KEY="your-api-key"

# Generate developer profile (default: analyze last 6 months)
git-work-profile --analysis profile

# Generate complete developer profile and project experience
git-work-profile --analysis profile

# Generate project experience summary
git-work-profile --analysis experience

# Generate tech stack analysis
git-work-profile --analysis techstack

# Specify analysis time range
git-work-profile --range 3m   # Last 3 months
git-work-profile --range 6m   # Last 6 months (default)
git-work-profile --range 1y   # Last 1 year
git-work-profile --range 2y   # Last 2 years

# Custom date range
git-work-profile --from 2024-01-01 --to 2025-11-30

# Specify output format
git-work-profile --format markdown
git-work-profile --format json

# Specify output file
git-work-profile --output developer-profile.md

# Analyze specific repository
git-work-profile --repo /path/to/your/repo

# Analyze multiple repositories (recommended)
git-work-profile --repos /path/to/projects

# Specify developer
git-work-profile --author "Your Name"

# Complete example
git-work-profile --repos ~/projects --range 1y --analysis profile --format markdown --output my-profile.md
```

## Command Line Options

```
Usage:
  git-work-profile [flags]

Flags:
  --analysis string  Analysis type (profile=developer profile, experience=project experience, techstack=tech stack) (default "profile")
  --author string    Git author name (default: current user)
  --from string      Start date (YYYY-MM-DD format)
  --to string        End date (YYYY-MM-DD format)
  --range string     Time range (3m=3 months, 6m=6 months, 1y=1 year, 2y=2 years) (default "6m")
  --format string    Output format (text, markdown, json) (default "markdown")
  --output string    Output file path (default: stdout)
  --repo string      Git repository path (default: current directory)
  --repos string     Repository directory path, analyze all Git repos in this directory
  --model string     Gemini model name (default: gemini-2.5-pro)
  -h, --help         Show help information
```

## Analysis Types

### Developer Profile (profile)
Comprehensive analysis of developer's technical capabilities and work characteristics:
- Tech stack profile (languages, frameworks, tools)
- Work style analysis (commit habits, code quality)
- Professional domain positioning (frontend/backend/fullstack, etc.)
- Core competency identification
- Technical growth trajectory
- Collaboration ability assessment

### Project Experience (experience)
Summarize developer's project experience and practical abilities:
- Types and scale of projects participated in
- Technical practice experience
- Engineering capability demonstration
- Business understanding depth
- Project highlights and contributions
- Reusable technical solutions

### Tech Stack Analysis (techstack)
In-depth analysis of developer's tech stack composition:
- Programming language proficiency assessment
- Framework and library usage
- Development toolchain mastery
- Infrastructure and DevOps capabilities
- Database and storage solutions
- Frontend and backend technology breakdown
- Tech stack modernization level

## Use Cases

### Resume Optimization
Use developer profile and project experience analysis to quickly generate technical skills and project experience sections for your resume:
```bash
git-work-profile --repos ~/projects --range 2y --analysis experience --output resume-projects.md
```

### Technical Interview Preparation
Fully understand your tech stack and project experience to prepare for technical interviews:
```bash
git-work-profile --repos ~/work --range 1y --analysis profile --format markdown
```

### Annual Technical Summary
Generate annual technical growth report to review technical development trajectory:
```bash
git-work-profile --repos ~/projects --range 1y --analysis profile --output annual-review.md
```

### Skill Assessment
Understand your tech stack composition and technical breadth:
```bash
git-work-profile --repos ~/projects --range 6m --analysis techstack
```

### Team Member Assessment
Analyze team members' technical capabilities and contributions (requires specifying author):
```bash
git-work-profile --repos /team/projects --author "Team Member" --range 6m --analysis profile
```

## Configuration

### API Key

The tool requires a Google Gemini API key to run. To get an API key:
1. Visit [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Create or select a project
3. Generate an API key

Set environment variable:
```bash
export GEMINI_API_KEY="your-api-key"
```

Or set permanently in `~/.zshrc` or `~/.bashrc`:
```bash
echo 'export GEMINI_API_KEY="your-api-key"' >> ~/.zshrc
source ~/.zshrc
```

### Gemini Model

By default, uses the `gemini-2.5-pro` model, which is the latest high-performance model.

You can also specify other models using the `--model` parameter:
```bash
git-work-profile --model gemini-pro
```

## Output Formats

### Markdown Format (Recommended)
Generate formatted Markdown documents with emoji icons and clear structure:
```bash
git-work-profile --format markdown --output profile.md
```

### JSON Format
Generate structured JSON data for programmatic processing:
```bash
git-work-profile --format json --output profile.json
```

### Text Format
Generate plain text reports suitable for terminal viewing:
```bash
git-work-profile --format text
```

## Examples

- See [EXAMPLES.md](EXAMPLES.md) for more usage examples and real-world scenarios
- See [INTERACTIVE_MODE.md](INTERACTIVE_MODE.md) for detailed interactive mode instructions

## FAQ

### How to get a Gemini API key?
Visit [Google AI Studio](https://makersuite.google.com/app/apikey) to create an API key.

### How long does analysis take?
Depends on the number of commits, usually a few seconds to tens of seconds. Large numbers of commits (1000+) may take 1-2 minutes.

### Which Git hosting platforms are supported?
Supports all Git repositories, including GitHub, GitLab, Bitbucket, etc., as long as they are locally cloned repositories.

### Can I analyze private repositories?
Yes, the tool only reads commit history from local Git repositories and does not upload code to any server.

### How accurate is the AI analysis?
AI analysis is based on multi-dimensional data such as commit content, file types, commit frequency, etc., and has high accuracy. However, it is recommended to adjust based on actual circumstances.

## Contributing

Issues and Pull Requests are welcome!

## License

MIT
