package parser

import (
	. "MySQLParser-go/parser/antlr4"
	"MySQLParser-go/parser/tree"
	"MySQLParser-go/parser/tree/util"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type MySQLListenerGo struct {
	stat        tree.Stat
	parentStack *util.StatStack
}

var _ MySQLListener = &MySQLListenerGo{}

func NewMySQLListener() MySQLListener {
	return new(MySQLListenerGo)
}

// EnterEveryRule is called when any rule is entered.
func (this *MySQLListenerGo) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *MySQLListenerGo) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (this *MySQLListenerGo) EnterSelectStat(ctx *SelectStatContext) {
	selectStat := &tree.SelectStat{}
	if this.parentStack.Top() == nil {
		this.stat = selectStat
	} else {
		// todo
	}
	this.parentStack.Push(selectStat)
}
func (this *MySQLListenerGo) ExitSelectStat(ctx *SelectStatContext) {
	this.parentStack.Pop()
}

func (this *MySQLListenerGo) EnterSelectInner(ctx *SelectInnerContext) {
	selectInner := &tree.SelectInner{}
	selectStat := this.parentStack.Top().(*tree.SelectStat)
	selectStat.SetSelectInner(selectInner)

	this.parentStack.Push(selectInner)
}

func (this *MySQLListenerGo) ExitSelectInner(ctx *SelectInnerContext) {
	this.parentStack.Pop()
}

// EnterSelectPrefix is called when production selectPrefix is entered.
func (s *MySQLListenerGo) EnterSelectPrefix(ctx *SelectPrefixContext) {}

// ExitSelectPrefix is called when production selectPrefix is exited.
func (s *MySQLListenerGo) ExitSelectPrefix(ctx *SelectPrefixContext) {}

// VisitTerminal is called when a terminal node is visited.
func (s *MySQLListenerGo) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *MySQLListenerGo) VisitErrorNode(node antlr.ErrorNode) {}

// EnterStat is called when production stat is entered.
func (s *MySQLListenerGo) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *MySQLListenerGo) ExitStat(ctx *StatContext) {}

// EnterTranscationStat is called when production transcationStat is entered.
func (s *MySQLListenerGo) EnterTranscationStat(ctx *TranscationStatContext) {}

// ExitTranscationStat is called when production transcationStat is exited.
func (s *MySQLListenerGo) ExitTranscationStat(ctx *TranscationStatContext) {}

// EnterCommit is called when production commit is entered.
func (s *MySQLListenerGo) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *MySQLListenerGo) ExitCommit(ctx *CommitContext) {}

// EnterRollback is called when production rollback is entered.
func (s *MySQLListenerGo) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *MySQLListenerGo) ExitRollback(ctx *RollbackContext) {}

// EnterInsertStat is called when production insertStat is entered.
func (s *MySQLListenerGo) EnterInsertStat(ctx *InsertStatContext) {}

// ExitInsertStat is called when production insertStat is exited.
func (s *MySQLListenerGo) ExitInsertStat(ctx *InsertStatContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *MySQLListenerGo) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *MySQLListenerGo) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterValueList is called when production valueList is entered.
func (s *MySQLListenerGo) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *MySQLListenerGo) ExitValueList(ctx *ValueListContext) {}

// EnterSelectSuffix is called when production selectSuffix is entered.
func (s *MySQLListenerGo) EnterSelectSuffix(ctx *SelectSuffixContext) {}

// ExitSelectSuffix is called when production selectSuffix is exited.
func (s *MySQLListenerGo) ExitSelectSuffix(ctx *SelectSuffixContext) {}

// EnterSelectUnionSuffix is called when production selectUnionSuffix is entered.
func (s *MySQLListenerGo) EnterSelectUnionSuffix(ctx *SelectUnionSuffixContext) {}

// ExitSelectUnionSuffix is called when production selectUnionSuffix is exited.
func (s *MySQLListenerGo) ExitSelectUnionSuffix(ctx *SelectUnionSuffixContext) {}

// EnterSelectExprs is called when production selectExprs is entered.
func (s *MySQLListenerGo) EnterSelectExprs(ctx *SelectExprsContext) {}

// ExitSelectExprs is called when production selectExprs is exited.
func (s *MySQLListenerGo) ExitSelectExprs(ctx *SelectExprsContext) {}

// EnterTables is called when production tables is entered.
func (s *MySQLListenerGo) EnterTables(ctx *TablesContext) {}

// ExitTables is called when production tables is exited.
func (s *MySQLListenerGo) ExitTables(ctx *TablesContext) {}

// EnterTableRel is called when production tableRel is entered.
func (s *MySQLListenerGo) EnterTableRel(ctx *TableRelContext) {}

// ExitTableRel is called when production tableRel is exited.
func (s *MySQLListenerGo) ExitTableRel(ctx *TableRelContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *MySQLListenerGo) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *MySQLListenerGo) ExitTableFactor(ctx *TableFactorContext) {}

// EnterTableSubQuery is called when production tableSubQuery is entered.
func (s *MySQLListenerGo) EnterTableSubQuery(ctx *TableSubQueryContext) {}

// ExitTableSubQuery is called when production tableSubQuery is exited.
func (s *MySQLListenerGo) ExitTableSubQuery(ctx *TableSubQueryContext) {}

// EnterTableRecu is called when production tableRecu is entered.
func (s *MySQLListenerGo) EnterTableRecu(ctx *TableRecuContext) {}

// ExitTableRecu is called when production tableRecu is exited.
func (s *MySQLListenerGo) ExitTableRecu(ctx *TableRecuContext) {}

// EnterTableJoin is called when production tableJoin is entered.
func (s *MySQLListenerGo) EnterTableJoin(ctx *TableJoinContext) {}

// ExitTableJoin is called when production tableJoin is exited.
func (s *MySQLListenerGo) ExitTableJoin(ctx *TableJoinContext) {}

// EnterTableJoinSuffix is called when production tableJoinSuffix is entered.
func (s *MySQLListenerGo) EnterTableJoinSuffix(ctx *TableJoinSuffixContext) {}

// ExitTableJoinSuffix is called when production tableJoinSuffix is exited.
func (s *MySQLListenerGo) ExitTableJoinSuffix(ctx *TableJoinSuffixContext) {}

// EnterTableJoinMod is called when production tableJoinMod is entered.
func (s *MySQLListenerGo) EnterTableJoinMod(ctx *TableJoinModContext) {}

// ExitTableJoinMod is called when production tableJoinMod is exited.
func (s *MySQLListenerGo) ExitTableJoinMod(ctx *TableJoinModContext) {}

// EnterJoinCondition is called when production joinCondition is entered.
func (s *MySQLListenerGo) EnterJoinCondition(ctx *JoinConditionContext) {}

// ExitJoinCondition is called when production joinCondition is exited.
func (s *MySQLListenerGo) ExitJoinCondition(ctx *JoinConditionContext) {}

// EnterGbobExprs is called when production gbobExprs is entered.
func (s *MySQLListenerGo) EnterGbobExprs(ctx *GbobExprsContext) {}

// ExitGbobExprs is called when production gbobExprs is exited.
func (s *MySQLListenerGo) ExitGbobExprs(ctx *GbobExprsContext) {}

// EnterUpdateStat is called when production updateStat is entered.
func (s *MySQLListenerGo) EnterUpdateStat(ctx *UpdateStatContext) {}

// ExitUpdateStat is called when production updateStat is exited.
func (s *MySQLListenerGo) ExitUpdateStat(ctx *UpdateStatContext) {}

// EnterUpdateSingleTable is called when production updateSingleTable is entered.
func (s *MySQLListenerGo) EnterUpdateSingleTable(ctx *UpdateSingleTableContext) {}

// ExitUpdateSingleTable is called when production updateSingleTable is exited.
func (s *MySQLListenerGo) ExitUpdateSingleTable(ctx *UpdateSingleTableContext) {}

// EnterUpdateMultipleTable is called when production updateMultipleTable is entered.
func (s *MySQLListenerGo) EnterUpdateMultipleTable(ctx *UpdateMultipleTableContext) {}

// ExitUpdateMultipleTable is called when production updateMultipleTable is exited.
func (s *MySQLListenerGo) ExitUpdateMultipleTable(ctx *UpdateMultipleTableContext) {}

// EnterSetExprs is called when production setExprs is entered.
func (s *MySQLListenerGo) EnterSetExprs(ctx *SetExprsContext) {}

// ExitSetExprs is called when production setExprs is exited.
func (s *MySQLListenerGo) ExitSetExprs(ctx *SetExprsContext) {}

// EnterSetExpr is called when production setExpr is entered.
func (s *MySQLListenerGo) EnterSetExpr(ctx *SetExprContext) {}

// ExitSetExpr is called when production setExpr is exited.
func (s *MySQLListenerGo) ExitSetExpr(ctx *SetExprContext) {}

// EnterDeleteStat is called when production deleteStat is entered.
func (s *MySQLListenerGo) EnterDeleteStat(ctx *DeleteStatContext) {}

// ExitDeleteStat is called when production deleteStat is exited.
func (s *MySQLListenerGo) ExitDeleteStat(ctx *DeleteStatContext) {}

// EnterTableNameAndAlias is called when production tableNameAndAlias is entered.
func (s *MySQLListenerGo) EnterTableNameAndAlias(ctx *TableNameAndAliasContext) {}

// ExitTableNameAndAlias is called when production tableNameAndAlias is exited.
func (s *MySQLListenerGo) ExitTableNameAndAlias(ctx *TableNameAndAliasContext) {}

// EnterTableNameAndAliases is called when production tableNameAndAliases is entered.
func (s *MySQLListenerGo) EnterTableNameAndAliases(ctx *TableNameAndAliasesContext) {}

// ExitTableNameAndAliases is called when production tableNameAndAliases is exited.
func (s *MySQLListenerGo) ExitTableNameAndAliases(ctx *TableNameAndAliasesContext) {}

// EnterWhereCondition is called when production whereCondition is entered.
func (s *MySQLListenerGo) EnterWhereCondition(ctx *WhereConditionContext) {}

// ExitWhereCondition is called when production whereCondition is exited.
func (s *MySQLListenerGo) ExitWhereCondition(ctx *WhereConditionContext) {}

// EnterWhereCondSub is called when production whereCondSub is entered.
func (s *MySQLListenerGo) EnterWhereCondSub(ctx *WhereCondSubContext) {}

// ExitWhereCondSub is called when production whereCondSub is exited.
func (s *MySQLListenerGo) ExitWhereCondSub(ctx *WhereCondSubContext) {}

// EnterWhereCondOp is called when production whereCondOp is entered.
func (s *MySQLListenerGo) EnterWhereCondOp(ctx *WhereCondOpContext) {}

// ExitWhereCondOp is called when production whereCondOp is exited.
func (s *MySQLListenerGo) ExitWhereCondOp(ctx *WhereCondOpContext) {}

// EnterExpression is called when production expression is entered.
func (s *MySQLListenerGo) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *MySQLListenerGo) ExitExpression(ctx *ExpressionContext) {}

// EnterExprRelational is called when production exprRelational is entered.
func (s *MySQLListenerGo) EnterExprRelational(ctx *ExprRelationalContext) {}

// ExitExprRelational is called when production exprRelational is exited.
func (s *MySQLListenerGo) ExitExprRelational(ctx *ExprRelationalContext) {}

// EnterExprBetweenAnd is called when production exprBetweenAnd is entered.
func (s *MySQLListenerGo) EnterExprBetweenAnd(ctx *ExprBetweenAndContext) {}

// ExitExprBetweenAnd is called when production exprBetweenAnd is exited.
func (s *MySQLListenerGo) ExitExprBetweenAnd(ctx *ExprBetweenAndContext) {}

// EnterExprIsOrIsNotNull is called when production exprIsOrIsNotNull is entered.
func (s *MySQLListenerGo) EnterExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {}

// ExitExprIsOrIsNotNull is called when production exprIsOrIsNotNull is exited.
func (s *MySQLListenerGo) ExitExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {}

// EnterExprInSelect is called when production exprInSelect is entered.
func (s *MySQLListenerGo) EnterExprInSelect(ctx *ExprInSelectContext) {}

// ExitExprInSelect is called when production exprInSelect is exited.
func (s *MySQLListenerGo) ExitExprInSelect(ctx *ExprInSelectContext) {}

// EnterExprInValues is called when production exprInValues is entered.
func (s *MySQLListenerGo) EnterExprInValues(ctx *ExprInValuesContext) {}

// ExitExprInValues is called when production exprInValues is exited.
func (s *MySQLListenerGo) ExitExprInValues(ctx *ExprInValuesContext) {}

// EnterExprExists is called when production exprExists is entered.
func (s *MySQLListenerGo) EnterExprExists(ctx *ExprExistsContext) {}

// ExitExprExists is called when production exprExists is exited.
func (s *MySQLListenerGo) ExitExprExists(ctx *ExprExistsContext) {}

// EnterExprNot is called when production exprNot is entered.
func (s *MySQLListenerGo) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production exprNot is exited.
func (s *MySQLListenerGo) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprLike is called when production exprLike is entered.
func (s *MySQLListenerGo) EnterExprLike(ctx *ExprLikeContext) {}

// ExitExprLike is called when production exprLike is exited.
func (s *MySQLListenerGo) ExitExprLike(ctx *ExprLikeContext) {}

// EnterElement is called when production element is entered.
func (s *MySQLListenerGo) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *MySQLListenerGo) ExitElement(ctx *ElementContext) {}

// EnterElementOpFactory is called when production elementOpFactory is entered.
func (s *MySQLListenerGo) EnterElementOpFactory(ctx *ElementOpFactoryContext) {}

// ExitElementOpFactory is called when production elementOpFactory is exited.
func (s *MySQLListenerGo) ExitElementOpFactory(ctx *ElementOpFactoryContext) {}

// EnterElementText is called when production elementText is entered.
func (s *MySQLListenerGo) EnterElementText(ctx *ElementTextContext) {}

// ExitElementText is called when production elementText is exited.
func (s *MySQLListenerGo) ExitElementText(ctx *ElementTextContext) {}

// EnterElementTextParam is called when production elementTextParam is entered.
func (s *MySQLListenerGo) EnterElementTextParam(ctx *ElementTextParamContext) {}

// ExitElementTextParam is called when production elementTextParam is exited.
func (s *MySQLListenerGo) ExitElementTextParam(ctx *ElementTextParamContext) {}

// EnterElementDate is called when production elementDate is entered.
func (s *MySQLListenerGo) EnterElementDate(ctx *ElementDateContext) {}

// ExitElementDate is called when production elementDate is exited.
func (s *MySQLListenerGo) ExitElementDate(ctx *ElementDateContext) {}

// EnterElementSubQuery is called when production elementSubQuery is entered.
func (s *MySQLListenerGo) EnterElementSubQuery(ctx *ElementSubQueryContext) {}

// ExitElementSubQuery is called when production elementSubQuery is exited.
func (s *MySQLListenerGo) ExitElementSubQuery(ctx *ElementSubQueryContext) {}

// EnterElementWapperBkt is called when production elementWapperBkt is entered.
func (s *MySQLListenerGo) EnterElementWapperBkt(ctx *ElementWapperBktContext) {}

// ExitElementWapperBkt is called when production elementWapperBkt is exited.
func (s *MySQLListenerGo) ExitElementWapperBkt(ctx *ElementWapperBktContext) {}

// EnterElementListFactor is called when production elementListFactor is entered.
func (s *MySQLListenerGo) EnterElementListFactor(ctx *ElementListFactorContext) {}

// ExitElementListFactor is called when production elementListFactor is exited.
func (s *MySQLListenerGo) ExitElementListFactor(ctx *ElementListFactorContext) {}

// EnterElementList is called when production elementList is entered.
func (s *MySQLListenerGo) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *MySQLListenerGo) ExitElementList(ctx *ElementListContext) {}

// EnterElementOpEle is called when production elementOpEle is entered.
func (s *MySQLListenerGo) EnterElementOpEle(ctx *ElementOpEleContext) {}

// ExitElementOpEle is called when production elementOpEle is exited.
func (s *MySQLListenerGo) ExitElementOpEle(ctx *ElementOpEleContext) {}

// EnterElementOpEleSuffix is called when production elementOpEleSuffix is entered.
func (s *MySQLListenerGo) EnterElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {}

// ExitElementOpEleSuffix is called when production elementOpEleSuffix is exited.
func (s *MySQLListenerGo) ExitElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {}

// EnterElementCase is called when production elementCase is entered.
func (s *MySQLListenerGo) EnterElementCase(ctx *ElementCaseContext) {}

// ExitElementCase is called when production elementCase is exited.
func (s *MySQLListenerGo) ExitElementCase(ctx *ElementCaseContext) {}

// EnterCaseWhenPart is called when production caseWhenPart is entered.
func (s *MySQLListenerGo) EnterCaseWhenPart(ctx *CaseWhenPartContext) {}

// ExitCaseWhenPart is called when production caseWhenPart is exited.
func (s *MySQLListenerGo) ExitCaseWhenPart(ctx *CaseWhenPartContext) {}

// EnterElementWithPrefix is called when production elementWithPrefix is entered.
func (s *MySQLListenerGo) EnterElementWithPrefix(ctx *ElementWithPrefixContext) {}

// ExitElementWithPrefix is called when production elementWithPrefix is exited.
func (s *MySQLListenerGo) ExitElementWithPrefix(ctx *ElementWithPrefixContext) {}

// EnterElementRow is called when production elementRow is entered.
func (s *MySQLListenerGo) EnterElementRow(ctx *ElementRowContext) {}

// ExitElementRow is called when production elementRow is exited.
func (s *MySQLListenerGo) ExitElementRow(ctx *ElementRowContext) {}

// EnterFunCall is called when production funCall is entered.
func (s *MySQLListenerGo) EnterFunCall(ctx *FunCallContext) {}

// ExitFunCall is called when production funCall is exited.
func (s *MySQLListenerGo) ExitFunCall(ctx *FunCallContext) {}

// EnterParamList is called when production paramList is entered.
func (s *MySQLListenerGo) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *MySQLListenerGo) ExitParamList(ctx *ParamListContext) {}
