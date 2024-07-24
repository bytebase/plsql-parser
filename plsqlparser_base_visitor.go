// Code generated from PlSqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PlSqlParser
import "github.com/antlr4-go/antlr/v4"

type BasePlSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlSqlParserVisitor) VisitSql_script(ctx *Sql_scriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnit_statement(ctx *Unit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_diskgroup(ctx *Alter_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_disk_clause(ctx *Add_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_disk_clause(ctx *Drop_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitResize_disk_clause(ctx *Resize_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReplace_disk_clause(ctx *Replace_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWait_nowait(ctx *Wait_nowaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRename_disk_clause(ctx *Rename_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisk_online_clause(ctx *Disk_online_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisk_offline_clause(ctx *Disk_offline_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTimeout_clause(ctx *Timeout_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRebalance_diskgroup_clause(ctx *Rebalance_diskgroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPhase(ctx *PhaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCheck_diskgroup_clause(ctx *Check_diskgroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_template_clauses(ctx *Diskgroup_template_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQualified_template_clause(ctx *Qualified_template_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRedundancy_clause(ctx *Redundancy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStriping_clause(ctx *Striping_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitForce_noforce(ctx *Force_noforceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_directory_clauses(ctx *Diskgroup_directory_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDir_name(ctx *Dir_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_alias_clauses(ctx *Diskgroup_alias_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_volume_clauses(ctx *Diskgroup_volume_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_volume_clause(ctx *Add_volume_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_volume_clause(ctx *Modify_volume_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_attributes(ctx *Diskgroup_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_diskgroup_file(ctx *Modify_diskgroup_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisk_region_clause(ctx *Disk_region_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_diskgroup_file_clause(ctx *Drop_diskgroup_file_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConvert_redundancy_clause(ctx *Convert_redundancy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsergroup_clauses(ctx *Usergroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_clauses(ctx *User_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_permissions_clause(ctx *File_permissions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_owner_clause(ctx *File_owner_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitScrub_clause(ctx *Scrub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuotagroup_clauses(ctx *Quotagroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProperty_name(ctx *Property_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProperty_value(ctx *Property_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFilegroup_clauses(ctx *Filegroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_filegroup_clause(ctx *Add_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_filegroup_clause(ctx *Modify_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMove_to_filegroup_clause(ctx *Move_to_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_filegroup_clause(ctx *Drop_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuorum_regular(ctx *Quorum_regularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUndrop_disk_clause(ctx *Undrop_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDiskgroup_availability(ctx *Diskgroup_availabilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnable_disable_volume(ctx *Enable_disable_volumeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_function(ctx *Drop_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_flashback_archive(ctx *Alter_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_hierarchy(ctx *Alter_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_function(ctx *Alter_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_java(ctx *Alter_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMatch_string(ctx *Match_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_function_body(ctx *Create_function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSql_macro_body(ctx *Sql_macro_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParallel_enable_clause(ctx *Parallel_enable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRelies_on_part(ctx *Relies_on_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStreaming_clause(ctx *Streaming_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_outline(ctx *Alter_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOutline_options(ctx *Outline_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_lockdown_profile(ctx *Alter_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLockdown_feature(ctx *Lockdown_featureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLockdown_options(ctx *Lockdown_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLockdown_statements(ctx *Lockdown_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStatement_clauses(ctx *Statement_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClause_options(ctx *Clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOption_values(ctx *Option_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitString_list(ctx *String_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisable_enable(ctx *Disable_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_lockdown_profile(ctx *Drop_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_package(ctx *Drop_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_package(ctx *Alter_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_package(ctx *Create_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_package_body(ctx *Create_package_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPackage_obj_spec(ctx *Package_obj_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProcedure_spec(ctx *Procedure_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_spec(ctx *Function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPackage_obj_body(ctx *Package_obj_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_pmem_filestore(ctx *Alter_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_pmem_filestore(ctx *Drop_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_procedure(ctx *Drop_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_procedure(ctx *Alter_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_body(ctx *Function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProcedure_body(ctx *Procedure_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_procedure_body(ctx *Create_procedure_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_resource_cost(ctx *Alter_resource_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_outline(ctx *Drop_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_rollback_segment(ctx *Alter_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_restore_point(ctx *Drop_restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_rollback_segment(ctx *Drop_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_role(ctx *Drop_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_pmem_filestore(ctx *Create_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPmem_filestore_options(ctx *Pmem_filestore_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_path(ctx *File_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_rollback_segment(ctx *Create_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_trigger(ctx *Drop_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_trigger(ctx *Alter_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_trigger(ctx *Create_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrigger_follows_clause(ctx *Trigger_follows_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrigger_when_clause(ctx *Trigger_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSimple_dml_trigger(ctx *Simple_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFor_each_row(ctx *For_each_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompound_dml_trigger(ctx *Compound_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNon_dml_trigger(ctx *Non_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrigger_body(ctx *Trigger_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRoutine_clause(ctx *Routine_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompound_trigger_block(ctx *Compound_trigger_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTiming_point_section(ctx *Timing_point_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNon_dml_event(ctx *Non_dml_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDml_event_clause(ctx *Dml_event_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDml_event_element(ctx *Dml_event_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDml_event_nested_clause(ctx *Dml_event_nested_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReferencing_clause(ctx *Referencing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReferencing_element(ctx *Referencing_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_type(ctx *Drop_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_type(ctx *Alter_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompile_type_clause(ctx *Compile_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReplace_type_clause(ctx *Replace_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_method_spec(ctx *Alter_method_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_method_element(ctx *Alter_method_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_attribute_definition(ctx *Alter_attribute_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAttribute_definition(ctx *Attribute_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_collection_clauses(ctx *Alter_collection_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDependent_handling_clause(ctx *Dependent_handling_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDependent_exceptions_part(ctx *Dependent_exceptions_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_type(ctx *Create_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_definition(ctx *Type_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_type_def(ctx *Object_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_as_part(ctx *Object_as_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_under_part(ctx *Object_under_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNested_table_type_def(ctx *Nested_table_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSqlj_object_type(ctx *Sqlj_object_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_body(ctx *Type_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_body_elements(ctx *Type_body_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMap_order_func_declaration(ctx *Map_order_func_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubprog_decl_in_type(ctx *Subprog_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProc_decl_in_type(ctx *Proc_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunc_decl_in_type(ctx *Func_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstructor_declaration(ctx *Constructor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModifier_clause(ctx *Modifier_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_member_spec(ctx *Object_member_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSqlj_object_type_attr(ctx *Sqlj_object_type_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitElement_spec(ctx *Element_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitElement_spec_options(ctx *Element_spec_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubprogram_spec(ctx *Subprogram_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOverriding_subprogram_spec(ctx *Overriding_subprogram_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOverriding_function_spec(ctx *Overriding_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_procedure_spec(ctx *Type_procedure_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_function_spec(ctx *Type_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstructor_spec(ctx *Constructor_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMap_order_function_spec(ctx *Map_order_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPragma_clause(ctx *Pragma_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPragma_elements(ctx *Pragma_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_elements_parameter(ctx *Type_elements_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_sequence(ctx *Drop_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_sequence(ctx *Alter_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_session(ctx *Alter_sessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_session_set_clause(ctx *Alter_session_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_sequence(ctx *Create_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSequence_spec(ctx *Sequence_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSequence_start_clause(ctx *Sequence_start_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_analytic_view(ctx *Create_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClassification_clause(ctx *Classification_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCaption_clause(ctx *Caption_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDescription_clause(ctx *Description_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClassification_item(ctx *Classification_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLanguage(ctx *LanguageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCav_using_clause(ctx *Cav_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDim_by_clause(ctx *Dim_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDim_key(ctx *Dim_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDim_ref(ctx *Dim_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHier_ref(ctx *Hier_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMeasures_clause(ctx *Measures_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAv_measure(ctx *Av_measureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBase_meas_clause(ctx *Base_meas_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMeas_aggregate_clause(ctx *Meas_aggregate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCalc_meas_clause(ctx *Calc_meas_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_measure_clause(ctx *Default_measure_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_aggregate_clause(ctx *Default_aggregate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCache_clause(ctx *Cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCache_specification(ctx *Cache_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevels_clause(ctx *Levels_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevel_specification(ctx *Level_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevel_group_type(ctx *Level_group_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFact_columns_clause(ctx *Fact_columns_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQry_transform_clause(ctx *Qry_transform_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_attribute_dimension(ctx *Create_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAd_using_clause(ctx *Ad_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSource_clause(ctx *Source_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_path_clause(ctx *Join_path_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_condition(ctx *Join_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_condition_item(ctx *Join_condition_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAttributes_clause(ctx *Attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAd_attributes_clause(ctx *Ad_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAd_level_clause(ctx *Ad_level_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKey_clause(ctx *Key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlternate_key_clause(ctx *Alternate_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDim_order_clause(ctx *Dim_order_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAll_clause(ctx *All_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_audit_policy(ctx *Create_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrivilege_audit_clause(ctx *Privilege_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAction_audit_clause(ctx *Action_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStandard_actions(ctx *Standard_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitActions_clause(ctx *Actions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_action(ctx *Object_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSystem_action(ctx *System_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComponent_actions(ctx *Component_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComponent_action(ctx *Component_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRole_audit_clause(ctx *Role_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_controlfile(ctx *Create_controlfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitControlfile_options(ctx *Controlfile_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogfile_clause(ctx *Logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCharacter_set_clause(ctx *Character_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_specification(ctx *File_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_diskgroup(ctx *Create_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQualified_disk_clause(ctx *Qualified_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_edition(ctx *Create_editionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_flashback_archive(ctx *Create_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_archive_quota(ctx *Flashback_archive_quotaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_archive_retention(ctx *Flashback_archive_retentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_hierarchy(ctx *Create_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHier_using_clause(ctx *Hier_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevel_hier_clause(ctx *Level_hier_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHier_attrs_clause(ctx *Hier_attrs_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHier_attr_clause(ctx *Hier_attr_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHier_attr_name(ctx *Hier_attr_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_index(ctx *Create_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCluster_index_clause(ctx *Cluster_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCluster_name(ctx *Cluster_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_index_clause(ctx *Table_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBitmap_join_index_clause(ctx *Bitmap_join_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_expr_option(ctx *Index_expr_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_expr(ctx *Index_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_properties(ctx *Index_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDomain_index_clause(ctx *Domain_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLocal_domain_index_clause(ctx *Local_domain_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlindex_clause(ctx *Xmlindex_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLocal_xmlindex_clause(ctx *Local_xmlindex_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGlobal_partitioned_index(ctx *Global_partitioned_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_partitioning_clause(ctx *Index_partitioning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLocal_partitioned_index(ctx *Local_partitioned_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_range_partitioned_table(ctx *On_range_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_list_partitioned_table(ctx *On_list_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartitioned_table(ctx *Partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_hash_partitioned_table(ctx *On_hash_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_hash_partitioned_clause(ctx *On_hash_partitioned_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_comp_partitioned_table(ctx *On_comp_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_comp_partitioned_clause(ctx *On_comp_partitioned_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_subpartition_clause(ctx *Index_subpartition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_subpartition_subclause(ctx *Index_subpartition_subclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOdci_parameters(ctx *Odci_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndextype(ctx *IndextypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_index(ctx *Alter_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_index_ops_set1(ctx *Alter_index_ops_set1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_index_ops_set2(ctx *Alter_index_ops_set2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVisible_or_invisible(ctx *Visible_or_invisibleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMonitoring_nomonitoring(ctx *Monitoring_nomonitoringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRebuild_clause(ctx *Rebuild_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_index_partitioning(ctx *Alter_index_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_index_default_attrs(ctx *Modify_index_default_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_hash_index_partition(ctx *Add_hash_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCoalesce_index_partition(ctx *Coalesce_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_index_partition(ctx *Modify_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_index_partitions_ops(ctx *Modify_index_partitions_opsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRename_index_partition(ctx *Rename_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_index_partition(ctx *Drop_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSplit_index_partition(ctx *Split_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_partition_description(ctx *Index_partition_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_index_subpartition(ctx *Modify_index_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_name_old(ctx *Partition_name_oldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_partition_name(ctx *New_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_index_name(ctx *New_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_inmemory_join_group(ctx *Alter_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_user(ctx *Create_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_user(ctx *Alter_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_user(ctx *Drop_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_identified_by(ctx *Alter_identified_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentified_by(ctx *Identified_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentified_other_clause(ctx *Identified_other_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_tablespace_clause(ctx *User_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuota_clause(ctx *Quota_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProfile_clause(ctx *Profile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRole_clause(ctx *Role_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_default_role_clause(ctx *User_default_role_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPassword_expire_clause(ctx *Password_expire_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_lock_clause(ctx *User_lock_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_editions_clause(ctx *User_editions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_user_editions_clause(ctx *Alter_user_editions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProxy_clause(ctx *Proxy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitContainer_names(ctx *Container_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_container_data(ctx *Set_container_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_rem_container_data(ctx *Add_rem_container_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitContainer_data_clause(ctx *Container_data_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdminister_key_management(ctx *Administer_key_managementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKeystore_management_clauses(ctx *Keystore_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_keystore(ctx *Create_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOpen_keystore(ctx *Open_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitForce_keystore(ctx *Force_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClose_keystore(ctx *Close_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBackup_keystore(ctx *Backup_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_keystore_password(ctx *Alter_keystore_passwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_into_new_keystore(ctx *Merge_into_new_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_into_existing_keystore(ctx *Merge_into_existing_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIsolate_keystore(ctx *Isolate_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnite_keystore(ctx *Unite_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKey_management_clauses(ctx *Key_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_key(ctx *Set_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_key(ctx *Create_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMkid(ctx *MkidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMk(ctx *MkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUse_key(ctx *Use_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_key_tag(ctx *Set_key_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExport_keys(ctx *Export_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImport_keys(ctx *Import_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMigrate_keys(ctx *Migrate_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReverse_migrate_keys(ctx *Reverse_migrate_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMove_keys(ctx *Move_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentified_by_store(ctx *Identified_by_storeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_algorithm_clause(ctx *Using_algorithm_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_tag_clause(ctx *Using_tag_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSecret_management_clauses(ctx *Secret_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_update_secret(ctx *Add_update_secretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDelete_secret(ctx *Delete_secretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_update_secret_seps(ctx *Add_update_secret_sepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDelete_secret_seps(ctx *Delete_secret_sepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitZero_downtime_software_patching_clauses(ctx *Zero_downtime_software_patching_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWith_backup_clause(ctx *With_backup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentified_by_password_clause(ctx *Identified_by_password_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKeystore_password(ctx *Keystore_passwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSecret(ctx *SecretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAnalyze(ctx *AnalyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_extention_clause(ctx *Partition_extention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitValidation_clauses(ctx *Validation_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompute_clauses(ctx *Compute_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFor_clause(ctx *For_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOnline_or_offline(ctx *Online_or_offlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInto_clause1(ctx *Into_clause1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_key_value(ctx *Partition_key_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_key_value(ctx *Subpartition_key_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAssociate_statistics(ctx *Associate_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_association(ctx *Column_associationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_association(ctx *Function_associationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndextype_name(ctx *Indextype_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_statistics_type(ctx *Using_statistics_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStatistics_type_name(ctx *Statistics_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_cost_clause(ctx *Default_cost_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCpu_cost(ctx *Cpu_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIo_cost(ctx *Io_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNetwork_cost(ctx *Network_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_selectivity_clause(ctx *Default_selectivity_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_selectivity(ctx *Default_selectivityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStorage_table_clause(ctx *Storage_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnified_auditing(ctx *Unified_auditingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPolicy_name(ctx *Policy_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_traditional(ctx *Audit_traditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_direct_path(ctx *Audit_direct_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_container_clause(ctx *Audit_container_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_operation_clause(ctx *Audit_operation_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAuditing_by_clause(ctx *Auditing_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_user(ctx *Audit_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAudit_schema_object_clause(ctx *Audit_schema_object_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSql_operation(ctx *Sql_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAuditing_on_clause(ctx *Auditing_on_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_name(ctx *Model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_name(ctx *Object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProfile_name(ctx *Profile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSql_statement_shortcut(ctx *Sql_statement_shortcutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_index(ctx *Drop_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisassociate_statistics(ctx *Disassociate_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_indextype(ctx *Drop_indextypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_inmemory_join_group(ctx *Drop_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_table(ctx *Flashback_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRestore_point(ctx *Restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPurge_statement(ctx *Purge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNoaudit_statement(ctx *Noaudit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRename_object(ctx *Rename_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGrant_statement(ctx *Grant_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitContainer_clause(ctx *Container_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRevoke_statement(ctx *Revoke_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRevoke_system_privilege(ctx *Revoke_system_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRevokee_clause(ctx *Revokee_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRevoke_object_privileges(ctx *Revoke_object_privilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_object_clause(ctx *On_object_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRevoke_roles_from_programs(ctx *Revoke_roles_from_programsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProgram_unit(ctx *Program_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_dimension(ctx *Create_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_directory(ctx *Create_directoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDirectory_name(ctx *Directory_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDirectory_path(ctx *Directory_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_inmemory_join_group(ctx *Create_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_hierarchy(ctx *Drop_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_library(ctx *Alter_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_java(ctx *Drop_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_library(ctx *Drop_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_java(ctx *Create_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_library(ctx *Create_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPlsql_library_source(ctx *Plsql_library_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCredential_name(ctx *Credential_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLibrary_editionable(ctx *Library_editionableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLibrary_debug(ctx *Library_debugContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompiler_parameters_clause(ctx *Compiler_parameters_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParameter_value(ctx *Parameter_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLibrary_name(ctx *Library_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_dimension(ctx *Alter_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevel_clause(ctx *Level_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHierarchy_clause(ctx *Hierarchy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDimension_join_clause(ctx *Dimension_join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAttribute_clause(ctx *Attribute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExtended_attribute_clause(ctx *Extended_attribute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_one_or_more_sub_clause(ctx *Column_one_or_more_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_view(ctx *Alter_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_view_editionable(ctx *Alter_view_editionableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_view(ctx *Create_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEditioning_clause(ctx *Editioning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitView_options(ctx *View_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitView_alias_constraint(ctx *View_alias_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_view_clause(ctx *Object_view_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInline_constraint(ctx *Inline_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInline_ref_constraint(ctx *Inline_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOut_of_line_ref_constraint(ctx *Out_of_line_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOut_of_line_constraint(ctx *Out_of_line_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstraint_state(ctx *Constraint_stateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltype_view_clause(ctx *Xmltype_view_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_schema_spec(ctx *Xml_schema_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_schema_url(ctx *Xml_schema_urlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_tablespace(ctx *Alter_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatafile_tempfile_clauses(ctx *Datafile_tempfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_logging_clauses(ctx *Tablespace_logging_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_group_clause(ctx *Tablespace_group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_group_name(ctx *Tablespace_group_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_state_clauses(ctx *Tablespace_state_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_mode_clause(ctx *Flashback_mode_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_tablespace_name(ctx *New_tablespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_tablespace(ctx *Create_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPermanent_tablespace_clause(ctx *Permanent_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_encryption_spec(ctx *Tablespace_encryption_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogging_clause(ctx *Logging_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExtent_management_clause(ctx *Extent_management_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSegment_management_clause(ctx *Segment_management_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTemporary_tablespace_clause(ctx *Temporary_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUndo_tablespace_clause(ctx *Undo_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_retention_clause(ctx *Tablespace_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_tablespace_set(ctx *Create_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPermanent_tablespace_attrs(ctx *Permanent_tablespace_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_encryption_clause(ctx *Tablespace_encryption_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_tablespace_params(ctx *Default_tablespace_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_table_compression(ctx *Default_table_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLow_high(ctx *Low_highContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_index_compression(ctx *Default_index_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmmemory_clause(ctx *Inmmemory_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatafile_specification(ctx *Datafile_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTempfile_specification(ctx *Tempfile_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatafile_tempfile_spec(ctx *Datafile_tempfile_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRedo_log_file_spec(ctx *Redo_log_file_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAutoextend_clause(ctx *Autoextend_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMaxsize_clause(ctx *Maxsize_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBuild_clause(ctx *Build_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartial_index_clause(ctx *Partial_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParallel_clause(ctx *Parallel_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_materialized_view(ctx *Alter_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_mv_option1(ctx *Alter_mv_option1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_mv_refresh(ctx *Alter_mv_refreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRollback_segment(ctx *Rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_mv_column_clause(ctx *Modify_mv_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_materialized_view_log(ctx *Alter_materialized_view_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_mv_log_column_clause(ctx *Add_mv_log_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMove_mv_log_clause(ctx *Move_mv_log_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMv_log_augmentation(ctx *Mv_log_augmentationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatetime_expr(ctx *Datetime_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInterval_expr(ctx *Interval_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSynchronous_or_asynchronous(ctx *Synchronous_or_asynchronousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIncluding_or_excluding(ctx *Including_or_excludingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_materialized_view_log(ctx *Create_materialized_view_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_values_clause(ctx *New_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_materialized_zonemap(ctx *Create_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_materialized_zonemap(ctx *Alter_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_materialized_zonemap(ctx *Drop_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitZonemap_refresh_clause(ctx *Zonemap_refresh_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitZonemap_attributes(ctx *Zonemap_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitZonemap_name(ctx *Zonemap_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOperator_name(ctx *Operator_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOperator_function_name(ctx *Operator_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_zonemap_on_table(ctx *Create_zonemap_on_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_zonemap_as_subquery(ctx *Create_zonemap_as_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_operator(ctx *Alter_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_operator(ctx *Drop_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_operator(ctx *Create_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBinding_clause(ctx *Binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_binding_clause(ctx *Add_binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImplementation_clause(ctx *Implementation_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrimary_operator_list(ctx *Primary_operator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrimary_operator_item(ctx *Primary_operator_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOperator_context_clause(ctx *Operator_context_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_function_clause(ctx *Using_function_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_binding_clause(ctx *Drop_binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_materialized_view(ctx *Create_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitScoped_table_ref_constraint(ctx *Scoped_table_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMv_column_alias(ctx *Mv_column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_mv_refresh(ctx *Create_mv_refreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_materialized_view(ctx *Drop_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_context(ctx *Create_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOracle_namespace(ctx *Oracle_namespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_cluster(ctx *Create_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_profile(ctx *Create_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitResource_parameters(ctx *Resource_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPassword_parameters(ctx *Password_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_lockdown_profile(ctx *Create_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStatic_base_profile(ctx *Static_base_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDynamic_base_profile(ctx *Dynamic_base_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_outline(ctx *Create_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_restore_point(ctx *Create_restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_role(ctx *Create_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_table(ctx *Create_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltype_table(ctx *Xmltype_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltype_virtual_columns(ctx *Xmltype_virtual_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltype_column_properties(ctx *Xmltype_column_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltype_storage(ctx *Xmltype_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlschema_spec(ctx *Xmlschema_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_table(ctx *Object_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_type(ctx *Object_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOid_index_clause(ctx *Oid_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOid_clause(ctx *Oid_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_properties(ctx *Object_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_table_substitution(ctx *Object_table_substitutionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRelational_table(ctx *Relational_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImmutable_table_clauses(ctx *Immutable_table_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImmutable_table_no_drop_clause(ctx *Immutable_table_no_drop_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImmutable_table_no_delete_clause(ctx *Immutable_table_no_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBlockchain_table_clauses(ctx *Blockchain_table_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBlockchain_drop_table_clause(ctx *Blockchain_drop_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBlockchain_row_retention_clause(ctx *Blockchain_row_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBlockchain_hash_and_data_format_clause(ctx *Blockchain_hash_and_data_format_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_properties(ctx *Table_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRead_only_clause(ctx *Read_only_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndexing_clause(ctx *Indexing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAttribute_clustering_clause(ctx *Attribute_clustering_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClustering_join(ctx *Clustering_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClustering_join_item(ctx *Clustering_join_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEquijoin_condition(ctx *Equijoin_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCluster_clause(ctx *Cluster_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClustering_columns(ctx *Clustering_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClustering_column_group(ctx *Clustering_column_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitYes_no(ctx *Yes_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitZonemap_clause(ctx *Zonemap_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogical_replication_clause(ctx *Logical_replication_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRelational_property(ctx *Relational_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_partitioning_clauses(ctx *Table_partitioning_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRange_partitions(ctx *Range_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitList_partitions(ctx *List_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHash_partitions(ctx *Hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndividual_hash_partitions(ctx *Individual_hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHash_partitions_by_quantity(ctx *Hash_partitions_by_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHash_partition_quantity(ctx *Hash_partition_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComposite_range_partitions(ctx *Composite_range_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComposite_list_partitions(ctx *Composite_list_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComposite_hash_partitions(ctx *Composite_hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReference_partitioning(ctx *Reference_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReference_partition_desc(ctx *Reference_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSystem_partitioning(ctx *System_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRange_partition_desc(ctx *Range_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitList_partition_desc(ctx *List_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_template(ctx *Subpartition_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHash_subpartition_quantity(ctx *Hash_subpartition_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_by_range(ctx *Subpartition_by_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_by_list(ctx *Subpartition_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_by_hash(ctx *Subpartition_by_hashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_name(ctx *Subpartition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRange_subpartition_desc(ctx *Range_subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitList_subpartition_desc(ctx *List_subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndividual_hash_subparts(ctx *Individual_hash_subpartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHash_subparts_by_quantity(ctx *Hash_subparts_by_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRange_values_clause(ctx *Range_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitList_values_clause(ctx *List_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_partition_description(ctx *Table_partition_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartitioning_storage_clause(ctx *Partitioning_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_partitioning_storage(ctx *Lob_partitioning_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatatype_null_enable(ctx *Datatype_null_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSize_clause(ctx *Size_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_compression(ctx *Table_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_table_clause(ctx *Inmemory_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_attributes(ctx *Inmemory_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_memcompress(ctx *Inmemory_memcompressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_priority(ctx *Inmemory_priorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_distribute(ctx *Inmemory_distributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_duplicate(ctx *Inmemory_duplicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInmemory_column_clause(ctx *Inmemory_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPhysical_attributes_clause(ctx *Physical_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStorage_clause(ctx *Storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDeferred_segment_creation(ctx *Deferred_segment_creationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSegment_attributes_clause(ctx *Segment_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPhysical_properties(ctx *Physical_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_clause(ctx *Ilm_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_policy_clause(ctx *Ilm_policy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_compression_policy(ctx *Ilm_compression_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_tiering_policy(ctx *Ilm_tiering_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_after_on(ctx *Ilm_after_onContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSegment_group(ctx *Segment_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_inmemory_policy(ctx *Ilm_inmemory_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIlm_time_period(ctx *Ilm_time_periodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHeap_org_table_clause(ctx *Heap_org_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExternal_table_clause(ctx *External_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAccess_driver_type(ctx *Access_driver_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExternal_table_data_props(ctx *External_table_data_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOpaque_format_spec(ctx *Opaque_format_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRecord_format_info(ctx *Record_format_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEt_record_spec_options(ctx *Et_record_spec_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEt_record_spec_option(ctx *Et_record_spec_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEt_output_files(ctx *Et_output_filesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEt_output_file(ctx *Et_output_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDirectory_spec(ctx *Directory_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_spec(ctx *File_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_definitions(ctx *Field_definitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_options(ctx *Field_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_option(ctx *Field_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_list(ctx *Field_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_item(ctx *Field_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_name(ctx *Field_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPos_spec(ctx *Pos_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPos_start(ctx *Pos_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPos_increment(ctx *Pos_incrementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPos_end(ctx *Pos_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPos_length(ctx *Pos_lengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatatype_spec(ctx *Datatype_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInit_spec(ctx *Init_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLls_clause(ctx *Lls_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDelim_spec(ctx *Delim_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrim_spec(ctx *Trim_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_date_format(ctx *Field_date_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_transforms(ctx *Column_transformsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTransform(ctx *TransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSource_field(ctx *Source_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLobfile_item(ctx *Lobfile_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLobfile_attr_list(ctx *Lobfile_attr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConcat_item(ctx *Concat_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRow_movement_clause(ctx *Row_movement_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_archive_clause(ctx *Flashback_archive_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLog_grp(ctx *Log_grpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_table_logging(ctx *Supplemental_table_loggingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_log_grp_clause(ctx *Supplemental_log_grp_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_id_key_clause(ctx *Supplemental_id_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAllocate_extent_clause(ctx *Allocate_extent_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDeallocate_unused_clause(ctx *Deallocate_unused_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitShrink_clause(ctx *Shrink_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRecords_per_block_clause(ctx *Records_per_block_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpgrade_table_clause(ctx *Upgrade_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTruncate_table(ctx *Truncate_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_table(ctx *Drop_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_tablespace(ctx *Drop_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_tablespace_set(ctx *Drop_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIncluding_contents_clause(ctx *Including_contents_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_view(ctx *Drop_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComment_on_column(ctx *Comment_on_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnable_or_disable(ctx *Enable_or_disableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAllow_or_disallow(ctx *Allow_or_disallowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_synonym(ctx *Alter_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_synonym(ctx *Create_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_synonym(ctx *Drop_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_spfile(ctx *Create_spfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSpfile_name(ctx *Spfile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPfile_name(ctx *Pfile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComment_on_table(ctx *Comment_on_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitComment_on_materialized(ctx *Comment_on_materializedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_analytic_view(ctx *Alter_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_add_cache_clause(ctx *Alter_add_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLevels_item(ctx *Levels_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMeasure_list(ctx *Measure_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_drop_cache_clause(ctx *Alter_drop_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_attribute_dimension(ctx *Alter_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_audit_policy(ctx *Alter_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_cluster(ctx *Alter_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_analytic_view(ctx *Drop_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_attribute_dimension(ctx *Drop_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_audit_policy(ctx *Drop_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_flashback_archive(ctx *Drop_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_cluster(ctx *Drop_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_context(ctx *Drop_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_directory(ctx *Drop_directoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_diskgroup(ctx *Drop_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_edition(ctx *Drop_editionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTruncate_cluster(ctx *Truncate_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCache_or_nocache(ctx *Cache_or_nocacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_database(ctx *Alter_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase_clause(ctx *Database_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStartup_clauses(ctx *Startup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitResetlogs_or_noresetlogs(ctx *Resetlogs_or_noresetlogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpgrade_or_downgrade(ctx *Upgrade_or_downgradeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRecovery_clauses(ctx *Recovery_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBegin_or_end(ctx *Begin_or_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGeneral_recovery(ctx *General_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFull_database_recovery(ctx *Full_database_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartial_database_recovery(ctx *Partial_database_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartial_database_recovery_10g(ctx *Partial_database_recovery_10gContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitManaged_standby_recovery(ctx *Managed_standby_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDb_name(ctx *Db_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase_file_clauses(ctx *Database_file_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_datafile_clause(ctx *Create_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_datafile_clause(ctx *Alter_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_tempfile_clause(ctx *Alter_tempfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMove_datafile_clause(ctx *Move_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogfile_clauses(ctx *Logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_logfile_clauses(ctx *Add_logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGroup_redo_logfile(ctx *Group_redo_logfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_logfile_clauses(ctx *Drop_logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSwitch_logfile_clause(ctx *Switch_logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_db_logging(ctx *Supplemental_db_loggingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_or_drop(ctx *Add_or_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_plsql_clause(ctx *Supplemental_plsql_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogfile_descriptor(ctx *Logfile_descriptorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitControlfile_clauses(ctx *Controlfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrace_file_clause(ctx *Trace_file_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStandby_database_clauses(ctx *Standby_database_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitActivate_standby_db_clause(ctx *Activate_standby_db_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMaximize_standby_db_clause(ctx *Maximize_standby_db_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRegister_logfile_clause(ctx *Register_logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCommit_switchover_clause(ctx *Commit_switchover_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStart_standby_clause(ctx *Start_standby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStop_standby_clause(ctx *Stop_standby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConvert_database_clause(ctx *Convert_database_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_settings_clause(ctx *Default_settings_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_time_zone_clause(ctx *Set_time_zone_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInstance_clauses(ctx *Instance_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSecurity_clause(ctx *Security_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDomain(ctx *DomainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase(ctx *DatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEdition_name(ctx *Edition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFilenumber(ctx *FilenumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFilename(ctx *FilenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrepare_clause(ctx *Prepare_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_mirror_clause(ctx *Drop_mirror_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLost_write_protection(ctx *Lost_write_protectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCdb_fleet_clauses(ctx *Cdb_fleet_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLead_cdb_clause(ctx *Lead_cdb_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLead_cdb_uri_clause(ctx *Lead_cdb_uri_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProperty_clauses(ctx *Property_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReplay_upgrade_clauses(ctx *Replay_upgrade_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_database_link(ctx *Alter_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPassword_value(ctx *Password_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLink_authentication(ctx *Link_authenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_database(ctx *Create_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase_logging_clauses(ctx *Database_logging_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatabase_logging_sub_clause(ctx *Database_logging_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_clauses(ctx *Tablespace_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnable_pluggable_database(ctx *Enable_pluggable_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFile_name_convert(ctx *File_name_convertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFilename_convert_sub_clause(ctx *Filename_convert_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace_datafile_clauses(ctx *Tablespace_datafile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUndo_mode_clause(ctx *Undo_mode_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_tablespace(ctx *Default_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_temp_tablespace(ctx *Default_temp_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUndo_tablespace(ctx *Undo_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_database(ctx *Drop_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCreate_database_link(ctx *Create_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDblink(ctx *DblinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_database_link(ctx *Drop_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_tablespace_set(ctx *Alter_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_tablespace_attrs(ctx *Alter_tablespace_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_tablespace_encryption(ctx *Alter_tablespace_encryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTs_file_name_convert(ctx *Ts_file_name_convertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_role(ctx *Alter_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRole_identified_clause(ctx *Role_identified_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_table(ctx *Alter_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMemoptimize_read_write_clause(ctx *Memoptimize_read_write_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_table_properties(ctx *Alter_table_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_table_partitioning(ctx *Alter_table_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_table_partition(ctx *Add_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_table_partition(ctx *Drop_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_table_partition(ctx *Merge_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_table_partition(ctx *Modify_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSplit_table_partition(ctx *Split_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTruncate_table_partition(ctx *Truncate_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExchange_table_partition(ctx *Exchange_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCoalesce_table_partition(ctx *Coalesce_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_interval_partition(ctx *Alter_interval_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_extended_names(ctx *Partition_extended_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubpartition_extended_names(ctx *Subpartition_extended_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_table_properties_1(ctx *Alter_table_properties_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_iot_clauses(ctx *Alter_iot_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_mapping_table_clause(ctx *Alter_mapping_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_overflow_clause(ctx *Alter_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_overflow_clause(ctx *Add_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_index_clauses(ctx *Update_index_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_global_index_clause(ctx *Update_global_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_all_indexes_clause(ctx *Update_all_indexes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_all_indexes_index_clause(ctx *Update_all_indexes_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_index_partition(ctx *Update_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_index_subpartition(ctx *Update_index_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnable_disable_clause(ctx *Enable_disable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_index_clause(ctx *Using_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_attributes(ctx *Index_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSort_or_nosort(ctx *Sort_or_nosortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExceptions_clause(ctx *Exceptions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMove_table_clause(ctx *Move_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_org_table_clause(ctx *Index_org_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMapping_table_clause(ctx *Mapping_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKey_compression(ctx *Key_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_org_overflow_clause(ctx *Index_org_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_clauses(ctx *Column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_collection_retrieval(ctx *Modify_collection_retrievalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCollection_item(ctx *Collection_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRename_column_clause(ctx *Rename_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOld_column_name(ctx *Old_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_column_name(ctx *New_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_modify_drop_column_clauses(ctx *Add_modify_drop_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_column_clause(ctx *Drop_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_column_clauses(ctx *Modify_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_col_visibility(ctx *Modify_col_visibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_col_properties(ctx *Modify_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_col_substitutable(ctx *Modify_col_substitutableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_column_clause(ctx *Add_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAlter_varray_col_properties(ctx *Alter_varray_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVarray_col_properties(ctx *Varray_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVarray_storage_clause(ctx *Varray_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_segname(ctx *Lob_segnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_item(ctx *Lob_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_storage_parameters(ctx *Lob_storage_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_storage_clause(ctx *Lob_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_lob_storage_clause(ctx *Modify_lob_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModify_lob_parameters(ctx *Modify_lob_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_parameters(ctx *Lob_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_deduplicate_clause(ctx *Lob_deduplicate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_compression_clause(ctx *Lob_compression_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_retention_clause(ctx *Lob_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEncryption_spec(ctx *Encryption_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTablespace(ctx *TablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVarray_item(ctx *Varray_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_properties(ctx *Column_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLob_partition_storage(ctx *Lob_partition_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPeriod_definition(ctx *Period_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStart_time_column(ctx *Start_time_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnd_time_column(ctx *End_time_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_definition(ctx *Column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_collation_name(ctx *Column_collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentity_clause(ctx *Identity_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentity_options_parentheses(ctx *Identity_options_parenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentity_options(ctx *Identity_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVirtual_column_definition(ctx *Virtual_column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAutogenerated_sequence_definition(ctx *Autogenerated_sequence_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEvaluation_edition_clause(ctx *Evaluation_edition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOut_of_line_part_storage(ctx *Out_of_line_part_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNested_table_col_properties(ctx *Nested_table_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNested_item(ctx *Nested_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubstitutable_column_clause(ctx *Substitutable_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_name(ctx *Partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSupplemental_logging_props(ctx *Supplemental_logging_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_or_attribute(ctx *Column_or_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_type_col_properties(ctx *Object_type_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstraint_clauses(ctx *Constraint_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOld_constraint_name(ctx *Old_constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNew_constraint_name(ctx *New_constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_constraint_clause(ctx *Drop_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_primary_key_or_unique_or_generic_clause(ctx *Drop_primary_key_or_unique_or_generic_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_constraint(ctx *Add_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAdd_constraint_clause(ctx *Add_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCheck_constraint(ctx *Check_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDrop_constraint(ctx *Drop_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitEnable_constraint(ctx *Enable_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDisable_constraint(ctx *Disable_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReferences_clause(ctx *References_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOn_delete_clause(ctx *On_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnique_key_clause(ctx *Unique_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrimary_key_clause(ctx *Primary_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAnonymous_block(ctx *Anonymous_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCall_spec(ctx *Call_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJava_spec(ctx *Java_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitC_spec(ctx *C_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitC_agent_in_clause(ctx *C_agent_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitC_parameters_clause(ctx *C_parameters_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDefault_value_part(ctx *Default_value_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDeclare_spec(ctx *Declare_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCursor_declaration(ctx *Cursor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParameter_spec(ctx *Parameter_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitException_declaration(ctx *Exception_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPragma_declaration(ctx *Pragma_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRecord_type_def(ctx *Record_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitField_spec(ctx *Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_type_def(ctx *Table_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_indexed_by_part(ctx *Table_indexed_by_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVarray_type_def(ctx *Varray_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSeq_of_statements(ctx *Seq_of_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLabel_declaration(ctx *Label_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSwallow_to_semi(ctx *Swallow_to_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAssignment_statement(ctx *Assignment_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExit_statement(ctx *Exit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGoto_statement(ctx *Goto_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitElsif_part(ctx *Elsif_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitElse_part(ctx *Else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCursor_loop_param(ctx *Cursor_loop_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitForall_statement(ctx *Forall_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBounds_clause(ctx *Bounds_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBetween_bound(ctx *Between_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLower_bound(ctx *Lower_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpper_bound(ctx *Upper_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNull_statement(ctx *Null_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRaise_statement(ctx *Raise_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCall_statement(ctx *Call_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPipe_row_statement(ctx *Pipe_row_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitException_handler(ctx *Exception_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrigger_block(ctx *Trigger_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSql_statement(ctx *Sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExecute_immediate(ctx *Execute_immediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitClose_statement(ctx *Close_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOpen_statement(ctx *Open_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFetch_statement(ctx *Fetch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOpen_for_statement(ctx *Open_for_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTransaction_control_statements(ctx *Transaction_control_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_transaction_command(ctx *Set_transaction_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_constraint_command(ctx *Set_constraint_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCommit_statement(ctx *Commit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWrite_clause(ctx *Write_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRollback_statement(ctx *Rollback_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSavepoint_statement(ctx *Savepoint_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExplain_statement(ctx *Explain_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSelect_only_statement(ctx *Select_only_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubquery_factoring_clause(ctx *Subquery_factoring_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFactoring_element(ctx *Factoring_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSearch_clause(ctx *Search_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCycle_clause(ctx *Cycle_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubquery_basic_elements(ctx *Subquery_basic_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubquery_operation_part(ctx *Subquery_operation_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuery_block(ctx *Query_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSelected_list(ctx *Selected_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSelect_list_elements(ctx *Select_list_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_wild(ctx *Table_wildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref_list(ctx *Table_ref_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref(ctx *Table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref_aux(ctx *Table_ref_auxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref_aux_internal_one(ctx *Table_ref_aux_internal_oneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref_aux_internal_two(ctx *Table_ref_aux_internal_twoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_ref_aux_internal_three(ctx *Table_ref_aux_internal_threeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_clause(ctx *Join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_on_part(ctx *Join_on_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJoin_using_part(ctx *Join_using_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOuter_join_type(ctx *Outer_join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuery_partition_clause(ctx *Query_partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFlashback_query_clause(ctx *Flashback_query_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_clause(ctx *Pivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_element(ctx *Pivot_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_for_clause(ctx *Pivot_for_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_in_clause(ctx *Pivot_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_in_clause_element(ctx *Pivot_in_clause_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPivot_in_clause_elements(ctx *Pivot_in_clause_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnpivot_in_clause(ctx *Unpivot_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnpivot_in_elements(ctx *Unpivot_in_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStart_part(ctx *Start_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGroup_by_elements(ctx *Group_by_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRollup_cube_clause(ctx *Rollup_cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGrouping_sets_elements(ctx *Grouping_sets_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_clause(ctx *Model_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCell_reference_options(ctx *Cell_reference_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReturn_rows_clause(ctx *Return_rows_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReference_model(ctx *Reference_modelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMain_model(ctx *Main_modelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_column_clauses(ctx *Model_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_column_partition_part(ctx *Model_column_partition_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_column_list(ctx *Model_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_column(ctx *Model_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_rules_clause(ctx *Model_rules_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_rules_part(ctx *Model_rules_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_rules_element(ctx *Model_rules_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCell_assignment(ctx *Cell_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_iterate_clause(ctx *Model_iterate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUntil_part(ctx *Until_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOrder_by_elements(ctx *Order_by_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFetch_clause(ctx *Fetch_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFor_update_clause(ctx *For_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFor_update_of_part(ctx *For_update_of_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFor_update_options(ctx *For_update_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUpdate_set_clause(ctx *Update_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_based_update_set_clause(ctx *Column_based_update_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInsert_statement(ctx *Insert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSingle_table_insert(ctx *Single_table_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMulti_table_insert(ctx *Multi_table_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMulti_table_element(ctx *Multi_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConditional_insert_clause(ctx *Conditional_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConditional_insert_when_part(ctx *Conditional_insert_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConditional_insert_else_part(ctx *Conditional_insert_else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInsert_into_clause(ctx *Insert_into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitValues_clause(ctx *Values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_statement(ctx *Merge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_element(ctx *Merge_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_update_delete_part(ctx *Merge_update_delete_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSelected_tableview(ctx *Selected_tableviewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLock_table_statement(ctx *Lock_table_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWait_nowait_part(ctx *Wait_nowait_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLock_table_element(ctx *Lock_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLock_mode(ctx *Lock_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGeneral_table_ref(ctx *General_table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStatic_returning_clause(ctx *Static_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitError_logging_clause(ctx *Error_logging_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitError_logging_into_part(ctx *Error_logging_into_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitError_logging_reject_part(ctx *Error_logging_reject_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDml_table_expression_clause(ctx *Dml_table_expression_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_collection_expression(ctx *Table_collection_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSubquery_restriction_clause(ctx *Subquery_restriction_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSample_clause(ctx *Sample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSeed_part(ctx *Seed_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_condition(ctx *Json_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCursor_expression(ctx *Cursor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogical_expression(ctx *Logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnary_logical_expression(ctx *Unary_logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLogical_operation(ctx *Logical_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMultiset_expression(ctx *Multiset_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRelational_expression(ctx *Relational_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCompound_expression(ctx *Compound_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRelational_operator(ctx *Relational_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIn_elements(ctx *In_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBetween_elements(ctx *Between_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInterval_expression(ctx *Interval_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_expression(ctx *Model_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitModel_expression_element(ctx *Model_expression_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSingle_column_for_loop(ctx *Single_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMulti_column_for_loop(ctx *Multi_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUnary_expression(ctx *Unary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSimple_case_statement(ctx *Simple_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSimple_case_when_part(ctx *Simple_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSearched_case_statement(ctx *Searched_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSearched_case_when_part(ctx *Searched_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCase_else_part(ctx *Case_else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuantified_expression(ctx *Quantified_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitString_function(ctx *String_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStandard_function(ctx *Standard_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_function(ctx *Json_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_object_content(ctx *Json_object_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_object_entry(ctx *Json_object_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_table_clause(ctx *Json_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_array_element(ctx *Json_array_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_on_null_clause(ctx *Json_on_null_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_return_clause(ctx *Json_return_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_transform_op(ctx *Json_transform_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_column_clause(ctx *Json_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_column_definition(ctx *Json_column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_query_returning_clause(ctx *Json_query_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_query_return_type(ctx *Json_query_return_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_query_wrapper_clause(ctx *Json_query_wrapper_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_query_on_error_clause(ctx *Json_query_on_error_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_query_on_empty_clause(ctx *Json_query_on_empty_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_value_return_clause(ctx *Json_value_return_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_value_return_type(ctx *Json_value_return_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitJson_value_on_mismatch_clause(ctx *Json_value_on_mismatch_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNumeric_function_wrapper(ctx *Numeric_function_wrapperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNumeric_function(ctx *Numeric_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitListagg_overflow_clause(ctx *Listagg_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOther_function(ctx *Other_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOver_clause_keyword(ctx *Over_clause_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWithin_or_over_clause_keyword(ctx *Within_or_over_clause_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitStandard_prediction_function_keyword(ctx *Standard_prediction_function_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWindowing_clause(ctx *Windowing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWindowing_type(ctx *Windowing_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWindowing_elements(ctx *Windowing_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUsing_element(ctx *Using_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCollect_order_by_part(ctx *Collect_order_by_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWithin_or_over_part(ctx *Within_or_over_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCost_matrix_clause(ctx *Cost_matrix_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_passing_clause(ctx *Xml_passing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_attributes_clause(ctx *Xml_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_namespaces_clause(ctx *Xml_namespaces_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_table_column(ctx *Xml_table_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_general_default_part(ctx *Xml_general_default_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_multiuse_expression_element(ctx *Xml_multiuse_expression_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlroot_param_version_part(ctx *Xmlroot_param_version_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlroot_param_standalone_part(ctx *Xmlroot_param_standalone_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlserialize_param_enconding_part(ctx *Xmlserialize_param_enconding_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlserialize_param_version_part(ctx *Xmlserialize_param_version_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmlserialize_param_ident_part(ctx *Xmlserialize_param_ident_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSql_plus_command(ctx *Sql_plus_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWhenever_command(ctx *Whenever_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSet_command(ctx *Set_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTiming_command(ctx *Timing_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPartition_extension_clause(ctx *Partition_extension_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuantitative_where_stmt(ctx *Quantitative_where_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitInto_clause(ctx *Into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXml_column_name(ctx *Xml_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCost_class_name(ctx *Cost_class_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAttribute_name(ctx *Attribute_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSavepoint_name(ctx *Savepoint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRollback_segment_name(ctx *Rollback_segment_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_var_name(ctx *Table_var_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSchema_name(ctx *Schema_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRoutine_name(ctx *Routine_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPackage_name(ctx *Package_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitImplementation_type_name(ctx *Implementation_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParameter_name(ctx *Parameter_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitReference_model_name(ctx *Reference_model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitMain_model_name(ctx *Main_model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitContainer_tableview_name(ctx *Container_tableview_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitAggregate_function_name(ctx *Aggregate_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuery_name(ctx *Query_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGrantee_name(ctx *Grantee_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRole_name(ctx *Role_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstraint_name(ctx *Constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLabel_name(ctx *Label_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSequence_name(ctx *Sequence_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitException_name(ctx *Exception_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitProcedure_name(ctx *Procedure_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTrigger_name(ctx *Trigger_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRecord_name(ctx *Record_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitCollection_name(ctx *Collection_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitLink_name(ctx *Link_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTableview_name(ctx *Tableview_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitXmltable(ctx *XmltableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitChar_set_name(ctx *Char_set_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSynonym_name(ctx *Synonym_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSchema_object_name(ctx *Schema_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDir_object_name(ctx *Dir_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitUser_object_name(ctx *User_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGrant_object_name(ctx *Grant_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitColumn_list(ctx *Column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitParen_column_list(ctx *Paren_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitKeep_clause(ctx *Keep_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_argument(ctx *Function_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_argument_analytic(ctx *Function_argument_analyticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitFunction_argument_modeling(ctx *Function_argument_modelingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRespect_or_ignore_nulls(ctx *Respect_or_ignore_nullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitType_spec(ctx *Type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitPrecision_part(ctx *Precision_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNative_datatype_element(ctx *Native_datatype_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitBind_variable(ctx *Bind_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitGeneral_element_part(ctx *General_element_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitTable_element(ctx *Table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitObject_privilege(ctx *Object_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitSystem_privilege(ctx *System_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitConstant_without_variable(ctx *Constant_without_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNumeric(ctx *NumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNumeric_negative(ctx *Numeric_negativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitQuoted_string(ctx *Quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitChar_str(ctx *Char_strContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitId_expression(ctx *Id_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitOuter_join_sign(ctx *Outer_join_signContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitRegular_id(ctx *Regular_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNon_reserved_keywords_in_12c(ctx *Non_reserved_keywords_in_12cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNon_reserved_keywords_pre12c(ctx *Non_reserved_keywords_pre12cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitString_function_name(ctx *String_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlSqlParserVisitor) VisitNumeric_function_name(ctx *Numeric_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}
