# 第七轮：72 题，v0.4.2 与 v0.1 各两次采样、不装 skill

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.4.2 | 72 | 66/6/0 | 4.54 | 4.43 | 4.74 | **4.57** | 0.39 | 61% | 0.00 | 2.7 | 20 | 8 | 7 | 10 |
| haohao-v0.4.2b | 72 | 64/8/0 | 4.51 | 4.43 | 4.76 | **4.57** | 0.43 | 57% | 0.00 | 3.3 | 21 | 7 | 5 | 5 |
| haohao-v0.1 | 72 | 68/4/0 | 4.60 | 4.39 | 4.81 | **4.60** | 0.40 | 60% | 0.00 | 2.5 | 16 | 7 | 8 | 6 |
| haohao-v0.1b | 72 | 62/10/0 | 4.62 | 4.43 | 4.82 | **4.63** | 0.46 | 54% | 0.00 | 3.0 | 22 | 7 | 44 | 7 |
| baseline | 72 | 43/26/3 | 4.33 | 3.86 | 4.57 | **4.25** | 0.82 | 18% | 0.64 | 3.6 | 24 | 12 | 9 | 17 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.4.2 | 61% | 54%–68% |
| haohao-v0.4.2b | 57% | 49%–65% |
| haohao-v0.1 | 60% | 52%–68% |
| haohao-v0.1b | 54% | 47%–61% |
| baseline | 18% | 12%–25% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.4.2 | haohao-v0.4.2b | 38 | 34 | 0.724 |
| haohao-v0.4.2 | haohao-v0.1 | 35 | 37 | 0.906 |
| haohao-v0.4.2 | haohao-v0.1b | 43 | 29 | 0.125 |
| haohao-v0.4.2 | baseline | 60 | 12 | 0.000 |
| haohao-v0.4.2b | haohao-v0.1 | 35 | 37 | 0.906 |
| haohao-v0.4.2b | haohao-v0.1b | 37 | 35 | 0.906 |
| haohao-v0.4.2b | baseline | 57 | 15 | 0.000 |
| haohao-v0.1 | haohao-v0.1b | 39 | 33 | 0.556 |
| haohao-v0.1 | baseline | 60 | 12 | 0.000 |
| haohao-v0.1b | baseline | 59 | 13 | 0.000 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.4.2 | 4.40 / 0.42 | 5.00 / 0.29 | 4.67 / 0.39 | 0% | 74% |
| haohao-v0.4.2b | 4.37 / 0.47 | 4.98 / 0.25 | 4.79 / 0.50 | 0% | 74% |
| haohao-v0.1 | 4.48 / 0.38 | 4.83 / 0.55 | 4.71 / 0.32 | 0% | 72% |
| haohao-v0.1b | 4.52 / 0.43 | 4.90 / 0.52 | 4.67 / 0.50 | 26% | 69% |
| baseline | 4.27 / 0.81 | 4.36 / 0.89 | 4.12 / 0.79 | 4% | 79% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > baseline |
| N02 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > baseline |
| N03 | 对照：原文本来就自然 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.4.2b > baseline |
| N04 | 对照：原文本来就自然 | baseline > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 |
| N05 | 对照：原文本来就自然 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 > baseline |
| N06 | 对照：原文本来就自然 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 > baseline |
| N07 | 对照：原文本来就自然 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.4.2 > baseline > haohao-v0.1b |
| N08 | 对照：原文本来就自然 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.1b > baseline |
| N09 | 对照：原文本来就自然 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > baseline |
| N10 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > baseline |
| N11 | 对照：原文本来就自然 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 > baseline |
| N12 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > baseline |
| N13 | 对照：原文本来就自然 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > baseline > haohao-v0.1b |
| N14 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1 > haohao-v0.1b > baseline |
| R01 | 改稿 | haohao-v0.1 > baseline > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b |
| R02 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > baseline |
| R03 | 改稿 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R04 | 改稿 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > baseline |
| R05 | 改稿 | haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R06 | 改稿 | haohao-v0.1 > haohao-v0.4.2 > baseline > haohao-v0.1b > haohao-v0.4.2b |
| R07 | 改稿 | haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > baseline |
| R08 | 改稿 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.1b > baseline > haohao-v0.4.2 |
| R09 | 改稿 | haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > baseline |
| R10 | 改稿 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R11 | 改稿 | haohao-v0.4.2 > haohao-v0.1b > baseline > haohao-v0.4.2b > haohao-v0.1 |
| R12 | 改稿 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > baseline |
| R13 | 改稿 | haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R14 | 改稿 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline > haohao-v0.1 |
| R15 | 改稿 | haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > baseline |
| R16 | 改稿 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > baseline > haohao-v0.4.2 |
| R17 | 改稿 | haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| R18 | 改稿 | haohao-v0.1b > haohao-v0.4.2 > baseline > haohao-v0.1 > haohao-v0.4.2b |
| R19 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > baseline |
| R20 | 改稿 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.1 > baseline |
| R21 | 改稿 | haohao-v0.1b > haohao-v0.1 > baseline > haohao-v0.4.2b > haohao-v0.4.2 |
| R22 | 改稿 | haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| R23 | 改稿 | haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > baseline |
| R24 | 改稿 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R25 | 改稿 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.4.2 > baseline > haohao-v0.1b |
| R26 | 改稿 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R27 | 改稿 | haohao-v0.1 > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R28 | 改稿 | haohao-v0.1 > baseline > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 |
| R29 | 改稿 | haohao-v0.4.2 > haohao-v0.1b > baseline > haohao-v0.1 > haohao-v0.4.2b |
| R30 | 改稿 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R31 | 改稿 | haohao-v0.1b > baseline > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b |
| R32 | 改稿 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 > baseline > haohao-v0.4.2b |
| R33 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.4.2b > haohao-v0.1 > haohao-v0.1b |
| R34 | 改稿 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R35 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > baseline > haohao-v0.1 |
| R36 | 改稿 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2b > baseline |
| R37 | 改稿 | baseline > haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b |
| R38 | 改稿 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R39 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > baseline |
| R40 | 改稿 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > baseline |
| R41 | 改稿 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.4.2 > baseline > haohao-v0.1 |
| R42 | 改稿 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.1b > baseline > haohao-v0.4.2 |
| R43 | 改稿 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| R44 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > baseline |
| W01 | 从零写作 | haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| W02 | 从零写作 | haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline > haohao-v0.1b |
| W03 | 从零写作 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.1b > baseline |
| W04 | 从零写作 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > baseline |
| W05 | 从零写作 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > baseline |
| W06 | 从零写作 | haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline > haohao-v0.1b |
| W07 | 从零写作 | baseline > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 |
| W08 | 从零写作 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > baseline > haohao-v0.4.2b |
| W09 | 从零写作 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > baseline |
| W10 | 从零写作 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > baseline > haohao-v0.4.2 |
| W11 | 从零写作 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > baseline |
| W12 | 从零写作 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| W13 | 从零写作 | haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1 > baseline |
| W14 | 从零写作 | baseline > haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| N05 | baseline | minor | “老陈说他”的转述口吻丢了；“就我们四个”“腰不行”“给准话”等口语被改成书面说法 |  |
| R01 | haohao-v0.1b | minor | “一线”值班同学 | “为双十一大促做好了准备”语气强于“奠定基础” |
| R01 | haohao-v0.4.2b | minor | 为双十一大促做准备的目的被改写 | “压力验证做完了”“大量无效告警”“不用再”属新增判断 |
| R02 | baseline | fail | 国内首个支持离线多人协作的声明 | 不用再一笔笔手输；已在用的到应用商店更新、欢迎反馈 |
| R03 | baseline | minor |  | 接口变慢或报错、翻日志难拼出经过的场景描述；每一步花了多久都看得见 |
| R04 | baseline | minor | “认知”一层的观点被换掉 | 焦虑来源的新解释（看不到进步又停不下来）；结尾新增具体建议；“忙到晚上人很累”等细节 |
| R05 | baseline | minor | “有理有据、态度诚恳”弱化为“拿数据说话，态度好一点” | “不用直接答应”；主语改为“咱们”（说明里已提示）；房东“两年内不用操心找新租客”的展开 |
| R06 | baseline | minor | “再次感谢您的耐心与理解”；“期待您的回复”；“如有问题随时联系”被收窄为安排不妥时告知 | “每周五下午的时间不便请告知” |
| R08 | haohao-v0.4.2 | minor | 监控“如果需要”的提议语气弱化，读着像待办 |  |
| R10 | baseline | minor | 贡献中“补充 AI 教育应用的理论研究”被换掉 | “提供了……实证依据”的新贡献表述 |
| R11 | baseline | minor |  | 茉莉花茶“配着吃刚好” |
| R12 | baseline | minor |  | “这件事我一直记着”；“我心里清楚，这不是我一个人做出来的” |
| R12 | haohao-v0.1b | minor | “2026”改成“今年”，年会若在次年开会有歧义 |  |
| R12 | haohao-v0.4.2b | minor |  | “我都看在眼里” |
| R13 | baseline | minor |  | 避免再出现同样的情况 |
| R15 | baseline | fail | 刘洋完成日期的“预计”被删，写成确定承诺；“暂不调整”的“暂”丢失 | 待办中新增“法务确认”条目及“负责人会上未指定”的断言 |
| R15 | haohao-v0.1 | minor | “暂不调整”的“暂”丢失 |  |
| R15 | haohao-v0.4.2 | minor | “暂不调整”的“暂”丢失 |  |
| R15 | haohao-v0.4.2b | minor | “暂不调整”的“暂”丢失 |  |
| R16 | haohao-v0.4.2 | minor |  |  |
| R17 | baseline | fail | 目标三原口径“退单比例降低一半”；“拍一下即可完成录入”；“一期”原词 | 目标三改为“被退回报销单数量较第三季度减少一半”；背景新增“单张平均耗时约 90 秒”现状陈述 |
| R17 | haohao-v0.1b | minor |  | “报销流程里这一环问题最集中”为新增论断 |
| R17 | haohao-v0.4.2b | minor |  | “主要靠手工”写成“全靠手工填写”，程度加重 |
| R19 | baseline | minor | GRAI 段名“总结规律”被改为“总结与改进”；“影响第一印象”改写成“先等了很久” |  |
| R19 | haohao-v0.1b | minor | GRAI 段名“分析原因”被改为“原因分析” |  |
| R23 | baseline | minor | “经核查”丢失 | 概述新增原因摘要，“没有设置合理的超时”改变了原意（原为未设置超时） |
| R23 | haohao-v0.1 | minor | “未发现资金损失”改成“没有资金损失”，口径变强 |  |
| R24 | baseline | minor |  | 脱敏示例“比如只保留后四位”；结尾“改完之后 @ 我” |
| R26 | baseline | minor |  | 快速失败、按预案降级、避免被拖慢的作用描述；以依赖库方式接入、不需要额外部署组件；用户能马上感觉到；引入成本不划算 |
| R27 | baseline | minor | “雪天”被改成“小雪” | 水珠不易立刻渗入；体感因人而异，气温偏低时可搭配厚内搭；有助于恢复蓬松 |
| R30 | baseline | minor |  |  |
| R33 | haohao-v0.1b | minor | 成都新开始/新工作只剩一句“那边还顺吧”，交代偏弱 |  |
| R34 | baseline | minor | 老周“带我做”第一个项目的关系表述弱化 |  |
| R35 | baseline | minor | “11 月 6 日前”变“到 11 月 6 日” | “多次报修后一直没有得到明确答复”；“向业主提供” |
| R35 | haohao-v0.1 | minor |  | “4 次”改成“至少 4 次” |
| R35 | haohao-v0.1b | minor | “整整两天”的“整整” |  |
| R38 | baseline | minor |  | 活计和人情一点点铺开，很实在；“有的”方言未注变成“不少” |
| R39 | baseline | minor | “普遍”被删；“提供依据”弱化为“提供参考” |  |
| R39 | haohao-v0.1b | minor | “普遍表明”删去“普遍” |  |
| R40 | baseline | minor |  | 把“有望填补空白”升格为断言“研究尚属空白”；监测滞后成“主要制约” |
| R40 | haohao-v0.1b | minor | “从源头”减少农药使用 |  |
| R43 | baseline | minor |  | 以免干烧；以防触电 |
| R44 | baseline | minor |  | 自行阐释团体格局与差序格局的区别；“只是一个家族的个案，是否普遍还需更多材料”的新论证 |
| W03 | baseline | minor |  | 有问题欢迎随时反馈 |
| W03 | haohao-v0.1b | minor |  | 欢迎在应用内反馈（已在说明中标出） |
| W04 | baseline | minor |  | 承诺后续对接节点提前确认、避免类似偏差 |
| W04 | haohao-v0.4.2 | minor |  | “我们的工作要等文档到位”的推断（已说明） |
| W07 | haohao-v0.1 | minor | 仓库名 billing-service 未写 |  |
| W07 | haohao-v0.1b | minor | 仓库名 billing-service 未写 |  |
| W07 | haohao-v0.4.2 | minor | 仓库名 billing-service 未写 |  |
| W07 | haohao-v0.4.2b | minor | 仓库名 billing-service 未写 |  |
| W08 | haohao-v0.4.2b | minor | 周三 |  |
| W10 | baseline | minor | 到期日期省了“2026 年” |  |
| W10 | haohao-v0.4.2 | minor | “12 月 31 日前到期”改成“2026 年内到期”，范围有变化；“续年卡”写成“续卡” |  |
| W10 | haohao-v0.4.2b | minor | 到期日期省了“2026 年” |  |
| W11 | baseline | minor |  | “以后回县里，我去看您”替作者许诺；批作文习惯“从您那儿来的”；落款加班级 |
| W11 | haohao-v0.4.2b | minor | 篇幅约450字，明显短于600字要求 |  |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| N05 | baseline | 四人参会；调休顺延；给确定答复 |
| N06 | baseline | 已加；至今；遗留一个 |
| R03 | baseline | 希望能帮你少走些弯路 |
| R04 | haohao-v0.1 | 认知上的突破 |
| R05 | haohao-v0.4.2b | 整体上 |
| R06 | haohao-v0.4.2b | 会上大家都很专业；让项目推进得更顺 |
| R12 | baseline | 我不想说大话，就一句；我心里清楚；这件事我一直记着 |
| R12 | haohao-v0.4.2b | 我都看在眼里 |
| R15 | haohao-v0.4.2b | 让续费更透明 |
| R17 | haohao-v0.4.2 | 财务侧的数据也能对上 |
| R19 | haohao-v0.1b | 根本原因 |
| R23 | haohao-v0.4.2 | 根子是 |
| R24 | haohao-v0.4.2 | 你看着办 |
| R27 | baseline | 温馨提示；一件衣服多种穿法；穿着没有负担感 |
| R29 | baseline | 全看缘分 |
| R31 | haohao-v0.1 | 谢谢关注 |
| R32 | baseline | 酱油不只是一瓶调味品，是…… |
| R32 | haohao-v0.4.2b | 我们想把这份时间的味道，送到更多人的餐桌上 |
| R33 | haohao-v0.1 | 顺顺利利的 |
| R37 | haohao-v0.4.2 | “爱情不用轰轰烈烈……就够了”换壳双向奔赴；风光/狼狈对仗大道理 |
| R37 | haohao-v0.4.2b | “爱情不用轰轰烈烈……就够了”换壳双向奔赴；室友、战友、家人排比残留 |
| R38 | haohao-v0.4.2b | 写父女，也写那个时代 |
| W01 | baseline | 【数据库迁移通知】；需要大家配合 |
| W04 | baseline | 想跟您同步一下；感谢您的理解与支持；祝工作顺利！；由此给您的安排带来不便，还请谅解 |
| W06 | baseline | 原因有两个 |
| W08 | baseline | 【】分节模板 |
| W11 | baseline | 不是分析，是这样一句话；讲台上/讲台下对照 |
| W12 | baseline | 说实话；心里很不是滋味 |
