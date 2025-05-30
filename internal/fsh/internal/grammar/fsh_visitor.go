// Code generated from FSH.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // FSH
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FSHParser.
type FSHVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FSHParser#doc.
	VisitDoc(ctx *DocContext) interface{}

	// Visit a parse tree produced by FSHParser#entity.
	VisitEntity(ctx *EntityContext) interface{}

	// Visit a parse tree produced by FSHParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by FSHParser#profile.
	VisitProfile(ctx *ProfileContext) interface{}

	// Visit a parse tree produced by FSHParser#extension.
	VisitExtension(ctx *ExtensionContext) interface{}

	// Visit a parse tree produced by FSHParser#logical.
	VisitLogical(ctx *LogicalContext) interface{}

	// Visit a parse tree produced by FSHParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by FSHParser#sdMetadata.
	VisitSdMetadata(ctx *SdMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#sdRule.
	VisitSdRule(ctx *SdRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#lrRule.
	VisitLrRule(ctx *LrRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#instance.
	VisitInstance(ctx *InstanceContext) interface{}

	// Visit a parse tree produced by FSHParser#instanceMetadata.
	VisitInstanceMetadata(ctx *InstanceMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#instanceRule.
	VisitInstanceRule(ctx *InstanceRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#invariant.
	VisitInvariant(ctx *InvariantContext) interface{}

	// Visit a parse tree produced by FSHParser#invariantMetadata.
	VisitInvariantMetadata(ctx *InvariantMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#invariantRule.
	VisitInvariantRule(ctx *InvariantRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#valueSet.
	VisitValueSet(ctx *ValueSetContext) interface{}

	// Visit a parse tree produced by FSHParser#vsMetadata.
	VisitVsMetadata(ctx *VsMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#vsRule.
	VisitVsRule(ctx *VsRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#codeSystem.
	VisitCodeSystem(ctx *CodeSystemContext) interface{}

	// Visit a parse tree produced by FSHParser#csMetadata.
	VisitCsMetadata(ctx *CsMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#csRule.
	VisitCsRule(ctx *CsRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#ruleSet.
	VisitRuleSet(ctx *RuleSetContext) interface{}

	// Visit a parse tree produced by FSHParser#ruleSetRule.
	VisitRuleSetRule(ctx *RuleSetRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#paramRuleSet.
	VisitParamRuleSet(ctx *ParamRuleSetContext) interface{}

	// Visit a parse tree produced by FSHParser#paramRuleSetRef.
	VisitParamRuleSetRef(ctx *ParamRuleSetRefContext) interface{}

	// Visit a parse tree produced by FSHParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by FSHParser#lastParameter.
	VisitLastParameter(ctx *LastParameterContext) interface{}

	// Visit a parse tree produced by FSHParser#paramRuleSetContent.
	VisitParamRuleSetContent(ctx *ParamRuleSetContentContext) interface{}

	// Visit a parse tree produced by FSHParser#mapping.
	VisitMapping(ctx *MappingContext) interface{}

	// Visit a parse tree produced by FSHParser#mappingMetadata.
	VisitMappingMetadata(ctx *MappingMetadataContext) interface{}

	// Visit a parse tree produced by FSHParser#mappingEntityRule.
	VisitMappingEntityRule(ctx *MappingEntityRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#parent.
	VisitParent(ctx *ParentContext) interface{}

	// Visit a parse tree produced by FSHParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by FSHParser#title.
	VisitTitle(ctx *TitleContext) interface{}

	// Visit a parse tree produced by FSHParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by FSHParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by FSHParser#xpath.
	VisitXpath(ctx *XpathContext) interface{}

	// Visit a parse tree produced by FSHParser#severity.
	VisitSeverity(ctx *SeverityContext) interface{}

	// Visit a parse tree produced by FSHParser#instanceOf.
	VisitInstanceOf(ctx *InstanceOfContext) interface{}

	// Visit a parse tree produced by FSHParser#usage.
	VisitUsage(ctx *UsageContext) interface{}

	// Visit a parse tree produced by FSHParser#source.
	VisitSource(ctx *SourceContext) interface{}

	// Visit a parse tree produced by FSHParser#target.
	VisitTarget(ctx *TargetContext) interface{}

	// Visit a parse tree produced by FSHParser#context.
	VisitContext(ctx *ContextContext) interface{}

	// Visit a parse tree produced by FSHParser#contextItem.
	VisitContextItem(ctx *ContextItemContext) interface{}

	// Visit a parse tree produced by FSHParser#lastContextItem.
	VisitLastContextItem(ctx *LastContextItemContext) interface{}

	// Visit a parse tree produced by FSHParser#characteristics.
	VisitCharacteristics(ctx *CharacteristicsContext) interface{}

	// Visit a parse tree produced by FSHParser#cardRule.
	VisitCardRule(ctx *CardRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#flagRule.
	VisitFlagRule(ctx *FlagRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#valueSetRule.
	VisitValueSetRule(ctx *ValueSetRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#fixedValueRule.
	VisitFixedValueRule(ctx *FixedValueRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#containsRule.
	VisitContainsRule(ctx *ContainsRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#onlyRule.
	VisitOnlyRule(ctx *OnlyRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#obeysRule.
	VisitObeysRule(ctx *ObeysRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#caretValueRule.
	VisitCaretValueRule(ctx *CaretValueRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#codeCaretValueRule.
	VisitCodeCaretValueRule(ctx *CodeCaretValueRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#mappingRule.
	VisitMappingRule(ctx *MappingRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#insertRule.
	VisitInsertRule(ctx *InsertRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#codeInsertRule.
	VisitCodeInsertRule(ctx *CodeInsertRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#addCRElementRule.
	VisitAddCRElementRule(ctx *AddCRElementRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#addElementRule.
	VisitAddElementRule(ctx *AddElementRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#pathRule.
	VisitPathRule(ctx *PathRuleContext) interface{}

	// Visit a parse tree produced by FSHParser#vsComponent.
	VisitVsComponent(ctx *VsComponentContext) interface{}

	// Visit a parse tree produced by FSHParser#vsConceptComponent.
	VisitVsConceptComponent(ctx *VsConceptComponentContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFilterComponent.
	VisitVsFilterComponent(ctx *VsFilterComponentContext) interface{}

	// Visit a parse tree produced by FSHParser#vsComponentFrom.
	VisitVsComponentFrom(ctx *VsComponentFromContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFromSystem.
	VisitVsFromSystem(ctx *VsFromSystemContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFromValueset.
	VisitVsFromValueset(ctx *VsFromValuesetContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFilterList.
	VisitVsFilterList(ctx *VsFilterListContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFilterDefinition.
	VisitVsFilterDefinition(ctx *VsFilterDefinitionContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFilterOperator.
	VisitVsFilterOperator(ctx *VsFilterOperatorContext) interface{}

	// Visit a parse tree produced by FSHParser#vsFilterValue.
	VisitVsFilterValue(ctx *VsFilterValueContext) interface{}

	// Visit a parse tree produced by FSHParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by FSHParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by FSHParser#caretPath.
	VisitCaretPath(ctx *CaretPathContext) interface{}

	// Visit a parse tree produced by FSHParser#flag.
	VisitFlag(ctx *FlagContext) interface{}

	// Visit a parse tree produced by FSHParser#strength.
	VisitStrength(ctx *StrengthContext) interface{}

	// Visit a parse tree produced by FSHParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FSHParser#item.
	VisitItem(ctx *ItemContext) interface{}

	// Visit a parse tree produced by FSHParser#code.
	VisitCode(ctx *CodeContext) interface{}

	// Visit a parse tree produced by FSHParser#concept.
	VisitConcept(ctx *ConceptContext) interface{}

	// Visit a parse tree produced by FSHParser#quantity.
	VisitQuantity(ctx *QuantityContext) interface{}

	// Visit a parse tree produced by FSHParser#ratio.
	VisitRatio(ctx *RatioContext) interface{}

	// Visit a parse tree produced by FSHParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by FSHParser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by FSHParser#codeableReferenceType.
	VisitCodeableReferenceType(ctx *CodeableReferenceTypeContext) interface{}

	// Visit a parse tree produced by FSHParser#canonical.
	VisitCanonical(ctx *CanonicalContext) interface{}

	// Visit a parse tree produced by FSHParser#ratioPart.
	VisitRatioPart(ctx *RatioPartContext) interface{}

	// Visit a parse tree produced by FSHParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by FSHParser#targetType.
	VisitTargetType(ctx *TargetTypeContext) interface{}

	// Visit a parse tree produced by FSHParser#mostAlphaKeywords.
	VisitMostAlphaKeywords(ctx *MostAlphaKeywordsContext) interface{}
}
