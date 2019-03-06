# tslat
As of first step to do performance review of your application is look to your log files. But it may be tedious to understand time delta between two adjacent log lines. 
This tool is to resquie! It's calculates timestamp delta between to adjacent log lines and out in format you may specify. 

Let's suppose we have following log lines in file:
```console 
2018-09-21 14:43:59,832 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Get evolution scripts by: META-INF/dbevolution/accounting/.index
2018-09-21 14:43:59,845 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Got evolution files: List(1.sql, 2.sql, 3.sql, 4.sql)
2018-09-21 14:43:59,846 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Found: META-INF/dbevolution/accounting/1.sql
2018-09-21 14:43:59,852 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Found: META-INF/dbevolution/accounting/2.sql
```

And we want to see delta between timestamps on each line. We can estimate performance of each operation on high level, for example.
Download and run this nifty tool:
```console
$cat server.log | tslat -delta-format "%9d|" 
        0| 2018-09-21 14:43:59,832 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Get evolution scripts by: META-INF/dbevolution/accounting/.index
       13| 2018-09-21 14:43:59,845 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Got evolution files: List(1.sql, 2.sql, 3.sql, 4.sql)
        1| 2018-09-21 14:43:59,846 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Found: META-INF/dbevolution/accounting/1.sql
        6| 2018-09-21 14:43:59,852 DEBUG [main] r.k.u.d.ClasspathScriptLocator - Found: META-INF/dbevolution/accounting/2.sql
```
I've love to separate delta value from the rest of logs by `|` char, so I added parameter `-delta-format`. It's simplifies further logs processing by other tools like `sed` and `awk`. 
