package profile

import (
	"time"

	"github.com/kway-teow/git-work-profile/internal/git"
)

// DeveloperProfile 开发者画像
type DeveloperProfile struct {
	Author      string     `json:"author"`
	TimeRange   TimeRange  `json:"time_range"`
	Statistics  Statistics `json:"statistics"`
	TechStack   TechStack  `json:"tech_stack"`
	WorkStyle   WorkStyle  `json:"work_style"`
	Expertise   Expertise  `json:"expertise"`
	AIAnalysis  string     `json:"ai_analysis"`
	GeneratedAt time.Time  `json:"generated_at"`
}

// TimeRange 时间范围
type TimeRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// Statistics 统计数据
type Statistics struct {
	TotalCommits   int            `json:"total_commits"`
	TotalRepos     int            `json:"total_repos"`
	LinesAdded     int            `json:"lines_added"`
	LinesDeleted   int            `json:"lines_deleted"`
	FilesChanged   int            `json:"files_changed"`
	FileTypeStats  map[string]int `json:"file_type_stats"`
	RepoStats      map[string]int `json:"repo_stats"`
	CommitsByMonth map[string]int `json:"commits_by_month"`
	CommitsByHour  map[int]int    `json:"commits_by_hour"`
}

// TechStack 技术栈
type TechStack struct {
	Languages  map[string]int `json:"languages"`
	Frameworks []string       `json:"frameworks"`
	Tools      []string       `json:"tools"`
	Platforms  []string       `json:"platforms"`
}

// WorkStyle 工作风格
type WorkStyle struct {
	AvgCommitsPerDay    float64 `json:"avg_commits_per_day"`
	AvgLinesPerCommit   float64 `json:"avg_lines_per_commit"`
	MostActiveHour      int     `json:"most_active_hour"`
	MostActiveDay       string  `json:"most_active_day"`
	CommitMessageLength float64 `json:"commit_message_length"`
}

// Expertise 专业领域
type Expertise struct {
	PrimaryDomain    string   `json:"primary_domain"`
	SecondaryDomains []string `json:"secondary_domains"`
	KeySkills        []string `json:"key_skills"`
}

// AnalyzeProfile 分析开发者画像
func AnalyzeProfile(commits []git.CommitInfo, from, to time.Time, author string) *DeveloperProfile {
	profile := &DeveloperProfile{
		Author: author,
		TimeRange: TimeRange{
			From: from,
			To:   to,
		},
		Statistics:  calculateStatistics(commits),
		TechStack:   analyzeTechStack(commits),
		WorkStyle:   analyzeWorkStyle(commits, from, to),
		Expertise:   analyzeExpertise(commits),
		GeneratedAt: time.Now(),
	}

	return profile
}

// calculateStatistics 计算统计数据
func calculateStatistics(commits []git.CommitInfo) Statistics {
	stats := Statistics{
		TotalCommits:   len(commits),
		FileTypeStats:  make(map[string]int),
		RepoStats:      make(map[string]int),
		CommitsByMonth: make(map[string]int),
		CommitsByHour:  make(map[int]int),
	}

	repoSet := make(map[string]bool)
	filesSet := make(map[string]bool)

	for _, commit := range commits {
		// 仓库统计
		if commit.RepoPath != "" {
			stats.RepoStats[commit.RepoPath]++
			repoSet[commit.RepoPath] = true
		}

		// 文件类型统计
		for _, file := range commit.ChangedFiles {
			filesSet[file] = true
			ext := getFileExtension(file)
			if ext != "" {
				stats.FileTypeStats[ext]++
			}
		}

		// 按月统计
		monthKey := commit.Date.Format("2006-01")
		stats.CommitsByMonth[monthKey]++

		// 按小时统计
		stats.CommitsByHour[commit.Date.Hour()]++
	}

	stats.TotalRepos = len(repoSet)
	stats.FilesChanged = len(filesSet)

	return stats
}

// analyzeTechStack 分析技术栈
func analyzeTechStack(commits []git.CommitInfo) TechStack {
	techStack := TechStack{
		Languages:  make(map[string]int),
		Frameworks: []string{},
		Tools:      []string{},
		Platforms:  []string{},
	}

	// 语言映射
	langMap := map[string]string{
		".go":    "Go",
		".js":    "JavaScript",
		".ts":    "TypeScript",
		".py":    "Python",
		".java":  "Java",
		".rb":    "Ruby",
		".php":   "PHP",
		".c":     "C",
		".cpp":   "C++",
		".cs":    "C#",
		".swift": "Swift",
		".kt":    "Kotlin",
		".rs":    "Rust",
		".scala": "Scala",
		".sh":    "Shell",
		".sql":   "SQL",
		".html":  "HTML",
		".css":   "CSS",
		".vue":   "Vue",
		".jsx":   "React",
		".tsx":   "React",
	}

	for _, commit := range commits {
		for _, file := range commit.ChangedFiles {
			ext := getFileExtension(file)
			if lang, ok := langMap[ext]; ok {
				techStack.Languages[lang]++
			}
		}
	}

	return techStack
}

// analyzeWorkStyle 分析工作风格
func analyzeWorkStyle(commits []git.CommitInfo, from, to time.Time) WorkStyle {
	if len(commits) == 0 {
		return WorkStyle{}
	}

	totalDays := to.Sub(from).Hours() / 24
	if totalDays == 0 {
		totalDays = 1
	}

	totalMessageLength := 0
	hourCounts := make(map[int]int)

	for _, commit := range commits {
		totalMessageLength += len(commit.Message)
		hourCounts[commit.Date.Hour()]++
	}

	// 找出最活跃的小时
	mostActiveHour := 0
	maxCount := 0
	for hour, count := range hourCounts {
		if count > maxCount {
			maxCount = count
			mostActiveHour = hour
		}
	}

	return WorkStyle{
		AvgCommitsPerDay:    float64(len(commits)) / totalDays,
		MostActiveHour:      mostActiveHour,
		CommitMessageLength: float64(totalMessageLength) / float64(len(commits)),
	}
}

// analyzeExpertise 分析专业领域
func analyzeExpertise(commits []git.CommitInfo) Expertise {
	expertise := Expertise{
		KeySkills: []string{},
	}

	// 基于文件类型判断领域
	frontendCount := 0
	backendCount := 0
	devopsCount := 0

	frontendExts := map[string]bool{
		".js": true, ".ts": true, ".jsx": true, ".tsx": true,
		".vue": true, ".html": true, ".css": true, ".scss": true,
	}

	backendExts := map[string]bool{
		".go": true, ".py": true, ".java": true, ".rb": true,
		".php": true, ".cs": true, ".rs": true,
	}

	devopsExts := map[string]bool{
		".yml": true, ".yaml": true, ".sh": true, ".dockerfile": true,
	}

	for _, commit := range commits {
		for _, file := range commit.ChangedFiles {
			ext := getFileExtension(file)
			if frontendExts[ext] {
				frontendCount++
			}
			if backendExts[ext] {
				backendCount++
			}
			if devopsExts[ext] {
				devopsCount++
			}
		}
	}

	// 判断主要领域
	if frontendCount > backendCount && frontendCount > devopsCount {
		expertise.PrimaryDomain = "前端开发"
	} else if backendCount > frontendCount && backendCount > devopsCount {
		expertise.PrimaryDomain = "后端开发"
	} else if frontendCount > 0 && backendCount > 0 {
		expertise.PrimaryDomain = "全栈开发"
	} else if devopsCount > 0 {
		expertise.PrimaryDomain = "DevOps"
	} else {
		expertise.PrimaryDomain = "软件开发"
	}

	return expertise
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
