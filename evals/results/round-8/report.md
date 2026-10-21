# 第八轮：v0.5.1 在 72 条通用用例上的回归检查

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.5.1 | 72 | 68/4/0 | 4.54 | 4.43 | 4.78 | **4.58** | 0.38 | 62% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |
| haohao-v0.4.2 | 72 | 70/2/0 | 4.51 | 4.47 | 4.69 | **4.56** | 0.42 | 58% | 0.00 | 2.7 | 20 | 8 | 7 | 10 |
| baseline | 72 | 50/19/3 | 4.50 | 4.10 | 4.65 | **4.42** | 0.71 | 29% | 0.50 | 3.6 | 24 | 12 | 9 | 17 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.5.1 | 62% | 53%–72% |
| haohao-v0.4.2 | 58% | 51%–65% |
| baseline | 29% | 20%–39% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.5.1 | haohao-v0.4.2 | 42 | 30 | 0.195 |
| haohao-v0.5.1 | baseline | 48 | 24 | 0.006 |
| haohao-v0.4.2 | baseline | 54 | 18 | 0.000 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.5.1 | 4.43 / 0.41 | 5.00 / 0.21 | 4.64 / 0.43 | 0% | 73% |
| haohao-v0.4.2 | 4.39 / 0.39 | 5.00 / 0.43 | 4.67 / 0.50 | 0% | 74% |
| baseline | 4.42 / 0.70 | 4.40 / 0.86 | 4.40 / 0.57 | 4% | 79% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| N02 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| N03 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N04 | 对照：原文本来就自然 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| N05 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N06 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N07 | 对照：原文本来就自然 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| N08 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N09 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N11 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N12 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| N13 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| N14 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R01 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R02 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R03 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R05 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R06 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R08 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R09 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R11 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R12 | 改稿 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| R13 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R15 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R17 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R18 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R19 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R20 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R21 | 改稿 | haohao-v0.5.1 > baseline > haohao-v0.4.2 |
| R22 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R23 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R24 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R25 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R26 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R28 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| R29 | 改稿 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| R30 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R31 | 改稿 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| R32 | 改稿 | haohao-v0.5.1 > baseline > haohao-v0.4.2 |
| R33 | 改稿 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| R34 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R35 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R36 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R37 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| R38 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| R40 | 改稿 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| R41 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R42 | 改稿 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| R43 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| R44 | 改稿 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| W01 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W02 | 从零写作 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| W03 | 从零写作 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| W04 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W05 | 从零写作 | haohao-v0.4.2 > haohao-v0.5.1 > baseline |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W07 | 从零写作 | baseline > haohao-v0.4.2 > haohao-v0.5.1 |
| W08 | 从零写作 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| W09 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W10 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W11 | 从零写作 | haohao-v0.5.1 > haohao-v0.4.2 > baseline |
| W12 | 从零写作 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |
| W13 | 从零写作 | haohao-v0.4.2 > baseline > haohao-v0.5.1 |
| W14 | 从零写作 | baseline > haohao-v0.5.1 > haohao-v0.4.2 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| N05 | baseline | minor | “老陈说他”的转述口吻丢失；“给准话”改成“给确定答复”，口语味丢失 |  |
| R01 | baseline | minor |  | 告警分级具体做法“立即处理/一般告警区分开”为推断 |
| R02 | baseline | fail | 国内首个支持离线多人协作的记账应用（声明被删）；下载号召改成了更新提示 | 不用再一笔笔手输；用着有什么问题或想法，欢迎告诉我们 |
| R03 | baseline | minor |  | 接口变慢或报错、光翻日志拼不出经过的场景铺陈；每一步花了多久都看得见 |
| R04 | baseline | minor |  | 焦虑来源的新解释：看不到进步又停不下来；对朋友做法的解读：留给真正想提高的事；结尾新增建议：先别急着加任务，看哪些事可以不做 |
| R05 | baseline | minor | “有理有据”变成“拿数据说话”，“态度诚恳”弱化为“态度好一点” | 不用直接答应；房东两年内不用操心找新租客的推理 |
| R05 | haohao-v0.5.1 | minor | 房东要涨 15% 的背景没写 |  |
| R07 | haohao-v0.5.1 | minor | “巷子深处”弱化为“巷子里” |  |
| R10 | baseline | minor | 贡献中“补充 AI 教育应用的理论研究”被改写为“提供实证依据” |  |
| R13 | baseline | minor |  | 避免再出现同样的情况 |
| R15 | baseline | fail | 刘洋“预计”11月5日前完成的不确定语气被删，变成确定承诺 | 单列“法务确认”待办项 |
| R17 | baseline | fail | 目标三“退单比例降低一半”被改成“报销单数量较第三季度减少一半”，口径和范围都变了；“拍照即可完成录入”被改为“用户提供发票后”；“一期”改成“本期” | 以第三季度为基线；“这两个问题的根源都在手工录入”的推断 |
| R19 | baseline | minor | 签到排队影响第一印象；固定段名“总结规律”被改成“总结与改进” |  |
| R23 | baseline | minor | “经核查”限定丢失 | 概述里“没有设置合理的超时”改变了“未设置超时”的说法 |
| R24 | baseline | minor |  | 脱敏举例“只保留后四位”；结尾“改完之后 @ 我，我再看一遍” |
| R26 | baseline | minor |  | 以依赖库方式接入、不需要额外部署组件；快速失败、按预案降级、避免被渠道拖慢；引入成本不划算 |
| R27 | baseline | minor | 帽子可拆的“约会”场景 | 水珠不易立刻渗入；体感因人而异，气温偏低时可搭配厚内搭；有助于恢复蓬松；保暖和轻便兼顾 |
| R35 | baseline | minor | “我们”统一改为“我” | “多次报修后一直没有得到明确答复”；“向业主提供”书面方案 |
| R35 | haohao-v0.5.1 | minor | “整整两天”丢了“整整” |  |
| R38 | haohao-v0.4.2 | minor |  | 也是在写粮站里那一代人（保留了原稿拔高） |
| R39 | baseline | minor | “普遍表明”的“普遍”；“提供依据”弱化为“提供参考” |  |
| R40 | baseline | minor |  | 把“研究尚属空白”写成既定事实，强于原文的“有望填补”；开头补了“虫情监测及时性尤为重要”的过渡判断 |
| R43 | baseline | minor |  | 以免干烧；以防触电；开启电源 |
| R44 | baseline | minor |  | 对团体格局界限分明的解释；个案不能代表整体的结尾判断 |
| W04 | baseline | minor |  | 承诺后续对接节点提前确认、避免类似偏差；主动提出一起协调新日期 |
| W07 | haohao-v0.4.2 | minor | 仓库名 billing-service 未写 |  |
| W07 | haohao-v0.5.1 | minor | 仓库名 billing-service 未写 |  |
| W11 | baseline | minor |  | “以后回县里，我去看您”承诺；批作文习惯来自老师；落款加班级 |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| N05 | baseline | 四人参会；调休顺延；给确定答复 |
| R03 | baseline | 也不稀奇；折腾；希望能帮你少走些弯路 |
| R04 | baseline | 其实 |
| R06 | haohao-v0.5.1 | 来信是想 |
| R07 | baseline | 所谓远方，其实就在自己心里 |
| R08 | haohao-v0.5.1 | 先把线上顶住；往深查一轮 |
| R12 | haohao-v0.5.1 | 踏踏实实往前走 |
| R19 | haohao-v0.4.2 | 看起来是……其实在…… |
| R20 | baseline | 关于……的通知/各位同事：公文抬头 |
| R23 | haohao-v0.4.2 | 根子是 |
| R24 | haohao-v0.4.2 | 你看着办 |
| R27 | baseline | 保留“温馨提示”；每段加粗句 |
| R28 | haohao-v0.4.2 | 城市还没醒 |
| R35 | baseline | 不能再这样拖下去 |
| R38 | haohao-v0.4.2 | 写粮站也是写一代人 |
| W01 | baseline | 【数据库迁移通知】标题 |
| W02 | baseline | 那是生活质量的账，不是效率的账 |
| W02 | haohao-v0.5.1 | 往往不是……而是等 |
| W03 | baseline | 有问题欢迎随时反馈 |
| W04 | baseline | 同步一下 |
| W11 | baseline | 不是分析，是这样一句话 |
| W12 | baseline | 说实话 |
