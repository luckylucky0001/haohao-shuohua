# 第 11 轮：v0.6.2 对 v0.5.1、v0.6（72 题）

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.6.2 | 72 | 69/3/0 | 4.42 | 4.33 | 4.78 | **4.51** | 0.49 | 51% | 0.00 | 2.5 | 21 | 5 | 7 | 7 |
| haohao-v0.5.1 | 72 | 66/6/0 | 4.62 | 4.58 | 4.79 | **4.67** | 0.40 | 60% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |
| haohao-v0.6 | 72 | 67/5/0 | 4.46 | 4.39 | 4.76 | **4.54** | 0.60 | 40% | 0.07 | 3.0 | 29 | 3 | 12 | 4 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.6.2 | 51% | 42%–60% |
| haohao-v0.5.1 | 60% | 51%–69% |
| haohao-v0.6 | 40% | 31%–49% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.6.2 | haohao-v0.5.1 | 32 | 40 | 0.410 |
| haohao-v0.6.2 | haohao-v0.6 | 41 | 31 | 0.289 |
| haohao-v0.5.1 | haohao-v0.6 | 46 | 26 | 0.024 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.6.2 | 4.37 / 0.45 | 4.95 / 0.57 | 4.50 / 0.54 | 0% | 74% |
| haohao-v0.5.1 | 4.54 / 0.39 | 5.00 / 0.29 | 4.74 / 0.57 | 0% | 73% |
| haohao-v0.6 | 4.32 / 0.66 | 4.98 / 0.64 | 4.79 / 0.39 | 0% | 74% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| N02 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| N03 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| N04 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| N05 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| N06 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| N07 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| N08 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| N09 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| N11 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| N12 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| N13 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| N14 | 对照：原文本来就自然 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R01 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R02 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R04 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R05 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R06 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R08 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R11 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R12 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R13 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R15 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R17 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R18 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| R19 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R20 | 改稿 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R22 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R23 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R24 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R25 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R26 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R28 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R29 | 改稿 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| R30 | 改稿 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| R31 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R32 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| R33 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R34 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R35 | 改稿 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| R36 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R37 | 改稿 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| R38 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R40 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| R41 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R42 | 改稿 | haohao-v0.5.1 > haohao-v0.6.2 > haohao-v0.6 |
| R43 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| R44 | 改稿 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| W01 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| W02 | 从零写作 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| W03 | 从零写作 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| W04 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| W05 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.2 |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| W07 | 从零写作 | haohao-v0.6 > haohao-v0.6.2 > haohao-v0.5.1 |
| W08 | 从零写作 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| W09 | 从零写作 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |
| W10 | 从零写作 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| W11 | 从零写作 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| W12 | 从零写作 | haohao-v0.6.2 > haohao-v0.6 > haohao-v0.5.1 |
| W13 | 从零写作 | haohao-v0.6.2 > haohao-v0.5.1 > haohao-v0.6 |
| W14 | 从零写作 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.2 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.6.2 | minor | 告警分级的效果（告警更有用、噪音少） |  |
| R03 | haohao-v0.6 | minor |  | “微服务拆得越细”把“普及”改成拆分变细的前提 |
| R10 | haohao-v0.6 | minor | 背景句“大语言模型在教育领域的应用受到关注”被并掉 |  |
| R23 | haohao-v0.5.1 | minor | 概述中“严重”“大量”的程度描述 |  |
| R27 | haohao-v0.6 | minor | “雪天”改成“小雪天” |  |
| R35 | haohao-v0.5.1 | minor | “整整两天”的“整整”被删 |  |
| R37 | haohao-v0.6 | minor |  | “许晴答应了，才有了今天” |
| R38 | haohao-v0.6.2 | minor | 删“瑕不掩瑜”后又换成同义软化，削弱三星的批评态度 |  |
| R39 | haohao-v0.6.2 | minor | “普遍表明”删了“普遍”，削弱了原文对既有研究的概括 |  |
| R40 | haohao-v0.5.1 | minor | “从源头上”减少农药使用被删；“茶业增值”未说明即删去 |  |
| W07 | haohao-v0.5.1 | minor | 仓库名 billing-service 未出现 |  |
| W08 | haohao-v0.6 | minor | 周三 |  |
| W10 | haohao-v0.5.1 | minor | 到期日期省去年份2026年 |  |
| W11 | haohao-v0.5.1 | minor |  | “我这才知道您要退休”的知情经过 |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R01 | haohao-v0.6 | 做好了 |
| R02 | haohao-v0.6 | 融入了 AI 能力 |
| R08 | haohao-v0.5.1 | 先把线上顶住；往深查一轮 |
| R12 | haohao-v0.5.1 | 踏踏实实往前走 |
| R16 | haohao-v0.6 | 围绕 |
| R17 | haohao-v0.5.1 | 最让人头疼 |
| R31 | haohao-v0.6 | 有问题随时问我们；有问题随时联系客服 |
| R33 | haohao-v0.5.1 | 新工作新城市，一切顺利 |
| R34 | haohao-v0.6 | 承蒙 |
| R37 | haohao-v0.5.1 | 感情不靠轰轰烈烈的誓言……都肯往对方那边走（“不是……而是……双向奔赴”换壳） |
| R38 | haohao-v0.6.2 | 缺点没盖过优点（瑕不掩瑜换壳） |
| R44 | haohao-v0.5.1 | 这段话说的是 |
| W02 | haohao-v0.5.1 | 不是……而是…… |
| W02 | haohao-v0.6.2 | 好处落在个人/团队的对仗句 |
| W04 | haohao-v0.6.2 | 特此向您说明；方便您随时掌握情况 |
| W11 | haohao-v0.5.1 | 我不太会写漂亮话 |
| W11 | haohao-v0.6 | 第一件……第二件……第三件 |
| W11 | haohao-v0.6.2 | 我不太会写漂亮话，就把事情说清楚 |
