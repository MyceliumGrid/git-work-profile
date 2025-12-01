package interactive

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MyceliumGrid/git-work-profile/internal/i18n"
	"github.com/manifoldco/promptui"
)

// Config 交互式配置
type Config struct {
	AnalysisType   string
	TimeRange      string
	CustomFromDate string // 自定义开始日期
	CustomToDate   string // 自定义结束日期
	RepoMode       string
	RepoPath       string
	OutputFormat   string
	OutputFile     string
	AuthorName     string
}

// RunInteractive 运行交互式配置
func RunInteractive() (*Config, error) {
	config := &Config{}

	// 先选择语言
	if err := selectLanguage(); err != nil {
		return nil, err
	}

	// 获取当前语言的消息
	msg := i18n.T()

	fmt.Println(msg.InteractiveTitle)
	fmt.Println("================================================")
	fmt.Println()

	// 1. 选择分析类型
	analysisPrompt := promptui.Select{
		Label: msg.SelectAnalysisType,
		Items: []string{
			msg.AnalysisProfile,
			msg.AnalysisExperience,
			msg.AnalysisTechStack,
		},
		Size: 3,
	}

	idx, _, err := analysisPrompt.Run()
	if err != nil {
		return nil, err
	}

	analysisTypes := []string{"profile", "experience", "techstack"}
	config.AnalysisType = analysisTypes[idx]

	// 2. 选择时间范围
	timeRangePrompt := promptui.Select{
		Label: msg.SelectTimeRange,
		Items: []string{
			msg.TimeRange3Months,
			msg.TimeRange6Months,
			msg.TimeRange1Year,
			msg.TimeRange2Years,
			msg.TimeRangeCustom,
		},
		Size:      5,
		CursorPos: 1, // 默认选中6个月
	}

	idx, _, err = timeRangePrompt.Run()
	if err != nil {
		return nil, err
	}

	// 处理时间范围选择
	if idx == 4 {
		// 自定义日期范围
		config.TimeRange = "custom"

		// 输入开始日期
		fromDatePrompt := promptui.Prompt{
			Label:    msg.InputFromDate,
			Validate: validateDate,
		}
		config.CustomFromDate, err = fromDatePrompt.Run()
		if err != nil {
			return nil, err
		}

		// 输入结束日期
		toDatePrompt := promptui.Prompt{
			Label:    msg.InputToDate,
			Validate: validateDate,
		}
		config.CustomToDate, err = toDatePrompt.Run()
		if err != nil {
			return nil, err
		}
	} else {
		timeRanges := []string{"3m", "6m", "1y", "2y"}
		config.TimeRange = timeRanges[idx]
	}

	// 3. 选择仓库模式
	repoModePrompt := promptui.Select{
		Label: msg.SelectRepoMode,
		Items: []string{
			msg.RepoModeCurrent,
			msg.RepoModeSingle,
			msg.RepoModeMultiple,
		},
		Size: 3,
	}

	idx, _, err = repoModePrompt.Run()
	if err != nil {
		return nil, err
	}

	switch idx {
	case 0:
		// 当前目录
		config.RepoMode = "current"
		config.RepoPath = "."
		// 验证当前目录是否是Git仓库
		if err := validateGitRepo("."); err != nil {
			msg := i18n.T()
			return nil, fmt.Errorf(msg.ErrorCurrentDirNotGitRepo, err)
		}
	case 1:
		// 单个仓库
		config.RepoMode = "single"
		repoPathPrompt := promptui.Prompt{
			Label:    msg.InputRepoPath,
			Default:  ".",
			Validate: validateGitRepo,
		}
		config.RepoPath, err = repoPathPrompt.Run()
		if err != nil {
			return nil, err
		}
	case 2:
		// 多仓库
		config.RepoMode = "multiple"
		reposPathPrompt := promptui.Prompt{
			Label:    msg.InputReposPath,
			Default:  getDefaultProjectsPath(),
			Validate: validateDirectory,
		}
		config.RepoPath, err = reposPathPrompt.Run()
		if err != nil {
			return nil, err
		}
	}

	// 4. 选择输出格式
	formatPrompt := promptui.Select{
		Label: msg.SelectOutputFormat,
		Items: []string{
			msg.FormatMarkdown,
			msg.FormatJSON,
			msg.FormatText,
		},
		Size:      3,
		CursorPos: 0,
	}

	idx, _, err = formatPrompt.Run()
	if err != nil {
		return nil, err
	}

	formats := []string{"markdown", "json", "text"}
	config.OutputFormat = formats[idx]

	// 5. 输出文件
	defaultFileName := generateDefaultOutputFile(config.AnalysisType, config.OutputFormat)
	outputPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%s)", msg.InputOutputFile, msg.OutputFileHint),
		Default: defaultFileName, // 默认保存到文件
	}

	config.OutputFile, err = outputPrompt.Run()
	if err != nil {
		return nil, err
	}

	// Trim 空格
	config.OutputFile = strings.TrimSpace(config.OutputFile)

	// 如果用户输入了文件名，验证路径
	if config.OutputFile != "" {
		if err := validateOutputPath(config.OutputFile); err != nil {
			return nil, err
		}
	}

	// 6. 作者名称（可选）
	// 获取当前Git用户名作为默认值
	defaultAuthor := getCurrentGitUserName(config.RepoPath)

	authorPrompt := promptui.Prompt{
		Label:   msg.InputAuthor,
		Default: defaultAuthor, // 默认使用当前Git用户名
	}

	config.AuthorName, err = authorPrompt.Run()
	if err != nil {
		return nil, err
	}

	// Trim 空格
	config.AuthorName = strings.TrimSpace(config.AuthorName)

	// 显示配置摘要
	fmt.Println()
	fmt.Println(msg.ConfigSummary)
	fmt.Println("================================================")
	fmt.Printf("%s: %s\n", msg.LabelAnalysisType, config.AnalysisType)

	// 显示时间范围
	if config.TimeRange == "custom" {
		fmt.Printf("%s: %s ~ %s\n", msg.LabelTimeRange, config.CustomFromDate, config.CustomToDate)
	} else {
		fmt.Printf("%s: %s\n", msg.LabelTimeRange, config.TimeRange)
	}

	fmt.Printf("%s: %s\n", msg.LabelRepoPath, config.RepoPath)
	fmt.Printf("%s: %s\n", msg.LabelOutputFormat, config.OutputFormat)

	if config.OutputFile != "" {
		fmt.Printf("%s: %s\n", msg.LabelOutputFile, config.OutputFile)
	} else {
		fmt.Printf("%s: %s\n", msg.LabelOutputFile, msg.LabelTerminalOutput)
	}

	if config.AuthorName != "" {
		fmt.Printf("%s: %s\n", msg.LabelAuthorFilter, config.AuthorName)
	}
	fmt.Println()

	// 确认（默认Y）
	confirmPrompt := promptui.Prompt{
		Label:     msg.ConfirmAnalysis,
		Default:   "Y",
		IsConfirm: false,
	}

	result, err := confirmPrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("%s", msg.Canceled)
	}

	// 检查是否拒绝
	result = strings.ToLower(strings.TrimSpace(result))
	if result == "n" || result == "no" {
		return nil, fmt.Errorf("%s", msg.Canceled)
	}

	return config, nil
}

// getDefaultProjectsPath 获取默认的项目目录路径
func getDefaultProjectsPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}

	// 尝试常见的项目目录
	commonPaths := []string{
		filepath.Join(home, "projects"),
		filepath.Join(home, "work"),
		filepath.Join(home, "code"),
		filepath.Join(home, "dev"),
		filepath.Join(home, "workspace"),
	}

	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return "."
}

// generateDefaultOutputFile 生成默认的输出文件名
func generateDefaultOutputFile(analysisType, format string) string {
	var prefix string
	switch analysisType {
	case "profile":
		prefix = "developer-profile"
	case "experience":
		prefix = "project-experience"
	case "techstack":
		prefix = "techstack-analysis"
	default:
		prefix = "analysis"
	}

	var ext string
	switch format {
	case "markdown":
		ext = "md"
	case "json":
		ext = "json"
	default:
		ext = "txt"
	}

	return fmt.Sprintf("%s.%s", prefix, ext)
}

// PromptForAPIKey 提示输入API密钥
func PromptForAPIKey() (string, error) {
	msg := i18n.T()

	fmt.Println()
	fmt.Println(msg.APIKeyNotFound)
	fmt.Println()
	fmt.Println(msg.APIKeyInstructions)
	fmt.Println()

	prompt := promptui.Prompt{
		Label: msg.InputAPIKey,
		Mask:  '*',
		Validate: func(input string) error {
			input = strings.TrimSpace(input)
			if input == "" {
				return fmt.Errorf("%s", msg.APIKeyEmpty)
			}
			if len(input) < 20 {
				return fmt.Errorf("%s", msg.APIKeyInvalid)
			}
			return nil
		},
	}

	apiKey, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(apiKey), nil
}

// PromptYesNo 提示是否继续
func PromptYesNo(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	return err == nil
}

// getCurrentGitUserName 获取当前Git用户名
func getCurrentGitUserName(repoPath string) string {
	// 尝试从指定路径获取
	if repoPath != "" && repoPath != "." {
		cmd := exec.Command("git", "config", "user.name")
		cmd.Dir = repoPath
		if output, err := cmd.Output(); err == nil {
			if name := strings.TrimSpace(string(output)); name != "" {
				return name
			}
		}
	}

	// 尝试从当前目录获取
	cmd := exec.Command("git", "config", "user.name")
	if output, err := cmd.Output(); err == nil {
		if name := strings.TrimSpace(string(output)); name != "" {
			return name
		}
	}

	// 尝试获取全局配置
	cmd = exec.Command("git", "config", "--global", "user.name")
	if output, err := cmd.Output(); err == nil {
		if name := strings.TrimSpace(string(output)); name != "" {
			return name
		}
	}

	return ""
}

// selectLanguage 选择语言
func selectLanguage() error {
	// 检查环境变量
	envLang := os.Getenv(i18n.LangEnvVar)
	if envLang != "" {
		// 环境变量已设置，直接使用
		if envLang == string(i18n.Chinese) {
			i18n.SetLanguage(i18n.Chinese)
		} else {
			i18n.SetLanguage(i18n.English)
		}
		return nil
	}

	// 交互式选择语言
	langPrompt := promptui.Select{
		Label: "Select Language / 选择语言",
		Items: []string{
			"English",
			"中文",
		},
		Size:      2,
		CursorPos: 0, // 默认英文
	}

	idx, _, err := langPrompt.Run()
	if err != nil {
		return err
	}

	if idx == 0 {
		i18n.SetLanguage(i18n.English)
	} else {
		i18n.SetLanguage(i18n.Chinese)
	}

	fmt.Println()
	return nil
}

// validateDate 验证日期格式
func validateDate(input string) error {
	msg := i18n.T()
	input = strings.TrimSpace(input)
	if input == "" {
		return fmt.Errorf("%s", msg.ErrorDateEmpty)
	}

	// 验证格式 YYYY-MM-DD
	if len(input) != 10 {
		return fmt.Errorf("%s", msg.ErrorDateFormatInvalid)
	}

	// 简单验证
	parts := strings.Split(input, "-")
	if len(parts) != 3 {
		return fmt.Errorf("%s", msg.ErrorDateFormatInvalid)
	}

	return nil
}

// validateGitRepo 验证是否是Git仓库
func validateGitRepo(path string) error {
	msg := i18n.T()
	path = strings.TrimSpace(path)
	if path == "" {
		return fmt.Errorf("%s", msg.ErrorPathEmpty)
	}

	// 检查路径是否存在
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(msg.ErrorPathNotExist, path)
		}
		return fmt.Errorf(msg.ErrorPathNotAccessible, err)
	}

	if !info.IsDir() {
		return fmt.Errorf(msg.ErrorPathNotDirectory, path)
	}

	// 检查是否是Git仓库
	gitPath := filepath.Join(path, ".git")
	if _, err := os.Stat(gitPath); err != nil {
		return fmt.Errorf(msg.ErrorNotGitRepo, path)
	}

	return nil
}

// validateDirectory 验证目录是否存在
func validateDirectory(path string) error {
	msg := i18n.T()
	path = strings.TrimSpace(path)
	if path == "" {
		return fmt.Errorf("%s", msg.ErrorPathEmpty)
	}

	// 检查路径是否存在
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(msg.ErrorPathNotExist, path)
		}
		return fmt.Errorf(msg.ErrorPathNotAccessible, err)
	}

	if !info.IsDir() {
		return fmt.Errorf(msg.ErrorPathNotDirectory, path)
	}

	return nil
}

// validateOutputPath 验证输出文件路径
func validateOutputPath(path string) error {
	msg := i18n.T()
	path = strings.TrimSpace(path)
	if path == "" {
		return nil // 空路径表示输出到终端
	}

	// 检查目录是否存在
	dir := filepath.Dir(path)
	if dir != "." && dir != "" {
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf(msg.ErrorDirectoryNotExist, dir)
			}
			return fmt.Errorf(msg.ErrorDirectoryNotAccessible, err)
		}
	}

	return nil
}
