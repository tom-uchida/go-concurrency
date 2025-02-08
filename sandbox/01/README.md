# 01

並行処理のシンプルな実装例。

```bash
6 tasks
2025/02/09 01:43:52 func1: task1 done.(1sec)
2025/02/09 01:43:52 func1: task4 done.(1sec)
2025/02/09 01:43:53 func1: task2 done.(2sec)
2025/02/09 01:43:53 func1: task6 done.(2sec)
2025/02/09 01:43:54 func1: task5 done.(3sec)
2025/02/09 01:43:54 func1: task3 done.(3sec)
2025/02/09 01:43:55 func2: task3 done.(1sec)
2025/02/09 01:43:55 func2: task6 done.(1sec)
2025/02/09 01:43:56 func2: task5 done.(2sec)
2025/02/09 01:43:56 func2: task1 done.(2sec)
2025/02/09 01:43:56 func2: task4 done.(2sec)
2025/02/09 01:43:57 func2: task2 done.(3sec)
2025/02/09 01:43:59 func3: task3 done.(2sec)
2025/02/09 01:43:59 func3: task5 done.(2sec)
2025/02/09 01:43:59 func3: task4 done.(2sec)
2025/02/09 01:43:59 func3: task2 done.(2sec)
2025/02/09 01:43:59 func3: task1 done.(2sec)
2025/02/09 01:44:00 func3: task6 done.(3sec)
```