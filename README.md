# Git Developer Profile Analyzer

一个基于Git提交记录分析开发者画像和项目经验的智能工具，使用Google Gemini AI进行深度分析。

中文 | [English](README_EN.md)

> 💡 **快速开始**: 查看 [快速开始指南](QUICKSTART.md) 5分钟上手

## 功能

- 🌍 **多语言支持**：完整的中英文双语界面
- 🤖 自动分析Git提交记录，生成开发者画像
- 🧠 使用Google Gemini AI进行智能分析
- 📦 支持多种分析模式：
  - **单仓库分析**：分析指定的单个Git仓库
  - **多仓库分析**：自动发现并分析目录下的所有Git仓库
- 👤 多维度开发者画像：
  - 技术栈分析（编程语言、框架、工具）
  - 工作习惯（提交频率、时间分布、代码风格）
  - 专业领域（前端/后端/全栈/DevOps等）
  - 贡献统计（代码量、活跃度、项目参与）
- 📝 项目经验总结：
  - 参与的项目类型和规模
  - 解决的技术难题
  - 技术成长轨迹
  - 核心竞争力
- 📄 支持多种输出格式（文本、Markdown、JSON）
- ⏰ 灵活的时间范围选择（最近3个月、6个月、1年或自定义）

## 安装

```bash
go install github.com/kway-teow/git-work-profile/cmd/git-work-profile@latest
```

## 使用方法

### 交互式模式（推荐）

直接运行命令，通过交互式界面配置：

```bash
# 启动交互式配置
git-work-profile
```

交互式模式会引导你：
1. 选择分析类型（开发者画像/项目经验/技术栈）
2. 选择时间范围（3个月/6个月/1年/2年）
3. 选择仓库模式（当前目录/单个仓库/多仓库）
4. 选择输出格式（Markdown/JSON/Text）
5. 设置输出文件和作者筛选

如果未设置API密钥，会提示输入。

### 多语言支持

工具支持中英文双语界面，通过环境变量切换：

```bash
# 使用中文界面（默认）
git-work-profile

# 使用英文界面
export GIT_PROFILE_LANG=en
git-work-profile

# 或者临时使用英文
GIT_PROFILE_LANG=en git-work-profile --help
```

### 命令行模式

```bash
# 设置Gemini API密钥
export GEMINI_API_KEY="your-api-key"

# 生成开发者画像（默认分析最近6个月）
git-work-profile --analysis profile

# 生成完整的开发者画像和项目经验
git-work-profile --analysis profile

# 生成项目经验总结
git-work-profile --analysis experience

# 生成技术栈分析
git-work-profile --analysis techstack

# 指定分析时间范围
git-work-profile --range 3m   # 最近3个月
git-work-profile --range 6m   # 最近6个月（默认）
git-work-profile --range 1y   # 最近1年
git-work-profile --range 2y   # 最近2年

# 自定义日期范围
git-work-profile --from 2024-01-01 --to 2025-11-30

# 指定输出格式
git-work-profile --format markdown
git-work-profile --format json

# 指定输出文件
git-work-profile --output developer-profile.md

# 分析指定仓库
git-work-profile --repo /path/to/your/repo

# 分析多个仓库（推荐）
git-work-profile --repos /path/to/projects

# 指定开发者
git-work-profile --author "Your Name"

# 完整示例
git-work-profile --repos ~/projects --range 1y --analysis profile --format markdown --output my-profile.md
```

## 命令行选项

```
Usage:
  git-work-profile [flags]

Flags:
  --analysis string  分析类型 (profile=开发者画像, experience=项目经验, techstack=技术栈) (default "profile")
  --author string    Git作者名称 (默认使用当前用户名)
  --from string      开始日期 (YYYY-MM-DD 格式)
  --to string        结束日期 (YYYY-MM-DD 格式)
  --range string     时间范围 (3m=3个月, 6m=6个月, 1y=1年, 2y=2年) (default "6m")
  --format string    输出格式 (text, markdown, json) (default "markdown")
  --output string    输出文件路径 (默认为标准输出)
  --repo string      Git仓库路径 (默认为当前目录)
  --repos string     仓库目录路径，分析该目录下的所有Git仓库
  --model string     Gemini模型名称 (默认为gemini-2.5-pro)
  -h, --help         显示帮助信息
```

## 分析类型说明

### 开发者画像 (profile)
全面分析开发者的技术能力和工作特点：
- 技术栈画像（语言、框架、工具）
- 工作风格分析（提交习惯、代码质量）
- 专业领域定位（前端/后端/全栈等）
- 核心竞争力识别
- 技术成长轨迹
- 协作能力评估

### 项目经验 (experience)
总结开发者的项目经验和实践能力：
- 参与的项目类型和规模
- 技术实践经验
- 工程能力体现
- 业务理解深度
- 项目亮点和贡献
- 可复用的技术方案

### 技术栈分析 (techstack)
深度分析开发者的技术栈构成：
- 编程语言能力评估
- 框架和库的使用
- 开发工具链掌握
- 基础设施和运维能力
- 数据库和存储方案
- 前后端技术细分
- 技术栈现代化程度

## 使用场景

### 个人简历优化
使用开发者画像和项目经验分析，快速生成简历中的技术能力和项目经验部分：
```bash
git-work-profile --repos ~/projects --range 2y --analysis experience --output resume-projects.md
```

### 技术面试准备
全面了解自己的技术栈和项目经验，为技术面试做准备：
```bash
git-work-profile --repos ~/work --range 1y --analysis profile --format markdown
```

### 年度技术总结
生成年度技术成长报告，回顾技术发展轨迹：
```bash
git-work-profile --repos ~/projects --range 1y --analysis profile --output annual-review.md
```

### 技能评估
了解自己的技术栈构成和技术广度：
```bash
git-work-profile --repos ~/projects --range 6m --analysis techstack
```

### 团队成员评估
分析团队成员的技术能力和贡献（需要指定作者）：
```bash
git-work-profile --repos /team/projects --author "Team Member" --range 6m --analysis profile
```

## 配置

### API密钥

工具需要Google Gemini API密钥才能运行。获取API密钥：
1. 访问 [Google AI Studio](https://makersuite.google.com/app/apikey)
2. 创建或选择一个项目
3. 生成API密钥

设置环境变量：
```bash
export GEMINI_API_KEY="your-api-key"
```

或者在 `~/.zshrc` 或 `~/.bashrc` 中永久设置：
```bash
echo 'export GEMINI_API_KEY="your-api-key"' >> ~/.zshrc
source ~/.zshrc
```

### Gemini模型

默认使用 `gemini-2.5-pro` 模型，这是最新的高性能模型。

你也可以使用 `--model` 参数指定其他模型：
```bash
git-work-profile --model gemini-pro
```

## 输出格式

### Markdown格式（推荐）
生成格式化的Markdown文档，包含emoji图标和清晰的结构：
```bash
git-work-profile --format markdown --output profile.md
```

### JSON格式
生成结构化的JSON数据，便于程序处理：
```bash
git-work-profile --format json --output profile.json
```

### 文本格式
生成纯文本报告，适合终端查看：
```bash
git-work-profile --format text
```

## 示例

- 查看 [EXAMPLES.md](EXAMPLES.md) 了解更多使用示例和实际场景
- 查看 [INTERACTIVE_MODE.md](INTERACTIVE_MODE.md) 了解交互式模式详细说明

## 常见问题

### 如何获取Gemini API密钥？
访问 [Google AI Studio](https://makersuite.google.com/app/apikey) 创建API密钥。

### 分析需要多长时间？
取决于提交记录的数量，通常几秒到几十秒。大量提交（1000+）可能需要1-2分钟。

### 支持哪些Git托管平台？
支持所有Git仓库，包括GitHub、GitLab、Bitbucket等，只要是本地克隆的仓库即可。

### 可以分析私有仓库吗？
可以，工具只读取本地Git仓库的提交记录，不会上传代码到任何服务器。

### AI分析的准确性如何？
AI分析基于提交记录的内容、文件类型、提交频率等多维度数据，准确性较高。但建议结合实际情况进行调整。

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT
