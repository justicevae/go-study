
-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');

-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
SELECT * FROM students WHERE age > 18;

-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
UPDATE students SET grade = '四年级' WHERE name = '张三';

-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
DELETE FROM students WHERE age < 15;





-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- 要求 ：
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务
BEGIN TRANSACTION;

DECLARE @accountA_balance DECIMAL(18, 2);
DECLARE @accountA_id INT = 1;
DECLARE @accountB_id INT = 2;
DECLARE @transfer_amount DECIMAL(18, 2) = 100.00;

-- 检查账户A的余额是否足够
SELECT @accountA_balance = balance
FROM accounts
WHERE id = @accountA_id;

-- 余额不足，回滚
IF @accountA_balance < @transfer_amount
BEGIN
ROLLBACK TRANSACTION;
    RETURN;
END

-- 从账户A扣除金额
UPDATE accounts
SET balance = balance - @transfer_amount
WHERE id = @accountA_id;

-- 向账户B增加金额
UPDATE accounts
SET balance = balance + @transfer_amount
WHERE id = @accountB_id;

-- 记录转账交易
INSERT INTO transactions (from_account_id, to_account_id, amount, transaction_time)
VALUES (@accountA_id, @accountB_id, @transfer_amount, GETDATE());

-- 提交事务
COMMIT TRANSACTION;