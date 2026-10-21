# 第七轮（只看 50 条新用例）

判分 50 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.4.2 | 50 | 46/4/0 | 4.58 | 4.46 | 4.74 | **4.59** | 0.40 | 60% | 0.00 | 1.2 | 10 | 4 | 5 | 6 |
| haohao-v0.4.2b | 50 | 44/6/0 | 4.58 | 4.58 | 4.82 | **4.66** | 0.41 | 58% | 0.00 | 1.2 | 11 | 2 | 3 | 2 |
| haohao-v0.1 | 50 | 46/4/0 | 4.68 | 4.48 | 4.88 | **4.68** | 0.40 | 60% | 0.00 | 1.1 | 9 | 3 | 7 | 3 |
| haohao-v0.1b | 50 | 43/7/0 | 4.70 | 4.42 | 4.86 | **4.66** | 0.48 | 52% | 0.00 | 0.9 | 8 | 3 | 42 | 4 |
| baseline | 50 | 32/16/2 | 4.48 | 3.84 | 4.58 | **4.30** | 0.81 | 19% | 0.90 | 1.9 | 9 | 7 | 8 | 9 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.4.2 | 60% | 52%–68% |
| haohao-v0.4.2b | 58% | 49%–68% |
| haohao-v0.1 | 60% | 52%–69% |
| haohao-v0.1b | 52% | 44%–60% |
| baseline | 19% | 12%–28% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.4.2 | haohao-v0.4.2b | 24 | 26 | 0.888 |
| haohao-v0.4.2 | haohao-v0.1 | 25 | 25 | 1.000 |
| haohao-v0.4.2 | haohao-v0.1b | 30 | 20 | 0.203 |
| haohao-v0.4.2 | baseline | 41 | 9 | 0.000 |
| haohao-v0.4.2b | haohao-v0.1 | 24 | 26 | 0.888 |
| haohao-v0.4.2b | haohao-v0.1b | 28 | 22 | 0.480 |
| haohao-v0.4.2b | baseline | 39 | 11 | 0.000 |
| haohao-v0.1 | haohao-v0.1b | 29 | 21 | 0.322 |
| haohao-v0.1 | baseline | 41 | 9 | 0.000 |
| haohao-v0.1b | baseline | 41 | 9 | 0.000 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.4.2 | 4.40 / 0.43 | 5.00 / 0.33 | 4.77 / 0.38 | 0% | 74% |
| haohao-v0.4.2b | 4.48 / 0.44 | 5.00 / 0.15 | 4.87 / 0.60 | 0% | 74% |
| haohao-v0.1 | 4.58 / 0.38 | 4.90 / 0.47 | 4.77 / 0.38 | 0% | 73% |
| haohao-v0.1b | 4.56 / 0.46 | 4.87 / 0.60 | 4.77 / 0.42 | 37% | 68% |
| baseline | 4.33 / 0.79 | 4.20 / 0.95 | 4.30 / 0.72 | 6% | 76% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
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
| W06 | baseline | 原因有两个 |
| W08 | baseline | 【】分节模板 |
| W11 | baseline | 不是分析，是这样一句话；讲台上/讲台下对照 |
| W12 | baseline | 说实话；心里很不是滋味 |
