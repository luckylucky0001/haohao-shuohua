# 第十轮：v0.6.1 在 44 条改稿用例上对 v0.6、v0.5.1

判分 44 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.6.1 | 44 | 39/5/0 | 4.52 | 4.20 | 4.66 | **4.46** | 0.57 | 43% | - | 5.8 | 22 | 6 | 8 | 3 |
| haohao-v0.6 | 44 | 39/5/0 | 4.27 | 4.20 | 4.57 | **4.35** | 0.56 | 44% | - | 3.9 | 28 | 0 | 12 | 4 |
| haohao-v0.5.1 | 44 | 40/4/0 | 4.59 | 4.48 | 4.75 | **4.61** | 0.38 | 62% | - | 4.0 | 23 | 1 | 6 | 3 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.6.1 | 43% | 33%–53% |
| haohao-v0.6 | 44% | 31%–57% |
| haohao-v0.5.1 | 62% | 50%–75% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.6.1 | haohao-v0.6 | 22 | 22 | 1.000 |
| haohao-v0.6.1 | haohao-v0.5.1 | 16 | 28 | 0.096 |
| haohao-v0.6 | haohao-v0.5.1 | 17 | 27 | 0.174 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.6.1 | 4.46 / 0.57 | - | - | - | 65% |
| haohao-v0.6 | 4.35 / 0.56 | - | - | - | 65% |
| haohao-v0.5.1 | 4.61 / 0.38 | - | - | - | 65% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| R01 | 改稿 | haohao-v0.6.1 > haohao-v0.5.1 > haohao-v0.6 |
| R02 | 改稿 | haohao-v0.6.1 > haohao-v0.6 > haohao-v0.5.1 |
| R03 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R04 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R05 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R06 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R07 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R08 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R09 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R10 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R11 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R12 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R13 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R14 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R15 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R16 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R17 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R18 | 改稿 | haohao-v0.6.1 > haohao-v0.5.1 > haohao-v0.6 |
| R19 | 改稿 | haohao-v0.6.1 > haohao-v0.5.1 > haohao-v0.6 |
| R20 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R21 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R22 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R23 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R24 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R25 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R26 | 改稿 | haohao-v0.6.1 > haohao-v0.5.1 > haohao-v0.6 |
| R27 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R28 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R29 | 改稿 | haohao-v0.6.1 > haohao-v0.6 > haohao-v0.5.1 |
| R30 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R31 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R32 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R33 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R34 | 改稿 | haohao-v0.6.1 > haohao-v0.6 > haohao-v0.5.1 |
| R35 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R36 | 改稿 | haohao-v0.6 > haohao-v0.5.1 > haohao-v0.6.1 |
| R37 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R38 | 改稿 | haohao-v0.5.1 > haohao-v0.6 > haohao-v0.6.1 |
| R39 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R40 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |
| R41 | 改稿 | haohao-v0.6.1 > haohao-v0.5.1 > haohao-v0.6 |
| R42 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R43 | 改稿 | haohao-v0.5.1 > haohao-v0.6.1 > haohao-v0.6 |
| R44 | 改稿 | haohao-v0.6 > haohao-v0.6.1 > haohao-v0.5.1 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| R01 | haohao-v0.5.1 | minor |  | “做好了准备”比原文“奠定基础”更像已就绪；“不再被噪音淹没”略有发挥 |
| R01 | haohao-v0.6 | minor | “一线” | “做好准备”略加强；“用起来更省事”为推断 |
| R03 | haohao-v0.6 | minor |  | “微服务拆得越细”替换了“架构普及”的起因 |
| R03 | haohao-v0.6.1 | minor |  | “微服务拆多了”把“架构普及”换成拆分过多，原因略有改动 |
| R04 | haohao-v0.6.1 | minor | “最高级的自律”改成“最难的自律”，判断变了 |  |
| R05 | haohao-v0.6.1 | minor | 核心：有理有据、态度诚恳 |  |
| R10 | haohao-v0.6 | minor | 背景句“大语言模型在教育领域的应用受到关注”被并入研究空白，背景信息弱化 |  |
| R23 | haohao-v0.6.1 | minor | “未发现资金损失”改成“没有资金损失”，否定变成确定 |  |
| R27 | haohao-v0.6 | minor |  | “雪天”收窄为“小雪天” |
| R35 | haohao-v0.5.1 | minor | “整整两天”的“整整”被删 |  |
| R37 | haohao-v0.6 | minor |  | “许晴答应了，才有了今天”（说明中已标出） |
| R37 | haohao-v0.6.1 | minor |  | “许晴答应了，所以今天我们才能坐在这儿喝喜酒”（说明中已标出） |
| R40 | haohao-v0.5.1 | minor | 从源头减少农药使用中的“从源头” |  |
| R44 | haohao-v0.5.1 | minor |  | 新增对引文的解释句：中国人的人际关系以自己为中心，一圈圈往外推 |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| R01 | haohao-v0.5.1 | 不再被噪音淹没 |
| R01 | haohao-v0.6 | 做好了；用起来更省事 |
| R06 | haohao-v0.6 | 专业的投入 |
| R08 | haohao-v0.5.1 | 顶住 |
| R24 | haohao-v0.6.1 | 你看着办 |
| R31 | haohao-v0.5.1 | 谢谢您的耐心等待；谢谢您的理解；谢谢咨询 |
| R31 | haohao-v0.6 | 亲～ |
| R34 | haohao-v0.6 | 承蒙 |
| R36 | haohao-v0.6.1 | 先说要您做的 |
