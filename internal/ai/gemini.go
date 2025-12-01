package ai

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/MyceliumGrid/git-work-profile/internal/git"
	"github.com/MyceliumGrid/git-work-profile/internal/i18n"
	"google.golang.org/api/option"
)

// 默认模型名称
const DefaultModelName = "gemini-2.5-pro"

// GeminiClient 是Gemini AI API的客户端
type GeminiClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewGeminiClient 创建一个新的Gemini客户端
func NewGeminiClient() (*GeminiClient, error) {
	return NewGeminiClientWithModel(DefaultModelName) // 默认使用gemini-2.5-pro模型
}

// NewGeminiClientWithModel 使用指定模型创建一个新的Gemini客户端
func NewGeminiClientWithModel(modelName string) (*GeminiClient, error) {
	// 如果没有指定模型名称，使用默认模型
	if modelName == "" {
		modelName = DefaultModelName
	}
	// 从环境变量获取API密钥
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		msg := i18n.T()
		return nil, fmt.Errorf("%s", msg.ErrorAPIKeyNotSet)
	}

	// 创建Gemini客户端
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		msg := i18n.T()
		return nil, fmt.Errorf("%s: %w", msg.ErrorGeminiClientFailed, err)
	}

	// 使用指定的模型
	model := client.GenerativeModel(modelName)

	return &GeminiClient{
		client: client,
		model:  model,
	}, nil
}

// SummarizeCommits 使用AI总结提交记录
func (g *GeminiClient) SummarizeCommits(commits []git.CommitInfo) (string, error) {
	return g.SummarizeCommitsWithPrompt(commits, DeveloperProfilePrompt)
}

// SummarizeCommitsWithPrompt 使用指定的提示词类型总结提交记录
func (g *GeminiClient) SummarizeCommitsWithPrompt(commits []git.CommitInfo, promptType PromptType) (string, error) {
	if len(commits) == 0 {
		return "没有找到提交记录。", nil
	}

	// 获取时间范围
	var earliestDate, latestDate time.Time
	if len(commits) > 0 {
		earliestDate = commits[len(commits)-1].Date
		latestDate = commits[0].Date

		// 遍历所有提交，找出最早和最晚的日期
		for _, commit := range commits {
			if commit.Date.Before(earliestDate) {
				earliestDate = commit.Date
			}
			if commit.Date.After(latestDate) {
				latestDate = commit.Date
			}
		}
	}

	// 构建提示词
	prompt := buildPromptWithTemplate(commits, earliestDate, latestDate, promptType)

	// 调用Gemini API
	ctx := context.Background()
	resp, err := g.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		msg := i18n.T()
		return "", fmt.Errorf("%s: %w", msg.ErrorGeminiAPIFailed, err)
	}

	// 提取回复
	var result strings.Builder
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			result.WriteString(fmt.Sprintf("%v", part))
		}
	}

	return result.String(), nil
}

// buildPromptWithTemplate 使用指定的提示词模板构建提示词
func buildPromptWithTemplate(commits []git.CommitInfo, fromDate, toDate time.Time, promptType PromptType) string {
	// 获取提示词模板
	template, err := loadPromptTemplate(promptType)
	if err != nil {
		// 如果加载模板失败，使用默认的提示词
		msg := i18n.T()
		fmt.Printf(msg.WarningPromptLoadFailed+"\n", err)
		template = defaultPromptTemplate
	}

	// 构建提交记录字符串
	var commitMessages strings.Builder

	// 统计数据
	totalCommits := len(commits)
	repoSet := make(map[string]bool)
	fileTypeMap := make(map[string]int)

	for i, commit := range commits {
		// 添加提交记录
		fmt.Fprintf(&commitMessages, "提交 %d:\n", i+1)
		fmt.Fprintf(&commitMessages, "- 哈希值: %s\n", commit.Hash[:8])
		fmt.Fprintf(&commitMessages, "- 作者: %s\n", commit.Author)
		fmt.Fprintf(&commitMessages, "- 日期: %s\n", commit.Date.Format("2006-01-02 15:04:05"))

		// 统计仓库
		if commit.RepoPath != "" {
			repoSet[commit.RepoPath] = true
		}

		// 添加分支信息
		if len(commit.Branches) > 0 {
			fmt.Fprintf(&commitMessages, "- 分支: %s\n", strings.Join(commit.Branches, ", "))
		}

		// 添加提交消息
		fmt.Fprintf(&commitMessages, "- 消息: %s\n", commit.Message)

		// 添加变更文件
		if len(commit.ChangedFiles) > 0 {
			fmt.Fprintf(&commitMessages, "- 变更文件:\n")
			// 最多显示10个文件
			maxFiles := 10
			if len(commit.ChangedFiles) < maxFiles {
				maxFiles = len(commit.ChangedFiles)
			}
			for j := 0; j < maxFiles; j++ {
				file := commit.ChangedFiles[j]
				fmt.Fprintf(&commitMessages, "  * %s\n", file)

				// 统计文件类型
				ext := getFileExtension(file)
				if ext != "" {
					fileTypeMap[ext]++
				}
			}
			if len(commit.ChangedFiles) > maxFiles {
				fmt.Fprintf(&commitMessages, "  * ... 以及其他 %d 个文件\n", len(commit.ChangedFiles)-maxFiles)
			}
		}

		// 添加空行分隔不同提交
		fmt.Fprintf(&commitMessages, "\n")
	}

	// 构建文件类型统计字符串
	var fileTypes strings.Builder
	for ext, count := range fileTypeMap {
		fmt.Fprintf(&fileTypes, "%s(%d) ", ext, count)
	}

	// 替换模板中的变量
	prompt := template
	prompt = strings.ReplaceAll(prompt, "{{.CommitMessages}}", commitMessages.String())
	prompt = strings.ReplaceAll(prompt, "{{.TotalCommits}}", fmt.Sprintf("%d", totalCommits))
	prompt = strings.ReplaceAll(prompt, "{{.TimeRange}}", fmt.Sprintf("%s 至 %s", fromDate.Format("2006-01-02"), toDate.Format("2006-01-02")))
	prompt = strings.ReplaceAll(prompt, "{{.RepoCount}}", fmt.Sprintf("%d", len(repoSet)))
	prompt = strings.ReplaceAll(prompt, "{{.LinesAdded}}", "N/A")
	prompt = strings.ReplaceAll(prompt, "{{.LinesDeleted}}", "N/A")
	prompt = strings.ReplaceAll(prompt, "{{.FileTypes}}", fileTypes.String())

	return prompt
}

// getFileExtension 获取文件扩展名
func getFileExtension(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i:]
		}
		if filename[i] == '/' {
			return ""
		}
	}
	return ""
}

// GenerateReport 根据提交记录和时间范围生成报告
func (g *GeminiClient) GenerateReport(commits []git.CommitInfo, fromDate, toDate time.Time) (string, error) {
	return g.GenerateReportWithPrompt(commits, fromDate, toDate, DeveloperProfilePrompt)
}

// GenerateReportWithPrompt 使用指定的提示词类型生成报告
func (g *GeminiClient) GenerateReportWithPrompt(commits []git.CommitInfo, fromDate, toDate time.Time, promptType PromptType) (string, error) {
	// 这个方法实际上是对SummarizeCommits的封装，提供更明确的接口
	if len(commits) == 0 {
		// 根据时间范围返回不同的消息
		daysDiff := toDate.Sub(fromDate).Hours() / 24
		var periodType string

		switch {
		case daysDiff <= 1:
			periodType = "今日"
		case daysDiff <= 7:
			periodType = "本周"
		case daysDiff <= 31:
			periodType = "本月"
		case daysDiff <= 366:
			periodType = "本年"
		default:
			periodType = "指定时间范围内"
		}

		return fmt.Sprintf("%s没有提交记录。", periodType), nil
	}

	// 构建提示词
	prompt := buildPromptWithTemplate(commits, fromDate, toDate, promptType)

	// 调用Gemini API
	ctx := context.Background()
	resp, err := g.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		msg := i18n.T()
		return "", fmt.Errorf("%s: %w", msg.ErrorGeminiAPIFailed, err)
	}

	// 提取回复
	var result strings.Builder
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			result.WriteString(fmt.Sprintf("%v", part))
		}
	}

	return result.String(), nil
}

// Close 关闭Gemini客户端
func (g *GeminiClient) Close() {
	if g.client != nil {
		g.client.Close()
	}
}

// 默认提示词模板
const defaultPromptTemplate = `你是一位专业的技术人才分析师。请根据以下Git提交记录，生成一份开发者画像报告。

提交记录：
{{.CommitMessages}}

统计数据：
- 总提交数：{{.TotalCommits}}
- 分析时间范围：{{.TimeRange}}
- 涉及仓库数：{{.RepoCount}}

请分析开发者的技术栈、工作风格、专业领域和核心竞争力。`

// loadPromptTemplate 从文件加载提示词模板
func loadPromptTemplate(promptType PromptType) (string, error) {
	// 根据提示词类型确定文件名
	var filename string
	switch promptType {
	case DeveloperProfilePrompt:
		filename = "developer-profile.txt"
	case ProjectExperiencePrompt:
		filename = "project-experience.txt"
	case TechStackPrompt:
		filename = "techstack-analysis.txt"
	default:
		filename = "developer-profile.txt"
	}

	// 尝试从多个可能的位置加载模板
	cwd, err := os.Getwd()
	if err != nil {
		msg := i18n.T()
		return "", fmt.Errorf("%s: %w", msg.ErrorGetCurrentDir, err)
	}

	// 尝试从多个可能的位置加载模板
	paths := []string{
		fmt.Sprintf("%s/prompts/%s", cwd, filename),       // 当前目录下的prompts目录
		fmt.Sprintf("%s/../prompts/%s", cwd, filename),    // 上级目录下的prompts目录
		fmt.Sprintf("%s/../../prompts/%s", cwd, filename), // 上上级目录下的prompts目录
	}

	var content []byte
	var loadErr error

	// 尝试每个路径
	for _, path := range paths {
		content, loadErr = loadPromptTemplateFromPath(path)
		if loadErr == nil {
			// 成功加载模板
			return string(content), nil
		}
	}

	// 所有路径都失败了，返回最后一个错误
	msg := i18n.T()
	return "", fmt.Errorf("%s: %w", msg.ErrorLoadPromptTemplate, loadErr)
}

// loadPromptTemplateFromPath 从指定路径加载提示词模板
func loadPromptTemplateFromPath(path string) ([]byte, error) {
	return os.ReadFile(path)
}
