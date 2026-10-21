# 第九轮：v0.6 在 72 条通用用例上对 v0.5.1、不装 skill

判分 72 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.6 | 72 | 67/5/0 | 4.46 | 4.42 | 4.72 | **4.53** | 0.36 | 64% | 0.07 | 3.0 | 29 | 3 | 12 | 4 |
| haohao-v0.5.1 | 72 | 64/8/0 | 4.53 | 4.44 | 4.82 | **4.60** | 0.37 | 63% | 0.00 | 2.7 | 24 | 2 | 6 | 3 |
| baseline | 72 | 44/25/3 | 4.44 | 3.96 | 4.61 | **4.34** | 0.77 | 23% | 0.50 | 3.6 | 24 | 12 | 9 | 17 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.6 | 64% | 56%–72% |
| haohao-v0.5.1 | 63% | 55%–72% |
| baseline | 23% | 15%–33% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.6 | haohao-v0.5.1 | 34 | 38 | 0.724 |
| haohao-v0.6 | baseline | 58 | 14 | 0.000 |
| haohao-v0.5.1 | baseline | 53 | 19 | 0.000 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.6 | 4.29 / 0.42 | 4.98 / 0.36 | 4.86 / 0.18 | 0% | 74% |
| haohao-v0.5.1 | 4.45 / 0.35 | 5.00 / 0.21 | 4.67 / 0.57 | 0% | 73% |
| baseline | 4.33 / 0.73 | 4.48 / 0.93 | 4.24 / 0.75 | 4% | 79% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| N02 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N03 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N04 | 对照：原文本来就自然 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| N05 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N06 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| N07 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N08 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| N09 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N10 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N11 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| N12 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| N13 | 对照：原文本来就自然 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| N14 | 对照：原文本来就自然 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R01 | 改稿 | baseline > haohao-v0.6 > haohao-v0.5.1 |
| R02 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R05 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R06 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R08 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R11 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R12 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R13 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R15 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R17 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R18 | 改稿 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| R19 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R20 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R22 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R23 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R24 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R25 | 改稿 | baseline > haohao-v0.6 > haohao-v0.5.1 |
| R26 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R28 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R29 | 改稿 | baseline > haohao-v0.6 > haohao-v0.5.1 |
| R30 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R31 | 改稿 | baseline > haohao-v0.6 > haohao-v0.5.1 |
| R32 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R33 | 改稿 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| R34 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R35 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R36 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R37 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R38 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R40 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| R41 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R42 | 改稿 | baseline > haohao-v0.5.1 > haohao-v0.6 |
| R43 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| R44 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W01 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W02 | 从零写作 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| W03 | 从零写作 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| W04 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W05 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W06 | 从零写作 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| W07 | 从零写作 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| W08 | 从零写作 | haohao-v0.5.1 > baseline > haohao-v0.6 |
| W09 | 从零写作 | haohao-v0.5.1 > haohao-v0.6 > baseline |
| W10 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W11 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W12 | 从零写作 | haohao-v0.6 > haohao-v0.5.1 > baseline |
| W13 | 从零写作 | haohao-v0.6 > baseline > haohao-v0.5.1 |
| W14 | 从零写作 | baseline > haohao-v0.6 > haohao-v0.5.1 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| N05 | baseline | minor | “就我们四个”“腰不行”“给准话”等口语原话被改掉 | “表”具体化为“排班表” |
| R01 | haohao-v0.5.1 | minor |  | “为双十一大促做好了准备”语气加强；“不再被噪音淹没”把噪音少说成了绝对 |
| R02 | baseline | fail | 国内首个支持离线多人协作；下载号召改成了老用户更新提示 | 不用再一笔笔手输；欢迎反馈问题或想法 |
| R03 | baseline | minor |  | 接口变慢或报错、光翻日志拼不出经过等场景铺陈；每一步花了多久都看得见；可观测性里专门解决这个问题的一块 |
| R03 | haohao-v0.6 | minor |  | “微服务拆得越细”替换“微服务架构广泛普及”，属轻微推断（已在说明中提示） |
| R04 | baseline | minor |  | 忙到晚上说不清有没有往前走、焦虑由此而来的新论述；结尾新增“看看哪些事可以不做”的建议 |
| R05 | baseline | minor |  | 开头新增“不用直接答应，可以谈”；把双赢具体化为“两年内也不用操心找新租客” |
| R10 | baseline | minor | 贡献“补充理论研究”被改写 | “提供了……的实证依据”（新的贡献表述，已在说明中提示确认） |
| R10 | haohao-v0.6 | minor | 背景句“大语言模型教育应用受到关注”被并掉 |  |
| R11 | baseline | minor |  | 茉莉花茶“配着吃刚好” |
| R12 | baseline | minor |  | “这件事我一直记着”；“这不是我一个人做出来的”等老板心声 |
| R12 | haohao-v0.5.1 | minor | “2026”年份没有点出 | “很多个深夜”的“很多” |
| R12 | haohao-v0.6 | minor | “2026”年份没有点出 |  |
| R13 | baseline | minor |  | 避免再出现同样的情况 |
| R15 | baseline | fail | 刘洋“预计”11月5日前完成，被改成确定日期 | 新增“法务确认”待办条目及“负责人会上未指定”标注 |
| R17 | baseline | fail | 拍照即可完成录入（改成“用户提供发票后”）；目标三“退单比例降低一半”的原口径；“一期”被改成“本期” | 目标三改为“报销单数量较第三季度减少一半”；背景新增“单张平均耗时约 90 秒”；“两个问题的根源都在手工录入”的推断 |
| R19 | baseline | minor | 签到排队影响第一印象 | “很多人一进场就先等了很久” |
| R19 | haohao-v0.5.1 | minor |  | “拉低第一印象”加重为“第一印象很差” |
| R23 | baseline | minor | “经核查”被删，未发现资金损失变成直接断言；删掉对客服团队的感谢（说明里已告知） | 概述新增原因摘要句，且把“未设置超时”说成“没有设置合理的超时” |
| R24 | baseline | minor |  | 脱敏方案举例“只保留后四位”；结尾“改完之后 @ 我，我再看一遍” |
| R26 | baseline | minor |  | 容错目的：快速失败、按预案降级、避免被渠道拖慢；Resilience4j以依赖库接入、不需要额外部署组件；引入成本不划算 |
| R27 | baseline | minor | “雪天”收窄为“小雪” | 水珠不易立刻渗入；体感因人而异，可搭配厚内搭；有助于恢复蓬松 |
| R30 | baseline | minor |  |  |
| R32 | haohao-v0.5.1 | minor | “或许不会想到”改成肯定的“没想到” |  |
| R34 | baseline | minor | 老周“带”他做第一个项目的关系表述变弱 |  |
| R35 | baseline | minor | “11 月 6 日前”改为“到 11 月 6 日”；“我们”统一改为“我” | 多次报修后一直没有得到明确答复；不能再这样拖下去 |
| R35 | haohao-v0.5.1 | minor | “整整两天”的“整整” |  |
| R37 | haohao-v0.6 | minor |  | “我和陈默的事，得从大学说起”；“许晴答应了，才有了今天” |
| R38 | baseline | minor |  | “有的地方没加注”加重为“不少方言没加注释”；补了“活计和人情一点点铺开，很实在”的评价 |
| R39 | baseline | minor | “提供依据”弱化为“提供参考”；“普遍表明”去掉“普遍” | 公园绿地是影响体力活动的重要环境因素（改写原判断） |
| R40 | baseline | minor |  | 把“有望填补空白”改成断言“研究尚属空白”；新增“虫情监测的及时性尤为重要”一类铺垫 |
| R43 | baseline | minor |  | 以免干烧；以防触电；开启电源（原为开启） |
| R44 | baseline | minor |  | 自行解释团体格局与差序格局的区别；“只是一个家族的个案，是否普遍还需更多材料”的限定；“至少在亲属往来场合”的收窄 |
| W04 | baseline | minor |  | 今后提前确认对接节点的承诺；“一起看看怎么协调”的新安排 |
| W07 | haohao-v0.5.1 | minor | 仓库名 billing-service 没写 |  |
| W08 | haohao-v0.6 | minor | 日期后的“周三” |  |
| W10 | baseline | minor | 到期日省略年份2026，未提示 |  |
| W10 | haohao-v0.5.1 | minor | 到期日省略年份2026（已在说明中提示） |  |
| W11 | baseline | minor |  | 批作文时多留意的习惯；“以后回县里，我去看您”的承诺；落款加班级 |
| W11 | haohao-v0.5.1 | minor |  | “我这才知道您今年12月就要退休”编了得知经过 |
| W12 | baseline | minor |  | “这个一定补上”加强了承诺语气；结尾“对不起啊兄弟” |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| N05 | baseline | 四人参会；调休顺延；确定答复 |
| N06 | baseline | “挂了”改成“没通过”的书面腔；“至今没再报错” |
| R02 | haohao-v0.6 | 融入了 AI 能力 |
| R03 | baseline | 也不稀奇；折腾；希望能帮你少走些弯路 |
| R07 | baseline | 所谓……其实 |
| R08 | haohao-v0.5.1 | 顶住；锁住 |
| R12 | baseline | 这几个数字挺好看；我不想说大话 |
| R17 | haohao-v0.5.1 | 最让人头疼 |
| R18 | haohao-v0.5.1 | 这一年我体会最深的是 |
| R27 | baseline | 保暖和轻便兼顾；穿着没有负担感 |
| R29 | haohao-v0.5.1 | 我们等你们 |
| R31 | haohao-v0.6 | 亲～ |
| R33 | haohao-v0.5.1 | 一切顺利 |
| R34 | haohao-v0.6 | 承蒙 |
| R35 | baseline | 不能再这样拖下去 |
| R36 | baseline | 具体安排如下 |
| R37 | haohao-v0.5.1 | “两个人合得来，都肯往对方那边走”是双向奔赴换壳 |
| R37 | haohao-v0.6 | “今天是你们新日子的开头”是新起点换壳 |
| W01 | baseline | 【数据库迁移通知】；需要大家配合 |
| W02 | baseline | 结尾“是……的账，不是……的账”收束句 |
| W02 | haohao-v0.5.1 | “往往不是……而是……”公式 |
| W03 | baseline | 有问题欢迎随时反馈 |
| W04 | baseline | 想跟您同步一下；方便您随时了解情况；感谢您的理解与支持 |
| W11 | baseline | “最需要的不是分析，是这样一句话” |
| W12 | baseline | 说实话；跟你说个事 |
