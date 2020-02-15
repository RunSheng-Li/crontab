



#### 项目简介



在传统的crontab中，至少有以下三个痛点：

- 机器故障，任务停止调度
- 任务数量过多，导致服务器硬件资源耗尽，要人工迁移到其他服务器
- 需要ssh到服务器配置crontab，任务的执行状态、日志也很不方便查看

所以一个理想的crontab至少要包含以下两点：

- 高可用：任务不会因为单点故障停止调度
- 扩展性：要有处理大量任务的能力。当服务器硬件资源不够的时候，单纯的加服务器就可以了

本项目正是为了解决上述场景而诞生的



#### 项目架构



![image](data:image/png;base64,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**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**ttp550Xu5UuF31t8DP6r+qzgvX9GZeP7mA7BALVLnOscQ2oGnGQJoCM7RVwaOxLfDpG6X6Q0f8MwIoMN6bttuDe2j2fvEegfRfpZ/8DNn/5r9GP6F1Izq9iVcRCSfT1X0L7xeL0LaSn0DTfztNPF7lF0rf70Anid1znIMoR14niHAhtsn2nPzNsj2M7n/fPzuA2IyxMUzvsu3yAzBN9Fue2gPmn2IQmn6W2Si/VzZRspeob3fx9tKtx9VH9qjW5x+dbcIpS5E+mBvph3OZ9x/tMUxhHbgiaa3VqTC9u3s7OwMdTSL3QbAadj70YTea2BNzYbfF9qD8WtaCrH98n8St8lc7jVvAvL0Fprc/Vvyd/iXYdsQHtbX//cSwLPubw/vZbcvzUVUHMynlp67cE87vKy23xgXxxDagacagl+omdA2vgVj+V7wcdu2fTtDHD82f/vK/aG9Ed1+E+o2rEaBd6h+P5dfa+7+JYV/DbZ/H5qLmnCx0L9fl/cj3u/kRc/E5SKp3UehHd7Tpb+YFMcQ2oGnikP27H3gozC5HuzaWerr+T6qlTC6S2gPmkB7De7pC43rjHeimvCemvHeFNqbsD79hpt+2VFovxj/NWL5/Wrads9dg/nSPe2XpreEdnhZQ18RFccQ2oETC6Hxel90tYEw7Ft3//Yl6HYP3+vm4qWbXY/dhvbW7YXPP5d72efuL48vvuZKaIfzSZ3rHENoBziR/l74UZ93CevLYXvurxxzf7W4CfnDerqLpKj6WX6hHc4n7hf64hhCO8Cp9B+EDTPjYea+e3jGWmgfdLfYtLfz3H5guF+Pe9rhvfS5Ki6OIbQDvLHLTHr410oX7vPP096/vk24Nz7Myt+7D8CjyViPI7QDALCJjPU4QjsAAJvIWI8jtAMAsImM9ThCOwAAm8hYjyO0AwCwiYz1OEI7AACbyFiPI7QDALCJjPU4QjsAAJvIWI8jtAMAsImM9ThCOwAAm8hYjyO0AwCwiYz1OEI7AACbyFiPI7QDALCJjPU4QjsAAJvIWI8jtAMAsImM9ThCOwAAm8hYjyO0AwCwiYz1OEI7AACbyFiP87ahXSmllFJK7V8cQ2hXSimllFK7FccQ2pVSSiml1G7FMYR2pZRSSim1W3EMoV0ppZRSSu1WHOMtQjsAALwyoR0AAContAMAQOWEdgAAqJzQDgAAlRPaAQCgckI7AABUTmgHAIDKCe0AAFA5oR0AAContAMAQOWEdgAAqJzQDgAAlRPaAQCgckI7AABUTmgHAIDKCe0AAFA5oR0AAContAMAQOWEdgAAqJzQDgAAlRPaAQCgckI7AABUTmgHAIDKCe0AAFA5oR0AAContAMAQOWEdgAAqJzQDgAAlRPaAQCgckI7AABUTmgHAIDKCe0AAFA5oR0AAContAMAQOWEdgAAqJzQDgAAlRPaAQCgckI7AABUTmgHAIDKCe0AAFA5oR0AAContAMAQOWEdgAAqNrHx/8HA76wBob+xZYAAAAASUVORK5CYII=)





![整体架构](image/整体架构.png)

- 利用etcd同步全量任务列表到所有worker节点
- 每个worker独立调度全量任务
- 各个worker利用分布式锁抢占，解决并发调度相同任务的问题



#### 主要接口







master功能
1.任务管理接口：新建、修改、查看、删除任务 --> etcd（/cron/jobs）
etcd结构：/cron/jobs/任务名
name 任务名
command shell命令
cronExpr cron表达式
保存到etcd的任务，会被实时同步到所有worker

2.任务日志接口：查看任务执行历史日志 -->mongodb
mongodb结构
jobName 任务名
command shell命令
err 执行报错
output 执行输出
startTime 开始时间
endTime 结束时间
master会请求mongodb，按任务名查看最近的执行日志

3.任务控制接口：强制结束任务接口 -->etcd（/cron/killer）
etcd结构：/cron/killer/任务名
worker监听/cron/killer/目录下put修改操作
msater将要结束的任务名put在/cron/killer/目录下，触发worker立即结束shell任务

4.web管理界面

worker功能

任务同步：监听etcd中/cron/jobs目录变化
任务调度：基于cron表达式计算，触发过期任务
任务执行：协程池并发执行多任务，基于etcd分布式锁抢占
日志保存：捕获任务执行输出，保存到mongodb

监听协程

利用watch API，监听/cron/jobs/和/cron/killer/目录的变化
将变化事件通过channel推送给调度协程，更新内存中的任务信息

调度协程

监听任务变更event，更新内存中维护的任务列表
检查任务cron表达式，扫描到期任务，交给执行协程运行
监听任务控制event，强制中断正在执行中的子进程
监听任务result，更新内存中任务状态，投递执行日志

执行协程

在etcd中抢占分布式乐观锁：/cron/lock/任务名
抢占成功则通过command类执行shell任务
捕获command输出并等待子进程结束，将执行结果投递给调度协程


日志协程

监听调度发来的执行日志，放入一个batch中
对新batch启动定时器，超时未满自动提交
若batch被放满，那么立即提交，并取消自定提交定时器