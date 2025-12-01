package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/MyceliumGrid/git-work-profile/internal/ai"
	"github.com/MyceliumGrid/git-work-profile/internal/git"
	"github.com/MyceliumGrid/git-work-profile/internal/i18n"
	"github.com/MyceliumGrid/git-work-profile/internal/interactive"
	"github.com/MyceliumGrid/git-work-profile/internal/report"
	"github.com/spf13/cobra"
)

// 版本信息，由 GoReleaser 在构建时注入
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	// 命令行参数
	fromDate     string
	toDate       string
	outputFormat string
	outputFile   string
	repoPath     string // Git仓库路径
	reposPath    string // 仓库目录路径，分析该目录下的所有Git仓库
	modelName    string // Gemini模型名称
	authorName   string // Git作者名称
	timeRange    string // 时间范围类型：3m(3个月)、6m(6个月)、1y(1年)、2y(2年)
	analysisType string // 分析类型：profile(开发者画像)、experience(项目经验)、techstack(技术栈)
)

// rootCmd 表示根命令
var rootCmd = &cobra.Command{
	Use: "git-work-profile",
	Run: func(_ *cobra.Command, _ []string) {
		// 如果没有指定任何参数，启动交互式模式
		if !hasAnyFlags() {
			runInteractiveMode()
		} else {
			// 执行生成报告的操作
			generateReport()
		}
	},
}

// updateCommandDescriptions 更新命令描述为当前语言
func updateCommandDescriptions() {
	msg := i18n.T()
	rootCmd.Short = msg.CmdShortDesc
	rootCmd.Long = msg.CmdLongDesc
	versionCmd.Short = msg.CmdVersionShort
}

// 版本子命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(_ *cobra.Command, _ []string) {
		msg := i18n.T()
		fmt.Printf(msg.VersionInfo+"\n", version)
		fmt.Printf(msg.CommitHash+"\n", commit)
		fmt.Printf(msg.BuildDate+"\n", date)
	},
}

func init() {
	// 从环境变量加载语言设置
	lang := i18n.LoadLanguageFromEnv()
	i18n.SetLanguage(lang)

	// 更新命令描述
	updateCommandDescriptions()

	// 添加版本子命令
	rootCmd.AddCommand(versionCmd)

	// 获取多语言消息
	msg := i18n.T()

	// 添加命令行参数
	rootCmd.PersistentFlags().StringVar(&analysisType, "analysis", "profile", msg.FlagAnalysis)
	rootCmd.PersistentFlags().StringVar(&fromDate, "from", "", msg.FlagFrom)
	rootCmd.PersistentFlags().StringVar(&toDate, "to", "", msg.FlagTo)
	rootCmd.PersistentFlags().StringVar(&timeRange, "range", "6m", msg.FlagRange)
	rootCmd.PersistentFlags().StringVar(&outputFormat, "format", "markdown", msg.FlagFormat)
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", msg.FlagOutput)
	rootCmd.PersistentFlags().StringVar(&repoPath, "repo", "", msg.FlagRepo)
	rootCmd.PersistentFlags().StringVar(&reposPath, "repos", "", msg.FlagRepos)
	rootCmd.PersistentFlags().StringVar(&modelName, "model", "", msg.FlagModel)
	rootCmd.PersistentFlags().StringVar(&authorName, "author", "", msg.FlagAuthor)
}

func main() {
	// 执行根命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// calculateTimeRange 根据时间范围类型计算开始和结束日期
func calculateTimeRange(rangeType string) (time.Time, time.Time) {
	now := time.Now()
	var from, to time.Time

	switch rangeType {
	case "3m":
		// 最近3个月
		from = now.AddDate(0, -3, 0)
		to = now
	case "6m":
		// 最近6个月
		from = now.AddDate(0, -6, 0)
		to = now
	case "1y":
		// 最近1年
		from = now.AddDate(-1, 0, 0)
		to = now
	case "2y":
		// 最近2年
		from = now.AddDate(-2, 0, 0)
		to = now
	default:
		// 默认使用最近6个月
		from = now.AddDate(0, -6, 0)
		to = now
	}

	return from, to
}

// generateReport 生成分析报告（支持开发者画像、项目经验、技术栈等类型）
func generateReport() {
	msg := i18n.T()

	// 检查环境变量
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println(msg.ErrorAPIKeyNotSet)
		os.Exit(1)
	}

	// 创建Gemini客户端
	geminiClient, err := ai.NewGeminiClientWithModel(modelName)
	if err != nil {
		fmt.Printf(msg.ErrorCreateClient+"\n", err)
		os.Exit(1)
	}

	// 判断使用何种时间范围
	var from, to time.Time
	var err1, err2 error

	switch {
	case fromDate != "" && toDate != "":
		// 使用自定义时间范围
		from, err1 = time.Parse("2006-01-02", fromDate)
		to, err2 = time.Parse("2006-01-02", toDate)
		if err1 != nil || err2 != nil {
			fmt.Println(msg.ErrorDateFormat)
			os.Exit(1)
		}
		to = to.Add(24*time.Hour - time.Second)
		fmt.Printf(msg.InfoCustomTimeRange+"\n", fromDate, toDate)
	default:
		// 使用预定义的时间范围
		from, to = calculateTimeRange(timeRange)
		fmt.Printf(msg.InfoTimeRange+"\n", timeRange, from.Format("2006-01-02"), to.Format("2006-01-02"))
	}

	// 判断使用何种分析模式：单仓库还是多仓库
	var repoPaths []string
	var discoveryErr error

	switch {
	case reposPath != "":
		// 多仓库模式：发现指定目录下的所有Git仓库
		repoPaths, discoveryErr = git.DiscoverGitRepos(reposPath)
		if discoveryErr != nil {
			fmt.Printf(msg.ErrorDiscoverRepos+"\n", discoveryErr)
			return
		}

		if len(repoPaths) == 0 {
			fmt.Printf(msg.ErrorNoReposFound+"\n", reposPath)
			return
		}
	case repoPath != "":
		// 单仓库模式：使用指定的仓库路径
		repoPaths = []string{repoPath}
	default:
		// 默认模式：使用当前目录
		repoPaths = []string{"."}
	}

	// 收集所有仓库的提交记录
	var allCommits []git.CommitInfo
	repoCommitCounts := make(map[string]int)

	fmt.Printf(msg.InfoProcessingRepos+"\n", len(repoPaths))

	for _, currentRepoPath := range repoPaths {
		fmt.Printf(msg.InfoAnalyzingRepo+"\n", currentRepoPath)

		// 创建Git选项
		gitOpts := git.NewGitOptions(currentRepoPath)

		// 如果命令行指定了作者名称，覆盖自动检测的用户名
		if authorName != "" {
			gitOpts.Author = authorName
		}

		// 获取提交记录
		commits, commitErr := git.GetCommitsBetween(from, to, gitOpts)
		if commitErr != nil {
			fmt.Printf(msg.ErrorRepoWarning+"\n", currentRepoPath, commitErr)
			continue
		}

		// 记录每个仓库的提交数量
		repoCommitCounts[currentRepoPath] = len(commits)

		// 为每个提交添加仓库信息
		for i := range commits {
			commits[i].RepoPath = currentRepoPath
		}

		// 合并到总的提交列表
		allCommits = append(allCommits, commits...)

		fmt.Printf(msg.InfoFoundCommits+"\n", len(commits))
	}

	// 显示汇总统计信息
	fmt.Println(msg.InfoCommitStats)
	totalCommits := 0
	for repoPath, count := range repoCommitCounts {
		// 显示相对路径，更清晰
		displayPath := repoPath
		if reposPath != "" {
			if rel, err := filepath.Rel(reposPath, repoPath); err == nil {
				displayPath = rel
			}
		}
		fmt.Printf("  %s: %d\n", displayPath, count)
		totalCommits += count
	}
	fmt.Printf(msg.InfoTotalCommits, totalCommits)

	if len(allCommits) == 0 {
		fmt.Printf(msg.ErrorNoCommitsFound+"\n", from.Format("2006-01-02"), to.Format("2006-01-02"))
		return
	}

	// 显示作者信息
	if authorName != "" {
		fmt.Printf(msg.InfoFilterAuthor+"\n", authorName)
	} else {
		fmt.Println(msg.InfoAllAuthors)
	}

	fmt.Println(msg.InfoAIAnalyzing)

	// 重用之前创建的客户端变量
	if geminiClient == nil {
		geminiClient, err = ai.NewGeminiClientWithModel(modelName)
		if err != nil {
			fmt.Printf(msg.ErrorCreateClient+"\n", err)
			return
		}
	}
	defer geminiClient.Close()

	// 根据分析类型确定使用哪种提示词
	aiPromptType := ai.GetPromptTypeFromString(analysisType)

	switch aiPromptType {
	case ai.DeveloperProfilePrompt:
		fmt.Println(msg.LabelAnalysisTypeProfile)
	case ai.ProjectExperiencePrompt:
		fmt.Println(msg.LabelAnalysisTypeExperience)
	case ai.TechStackPrompt:
		fmt.Println(msg.LabelAnalysisTypeTechStack)
	default:
		fmt.Println(msg.LabelAnalysisTypeDefault)
	}

	// 使用AI生成分析报告
	analysisResult, err := geminiClient.SummarizeCommitsWithPrompt(allCommits, aiPromptType)
	if err != nil {
		fmt.Printf(msg.ErrorAIAnalysisFailed+"\n", err)
		return
	}

	// 决定输出目标
	var output io.Writer = os.Stdout
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, msg.ErrorCreateOutputFile+"\n", err)
			return
		}
		defer file.Close()
		output = file
	}

	// 创建报告生成器
	reportFormat := report.Format(outputFormat)
	reportGenerator := report.NewGenerator(reportFormat, output)

	// 生成并输出报告
	err = reportGenerator.GenerateProfileReport(analysisResult, allCommits, from, to, analysisType)
	if err != nil {
		fmt.Printf(msg.ErrorOutputFailed+"\n", err)
		return
	}

	fmt.Println(msg.InfoAnalysisComplete)
	if outputFile != "" {
		fmt.Printf(msg.InfoReportSaved+"\n", outputFile)
	}
}

// hasAnyFlags 检查是否指定了任何参数
func hasAnyFlags() bool {
	return fromDate != "" ||
		toDate != "" ||
		timeRange != "6m" ||
		analysisType != "profile" ||
		outputFormat != "markdown" ||
		outputFile != "" ||
		repoPath != "" ||
		reposPath != "" ||
		authorName != "" ||
		modelName != ""
}

// runInteractiveMode 运行交互式模式
func runInteractiveMode() {
	fmt.Println()

	// 检查API密钥
	msg := i18n.T()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		var err error
		apiKey, err = interactive.PromptForAPIKey()
		if err != nil {
			fmt.Printf("%s: %v\n", msg.ErrorCancelled, err)
			os.Exit(1)
		}
		// 设置环境变量供后续使用
		os.Setenv("GEMINI_API_KEY", apiKey)

		// 询问是否保存
		if interactive.PromptYesNo(msg.SaveAPIKey) {
			fmt.Println()
			fmt.Printf(msg.SaveAPIKeyHint, apiKey)
			fmt.Println()
		}
	}

	// 运行交互式配置
	config, err := interactive.RunInteractive()
	if err != nil {
		msg := i18n.T()
		// 检查是否是取消操作
		errMsg := err.Error()
		if errMsg == msg.Canceled || errMsg == "已取消" || errMsg == "Canceled" {
			fmt.Println(msg.ErrorCancelled)
			os.Exit(0)
		}
		fmt.Printf("%s: %v\n", msg.ErrorCancelled, err)
		os.Exit(1)
	}

	// 应用配置到全局变量
	analysisType = config.AnalysisType
	outputFormat = config.OutputFormat
	outputFile = config.OutputFile
	authorName = config.AuthorName

	// 处理时间范围
	if config.TimeRange == "custom" {
		fromDate = config.CustomFromDate
		toDate = config.CustomToDate
		timeRange = "" // 清空timeRange，使用fromDate和toDate
	} else {
		timeRange = config.TimeRange
	}

	switch config.RepoMode {
	case "current":
		repoPath = "."
	case "single":
		repoPath = config.RepoPath
	case "multiple":
		reposPath = config.RepoPath
	}

	// 执行分析
	fmt.Println()
	fmt.Println(msg.AnalysisStarting)
	fmt.Println()
	generateReport()
}
