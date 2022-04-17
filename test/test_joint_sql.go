package main

import (
	"MySQLParser-go/parser"
	"MySQLParser-go/parser/tree"
	"strings"
)

func main() {
	sqls := []string{
		"select a, (select x from tab1 where id = t.y) as xx, b from tab2 t left outer join tab3 using(c,d), (select * from tab5) t5 where t.m = ? and exists (select 1 from tab4 t4 where t1.n  = t4.n) and t.tm >= '2016-11-11'",
		"select date '2016-08-08', time '10:20:30'",                                                                      //
		"SELECT 38.8, CAST(38.8 AS CHAR), x'0a0e', 0xabc, X'0a0e', 0Xabc, b'1010', 0B1010",                               //
		"SELECT * FROM t1 WHERE (col1,col2) = (SELECT col3, col4 FROM t2 WHERE id = 10)",                                 //
		"SELECT * FROM t1 WHERE ROW(col1,col2) = ANY (SELECT col3, col4 FROM t2 WHERE id = 10)",                          //
		"select distinct c1, c2 from t1 where (c1, c2) = (1, 2) lock in Share Mode",                                      //
		"select (select column1 from t1) + 5 from t2 FOR UPDATE",                                                         //
		"select 1 as a union all select x from t left join t1 on t.a = t1.b union select 'x' from t2 where x is UNKNOWN", //
		"select sum(a) from tab1 t1, tab2 t2 where t1.id = t2.id group by t1.a, t2.b having count(*) > 1 order by t1.id DESC, t2.id asc",
		"SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10)",                                                      //
		"(SELECT a FROM t1 WHERE a=10 AND B=1) UNION (SELECT a FROM t2 WHERE a=11 AND B=2) ORDER BY a LIMIT 10", //
		"SELECT 1 AS foo UNION SELECT 2 ORDER BY MAX(1)",                                                        //
		"SELECT 1 /* this is an in-line comment */ + 1",                                                         //

		"SELECT 102/(1-1)", //
		"SELECT CONV(10+'10'+'10'+X'0a',10,10),IF(1>2,2,3),COS(PI() / 2)",                                                      //
		"SELECT CASE WHEN 1>0 THEN 'true' ELSE 'false' END",                                                                    //
		"SELECT CASE 1 WHEN 1 THEN 'one' WHEN 2 THEN 'two' ELSE 'more' END",                                                    //
		"SELECT CASE BINARY 'B' WHEN 'a' THEN 1 WHEN 'b' THEN 2 END",                                                           //
		"select date '2016-10-01' + 1",                                                                                         //
		"InSeRt taba(abc_1,str1,str2, create_tiMe,n,ph1,ph2) \nVALUES(?,'a\"b''c',\"a'\"\"b\",current_timestamp,nUlL,:1,:ph2)", //
		"insert into tab (a, b, c) select 1, 2, 3",                                                                             //
		"delete from tt_order_status # test comment",                                                                           //
		"DELETE FROM t1 WHERE s11 > ANY (SELECT COUNT(*) /* no hint */ FROM t2 WHERE NOT EXISTS (SELECT * FROM t3 WHERE ROW(5*t2.s1,77)= (SELECT 50,11*s1 FROM t4 UNION SELECT 50,77 FROM(SELECT * FROM t5) AS t5)))", //

	}
	for _, sql := range sqls {
		println("\n原始SQL: ", sql)

		sqlTree := parser.ParseSQL(sql)
		ss := asldkfasdfalsdf(sqlTree)

		if len(ss) == 0 {
			println("SQL写得很棒！！！")
		} else {
			println("这几个元素应该参数化：", strings.Join(ss, ", "))
		}
	}
}

func asldkfasdfalsdf(stmt tree.Stat) []string {
	var result []string
	for _, stat := range stmt.Children() {
		switch stat.(type) {
		case *tree.ElementTextParam:
			result = append(result, stat.String())
		case *tree.ElementDate:
			result = append(result, stat.String())
		case *tree.SelectSuffix: // todo 未实现limit和offset
			result = append(result, asldkfasdfalsdf(stat)...)
		case *tree.UpdateSingleTable: // todo 未实现limit和offset
			result = append(result, asldkfasdfalsdf(stat)...)
		case *tree.DeleteStat: // todo 未实现limit和offset
			result = append(result, asldkfasdfalsdf(stat)...)
		default:
			result = append(result, asldkfasdfalsdf(stat)...)
		}
	}
	return result
}
