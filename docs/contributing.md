# 参与改进

这份说明写给想改规则、加用例、跑评测的人。好好说话的每条规则都来自评测里的具体失误，改动也要过评测才算数，下面讲具体怎么做。

## 目录

- [先读什么](#先读什么)
- [提一条新的腔调](#提一条新的腔调)
- [加评测用例](#加评测用例)
- [跑一轮评测](#跑一轮评测)
- [改 SKILL.md 的原则](#改-skillmd-的原则)
- [提交前检查](#提交前检查)

## 先读什么

- [SKILL.md](../SKILL.md) 和 [references/](../references/)：现在的规则和语气。
- [docs/rules.md](rules.md)：每条主要规则从哪来、有什么证据，免得把验证过的东西改掉。
- [docs/iterations.md](iterations.md)：每一版改了什么、为什么改，包括没成功的重写。
- [evals/README.md](../evals/README.md) 和 [evals/judge-prompt.md](../evals/judge-prompt.md)：评测方法、判官标准和已知局限。

## 提一条新的腔调

先确认它还没被收录：在 [references/patterns.md](../references/patterns.md) 里搜一下，再跑 `go run ./cmd/hhsh rules | grep 关键字` 查规则表。没有的话，开 issue 或直接提 PR，附上下面这些材料：

```markdown
腔调：（这个说法本身，或者这类句式的样子）
出现在：（文体和场合，例如周报、商务邮件、技术方案）
真实样本：（至少两处，来自不同的文本，脱敏后贴上下文，注明是哪类工具生成的）
平常说法：（换成最平常的说法后的句子，说明意思一点没少）
该留的情形：（它在哪些行当是术语、在哪些文体是惯例；没有就写“没有”）
建议落在：（patterns.md 的哪一层；要不要进 lexicon/zh.json，权重几）
```

几点要求：

- **样本要真实。** 自己编的句子只能当改法示例，证明不了这是常见的腔调。只出现过一次的说法先不收。
- **过一遍替换试法。** 把句子里最抢眼的词换成最平常的说法，意思没少才算腔调；意思少了，它在干活。
- **写清该留的情形。** 很多词在某个行当是真术语，比如工程师腔里的不少词在技术文档里说得出具体机制。写不出该留的情形，再考虑给高权重。
- **示例另编。** 写进 patterns.md 或规则测试的例句要自己新编，换领域、换说法，不能取自评测用例，也不要和真实样本只差几个字。
- **要进检查器的**，按 [checker.md](checker.md) 的“加一条规则并补测试”一节做，规则和测试放在同一个 PR 里。

如果你发现的是 skill 的失误（腔调本来就在规则里，模型没改掉，或者改过了头），附上评测里的用例 ID 和判官简评，用 `go run ./cmd/hheval notes -judge <判分目录> -cond <方案名>` 能列出某个方案逐题的打分和简评。

## 加评测用例

### 格式

每条用例是一个 Markdown 文件，开头是头信息，后面是原文。解析代码在 `internal/evalcase/evalcase.go`，规则很严：只认下面这些字段，多一个字段就报错；字段值只能写一行；列表项以两个空格加 `- ` 开头。

```markdown
---
id: R45
kind: rewrite
scene: 社区游泳馆给会员发的闭馆通知，原稿由 AI 起草
register: 正式而亲切的通知，不要公文腔，也不要聊天口吻
request: 这条闭馆通知是 AI 写的，帮我改自然点，下午要发给会员。
must_keep:
  - 闭馆时间：12 月 1 日至 12 月 7 日
  - 原因：更换池水循环泵
  - 已购次卡有效期顺延 7 天，月卡不受影响
  - 预计 12 月 8 日恢复开放（“预计”不能改成确定的说法）
targets:
  - 开场的“为了给您带来更加优质、舒适、安心的游泳体验”
  - “以全新的面貌迎接您的到来”
traps:
  - 顺延只针对次卡，不能扩大到所有会员卡
---
尊敬的各位会员：

为了给您带来更加优质、舒适、安心的游泳体验，本馆将于 12 月 1 日至 12 月 7 日暂停开放，全面升级池水循环系统，更换循环泵。……
```

（这条是为说明格式编的，不在用例集里。）

| 字段 | 含义 | 给谁看 |
|---|---|---|
| `id` | 编号，用目录里还没用过的；通用题 R 改稿、N 对照、W 从零写，文档题 P、PN、PW，留出题 H | 必填 |
| `kind` | `rewrite` 改稿、`control` 本来就自然的对照、`write` 从零写 | 必填 |
| `scene` | 场景：谁写给谁、什么文体 | 判官 |
| `register` | 语体期待 | 判官 |
| `request` | 用户的原话，一行 | 改写模型和判官，必填 |
| `must_keep` | 必须保留的信息，逐条写，带上情态和范围 | 判官 |
| `targets` | 应清掉的腔调，只作参考 | 判官 |
| `traps` | 陷阱：是术语的词、统计学含义、容易说满的程度词 | 判官 |
| 正文 | 原文；`rewrite` 和 `control` 必须有，`write` 留空 | 改写模型和判官 |

写完跑 `go run ./cmd/hheval inputs -cases <你的目录>`，能出任务单就说明格式没问题。通用题放 `evals/cases/`，PRD 和技术方案题放 `evals/cases-doc/`，留出题单开目录。

### 冻结

- 用例写完、在任何方案跑这批题之前，单独一笔提交，提交说明以 `evals:` 开头，写明“在任何方案跑分前提交”。git 历史就是冻结的凭证。
- 冻结以后不为分数改用例。发现用例本身有错（比如日期和星期对不上），不改文件，写进 [evals/README.md](../evals/README.md) 的局限一节，并说明判分时怎么处理。
- 留出题由没看过 SKILL.md 和迭代用例的人来写，只用来验证、不用来改规则。参考留出题的判官意见改过规则以后，这批题只能算半个留出，要在 README 里写明。

### 不能泄题

- 写题的人不看 SKILL.md 和 references/，只看已有用例的格式。
- 改写模型只拿 `hheval inputs` 导出的任务单，看不到 `must_keep`、`targets`、`traps`。
- SKILL.md、references/、dist/、规则测试和 docs/ 里的示例都不能取自用例。v0.1 就在这上面出过事：几处示例直接来自用例，那几题的成绩偏高，v0.2 起全部换掉。
- 提交前拿新用例里有辨识度的短语在 skill 文件里搜一遍，确认没撞上：`grep -rn "特征短语" SKILL.md references dist lexicon docs`。

## 跑一轮评测

以下命令都在仓库根目录运行，`<轮次>` 沿用 `round-13` 这样的写法。`-cases` 默认是 `evals/cases`；跑文档题时三条命令都加 `-cases evals/cases-doc`。

**出任务单。** 只导出用户请求和原文，评分要点不给改写模型看。`-kinds` 可以写类型，也可以写用例 ID：

```bash
go run ./cmd/hheval inputs -kinds rewrite,control -out tasks.md
```

**改写。** 每个方案交给一个独立的子代理，加载各自的 skill 后逐题完成；`baseline` 不加载任何 skill，只按用户原话做。正文写到 `evals/runs/<轮次>/<方案>/<用例ID>.md`，正文之外对用户说的话写到同目录的 `<用例ID>.notes.md`，没有就不建。要保证子代理读不到别的方案的输出目录，第九轮就出过这个问题。

**打包盲评。** 同一题的各方案输出被打乱成匿名的 A、B、C……，映射存进 `mapping.json`。一题少于两个方案有输出时跳过，一次最多比 8 个方案：

```bash
go run ./cmd/hheval blind -runs evals/runs/<轮次> -conds haohao-new,haohao-old,baseline -batches 4 -out evals/results/<轮次>
```

**判分。** 判官用和改写不同的模型，按 [judge-prompt.md](../evals/judge-prompt.md) 逐包打分，把结果写成同目录下的 `result-1.json`、`result-2.json`……包里的每条用例都要有结果，字母和包里一致。

**出报告。** 合并判官结果和检查器指标，给出保真、质量分、胜率和 95% 区间、两两比较：

```bash
go run ./cmd/hheval report -runs evals/runs/<轮次> -judge evals/results/<轮次> -order haohao-new,haohao-old,baseline -svg evals/results/<轮次>/chart.svg -out evals/results/<轮次>/report.md
```

`-skip` 可以排除部分用例（例如只看新题），`-labels` 和 `-highlight` 调整图里的名字和高亮。看完总表再用 `hheval notes` 读逐题简评，找失误类型。

**读结果要留余量。** 实测同一个 skill 重新采样一次，胜率能差 8 个百分点；同一份输出重新判分，能差 5 个点左右。5 到 8 个点以内的差别不能当真，要结合逐题记录和两两比较看。能跑两次采样就跑两次。

## 改 SKILL.md 的原则

- **每条规则都要有失误证据。** 新增或改动一条规则，要能指出评测里哪几题、哪个方案出了这个问题，写进 [iterations.md](iterations.md) 这一版的条目。凭感觉加的规则，常把本来没问题的地方改坏。
- **对着失误类型写，不对着单题打补丁。** 规则要写成一类问题的判断方法，不写成某道题的答案。
- **示例不能取自用例。** 规则里的例句换领域、换事情、换说法重新编；写完拿用例里的短语搜一遍。
- **SKILL.md 宁短。** 它每次整篇读进去，规则一多，模型在每句话上的注意力就少。细则、词表、对照例子放进 references/，SKILL.md 只留每次都要用的判断。第十轮的 v0.6.1 在整篇重写的基础上又加了规则，改稿题逐题 16 比 28 输给 v0.5.1，经过记在 iterations.md 里。
- **改完要跑评测，打平或更好才合并。** 至少跑受影响的用例类型，再跑一遍通用 72 题回归；动了文档一节，还要跑文档题。保真不能退（失败不能增加，小偏差不能明显增加），胜率差在噪声以内或更好，才算打平。没打平的改动不合并，但要把尝试和结果记进 iterations.md。
- **几处要一起改。** 规则改了，看 README 的规则骨架、[dist/haohao-shuohua-mini.md](../dist/haohao-shuohua-mini.md) 精简版、references/ 和规则表要不要跟着改，别让它们互相矛盾。发版时 `cmd/hhsh/main.go` 里的 `version` 和 `lexicon/zh.json` 的 `version` 一起改。
- **你写的说明本身也要合规。** 第一句说事，中文正文用全角标点，不用元话语和量词标题，不加总结升华。

## 提交前检查

```bash
gofmt -l .                     # 应该没有输出
go vet ./...
go test ./...
go run ./cmd/hhsh scan -skip-quoted README.md SKILL.md docs/*.md
go run ./cmd/hhsh lint -min 2 -skip-quoted 你改过的文档.md
```

- `go test` 必须全过。改了规则表，`lexicon` 和 `internal/lint` 的测试会检查规则合法、典型腔调能命中、自然文字不误报。
- `scan` 的半角标点必须清到零。元话语、编号和“不是 X 而是 Y”的命中如果出现在故意写坏的“改前”示例或规则说明里，可以留；出现在你自己的说明文字里，就改掉。
- `lint` 的命中逐条过一遍，引号里的例子不算，正文里的腔调要改。
- 改了用例或 skill 文件，再做一次泄题检查。
- 提交说明沿用仓库习惯：发版写 `v0.x.y：改了什么`，冻结用例写 `evals: ……`，文档写 `docs: ……`，修复写 `fix: ……`。
