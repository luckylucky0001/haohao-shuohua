# 文档类第二轮（D2）：v0.5 对 v0.4.2（各两次采样）、不装 skill

判分 26 条用例。质量分是判官给的“去味、自然、语体”三项（1–5）的平均；名次 0 表示每题都排第一，1 表示每题都垫底；胜率是两两比较中排在对方前面的比例。保真：通过 / 小偏差 / 失败。

## 总表

| 方案 | 题数 | 保真 | 去味 | 自然 | 语体 | 质量分 | 平均名次 | 胜率 | 对照组过度修改 | 输出腔调指数 | 残留腔调 | 新增腔调 | 丢失保护项 | 新增保护项 |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| haohao-v0.5 | 26 | 25/1/0 | 4.88 | 4.85 | 5.00 | **4.91** | 0.20 | 80% | 0.00 | 5.1 | 29 | 4 | 1 | 1 |
| haohao-v0.5b | 26 | 24/2/0 | 4.69 | 4.69 | 4.88 | **4.76** | 0.25 | 75% | 0.00 | 4.3 | 27 | 2 | 1 | 0 |
| haohao-v0.4.2 | 26 | 23/3/0 | 4.62 | 4.42 | 4.62 | **4.55** | 0.72 | 28% | 0.00 | 4.8 | 30 | 2 | 1 | 1 |
| haohao-v0.4.2b | 26 | 25/1/0 | 4.54 | 4.35 | 4.65 | **4.51** | 0.69 | 31% | 0.00 | 4.1 | 26 | 3 | 1 | 1 |
| baseline | 26 | 18/6/2 | 4.69 | 4.08 | 4.81 | **4.53** | 0.63 | 37% | 0.00 | 4.2 | 26 | 5 | 2 | 1 |

## 胜率的 95% 区间

按用例有放回重抽样 2000 次。两个方案的区间大幅重叠时，不能说谁更好。

| 方案 | 胜率 | 95% 区间 |
|---|---:|---|
| haohao-v0.5 | 80% | 69%–89% |
| haohao-v0.5b | 75% | 65%–85% |
| haohao-v0.4.2 | 28% | 19%–37% |
| haohao-v0.4.2b | 31% | 20%–42% |
| baseline | 37% | 24%–50% |

## 两两对比

逐题看谁排在前面。p 值是符号检验（假设两者一样好），小于 0.05 才算有差别。

| 方案 A | 方案 B | A 在前 | B 在前 | p 值 |
|---|---|---:|---:|---:|
| haohao-v0.5 | haohao-v0.5b | 13 | 13 | 1.000 |
| haohao-v0.5 | haohao-v0.4.2 | 25 | 1 | 0.000 |
| haohao-v0.5 | haohao-v0.4.2b | 24 | 2 | 0.000 |
| haohao-v0.5 | baseline | 21 | 5 | 0.002 |
| haohao-v0.5b | haohao-v0.4.2 | 23 | 3 | 0.000 |
| haohao-v0.5b | haohao-v0.4.2b | 22 | 4 | 0.001 |
| haohao-v0.5b | baseline | 20 | 6 | 0.009 |
| haohao-v0.4.2 | haohao-v0.4.2b | 14 | 12 | 0.845 |
| haohao-v0.4.2 | baseline | 11 | 15 | 0.557 |
| haohao-v0.4.2b | baseline | 14 | 12 | 0.845 |

“丢失/新增保护项”是检查器按字面核对数字、日期、代码、英文名的结果，中文数字与阿拉伯数字互换也会计入，只作参考；保真以判官逐条核对为准。

## 按用例类型

| 方案 | 改稿 质量/名次 | 对照 质量/名次 | 从零写 质量/名次 | 对照组改动率 | 篇幅比 |
|---|---|---|---|---:|---:|
| haohao-v0.5 | 4.92 / 0.20 | 4.83 / 0.44 | 4.94 / 0.04 | 0% | 74% |
| haohao-v0.5b | 4.73 / 0.25 | 5.00 / 0.12 | 4.67 / 0.33 | 0% | 75% |
| haohao-v0.4.2 | 4.46 / 0.75 | 5.00 / 0.75 | 4.50 / 0.62 | 0% | 77% |
| haohao-v0.4.2b | 4.44 / 0.70 | 5.00 / 0.75 | 4.39 / 0.62 | 0% | 78% |
| baseline | 4.52 / 0.59 | 4.83 / 0.44 | 4.33 / 0.88 | 0% | 78% |

## 逐题名次（从好到差）

| 用例 | 类型 | 名次 |
|---|---|---|
| P01 | 改稿 | baseline > haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 > haohao-v0.4.2b |
| P02 | 改稿 | haohao-v0.5b > haohao-v0.5 > baseline > haohao-v0.4.2 > haohao-v0.4.2b |
| P03 | 改稿 | haohao-v0.5b > baseline > haohao-v0.5 > haohao-v0.4.2 > haohao-v0.4.2b |
| P04 | 改稿 | haohao-v0.5 > haohao-v0.4.2b > baseline > haohao-v0.5b > haohao-v0.4.2 |
| P05 | 改稿 | haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 > baseline > haohao-v0.4.2b |
| P06 | 改稿 | haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2b > baseline > haohao-v0.4.2 |
| P07 | 改稿 | haohao-v0.5 > baseline > haohao-v0.4.2 > haohao-v0.5b > haohao-v0.4.2b |
| P08 | 改稿 | haohao-v0.4.2b > baseline > haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 |
| P09 | 改稿 | haohao-v0.5 > baseline > haohao-v0.5b > haohao-v0.4.2b > haohao-v0.4.2 |
| P10 | 改稿 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2 > baseline > haohao-v0.4.2b |
| P11 | 改稿 | haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| P12 | 改稿 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| P13 | 改稿 | haohao-v0.5 > haohao-v0.4.2b > haohao-v0.5b > haohao-v0.4.2 > baseline |
| P14 | 改稿 | haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2b > baseline > haohao-v0.4.2 |
| P15 | 改稿 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2b > haohao-v0.4.2 > baseline |
| P16 | 改稿 | haohao-v0.5b > haohao-v0.5 > baseline > haohao-v0.4.2 > haohao-v0.4.2b |
| PN01 | 对照：原文本来就自然 | haohao-v0.5b > haohao-v0.5 > baseline > haohao-v0.4.2b > haohao-v0.4.2 |
| PN02 | 对照：原文本来就自然 | baseline > haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 > haohao-v0.4.2b |
| PN03 | 对照：原文本来就自然 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2b > haohao-v0.4.2 > baseline |
| PN04 | 对照：原文本来就自然 | haohao-v0.5b > baseline > haohao-v0.4.2 > haohao-v0.4.2b > haohao-v0.5 |
| PW01 | 从零写作 | haohao-v0.5 > haohao-v0.4.2 > haohao-v0.5b > haohao-v0.4.2b > baseline |
| PW02 | 从零写作 | haohao-v0.5 > haohao-v0.4.2b > haohao-v0.4.2 > haohao-v0.5b > baseline |
| PW03 | 从零写作 | haohao-v0.5b > haohao-v0.5 > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| PW04 | 从零写作 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2 > haohao-v0.4.2b > baseline |
| PW05 | 从零写作 | haohao-v0.5 > haohao-v0.5b > haohao-v0.4.2b > baseline > haohao-v0.4.2 |
| PW06 | 从零写作 | haohao-v0.5 > haohao-v0.5b > baseline > haohao-v0.4.2b > haohao-v0.4.2 |

## 保真问题

| 用例 | 方案 | 判定 | 丢失或改变 | 新增 |
|---|---|---|---|---|
| P06 | haohao-v0.4.2 | minor | 目标引导句漏了“生产”限定，写成拦截未审批的发布 |  |
| P11 | baseline | minor | “告警收敛”术语被改为“告警去重”；“降级兜底”标题改为“降级” |  |
| P15 | baseline | fail |  | 准出新增“积分发放、商城兑换冒烟用例全部通过” |
| P16 | baseline | minor | 第3条“收敛到20路”改为“定为”，丢了收紧的含义 |  |
| P16 | haohao-v0.4.2 | minor | 第3条“收敛到20路”改为“设为”，丢了收紧的含义 |  |
| P16 | haohao-v0.4.2b | minor | 第3条“收敛到20路”改为“限制为”，丢了收紧的含义 |  |
| P16 | haohao-v0.5 | minor | 第3条“收敛到20路”改为“定为”，丢了收紧的含义 |  |
| PW01 | baseline | minor |  | 降至约2,060通以下的换算；1.2问题小节的概括（近四分之一） |
| PW01 | haohao-v0.5b | minor |  | G1衡量方式定义为到退款发起 |
| PW02 | baseline | fail | 超时时长可自定义 | RocketMQ不支持自定义；方案B外部依赖为无；方案A当前不可行 |
| PW02 | haohao-v0.5b | minor | 超时时长可自定义 | 每分钟500条能满足日增3万的推断 |
| PW03 | baseline | minor |  | 收益中新增“isolated 链接杜绝幽灵依赖”的解释 |
| PW04 | baseline | minor |  | 自定对照组为未放量流量；“立即”回滚；“随时可回切”；结论前放量计划不含双十一暂停，偏向已定 |
| PW05 | baseline | minor |  | 第11位看到明确提示；补位/退出后顺位前移；未满员不显示候补入口 |
| PW05 | haohao-v0.4.2 | minor |  | 课程未满员时不显示候补入口（推断，已在说明中标出） |

## 新长出来的腔调

| 用例 | 方案 | 判官指出的新腔调 |
|---|---|---|
| P01 | haohao-v0.4.2b | 错题本连着课堂和家庭 |
| P04 | haohao-v0.5b | 主要卖点 |
| P06 | haohao-v0.4.2 | 自己识别 |
| PW03 | baseline | 现状存在以下问题 |
| PW03 | haohao-v0.5 | 具体情况如下 |
