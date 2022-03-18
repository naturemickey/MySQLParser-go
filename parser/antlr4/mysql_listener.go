// Code generated from /Users/mickey/git/MySQLParser-go/parser/antlr4/MySQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package antlr4 // MySQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MySQLListener is a complete listener for a parse tree produced by MySQLParser.
type MySQLListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterTranscationStat is called when entering the transcationStat production.
	EnterTranscationStat(c *TranscationStatContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterInsertStat is called when entering the insertStat production.
	EnterInsertStat(c *InsertStatContext)

	// EnterColumnNames is called when entering the columnNames production.
	EnterColumnNames(c *ColumnNamesContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterSelectStat is called when entering the selectStat production.
	EnterSelectStat(c *SelectStatContext)

	// EnterSelectInner is called when entering the selectInner production.
	EnterSelectInner(c *SelectInnerContext)

	// EnterSelectPrefix is called when entering the selectPrefix production.
	EnterSelectPrefix(c *SelectPrefixContext)

	// EnterSelectSuffix is called when entering the selectSuffix production.
	EnterSelectSuffix(c *SelectSuffixContext)

	// EnterSelectUnionSuffix is called when entering the selectUnionSuffix production.
	EnterSelectUnionSuffix(c *SelectUnionSuffixContext)

	// EnterSelectExprs is called when entering the selectExprs production.
	EnterSelectExprs(c *SelectExprsContext)

	// EnterTables is called when entering the tables production.
	EnterTables(c *TablesContext)

	// EnterTableRel is called when entering the tableRel production.
	EnterTableRel(c *TableRelContext)

	// EnterTableFactor is called when entering the tableFactor production.
	EnterTableFactor(c *TableFactorContext)

	// EnterTableSubQuery is called when entering the tableSubQuery production.
	EnterTableSubQuery(c *TableSubQueryContext)

	// EnterTableRecu is called when entering the tableRecu production.
	EnterTableRecu(c *TableRecuContext)

	// EnterTableJoin is called when entering the tableJoin production.
	EnterTableJoin(c *TableJoinContext)

	// EnterTableJoinSuffix is called when entering the tableJoinSuffix production.
	EnterTableJoinSuffix(c *TableJoinSuffixContext)

	// EnterTableJoinMod is called when entering the tableJoinMod production.
	EnterTableJoinMod(c *TableJoinModContext)

	// EnterJoinCondition is called when entering the joinCondition production.
	EnterJoinCondition(c *JoinConditionContext)

	// EnterGbobExprs is called when entering the gbobExprs production.
	EnterGbobExprs(c *GbobExprsContext)

	// EnterUpdateStat is called when entering the updateStat production.
	EnterUpdateStat(c *UpdateStatContext)

	// EnterUpdateSingleTable is called when entering the updateSingleTable production.
	EnterUpdateSingleTable(c *UpdateSingleTableContext)

	// EnterUpdateMultipleTable is called when entering the updateMultipleTable production.
	EnterUpdateMultipleTable(c *UpdateMultipleTableContext)

	// EnterSetExprs is called when entering the setExprs production.
	EnterSetExprs(c *SetExprsContext)

	// EnterSetExpr is called when entering the setExpr production.
	EnterSetExpr(c *SetExprContext)

	// EnterDeleteStat is called when entering the deleteStat production.
	EnterDeleteStat(c *DeleteStatContext)

	// EnterTableNameAndAlias is called when entering the tableNameAndAlias production.
	EnterTableNameAndAlias(c *TableNameAndAliasContext)

	// EnterTableNameAndAliases is called when entering the tableNameAndAliases production.
	EnterTableNameAndAliases(c *TableNameAndAliasesContext)

	// EnterWhereCondition is called when entering the whereCondition production.
	EnterWhereCondition(c *WhereConditionContext)

	// EnterWhereCondSub is called when entering the whereCondSub production.
	EnterWhereCondSub(c *WhereCondSubContext)

	// EnterWhereCondOp is called when entering the whereCondOp production.
	EnterWhereCondOp(c *WhereCondOpContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExprRelational is called when entering the exprRelational production.
	EnterExprRelational(c *ExprRelationalContext)

	// EnterExprBetweenAnd is called when entering the exprBetweenAnd production.
	EnterExprBetweenAnd(c *ExprBetweenAndContext)

	// EnterExprIsOrIsNotNull is called when entering the exprIsOrIsNotNull production.
	EnterExprIsOrIsNotNull(c *ExprIsOrIsNotNullContext)

	// EnterExprInSelect is called when entering the exprInSelect production.
	EnterExprInSelect(c *ExprInSelectContext)

	// EnterExprInValues is called when entering the exprInValues production.
	EnterExprInValues(c *ExprInValuesContext)

	// EnterExprExists is called when entering the exprExists production.
	EnterExprExists(c *ExprExistsContext)

	// EnterExprNot is called when entering the exprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprLike is called when entering the exprLike production.
	EnterExprLike(c *ExprLikeContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterElementOpFactory is called when entering the elementOpFactory production.
	EnterElementOpFactory(c *ElementOpFactoryContext)

	// EnterElementText is called when entering the elementText production.
	EnterElementText(c *ElementTextContext)

	// EnterElementTextParam is called when entering the elementTextParam production.
	EnterElementTextParam(c *ElementTextParamContext)

	// EnterElementDate is called when entering the elementDate production.
	EnterElementDate(c *ElementDateContext)

	// EnterElementSubQuery is called when entering the elementSubQuery production.
	EnterElementSubQuery(c *ElementSubQueryContext)

	// EnterElementWapperBkt is called when entering the elementWapperBkt production.
	EnterElementWapperBkt(c *ElementWapperBktContext)

	// EnterElementListFactor is called when entering the elementListFactor production.
	EnterElementListFactor(c *ElementListFactorContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterElementOpEle is called when entering the elementOpEle production.
	EnterElementOpEle(c *ElementOpEleContext)

	// EnterElementOpEleSuffix is called when entering the elementOpEleSuffix production.
	EnterElementOpEleSuffix(c *ElementOpEleSuffixContext)

	// EnterElementCase is called when entering the elementCase production.
	EnterElementCase(c *ElementCaseContext)

	// EnterCaseWhenPart is called when entering the caseWhenPart production.
	EnterCaseWhenPart(c *CaseWhenPartContext)

	// EnterElementWithPrefix is called when entering the elementWithPrefix production.
	EnterElementWithPrefix(c *ElementWithPrefixContext)

	// EnterElementRow is called when entering the elementRow production.
	EnterElementRow(c *ElementRowContext)

	// EnterFunCall is called when entering the funCall production.
	EnterFunCall(c *FunCallContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitTranscationStat is called when exiting the transcationStat production.
	ExitTranscationStat(c *TranscationStatContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitInsertStat is called when exiting the insertStat production.
	ExitInsertStat(c *InsertStatContext)

	// ExitColumnNames is called when exiting the columnNames production.
	ExitColumnNames(c *ColumnNamesContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitSelectStat is called when exiting the selectStat production.
	ExitSelectStat(c *SelectStatContext)

	// ExitSelectInner is called when exiting the selectInner production.
	ExitSelectInner(c *SelectInnerContext)

	// ExitSelectPrefix is called when exiting the selectPrefix production.
	ExitSelectPrefix(c *SelectPrefixContext)

	// ExitSelectSuffix is called when exiting the selectSuffix production.
	ExitSelectSuffix(c *SelectSuffixContext)

	// ExitSelectUnionSuffix is called when exiting the selectUnionSuffix production.
	ExitSelectUnionSuffix(c *SelectUnionSuffixContext)

	// ExitSelectExprs is called when exiting the selectExprs production.
	ExitSelectExprs(c *SelectExprsContext)

	// ExitTables is called when exiting the tables production.
	ExitTables(c *TablesContext)

	// ExitTableRel is called when exiting the tableRel production.
	ExitTableRel(c *TableRelContext)

	// ExitTableFactor is called when exiting the tableFactor production.
	ExitTableFactor(c *TableFactorContext)

	// ExitTableSubQuery is called when exiting the tableSubQuery production.
	ExitTableSubQuery(c *TableSubQueryContext)

	// ExitTableRecu is called when exiting the tableRecu production.
	ExitTableRecu(c *TableRecuContext)

	// ExitTableJoin is called when exiting the tableJoin production.
	ExitTableJoin(c *TableJoinContext)

	// ExitTableJoinSuffix is called when exiting the tableJoinSuffix production.
	ExitTableJoinSuffix(c *TableJoinSuffixContext)

	// ExitTableJoinMod is called when exiting the tableJoinMod production.
	ExitTableJoinMod(c *TableJoinModContext)

	// ExitJoinCondition is called when exiting the joinCondition production.
	ExitJoinCondition(c *JoinConditionContext)

	// ExitGbobExprs is called when exiting the gbobExprs production.
	ExitGbobExprs(c *GbobExprsContext)

	// ExitUpdateStat is called when exiting the updateStat production.
	ExitUpdateStat(c *UpdateStatContext)

	// ExitUpdateSingleTable is called when exiting the updateSingleTable production.
	ExitUpdateSingleTable(c *UpdateSingleTableContext)

	// ExitUpdateMultipleTable is called when exiting the updateMultipleTable production.
	ExitUpdateMultipleTable(c *UpdateMultipleTableContext)

	// ExitSetExprs is called when exiting the setExprs production.
	ExitSetExprs(c *SetExprsContext)

	// ExitSetExpr is called when exiting the setExpr production.
	ExitSetExpr(c *SetExprContext)

	// ExitDeleteStat is called when exiting the deleteStat production.
	ExitDeleteStat(c *DeleteStatContext)

	// ExitTableNameAndAlias is called when exiting the tableNameAndAlias production.
	ExitTableNameAndAlias(c *TableNameAndAliasContext)

	// ExitTableNameAndAliases is called when exiting the tableNameAndAliases production.
	ExitTableNameAndAliases(c *TableNameAndAliasesContext)

	// ExitWhereCondition is called when exiting the whereCondition production.
	ExitWhereCondition(c *WhereConditionContext)

	// ExitWhereCondSub is called when exiting the whereCondSub production.
	ExitWhereCondSub(c *WhereCondSubContext)

	// ExitWhereCondOp is called when exiting the whereCondOp production.
	ExitWhereCondOp(c *WhereCondOpContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExprRelational is called when exiting the exprRelational production.
	ExitExprRelational(c *ExprRelationalContext)

	// ExitExprBetweenAnd is called when exiting the exprBetweenAnd production.
	ExitExprBetweenAnd(c *ExprBetweenAndContext)

	// ExitExprIsOrIsNotNull is called when exiting the exprIsOrIsNotNull production.
	ExitExprIsOrIsNotNull(c *ExprIsOrIsNotNullContext)

	// ExitExprInSelect is called when exiting the exprInSelect production.
	ExitExprInSelect(c *ExprInSelectContext)

	// ExitExprInValues is called when exiting the exprInValues production.
	ExitExprInValues(c *ExprInValuesContext)

	// ExitExprExists is called when exiting the exprExists production.
	ExitExprExists(c *ExprExistsContext)

	// ExitExprNot is called when exiting the exprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprLike is called when exiting the exprLike production.
	ExitExprLike(c *ExprLikeContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitElementOpFactory is called when exiting the elementOpFactory production.
	ExitElementOpFactory(c *ElementOpFactoryContext)

	// ExitElementText is called when exiting the elementText production.
	ExitElementText(c *ElementTextContext)

	// ExitElementTextParam is called when exiting the elementTextParam production.
	ExitElementTextParam(c *ElementTextParamContext)

	// ExitElementDate is called when exiting the elementDate production.
	ExitElementDate(c *ElementDateContext)

	// ExitElementSubQuery is called when exiting the elementSubQuery production.
	ExitElementSubQuery(c *ElementSubQueryContext)

	// ExitElementWapperBkt is called when exiting the elementWapperBkt production.
	ExitElementWapperBkt(c *ElementWapperBktContext)

	// ExitElementListFactor is called when exiting the elementListFactor production.
	ExitElementListFactor(c *ElementListFactorContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitElementOpEle is called when exiting the elementOpEle production.
	ExitElementOpEle(c *ElementOpEleContext)

	// ExitElementOpEleSuffix is called when exiting the elementOpEleSuffix production.
	ExitElementOpEleSuffix(c *ElementOpEleSuffixContext)

	// ExitElementCase is called when exiting the elementCase production.
	ExitElementCase(c *ElementCaseContext)

	// ExitCaseWhenPart is called when exiting the caseWhenPart production.
	ExitCaseWhenPart(c *CaseWhenPartContext)

	// ExitElementWithPrefix is called when exiting the elementWithPrefix production.
	ExitElementWithPrefix(c *ElementWithPrefixContext)

	// ExitElementRow is called when exiting the elementRow production.
	ExitElementRow(c *ElementRowContext)

	// ExitFunCall is called when exiting the funCall production.
	ExitFunCall(c *FunCallContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)
}
