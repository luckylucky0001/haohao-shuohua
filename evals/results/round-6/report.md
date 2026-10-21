# 第六轮：v0.4.2 对 v0.1（各两次采样）、不装 skill

判分 22 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.4.2 | 22 | 20/2/0 | 4.50 | 4.36 | 4.68 | **4.52** | 0.42 | 58% | 0.00 | 6.2 | 10 | 4 | 2 | 4 |
| haohao-v0.4.2b | 22 | 17/5/0 | 4.45 | 4.23 | 4.59 | **4.42** | 0.45 | 55% | 0.00 | 8.2 | 10 | 5 | 2 | 3 |
| haohao-v0.1 | 22 | 21/1/0 | 4.50 | 4.18 | 4.77 | **4.48** | 0.44 | 56% | 0.00 | 5.5 | 7 | 4 | 1 | 3 |
| haohao-v0.1b | 22 | 19/3/0 | 4.50 | 4.50 | 4.68 | **4.56** | 0.39 | 61% | 0.00 | 7.9 | 14 | 4 | 2 | 3 |
| baseline | 22 | 13/7/2 | 4.05 | 3.91 | 4.59 | **4.18** | 0.80 | 20% | 0.00 | 7.7 | 15 | 5 | 1 | 8 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.4.2 | 58% | 47%–69% |
| haohao-v0.4.2b | 55% | 41%–68% |
| haohao-v0.1 | 56% | 41%–69% |
| haohao-v0.1b | 61% | 47%–75% |
| baseline | 20% | 9%–33% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.4.2 | haohao-v0.4.2b | 14 | 8 | 0.286 |
| haohao-v0.4.2 | haohao-v0.1 | 11 | 11 | 1.000 |
| haohao-v0.4.2 | haohao-v0.1b | 10 | 12 | 0.832 |
| haohao-v0.4.2 | baseline | 16 | 6 | 0.052 |
| haohao-v0.4.2b | haohao-v0.1 | 12 | 10 | 0.832 |
| haohao-v0.4.2b | haohao-v0.1b | 11 | 11 | 1.000 |
| haohao-v0.4.2b | baseline | 17 | 5 | 0.017 |
| haohao-v0.1 | haohao-v0.1b | 9 | 13 | 0.523 |
| haohao-v0.1 | baseline | 19 | 3 | 0.001 |
| haohao-v0.1b | baseline | 18 | 4 | 0.004 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.4.2 | 4.38 / 0.43 | 5.00 / 0.25 | 4.50 / 0.56 | 0% | 73% |
| haohao-v0.4.2b | 4.21 / 0.50 | 4.92 / 0.44 | 4.67 / 0.31 | 0% | 74% |
| haohao-v0.1 | 4.33 / 0.39 | 4.67 / 0.81 | 4.83 / 0.25 | 0% | 72% |
| haohao-v0.1b | 4.45 / 0.36 | 5.00 / 0.25 | 4.50 / 0.62 | 0% | 71% |
| baseline | 4.10 / 0.82 | 4.75 / 0.75 | 3.92 / 0.75 | 0% | 85% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1 > baseline |
| N02 | 对照：原文本来就自然 | haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.1 > baseline |
| N03 | 对照：原文本来就自然 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1 > baseline |
| N04 | 对照：原文本来就自然 | baseline > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 |
| R01 | 改稿 | haohao-v0.1 > baseline > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.4.2b |
| R02 | 改稿 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.1b > baseline |
| R03 | 改稿 | haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.1b > haohao-v0.1 > baseline |
| R04 | 改稿 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.1 > haohao-v0.4.2b > baseline |
| R05 | 改稿 | haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R06 | 改稿 | haohao-v0.1 > haohao-v0.1b > baseline > haohao-v0.4.2 > haohao-v0.4.2b |
| R07 | 改稿 | haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| R08 | 改稿 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.1b > baseline > haohao-v0.4.2 |
| R09 | 改稿 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.1 > haohao-v0.4.2 > baseline |
| R10 | 改稿 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| R11 | 改稿 | haohao-v0.1b > haohao-v0.4.2 > haohao-v0.4.2b > baseline > haohao-v0.1 |
| R12 | 改稿 | haohao-v0.4.2 > haohao-v0.1 > haohao-v0.4.2b > haohao-v0.1b > baseline |
| R13 | 改稿 | haohao-v0.1b > haohao-v0.1 > baseline > haohao-v0.4.2 > haohao-v0.4.2b |
| R14 | 改稿 | haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b > baseline > haohao-v0.1 |
| W01 | 从零写作 | haohao-v0.4.2b > haohao-v0.1b > haohao-v0.4.2 > haohao-v0.1 > baseline |
| W02 | 从零写作 | haohao-v0.1 > baseline > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.1b |
| W03 | 从零写作 | haohao-v0.4.2b > haohao-v0.1 > haohao-v0.4.2 > baseline > haohao-v0.1b |
| W04 | 从零写作 | haohao-v0.1 > haohao-v0.1b > haohao-v0.4.2b > haohao-v0.4.2 > baseline |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.1b | minor | “一线”值班同学 |  |
| R01 | haohao-v0.4.2b | minor | “一线”值班同学 | “压力验证做完了”比“为大促做准备”说得满；“大量无效告警” |
| R02 | baseline | fail | 国内首个支持离线多人协作的声明；下载号召改成了更新提示 | 不用再一笔笔手输；欢迎反馈的结尾 |
| R03 | baseline | minor |  | 接口变慢或报错、翻日志难拼经过的场景描述；每一步花了多久都看得见；希望能帮你少走些弯路 |
| R04 | baseline | fail | “认知”这一论点被换掉 | 焦虑来源的新解释（看不到进步又停不下来）；忙到晚上很累的描写；结尾新增行动建议 |
| R04 | haohao-v0.1 | minor | “最高级的自律”改为“更高级”（已在说明中提示） |  |
| R04 | haohao-v0.4.2 | minor | “最高级的自律”弱化为“更像自律”（已提示） | “你已经够努力了”比原文判断更绝对 |
| R04 | haohao-v0.4.2b | minor | “最高级的自律”弱化为“也是一种自律”，说明称观点没变 |  |
| R05 | baseline | minor | “有理有据”变成“拿数据说话”，“态度诚恳”弱化为“态度好一点” | “不用直接答应”；两年内不用操心找新租客等展开；“咱们”视角 |
| R05 | haohao-v0.4.2b | minor |  | “就有底气跟房东说 15% 太高了”的推断 |
| R06 | baseline | minor | 合作跟进的开篇交代；泛化的“有问题可联系”收窄为“安排不妥请告知”；期待回复 | 为便于及时沟通项目进展 |
| R08 | haohao-v0.4.2 | minor | 监控由“需要的话我来做”的提议变成了泛泛待办 |  |
| R10 | baseline | minor | 贡献由“补充理论研究”改为“提供实证依据”（已提示确认） |  |
| R11 | baseline | minor |  | “配着吃刚好” |
| R11 | haohao-v0.4.2b | minor | 避坑提醒标识 | “吃得很满意” |
| R12 | baseline | minor |  | 这件事我一直记着；这几个数字挺好看；我不想说大话 |
| R12 | haohao-v0.1b | minor | 2026年改成“今年”，年会若在次年举办会出错 |  |
| R12 | haohao-v0.4.2b | minor |  | 我都看在眼里 |
| W03 | haohao-v0.1b | minor |  | 应用内反馈入口 |
| W04 | baseline | minor |  | 后续对接节点提前确认、避免类似偏差的新承诺；11月15日如有影响一起协调 |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R03 | baseline | 如果你也在折腾……希望能帮你少走些弯路 |
| R04 | baseline | 我现在觉得；你要是也总觉得忙得慌 |
| R04 | haohao-v0.4.2 | 所以我觉得 |
| R04 | haohao-v0.4.2b | 所以我想 |
| R05 | haohao-v0.4.2b | 整体上 |
| R11 | baseline | 氛围感就很足；梦回老北京；推荐大家去试试～ |
| R11 | haohao-v0.1 | 绝了；一口爆汁；太好吃了 |
| R11 | haohao-v0.4.2 | 梦回老北京 |
| R12 | baseline | 我心里清楚；我不想说大话，就一句 |
| W01 | baseline | 【数据库迁移通知】标题；大家好；需要大家配合 |
| W03 | baseline | 有问题欢迎随时反馈 |
| W04 | baseline | 感谢您的理解与支持；祝工作顺利 |
