# 好好说话 haohao-shuohua

一套中文表达规范，做成了给 AI 助手用的 skill。凡是要交给人看的中文都按它来：改写和润色、替你起草，也包括 AI 自己给出的回答、总结和汇报。给 Claude Code、Codex、Cursor 以及任何能读 Markdown 的 AI 助手用。

一句话：保住事实，不添原文没有的信息，去掉 AI 味，写回中文本来的样子，不造词装深刻。

它做三件事：

- **改稿**：把一段 AI 味重的中文改成能直接发出去的版本。先看全篇、从含义出发改，不逐句翻译；没毛病的句子不动。
- **从零写**：替你起草消息、邮件、通知、文档、文章时，一开始就按这套标准写。
- **只标问题**：不改原文，只指出哪里有腔调、该怎么改。

外加一个 Go 写的检查器 `hhsh`：四项零容忍扫描（中文里的半角标点、元话语、机械编号和量词标题、“不是 X 而是 Y”超过两处），腔调候选和句长节奏，改稿前后的数字、日期、代码核对。

## 规则的骨架

| | 好好说话 |
|---|---|
| 底线 | 保真 > 场合 > 力度 > 风格。数字、术语原词、因果方向、承诺、条件、否定、“计划/预计/力争”这类强弱，一条都不丢、不改、不加 |
| 改法 | 先读懂全篇，从含义出发整体改，最后才落到句子；每处改动说得出治的是什么毛病，说不出就放回原样 |
| 腔调 | 结构（铺垫、升华、车轱辘话、量词标题和排比）、句子（“不是 X 而是 Y”、空动词、憋气句）、词（黑话、夸饰、元话语、自造词、装饰性比喻、词不配位）、标点、姿态 |
| 场合 | 官方对外、敬语、要气势、私人随手、技术议论五档，口语化的程度跟着文体走 |
| 力度 | 默认标准档；说“大改”可以动结构，说“别大改”只换词 |
| 文档 | PRD、技术方案单独一节：结构留着，宣传结论照清，真术语照用，从零写用文档的惯例格式 |
| 标题 | 直给，一句陈述加一个判断，20 字以内，不用破折号和翻案句，题眼用读者认识的词 |
| 交付 | 默认只给正文，必要时附最多三行提醒；要修改报告时按模板出 |
| 风格 | 默认关；说“记住我的风格”才存，在底线以内照着写 |
| 评测 | 72 条通用用例、38 条 PRD / 技术方案用例、24 条规则定稿前冻结的新留出题，盲评、多次采样，带置信区间和逐题显著性检验 |

## 安装

### Claude Code

```bash
git clone https://github.com/luckylucky0001/haohao-shuohua ~/.claude/skills/haohao-shuohua
```

装好后直接说“把这段去 AI 味”“帮我改成人话”，会自动触发；也可以点名“用 haohao-shuohua 改一下”。

### Codex

```bash
git clone https://github.com/luckylucky0001/haohao-shuohua ~/.codex/skills/haohao-shuohua
```

或在任务里写明：读取 `~/.codex/skills/haohao-shuohua/SKILL.md`，按规则改写。

### Cursor / 其他工具

把 `SKILL.md` 的内容放进项目规则（如 `.cursor/rules/haohao-shuohua.mdc`），`references/` 按需引用。

### claude.ai

在设置里上传 skill 文件夹（打包成 zip），或把 `SKILL.md` 的正文贴进项目说明。

### ChatGPT、Gemini、Kimi、通义千问等没有 skill 机制的地方

把 [dist/haohao-shuohua-mini.md](dist/haohao-shuohua-mini.md) 的正文粘进自定义指令或项目说明。精简版 16 条，覆盖最常用的规则。

## 用法

```
把这段去掉 AI 味，保留信息和我的口气：……
这是 Codex 写的进度汇报，改成正常中文我要转给组里：……
帮我写一条工作消息，通知下周三晚上数据库迁移：……
这份 PRD 的背景和目标是 AI 起的稿，评审前帮我改实一点，结构别动：……
这封邮件别那么客套，但要正式：……
只标问题，不改写：……
记住我的风格：（贴几段你自己写的东西）
```

用户的要求优先：“只改措辞”“保留结构”“可以大改”“别大改”“别吹”“保留小红书语气”“要能念出来”都会照办。

## 检查器 hhsh

需要 Go 1.22 及以上：

```bash
cd ~/.claude/skills/haohao-shuohua
go run ./cmd/hhsh scan 稿子.md             # 四项零容忍扫描；-skip-quoted 不报引号里的例子，-strict 有半角标点时返回 1
go run ./cmd/hhsh lint 稿子.md             # 腔调候选、排版信号、句长节奏
go run ./cmd/hhsh check 原文.md 改稿.md     # 保护项核对、残留与新增腔调、改动率
go run ./cmd/hhsh rules -cat jargon        # 查看某一类规则
```

也可以编译成二进制：`go build -o ~/bin/hhsh ./cmd/hhsh`。

命中只是候选：同一个词在技术文档里可能是术语，在宣传稿里才是腔调。检查器帮你少漏，不替你判断，也不负责“降低 AI 检测率”。

## 目录

```
SKILL.md                 skill 入口：底线、改法、交付、场合和力度、从零写、文档类
references/
  patterns.md            腔调症状全谱：样子、为什么像 AI、怎么改、什么时候留
  fidelity.md            保真细则：程度词表、情态梯子、否定范围、换壳自查
  formats.md             排版按渠道走：消息、短信、公告、邮件、新闻稿、念稿、文档
  scenes.md              各场景的尺度
  writing.md             从零写：各体裁的结构和示范
  examples.md            31 组改写对照，按场景分组
  workflow.md            成稿流程、力度细则、修改报告模板、自检清单
  titles.md              标题和小标题
lexicon/zh.json          检查器规则表
cmd/hhsh                 检查器
cmd/hheval               评测工具（出任务单、打包盲评、汇总报告）
dist/                    给没有 skill 机制的平台用的精简版
evals/                   冻结用例、判官说明、每轮结果和原始输出
docs/                    使用指南、规则说明、检查器说明、参与改进、迭代记录、三页手册
```

## 文档

- [docs/guide.md](docs/guide.md)：使用指南。装到各平台、常用说法、交回什么、场合和力度怎么指定、个人风格、常见问题。
- [docs/rules.md](docs/rules.md)：规则为什么这样定，每条的证据和代价，试过但没留下的做法。
- [docs/checker.md](docs/checker.md)：检查器 hhsh 的命令、输出怎么读、规则表格式、在 CI 里用。
- [docs/contributing.md](docs/contributing.md)：怎么提新腔调、加用例、跑一轮评测。
- [docs/iterations.md](docs/iterations.md)：每一版改了什么、评测结果、为什么这样改。
- [docs/manual/haohao-shuohua-manual.pdf](docs/manual/haohao-shuohua-manual.pdf)：三页使用手册。

## 评测

同一个模型分别装 skill、不装 skill 去做，Claude Opus 盲评排名。用例有三组：72 条通用题（44 条改稿、14 条本来就自然的对照、14 条从零写），38 条 PRD 和技术方案题，以及 24 条在 v0.7 规则定稿前冻结的新留出题（写题的人看不到 skill）。一共十四轮。

**装和不装**（新留出集 H5，24 题，规则从没见过这批题）：

| 方案 | 保真 通过/小偏差/失败 | 质量分 | 胜率 |
|---|---|---:|---:|
| 好好说话 v0.7 | 19/4/1 | 4.71 | 71% |
| 好好说话 v0.6.3 | 19/5/0 | 4.71 | 65% |
| 不装 skill | 12/12/0 | 4.31 | 15% |

不装 skill 的 24 题里一半丢信息、加料或改了语气强弱；装了以后逐题 20 比 4、21 比 3。通用 72 题上也一样（第九轮 64% 对 23%）。

**当前版本 v0.7.1**（第十四轮通用 72 题、H6 新留出 24 题）：

| 方案 | 通用 保真 | 通用胜率 | 新留出 保真 | 新留出胜率 |
|---|---|---:|---|---:|
| 好好说话 v0.7.1 | 65/7/0 | 51% | 22/2/0 | 52% |
| 好好说话 v0.7 | 68/4/0 | 41% | 22/2/0 | 50% |
| 好好说话 v0.5.1 | 66/6/0 | 58% | | |
| 好好说话 v0.6.3 | | | 21/3/0 | 48% |

几个版本之间差距都在噪声内：v0.7.1 对 v0.5.1 逐题 34 比 38，新留出题上三版打平。扩充参考资料没有拖累效果，对照题（本来就自然的文字）表现最好。同一份输出换一批判官，胜率能差 10 个点以上，所以小差距不要当真。

每一版改了什么、为什么改，包括没成功的尝试，见 [docs/iterations.md](docs/iterations.md)；方法、局限和复现步骤见 [evals/README.md](evals/README.md)。

## 许可

MIT。调研和对照参考了 [说人话 shuorenhua](https://github.com/MrGeDiao/shuorenhua)、[Humanizer-zh](https://github.com/op7418/Humanizer-zh)、[qu-ai-wei](https://github.com/LifelongLazyLearner/qu-ai-wei)、[speak-human-tw](https://github.com/Raymondhou0917/speak-human-tw)、[humanizer](https://github.com/blader/humanizer)、[stop-slop](https://github.com/hardikpandya/stop-slop)，均为 MIT 许可；本项目的规则和文字为重新编写。
