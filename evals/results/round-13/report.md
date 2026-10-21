# 第 13 轮：v0.7 对 v0.6.3、v0.5.1（72 题）

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.7 | 72 | 66/6/0 | 4.47 | 4.26 | 4.62 | **4.45** | 0.55 | 45% | 0.00 | 2.7 | 26 | 3 | 13 | 7 |
| haohao-v0.6.3 | 72 | 65/7/0 | 4.53 | 4.33 | 4.71 | **4.52** | 0.58 | 42% | 0.00 | 3.2 | 24 | 7 | 6 | 6 |
| haohao-v0.5.1 | 72 | 66/6/0 | 4.61 | 4.56 | 4.82 | **4.66** | 0.37 | 63% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.7 | 45% | 36%–53% |
| haohao-v0.6.3 | 42% | 33%–51% |
| haohao-v0.5.1 | 63% | 54%–72% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.7 | haohao-v0.6.3 | 39 | 33 | 0.556 |
| haohao-v0.7 | haohao-v0.5.1 | 26 | 46 | 0.024 |
| haohao-v0.6.3 | haohao-v0.5.1 | 27 | 45 | 0.044 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.7 | 4.28 / 0.61 | 4.95 / 0.39 | 4.50 / 0.50 | 0% | 74% |
| haohao-v0.6.3 | 4.39 / 0.53 | 4.88 / 0.79 | 4.60 / 0.54 | 0% | 74% |
| haohao-v0.5.1 | 4.58 / 0.35 | 4.95 / 0.32 | 4.62 / 0.46 | 0% | 73% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| N02 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| N03 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| N04 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| N05 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| N06 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| N07 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| N08 | 对照：原文本来就自然 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| N09 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| N11 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| N12 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| N13 | 对照：原文本来就自然 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| N14 | 对照：原文本来就自然 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| R01 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R02 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R05 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R06 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R08 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R11 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R12 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R13 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R14 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R15 | 改稿 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R17 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R18 | 改稿 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| R19 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R20 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R22 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R23 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R24 | 改稿 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| R25 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R26 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R28 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R29 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R30 | 改稿 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| R31 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R32 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R33 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R34 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R35 | 改稿 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| R36 | 改稿 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| R37 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R38 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R40 | 改稿 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| R41 | 改稿 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| R42 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| R43 | 改稿 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| R44 | 改稿 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| W01 | 从零写作 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |
| W02 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| W03 | 从零写作 | haohao-v0.7 > haohao-v0.6.3 > haohao-v0.5.1 |
| W04 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| W05 | 从零写作 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| W07 | 从零写作 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| W08 | 从零写作 | haohao-v0.5.1 > haohao-v0.7 > haohao-v0.6.3 |
| W09 | 从零写作 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| W10 | 从零写作 | haohao-v0.6.3 > haohao-v0.7 > haohao-v0.5.1 |
| W11 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| W12 | 从零写作 | haohao-v0.7 > haohao-v0.5.1 > haohao-v0.6.3 |
| W13 | 从零写作 | haohao-v0.6.3 > haohao-v0.5.1 > haohao-v0.7 |
| W14 | 从零写作 | haohao-v0.5.1 > haohao-v0.6.3 > haohao-v0.7 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.7 | minor | 告警分级的效果（告警更有用、噪音少）被删，只剩“供一线值班同学使用” |  |
| R12 | haohao-v0.5.1 | minor | “回顾 2026 年”改成“这一年”，未说明；“阖家幸福”换成“全家幸福” | “很多个深夜”的程度推断（已说明） |
| R12 | haohao-v0.6.3 | minor | “回顾 2026 年”改成“这一年”（已在说明中告知） |  |
| R12 | haohao-v0.7 | minor | “回顾 2026 年”改成“这一年”，未说明 |  |
| R15 | haohao-v0.5.1 | minor | “暂不调整”的“暂”字省去（“本期”仍保留时限） | 把“价格展示有优化空间”明确归为陈立的观点 |
| R17 | haohao-v0.6.3 | minor | 用户长期不满的表述；“表单好看解决不了”的动机句 | “高度依赖手工”改成“全靠手工”，程度加重 |
| R18 | haohao-v0.5.1 | minor | “颗粒度/细化”字样被删，只写“对账从按天改成按小时”，时间粒度含义变模糊 |  |
| R28 | haohao-v0.7 | minor | “唤醒酵种”（喂酵种）改成“去看那罐酵种”，动作含义变了 |  |
| R33 | haohao-v0.6.3 | minor | 未提新工作，只说“刚去成都” |  |
| R35 | haohao-v0.5.1 | minor | “整整两天”的“整整”被删 |  |
| R37 | haohao-v0.7 | minor | “答应的事一定做到”被收窄为“答应你的事” |  |
| R38 | haohao-v0.7 | minor |  | “写不过来”为新增判断（说明里已交代） |
| R39 | haohao-v0.6.3 | minor | “相对匮乏”的“相对”丢失，改成“仍然不足”；“提供依据”弱化为“提供参考” |  |
| W04 | haohao-v0.6.3 | minor |  | “您可以随时看到推进情况”与每周一次不符，略夸大 |
| W07 | haohao-v0.5.1 | minor | 仓库名 billing-service |  |
| W07 | haohao-v0.6.3 | minor | 仓库名 billing-service |  |
| W07 | haohao-v0.7 | minor | 仓库名 billing-service |  |
| W10 | haohao-v0.5.1 | minor | 到期日期省略了年份2026 |  |
| W11 | haohao-v0.6.3 | minor | 篇幅约430字，明显短于600字要求 |  |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R01 | haohao-v0.5.1 | 不再被噪音淹没 |
| R03 | haohao-v0.6.3 | 必要一环；可以参考一下 |
| R03 | haohao-v0.7 | 绕不开的一环 |
| R05 | haohao-v0.5.1 | 总之 |
| R06 | haohao-v0.7 | 我来跟您同步一下 |
| R08 | haohao-v0.5.1 | 先把线上顶住；再往深查一轮 |
| R08 | haohao-v0.7 | 更彻底地排查一遍 |
| R12 | haohao-v0.5.1 | 踏踏实实往前走 |
| R14 | haohao-v0.7 | 先说改动的大小： |
| R26 | haohao-v0.6.3 | 够我们目前用 |
| R44 | haohao-v0.5.1 | 这段话说的是 |
| W02 | haohao-v0.6.3 | 说到底 |
| W03 | haohao-v0.5.1 | 这次更新做了三件事 |
| W03 | haohao-v0.6.3 | 之前给大家添麻烦了 |
| W08 | haohao-v0.6.3 | 先说结论 |
| W11 | haohao-v0.5.1 | 能……能……能……排比；谢谢您……谢谢您，孙老师 |
