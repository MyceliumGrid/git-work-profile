package ai

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kway-teow/git-work-profile/internal/i18n"
)

// PromptType 表示不同类型的提示词
type PromptType string

const (
	// DeveloperProfilePrompt 开发者画像分析
	DeveloperProfilePrompt PromptType = "profile"
	// ProjectExperiencePrompt 项目经验总结
	ProjectExperiencePrompt PromptType = "experience"
	// TechStackPrompt 技术栈分析
	TechStackPrompt PromptType = "techstack"
)

// GetPromptTypeFromString 根据字符串返回对应的提示词类型
func GetPromptTypeFromString(promptTypeStr string) PromptType {
	switch promptTypeStr {
	case "profile":
		return DeveloperProfilePrompt
	case "experience":
		return ProjectExperiencePrompt
	case "techstack":
		return TechStackPrompt
	default:
		return DeveloperProfilePrompt
	}
}

// LoadCustomPrompt 加载自定义提示词文件
func LoadCustomPrompt(filePath string) (string, error) {
	msg := i18n.T()

	// 检查文件是否存在
	if !filepath.IsAbs(filePath) {
		// 如果是相对路径，转换为绝对路径
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("%s: %w", msg.ErrorGetCurrentDir, err)
		}
		filePath = filepath.Join(cwd, filePath)
	}

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("%s: %s", msg.ErrorPromptFileNotExist, filePath)
	}

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("%s: %w", msg.ErrorReadPromptFile, err)
	}

	promptContent := strings.TrimSpace(string(content))
	if promptContent == "" {
		return "", fmt.Errorf("%s: %s", msg.ErrorPromptFileEmpty, filePath)
	}

	// 确保提示词末尾有一个换行符，以便后续添加提交记录
	if !strings.HasSuffix(promptContent, "\n") {
		promptContent += "\n"
	}

	return promptContent, nil
}
