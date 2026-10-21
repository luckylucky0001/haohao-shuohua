# 第 12 轮：v0.6.3 对 v0.5.1、v0.6.2（72 题）

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.6.3 | 72 | 68/4/0 | 4.58 | 4.49 | 4.82 | **4.63** | 0.51 | 49% | 0.00 | 3.2 | 24 | 7 | 6 | 6 |
| haohao-v0.5.1 | 72 | 64/8/0 | 4.60 | 4.43 | 4.85 | **4.62** | 0.47 | 53% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |
| haohao-v0.6.2 | 72 | 68/4/0 | 4.47 | 4.39 | 4.81 | **4.56** | 0.53 | 47% | 0.00 | 2.5 | 21 | 5 | 7 | 7 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.6.3 | 49% | 40%–59% |
| haohao-v0.5.1 | 53% | 44%–64% |
| haohao-v0.6.2 | 47% | 38%–56% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.6.3 | haohao-v0.5.1 | 35 | 37 | 0.906 |
| haohao-v0.6.3 | haohao-v0.6.2 | 36 | 36 | 1.000 |
| haohao-v0.5.1 | haohao-v0.6.2 | 40 | 32 | 0.410 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.6.3 | 4.47 / 0.52 | 4.95 / 0.64 | 4.81 / 0.32 | 0% | 74% |
| haohao-v0.5.1 | 4.50 / 0.47 | 4.98 / 0.21 | 4.67 / 0.71 | 0% | 73% |
| haohao-v0.6.2 | 4.37 / 0.51 | 4.98 / 0.64 | 4.71 / 0.46 | 0% | 74% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| N02 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| N03 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| N04 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| N05 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| N06 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| N07 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| N08 | 对照：原文本来就自然 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| N09 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| N11 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| N12 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| N13 | 对照：原文本来就自然 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| N14 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R01 | 改稿 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| R02 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R05 | 改稿 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| R06 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R07 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R08 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R11 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R12 | 改稿 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| R13 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R15 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R17 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R18 | 改稿 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| R19 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R20 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R22 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R23 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R24 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R25 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R26 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R28 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R29 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R30 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R31 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| R32 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R33 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R34 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R35 | 改稿 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| R36 | 改稿 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| R37 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| R38 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| R40 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R41 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R42 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R43 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| R44 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |
| W01 | 从零写作 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| W02 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W03 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W04 | 从零写作 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6.3 |
| W05 | 从零写作 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.6.2 |
| W07 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W08 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W09 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W10 | 从零写作 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| W11 | 从零写作 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6.3 |
| W12 | 从零写作 | haohao-v0.6.2 > haohao-v0.6.3 > haohao-v0.5.1 |
| W13 | 从零写作 | haohao-v0.6.3 > haohao-v0.6.2 > haohao-v0.5.1 |
| W14 | 从零写作 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.6.2 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.5.1 | minor |  | “为双十一做好了准备”语气强于原文；“不再被噪音淹没”说得过满 |
| R01 | haohao-v0.6.2 | minor | 告警分级的效果（告警更有用、噪音少） |  |
| R04 | haohao-v0.6.2 | minor |  | “多半”弱化了原文的断言 |
| R07 | haohao-v0.5.1 | minor | “巷子深处”弱化为“巷子里” |  |
| R12 | haohao-v0.5.1 | minor |  | “很多个深夜”的程度补充；“没有人放弃”说得过满 |
| R15 | haohao-v0.5.1 | minor | “暂不调整”的“暂” | 把“价格展示有优化空间”的判断归到陈立名下；方案描述放进“结论”，像是方案已定 |
| R17 | haohao-v0.6.3 | minor | “高度依赖手工”改成“全靠手工”，程度加重 |  |
| R18 | haohao-v0.5.1 | minor | “颗粒度”改为“对账从按天改成按小时”，粒度含义变模糊 |  |
| R35 | haohao-v0.5.1 | minor | “整整停运了两天”丢了“整整” |  |
| R38 | haohao-v0.6.2 | minor | “瑕不掩瑜”式软化未清掉，换成“缺点没盖过优点”，削弱了三星的批评态度 |  |
| R39 | haohao-v0.6.2 | minor | 删去“普遍”，改变既有研究结论的强度 |  |
| R39 | haohao-v0.6.3 | minor | “提供依据”弱化为“提供参考” |  |
| R40 | haohao-v0.5.1 | minor | “从源头”减少农药使用 |  |
| R40 | haohao-v0.6.3 | minor | “从源头”减少农药使用 |  |
| W10 | haohao-v0.5.1 | minor | 到期日省去年份“2026年” |  |
| W11 | haohao-v0.6.3 | minor | 篇幅约430字，明显短于600字左右 | 推测“纪念册里想必有不少人写您教得好” |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R01 | haohao-v0.5.1 | 不再被噪音淹没 |
| R02 | haohao-v0.6.3 | 这一版加入了 AI 能力 |
| R03 | haohao-v0.6.3 | 正在做这件事的话，可以参考一下 |
| R05 | haohao-v0.5.1 | 总之 |
| R08 | haohao-v0.5.1 | 先把线上顶住；往深查一轮 |
| R08 | haohao-v0.6.2 | 锁住这个值的取法 |
| R17 | haohao-v0.5.1 | 最让人头疼 |
| R17 | haohao-v0.6.3 | 最让人头疼；拍一下就能 |
| R33 | haohao-v0.5.1 | 新工作新城市 |
| R34 | haohao-v0.6.2 | 二话不说 |
| R34 | haohao-v0.6.3 | 二话不说 |
| R37 | haohao-v0.5.1 | 都肯往对方那边走（双向奔赴换壳） |
| R37 | haohao-v0.6.3 | 你朝我走一步，我朝你走一步 |
| R38 | haohao-v0.6.2 | 缺点没盖过优点（瑕不掩瑜换壳） |
| W02 | haohao-v0.5.1 | 不是……而是…… |
| W02 | haohao-v0.6.3 | 说到底 |
