# 第 14 轮：v0.7.1 对 v0.7、v0.5.1（72 题）

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.7.1 | 72 | 65/7/0 | 4.44 | 4.32 | 4.71 | **4.49** | 0.49 | 51% | 0.00 | 2.9 | 28 | 3 | 6 | 8 |
| haohao-v0.7 | 72 | 68/4/0 | 4.43 | 4.29 | 4.69 | **4.47** | 0.59 | 41% | 0.00 | 2.7 | 26 | 3 | 13 | 7 |
| haohao-v0.5.1 | 72 | 66/6/0 | 4.54 | 4.53 | 4.79 | **4.62** | 0.42 | 58% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.7.1 | 51% | 42%–60% |
| haohao-v0.7 | 41% | 32%–50% |
| haohao-v0.5.1 | 58% | 49%–67% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.7.1 | haohao-v0.7 | 40 | 32 | 0.410 |
| haohao-v0.7.1 | haohao-v0.5.1 | 34 | 38 | 0.724 |
| haohao-v0.7 | haohao-v0.5.1 | 27 | 45 | 0.044 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.7.1 | 4.32 / 0.56 | 4.98 / 0.21 | 4.55 / 0.54 | 0% | 74% |
| haohao-v0.7 | 4.31 / 0.56 | 4.90 / 0.71 | 4.55 / 0.57 | 0% | 74% |
| haohao-v0.5.1 | 4.49 / 0.39 | 4.95 / 0.57 | 4.69 / 0.39 | 0% | 73% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| N02 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| N03 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N04 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N05 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N06 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| N07 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| N08 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| N09 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| N11 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N12 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N13 | 对照：原文本来就自然 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| N14 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R01 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R02 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R05 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R06 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R08 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R11 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R12 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R13 | 改稿 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R15 | 改稿 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R17 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R18 | 改稿 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| R19 | 改稿 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| R20 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R22 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| R23 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R24 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R25 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R26 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R28 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R29 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R30 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| R31 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R32 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R33 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| R34 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R35 | 改稿 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| R36 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| R37 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R38 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R40 | 改稿 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| R41 | 改稿 | haohao-v0.7 > haohao-v0.7.1 > haohao-v0.5.1 |
| R42 | 改稿 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| R43 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| R44 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| W01 | 从零写作 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| W02 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| W03 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| W04 | 从零写作 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| W05 | 从零写作 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.7.1 > haohao-v0.7 |
| W07 | 从零写作 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| W08 | 从零写作 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| W09 | 从零写作 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| W10 | 从零写作 | haohao-v0.7.1 > haohao-v0.7 > haohao-v0.5.1 |
| W11 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |
| W12 | 从零写作 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.7.1 |
| W13 | 从零写作 | haohao-v0.7.1 > haohao-v0.5.1 > haohao-v0.7 |
| W14 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.7.1 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.7 | minor | 告警分级的效果（告警更有用、噪音少） |  |
| R12 | haohao-v0.5.1 | minor | 未明确提到2026年，只说“这一年” |  |
| R12 | haohao-v0.7 | minor | 未明确提到2026年，只说“这一年” |  |
| R14 | haohao-v0.7.1 | minor |  | “更高效”说成“可以快不少”，程度加重 |
| R15 | haohao-v0.5.1 | minor | “暂不调整”丢了“暂”；未决的默认勾选被列在“结论”标题下 |  |
| R18 | haohao-v0.5.1 | minor | “对账颗粒度”改成“对账从按天改成按小时”，粒度含义变模糊 |  |
| R24 | haohao-v0.7.1 | minor | 第3条“好像”的不确定语气被去掉，改成断言 |  |
| R27 | haohao-v0.7.1 | minor | “雪天”收窄成“小雪天气” | 帮助恢复蓬松 |
| R28 | haohao-v0.7 | minor | 喂酵种被改成“去看那罐酵种” |  |
| R30 | haohao-v0.7.1 | minor |  |  |
| R35 | haohao-v0.5.1 | minor | “整整两天”的“整整”被删 |  |
| R38 | haohao-v0.7 | minor | 外地读者会看不懂被弱化为“可能会有点吃力” | 写不过来 |
| R39 | haohao-v0.7.1 | minor | 公园绿地是城市重要的公共空间中的“城市”“重要”定位 |  |
| R40 | haohao-v0.5.1 | minor | “从源头”减少农药使用 |  |
| R40 | haohao-v0.7.1 | minor | “从源头”减少农药使用 |  |
| W10 | haohao-v0.5.1 | minor | 到期日省掉了年份2026 |  |
| W11 | haohao-v0.7.1 | minor | 篇幅约 470 字，明显短于 600 字左右 |  |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R01 | haohao-v0.5.1 | 不再被噪音淹没 |
| R03 | haohao-v0.7 | 绕不开的一环 |
| R05 | haohao-v0.5.1 | 总之 |
| R06 | haohao-v0.7 | 我来跟您同步一下 |
| R08 | haohao-v0.5.1 | 先把线上顶住 |
| R14 | haohao-v0.7.1 | 最后别忘了 |
| R17 | haohao-v0.5.1 | 最让人头疼 |
| R17 | haohao-v0.7 | 最让人头疼 |
| R23 | haohao-v0.7.1 | 辛苦客服同学了 |
| R31 | haohao-v0.7 | 每条结尾都挂“有问题随时问我们” |
| R37 | haohao-v0.5.1 | 都肯往对方那边走 |
| R37 | haohao-v0.7 | 互相往对方那边走 |
| W02 | haohao-v0.5.1 | 不是……而是…… |
| W02 | haohao-v0.7.1 | 第一……第二……分点套路 |
| W10 | haohao-v0.5.1 | 详询 |
| W11 | haohao-v0.5.1 | 能……能……能……；谢谢您，孙老师 |
| W11 | haohao-v0.7 | 一件在……一件在……一件在…… |
| W12 | haohao-v0.7.1 | 说实话 |
