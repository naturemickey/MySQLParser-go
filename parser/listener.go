package parser

import (
	. "MySQLParser-go/parser/antlr4"
	"MySQLParser-go/parser/tree"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type MySQLListenerGo struct {
	stack *tree.StatStack
}

func (this *MySQLListenerGo) Result() tree.Stat {
	return this.stack.Head()
}

var _ MySQLListener = &MySQLListenerGo{}

func NewMySQLListener() *MySQLListenerGo {
	return &MySQLListenerGo{stack: tree.NewStatStack()}
}

// EnterEveryRule is called when any rule is entered.
func (this *MySQLListenerGo) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println(ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (this *MySQLListenerGo) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (this *MySQLListenerGo) EnterSelectStat(ctx *SelectStatContext) {
	this.stack.PushMateriel(&tree.SelectStat{})
}
func (this *MySQLListenerGo) ExitSelectStat(ctx *SelectStatContext) {
	this.stack.MakeProduct()
}

func (this *MySQLListenerGo) EnterSelectInner(ctx *SelectInnerContext) {
	this.stack.PushMateriel(&tree.SelectInner{})
}

func (this *MySQLListenerGo) ExitSelectInner(ctx *SelectInnerContext) {
	this.stack.MakeProduct()
}

// EnterSelectPrefix is called when production selectPrefix is entered.
func (this *MySQLListenerGo) EnterSelectPrefix(ctx *SelectPrefixContext) {
	selectPrefix := &tree.SelectPrefix{}
	selectPrefix.SetDistinct(ctx.GetDistinct() != nil)
	this.stack.PushMateriel(selectPrefix)
}

// ExitSelectPrefix is called when production selectPrefix is exited.
func (this *MySQLListenerGo) ExitSelectPrefix(ctx *SelectPrefixContext) {
	this.stack.MakeProduct()
}

// VisitTerminal is called when a terminal node is visited.
func (this *MySQLListenerGo) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (this *MySQLListenerGo) VisitErrorNode(node antlr.ErrorNode) {}

// EnterStat is called when production stat is entered.
func (this *MySQLListenerGo) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (this *MySQLListenerGo) ExitStat(ctx *StatContext) {}

// EnterTranscationStat is called when production transcationStat is entered.
func (this *MySQLListenerGo) EnterTranscationStat(ctx *TranscationStatContext) {}

// ExitTranscationStat is called when production transcationStat is exited.
func (this *MySQLListenerGo) ExitTranscationStat(ctx *TranscationStatContext) {}

// EnterCommit is called when production commit is entered.
func (this *MySQLListenerGo) EnterCommit(ctx *CommitContext) {
	this.stack.PushMateriel(&tree.Commit{})
}

// ExitCommit is called when production commit is exited.
func (this *MySQLListenerGo) ExitCommit(ctx *CommitContext) {
	this.stack.MakeProduct()
}

// EnterRollback is called when production rollback is entered.
func (this *MySQLListenerGo) EnterRollback(ctx *RollbackContext) {
	this.stack.PushMateriel(&tree.Rollback{})
}

// ExitRollback is called when production rollback is exited.
func (this *MySQLListenerGo) ExitRollback(ctx *RollbackContext) {
	this.stack.MakeProduct()
}

// EnterInsertStat is called when production insertStat is entered.
func (this *MySQLListenerGo) EnterInsertStat(ctx *InsertStatContext) {
	insertStat := &tree.InsertStat{}
	insertStat.SetTableName(ctx.GetTableName().GetText())
	this.stack.PushMateriel(insertStat)
}

// ExitInsertStat is called when production insertStat is exited.
func (this *MySQLListenerGo) ExitInsertStat(ctx *InsertStatContext) {
	this.stack.MakeProduct()
}

// EnterColumnNames is called when production columnNames is entered.
func (this *MySQLListenerGo) EnterColumnNames(ctx *ColumnNamesContext) {
	columnNames := &tree.ColumnNames{}
	var names []string
	for _, id := range ctx.AllID() {
		names = append(names, id.GetText())
	}
	columnNames.SetNames(names)
	this.stack.PushMateriel(columnNames)
}

// ExitColumnNames is called when production columnNames is exited.
func (this *MySQLListenerGo) ExitColumnNames(ctx *ColumnNamesContext) {
	this.stack.MakeProduct()
}

// EnterValueList is called when production valueList is entered.
func (this *MySQLListenerGo) EnterValueList(ctx *ValueListContext) {
	this.stack.PushMateriel(&tree.ValueList{})
}

// ExitValueList is called when production valueList is exited.
func (this *MySQLListenerGo) ExitValueList(ctx *ValueListContext) {
	this.stack.MakeProduct()
}

// EnterSelectSuffix is called when production selectSuffix is entered.
func (this *MySQLListenerGo) EnterSelectSuffix(ctx *SelectSuffixContext) {
	selectSuffix := &tree.SelectSuffix{}
	hasLIMIT := ctx.LIMIT() != nil
	selectSuffix.SetHasLimit(hasLIMIT)
	if hasLIMIT {
		if ctx.GetOffset() != nil {
			selectSuffix.SetOffset(ctx.GetOffset().GetText())
		}
		selectSuffix.SetRowCount(ctx.GetRowCount().GetText())
		selectSuffix.SetHasOffSet(ctx.OFFSET() != nil)
	}
	if ctx.GetLock() != nil {
		selectSuffix.SetLock(ctx.GetLock().GetText())
	}

	this.stack.PushMateriel(selectSuffix)
}

// ExitSelectSuffix is called when production selectSuffix is exited.
func (this *MySQLListenerGo) ExitSelectSuffix(ctx *SelectSuffixContext) {
	this.stack.MakeProduct()
}

// EnterSelectUnionSuffix is called when production selectUnionSuffix is entered.
func (this *MySQLListenerGo) EnterSelectUnionSuffix(ctx *SelectUnionSuffixContext) {
	selectUnionSuffix := &tree.SelectUnionSuffix{}
	selectUnionSuffix.SetMethod(ctx.GetMethod().GetText())

	this.stack.PushMateriel(selectUnionSuffix)
}

// ExitSelectUnionSuffix is called when production selectUnionSuffix is exited.
func (this *MySQLListenerGo) ExitSelectUnionSuffix(ctx *SelectUnionSuffixContext) {
	this.stack.MakeProduct()
}

// EnterSelectExprs is called when production selectExprs is entered.
func (this *MySQLListenerGo) EnterSelectExprs(ctx *SelectExprsContext) {
	selectExprs := &tree.SelectExprs{}
	if ctx.GetAlias() != nil {
		selectExprs.SetAlias(ctx.GetAlias().GetText())
	}
	this.stack.PushMateriel(selectExprs)
}

// ExitSelectExprs is called when production selectExprs is exited.
func (this *MySQLListenerGo) ExitSelectExprs(ctx *SelectExprsContext) {
	this.stack.MakeProduct()
}

// EnterTables is called when production tables is entered.
func (this *MySQLListenerGo) EnterTables(ctx *TablesContext) {
	this.stack.PushMateriel(&tree.Tables{})
}

// ExitTables is called when production tables is exited.
func (this *MySQLListenerGo) ExitTables(ctx *TablesContext) {
	this.stack.MakeProduct()
}

// EnterTableRel is called when production tableRel is entered.
func (this *MySQLListenerGo) EnterTableRel(ctx *TableRelContext) {}

// ExitTableRel is called when production tableRel is exited.
func (this *MySQLListenerGo) ExitTableRel(ctx *TableRelContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (this *MySQLListenerGo) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (this *MySQLListenerGo) ExitTableFactor(ctx *TableFactorContext) {}

// EnterTableSubQuery is called when production tableSubQuery is entered.
func (this *MySQLListenerGo) EnterTableSubQuery(ctx *TableSubQueryContext) {
	tableSubQuery := &tree.TableSubQuery{}
	tableSubQuery.SetAlias(ctx.GetAlias().GetText())
	this.stack.PushMateriel(tableSubQuery)
}

// ExitTableSubQuery is called when production tableSubQuery is exited.
func (this *MySQLListenerGo) ExitTableSubQuery(ctx *TableSubQueryContext) {
	this.stack.MakeProduct()
}

// EnterTableRecu is called when production tableRecu is entered.
func (this *MySQLListenerGo) EnterTableRecu(ctx *TableRecuContext) {
	this.stack.PushMateriel(&tree.TableRecu{})
}

// ExitTableRecu is called when production tableRecu is exited.
func (this *MySQLListenerGo) ExitTableRecu(ctx *TableRecuContext) {
	this.stack.MakeProduct()
}

// EnterTableJoin is called when production tableJoin is entered.
func (this *MySQLListenerGo) EnterTableJoin(ctx *TableJoinContext) {
	this.stack.PushMateriel(&tree.TableJoin{})
}

// ExitTableJoin is called when production tableJoin is exited.
func (this *MySQLListenerGo) ExitTableJoin(ctx *TableJoinContext) {
	this.stack.MakeProduct()
}

// EnterTableJoinSuffix is called when production tableJoinSuffix is entered.
func (this *MySQLListenerGo) EnterTableJoinSuffix(ctx *TableJoinSuffixContext) {
	this.stack.PushMateriel(&tree.TableJoinSuffix{})
}

// ExitTableJoinSuffix is called when production tableJoinSuffix is exited.
func (this *MySQLListenerGo) ExitTableJoinSuffix(ctx *TableJoinSuffixContext) {
	this.stack.MakeProduct()
}

// EnterTableJoinMod is called when production tableJoinMod is entered.
func (this *MySQLListenerGo) EnterTableJoinMod(ctx *TableJoinModContext) {
	tableJoinMod := &tree.TableJoinMod{}
	tableJoinMod.SetMod(ctx.GetText())
	this.stack.PushMateriel(tableJoinMod)
}

// ExitTableJoinMod is called when production tableJoinMod is exited.
func (this *MySQLListenerGo) ExitTableJoinMod(ctx *TableJoinModContext) {
	this.stack.MakeProduct()
}

// EnterJoinCondition is called when production joinCondition is entered.
func (this *MySQLListenerGo) EnterJoinCondition(ctx *JoinConditionContext) {
	this.stack.PushMateriel(&tree.JoinCondition{})
}

// ExitJoinCondition is called when production joinCondition is exited.
func (this *MySQLListenerGo) ExitJoinCondition(ctx *JoinConditionContext) {
	this.stack.MakeProduct()
}

// EnterGbobExprs is called when production gbobExprs is entered.
func (this *MySQLListenerGo) EnterGbobExprs(ctx *GbobExprsContext) {
	gbobExprs := &tree.GbobExprs{}
	if ctx.GetSc() != nil {
		gbobExprs.SetSc(ctx.GetSc().GetText())
	}
	this.stack.PushMateriel(gbobExprs)
}

// ExitGbobExprs is called when production gbobExprs is exited.
func (this *MySQLListenerGo) ExitGbobExprs(ctx *GbobExprsContext) {
	this.stack.MakeProduct()
}

// EnterUpdateStat is called when production updateStat is entered.
func (this *MySQLListenerGo) EnterUpdateStat(ctx *UpdateStatContext) {}

// ExitUpdateStat is called when production updateStat is exited.
func (this *MySQLListenerGo) ExitUpdateStat(ctx *UpdateStatContext) {}

// EnterUpdateSingleTable is called when production updateSingleTable is entered.
func (this *MySQLListenerGo) EnterUpdateSingleTable(ctx *UpdateSingleTableContext) {
	updateSingleTable := &tree.UpdateSingleTable{}
	if ctx.GetRowCount() != nil {
		updateSingleTable.SetRowCount(ctx.GetRowCount().GetText())
	}
	this.stack.PushMateriel(updateSingleTable)
}

// ExitUpdateSingleTable is called when production updateSingleTable is exited.
func (this *MySQLListenerGo) ExitUpdateSingleTable(ctx *UpdateSingleTableContext) {
	this.stack.MakeProduct()
}

// EnterUpdateMultipleTable is called when production updateMultipleTable is entered.
func (this *MySQLListenerGo) EnterUpdateMultipleTable(ctx *UpdateMultipleTableContext) {
	this.stack.PushMateriel(&tree.UpdateMultipleTable{})
}

// ExitUpdateMultipleTable is called when production updateMultipleTable is exited.
func (this *MySQLListenerGo) ExitUpdateMultipleTable(ctx *UpdateMultipleTableContext) {
	this.stack.MakeProduct()
}

// EnterSetExprs is called when production setExprs is entered.
func (this *MySQLListenerGo) EnterSetExprs(ctx *SetExprsContext) {
	this.stack.PushMateriel(&tree.SetExprs{})
}

// ExitSetExprs is called when production setExprs is exited.
func (this *MySQLListenerGo) ExitSetExprs(ctx *SetExprsContext) {
	this.stack.MakeProduct()
}

// EnterSetExpr is called when production setExpr is entered.
func (this *MySQLListenerGo) EnterSetExpr(ctx *SetExprContext) {
	setExpr := &tree.SetExpr{}
	if ctx.GetRightDefault() != nil {
		setExpr.SetRightDefault("default")
	}
	this.stack.PushMateriel(setExpr)
}

// ExitSetExpr is called when production setExpr is exited.
func (this *MySQLListenerGo) ExitSetExpr(ctx *SetExprContext) {
	this.stack.MakeProduct()
}

// EnterDeleteStat is called when production deleteStat is entered.
func (this *MySQLListenerGo) EnterDeleteStat(ctx *DeleteStatContext) {
	deleteStat := &tree.DeleteStat{}
	if ctx.GetRowCount() != nil {
		deleteStat.SetRowCount(ctx.GetRowCount().GetText())
	}
	this.stack.PushMateriel(deleteStat)
}

// ExitDeleteStat is called when production deleteStat is exited.
func (this *MySQLListenerGo) ExitDeleteStat(ctx *DeleteStatContext) {
	this.stack.MakeProduct()
}

// EnterTableNameAndAlias is called when production tableNameAndAlias is entered.
func (this *MySQLListenerGo) EnterTableNameAndAlias(ctx *TableNameAndAliasContext) {
	tableNameAndAlias := &tree.TableNameAndAlias{}
	tableNameAndAlias.SetName(ctx.GetName().GetText())
	if ctx.GetAlias() != nil {
		tableNameAndAlias.SetAlias(ctx.GetAlias().GetText())
	}
	this.stack.PushMateriel(tableNameAndAlias)
}

// ExitTableNameAndAlias is called when production tableNameAndAlias is exited.
func (this *MySQLListenerGo) ExitTableNameAndAlias(ctx *TableNameAndAliasContext) {
	this.stack.MakeProduct()
}

// EnterTableNameAndAliases is called when production tableNameAndAliases is entered.
func (this *MySQLListenerGo) EnterTableNameAndAliases(ctx *TableNameAndAliasesContext) {
	this.stack.PushMateriel(&tree.TableNameAndAliases{})
}

// ExitTableNameAndAliases is called when production tableNameAndAliases is exited.
func (this *MySQLListenerGo) ExitTableNameAndAliases(ctx *TableNameAndAliasesContext) {
	this.stack.MakeProduct()
}

// EnterWhereCondition is called when production whereCondition is entered.
func (this *MySQLListenerGo) EnterWhereCondition(ctx *WhereConditionContext) {}

// ExitWhereCondition is called when production whereCondition is exited.
func (this *MySQLListenerGo) ExitWhereCondition(ctx *WhereConditionContext) {}

// EnterWhereCondSub is called when production whereCondSub is entered.
func (this *MySQLListenerGo) EnterWhereCondSub(ctx *WhereCondSubContext) {
	whereCondSub := &tree.WhereCondSub{}
	if ctx.GetExpressionOperator() != nil {
		whereCondSub.SetExpressionOperator(ctx.GetExpressionOperator().GetText())
	}
	this.stack.PushMateriel(whereCondSub)
}

// ExitWhereCondSub is called when production whereCondSub is exited.
func (this *MySQLListenerGo) ExitWhereCondSub(ctx *WhereCondSubContext) {
	this.stack.MakeProduct()
}

// EnterWhereCondOp is called when production whereCondOp is entered.
func (this *MySQLListenerGo) EnterWhereCondOp(ctx *WhereCondOpContext) {
	whereCondOp := &tree.WhereCondOp{}
	if ctx.GetExpressionOperator() != nil {
		whereCondOp.SetExpressionOperator(ctx.GetExpressionOperator().GetText())
	}
	this.stack.PushMateriel(whereCondOp)
}

// ExitWhereCondOp is called when production whereCondOp is exited.
func (this *MySQLListenerGo) ExitWhereCondOp(ctx *WhereCondOpContext) {
	this.stack.MakeProduct()
}

// EnterExpression is called when production expression is entered.
func (this *MySQLListenerGo) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (this *MySQLListenerGo) ExitExpression(ctx *ExpressionContext) {}

// EnterExprRelational is called when production exprRelational is entered.
func (this *MySQLListenerGo) EnterExprRelational(ctx *ExprRelationalContext) {
	exprRelational := &tree.ExprRelational{}
	exprRelational.SetRelationalOp(ctx.GetRelationalOp().GetText())
	this.stack.PushMateriel(exprRelational)
}

// ExitExprRelational is called when production exprRelational is exited.
func (this *MySQLListenerGo) ExitExprRelational(ctx *ExprRelationalContext) {
	this.stack.MakeProduct()
}

// EnterExprBetweenAnd is called when production exprBetweenAnd is entered.
func (this *MySQLListenerGo) EnterExprBetweenAnd(ctx *ExprBetweenAndContext) {
	exprBetweenAnd := &tree.ExprBetweenAnd{}
	if ctx.GetNot() != nil {
		exprBetweenAnd.SetNot("not")
	}
	this.stack.PushMateriel(exprBetweenAnd)
}

// ExitExprBetweenAnd is called when production exprBetweenAnd is exited.
func (this *MySQLListenerGo) ExitExprBetweenAnd(ctx *ExprBetweenAndContext) {
	this.stack.MakeProduct()
}

// EnterExprIsOrIsNotNull is called when production exprIsOrIsNotNull is entered.
func (this *MySQLListenerGo) EnterExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {
	exprIsOrIsNotNull := &tree.ExprIsOrIsNotNull{}
	if ctx.GetNot() != nil {
		exprIsOrIsNotNull.SetNot("not")
	}
	exprIsOrIsNotNull.SetWhat(ctx.GetWhat().GetText())
	this.stack.PushMateriel(exprIsOrIsNotNull)
}

// ExitExprIsOrIsNotNull is called when production exprIsOrIsNotNull is exited.
func (this *MySQLListenerGo) ExitExprIsOrIsNotNull(ctx *ExprIsOrIsNotNullContext) {
	this.stack.MakeProduct()
}

// EnterExprInSelect is called when production exprInSelect is entered.
func (this *MySQLListenerGo) EnterExprInSelect(ctx *ExprInSelectContext) {
	exprInSelect := &tree.ExprInSelect{}
	if ctx.GetNot() != nil {
		exprInSelect.SetNot("not")
	}
	this.stack.PushMateriel(exprInSelect)
}

// ExitExprInSelect is called when production exprInSelect is exited.
func (this *MySQLListenerGo) ExitExprInSelect(ctx *ExprInSelectContext) {
	this.stack.MakeProduct()
}

// EnterExprInValues is called when production exprInValues is entered.
func (this *MySQLListenerGo) EnterExprInValues(ctx *ExprInValuesContext) {
	exprInValues := &tree.ExprInValues{}
	if ctx.GetNot() != nil {
		exprInValues.SetNot("not")
	}
	this.stack.PushMateriel(exprInValues)
}

// ExitExprInValues is called when production exprInValues is exited.
func (this *MySQLListenerGo) ExitExprInValues(ctx *ExprInValuesContext) {
	this.stack.MakeProduct()
}

// EnterExprExists is called when production exprExists is entered.
func (this *MySQLListenerGo) EnterExprExists(ctx *ExprExistsContext) {
	exprExists := &tree.ExprExists{}
	if ctx.GetNot() != nil {
		exprExists.SetNot("not")
	}
	this.stack.PushMateriel(exprExists)
}

// ExitExprExists is called when production exprExists is exited.
func (this *MySQLListenerGo) ExitExprExists(ctx *ExprExistsContext) {
	this.stack.MakeProduct()
}

// EnterExprNot is called when production exprNot is entered.
func (this *MySQLListenerGo) EnterExprNot(ctx *ExprNotContext) {
	this.stack.PushMateriel(&tree.ExprNot{})
}

// ExitExprNot is called when production exprNot is exited.
func (this *MySQLListenerGo) ExitExprNot(ctx *ExprNotContext) {
	this.stack.MakeProduct()
}

// EnterExprLike is called when production exprLike is entered.
func (this *MySQLListenerGo) EnterExprLike(ctx *ExprLikeContext) {
	exprLike := &tree.ExprLike{}
	if ctx.GetNot() != nil {
		exprLike.SetNot("not")
	}
	this.stack.PushMateriel(exprLike)
}

// ExitExprLike is called when production exprLike is exited.
func (this *MySQLListenerGo) ExitExprLike(ctx *ExprLikeContext) {
	this.stack.MakeProduct()
}

// EnterElement is called when production element is entered.
func (this *MySQLListenerGo) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (this *MySQLListenerGo) ExitElement(ctx *ElementContext) {}

// EnterElementOpFactory is called when production elementOpFactory is entered.
func (this *MySQLListenerGo) EnterElementOpFactory(ctx *ElementOpFactoryContext) {}

// ExitElementOpFactory is called when production elementOpFactory is exited.
func (this *MySQLListenerGo) ExitElementOpFactory(ctx *ElementOpFactoryContext) {}

// EnterElementText is called when production elementText is entered.
func (this *MySQLListenerGo) EnterElementText(ctx *ElementTextContext) {
	elementText := &tree.ElementText{}
	elementText.SetText(ctx.GetText())
	this.stack.PushMateriel(elementText)
}

// ExitElementText is called when production elementText is exited.
func (this *MySQLListenerGo) ExitElementText(ctx *ElementTextContext) {
	this.stack.MakeProduct()
}

// EnterElementTextParam is called when production elementTextParam is entered.
func (this *MySQLListenerGo) EnterElementTextParam(ctx *ElementTextParamContext) {
	elementTextParam := &tree.ElementTextParam{}
	elementTextParam.SetText(ctx.GetText())
	this.stack.PushMateriel(elementTextParam)
}

// ExitElementTextParam is called when production elementTextParam is exited.
func (this *MySQLListenerGo) ExitElementTextParam(ctx *ElementTextParamContext) {
	this.stack.MakeProduct()
}

// EnterElementDate is called when production elementDate is entered.
func (this *MySQLListenerGo) EnterElementDate(ctx *ElementDateContext) {
	elementDate := &tree.ElementDate{}
	elementDate.SetDt(ctx.GetDt().GetText())
	elementDate.SetStr(ctx.STRING().GetText())
	this.stack.PushMateriel(elementDate)
}

// ExitElementDate is called when production elementDate is exited.
func (this *MySQLListenerGo) ExitElementDate(ctx *ElementDateContext) {
	this.stack.MakeProduct()
}

// EnterElementSubQuery is called when production elementSubQuery is entered.
func (this *MySQLListenerGo) EnterElementSubQuery(ctx *ElementSubQueryContext) {
	elementSubQuery := &tree.ElementSubQuery{}
	if ctx.GetSqWith() != nil {
		elementSubQuery.SetSqWith(ctx.GetSqWith().GetText())
	}
	this.stack.PushMateriel(elementSubQuery)
}

// ExitElementSubQuery is called when production elementSubQuery is exited.
func (this *MySQLListenerGo) ExitElementSubQuery(ctx *ElementSubQueryContext) {
	this.stack.MakeProduct()
}

// EnterElementWapperBkt is called when production elementWapperBkt is entered.
func (this *MySQLListenerGo) EnterElementWapperBkt(ctx *ElementWapperBktContext) {
	this.stack.PushMateriel(&tree.ElementWapperBkt{})
}

// ExitElementWapperBkt is called when production elementWapperBkt is exited.
func (this *MySQLListenerGo) ExitElementWapperBkt(ctx *ElementWapperBktContext) {
	this.stack.MakeProduct()
}

// EnterElementListFactor is called when production elementListFactor is entered.
func (this *MySQLListenerGo) EnterElementListFactor(ctx *ElementListFactorContext) {
	this.stack.PushMateriel(&tree.ElementListFactor{})
}

// ExitElementListFactor is called when production elementListFactor is exited.
func (this *MySQLListenerGo) ExitElementListFactor(ctx *ElementListFactorContext) {
	this.stack.MakeProduct()
}

// EnterElementList is called when production elementList is entered.
func (this *MySQLListenerGo) EnterElementList(ctx *ElementListContext) {
	this.stack.PushMateriel(&tree.ElementList{})
}

// ExitElementList is called when production elementList is exited.
func (this *MySQLListenerGo) ExitElementList(ctx *ElementListContext) {
	this.stack.MakeProduct()
}

// EnterElementOpEle is called when production elementOpEle is entered.
func (this *MySQLListenerGo) EnterElementOpEle(ctx *ElementOpEleContext) {
	this.stack.PushMateriel(&tree.ElementOpEle{})
}

// ExitElementOpEle is called when production elementOpEle is exited.
func (this *MySQLListenerGo) ExitElementOpEle(ctx *ElementOpEleContext) {
	this.stack.MakeProduct()
}

// EnterElementOpEleSuffix is called when production elementOpEleSuffix is entered.
func (this *MySQLListenerGo) EnterElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {
	elementOpEleSuffix := &tree.ElementOpEleSuffix{}
	if ctx.GetOp() != nil {
		elementOpEleSuffix.SetOp(ctx.GetOp().GetText())
	}
	this.stack.PushMateriel(elementOpEleSuffix)
}

// ExitElementOpEleSuffix is called when production elementOpEleSuffix is exited.
func (this *MySQLListenerGo) ExitElementOpEleSuffix(ctx *ElementOpEleSuffixContext) {
	this.stack.MakeProduct()
}

// EnterElementCase is called when production elementCase is entered.
func (this *MySQLListenerGo) EnterElementCase(ctx *ElementCaseContext) {
	this.stack.PushMateriel(&tree.ElementCase{})
}

// ExitElementCase is called when production elementCase is exited.
func (this *MySQLListenerGo) ExitElementCase(ctx *ElementCaseContext) {
	this.stack.MakeProduct()
}

// EnterCaseWhenPart is called when production caseWhenPart is entered.
func (this *MySQLListenerGo) EnterCaseWhenPart(ctx *CaseWhenPartContext) {
	this.stack.PushMateriel(&tree.CaseWhenPart{})
}

// ExitCaseWhenPart is called when production caseWhenPart is exited.
func (this *MySQLListenerGo) ExitCaseWhenPart(ctx *CaseWhenPartContext) {
	this.stack.MakeProduct()
}

// EnterElementWithPrefix is called when production elementWithPrefix is entered.
func (this *MySQLListenerGo) EnterElementWithPrefix(ctx *ElementWithPrefixContext) {
	elementWithPrefix := &tree.ElementWithPrefix{}
	elementWithPrefix.SetPrefix("binary")
	this.stack.PushMateriel(elementWithPrefix)
}

// ExitElementWithPrefix is called when production elementWithPrefix is exited.
func (this *MySQLListenerGo) ExitElementWithPrefix(ctx *ElementWithPrefixContext) {
	this.stack.MakeProduct()
}

// EnterElementRow is called when production elementRow is entered.
func (this *MySQLListenerGo) EnterElementRow(ctx *ElementRowContext) {
	this.stack.PushMateriel(&tree.ElementRow{})
}

// ExitElementRow is called when production elementRow is exited.
func (this *MySQLListenerGo) ExitElementRow(ctx *ElementRowContext) {
	this.stack.MakeProduct()
}

// EnterFunCall is called when production funCall is entered.
func (this *MySQLListenerGo) EnterFunCall(ctx *FunCallContext) {
	funCall := &tree.FunCall{}
	funCall.SetFunName(ctx.GetFunName().GetText())
	this.stack.PushMateriel(funCall)
}

// ExitFunCall is called when production funCall is exited.
func (this *MySQLListenerGo) ExitFunCall(ctx *FunCallContext) {
	this.stack.MakeProduct()
}

// EnterParamList is called when production paramList is entered.
func (this *MySQLListenerGo) EnterParamList(ctx *ParamListContext) {
	this.stack.PushMateriel(&tree.ParamList{})
}

// ExitParamList is called when production paramList is exited.
func (this *MySQLListenerGo) ExitParamList(ctx *ParamListContext) {
	this.stack.MakeProduct()
}
