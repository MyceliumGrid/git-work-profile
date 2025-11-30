# Quick Start Guide

[中文](QUICKSTART.md) | English

## Get Started in 5 Minutes

### 1. Installation

```bash
# Install using go install
go install github.com/kway-teow/git-work-profile/cmd/git-work-profile@latest

# Or clone the repository and build
git clone https://github.com/kway-teow/git-work-profile.git
cd git-work-profile
go build -o git-work-profile ./cmd/git-work-profile/
```

### 2. Language Settings (Optional)

```bash
# Use Chinese interface (default)
git-work-profile

# Use English interface
export GIT_PROFILE_LANG=en
git-work-profile

# Or temporarily use English
GIT_PROFILE_LANG=en git-work-profile
```

### 3. First Run (Interactive Mode)

```bash
# Run directly to start interactive configuration
git-work-profile
```

The interactive interface will:
- Prompt for API key (if not set)
- Guide you to select analysis type
- Select time range and repository
- Configure output format

### 4. Or Use Command Line Mode

```bash
# First set the API key
export GEMINI_API_KEY="your-api-key-here"

# Run in any Git repository directory
cd ~/your-project
git-work-profile --analysis profile --output my-profile.md
```

### 5. View Results

The report will be saved to the specified file or output to the terminal.

## Common Commands

### Analyze All Personal Projects

```bash
# Assuming all your projects are in ~/projects directory
git-work-profile --repos ~/projects --output full-profile.md
```

### Generate Project Experience for Resume

```bash
git-work-profile --repos ~/work --range 2y --analysis experience --output resume.md
```

### View Tech Stack

```bash
git-work-profile --analysis techstack
```

### Generate JSON Data

```bash
git-work-profile --format json --output profile.json
```

## Next Steps

- Check [README_EN.md](README_EN.md) for complete features
- Check [EXAMPLES.md](EXAMPLES.md) for more use cases
- Adjust analysis time range and type based on output

## Troubleshooting

### Issue: "GEMINI_API_KEY environment variable not set"

**Solution**:
```bash
export GEMINI_API_KEY="your-api-key"
```

### Issue: No commits found

**Solution**:
- Check if you're in a Git repository directory
- Try increasing the time range: `--range 1y`
- Check if wrong author is specified: `--author "Your Name"`

### Issue: Analysis takes too long

**Solution**:
- Reduce time range: `--range 3m`
- Analyze single repository instead of multiple: `--repo` instead of `--repos`

### Issue: Output doesn't meet expectations

**Solution**:
- Try different analysis types: `--analysis profile/experience/techstack`
- Increase number of commits (expand time range)
- Ensure commit messages are clear and explicit

## Tips

1. **Permanently set API key**: Add `export GEMINI_API_KEY="..."` to `~/.zshrc` or `~/.bashrc`
2. **Create alias**: `alias profile='git-work-profile --repos ~/projects'`
3. **Regular updates**: Run once a month to track technical growth
4. **Multiple format output**: Generate both Markdown and JSON for different purposes
