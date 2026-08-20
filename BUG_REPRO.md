# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

检察服务站有一条分办规则同时配置了“欠薪”和“工伤”两个关键词，原意是命中任意一个就交给权益组，但只写“欠薪”的诉求现在仍被判定为不匹配。请修复多关键词匹配，并兼容人工输入的空格和大小写差异。请不要修改测试文件，它们用于验证这次修复。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-11
- 仓库地址：https://github.com/11DingKing/rider-rights-task-11.git
- parent SHA：b12e1ba7512a6a2e7ca999a1652ce25b9c0b14c9

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-11.git bug-repro
cd bug-repro
git checkout --detach b12e1ba7512a6a2e7ca999a1652ce25b9c0b14c9
go test ./internal/domain -run "^TestRuleMatchesAnyConfiguredKeyword$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestRuleMatchesAnyConfiguredKeyword$" -count=1
--- FAIL: TestRuleMatchesAnyConfiguredKeyword (0.00s)
    task11_test.go:9: rule required every configured keyword instead of any keyword
FAIL
FAIL	riderguard/internal/domain	0.058s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestRuleMatchesAnyConfiguredKeyword$" -count=1
--- FAIL: TestRuleMatchesAnyConfiguredKeyword (0.00s)
    task11_test.go:9: rule required every configured keyword instead of any keyword
FAIL
FAIL	riderguard/internal/domain	0.010s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，一条规则配置多个关键词时，诉求命中任意一个规范化后的关键词即可匹配，不要求同时包含全部关键词；无关键词规则、类别条件和默认规则行为保持不变。定向验证、相关包测试和仓库全量回归必须通过，不得删除、跳过或削弱测试。
