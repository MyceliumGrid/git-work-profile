package report

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/kway-teow/git-work-profile/internal/git"
	"github.com/kway-teow/git-work-profile/internal/i18n"
)

// Format 表示报告输出格式
type Format string

const (
	// FormatText 纯文本格式
	FormatText Format = "text"
	// FormatMarkdown Markdown格式
	FormatMarkdown Format = "markdown"
)

// Generator 报告生成器
type Generator struct {
	Format Format
	Output io.Writer // 输出目标，可以是文件或标准输出
}

// NewGenerator 创建一个新的报告生成器
func NewGenerator(format Format, output io.Writer) *Generator {
	// 如果没有指定输出，默认使用标准输出
	if output == nil {
		output = os.Stdout
	}
	return &Generator{Format: format, Output: output}
}

// GenerateProfileReport 生成开发者画像报告
func (g *Generator) GenerateProfileReport(analysis string, commits []git.CommitInfo, fromDate, toDate time.Time, analysisType string) error {
	// 根据格式生成报告
	switch g.Format {
	case FormatMarkdown:
		return g.generateMarkdownProfileReport(analysis, commits, fromDate, toDate, analysisType)
	case "json":
		return g.generateJSONProfileReport(analysis, commits, fromDate, toDate, analysisType)
	default: // 默认使用文本格式
		return g.generateTextProfileReport(analysis, commits, fromDate, toDate, analysisType)
	}
}

// generateTextProfileReport 生成纯文本格式的开发者画像报告
func (g *Generator) generateTextProfileReport(analysis string, commits []git.CommitInfo, fromDate, toDate time.Time, analysisType string) error {
	msg := i18n.T()
	reportTitle := g.getAnalysisTitle(analysisType)

	fmt.Fprintf(g.Output, "%s\n", reportTitle)
	fmt.Fprintf(g.Output, msg.ReportTimeRange+": %s %s %s\n", fromDate.Format("2006-01-02"), msg.ReportTo, toDate.Format("2006-01-02"))
	fmt.Fprintln(g.Output, "==================================")
	fmt.Fprintln(g.Output)

	// 统计信息
	stats := g.calculateStats(commits)
	fmt.Fprintf(g.Output, "## %s\n", msg.ReportDataStats)
	fmt.Fprintf(g.Output, "- %s: %d\n", msg.ReportTotalCommits, stats["total_commits"])
	fmt.Fprintf(g.Output, "- %s: %d %s\n", msg.ReportTotalRepos, stats["total_repos"], msg.ReportRepoUnit)
	fmt.Fprintf(g.Output, "- %s: %d %s\n", msg.ReportTotalFiles, stats["total_files"], msg.ReportFileUnit)
	fmt.Fprintln(g.Output)

	// AI分析结果
	fmt.Fprintf(g.Output, "## %s\n", msg.ReportAIAnalysis)
	fmt.Fprintln(g.Output, analysis)
	fmt.Fprintln(g.Output)

	return nil
}

// generateMarkdownProfileReport 生成Markdown格式的开发者画像报告
func (g *Generator) generateMarkdownProfileReport(analysis string, commits []git.CommitInfo, fromDate, toDate time.Time, analysisType string) error {
	msg := i18n.T()
	reportTitle := g.getAnalysisTitle(analysisType)

	fmt.Fprintf(g.Output, "# %s\n\n", reportTitle)
	fmt.Fprintf(g.Output, "**%s**: %s %s %s\n\n", msg.ReportTimeRange, fromDate.Format("2006-01-02"), msg.ReportTo, toDate.Format("2006-01-02"))
	fmt.Fprintf(g.Output, "**%s**: %s\n\n", msg.ReportGeneratedAt, time.Now().Format("2006-01-02 15:04:05"))

	// 统计信息
	stats := g.calculateStats(commits)
	fmt.Fprintf(g.Output, "## 📊 %s\n\n", msg.ReportDataStats)
	fmt.Fprintf(g.Output, "- **%s**: %d\n", msg.ReportTotalCommits, stats["total_commits"])
	fmt.Fprintf(g.Output, "- **%s**: %d %s\n", msg.ReportTotalRepos, stats["total_repos"], msg.ReportRepoUnit)
	fmt.Fprintf(g.Output, "- **%s**: %d %s\n", msg.ReportTotalFiles, stats["total_files"], msg.ReportFileUnit)

	// 文件类型分布
	if fileTypes, ok := stats["file_types"].(map[string]int); ok && len(fileTypes) > 0 {
		fmt.Fprintf(g.Output, "- **%s**:\n", msg.ReportFileTypeDistribution)
		for ext, count := range fileTypes {
			fmt.Fprintf(g.Output, "  - `%s`: %d %s\n", ext, count, msg.ReportFileUnit)
		}
	}
	fmt.Fprintln(g.Output)

	// AI分析结果
	fmt.Fprintf(g.Output, "## 🤖 %s\n\n", msg.ReportAIAnalysis)
	fmt.Fprintln(g.Output, analysis)
	fmt.Fprintln(g.Output)

	// 页脚
	fmt.Fprintln(g.Output, "---")
	fmt.Fprintf(g.Output, "*%s*\n", msg.ReportFooter)

	return nil
}

// generateJSONProfileReport 生成JSON格式的开发者画像报告
func (g *Generator) generateJSONProfileReport(analysis string, commits []git.CommitInfo, fromDate, toDate time.Time, analysisType string) error {
	stats := g.calculateStats(commits)

	result := map[string]interface{}{
		"analysis_type": analysisType,
		"time_range": map[string]string{
			"from": fromDate.Format("2006-01-02"),
			"to":   toDate.Format("2006-01-02"),
		},
		"statistics":   stats,
		"ai_analysis":  analysis,
		"generated_at": time.Now().Format(time.RFC3339),
	}

	encoder := json.NewEncoder(g.Output)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

// getAnalysisTitle 根据分析类型获取标题
func (g *Generator) getAnalysisTitle(analysisType string) string {
	msg := i18n.T()
	switch analysisType {
	case "profile":
		return msg.ReportTitleProfile
	case "experience":
		return msg.ReportTitleExperience
	case "techstack":
		return msg.ReportTitleTechStack
	default:
		return msg.ReportTitleDefault
	}
}

// calculateStats 计算统计数据
func (g *Generator) calculateStats(commits []git.CommitInfo) map[string]interface{} {
	repoSet := make(map[string]bool)
	filesSet := make(map[string]bool)
	fileTypes := make(map[string]int)

	for _, commit := range commits {
		if commit.RepoPath != "" {
			repoSet[commit.RepoPath] = true
		}
		for _, file := range commit.ChangedFiles {
			filesSet[file] = true
			ext := getFileExtension(file)
			if ext != "" {
				fileTypes[ext]++
			}
		}
	}

	return map[string]interface{}{
		"total_commits": len(commits),
		"total_repos":   len(repoSet),
		"total_files":   len(filesSet),
		"file_types":    fileTypes,
	}
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
