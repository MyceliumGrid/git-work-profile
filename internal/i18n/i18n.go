package i18n

import (
	"os"
)

// Language 语言类型
type Language string

const (
	// English 英语
	English Language = "en"
	// Chinese 中文
	Chinese Language = "zh"
)

// 环境变量名
const LangEnvVar = "GIT_PROFILE_LANG"

// Messages 多语言消息
type Messages struct {
	// 交互式界面
	InteractiveTitle   string
	SelectAnalysisType string
	SelectTimeRange    string
	SelectRepoMode     string
	SelectOutputFormat string
	InputOutputFile    string
	InputAuthor        string
	ConfigSummary      string
	ConfirmAnalysis    string

	// 分析类型
	AnalysisProfile    string
	AnalysisExperience string
	AnalysisTechStack  string

	// 时间范围
	TimeRange3Months string
	TimeRange6Months string
	TimeRange1Year   string
	TimeRange2Years  string
	TimeRangeCustom  string

	// 仓库模式
	RepoModeCurrent  string
	RepoModeSingle   string
	RepoModeMultiple string

	// 输出格式
	FormatMarkdown string
	FormatJSON     string
	FormatText     string

	// 提示信息
	InputRepoPath  string
	InputReposPath string
	OutputFileHint string
	AuthorHint     string

	// 配置摘要标签
	LabelAnalysisType   string
	LabelTimeRange      string
	LabelRepoPath       string
	LabelOutputFormat   string
	LabelOutputFile     string
	LabelAuthorFilter   string
	LabelTerminalOutput string

	// 自定义日期
	InputFromDate string
	InputToDate   string

	// API密钥相关
	APIKeyNotFound     string
	APIKeyInstructions string
	InputAPIKey        string
	APIKeyEmpty        string
	APIKeyInvalid      string
	SaveAPIKey         string
	SaveAPIKeyHint     string

	// 错误消息
	ErrorAPIKeyNotSet           string
	ErrorCreateClient           string
	ErrorDateFormat             string
	ErrorDiscoverRepos          string
	ErrorNoReposFound           string
	ErrorNoCommitsFound         string
	ErrorAIAnalysisFailed       string
	ErrorOutputFailed           string
	ErrorCancelled              string
	ErrorDateEmpty              string
	ErrorDateFormatInvalid      string
	ErrorPathEmpty              string
	ErrorPathNotExist           string
	ErrorPathNotAccessible      string
	ErrorPathNotDirectory       string
	ErrorNotGitRepo             string
	ErrorCurrentDirNotGitRepo   string
	ErrorDirectoryNotExist      string
	ErrorDirectoryNotAccessible string
	ErrorRepoWarning            string

	// 信息提示
	InfoProcessingRepos  string
	InfoAnalyzingRepo    string
	InfoFoundCommits     string
	InfoCommitStats      string
	InfoTotalCommits     string
	InfoFilterAuthor     string
	InfoAllAuthors       string
	InfoAIAnalyzing      string
	InfoAnalysisComplete string
	InfoReportSaved      string
	InfoTimeRange        string
	InfoCustomTimeRange  string

	// 分析类型标签
	LabelAnalysisTypeProfile    string
	LabelAnalysisTypeExperience string
	LabelAnalysisTypeTechStack  string
	LabelAnalysisTypeDefault    string

	// 版本信息
	VersionInfo string
	CommitHash  string
	BuildDate   string

	// Git扫描相关
	InfoScanningDirectory string
	InfoFoundGitRepo      string
	InfoScanComplete      string

	// Report相关
	InfoReportGenerated        string
	ReportTitleProfile         string
	ReportTitleExperience      string
	ReportTitleTechStack       string
	ReportTitleDefault         string
	ReportTimeRange            string
	ReportTo                   string
	ReportGeneratedAt          string
	ReportDataStats            string
	ReportTotalCommits         string
	ReportTotalRepos           string
	ReportTotalFiles           string
	ReportRepoUnit             string
	ReportFileUnit             string
	ReportFileTypeDistribution string
	ReportAIAnalysis           string
	ReportFooter               string

	// Cobra命令描述
	CmdShortDesc    string
	CmdLongDesc     string
	CmdVersionShort string

	// 命令行参数描述
	FlagAnalysis string
	FlagFrom     string
	FlagTo       string
	FlagRange    string
	FlagFormat   string
	FlagOutput   string
	FlagRepo     string
	FlagRepos    string
	FlagModel    string
	FlagAuthor   string

	// 其他错误
	ErrorCreateOutputFile string

	// 警告消息
	WarningPromptLoadFailed string

	// Gemini相关错误
	ErrorGeminiClientFailed string
	ErrorGeminiAPIFailed    string
	ErrorLoadPromptTemplate string

	// 文件和目录错误
	ErrorGetCurrentDir      string
	ErrorPromptFileNotExist string
	ErrorReadPromptFile     string
	ErrorPromptFileEmpty    string
	ErrorGetAbsolutePath    string
	ErrorWalkDirectory      string

	// Git相关错误
	ErrorGitLogFailed       string
	ErrorGetCommitDetails   string
	ErrorParseCommitDetails string
	ErrorParseDateFailed    string
	ErrorGetChangedFiles    string
	ErrorGetGitUsername     string

	// 其他
	Canceled         string
	AnalysisStarting string
}

var (
	// 当前语言
	currentLang = English

	// 英文消息
	englishMessages = Messages{
		InteractiveTitle:   "🎨 Git Developer Profile Analyzer - Interactive Configuration",
		SelectAnalysisType: "Select analysis type",
		SelectTimeRange:    "Select time range",
		SelectRepoMode:     "Select repository mode",
		SelectOutputFormat: "Select output format",
		InputOutputFile:    "Output file path (leave empty for terminal output)",
		InputAuthor:        "Specify Git author name (default: %s, leave empty for all authors)",
		ConfigSummary:      "📋 Configuration Summary",
		ConfirmAnalysis:    "Confirm to start analysis (Y/n)",

		AnalysisProfile:    "Developer Profile - Comprehensive analysis of technical capabilities",
		AnalysisExperience: "Project Experience - Summarize project experience for resume",
		AnalysisTechStack:  "Tech Stack Analysis - In-depth analysis of technology stack",

		TimeRange3Months: "Last 3 months (3m)",
		TimeRange6Months: "Last 6 months (6m) - Recommended",
		TimeRange1Year:   "Last 1 year (1y)",
		TimeRange2Years:  "Last 2 years (2y)",
		TimeRangeCustom:  "Custom date range",

		RepoModeCurrent:  "Current directory - Analyze current Git repository",
		RepoModeSingle:   "Single repository - Analyze specified Git repository",
		RepoModeMultiple: "Multiple repositories - Analyze all Git repos in directory (Recommended)",

		FormatMarkdown: "Markdown - Formatted document (Recommended)",
		FormatJSON:     "JSON - Structured data",
		FormatText:     "Text - Plain text",

		InputRepoPath:  "Enter repository path",
		InputReposPath: "Enter directory path (containing multiple Git repositories)",
		OutputFileHint: "Suggested: %s",
		AuthorHint:     "Current user: %s",

		LabelAnalysisType:   "Analysis Type",
		LabelTimeRange:      "Time Range",
		LabelRepoPath:       "Repository Path",
		LabelOutputFormat:   "Output Format",
		LabelOutputFile:     "Output File",
		LabelAuthorFilter:   "Author Filter",
		LabelTerminalOutput: "Terminal output",

		InputFromDate: "Enter start date (YYYY-MM-DD)",
		InputToDate:   "Enter end date (YYYY-MM-DD)",

		Canceled:         "Canceled",
		AnalysisStarting: "🚀 Starting analysis...",
	}

	// 中文消息
	chineseMessages = Messages{
		InteractiveTitle:   "🎨 Git Developer Profile Analyzer - 交互式配置",
		SelectAnalysisType: "选择分析类型",
		SelectTimeRange:    "选择分析时间范围",
		SelectRepoMode:     "选择仓库分析模式",
		SelectOutputFormat: "选择输出格式",
		InputOutputFile:    "输出文件路径（留空则输出到终端）",
		InputAuthor:        "指定Git作者名称（默认: %s，留空则分析所有作者）",
		ConfigSummary:      "📋 配置摘要",
		ConfirmAnalysis:    "确认开始分析 (Y/n)",

		AnalysisProfile:    "开发者画像 - 全面分析技术能力和工作特点",
		AnalysisExperience: "项目经验 - 总结项目经验用于简历",
		AnalysisTechStack:  "技术栈分析 - 深度分析技术栈构成",

		TimeRange3Months: "最近3个月 (3m)",
		TimeRange6Months: "最近6个月 (6m) - 推荐",
		TimeRange1Year:   "最近1年 (1y)",
		TimeRange2Years:  "最近2年 (2y)",
		TimeRangeCustom:  "自定义日期范围",

		RepoModeCurrent:  "当前目录 - 分析当前Git仓库",
		RepoModeSingle:   "指定单个仓库 - 分析指定的Git仓库",
		RepoModeMultiple: "多仓库目录 - 分析目录下所有Git仓库（推荐）",

		FormatMarkdown: "Markdown - 格式化文档（推荐）",
		FormatJSON:     "JSON - 结构化数据",
		FormatText:     "Text - 纯文本",

		InputRepoPath:  "输入仓库路径",
		InputReposPath: "输入仓库目录路径（包含多个Git仓库）",
		OutputFileHint: "建议: %s",
		AuthorHint:     "当前用户: %s",

		LabelAnalysisType:   "分析类型",
		LabelTimeRange:      "时间范围",
		LabelRepoPath:       "仓库路径",
		LabelOutputFormat:   "输出格式",
		LabelOutputFile:     "输出文件",
		LabelAuthorFilter:   "作者筛选",
		LabelTerminalOutput: "终端输出",

		InputFromDate: "输入开始日期 (YYYY-MM-DD)",
		InputToDate:   "输入结束日期 (YYYY-MM-DD)",

		Canceled:         "已取消",
		AnalysisStarting: "🚀 开始分析...",
	}
)

// SetLanguage 设置语言
func SetLanguage(lang Language) {
	currentLang = lang
	// 同时设置环境变量
	os.Setenv(LangEnvVar, string(lang))
}

// GetLanguage 获取当前语言
func GetLanguage() Language {
	return currentLang
}

// LoadLanguageFromEnv 从环境变量加载语言设置
func LoadLanguageFromEnv() Language {
	envLang := os.Getenv(LangEnvVar)
	if envLang == string(Chinese) {
		return Chinese
	}
	return English
}

// GetMessages 获取当前语言的消息
func GetMessages() Messages {
	if currentLang == Chinese {
		return chineseMessages
	}
	return englishMessages
}

// T 翻译函数（简写）
func T() Messages {
	return GetMessages()
}

// API密钥相关消息
func init() {
	// 英文消息 - API密钥
	englishMessages.APIKeyNotFound = "⚠️  GEMINI_API_KEY environment variable not found"
	englishMessages.APIKeyInstructions = `How to get API key:
1. Visit https://makersuite.google.com/app/apikey
2. Create or select a project
3. Generate API key`
	englishMessages.InputAPIKey = "Enter Gemini API Key"
	englishMessages.APIKeyEmpty = "API key cannot be empty"
	englishMessages.APIKeyInvalid = "API key length is incorrect"
	englishMessages.SaveAPIKey = "Save API key to environment configuration file"
	englishMessages.SaveAPIKeyHint = `💡 Tip: Add the following command to ~/.zshrc or ~/.bashrc:
   export GEMINI_API_KEY="%s"`

	// 中文消息 - API密钥
	chineseMessages.APIKeyNotFound = "⚠️  未检测到 GEMINI_API_KEY 环境变量"
	chineseMessages.APIKeyInstructions = `获取API密钥:
1. 访问 https://makersuite.google.com/app/apikey
2. 创建或选择项目
3. 生成API密钥`
	chineseMessages.InputAPIKey = "请输入 Gemini API Key"
	chineseMessages.APIKeyEmpty = "API密钥不能为空"
	chineseMessages.APIKeyInvalid = "API密钥长度不正确"
	chineseMessages.SaveAPIKey = "是否将API密钥保存到环境变量配置文件"
	chineseMessages.SaveAPIKeyHint = `💡 提示: 将以下命令添加到 ~/.zshrc 或 ~/.bashrc:
   export GEMINI_API_KEY="%s"`
}

// 错误和提示消息
func init() {
	// 英文 - 错误和提示消息
	englishMessages.ErrorAPIKeyNotSet = "Error: GEMINI_API_KEY environment variable not set"
	englishMessages.ErrorCreateClient = "Error: Failed to create Gemini client: %v"
	englishMessages.ErrorDateFormat = "Error: Incorrect date format, please use YYYY-MM-DD"
	englishMessages.ErrorDiscoverRepos = "Error: Failed to discover Git repositories: %v"
	englishMessages.ErrorNoReposFound = "No Git repositories found in directory %s"
	englishMessages.ErrorNoCommitsFound = "No commits found in all repositories for time range %s to %s"
	englishMessages.ErrorAIAnalysisFailed = "Error: AI analysis failed: %v"
	englishMessages.ErrorOutputFailed = "Error: Failed to output report: %v"
	englishMessages.ErrorCancelled = "Operation canceled"

	englishMessages.InfoProcessingRepos = "\nProcessing %d repositories:"
	englishMessages.InfoAnalyzingRepo = "Analyzing repository: %s"
	englishMessages.InfoFoundCommits = "  Found %d commits"
	englishMessages.InfoCommitStats = "\n=== Commit Statistics ==="
	englishMessages.InfoTotalCommits = "Total: %d commits\n"
	englishMessages.InfoFilterAuthor = "Filtering author: %s"
	englishMessages.InfoAllAuthors = "Getting commits from all authors"
	englishMessages.InfoAIAnalyzing = "\nPerforming AI deep analysis..."
	englishMessages.InfoAnalysisComplete = "\n✓ Analysis complete!"
	englishMessages.InfoReportSaved = "Report saved to: %s"
	englishMessages.InfoTimeRange = "Analysis time range: %s (%s to %s)"
	englishMessages.InfoCustomTimeRange = "Analysis time range: %s to %s"

	englishMessages.LabelAnalysisTypeProfile = "Analysis type: Developer Profile"
	englishMessages.LabelAnalysisTypeExperience = "Analysis type: Project Experience"
	englishMessages.LabelAnalysisTypeTechStack = "Analysis type: Tech Stack Analysis"
	englishMessages.LabelAnalysisTypeDefault = "Analysis type: Developer Profile (default)"

	englishMessages.VersionInfo = "git-work-profile version: %s"
	englishMessages.CommitHash = "Commit hash: %s"
	englishMessages.BuildDate = "Build date: %s"

	// 验证错误消息
	englishMessages.ErrorDateEmpty = "Date cannot be empty"
	englishMessages.ErrorDateFormatInvalid = "Date format should be YYYY-MM-DD"
	englishMessages.ErrorPathEmpty = "Path cannot be empty"
	englishMessages.ErrorPathNotExist = "Path does not exist: %s"
	englishMessages.ErrorPathNotAccessible = "Cannot access path: %v"
	englishMessages.ErrorPathNotDirectory = "Path is not a directory: %s"
	englishMessages.ErrorNotGitRepo = "Not a Git repository: %s"
	englishMessages.ErrorCurrentDirNotGitRepo = "Current directory is not a Git repository: %v"
	englishMessages.ErrorDirectoryNotExist = "Directory does not exist: %s"
	englishMessages.ErrorDirectoryNotAccessible = "Cannot access directory: %v"
	englishMessages.ErrorRepoWarning = "  Warning: Repository %s failed to get Git commits: %v"

	// 中文 - 错误和提示消息
	chineseMessages.ErrorAPIKeyNotSet = "错误: 未设置GEMINI_API_KEY环境变量"
	chineseMessages.ErrorCreateClient = "错误: 创建Gemini客户端失败: %v"
	chineseMessages.ErrorDateFormat = "错误: 日期格式不正确，请使用YYYY-MM-DD格式"
	chineseMessages.ErrorDiscoverRepos = "错误: 发现Git仓库失败: %v"
	chineseMessages.ErrorNoReposFound = "在目录 %s 下没有发现任何Git仓库"
	chineseMessages.ErrorNoCommitsFound = "指定时间范围 %s 到 %s 在所有仓库中都没有找到提交记录"
	chineseMessages.ErrorAIAnalysisFailed = "错误: AI分析失败: %v"
	chineseMessages.ErrorOutputFailed = "错误: 输出报告失败: %v"
	chineseMessages.ErrorCancelled = "已取消操作"

	chineseMessages.InfoProcessingRepos = "\n处理 %d 个仓库:"
	chineseMessages.InfoAnalyzingRepo = "正在分析仓库: %s"
	chineseMessages.InfoFoundCommits = "  找到 %d 条提交记录"
	chineseMessages.InfoCommitStats = "\n=== 提交记录统计 ==="
	chineseMessages.InfoTotalCommits = "总计: %d 条提交\n"
	chineseMessages.InfoFilterAuthor = "筛选作者: %s"
	chineseMessages.InfoAllAuthors = "获取所有作者的提交"
	chineseMessages.InfoAIAnalyzing = "\n正在使用AI进行深度分析..."
	chineseMessages.InfoAnalysisComplete = "\n✓ 分析完成！"
	chineseMessages.InfoReportSaved = "报告已保存到: %s"
	chineseMessages.InfoTimeRange = "分析时间范围: %s (%s 到 %s)"
	chineseMessages.InfoCustomTimeRange = "分析时间范围: %s 到 %s"

	chineseMessages.LabelAnalysisTypeProfile = "分析类型: 开发者画像"
	chineseMessages.LabelAnalysisTypeExperience = "分析类型: 项目经验总结"
	chineseMessages.LabelAnalysisTypeTechStack = "分析类型: 技术栈分析"
	chineseMessages.LabelAnalysisTypeDefault = "分析类型: 开发者画像（默认）"

	chineseMessages.VersionInfo = "git-work-profile 版本: %s"
	chineseMessages.CommitHash = "提交哈希: %s"
	chineseMessages.BuildDate = "构建日期: %s"

	// 验证错误消息
	chineseMessages.ErrorDateEmpty = "日期不能为空"
	chineseMessages.ErrorDateFormatInvalid = "日期格式应为 YYYY-MM-DD"
	chineseMessages.ErrorPathEmpty = "路径不能为空"
	chineseMessages.ErrorPathNotExist = "路径不存在: %s"
	chineseMessages.ErrorPathNotAccessible = "无法访问路径: %v"
	chineseMessages.ErrorPathNotDirectory = "路径不是目录: %s"
	chineseMessages.ErrorNotGitRepo = "不是Git仓库: %s"
	chineseMessages.ErrorCurrentDirNotGitRepo = "当前目录不是Git仓库: %v"
	chineseMessages.ErrorDirectoryNotExist = "目录不存在: %s"
	chineseMessages.ErrorDirectoryNotAccessible = "无法访问目录: %v"
	chineseMessages.ErrorRepoWarning = "  警告: 仓库 %s 获取Git提交记录失败: %v"
}

// Git和Report相关消息
func init() {
	// 英文 - Git扫描消息
	englishMessages.InfoScanningDirectory = "Scanning directory: %s"
	englishMessages.InfoFoundGitRepo = "  Found Git repository: %s"
	englishMessages.InfoScanComplete = "Scan complete, found %d Git repositories"
	englishMessages.InfoReportGenerated = "Markdown report generated: %s"

	// 英文 - Report标题和内容
	englishMessages.ReportTitleProfile = "Developer Profile Analysis Report"
	englishMessages.ReportTitleExperience = "Project Experience Summary Report"
	englishMessages.ReportTitleTechStack = "Tech Stack Analysis Report"
	englishMessages.ReportTitleDefault = "Developer Analysis Report"
	englishMessages.ReportTimeRange = "Analysis Time Range"
	englishMessages.ReportTo = "to"
	englishMessages.ReportGeneratedAt = "Generated At"
	englishMessages.ReportDataStats = "Data Statistics"
	englishMessages.ReportTotalCommits = "Total Commits"
	englishMessages.ReportTotalRepos = "Repositories"
	englishMessages.ReportTotalFiles = "Changed Files"
	englishMessages.ReportRepoUnit = "repos"
	englishMessages.ReportFileUnit = "files"
	englishMessages.ReportFileTypeDistribution = "File Type Distribution"
	englishMessages.ReportAIAnalysis = "AI Deep Analysis"
	englishMessages.ReportFooter = "This report is automatically generated by Git Developer Profile Analyzer"

	// 中文 - Git扫描消息
	chineseMessages.InfoScanningDirectory = "正在扫描目录: %s"
	chineseMessages.InfoFoundGitRepo = "  发现Git仓库: %s"
	chineseMessages.InfoScanComplete = "扫描完成，共发现 %d 个Git仓库"
	chineseMessages.InfoReportGenerated = "已生成Markdown报告: %s"

	// 中文 - Report标题和内容
	chineseMessages.ReportTitleProfile = "开发者画像分析报告"
	chineseMessages.ReportTitleExperience = "项目经验总结报告"
	chineseMessages.ReportTitleTechStack = "技术栈分析报告"
	chineseMessages.ReportTitleDefault = "开发者分析报告"
	chineseMessages.ReportTimeRange = "分析时间范围"
	chineseMessages.ReportTo = "至"
	chineseMessages.ReportGeneratedAt = "生成时间"
	chineseMessages.ReportDataStats = "数据统计"
	chineseMessages.ReportTotalCommits = "总提交数"
	chineseMessages.ReportTotalRepos = "涉及仓库"
	chineseMessages.ReportTotalFiles = "变更文件"
	chineseMessages.ReportRepoUnit = "个"
	chineseMessages.ReportFileUnit = "个"
	chineseMessages.ReportFileTypeDistribution = "文件类型分布"
	chineseMessages.ReportAIAnalysis = "AI 深度分析"
	chineseMessages.ReportFooter = "本报告由 Git Developer Profile Analyzer 自动生成"
}

// Cobra命令描述
func init() {
	// 英文 - Cobra命令描述
	englishMessages.CmdShortDesc = "Analyze developer profile and project experience based on Git commit history"
	englishMessages.CmdLongDesc = `git-work-profile is an intelligent tool for analyzing developer profiles and project experience based on Git commit history.

It uses Google Gemini AI to perform deep analysis of commit records, generating developer profiles, project experience summaries, and tech stack analysis.
Supports multiple time ranges: 3 months (3m), 6 months (6m), 1 year (1y), 2 years (2y), or custom dates.
Supports single repository analysis (--repo) or multiple repositories analysis (--repos).
Default analysis of the last 6 months of commit records, generating developer profile.

Running without any parameters will start interactive configuration mode.`
	englishMessages.CmdVersionShort = "Show version information"

	// 中文 - Cobra命令描述
	chineseMessages.CmdShortDesc = "基于Git提交记录分析开发者画像和项目经验"
	chineseMessages.CmdLongDesc = `git-work-profile 是一个基于Git提交记录分析开发者画像和项目经验的智能工具。

它使用Google Gemini AI对提交记录进行深度分析，生成开发者画像、项目经验总结和技术栈分析。
支持多种时间范围：3个月(3m)、6个月(6m)、1年(1y)、2年(2y)或自定义日期。
支持单个仓库分析(--repo)或目录下所有仓库分析(--repos)。
默认分析最近6个月的提交记录，生成开发者画像。

不带任何参数运行将启动交互式配置模式。`
	chineseMessages.CmdVersionShort = "显示版本信息"
}

// 命令行参数描述
func init() {
	// 英文 - 命令行参数
	englishMessages.FlagAnalysis = "Analysis type (profile=developer profile, experience=project experience, techstack=tech stack)"
	englishMessages.FlagFrom = "Start date (YYYY-MM-DD format)"
	englishMessages.FlagTo = "End date (YYYY-MM-DD format)"
	englishMessages.FlagRange = "Time range (3m=3 months, 6m=6 months, 1y=1 year, 2y=2 years)"
	englishMessages.FlagFormat = "Output format (text, markdown, json)"
	englishMessages.FlagOutput = "Output file path (default: stdout)"
	englishMessages.FlagRepo = "Git repository path (default: current directory)"
	englishMessages.FlagRepos = "Repository directory path, analyze all Git repos in this directory"
	englishMessages.FlagModel = "Gemini model name (default: gemini-2.5-pro)"
	englishMessages.FlagAuthor = "Git author name"

	englishMessages.ErrorCreateOutputFile = "Error: Failed to create output file: %v"
	englishMessages.WarningPromptLoadFailed = "Warning: Failed to load prompt template: %v, using default prompt"

	// 英文 - Gemini相关错误
	englishMessages.ErrorGeminiClientFailed = "Failed to create Gemini client"
	englishMessages.ErrorGeminiAPIFailed = "Gemini API call failed"
	englishMessages.ErrorLoadPromptTemplate = "Failed to load prompt template"

	// 英文 - 文件和目录错误
	englishMessages.ErrorGetCurrentDir = "Failed to get current directory"
	englishMessages.ErrorPromptFileNotExist = "Prompt file does not exist"
	englishMessages.ErrorReadPromptFile = "Failed to read prompt file"
	englishMessages.ErrorPromptFileEmpty = "Prompt file is empty"
	englishMessages.ErrorGetAbsolutePath = "Failed to get absolute path"
	englishMessages.ErrorWalkDirectory = "Failed to walk directory"

	// 英文 - Git相关错误
	englishMessages.ErrorGitLogFailed = "Git log command failed"
	englishMessages.ErrorGetCommitDetails = "Failed to get commit details"
	englishMessages.ErrorParseCommitDetails = "Failed to parse commit details"
	englishMessages.ErrorParseDateFailed = "Failed to parse date"
	englishMessages.ErrorGetChangedFiles = "Failed to get changed files"
	englishMessages.ErrorGetGitUsername = "Failed to get Git username"

	// 中文 - 命令行参数
	chineseMessages.FlagAnalysis = "分析类型 (profile=开发者画像, experience=项目经验, techstack=技术栈)"
	chineseMessages.FlagFrom = "开始日期 (YYYY-MM-DD 格式)"
	chineseMessages.FlagTo = "结束日期 (YYYY-MM-DD 格式)"
	chineseMessages.FlagRange = "时间范围 (3m=3个月, 6m=6个月, 1y=1年, 2y=2年)"
	chineseMessages.FlagFormat = "输出格式 (text, markdown, json)"
	chineseMessages.FlagOutput = "输出文件路径 (默认为标准输出)"
	chineseMessages.FlagRepo = "Git仓库路径 (默认为当前目录)"
	chineseMessages.FlagRepos = "仓库目录路径，分析该目录下的所有Git仓库"
	chineseMessages.FlagModel = "Gemini模型名称 (默认为gemini-2.5-pro)"
	chineseMessages.FlagAuthor = "Git作者名称"

	chineseMessages.ErrorCreateOutputFile = "错误: 创建输出文件失败: %v"
	chineseMessages.WarningPromptLoadFailed = "警告: 加载提示词模板失败: %v, 使用默认提示词"

	// 中文 - Gemini相关错误
	chineseMessages.ErrorGeminiClientFailed = "创建Gemini客户端失败"
	chineseMessages.ErrorGeminiAPIFailed = "调用Gemini API失败"
	chineseMessages.ErrorLoadPromptTemplate = "无法加载提示词模板"

	// 中文 - 文件和目录错误
	chineseMessages.ErrorGetCurrentDir = "获取当前目录失败"
	chineseMessages.ErrorPromptFileNotExist = "自定义提示词文件不存在"
	chineseMessages.ErrorReadPromptFile = "读取自定义提示词文件失败"
	chineseMessages.ErrorPromptFileEmpty = "自定义提示词文件为空"
	chineseMessages.ErrorGetAbsolutePath = "无法获取绝对路径"
	chineseMessages.ErrorWalkDirectory = "遍历目录失败"

	// 中文 - Git相关错误
	chineseMessages.ErrorGitLogFailed = "执行git log失败"
	chineseMessages.ErrorGetCommitDetails = "获取提交详情失败"
	chineseMessages.ErrorParseCommitDetails = "解析提交详情失败"
	chineseMessages.ErrorParseDateFailed = "解析日期失败"
	chineseMessages.ErrorGetChangedFiles = "获取变更文件列表失败"
	chineseMessages.ErrorGetGitUsername = "获取Git用户名失败"
}
