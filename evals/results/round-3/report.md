# 第三轮：v0.3.1 对 v0.1、v0.2、不装 skill

判分 22 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.3.1 | 22 | 21/1/0 | 4.45 | 4.27 | 4.45 | **4.39** | 0.36 | 64% | 0.00 | 7.0 | 14 | 2 | 2 | 3 |
| haohao-v0.1 | 22 | 21/1/0 | 4.45 | 4.32 | 4.68 | **4.48** | 0.35 | 65% | 0.00 | 4.6 | 7 | 3 | 1 | 3 |
| haohao-v0.2 | 22 | 16/6/0 | 4.50 | 4.18 | 4.59 | **4.42** | 0.62 | 38% | 0.00 | 6.1 | 7 | 4 | 24 | 3 |
| baseline | 22 | 14/7/1 | 4.23 | 4.00 | 4.64 | **4.29** | 0.67 | 33% | 0.00 | 7.7 | 15 | 5 | 1 | 8 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.3.1 | 4.26 / 0.33 | 5.00 / 0.08 | 4.25 / 0.75 | 0% | 73% |
| haohao-v0.1 | 4.24 / 0.38 | 4.83 / 0.50 | 5.00 / 0.08 | 0% | 72% |
| haohao-v0.2 | 4.36 / 0.62 | 4.42 / 0.83 | 4.67 / 0.42 | 97% | 51% |
| baseline | 4.24 / 0.67 | 4.75 / 0.58 | 4.00 / 0.75 | 0% | 85% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| N01 | 对照：原文本来就自然 | haohao-v0.3.1 > haohao-v0.1 > baseline > haohao-v0.2 |
| N02 | 对照：原文本来就自然 | haohao-v0.3.1 > haohao-v0.2 > haohao-v0.1 > baseline |
| N03 | 对照：原文本来就自然 | haohao-v0.3.1 > haohao-v0.1 > baseline > haohao-v0.2 |
| N04 | 对照：原文本来就自然 | baseline > haohao-v0.3.1 > haohao-v0.1 > haohao-v0.2 |
| R01 | 改稿 | baseline > haohao-v0.3.1 > haohao-v0.2 > haohao-v0.1 |
| R02 | 改稿 | haohao-v0.1 > haohao-v0.3.1 > haohao-v0.2 > baseline |
| R03 | 改稿 | haohao-v0.1 > haohao-v0.3.1 > haohao-v0.2 > baseline |
| R04 | 改稿 | haohao-v0.3.1 > haohao-v0.2 > haohao-v0.1 > baseline |
| R05 | 改稿 | haohao-v0.1 > baseline > haohao-v0.3.1 > haohao-v0.2 |
| R06 | 改稿 | baseline > haohao-v0.1 > haohao-v0.3.1 > haohao-v0.2 |
| R07 | 改稿 | haohao-v0.2 > haohao-v0.3.1 > haohao-v0.1 > baseline |
| R08 | 改稿 | haohao-v0.2 > haohao-v0.1 > haohao-v0.3.1 > baseline |
| R09 | 改稿 | haohao-v0.1 > haohao-v0.2 > haohao-v0.3.1 > baseline |
| R10 | 改稿 | haohao-v0.3.1 > haohao-v0.1 > haohao-v0.2 > baseline |
| R11 | 改稿 | haohao-v0.3.1 > baseline > haohao-v0.2 > haohao-v0.1 |
| R12 | 改稿 | haohao-v0.1 > haohao-v0.3.1 > baseline > haohao-v0.2 |
| R13 | 改稿 | haohao-v0.3.1 > haohao-v0.1 > haohao-v0.2 > baseline |
| R14 | 改稿 | baseline > haohao-v0.3.1 > haohao-v0.1 > haohao-v0.2 |
| W01 | 从零写作 | haohao-v0.2 > haohao-v0.1 > baseline > haohao-v0.3.1 |
| W02 | 从零写作 | haohao-v0.1 > baseline > haohao-v0.2 > haohao-v0.3.1 |
| W03 | 从零写作 | haohao-v0.1 > haohao-v0.2 > haohao-v0.3.1 > baseline |
| W04 | 从零写作 | haohao-v0.1 > haohao-v0.3.1 > haohao-v0.2 > baseline |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| N03 | haohao-v0.2 | minor | 未返回正文，只给了一句结论 |  |
| R01 | haohao-v0.1 | minor | 告警‘噪音少’的效果未直接写出 | 一线值班同学能‘更快’判断告警（原文无此效果） |
| R02 | baseline | fail | 国内首个支持离线多人协作的声明被删；下载号召改成了更新提示 | 不用再一笔笔手输；用着有什么问题或想法，欢迎告诉我们 |
| R03 | baseline | minor |  | 光翻各服务日志难以拼出经过；每一步花了多久都看得见；分布式链路追踪是可观测性里专门解决这个问题的一块 |
| R04 | baseline | minor | “认知的跃迁/突破”这层观点被换成“想清楚往哪儿使劲” | 新增焦虑成因的推论“看不到进步又停不下来”；新增对朋友做法的解读；结尾新增给读者的具体建议 |
| R04 | haohao-v0.2 | minor |  | 结尾新增“比加班更需要自律”的对比 |
| R05 | baseline | minor |  | 开头“不用直接答应”；折中方案补了“两年内不用操心找新租客”的推断；“租金水平”变成“挂多少钱” |
| R05 | haohao-v0.2 | minor | 核心：有理有据、态度诚恳 |  |
| R05 | haohao-v0.3.1 | minor | 核心：有理有据、态度诚恳（说明中称已删） |  |
| R06 | haohao-v0.2 | minor | 有问题可以联系 |  |
| R10 | baseline | minor | 贡献中“丰富理论研究”被改写为“提供实证依据” | “提供了……实证依据”的新贡献表述 |
| R10 | haohao-v0.2 | minor | 背景句“大语言模型教育应用受到关注”被删，仅隐含 |  |
| R12 | baseline | minor |  | 这件事我一直记着；这不是我一个人做出来的 |
| R12 | haohao-v0.2 | minor | 遇到困难不放弃的意思 |  |
| R13 | baseline | minor |  | ‘避免再出现同样的情况’为新增的目的表述 |
| W04 | baseline | minor |  | 今后对接节点提前确认、避免类似偏差的承诺；主动提出一起协调 |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R03 | baseline | 希望能帮你少走些弯路；折腾 |
| R07 | baseline | 所谓……其实就在自己心里 |
| R08 | haohao-v0.3.1 | 把线上的问题压住 |
| R12 | baseline | 我不想说大话 |
| W01 | baseline | 【数据库迁移通知】标题 |
| W02 | baseline | “是生活质量的账，不是效率的账”对举收尾 |
| W03 | baseline | 谢谢你的使用，有问题欢迎随时反馈 |
| W04 | baseline | 还请谅解；感谢您的理解与支持；祝工作顺利 |
