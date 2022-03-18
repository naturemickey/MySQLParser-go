// Code generated from /Users/mickey/git/MySQLParser-go/parser/antlr4/MySQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package antlr4 // MySQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMySQLListener is a complete listener for a parse tree produced by MySQLParser.
type BaseMySQLListener struct{}

var _ MySQLListener = &BaseMySQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseMySQLListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseMySQLListener) ExitStat(ctx *StatContext) {}

// EnterTranscationStat is called when production transcationStat is entered.
func (s *BaseMySQLListener) EnterTranscationStat(ctx *TranscationStatContext) {}

// ExitTranscationStat is called when production transcationStat is exited.
func (s *BaseMySQLListener) ExitTranscationStat(ctx *TranscationStatContext) {}

// EnterCommit is called when production commit is entered.
func (s *BaseMySQLListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseMySQLListener) ExitCommit(ctx *CommitContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseMySQLListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseMySQLListener) ExitRollback(ctx *RollbackContext) {}

// EnterInsertStat is called when production insertStat is entered.
func (s *BaseMySQLListener) EnterInsertStat(ctx *InsertStatContext) {}

// ExitInsertStat is called when production insertStat is exited.
func (s *BaseMySQLListener) ExitInsertStat(ctx *InsertStatContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *BaseMySQLListener) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *BaseMySQLListener) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseMySQLListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseMySQLListener) ExitValueList(ctx *ValueListContext) {}

// EnterSelectStat is called when production selectStat is entered.
func (s *BaseMySQLListener) EnterSelectStat(ctx *SelectStatContext) {}

// ExitSelectStat is called when production selectStat is exited.
func (s *BaseMySQLListener) ExitSelectStat(ctx *SelectStatContext) {}

// EnterSelectInner is called when production selectInner is entered.
func (s *BaseMySQLListener) EnterSelectInner(ctx *SelectInnerContext) {}

// ExitSelectInner is called when production selectInner is exited.
func (s *BaseMySQLListener) ExitSelectInner(ctx *SelectInnerContext) {}

// EnterSelectPrefix is called when production selectPrefix is entered.
func (s *BaseMySQLListener) EnterSelectPrefix(ctx *SelectPrefixContext) {}

// ExitSelectPrefix is called when production selectPrefix is exited.
func (s *BaseMySQLListener) ExitSelectPrefix(ctx *SelectPrefixContext) {}

// EnterSelectSuffix is called when production selectSuffix is entered.
func (s *BaseMySQLListener) EnterSelectSuffix(ctx *SelectSuffixContext) {}

// ExitSelectSuffix is called when production selectSuffix is exited.
func (s *BaseMySQLListener) ExitSelectSuffix(ctx *SelectSuffixContext) {}

// EnterSelectUnionSuffix is called when production selectUnionSuffix is entered.
func (s *BaseMySQLListener) EnterSelectUnionSuffix(ctx *SelectUnionSuffixContext) {}

// ExitSelectUnionSuffix is called when production selectUnionSuffix is exited.
func (s *BaseMySQLListener) ExitSelectUnionSuffix(ctx *SelectUnionSuffixContext) {}

// EnterSelectExprs is called when production selectExprs is entered.
func (s *BaseMySQLListener) EnterSelectExprs(ctx *SelectExprsContext) {}

// ExitSelectExprs is called when production selectExprs is exited.
func (s *BaseMySQLListener) ExitSelectExprs(ctx *SelectExprsContext) {}

// EnterTables is called when production tables is entered.
func (s *BaseMySQLListener) EnterTables(ctx *TablesContext) {}

// ExitTables is called when production tables is exited.
func (s *BaseMySQLListener) ExitTables(ctx *TablesContext) {}

// EnterTableRel is called when production tableRel is entered.
func (s *BaseMySQLListener) EnterTableRel(ctx *TableRelContext) {}

// ExitTableRel is called when production tableRel is exited.
func (s *BaseMySQLListener) ExitTableRel(ctx *TableRelContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseMySQLListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseMySQLListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterTableSubQuery is called when production tableSubQuery is entered.
func (s *BaseMySQLListener) EnterTableSubQuery(ctx *TableSubQueryContext) {}

// ExitTableSubQuery is called when production tableSubQuery is exited.
func (s *BaseMySQLListener) ExitTableSubQuery(ctx *TableSubQueryContext) {}

// EnterTableRecu is called when production tableRecu is entered.
func (s *BaseMySQLListener) EnterTableRecu(ctx *TableRecuContext) {}

// ExitTableRecu is called when production tableRecu is exited.
func (s *BaseMySQLListener) ExitTableRecu(ctx *TableRecuContext) {}

// EnterTableJoin is called when production tableJoin is entered.
func (s *BaseMySQLListener) EnterTableJoin(ctx *TableJoinContext) {}

// ExitTableJoin is called when production tableJoin is exited.
func (s *BaseMySQLListener) ExitTableJoin(ctx *TableJoinContext) {}

// EnterTableJoinSuffix is called when production tableJoinSuffix is entered.
func (s *BaseMySQLListener) EnterTableJoinSuffix(ctx *TableJoinSuffixContext) {}

// ExitTableJoinSuffix is called when production tableJoinSuffix is exited.
func (s *BaseMySQLListener) ExitTableJoinSuffix(ctx *TableJoinSuffixContext) {}

// EnterTableJoinMod is called when production tableJoinMod is entered.
func (s *BaseMySQLListener) EnterTableJoinMod(ctx *TableJoinModContext) {}

// ExitTableJoinMod is called when production tableJoinMod is exited.
func (s *BaseMySQLListener) ExitTableJoinMod(ctx *TableJoinModContext) {}

// EnterJoinCondition is called when production joinCondition is entered.
func (s *BaseMySQLListener) EnterJoinCondition(ctx *JoinConditionContext) {}

// ExitJoinCondition is called when production joinCondition is exited.
func (s *BaseMySQLListener) ExitJoinCondition(ctx *JoinConditionContext) {}

// EnterGbobExprs is called when production gbobExprs is entered.
func (s *BaseMySQLListener) EnterGbobExprs(ctx *GbobExprsContext) {}

// ExitGbobExprs is called when production gbobExprs is exited.
func (s *BaseMySQLListener) ExitGbobExprs(ctx *GbobExprsContext) {}

// EnterUpdateStat is called when production updateStat is entered.
func (s *BaseMySQLListener) EnterUpdateStat(ctx *UpdateStatContext) {}

// ExitUpdateStat is called when production updateStat is exited.
func (s *BaseMySQLListener) ExitUpdateStat(ctx *UpdateStatContext) {}

// EnterUpdateSingleTable is called when production updateSingleTable is entered.
func (s *BaseMySQLListener) EnterUpdateSingleTable(ctx *UpdateSingleTableContext) {}

// ExitUpdateSingleTable is called when production updateSingleTable is exited.
func (s *BaseMySQLListener) ExitUpdateSingleTable(ctx *UpdateSingleTableContext) {}

// EnterUpdateMultipleTable is called when production updateMultipleTable is entered.
func (s *BaseMySQLListener) EnterUpdateMultipleTable(ctx *UpdateMultipleTableContext) {}

// ExitUpdateMultipleTable is called when production updateMultipleTable is exited.
func (s *BaseMySQLListener) ExitUpdateMultipleTable(ctx *UpdateMultipleTableContext) {}

// EnterSetExprs is called when production setExprs is entered.
func (s *BaseMySQLListener) EnterSetExprs(ctx *SetExprsContext) {}

// ExitSetExprs is called when production setExprs is exited.
func (s *BaseMySQLListener) ExitSetExprs(ctx *SetExprsContext) {}

// EnterSetExpr is called when production setExpr is entered.
func (s *BaseMySQLListener) EnterSetExpr(ctx *SetExprContext) {}

// ExitSetExpr is called when production setExpr is exited.
func (s *BaseMySQLListener) ExitSetExpr(ctx *SetExprContext) {}

// EnterDeleteStat is called when production deleteStat is entered.
func (s *BaseMySQLListener) EnterDeleteStat(ctx *DeleteStatContext) {}

// ExitDeleteStat is called when production deleteStat is exited.
func (s *BaseMySQLListener) ExitDeleteStat(ctx *DeleteStatContext) {}

// EnterTableNameAndAlias is called when production tableNameAndAlias is entered.
func (s *BaseMySQLListener) EnterTableNameAndAlias(ctx *TableNameAndAliasContext) {}

// ExitTableNameAndAlias is called when production tableNameAndAlias is exited.
func (s *BaseMySQLListener) ExitTableNameAndAlias(ctx *TableNameAndAliasContext) {}

// EnterTableNameAndAliases is called when production tableNameAndAliases is entered.
func (s *BaseMySQLListener) EnterTableNameAndAliases(ctx *TableNameAndAliasesContext) {}

// ExitTableNameAndAliases is called when production tableNameAndAliases is exited.
func (s *BaseMySQLListener) ExitTableNameAndAliases(ctx *TableNameAndAliasesContext) {}

// EnterWhereCondition is called when production whereCondition is entered.
func (s *BaseMySQLListener) EnterWhereCondition(ctx *WhereConditionContext) {}

// ExitWhereCondition is called when production whereCondition is exited.
func (s *BaseMySQLListener) ExitWhereCondition(ctx *WhereConditionContext) {}

// EnterWhereCondSub is called when production whereCondSub is entered.
func (s *BaseMySQLListener) EnterWhereCondSub(ctx *WhereCondSubContext) {}

// ExitWhereCondSub is called when production whereCondSub is exited.
func (s *BaseMySQLListener) ExitWhereCondSub(ctx *WhereCondSubContext) {}

// EnterWhereCondOp is called when production whereCondOp is entered.
func (s *BaseMySQLListener) EnterWhereCondOp(ctx *WhereCondOpContext) {}

// ExitWhereCondOp is called when production whereCondOp is exited.
func (s *BaseMySQLListener) ExitWhereCondOp(ctx *WhereCondOpContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMySQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMySQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExprRelational is called when production exprRelational is entered.
func (s *BaseMySQLListener) EnterExprRelational(ctx *ExprRelationalContext) {}

// ExitExprRelational is called when production exprRelational is exited.
func (s *BaseMySQLListener) ExitExprRelational(ctx *ExprRelationalContext) {}

// EnterExprBetweenAnd is called when production exprBetweenAnd is entered.
func (s *BaseMySQLListener) EnterExprBetweenAnd(ctx *ExprBetweenAndContext) {}

// ExitExprBetweenAnd is called when production exprBetweenAnd is exited.
func (s *BaseMySQLListener) ExitExprBetweenAnd(ctx *ExprBetweenAndContext) {}

// EnterExprIsOrIsNotNull is called when production exprIsOrIsNotNull is entered.
func (s *BaseMySQLListener) EnterExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {}

// ExitExprIsOrIsNotNull is called when production exprIsOrIsNotNull is exited.
func (s *BaseMySQLListener) ExitExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {}

// EnterExprInSelect is called when production exprInSelect is entered.
func (s *BaseMySQLListener) EnterExprInSelect(ctx *ExprInSelectContext) {}

// ExitExprInSelect is called when production exprInSelect is exited.
func (s *BaseMySQLListener) ExitExprInSelect(ctx *ExprInSelectContext) {}

// EnterExprInValues is called when production exprInValues is entered.
func (s *BaseMySQLListener) EnterExprInValues(ctx *ExprInValuesContext) {}

// ExitExprInValues is called when production exprInValues is exited.
func (s *BaseMySQLListener) ExitExprInValues(ctx *ExprInValuesContext) {}

// EnterExprExists is called when production exprExists is entered.
func (s *BaseMySQLListener) EnterExprExists(ctx *ExprExistsContext) {}

// ExitExprExists is called when production exprExists is exited.
func (s *BaseMySQLListener) ExitExprExists(ctx *ExprExistsContext) {}

// EnterExprNot is called when production exprNot is entered.
func (s *BaseMySQLListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production exprNot is exited.
func (s *BaseMySQLListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprLike is called when production exprLike is entered.
func (s *BaseMySQLListener) EnterExprLike(ctx *ExprLikeContext) {}

// ExitExprLike is called when production exprLike is exited.
func (s *BaseMySQLListener) ExitExprLike(ctx *ExprLikeContext) {}

// EnterElement is called when production element is entered.
func (s *BaseMySQLListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseMySQLListener) ExitElement(ctx *ElementContext) {}

// EnterElementOpFactory is called when production elementOpFactory is entered.
func (s *BaseMySQLListener) EnterElementOpFactory(ctx *ElementOpFactoryContext) {}

// ExitElementOpFactory is called when production elementOpFactory is exited.
func (s *BaseMySQLListener) ExitElementOpFactory(ctx *ElementOpFactoryContext) {}

// EnterElementText is called when production elementText is entered.
func (s *BaseMySQLListener) EnterElementText(ctx *ElementTextContext) {}

// ExitElementText is called when production elementText is exited.
func (s *BaseMySQLListener) ExitElementText(ctx *ElementTextContext) {}

// EnterElementTextParam is called when production elementTextParam is entered.
func (s *BaseMySQLListener) EnterElementTextParam(ctx *ElementTextParamContext) {}

// ExitElementTextParam is called when production elementTextParam is exited.
func (s *BaseMySQLListener) ExitElementTextParam(ctx *ElementTextParamContext) {}

// EnterElementDate is called when production elementDate is entered.
func (s *BaseMySQLListener) EnterElementDate(ctx *ElementDateContext) {}

// ExitElementDate is called when production elementDate is exited.
func (s *BaseMySQLListener) ExitElementDate(ctx *ElementDateContext) {}

// EnterElementSubQuery is called when production elementSubQuery is entered.
func (s *BaseMySQLListener) EnterElementSubQuery(ctx *ElementSubQueryContext) {}

// ExitElementSubQuery is called when production elementSubQuery is exited.
func (s *BaseMySQLListener) ExitElementSubQuery(ctx *ElementSubQueryContext) {}

// EnterElementWapperBkt is called when production elementWapperBkt is entered.
func (s *BaseMySQLListener) EnterElementWapperBkt(ctx *ElementWapperBktContext) {}

// ExitElementWapperBkt is called when production elementWapperBkt is exited.
func (s *BaseMySQLListener) ExitElementWapperBkt(ctx *ElementWapperBktContext) {}

// EnterElementListFactor is called when production elementListFactor is entered.
func (s *BaseMySQLListener) EnterElementListFactor(ctx *ElementListFactorContext) {}

// ExitElementListFactor is called when production elementListFactor is exited.
func (s *BaseMySQLListener) ExitElementListFactor(ctx *ElementListFactorContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseMySQLListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseMySQLListener) ExitElementList(ctx *ElementListContext) {}

// EnterElementOpEle is called when production elementOpEle is entered.
func (s *BaseMySQLListener) EnterElementOpEle(ctx *ElementOpEleContext) {}

// ExitElementOpEle is called when production elementOpEle is exited.
func (s *BaseMySQLListener) ExitElementOpEle(ctx *ElementOpEleContext) {}

// EnterElementOpEleSuffix is called when production elementOpEleSuffix is entered.
func (s *BaseMySQLListener) EnterElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {}

// ExitElementOpEleSuffix is called when production elementOpEleSuffix is exited.
func (s *BaseMySQLListener) ExitElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {}

// EnterElementCase is called when production elementCase is entered.
func (s *BaseMySQLListener) EnterElementCase(ctx *ElementCaseContext) {}

// ExitElementCase is called when production elementCase is exited.
func (s *BaseMySQLListener) ExitElementCase(ctx *ElementCaseContext) {}

// EnterCaseWhenPart is called when production caseWhenPart is entered.
func (s *BaseMySQLListener) EnterCaseWhenPart(ctx *CaseWhenPartContext) {}

// ExitCaseWhenPart is called when production caseWhenPart is exited.
func (s *BaseMySQLListener) ExitCaseWhenPart(ctx *CaseWhenPartContext) {}

// EnterElementWithPrefix is called when production elementWithPrefix is entered.
func (s *BaseMySQLListener) EnterElementWithPrefix(ctx *ElementWithPrefixContext) {}

// ExitElementWithPrefix is called when production elementWithPrefix is exited.
func (s *BaseMySQLListener) ExitElementWithPrefix(ctx *ElementWithPrefixContext) {}

// EnterElementRow is called when production elementRow is entered.
func (s *BaseMySQLListener) EnterElementRow(ctx *ElementRowContext) {}

// ExitElementRow is called when production elementRow is exited.
func (s *BaseMySQLListener) ExitElementRow(ctx *ElementRowContext) {}

// EnterFunCall is called when production funCall is entered.
func (s *BaseMySQLListener) EnterFunCall(ctx *FunCallContext) {}

// ExitFunCall is called when production funCall is exited.
func (s *BaseMySQLListener) ExitFunCall(ctx *FunCallContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseMySQLListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseMySQLListener) ExitParamList(ctx *ParamListContext) {}
