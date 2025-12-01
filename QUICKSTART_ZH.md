# 快速开始指南

[English](QUICKSTART.md) | 中文

## 5分钟上手

### 1. 安装

```bash
# 使用 go install 安装
go install github.com/kway-teow/git-work-profile/cmd/git-work-profile@latest

# 或者克隆仓库编译
git clone https://github.com/kway-teow/git-work-profile.git
cd git-work-profile
go build -o git-work-profile ./cmd/git-work-profile/
```

### 2. 语言设置（可选）

```bash
# 默认使用英文界面
git-work-profile

# 使用中文界面
export GIT_PROFILE_LANG=zh
git-work-profile

# 或者临时使用中文
GIT_PROFILE_LANG=zh git-work-profile
```

### 3. 第一次运行（交互式模式）

```bash
# 直接运行，启动交互式配置
git-work-profile
```

交互式界面会：
- 提示输入API密钥（如果未设置）
- 引导你选择分析类型
- 选择时间范围和仓库
- 配置输出格式

### 4. 或者使用命令行模式

```bash
# 先设置API密钥
export GEMINI_API_KEY="your-api-key-here"

# 在任意Git仓库目录下运行
cd ~/your-project
git-work-profile --analysis profile --output my-profile.md
```

### 5. 查看结果

报告会保存到指定文件或输出到终端。

## 常用命令

### 分析个人所有项目

```bash
# 假设你的项目都在 ~/projects 目录下
git-work-profile --repos ~/projects --output full-profile.md
```

### 生成简历用的项目经验

```bash
git-work-profile --repos ~/work --range 2y --analysis experience --output resume.md
```

### 查看技术栈

```bash
git-work-profile --analysis techstack
```

### 生成JSON数据

```bash
git-work-profile --format json --output profile.json
```

## 下一步

- 查看 [README.md](README.md) 了解完整功能
- 查看 [EXAMPLES.md](EXAMPLES.md) 了解更多使用场景
- 根据输出调整分析时间范围和类型

## 故障排除

### 问题：提示 "GEMINI_API_KEY环境变量未设置"

**解决方案**：
```bash
export GEMINI_API_KEY="your-api-key"
```

### 问题：没有找到提交记录

**解决方案**：
- 检查是否在Git仓库目录下
- 尝试增加时间范围：`--range 1y`
- 检查是否指定了错误的作者：`--author "Your Name"`

### 问题：分析时间太长

**解决方案**：
- 减少时间范围：`--range 3m`
- 分析单个仓库而不是多个：`--repo` 而不是 `--repos`

### 问题：输出内容不符合预期

**解决方案**：
- 尝试不同的分析类型：`--analysis profile/experience/techstack`
- 增加提交记录数量（扩大时间范围）
- 确保提交信息清晰明确

## 技巧

1. **永久设置API密钥**：将 `export GEMINI_API_KEY="..."` 添加到 `~/.zshrc` 或 `~/.bashrc`
2. **创建别名**：`alias profile='git-work-profile --repos ~/projects'`
3. **定期更新**：每月运行一次，跟踪技术成长
4. **多格式输出**：同时生成Markdown和JSON，方便不同用途
