// Code generated from PlSqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PlSqlParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PlSqlParser.
type PlSqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlSqlParser#sql_script.
	VisitSql_script(ctx *Sql_scriptContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unit_statement.
	VisitUnit_statement(ctx *Unit_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_diskgroup.
	VisitAlter_diskgroup(ctx *Alter_diskgroupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_disk_clause.
	VisitAdd_disk_clause(ctx *Add_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_disk_clause.
	VisitDrop_disk_clause(ctx *Drop_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#resize_disk_clause.
	VisitResize_disk_clause(ctx *Resize_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#replace_disk_clause.
	VisitReplace_disk_clause(ctx *Replace_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#wait_nowait.
	VisitWait_nowait(ctx *Wait_nowaitContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rename_disk_clause.
	VisitRename_disk_clause(ctx *Rename_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disk_online_clause.
	VisitDisk_online_clause(ctx *Disk_online_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disk_offline_clause.
	VisitDisk_offline_clause(ctx *Disk_offline_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#timeout_clause.
	VisitTimeout_clause(ctx *Timeout_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rebalance_diskgroup_clause.
	VisitRebalance_diskgroup_clause(ctx *Rebalance_diskgroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#phase.
	VisitPhase(ctx *PhaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#check_diskgroup_clause.
	VisitCheck_diskgroup_clause(ctx *Check_diskgroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_template_clauses.
	VisitDiskgroup_template_clauses(ctx *Diskgroup_template_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#qualified_template_clause.
	VisitQualified_template_clause(ctx *Qualified_template_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#redundancy_clause.
	VisitRedundancy_clause(ctx *Redundancy_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#striping_clause.
	VisitStriping_clause(ctx *Striping_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#force_noforce.
	VisitForce_noforce(ctx *Force_noforceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_directory_clauses.
	VisitDiskgroup_directory_clauses(ctx *Diskgroup_directory_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dir_name.
	VisitDir_name(ctx *Dir_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_alias_clauses.
	VisitDiskgroup_alias_clauses(ctx *Diskgroup_alias_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_volume_clauses.
	VisitDiskgroup_volume_clauses(ctx *Diskgroup_volume_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_volume_clause.
	VisitAdd_volume_clause(ctx *Add_volume_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_volume_clause.
	VisitModify_volume_clause(ctx *Modify_volume_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_attributes.
	VisitDiskgroup_attributes(ctx *Diskgroup_attributesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_diskgroup_file.
	VisitModify_diskgroup_file(ctx *Modify_diskgroup_fileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disk_region_clause.
	VisitDisk_region_clause(ctx *Disk_region_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_diskgroup_file_clause.
	VisitDrop_diskgroup_file_clause(ctx *Drop_diskgroup_file_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#convert_redundancy_clause.
	VisitConvert_redundancy_clause(ctx *Convert_redundancy_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#usergroup_clauses.
	VisitUsergroup_clauses(ctx *Usergroup_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_clauses.
	VisitUser_clauses(ctx *User_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_permissions_clause.
	VisitFile_permissions_clause(ctx *File_permissions_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_owner_clause.
	VisitFile_owner_clause(ctx *File_owner_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#scrub_clause.
	VisitScrub_clause(ctx *Scrub_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quotagroup_clauses.
	VisitQuotagroup_clauses(ctx *Quotagroup_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#property_name.
	VisitProperty_name(ctx *Property_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#property_value.
	VisitProperty_value(ctx *Property_valueContext) interface{}

	// Visit a parse tree produced by PlSqlParser#filegroup_clauses.
	VisitFilegroup_clauses(ctx *Filegroup_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_filegroup_clause.
	VisitAdd_filegroup_clause(ctx *Add_filegroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_filegroup_clause.
	VisitModify_filegroup_clause(ctx *Modify_filegroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#move_to_filegroup_clause.
	VisitMove_to_filegroup_clause(ctx *Move_to_filegroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_filegroup_clause.
	VisitDrop_filegroup_clause(ctx *Drop_filegroup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quorum_regular.
	VisitQuorum_regular(ctx *Quorum_regularContext) interface{}

	// Visit a parse tree produced by PlSqlParser#undrop_disk_clause.
	VisitUndrop_disk_clause(ctx *Undrop_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#diskgroup_availability.
	VisitDiskgroup_availability(ctx *Diskgroup_availabilityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#enable_disable_volume.
	VisitEnable_disable_volume(ctx *Enable_disable_volumeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_function.
	VisitDrop_function(ctx *Drop_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_flashback_archive.
	VisitAlter_flashback_archive(ctx *Alter_flashback_archiveContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_hierarchy.
	VisitAlter_hierarchy(ctx *Alter_hierarchyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_function.
	VisitAlter_function(ctx *Alter_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_java.
	VisitAlter_java(ctx *Alter_javaContext) interface{}

	// Visit a parse tree produced by PlSqlParser#match_string.
	VisitMatch_string(ctx *Match_stringContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_function_body.
	VisitCreate_function_body(ctx *Create_function_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sql_macro_body.
	VisitSql_macro_body(ctx *Sql_macro_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parallel_enable_clause.
	VisitParallel_enable_clause(ctx *Parallel_enable_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_by_clause.
	VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#result_cache_clause.
	VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#relies_on_part.
	VisitRelies_on_part(ctx *Relies_on_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#streaming_clause.
	VisitStreaming_clause(ctx *Streaming_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_outline.
	VisitAlter_outline(ctx *Alter_outlineContext) interface{}

	// Visit a parse tree produced by PlSqlParser#outline_options.
	VisitOutline_options(ctx *Outline_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_lockdown_profile.
	VisitAlter_lockdown_profile(ctx *Alter_lockdown_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lockdown_feature.
	VisitLockdown_feature(ctx *Lockdown_featureContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lockdown_options.
	VisitLockdown_options(ctx *Lockdown_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lockdown_statements.
	VisitLockdown_statements(ctx *Lockdown_statementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#statement_clauses.
	VisitStatement_clauses(ctx *Statement_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#clause_options.
	VisitClause_options(ctx *Clause_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#option_values.
	VisitOption_values(ctx *Option_valuesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#string_list.
	VisitString_list(ctx *String_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disable_enable.
	VisitDisable_enable(ctx *Disable_enableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_lockdown_profile.
	VisitDrop_lockdown_profile(ctx *Drop_lockdown_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_package.
	VisitDrop_package(ctx *Drop_packageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_package.
	VisitAlter_package(ctx *Alter_packageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_package.
	VisitCreate_package(ctx *Create_packageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_package_body.
	VisitCreate_package_body(ctx *Create_package_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#package_obj_spec.
	VisitPackage_obj_spec(ctx *Package_obj_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#procedure_spec.
	VisitProcedure_spec(ctx *Procedure_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_spec.
	VisitFunction_spec(ctx *Function_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#package_obj_body.
	VisitPackage_obj_body(ctx *Package_obj_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_pmem_filestore.
	VisitAlter_pmem_filestore(ctx *Alter_pmem_filestoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_pmem_filestore.
	VisitDrop_pmem_filestore(ctx *Drop_pmem_filestoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_procedure.
	VisitDrop_procedure(ctx *Drop_procedureContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_procedure.
	VisitAlter_procedure(ctx *Alter_procedureContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_body.
	VisitFunction_body(ctx *Function_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#procedure_body.
	VisitProcedure_body(ctx *Procedure_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_procedure_body.
	VisitCreate_procedure_body(ctx *Create_procedure_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_resource_cost.
	VisitAlter_resource_cost(ctx *Alter_resource_costContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_outline.
	VisitDrop_outline(ctx *Drop_outlineContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_rollback_segment.
	VisitAlter_rollback_segment(ctx *Alter_rollback_segmentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_restore_point.
	VisitDrop_restore_point(ctx *Drop_restore_pointContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_rollback_segment.
	VisitDrop_rollback_segment(ctx *Drop_rollback_segmentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_role.
	VisitDrop_role(ctx *Drop_roleContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_pmem_filestore.
	VisitCreate_pmem_filestore(ctx *Create_pmem_filestoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pmem_filestore_options.
	VisitPmem_filestore_options(ctx *Pmem_filestore_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_path.
	VisitFile_path(ctx *File_pathContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_rollback_segment.
	VisitCreate_rollback_segment(ctx *Create_rollback_segmentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_trigger.
	VisitDrop_trigger(ctx *Drop_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_trigger.
	VisitAlter_trigger(ctx *Alter_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_trigger.
	VisitCreate_trigger(ctx *Create_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trigger_follows_clause.
	VisitTrigger_follows_clause(ctx *Trigger_follows_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trigger_when_clause.
	VisitTrigger_when_clause(ctx *Trigger_when_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#simple_dml_trigger.
	VisitSimple_dml_trigger(ctx *Simple_dml_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#for_each_row.
	VisitFor_each_row(ctx *For_each_rowContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compound_dml_trigger.
	VisitCompound_dml_trigger(ctx *Compound_dml_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#non_dml_trigger.
	VisitNon_dml_trigger(ctx *Non_dml_triggerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trigger_body.
	VisitTrigger_body(ctx *Trigger_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#routine_clause.
	VisitRoutine_clause(ctx *Routine_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compound_trigger_block.
	VisitCompound_trigger_block(ctx *Compound_trigger_blockContext) interface{}

	// Visit a parse tree produced by PlSqlParser#timing_point_section.
	VisitTiming_point_section(ctx *Timing_point_sectionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#non_dml_event.
	VisitNon_dml_event(ctx *Non_dml_eventContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dml_event_clause.
	VisitDml_event_clause(ctx *Dml_event_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dml_event_element.
	VisitDml_event_element(ctx *Dml_event_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dml_event_nested_clause.
	VisitDml_event_nested_clause(ctx *Dml_event_nested_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#referencing_clause.
	VisitReferencing_clause(ctx *Referencing_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#referencing_element.
	VisitReferencing_element(ctx *Referencing_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_type.
	VisitDrop_type(ctx *Drop_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_type.
	VisitAlter_type(ctx *Alter_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compile_type_clause.
	VisitCompile_type_clause(ctx *Compile_type_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#replace_type_clause.
	VisitReplace_type_clause(ctx *Replace_type_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_method_spec.
	VisitAlter_method_spec(ctx *Alter_method_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_method_element.
	VisitAlter_method_element(ctx *Alter_method_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_attribute_definition.
	VisitAlter_attribute_definition(ctx *Alter_attribute_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#attribute_definition.
	VisitAttribute_definition(ctx *Attribute_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_collection_clauses.
	VisitAlter_collection_clauses(ctx *Alter_collection_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dependent_handling_clause.
	VisitDependent_handling_clause(ctx *Dependent_handling_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dependent_exceptions_part.
	VisitDependent_exceptions_part(ctx *Dependent_exceptions_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_type.
	VisitCreate_type(ctx *Create_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_definition.
	VisitType_definition(ctx *Type_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_type_def.
	VisitObject_type_def(ctx *Object_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_as_part.
	VisitObject_as_part(ctx *Object_as_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_under_part.
	VisitObject_under_part(ctx *Object_under_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#nested_table_type_def.
	VisitNested_table_type_def(ctx *Nested_table_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sqlj_object_type.
	VisitSqlj_object_type(ctx *Sqlj_object_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_body.
	VisitType_body(ctx *Type_bodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_body_elements.
	VisitType_body_elements(ctx *Type_body_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#map_order_func_declaration.
	VisitMap_order_func_declaration(ctx *Map_order_func_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subprog_decl_in_type.
	VisitSubprog_decl_in_type(ctx *Subprog_decl_in_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#proc_decl_in_type.
	VisitProc_decl_in_type(ctx *Proc_decl_in_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#func_decl_in_type.
	VisitFunc_decl_in_type(ctx *Func_decl_in_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constructor_declaration.
	VisitConstructor_declaration(ctx *Constructor_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modifier_clause.
	VisitModifier_clause(ctx *Modifier_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_member_spec.
	VisitObject_member_spec(ctx *Object_member_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sqlj_object_type_attr.
	VisitSqlj_object_type_attr(ctx *Sqlj_object_type_attrContext) interface{}

	// Visit a parse tree produced by PlSqlParser#element_spec.
	VisitElement_spec(ctx *Element_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#element_spec_options.
	VisitElement_spec_options(ctx *Element_spec_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subprogram_spec.
	VisitSubprogram_spec(ctx *Subprogram_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#overriding_subprogram_spec.
	VisitOverriding_subprogram_spec(ctx *Overriding_subprogram_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#overriding_function_spec.
	VisitOverriding_function_spec(ctx *Overriding_function_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_procedure_spec.
	VisitType_procedure_spec(ctx *Type_procedure_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_function_spec.
	VisitType_function_spec(ctx *Type_function_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constructor_spec.
	VisitConstructor_spec(ctx *Constructor_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#map_order_function_spec.
	VisitMap_order_function_spec(ctx *Map_order_function_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pragma_clause.
	VisitPragma_clause(ctx *Pragma_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pragma_elements.
	VisitPragma_elements(ctx *Pragma_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_elements_parameter.
	VisitType_elements_parameter(ctx *Type_elements_parameterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_sequence.
	VisitDrop_sequence(ctx *Drop_sequenceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_sequence.
	VisitAlter_sequence(ctx *Alter_sequenceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_session.
	VisitAlter_session(ctx *Alter_sessionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_session_set_clause.
	VisitAlter_session_set_clause(ctx *Alter_session_set_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_sequence.
	VisitCreate_sequence(ctx *Create_sequenceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sequence_spec.
	VisitSequence_spec(ctx *Sequence_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sequence_start_clause.
	VisitSequence_start_clause(ctx *Sequence_start_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_analytic_view.
	VisitCreate_analytic_view(ctx *Create_analytic_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#classification_clause.
	VisitClassification_clause(ctx *Classification_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#caption_clause.
	VisitCaption_clause(ctx *Caption_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#description_clause.
	VisitDescription_clause(ctx *Description_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#classification_item.
	VisitClassification_item(ctx *Classification_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#language.
	VisitLanguage(ctx *LanguageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cav_using_clause.
	VisitCav_using_clause(ctx *Cav_using_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dim_by_clause.
	VisitDim_by_clause(ctx *Dim_by_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dim_key.
	VisitDim_key(ctx *Dim_keyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dim_ref.
	VisitDim_ref(ctx *Dim_refContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hier_ref.
	VisitHier_ref(ctx *Hier_refContext) interface{}

	// Visit a parse tree produced by PlSqlParser#measures_clause.
	VisitMeasures_clause(ctx *Measures_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#av_measure.
	VisitAv_measure(ctx *Av_measureContext) interface{}

	// Visit a parse tree produced by PlSqlParser#base_meas_clause.
	VisitBase_meas_clause(ctx *Base_meas_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#meas_aggregate_clause.
	VisitMeas_aggregate_clause(ctx *Meas_aggregate_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#calc_meas_clause.
	VisitCalc_meas_clause(ctx *Calc_meas_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_measure_clause.
	VisitDefault_measure_clause(ctx *Default_measure_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_aggregate_clause.
	VisitDefault_aggregate_clause(ctx *Default_aggregate_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cache_clause.
	VisitCache_clause(ctx *Cache_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cache_specification.
	VisitCache_specification(ctx *Cache_specificationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#levels_clause.
	VisitLevels_clause(ctx *Levels_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#level_specification.
	VisitLevel_specification(ctx *Level_specificationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#level_group_type.
	VisitLevel_group_type(ctx *Level_group_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#fact_columns_clause.
	VisitFact_columns_clause(ctx *Fact_columns_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#qry_transform_clause.
	VisitQry_transform_clause(ctx *Qry_transform_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_attribute_dimension.
	VisitCreate_attribute_dimension(ctx *Create_attribute_dimensionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ad_using_clause.
	VisitAd_using_clause(ctx *Ad_using_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#source_clause.
	VisitSource_clause(ctx *Source_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_path_clause.
	VisitJoin_path_clause(ctx *Join_path_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_condition.
	VisitJoin_condition(ctx *Join_conditionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_condition_item.
	VisitJoin_condition_item(ctx *Join_condition_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#attributes_clause.
	VisitAttributes_clause(ctx *Attributes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ad_attributes_clause.
	VisitAd_attributes_clause(ctx *Ad_attributes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ad_level_clause.
	VisitAd_level_clause(ctx *Ad_level_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#key_clause.
	VisitKey_clause(ctx *Key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alternate_key_clause.
	VisitAlternate_key_clause(ctx *Alternate_key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dim_order_clause.
	VisitDim_order_clause(ctx *Dim_order_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#all_clause.
	VisitAll_clause(ctx *All_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_audit_policy.
	VisitCreate_audit_policy(ctx *Create_audit_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#privilege_audit_clause.
	VisitPrivilege_audit_clause(ctx *Privilege_audit_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#action_audit_clause.
	VisitAction_audit_clause(ctx *Action_audit_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#standard_actions.
	VisitStandard_actions(ctx *Standard_actionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#actions_clause.
	VisitActions_clause(ctx *Actions_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_action.
	VisitObject_action(ctx *Object_actionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#system_action.
	VisitSystem_action(ctx *System_actionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#component_actions.
	VisitComponent_actions(ctx *Component_actionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#component_action.
	VisitComponent_action(ctx *Component_actionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#role_audit_clause.
	VisitRole_audit_clause(ctx *Role_audit_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_controlfile.
	VisitCreate_controlfile(ctx *Create_controlfileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#controlfile_options.
	VisitControlfile_options(ctx *Controlfile_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logfile_clause.
	VisitLogfile_clause(ctx *Logfile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#character_set_clause.
	VisitCharacter_set_clause(ctx *Character_set_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_specification.
	VisitFile_specification(ctx *File_specificationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_diskgroup.
	VisitCreate_diskgroup(ctx *Create_diskgroupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#qualified_disk_clause.
	VisitQualified_disk_clause(ctx *Qualified_disk_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_edition.
	VisitCreate_edition(ctx *Create_editionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_flashback_archive.
	VisitCreate_flashback_archive(ctx *Create_flashback_archiveContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_archive_quota.
	VisitFlashback_archive_quota(ctx *Flashback_archive_quotaContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_archive_retention.
	VisitFlashback_archive_retention(ctx *Flashback_archive_retentionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_hierarchy.
	VisitCreate_hierarchy(ctx *Create_hierarchyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hier_using_clause.
	VisitHier_using_clause(ctx *Hier_using_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#level_hier_clause.
	VisitLevel_hier_clause(ctx *Level_hier_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hier_attrs_clause.
	VisitHier_attrs_clause(ctx *Hier_attrs_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hier_attr_clause.
	VisitHier_attr_clause(ctx *Hier_attr_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hier_attr_name.
	VisitHier_attr_name(ctx *Hier_attr_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_index.
	VisitCreate_index(ctx *Create_indexContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cluster_index_clause.
	VisitCluster_index_clause(ctx *Cluster_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cluster_name.
	VisitCluster_name(ctx *Cluster_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_index_clause.
	VisitTable_index_clause(ctx *Table_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#bitmap_join_index_clause.
	VisitBitmap_join_index_clause(ctx *Bitmap_join_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_expr_option.
	VisitIndex_expr_option(ctx *Index_expr_optionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_expr.
	VisitIndex_expr(ctx *Index_exprContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_properties.
	VisitIndex_properties(ctx *Index_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#domain_index_clause.
	VisitDomain_index_clause(ctx *Domain_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#local_domain_index_clause.
	VisitLocal_domain_index_clause(ctx *Local_domain_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlindex_clause.
	VisitXmlindex_clause(ctx *Xmlindex_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#local_xmlindex_clause.
	VisitLocal_xmlindex_clause(ctx *Local_xmlindex_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#global_partitioned_index.
	VisitGlobal_partitioned_index(ctx *Global_partitioned_indexContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_partitioning_clause.
	VisitIndex_partitioning_clause(ctx *Index_partitioning_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#local_partitioned_index.
	VisitLocal_partitioned_index(ctx *Local_partitioned_indexContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_range_partitioned_table.
	VisitOn_range_partitioned_table(ctx *On_range_partitioned_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_list_partitioned_table.
	VisitOn_list_partitioned_table(ctx *On_list_partitioned_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partitioned_table.
	VisitPartitioned_table(ctx *Partitioned_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_hash_partitioned_table.
	VisitOn_hash_partitioned_table(ctx *On_hash_partitioned_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_hash_partitioned_clause.
	VisitOn_hash_partitioned_clause(ctx *On_hash_partitioned_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_comp_partitioned_table.
	VisitOn_comp_partitioned_table(ctx *On_comp_partitioned_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_comp_partitioned_clause.
	VisitOn_comp_partitioned_clause(ctx *On_comp_partitioned_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_subpartition_clause.
	VisitIndex_subpartition_clause(ctx *Index_subpartition_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_subpartition_subclause.
	VisitIndex_subpartition_subclause(ctx *Index_subpartition_subclauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#odci_parameters.
	VisitOdci_parameters(ctx *Odci_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#indextype.
	VisitIndextype(ctx *IndextypeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_index.
	VisitAlter_index(ctx *Alter_indexContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_index_ops_set1.
	VisitAlter_index_ops_set1(ctx *Alter_index_ops_set1Context) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_index_ops_set2.
	VisitAlter_index_ops_set2(ctx *Alter_index_ops_set2Context) interface{}

	// Visit a parse tree produced by PlSqlParser#visible_or_invisible.
	VisitVisible_or_invisible(ctx *Visible_or_invisibleContext) interface{}

	// Visit a parse tree produced by PlSqlParser#monitoring_nomonitoring.
	VisitMonitoring_nomonitoring(ctx *Monitoring_nomonitoringContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rebuild_clause.
	VisitRebuild_clause(ctx *Rebuild_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_index_partitioning.
	VisitAlter_index_partitioning(ctx *Alter_index_partitioningContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_index_default_attrs.
	VisitModify_index_default_attrs(ctx *Modify_index_default_attrsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_hash_index_partition.
	VisitAdd_hash_index_partition(ctx *Add_hash_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#coalesce_index_partition.
	VisitCoalesce_index_partition(ctx *Coalesce_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_index_partition.
	VisitModify_index_partition(ctx *Modify_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_index_partitions_ops.
	VisitModify_index_partitions_ops(ctx *Modify_index_partitions_opsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rename_index_partition.
	VisitRename_index_partition(ctx *Rename_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_index_partition.
	VisitDrop_index_partition(ctx *Drop_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#split_index_partition.
	VisitSplit_index_partition(ctx *Split_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_partition_description.
	VisitIndex_partition_description(ctx *Index_partition_descriptionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_index_subpartition.
	VisitModify_index_subpartition(ctx *Modify_index_subpartitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_name_old.
	VisitPartition_name_old(ctx *Partition_name_oldContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_partition_name.
	VisitNew_partition_name(ctx *New_partition_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_index_name.
	VisitNew_index_name(ctx *New_index_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_inmemory_join_group.
	VisitAlter_inmemory_join_group(ctx *Alter_inmemory_join_groupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_user.
	VisitCreate_user(ctx *Create_userContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_user.
	VisitAlter_user(ctx *Alter_userContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_user.
	VisitDrop_user(ctx *Drop_userContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_identified_by.
	VisitAlter_identified_by(ctx *Alter_identified_byContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identified_by.
	VisitIdentified_by(ctx *Identified_byContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identified_other_clause.
	VisitIdentified_other_clause(ctx *Identified_other_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_tablespace_clause.
	VisitUser_tablespace_clause(ctx *User_tablespace_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quota_clause.
	VisitQuota_clause(ctx *Quota_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#profile_clause.
	VisitProfile_clause(ctx *Profile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#role_clause.
	VisitRole_clause(ctx *Role_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_default_role_clause.
	VisitUser_default_role_clause(ctx *User_default_role_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#password_expire_clause.
	VisitPassword_expire_clause(ctx *Password_expire_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_lock_clause.
	VisitUser_lock_clause(ctx *User_lock_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_editions_clause.
	VisitUser_editions_clause(ctx *User_editions_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_user_editions_clause.
	VisitAlter_user_editions_clause(ctx *Alter_user_editions_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#proxy_clause.
	VisitProxy_clause(ctx *Proxy_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#container_names.
	VisitContainer_names(ctx *Container_namesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_container_data.
	VisitSet_container_data(ctx *Set_container_dataContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_rem_container_data.
	VisitAdd_rem_container_data(ctx *Add_rem_container_dataContext) interface{}

	// Visit a parse tree produced by PlSqlParser#container_data_clause.
	VisitContainer_data_clause(ctx *Container_data_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#administer_key_management.
	VisitAdminister_key_management(ctx *Administer_key_managementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#keystore_management_clauses.
	VisitKeystore_management_clauses(ctx *Keystore_management_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_keystore.
	VisitCreate_keystore(ctx *Create_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#open_keystore.
	VisitOpen_keystore(ctx *Open_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#force_keystore.
	VisitForce_keystore(ctx *Force_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#close_keystore.
	VisitClose_keystore(ctx *Close_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#backup_keystore.
	VisitBackup_keystore(ctx *Backup_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_keystore_password.
	VisitAlter_keystore_password(ctx *Alter_keystore_passwordContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_into_new_keystore.
	VisitMerge_into_new_keystore(ctx *Merge_into_new_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_into_existing_keystore.
	VisitMerge_into_existing_keystore(ctx *Merge_into_existing_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#isolate_keystore.
	VisitIsolate_keystore(ctx *Isolate_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unite_keystore.
	VisitUnite_keystore(ctx *Unite_keystoreContext) interface{}

	// Visit a parse tree produced by PlSqlParser#key_management_clauses.
	VisitKey_management_clauses(ctx *Key_management_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_key.
	VisitSet_key(ctx *Set_keyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_key.
	VisitCreate_key(ctx *Create_keyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mkid.
	VisitMkid(ctx *MkidContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mk.
	VisitMk(ctx *MkContext) interface{}

	// Visit a parse tree produced by PlSqlParser#use_key.
	VisitUse_key(ctx *Use_keyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_key_tag.
	VisitSet_key_tag(ctx *Set_key_tagContext) interface{}

	// Visit a parse tree produced by PlSqlParser#export_keys.
	VisitExport_keys(ctx *Export_keysContext) interface{}

	// Visit a parse tree produced by PlSqlParser#import_keys.
	VisitImport_keys(ctx *Import_keysContext) interface{}

	// Visit a parse tree produced by PlSqlParser#migrate_keys.
	VisitMigrate_keys(ctx *Migrate_keysContext) interface{}

	// Visit a parse tree produced by PlSqlParser#reverse_migrate_keys.
	VisitReverse_migrate_keys(ctx *Reverse_migrate_keysContext) interface{}

	// Visit a parse tree produced by PlSqlParser#move_keys.
	VisitMove_keys(ctx *Move_keysContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identified_by_store.
	VisitIdentified_by_store(ctx *Identified_by_storeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_algorithm_clause.
	VisitUsing_algorithm_clause(ctx *Using_algorithm_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_tag_clause.
	VisitUsing_tag_clause(ctx *Using_tag_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#secret_management_clauses.
	VisitSecret_management_clauses(ctx *Secret_management_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_update_secret.
	VisitAdd_update_secret(ctx *Add_update_secretContext) interface{}

	// Visit a parse tree produced by PlSqlParser#delete_secret.
	VisitDelete_secret(ctx *Delete_secretContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_update_secret_seps.
	VisitAdd_update_secret_seps(ctx *Add_update_secret_sepsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#delete_secret_seps.
	VisitDelete_secret_seps(ctx *Delete_secret_sepsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#zero_downtime_software_patching_clauses.
	VisitZero_downtime_software_patching_clauses(ctx *Zero_downtime_software_patching_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#with_backup_clause.
	VisitWith_backup_clause(ctx *With_backup_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identified_by_password_clause.
	VisitIdentified_by_password_clause(ctx *Identified_by_password_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#keystore_password.
	VisitKeystore_password(ctx *Keystore_passwordContext) interface{}

	// Visit a parse tree produced by PlSqlParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by PlSqlParser#secret.
	VisitSecret(ctx *SecretContext) interface{}

	// Visit a parse tree produced by PlSqlParser#analyze.
	VisitAnalyze(ctx *AnalyzeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_extention_clause.
	VisitPartition_extention_clause(ctx *Partition_extention_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#validation_clauses.
	VisitValidation_clauses(ctx *Validation_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compute_clauses.
	VisitCompute_clauses(ctx *Compute_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#for_clause.
	VisitFor_clause(ctx *For_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#online_or_offline.
	VisitOnline_or_offline(ctx *Online_or_offlineContext) interface{}

	// Visit a parse tree produced by PlSqlParser#into_clause1.
	VisitInto_clause1(ctx *Into_clause1Context) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_key_value.
	VisitPartition_key_value(ctx *Partition_key_valueContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_key_value.
	VisitSubpartition_key_value(ctx *Subpartition_key_valueContext) interface{}

	// Visit a parse tree produced by PlSqlParser#associate_statistics.
	VisitAssociate_statistics(ctx *Associate_statisticsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_association.
	VisitColumn_association(ctx *Column_associationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_association.
	VisitFunction_association(ctx *Function_associationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#indextype_name.
	VisitIndextype_name(ctx *Indextype_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_statistics_type.
	VisitUsing_statistics_type(ctx *Using_statistics_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#statistics_type_name.
	VisitStatistics_type_name(ctx *Statistics_type_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_cost_clause.
	VisitDefault_cost_clause(ctx *Default_cost_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cpu_cost.
	VisitCpu_cost(ctx *Cpu_costContext) interface{}

	// Visit a parse tree produced by PlSqlParser#io_cost.
	VisitIo_cost(ctx *Io_costContext) interface{}

	// Visit a parse tree produced by PlSqlParser#network_cost.
	VisitNetwork_cost(ctx *Network_costContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_selectivity_clause.
	VisitDefault_selectivity_clause(ctx *Default_selectivity_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_selectivity.
	VisitDefault_selectivity(ctx *Default_selectivityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#storage_table_clause.
	VisitStorage_table_clause(ctx *Storage_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unified_auditing.
	VisitUnified_auditing(ctx *Unified_auditingContext) interface{}

	// Visit a parse tree produced by PlSqlParser#policy_name.
	VisitPolicy_name(ctx *Policy_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_traditional.
	VisitAudit_traditional(ctx *Audit_traditionalContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_direct_path.
	VisitAudit_direct_path(ctx *Audit_direct_pathContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_container_clause.
	VisitAudit_container_clause(ctx *Audit_container_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_operation_clause.
	VisitAudit_operation_clause(ctx *Audit_operation_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#auditing_by_clause.
	VisitAuditing_by_clause(ctx *Auditing_by_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_user.
	VisitAudit_user(ctx *Audit_userContext) interface{}

	// Visit a parse tree produced by PlSqlParser#audit_schema_object_clause.
	VisitAudit_schema_object_clause(ctx *Audit_schema_object_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sql_operation.
	VisitSql_operation(ctx *Sql_operationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#auditing_on_clause.
	VisitAuditing_on_clause(ctx *Auditing_on_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_name.
	VisitModel_name(ctx *Model_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_name.
	VisitObject_name(ctx *Object_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#profile_name.
	VisitProfile_name(ctx *Profile_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sql_statement_shortcut.
	VisitSql_statement_shortcut(ctx *Sql_statement_shortcutContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_index.
	VisitDrop_index(ctx *Drop_indexContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disassociate_statistics.
	VisitDisassociate_statistics(ctx *Disassociate_statisticsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_indextype.
	VisitDrop_indextype(ctx *Drop_indextypeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_inmemory_join_group.
	VisitDrop_inmemory_join_group(ctx *Drop_inmemory_join_groupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_table.
	VisitFlashback_table(ctx *Flashback_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#restore_point.
	VisitRestore_point(ctx *Restore_pointContext) interface{}

	// Visit a parse tree produced by PlSqlParser#purge_statement.
	VisitPurge_statement(ctx *Purge_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#noaudit_statement.
	VisitNoaudit_statement(ctx *Noaudit_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rename_object.
	VisitRename_object(ctx *Rename_objectContext) interface{}

	// Visit a parse tree produced by PlSqlParser#grant_statement.
	VisitGrant_statement(ctx *Grant_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#container_clause.
	VisitContainer_clause(ctx *Container_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#revoke_statement.
	VisitRevoke_statement(ctx *Revoke_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#revoke_system_privilege.
	VisitRevoke_system_privilege(ctx *Revoke_system_privilegeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#revokee_clause.
	VisitRevokee_clause(ctx *Revokee_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#revoke_object_privileges.
	VisitRevoke_object_privileges(ctx *Revoke_object_privilegesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_object_clause.
	VisitOn_object_clause(ctx *On_object_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#revoke_roles_from_programs.
	VisitRevoke_roles_from_programs(ctx *Revoke_roles_from_programsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#program_unit.
	VisitProgram_unit(ctx *Program_unitContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_dimension.
	VisitCreate_dimension(ctx *Create_dimensionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_directory.
	VisitCreate_directory(ctx *Create_directoryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#directory_name.
	VisitDirectory_name(ctx *Directory_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#directory_path.
	VisitDirectory_path(ctx *Directory_pathContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_inmemory_join_group.
	VisitCreate_inmemory_join_group(ctx *Create_inmemory_join_groupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_hierarchy.
	VisitDrop_hierarchy(ctx *Drop_hierarchyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_library.
	VisitAlter_library(ctx *Alter_libraryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_java.
	VisitDrop_java(ctx *Drop_javaContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_library.
	VisitDrop_library(ctx *Drop_libraryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_java.
	VisitCreate_java(ctx *Create_javaContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_library.
	VisitCreate_library(ctx *Create_libraryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#plsql_library_source.
	VisitPlsql_library_source(ctx *Plsql_library_sourceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#credential_name.
	VisitCredential_name(ctx *Credential_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#library_editionable.
	VisitLibrary_editionable(ctx *Library_editionableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#library_debug.
	VisitLibrary_debug(ctx *Library_debugContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compiler_parameters_clause.
	VisitCompiler_parameters_clause(ctx *Compiler_parameters_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parameter_value.
	VisitParameter_value(ctx *Parameter_valueContext) interface{}

	// Visit a parse tree produced by PlSqlParser#library_name.
	VisitLibrary_name(ctx *Library_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_dimension.
	VisitAlter_dimension(ctx *Alter_dimensionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#level_clause.
	VisitLevel_clause(ctx *Level_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hierarchy_clause.
	VisitHierarchy_clause(ctx *Hierarchy_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dimension_join_clause.
	VisitDimension_join_clause(ctx *Dimension_join_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#attribute_clause.
	VisitAttribute_clause(ctx *Attribute_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#extended_attribute_clause.
	VisitExtended_attribute_clause(ctx *Extended_attribute_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_one_or_more_sub_clause.
	VisitColumn_one_or_more_sub_clause(ctx *Column_one_or_more_sub_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_view.
	VisitAlter_view(ctx *Alter_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_view_editionable.
	VisitAlter_view_editionable(ctx *Alter_view_editionableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_view.
	VisitCreate_view(ctx *Create_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#editioning_clause.
	VisitEditioning_clause(ctx *Editioning_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#view_options.
	VisitView_options(ctx *View_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#view_alias_constraint.
	VisitView_alias_constraint(ctx *View_alias_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_view_clause.
	VisitObject_view_clause(ctx *Object_view_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inline_constraint.
	VisitInline_constraint(ctx *Inline_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inline_ref_constraint.
	VisitInline_ref_constraint(ctx *Inline_ref_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#out_of_line_ref_constraint.
	VisitOut_of_line_ref_constraint(ctx *Out_of_line_ref_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#out_of_line_constraint.
	VisitOut_of_line_constraint(ctx *Out_of_line_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constraint_state.
	VisitConstraint_state(ctx *Constraint_stateContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltype_view_clause.
	VisitXmltype_view_clause(ctx *Xmltype_view_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_schema_spec.
	VisitXml_schema_spec(ctx *Xml_schema_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_schema_url.
	VisitXml_schema_url(ctx *Xml_schema_urlContext) interface{}

	// Visit a parse tree produced by PlSqlParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_tablespace.
	VisitAlter_tablespace(ctx *Alter_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datafile_tempfile_clauses.
	VisitDatafile_tempfile_clauses(ctx *Datafile_tempfile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_logging_clauses.
	VisitTablespace_logging_clauses(ctx *Tablespace_logging_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_group_clause.
	VisitTablespace_group_clause(ctx *Tablespace_group_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_group_name.
	VisitTablespace_group_name(ctx *Tablespace_group_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_state_clauses.
	VisitTablespace_state_clauses(ctx *Tablespace_state_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_mode_clause.
	VisitFlashback_mode_clause(ctx *Flashback_mode_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_tablespace_name.
	VisitNew_tablespace_name(ctx *New_tablespace_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_tablespace.
	VisitCreate_tablespace(ctx *Create_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#permanent_tablespace_clause.
	VisitPermanent_tablespace_clause(ctx *Permanent_tablespace_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_encryption_spec.
	VisitTablespace_encryption_spec(ctx *Tablespace_encryption_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logging_clause.
	VisitLogging_clause(ctx *Logging_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#extent_management_clause.
	VisitExtent_management_clause(ctx *Extent_management_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#segment_management_clause.
	VisitSegment_management_clause(ctx *Segment_management_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#temporary_tablespace_clause.
	VisitTemporary_tablespace_clause(ctx *Temporary_tablespace_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#undo_tablespace_clause.
	VisitUndo_tablespace_clause(ctx *Undo_tablespace_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_retention_clause.
	VisitTablespace_retention_clause(ctx *Tablespace_retention_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_tablespace_set.
	VisitCreate_tablespace_set(ctx *Create_tablespace_setContext) interface{}

	// Visit a parse tree produced by PlSqlParser#permanent_tablespace_attrs.
	VisitPermanent_tablespace_attrs(ctx *Permanent_tablespace_attrsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_encryption_clause.
	VisitTablespace_encryption_clause(ctx *Tablespace_encryption_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_tablespace_params.
	VisitDefault_tablespace_params(ctx *Default_tablespace_paramsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_table_compression.
	VisitDefault_table_compression(ctx *Default_table_compressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#low_high.
	VisitLow_high(ctx *Low_highContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_index_compression.
	VisitDefault_index_compression(ctx *Default_index_compressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmmemory_clause.
	VisitInmmemory_clause(ctx *Inmmemory_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datafile_specification.
	VisitDatafile_specification(ctx *Datafile_specificationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tempfile_specification.
	VisitTempfile_specification(ctx *Tempfile_specificationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datafile_tempfile_spec.
	VisitDatafile_tempfile_spec(ctx *Datafile_tempfile_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#redo_log_file_spec.
	VisitRedo_log_file_spec(ctx *Redo_log_file_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#autoextend_clause.
	VisitAutoextend_clause(ctx *Autoextend_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#maxsize_clause.
	VisitMaxsize_clause(ctx *Maxsize_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#build_clause.
	VisitBuild_clause(ctx *Build_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partial_index_clause.
	VisitPartial_index_clause(ctx *Partial_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parallel_clause.
	VisitParallel_clause(ctx *Parallel_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_materialized_view.
	VisitAlter_materialized_view(ctx *Alter_materialized_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_mv_option1.
	VisitAlter_mv_option1(ctx *Alter_mv_option1Context) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_mv_refresh.
	VisitAlter_mv_refresh(ctx *Alter_mv_refreshContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rollback_segment.
	VisitRollback_segment(ctx *Rollback_segmentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_mv_column_clause.
	VisitModify_mv_column_clause(ctx *Modify_mv_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_materialized_view_log.
	VisitAlter_materialized_view_log(ctx *Alter_materialized_view_logContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_mv_log_column_clause.
	VisitAdd_mv_log_column_clause(ctx *Add_mv_log_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#move_mv_log_clause.
	VisitMove_mv_log_clause(ctx *Move_mv_log_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mv_log_augmentation.
	VisitMv_log_augmentation(ctx *Mv_log_augmentationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datetime_expr.
	VisitDatetime_expr(ctx *Datetime_exprContext) interface{}

	// Visit a parse tree produced by PlSqlParser#interval_expr.
	VisitInterval_expr(ctx *Interval_exprContext) interface{}

	// Visit a parse tree produced by PlSqlParser#synchronous_or_asynchronous.
	VisitSynchronous_or_asynchronous(ctx *Synchronous_or_asynchronousContext) interface{}

	// Visit a parse tree produced by PlSqlParser#including_or_excluding.
	VisitIncluding_or_excluding(ctx *Including_or_excludingContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_materialized_view_log.
	VisitCreate_materialized_view_log(ctx *Create_materialized_view_logContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_values_clause.
	VisitNew_values_clause(ctx *New_values_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mv_log_purge_clause.
	VisitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_materialized_zonemap.
	VisitCreate_materialized_zonemap(ctx *Create_materialized_zonemapContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_materialized_zonemap.
	VisitAlter_materialized_zonemap(ctx *Alter_materialized_zonemapContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_materialized_zonemap.
	VisitDrop_materialized_zonemap(ctx *Drop_materialized_zonemapContext) interface{}

	// Visit a parse tree produced by PlSqlParser#zonemap_refresh_clause.
	VisitZonemap_refresh_clause(ctx *Zonemap_refresh_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#zonemap_attributes.
	VisitZonemap_attributes(ctx *Zonemap_attributesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#zonemap_name.
	VisitZonemap_name(ctx *Zonemap_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#operator_name.
	VisitOperator_name(ctx *Operator_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#operator_function_name.
	VisitOperator_function_name(ctx *Operator_function_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_zonemap_on_table.
	VisitCreate_zonemap_on_table(ctx *Create_zonemap_on_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_zonemap_as_subquery.
	VisitCreate_zonemap_as_subquery(ctx *Create_zonemap_as_subqueryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_operator.
	VisitAlter_operator(ctx *Alter_operatorContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_operator.
	VisitDrop_operator(ctx *Drop_operatorContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_operator.
	VisitCreate_operator(ctx *Create_operatorContext) interface{}

	// Visit a parse tree produced by PlSqlParser#binding_clause.
	VisitBinding_clause(ctx *Binding_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_binding_clause.
	VisitAdd_binding_clause(ctx *Add_binding_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#implementation_clause.
	VisitImplementation_clause(ctx *Implementation_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#primary_operator_list.
	VisitPrimary_operator_list(ctx *Primary_operator_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#primary_operator_item.
	VisitPrimary_operator_item(ctx *Primary_operator_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#operator_context_clause.
	VisitOperator_context_clause(ctx *Operator_context_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_function_clause.
	VisitUsing_function_clause(ctx *Using_function_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_binding_clause.
	VisitDrop_binding_clause(ctx *Drop_binding_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_materialized_view.
	VisitCreate_materialized_view(ctx *Create_materialized_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#scoped_table_ref_constraint.
	VisitScoped_table_ref_constraint(ctx *Scoped_table_ref_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mv_column_alias.
	VisitMv_column_alias(ctx *Mv_column_aliasContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_mv_refresh.
	VisitCreate_mv_refresh(ctx *Create_mv_refreshContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_materialized_view.
	VisitDrop_materialized_view(ctx *Drop_materialized_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_context.
	VisitCreate_context(ctx *Create_contextContext) interface{}

	// Visit a parse tree produced by PlSqlParser#oracle_namespace.
	VisitOracle_namespace(ctx *Oracle_namespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_cluster.
	VisitCreate_cluster(ctx *Create_clusterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_profile.
	VisitCreate_profile(ctx *Create_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#resource_parameters.
	VisitResource_parameters(ctx *Resource_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#password_parameters.
	VisitPassword_parameters(ctx *Password_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_lockdown_profile.
	VisitCreate_lockdown_profile(ctx *Create_lockdown_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#static_base_profile.
	VisitStatic_base_profile(ctx *Static_base_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dynamic_base_profile.
	VisitDynamic_base_profile(ctx *Dynamic_base_profileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_outline.
	VisitCreate_outline(ctx *Create_outlineContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_restore_point.
	VisitCreate_restore_point(ctx *Create_restore_pointContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_role.
	VisitCreate_role(ctx *Create_roleContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_table.
	VisitCreate_table(ctx *Create_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltype_table.
	VisitXmltype_table(ctx *Xmltype_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltype_virtual_columns.
	VisitXmltype_virtual_columns(ctx *Xmltype_virtual_columnsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltype_column_properties.
	VisitXmltype_column_properties(ctx *Xmltype_column_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltype_storage.
	VisitXmltype_storage(ctx *Xmltype_storageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlschema_spec.
	VisitXmlschema_spec(ctx *Xmlschema_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_table.
	VisitObject_table(ctx *Object_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_type.
	VisitObject_type(ctx *Object_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#oid_index_clause.
	VisitOid_index_clause(ctx *Oid_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#oid_clause.
	VisitOid_clause(ctx *Oid_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_properties.
	VisitObject_properties(ctx *Object_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_table_substitution.
	VisitObject_table_substitution(ctx *Object_table_substitutionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#relational_table.
	VisitRelational_table(ctx *Relational_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#immutable_table_clauses.
	VisitImmutable_table_clauses(ctx *Immutable_table_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#immutable_table_no_drop_clause.
	VisitImmutable_table_no_drop_clause(ctx *Immutable_table_no_drop_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#immutable_table_no_delete_clause.
	VisitImmutable_table_no_delete_clause(ctx *Immutable_table_no_delete_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#blockchain_table_clauses.
	VisitBlockchain_table_clauses(ctx *Blockchain_table_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#blockchain_drop_table_clause.
	VisitBlockchain_drop_table_clause(ctx *Blockchain_drop_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#blockchain_row_retention_clause.
	VisitBlockchain_row_retention_clause(ctx *Blockchain_row_retention_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#blockchain_hash_and_data_format_clause.
	VisitBlockchain_hash_and_data_format_clause(ctx *Blockchain_hash_and_data_format_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#collation_name.
	VisitCollation_name(ctx *Collation_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_properties.
	VisitTable_properties(ctx *Table_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#read_only_clause.
	VisitRead_only_clause(ctx *Read_only_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#indexing_clause.
	VisitIndexing_clause(ctx *Indexing_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#attribute_clustering_clause.
	VisitAttribute_clustering_clause(ctx *Attribute_clustering_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#clustering_join.
	VisitClustering_join(ctx *Clustering_joinContext) interface{}

	// Visit a parse tree produced by PlSqlParser#clustering_join_item.
	VisitClustering_join_item(ctx *Clustering_join_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#equijoin_condition.
	VisitEquijoin_condition(ctx *Equijoin_conditionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cluster_clause.
	VisitCluster_clause(ctx *Cluster_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#clustering_columns.
	VisitClustering_columns(ctx *Clustering_columnsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#clustering_column_group.
	VisitClustering_column_group(ctx *Clustering_column_groupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#yes_no.
	VisitYes_no(ctx *Yes_noContext) interface{}

	// Visit a parse tree produced by PlSqlParser#zonemap_clause.
	VisitZonemap_clause(ctx *Zonemap_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logical_replication_clause.
	VisitLogical_replication_clause(ctx *Logical_replication_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#relational_property.
	VisitRelational_property(ctx *Relational_propertyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_partitioning_clauses.
	VisitTable_partitioning_clauses(ctx *Table_partitioning_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#range_partitions.
	VisitRange_partitions(ctx *Range_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#list_partitions.
	VisitList_partitions(ctx *List_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hash_partitions.
	VisitHash_partitions(ctx *Hash_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#individual_hash_partitions.
	VisitIndividual_hash_partitions(ctx *Individual_hash_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hash_partitions_by_quantity.
	VisitHash_partitions_by_quantity(ctx *Hash_partitions_by_quantityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hash_partition_quantity.
	VisitHash_partition_quantity(ctx *Hash_partition_quantityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#composite_range_partitions.
	VisitComposite_range_partitions(ctx *Composite_range_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#composite_list_partitions.
	VisitComposite_list_partitions(ctx *Composite_list_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#composite_hash_partitions.
	VisitComposite_hash_partitions(ctx *Composite_hash_partitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#reference_partitioning.
	VisitReference_partitioning(ctx *Reference_partitioningContext) interface{}

	// Visit a parse tree produced by PlSqlParser#reference_partition_desc.
	VisitReference_partition_desc(ctx *Reference_partition_descContext) interface{}

	// Visit a parse tree produced by PlSqlParser#system_partitioning.
	VisitSystem_partitioning(ctx *System_partitioningContext) interface{}

	// Visit a parse tree produced by PlSqlParser#range_partition_desc.
	VisitRange_partition_desc(ctx *Range_partition_descContext) interface{}

	// Visit a parse tree produced by PlSqlParser#list_partition_desc.
	VisitList_partition_desc(ctx *List_partition_descContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_template.
	VisitSubpartition_template(ctx *Subpartition_templateContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hash_subpartition_quantity.
	VisitHash_subpartition_quantity(ctx *Hash_subpartition_quantityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_by_range.
	VisitSubpartition_by_range(ctx *Subpartition_by_rangeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_by_list.
	VisitSubpartition_by_list(ctx *Subpartition_by_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_by_hash.
	VisitSubpartition_by_hash(ctx *Subpartition_by_hashContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_name.
	VisitSubpartition_name(ctx *Subpartition_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#range_subpartition_desc.
	VisitRange_subpartition_desc(ctx *Range_subpartition_descContext) interface{}

	// Visit a parse tree produced by PlSqlParser#list_subpartition_desc.
	VisitList_subpartition_desc(ctx *List_subpartition_descContext) interface{}

	// Visit a parse tree produced by PlSqlParser#individual_hash_subparts.
	VisitIndividual_hash_subparts(ctx *Individual_hash_subpartsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hash_subparts_by_quantity.
	VisitHash_subparts_by_quantity(ctx *Hash_subparts_by_quantityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#range_values_clause.
	VisitRange_values_clause(ctx *Range_values_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#list_values_clause.
	VisitList_values_clause(ctx *List_values_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_partition_description.
	VisitTable_partition_description(ctx *Table_partition_descriptionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partitioning_storage_clause.
	VisitPartitioning_storage_clause(ctx *Partitioning_storage_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_partitioning_storage.
	VisitLob_partitioning_storage(ctx *Lob_partitioning_storageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datatype_null_enable.
	VisitDatatype_null_enable(ctx *Datatype_null_enableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#size_clause.
	VisitSize_clause(ctx *Size_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_compression.
	VisitTable_compression(ctx *Table_compressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_table_clause.
	VisitInmemory_table_clause(ctx *Inmemory_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_attributes.
	VisitInmemory_attributes(ctx *Inmemory_attributesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_memcompress.
	VisitInmemory_memcompress(ctx *Inmemory_memcompressContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_priority.
	VisitInmemory_priority(ctx *Inmemory_priorityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_distribute.
	VisitInmemory_distribute(ctx *Inmemory_distributeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_duplicate.
	VisitInmemory_duplicate(ctx *Inmemory_duplicateContext) interface{}

	// Visit a parse tree produced by PlSqlParser#inmemory_column_clause.
	VisitInmemory_column_clause(ctx *Inmemory_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#physical_attributes_clause.
	VisitPhysical_attributes_clause(ctx *Physical_attributes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#storage_clause.
	VisitStorage_clause(ctx *Storage_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#deferred_segment_creation.
	VisitDeferred_segment_creation(ctx *Deferred_segment_creationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#segment_attributes_clause.
	VisitSegment_attributes_clause(ctx *Segment_attributes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#physical_properties.
	VisitPhysical_properties(ctx *Physical_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_clause.
	VisitIlm_clause(ctx *Ilm_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_policy_clause.
	VisitIlm_policy_clause(ctx *Ilm_policy_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_compression_policy.
	VisitIlm_compression_policy(ctx *Ilm_compression_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_tiering_policy.
	VisitIlm_tiering_policy(ctx *Ilm_tiering_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_after_on.
	VisitIlm_after_on(ctx *Ilm_after_onContext) interface{}

	// Visit a parse tree produced by PlSqlParser#segment_group.
	VisitSegment_group(ctx *Segment_groupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_inmemory_policy.
	VisitIlm_inmemory_policy(ctx *Ilm_inmemory_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ilm_time_period.
	VisitIlm_time_period(ctx *Ilm_time_periodContext) interface{}

	// Visit a parse tree produced by PlSqlParser#heap_org_table_clause.
	VisitHeap_org_table_clause(ctx *Heap_org_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#external_table_clause.
	VisitExternal_table_clause(ctx *External_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#access_driver_type.
	VisitAccess_driver_type(ctx *Access_driver_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#external_table_data_props.
	VisitExternal_table_data_props(ctx *External_table_data_propsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#opaque_format_spec.
	VisitOpaque_format_spec(ctx *Opaque_format_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#record_format_info.
	VisitRecord_format_info(ctx *Record_format_infoContext) interface{}

	// Visit a parse tree produced by PlSqlParser#et_record_spec_options.
	VisitEt_record_spec_options(ctx *Et_record_spec_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#et_record_spec_option.
	VisitEt_record_spec_option(ctx *Et_record_spec_optionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#et_output_files.
	VisitEt_output_files(ctx *Et_output_filesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#et_output_file.
	VisitEt_output_file(ctx *Et_output_fileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#directory_spec.
	VisitDirectory_spec(ctx *Directory_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_spec.
	VisitFile_spec(ctx *File_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_definitions.
	VisitField_definitions(ctx *Field_definitionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_options.
	VisitField_options(ctx *Field_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_option.
	VisitField_option(ctx *Field_optionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_list.
	VisitField_list(ctx *Field_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_item.
	VisitField_item(ctx *Field_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_name.
	VisitField_name(ctx *Field_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pos_spec.
	VisitPos_spec(ctx *Pos_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pos_start.
	VisitPos_start(ctx *Pos_startContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pos_increment.
	VisitPos_increment(ctx *Pos_incrementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pos_end.
	VisitPos_end(ctx *Pos_endContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pos_length.
	VisitPos_length(ctx *Pos_lengthContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datatype_spec.
	VisitDatatype_spec(ctx *Datatype_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#init_spec.
	VisitInit_spec(ctx *Init_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lls_clause.
	VisitLls_clause(ctx *Lls_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#delim_spec.
	VisitDelim_spec(ctx *Delim_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trim_spec.
	VisitTrim_spec(ctx *Trim_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_date_format.
	VisitField_date_format(ctx *Field_date_formatContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_transforms.
	VisitColumn_transforms(ctx *Column_transformsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#transform.
	VisitTransform(ctx *TransformContext) interface{}

	// Visit a parse tree produced by PlSqlParser#source_field.
	VisitSource_field(ctx *Source_fieldContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lobfile_item.
	VisitLobfile_item(ctx *Lobfile_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lobfile_attr_list.
	VisitLobfile_attr_list(ctx *Lobfile_attr_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#concat_item.
	VisitConcat_item(ctx *Concat_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#row_movement_clause.
	VisitRow_movement_clause(ctx *Row_movement_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_archive_clause.
	VisitFlashback_archive_clause(ctx *Flashback_archive_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#log_grp.
	VisitLog_grp(ctx *Log_grpContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_table_logging.
	VisitSupplemental_table_logging(ctx *Supplemental_table_loggingContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_log_grp_clause.
	VisitSupplemental_log_grp_clause(ctx *Supplemental_log_grp_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_id_key_clause.
	VisitSupplemental_id_key_clause(ctx *Supplemental_id_key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#allocate_extent_clause.
	VisitAllocate_extent_clause(ctx *Allocate_extent_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#deallocate_unused_clause.
	VisitDeallocate_unused_clause(ctx *Deallocate_unused_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#shrink_clause.
	VisitShrink_clause(ctx *Shrink_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#records_per_block_clause.
	VisitRecords_per_block_clause(ctx *Records_per_block_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#upgrade_table_clause.
	VisitUpgrade_table_clause(ctx *Upgrade_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#truncate_table.
	VisitTruncate_table(ctx *Truncate_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_table.
	VisitDrop_table(ctx *Drop_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_tablespace.
	VisitDrop_tablespace(ctx *Drop_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_tablespace_set.
	VisitDrop_tablespace_set(ctx *Drop_tablespace_setContext) interface{}

	// Visit a parse tree produced by PlSqlParser#including_contents_clause.
	VisitIncluding_contents_clause(ctx *Including_contents_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_view.
	VisitDrop_view(ctx *Drop_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#comment_on_column.
	VisitComment_on_column(ctx *Comment_on_columnContext) interface{}

	// Visit a parse tree produced by PlSqlParser#enable_or_disable.
	VisitEnable_or_disable(ctx *Enable_or_disableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#allow_or_disallow.
	VisitAllow_or_disallow(ctx *Allow_or_disallowContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_synonym.
	VisitAlter_synonym(ctx *Alter_synonymContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_synonym.
	VisitCreate_synonym(ctx *Create_synonymContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_synonym.
	VisitDrop_synonym(ctx *Drop_synonymContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_spfile.
	VisitCreate_spfile(ctx *Create_spfileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#spfile_name.
	VisitSpfile_name(ctx *Spfile_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pfile_name.
	VisitPfile_name(ctx *Pfile_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#comment_on_table.
	VisitComment_on_table(ctx *Comment_on_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#comment_on_materialized.
	VisitComment_on_materialized(ctx *Comment_on_materializedContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_analytic_view.
	VisitAlter_analytic_view(ctx *Alter_analytic_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_add_cache_clause.
	VisitAlter_add_cache_clause(ctx *Alter_add_cache_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#levels_item.
	VisitLevels_item(ctx *Levels_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#measure_list.
	VisitMeasure_list(ctx *Measure_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_drop_cache_clause.
	VisitAlter_drop_cache_clause(ctx *Alter_drop_cache_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_attribute_dimension.
	VisitAlter_attribute_dimension(ctx *Alter_attribute_dimensionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_audit_policy.
	VisitAlter_audit_policy(ctx *Alter_audit_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_cluster.
	VisitAlter_cluster(ctx *Alter_clusterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_analytic_view.
	VisitDrop_analytic_view(ctx *Drop_analytic_viewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_attribute_dimension.
	VisitDrop_attribute_dimension(ctx *Drop_attribute_dimensionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_audit_policy.
	VisitDrop_audit_policy(ctx *Drop_audit_policyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_flashback_archive.
	VisitDrop_flashback_archive(ctx *Drop_flashback_archiveContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_cluster.
	VisitDrop_cluster(ctx *Drop_clusterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_context.
	VisitDrop_context(ctx *Drop_contextContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_directory.
	VisitDrop_directory(ctx *Drop_directoryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_diskgroup.
	VisitDrop_diskgroup(ctx *Drop_diskgroupContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_edition.
	VisitDrop_edition(ctx *Drop_editionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#truncate_cluster.
	VisitTruncate_cluster(ctx *Truncate_clusterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cache_or_nocache.
	VisitCache_or_nocache(ctx *Cache_or_nocacheContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_database.
	VisitAlter_database(ctx *Alter_databaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database_clause.
	VisitDatabase_clause(ctx *Database_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#startup_clauses.
	VisitStartup_clauses(ctx *Startup_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#resetlogs_or_noresetlogs.
	VisitResetlogs_or_noresetlogs(ctx *Resetlogs_or_noresetlogsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#upgrade_or_downgrade.
	VisitUpgrade_or_downgrade(ctx *Upgrade_or_downgradeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#recovery_clauses.
	VisitRecovery_clauses(ctx *Recovery_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#begin_or_end.
	VisitBegin_or_end(ctx *Begin_or_endContext) interface{}

	// Visit a parse tree produced by PlSqlParser#general_recovery.
	VisitGeneral_recovery(ctx *General_recoveryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#full_database_recovery.
	VisitFull_database_recovery(ctx *Full_database_recoveryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partial_database_recovery.
	VisitPartial_database_recovery(ctx *Partial_database_recoveryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partial_database_recovery_10g.
	VisitPartial_database_recovery_10g(ctx *Partial_database_recovery_10gContext) interface{}

	// Visit a parse tree produced by PlSqlParser#managed_standby_recovery.
	VisitManaged_standby_recovery(ctx *Managed_standby_recoveryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#db_name.
	VisitDb_name(ctx *Db_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database_file_clauses.
	VisitDatabase_file_clauses(ctx *Database_file_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_datafile_clause.
	VisitCreate_datafile_clause(ctx *Create_datafile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_datafile_clause.
	VisitAlter_datafile_clause(ctx *Alter_datafile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_tempfile_clause.
	VisitAlter_tempfile_clause(ctx *Alter_tempfile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#move_datafile_clause.
	VisitMove_datafile_clause(ctx *Move_datafile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logfile_clauses.
	VisitLogfile_clauses(ctx *Logfile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_logfile_clauses.
	VisitAdd_logfile_clauses(ctx *Add_logfile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#group_redo_logfile.
	VisitGroup_redo_logfile(ctx *Group_redo_logfileContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_logfile_clauses.
	VisitDrop_logfile_clauses(ctx *Drop_logfile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#switch_logfile_clause.
	VisitSwitch_logfile_clause(ctx *Switch_logfile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_db_logging.
	VisitSupplemental_db_logging(ctx *Supplemental_db_loggingContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_or_drop.
	VisitAdd_or_drop(ctx *Add_or_dropContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_plsql_clause.
	VisitSupplemental_plsql_clause(ctx *Supplemental_plsql_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logfile_descriptor.
	VisitLogfile_descriptor(ctx *Logfile_descriptorContext) interface{}

	// Visit a parse tree produced by PlSqlParser#controlfile_clauses.
	VisitControlfile_clauses(ctx *Controlfile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trace_file_clause.
	VisitTrace_file_clause(ctx *Trace_file_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#standby_database_clauses.
	VisitStandby_database_clauses(ctx *Standby_database_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#activate_standby_db_clause.
	VisitActivate_standby_db_clause(ctx *Activate_standby_db_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#maximize_standby_db_clause.
	VisitMaximize_standby_db_clause(ctx *Maximize_standby_db_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#register_logfile_clause.
	VisitRegister_logfile_clause(ctx *Register_logfile_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#commit_switchover_clause.
	VisitCommit_switchover_clause(ctx *Commit_switchover_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#start_standby_clause.
	VisitStart_standby_clause(ctx *Start_standby_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#stop_standby_clause.
	VisitStop_standby_clause(ctx *Stop_standby_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#convert_database_clause.
	VisitConvert_database_clause(ctx *Convert_database_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_settings_clause.
	VisitDefault_settings_clause(ctx *Default_settings_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_time_zone_clause.
	VisitSet_time_zone_clause(ctx *Set_time_zone_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#instance_clauses.
	VisitInstance_clauses(ctx *Instance_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#security_clause.
	VisitSecurity_clause(ctx *Security_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#domain.
	VisitDomain(ctx *DomainContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database.
	VisitDatabase(ctx *DatabaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#edition_name.
	VisitEdition_name(ctx *Edition_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#filenumber.
	VisitFilenumber(ctx *FilenumberContext) interface{}

	// Visit a parse tree produced by PlSqlParser#filename.
	VisitFilename(ctx *FilenameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#prepare_clause.
	VisitPrepare_clause(ctx *Prepare_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_mirror_clause.
	VisitDrop_mirror_clause(ctx *Drop_mirror_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lost_write_protection.
	VisitLost_write_protection(ctx *Lost_write_protectionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cdb_fleet_clauses.
	VisitCdb_fleet_clauses(ctx *Cdb_fleet_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lead_cdb_clause.
	VisitLead_cdb_clause(ctx *Lead_cdb_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lead_cdb_uri_clause.
	VisitLead_cdb_uri_clause(ctx *Lead_cdb_uri_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#property_clauses.
	VisitProperty_clauses(ctx *Property_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#replay_upgrade_clauses.
	VisitReplay_upgrade_clauses(ctx *Replay_upgrade_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_database_link.
	VisitAlter_database_link(ctx *Alter_database_linkContext) interface{}

	// Visit a parse tree produced by PlSqlParser#password_value.
	VisitPassword_value(ctx *Password_valueContext) interface{}

	// Visit a parse tree produced by PlSqlParser#link_authentication.
	VisitLink_authentication(ctx *Link_authenticationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_database.
	VisitCreate_database(ctx *Create_databaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database_logging_clauses.
	VisitDatabase_logging_clauses(ctx *Database_logging_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#database_logging_sub_clause.
	VisitDatabase_logging_sub_clause(ctx *Database_logging_sub_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_clauses.
	VisitTablespace_clauses(ctx *Tablespace_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#enable_pluggable_database.
	VisitEnable_pluggable_database(ctx *Enable_pluggable_databaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#file_name_convert.
	VisitFile_name_convert(ctx *File_name_convertContext) interface{}

	// Visit a parse tree produced by PlSqlParser#filename_convert_sub_clause.
	VisitFilename_convert_sub_clause(ctx *Filename_convert_sub_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace_datafile_clauses.
	VisitTablespace_datafile_clauses(ctx *Tablespace_datafile_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#undo_mode_clause.
	VisitUndo_mode_clause(ctx *Undo_mode_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_tablespace.
	VisitDefault_tablespace(ctx *Default_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_temp_tablespace.
	VisitDefault_temp_tablespace(ctx *Default_temp_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#undo_tablespace.
	VisitUndo_tablespace(ctx *Undo_tablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_database.
	VisitDrop_database(ctx *Drop_databaseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#create_database_link.
	VisitCreate_database_link(ctx *Create_database_linkContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dblink.
	VisitDblink(ctx *DblinkContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_database_link.
	VisitDrop_database_link(ctx *Drop_database_linkContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_tablespace_set.
	VisitAlter_tablespace_set(ctx *Alter_tablespace_setContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_tablespace_attrs.
	VisitAlter_tablespace_attrs(ctx *Alter_tablespace_attrsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_tablespace_encryption.
	VisitAlter_tablespace_encryption(ctx *Alter_tablespace_encryptionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ts_file_name_convert.
	VisitTs_file_name_convert(ctx *Ts_file_name_convertContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_role.
	VisitAlter_role(ctx *Alter_roleContext) interface{}

	// Visit a parse tree produced by PlSqlParser#role_identified_clause.
	VisitRole_identified_clause(ctx *Role_identified_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_table.
	VisitAlter_table(ctx *Alter_tableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#memoptimize_read_write_clause.
	VisitMemoptimize_read_write_clause(ctx *Memoptimize_read_write_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_table_properties.
	VisitAlter_table_properties(ctx *Alter_table_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_table_partitioning.
	VisitAlter_table_partitioning(ctx *Alter_table_partitioningContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_table_partition.
	VisitAdd_table_partition(ctx *Add_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_table_partition.
	VisitDrop_table_partition(ctx *Drop_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_table_partition.
	VisitMerge_table_partition(ctx *Merge_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_table_partition.
	VisitModify_table_partition(ctx *Modify_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#split_table_partition.
	VisitSplit_table_partition(ctx *Split_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#truncate_table_partition.
	VisitTruncate_table_partition(ctx *Truncate_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exchange_table_partition.
	VisitExchange_table_partition(ctx *Exchange_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#coalesce_table_partition.
	VisitCoalesce_table_partition(ctx *Coalesce_table_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_interval_partition.
	VisitAlter_interval_partition(ctx *Alter_interval_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_extended_names.
	VisitPartition_extended_names(ctx *Partition_extended_namesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subpartition_extended_names.
	VisitSubpartition_extended_names(ctx *Subpartition_extended_namesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_table_properties_1.
	VisitAlter_table_properties_1(ctx *Alter_table_properties_1Context) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_iot_clauses.
	VisitAlter_iot_clauses(ctx *Alter_iot_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_mapping_table_clause.
	VisitAlter_mapping_table_clause(ctx *Alter_mapping_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_overflow_clause.
	VisitAlter_overflow_clause(ctx *Alter_overflow_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_overflow_clause.
	VisitAdd_overflow_clause(ctx *Add_overflow_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_index_clauses.
	VisitUpdate_index_clauses(ctx *Update_index_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_global_index_clause.
	VisitUpdate_global_index_clause(ctx *Update_global_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_all_indexes_clause.
	VisitUpdate_all_indexes_clause(ctx *Update_all_indexes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_all_indexes_index_clause.
	VisitUpdate_all_indexes_index_clause(ctx *Update_all_indexes_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_index_partition.
	VisitUpdate_index_partition(ctx *Update_index_partitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_index_subpartition.
	VisitUpdate_index_subpartition(ctx *Update_index_subpartitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#enable_disable_clause.
	VisitEnable_disable_clause(ctx *Enable_disable_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_index_clause.
	VisitUsing_index_clause(ctx *Using_index_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_attributes.
	VisitIndex_attributes(ctx *Index_attributesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sort_or_nosort.
	VisitSort_or_nosort(ctx *Sort_or_nosortContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exceptions_clause.
	VisitExceptions_clause(ctx *Exceptions_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#move_table_clause.
	VisitMove_table_clause(ctx *Move_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_org_table_clause.
	VisitIndex_org_table_clause(ctx *Index_org_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#mapping_table_clause.
	VisitMapping_table_clause(ctx *Mapping_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#key_compression.
	VisitKey_compression(ctx *Key_compressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_org_overflow_clause.
	VisitIndex_org_overflow_clause(ctx *Index_org_overflow_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_clauses.
	VisitColumn_clauses(ctx *Column_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_collection_retrieval.
	VisitModify_collection_retrieval(ctx *Modify_collection_retrievalContext) interface{}

	// Visit a parse tree produced by PlSqlParser#collection_item.
	VisitCollection_item(ctx *Collection_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rename_column_clause.
	VisitRename_column_clause(ctx *Rename_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#old_column_name.
	VisitOld_column_name(ctx *Old_column_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_column_name.
	VisitNew_column_name(ctx *New_column_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_modify_drop_column_clauses.
	VisitAdd_modify_drop_column_clauses(ctx *Add_modify_drop_column_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_column_clause.
	VisitDrop_column_clause(ctx *Drop_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_column_clauses.
	VisitModify_column_clauses(ctx *Modify_column_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_col_visibility.
	VisitModify_col_visibility(ctx *Modify_col_visibilityContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_col_properties.
	VisitModify_col_properties(ctx *Modify_col_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_col_substitutable.
	VisitModify_col_substitutable(ctx *Modify_col_substitutableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_column_clause.
	VisitAdd_column_clause(ctx *Add_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#alter_varray_col_properties.
	VisitAlter_varray_col_properties(ctx *Alter_varray_col_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#varray_col_properties.
	VisitVarray_col_properties(ctx *Varray_col_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#varray_storage_clause.
	VisitVarray_storage_clause(ctx *Varray_storage_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_segname.
	VisitLob_segname(ctx *Lob_segnameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_item.
	VisitLob_item(ctx *Lob_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_storage_parameters.
	VisitLob_storage_parameters(ctx *Lob_storage_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_storage_clause.
	VisitLob_storage_clause(ctx *Lob_storage_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_lob_storage_clause.
	VisitModify_lob_storage_clause(ctx *Modify_lob_storage_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#modify_lob_parameters.
	VisitModify_lob_parameters(ctx *Modify_lob_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_parameters.
	VisitLob_parameters(ctx *Lob_parametersContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_deduplicate_clause.
	VisitLob_deduplicate_clause(ctx *Lob_deduplicate_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_compression_clause.
	VisitLob_compression_clause(ctx *Lob_compression_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_retention_clause.
	VisitLob_retention_clause(ctx *Lob_retention_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#encryption_spec.
	VisitEncryption_spec(ctx *Encryption_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tablespace.
	VisitTablespace(ctx *TablespaceContext) interface{}

	// Visit a parse tree produced by PlSqlParser#varray_item.
	VisitVarray_item(ctx *Varray_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_properties.
	VisitColumn_properties(ctx *Column_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lob_partition_storage.
	VisitLob_partition_storage(ctx *Lob_partition_storageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#period_definition.
	VisitPeriod_definition(ctx *Period_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#start_time_column.
	VisitStart_time_column(ctx *Start_time_columnContext) interface{}

	// Visit a parse tree produced by PlSqlParser#end_time_column.
	VisitEnd_time_column(ctx *End_time_columnContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_definition.
	VisitColumn_definition(ctx *Column_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_collation_name.
	VisitColumn_collation_name(ctx *Column_collation_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identity_clause.
	VisitIdentity_clause(ctx *Identity_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identity_options_parentheses.
	VisitIdentity_options_parentheses(ctx *Identity_options_parenthesesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identity_options.
	VisitIdentity_options(ctx *Identity_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#virtual_column_definition.
	VisitVirtual_column_definition(ctx *Virtual_column_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#autogenerated_sequence_definition.
	VisitAutogenerated_sequence_definition(ctx *Autogenerated_sequence_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#evaluation_edition_clause.
	VisitEvaluation_edition_clause(ctx *Evaluation_edition_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#out_of_line_part_storage.
	VisitOut_of_line_part_storage(ctx *Out_of_line_part_storageContext) interface{}

	// Visit a parse tree produced by PlSqlParser#nested_table_col_properties.
	VisitNested_table_col_properties(ctx *Nested_table_col_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#nested_item.
	VisitNested_item(ctx *Nested_itemContext) interface{}

	// Visit a parse tree produced by PlSqlParser#substitutable_column_clause.
	VisitSubstitutable_column_clause(ctx *Substitutable_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_name.
	VisitPartition_name(ctx *Partition_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#supplemental_logging_props.
	VisitSupplemental_logging_props(ctx *Supplemental_logging_propsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_or_attribute.
	VisitColumn_or_attribute(ctx *Column_or_attributeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_type_col_properties.
	VisitObject_type_col_properties(ctx *Object_type_col_propertiesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constraint_clauses.
	VisitConstraint_clauses(ctx *Constraint_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#old_constraint_name.
	VisitOld_constraint_name(ctx *Old_constraint_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#new_constraint_name.
	VisitNew_constraint_name(ctx *New_constraint_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_constraint_clause.
	VisitDrop_constraint_clause(ctx *Drop_constraint_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_primary_key_or_unique_or_generic_clause.
	VisitDrop_primary_key_or_unique_or_generic_clause(ctx *Drop_primary_key_or_unique_or_generic_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_constraint.
	VisitAdd_constraint(ctx *Add_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#add_constraint_clause.
	VisitAdd_constraint_clause(ctx *Add_constraint_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#check_constraint.
	VisitCheck_constraint(ctx *Check_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#drop_constraint.
	VisitDrop_constraint(ctx *Drop_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#enable_constraint.
	VisitEnable_constraint(ctx *Enable_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#disable_constraint.
	VisitDisable_constraint(ctx *Disable_constraintContext) interface{}

	// Visit a parse tree produced by PlSqlParser#foreign_key_clause.
	VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#references_clause.
	VisitReferences_clause(ctx *References_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#on_delete_clause.
	VisitOn_delete_clause(ctx *On_delete_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unique_key_clause.
	VisitUnique_key_clause(ctx *Unique_key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#primary_key_clause.
	VisitPrimary_key_clause(ctx *Primary_key_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#anonymous_block.
	VisitAnonymous_block(ctx *Anonymous_blockContext) interface{}

	// Visit a parse tree produced by PlSqlParser#invoker_rights_clause.
	VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#call_spec.
	VisitCall_spec(ctx *Call_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#java_spec.
	VisitJava_spec(ctx *Java_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#c_spec.
	VisitC_spec(ctx *C_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#c_agent_in_clause.
	VisitC_agent_in_clause(ctx *C_agent_in_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#c_parameters_clause.
	VisitC_parameters_clause(ctx *C_parameters_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by PlSqlParser#default_value_part.
	VisitDefault_value_part(ctx *Default_value_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#seq_of_declare_specs.
	VisitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#declare_spec.
	VisitDeclare_spec(ctx *Declare_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#variable_declaration.
	VisitVariable_declaration(ctx *Variable_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subtype_declaration.
	VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cursor_declaration.
	VisitCursor_declaration(ctx *Cursor_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parameter_spec.
	VisitParameter_spec(ctx *Parameter_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exception_declaration.
	VisitException_declaration(ctx *Exception_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pragma_declaration.
	VisitPragma_declaration(ctx *Pragma_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#record_type_def.
	VisitRecord_type_def(ctx *Record_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#field_spec.
	VisitField_spec(ctx *Field_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#ref_cursor_type_def.
	VisitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_type_def.
	VisitTable_type_def(ctx *Table_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_indexed_by_part.
	VisitTable_indexed_by_part(ctx *Table_indexed_by_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#varray_type_def.
	VisitVarray_type_def(ctx *Varray_type_defContext) interface{}

	// Visit a parse tree produced by PlSqlParser#seq_of_statements.
	VisitSeq_of_statements(ctx *Seq_of_statementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#label_declaration.
	VisitLabel_declaration(ctx *Label_declarationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#swallow_to_semi.
	VisitSwallow_to_semi(ctx *Swallow_to_semiContext) interface{}

	// Visit a parse tree produced by PlSqlParser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exit_statement.
	VisitExit_statement(ctx *Exit_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#goto_statement.
	VisitGoto_statement(ctx *Goto_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#elsif_part.
	VisitElsif_part(ctx *Elsif_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#else_part.
	VisitElse_part(ctx *Else_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cursor_loop_param.
	VisitCursor_loop_param(ctx *Cursor_loop_paramContext) interface{}

	// Visit a parse tree produced by PlSqlParser#forall_statement.
	VisitForall_statement(ctx *Forall_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#bounds_clause.
	VisitBounds_clause(ctx *Bounds_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#between_bound.
	VisitBetween_bound(ctx *Between_boundContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lower_bound.
	VisitLower_bound(ctx *Lower_boundContext) interface{}

	// Visit a parse tree produced by PlSqlParser#upper_bound.
	VisitUpper_bound(ctx *Upper_boundContext) interface{}

	// Visit a parse tree produced by PlSqlParser#null_statement.
	VisitNull_statement(ctx *Null_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#raise_statement.
	VisitRaise_statement(ctx *Raise_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#call_statement.
	VisitCall_statement(ctx *Call_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pipe_row_statement.
	VisitPipe_row_statement(ctx *Pipe_row_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exception_handler.
	VisitException_handler(ctx *Exception_handlerContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trigger_block.
	VisitTrigger_block(ctx *Trigger_blockContext) interface{}

	// Visit a parse tree produced by PlSqlParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sql_statement.
	VisitSql_statement(ctx *Sql_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#execute_immediate.
	VisitExecute_immediate(ctx *Execute_immediateContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dynamic_returning_clause.
	VisitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#data_manipulation_language_statements.
	VisitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cursor_manipulation_statements.
	VisitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#close_statement.
	VisitClose_statement(ctx *Close_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#open_statement.
	VisitOpen_statement(ctx *Open_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#fetch_statement.
	VisitFetch_statement(ctx *Fetch_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#open_for_statement.
	VisitOpen_for_statement(ctx *Open_for_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#transaction_control_statements.
	VisitTransaction_control_statements(ctx *Transaction_control_statementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_transaction_command.
	VisitSet_transaction_command(ctx *Set_transaction_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_constraint_command.
	VisitSet_constraint_command(ctx *Set_constraint_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#commit_statement.
	VisitCommit_statement(ctx *Commit_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#write_clause.
	VisitWrite_clause(ctx *Write_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rollback_statement.
	VisitRollback_statement(ctx *Rollback_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#savepoint_statement.
	VisitSavepoint_statement(ctx *Savepoint_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#explain_statement.
	VisitExplain_statement(ctx *Explain_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#select_only_statement.
	VisitSelect_only_statement(ctx *Select_only_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subquery_factoring_clause.
	VisitSubquery_factoring_clause(ctx *Subquery_factoring_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#factoring_element.
	VisitFactoring_element(ctx *Factoring_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#search_clause.
	VisitSearch_clause(ctx *Search_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cycle_clause.
	VisitCycle_clause(ctx *Cycle_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subquery_basic_elements.
	VisitSubquery_basic_elements(ctx *Subquery_basic_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subquery_operation_part.
	VisitSubquery_operation_part(ctx *Subquery_operation_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#query_block.
	VisitQuery_block(ctx *Query_blockContext) interface{}

	// Visit a parse tree produced by PlSqlParser#selected_list.
	VisitSelected_list(ctx *Selected_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#select_list_elements.
	VisitSelect_list_elements(ctx *Select_list_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_wild.
	VisitTable_wild(ctx *Table_wildContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref_list.
	VisitTable_ref_list(ctx *Table_ref_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref.
	VisitTable_ref(ctx *Table_refContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref_aux.
	VisitTable_ref_aux(ctx *Table_ref_auxContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref_aux_internal_one.
	VisitTable_ref_aux_internal_one(ctx *Table_ref_aux_internal_oneContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref_aux_internal_two.
	VisitTable_ref_aux_internal_two(ctx *Table_ref_aux_internal_twoContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_ref_aux_internal_three.
	VisitTable_ref_aux_internal_three(ctx *Table_ref_aux_internal_threeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_clause.
	VisitJoin_clause(ctx *Join_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_on_part.
	VisitJoin_on_part(ctx *Join_on_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#join_using_part.
	VisitJoin_using_part(ctx *Join_using_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#outer_join_type.
	VisitOuter_join_type(ctx *Outer_join_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#query_partition_clause.
	VisitQuery_partition_clause(ctx *Query_partition_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#flashback_query_clause.
	VisitFlashback_query_clause(ctx *Flashback_query_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_clause.
	VisitPivot_clause(ctx *Pivot_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_element.
	VisitPivot_element(ctx *Pivot_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_for_clause.
	VisitPivot_for_clause(ctx *Pivot_for_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_in_clause.
	VisitPivot_in_clause(ctx *Pivot_in_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_in_clause_element.
	VisitPivot_in_clause_element(ctx *Pivot_in_clause_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#pivot_in_clause_elements.
	VisitPivot_in_clause_elements(ctx *Pivot_in_clause_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unpivot_clause.
	VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unpivot_in_clause.
	VisitUnpivot_in_clause(ctx *Unpivot_in_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unpivot_in_elements.
	VisitUnpivot_in_elements(ctx *Unpivot_in_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#hierarchical_query_clause.
	VisitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#start_part.
	VisitStart_part(ctx *Start_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#group_by_clause.
	VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#group_by_elements.
	VisitGroup_by_elements(ctx *Group_by_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rollup_cube_clause.
	VisitRollup_cube_clause(ctx *Rollup_cube_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#grouping_sets_clause.
	VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#grouping_sets_elements.
	VisitGrouping_sets_elements(ctx *Grouping_sets_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_clause.
	VisitModel_clause(ctx *Model_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cell_reference_options.
	VisitCell_reference_options(ctx *Cell_reference_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#return_rows_clause.
	VisitReturn_rows_clause(ctx *Return_rows_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#reference_model.
	VisitReference_model(ctx *Reference_modelContext) interface{}

	// Visit a parse tree produced by PlSqlParser#main_model.
	VisitMain_model(ctx *Main_modelContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_column_clauses.
	VisitModel_column_clauses(ctx *Model_column_clausesContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_column_partition_part.
	VisitModel_column_partition_part(ctx *Model_column_partition_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_column_list.
	VisitModel_column_list(ctx *Model_column_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_column.
	VisitModel_column(ctx *Model_columnContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_rules_clause.
	VisitModel_rules_clause(ctx *Model_rules_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_rules_part.
	VisitModel_rules_part(ctx *Model_rules_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_rules_element.
	VisitModel_rules_element(ctx *Model_rules_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cell_assignment.
	VisitCell_assignment(ctx *Cell_assignmentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_iterate_clause.
	VisitModel_iterate_clause(ctx *Model_iterate_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#until_part.
	VisitUntil_part(ctx *Until_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#order_by_clause.
	VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#order_by_elements.
	VisitOrder_by_elements(ctx *Order_by_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#offset_clause.
	VisitOffset_clause(ctx *Offset_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#fetch_clause.
	VisitFetch_clause(ctx *Fetch_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#for_update_clause.
	VisitFor_update_clause(ctx *For_update_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#for_update_of_part.
	VisitFor_update_of_part(ctx *For_update_of_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#for_update_options.
	VisitFor_update_options(ctx *For_update_optionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_statement.
	VisitUpdate_statement(ctx *Update_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#update_set_clause.
	VisitUpdate_set_clause(ctx *Update_set_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_based_update_set_clause.
	VisitColumn_based_update_set_clause(ctx *Column_based_update_set_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#delete_statement.
	VisitDelete_statement(ctx *Delete_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#insert_statement.
	VisitInsert_statement(ctx *Insert_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#single_table_insert.
	VisitSingle_table_insert(ctx *Single_table_insertContext) interface{}

	// Visit a parse tree produced by PlSqlParser#multi_table_insert.
	VisitMulti_table_insert(ctx *Multi_table_insertContext) interface{}

	// Visit a parse tree produced by PlSqlParser#multi_table_element.
	VisitMulti_table_element(ctx *Multi_table_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#conditional_insert_clause.
	VisitConditional_insert_clause(ctx *Conditional_insert_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#conditional_insert_when_part.
	VisitConditional_insert_when_part(ctx *Conditional_insert_when_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#conditional_insert_else_part.
	VisitConditional_insert_else_part(ctx *Conditional_insert_else_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#insert_into_clause.
	VisitInsert_into_clause(ctx *Insert_into_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#values_clause.
	VisitValues_clause(ctx *Values_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_statement.
	VisitMerge_statement(ctx *Merge_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_update_clause.
	VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_element.
	VisitMerge_element(ctx *Merge_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_update_delete_part.
	VisitMerge_update_delete_part(ctx *Merge_update_delete_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#merge_insert_clause.
	VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#selected_tableview.
	VisitSelected_tableview(ctx *Selected_tableviewContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lock_table_statement.
	VisitLock_table_statement(ctx *Lock_table_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#wait_nowait_part.
	VisitWait_nowait_part(ctx *Wait_nowait_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lock_table_element.
	VisitLock_table_element(ctx *Lock_table_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#lock_mode.
	VisitLock_mode(ctx *Lock_modeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#general_table_ref.
	VisitGeneral_table_ref(ctx *General_table_refContext) interface{}

	// Visit a parse tree produced by PlSqlParser#static_returning_clause.
	VisitStatic_returning_clause(ctx *Static_returning_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#error_logging_clause.
	VisitError_logging_clause(ctx *Error_logging_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#error_logging_into_part.
	VisitError_logging_into_part(ctx *Error_logging_into_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#error_logging_reject_part.
	VisitError_logging_reject_part(ctx *Error_logging_reject_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dml_table_expression_clause.
	VisitDml_table_expression_clause(ctx *Dml_table_expression_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_collection_expression.
	VisitTable_collection_expression(ctx *Table_collection_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#subquery_restriction_clause.
	VisitSubquery_restriction_clause(ctx *Subquery_restriction_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sample_clause.
	VisitSample_clause(ctx *Sample_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#seed_part.
	VisitSeed_part(ctx *Seed_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_condition.
	VisitJson_condition(ctx *Json_conditionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cursor_expression.
	VisitCursor_expression(ctx *Cursor_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logical_expression.
	VisitLogical_expression(ctx *Logical_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unary_logical_expression.
	VisitUnary_logical_expression(ctx *Unary_logical_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#logical_operation.
	VisitLogical_operation(ctx *Logical_operationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#multiset_expression.
	VisitMultiset_expression(ctx *Multiset_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#relational_expression.
	VisitRelational_expression(ctx *Relational_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#compound_expression.
	VisitCompound_expression(ctx *Compound_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#relational_operator.
	VisitRelational_operator(ctx *Relational_operatorContext) interface{}

	// Visit a parse tree produced by PlSqlParser#in_elements.
	VisitIn_elements(ctx *In_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#between_elements.
	VisitBetween_elements(ctx *Between_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#concatenation.
	VisitConcatenation(ctx *ConcatenationContext) interface{}

	// Visit a parse tree produced by PlSqlParser#interval_expression.
	VisitInterval_expression(ctx *Interval_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_expression.
	VisitModel_expression(ctx *Model_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#model_expression_element.
	VisitModel_expression_element(ctx *Model_expression_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#single_column_for_loop.
	VisitSingle_column_for_loop(ctx *Single_column_for_loopContext) interface{}

	// Visit a parse tree produced by PlSqlParser#multi_column_for_loop.
	VisitMulti_column_for_loop(ctx *Multi_column_for_loopContext) interface{}

	// Visit a parse tree produced by PlSqlParser#unary_expression.
	VisitUnary_expression(ctx *Unary_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#simple_case_statement.
	VisitSimple_case_statement(ctx *Simple_case_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#simple_case_when_part.
	VisitSimple_case_when_part(ctx *Simple_case_when_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#searched_case_statement.
	VisitSearched_case_statement(ctx *Searched_case_statementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#searched_case_when_part.
	VisitSearched_case_when_part(ctx *Searched_case_when_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#case_else_part.
	VisitCase_else_part(ctx *Case_else_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quantified_expression.
	VisitQuantified_expression(ctx *Quantified_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#string_function.
	VisitString_function(ctx *String_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#standard_function.
	VisitStandard_function(ctx *Standard_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_function.
	VisitJson_function(ctx *Json_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_object_content.
	VisitJson_object_content(ctx *Json_object_contentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_object_entry.
	VisitJson_object_entry(ctx *Json_object_entryContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_table_clause.
	VisitJson_table_clause(ctx *Json_table_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_array_element.
	VisitJson_array_element(ctx *Json_array_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_on_null_clause.
	VisitJson_on_null_clause(ctx *Json_on_null_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_return_clause.
	VisitJson_return_clause(ctx *Json_return_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_transform_op.
	VisitJson_transform_op(ctx *Json_transform_opContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_column_clause.
	VisitJson_column_clause(ctx *Json_column_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_column_definition.
	VisitJson_column_definition(ctx *Json_column_definitionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_query_returning_clause.
	VisitJson_query_returning_clause(ctx *Json_query_returning_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_query_return_type.
	VisitJson_query_return_type(ctx *Json_query_return_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_query_wrapper_clause.
	VisitJson_query_wrapper_clause(ctx *Json_query_wrapper_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_query_on_error_clause.
	VisitJson_query_on_error_clause(ctx *Json_query_on_error_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_query_on_empty_clause.
	VisitJson_query_on_empty_clause(ctx *Json_query_on_empty_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_value_return_clause.
	VisitJson_value_return_clause(ctx *Json_value_return_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_value_return_type.
	VisitJson_value_return_type(ctx *Json_value_return_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#json_value_on_mismatch_clause.
	VisitJson_value_on_mismatch_clause(ctx *Json_value_on_mismatch_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PlSqlParser#numeric_function_wrapper.
	VisitNumeric_function_wrapper(ctx *Numeric_function_wrapperContext) interface{}

	// Visit a parse tree produced by PlSqlParser#numeric_function.
	VisitNumeric_function(ctx *Numeric_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#listagg_overflow_clause.
	VisitListagg_overflow_clause(ctx *Listagg_overflow_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#other_function.
	VisitOther_function(ctx *Other_functionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#over_clause_keyword.
	VisitOver_clause_keyword(ctx *Over_clause_keywordContext) interface{}

	// Visit a parse tree produced by PlSqlParser#within_or_over_clause_keyword.
	VisitWithin_or_over_clause_keyword(ctx *Within_or_over_clause_keywordContext) interface{}

	// Visit a parse tree produced by PlSqlParser#standard_prediction_function_keyword.
	VisitStandard_prediction_function_keyword(ctx *Standard_prediction_function_keywordContext) interface{}

	// Visit a parse tree produced by PlSqlParser#over_clause.
	VisitOver_clause(ctx *Over_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#windowing_clause.
	VisitWindowing_clause(ctx *Windowing_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#windowing_type.
	VisitWindowing_type(ctx *Windowing_typeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#windowing_elements.
	VisitWindowing_elements(ctx *Windowing_elementsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#using_element.
	VisitUsing_element(ctx *Using_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#collect_order_by_part.
	VisitCollect_order_by_part(ctx *Collect_order_by_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#within_or_over_part.
	VisitWithin_or_over_part(ctx *Within_or_over_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cost_matrix_clause.
	VisitCost_matrix_clause(ctx *Cost_matrix_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_passing_clause.
	VisitXml_passing_clause(ctx *Xml_passing_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_attributes_clause.
	VisitXml_attributes_clause(ctx *Xml_attributes_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_namespaces_clause.
	VisitXml_namespaces_clause(ctx *Xml_namespaces_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_table_column.
	VisitXml_table_column(ctx *Xml_table_columnContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_general_default_part.
	VisitXml_general_default_part(ctx *Xml_general_default_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_multiuse_expression_element.
	VisitXml_multiuse_expression_element(ctx *Xml_multiuse_expression_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlroot_param_version_part.
	VisitXmlroot_param_version_part(ctx *Xmlroot_param_version_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlroot_param_standalone_part.
	VisitXmlroot_param_standalone_part(ctx *Xmlroot_param_standalone_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlserialize_param_enconding_part.
	VisitXmlserialize_param_enconding_part(ctx *Xmlserialize_param_enconding_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlserialize_param_version_part.
	VisitXmlserialize_param_version_part(ctx *Xmlserialize_param_version_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmlserialize_param_ident_part.
	VisitXmlserialize_param_ident_part(ctx *Xmlserialize_param_ident_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sql_plus_command.
	VisitSql_plus_command(ctx *Sql_plus_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#whenever_command.
	VisitWhenever_command(ctx *Whenever_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#set_command.
	VisitSet_command(ctx *Set_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#timing_command.
	VisitTiming_command(ctx *Timing_commandContext) interface{}

	// Visit a parse tree produced by PlSqlParser#partition_extension_clause.
	VisitPartition_extension_clause(ctx *Partition_extension_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by PlSqlParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quantitative_where_stmt.
	VisitQuantitative_where_stmt(ctx *Quantitative_where_stmtContext) interface{}

	// Visit a parse tree produced by PlSqlParser#into_clause.
	VisitInto_clause(ctx *Into_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xml_column_name.
	VisitXml_column_name(ctx *Xml_column_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cost_class_name.
	VisitCost_class_name(ctx *Cost_class_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#attribute_name.
	VisitAttribute_name(ctx *Attribute_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#savepoint_name.
	VisitSavepoint_name(ctx *Savepoint_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#rollback_segment_name.
	VisitRollback_segment_name(ctx *Rollback_segment_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_var_name.
	VisitTable_var_name(ctx *Table_var_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#schema_name.
	VisitSchema_name(ctx *Schema_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#routine_name.
	VisitRoutine_name(ctx *Routine_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#package_name.
	VisitPackage_name(ctx *Package_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#implementation_type_name.
	VisitImplementation_type_name(ctx *Implementation_type_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#parameter_name.
	VisitParameter_name(ctx *Parameter_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#reference_model_name.
	VisitReference_model_name(ctx *Reference_model_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#main_model_name.
	VisitMain_model_name(ctx *Main_model_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#container_tableview_name.
	VisitContainer_tableview_name(ctx *Container_tableview_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#aggregate_function_name.
	VisitAggregate_function_name(ctx *Aggregate_function_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#query_name.
	VisitQuery_name(ctx *Query_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#grantee_name.
	VisitGrantee_name(ctx *Grantee_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#role_name.
	VisitRole_name(ctx *Role_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constraint_name.
	VisitConstraint_name(ctx *Constraint_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#label_name.
	VisitLabel_name(ctx *Label_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#sequence_name.
	VisitSequence_name(ctx *Sequence_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#exception_name.
	VisitException_name(ctx *Exception_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#procedure_name.
	VisitProcedure_name(ctx *Procedure_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#trigger_name.
	VisitTrigger_name(ctx *Trigger_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#variable_name.
	VisitVariable_name(ctx *Variable_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#cursor_name.
	VisitCursor_name(ctx *Cursor_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#record_name.
	VisitRecord_name(ctx *Record_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#collection_name.
	VisitCollection_name(ctx *Collection_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#link_name.
	VisitLink_name(ctx *Link_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#tableview_name.
	VisitTableview_name(ctx *Tableview_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#xmltable.
	VisitXmltable(ctx *XmltableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#char_set_name.
	VisitChar_set_name(ctx *Char_set_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#synonym_name.
	VisitSynonym_name(ctx *Synonym_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#schema_object_name.
	VisitSchema_object_name(ctx *Schema_object_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#dir_object_name.
	VisitDir_object_name(ctx *Dir_object_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#user_object_name.
	VisitUser_object_name(ctx *User_object_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#grant_object_name.
	VisitGrant_object_name(ctx *Grant_object_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#column_list.
	VisitColumn_list(ctx *Column_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#paren_column_list.
	VisitParen_column_list(ctx *Paren_column_listContext) interface{}

	// Visit a parse tree produced by PlSqlParser#keep_clause.
	VisitKeep_clause(ctx *Keep_clauseContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_argument.
	VisitFunction_argument(ctx *Function_argumentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_argument_analytic.
	VisitFunction_argument_analytic(ctx *Function_argument_analyticContext) interface{}

	// Visit a parse tree produced by PlSqlParser#function_argument_modeling.
	VisitFunction_argument_modeling(ctx *Function_argument_modelingContext) interface{}

	// Visit a parse tree produced by PlSqlParser#respect_or_ignore_nulls.
	VisitRespect_or_ignore_nulls(ctx *Respect_or_ignore_nullsContext) interface{}

	// Visit a parse tree produced by PlSqlParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by PlSqlParser#type_spec.
	VisitType_spec(ctx *Type_specContext) interface{}

	// Visit a parse tree produced by PlSqlParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#precision_part.
	VisitPrecision_part(ctx *Precision_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#native_datatype_element.
	VisitNative_datatype_element(ctx *Native_datatype_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#bind_variable.
	VisitBind_variable(ctx *Bind_variableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#general_element_part.
	VisitGeneral_element_part(ctx *General_element_partContext) interface{}

	// Visit a parse tree produced by PlSqlParser#table_element.
	VisitTable_element(ctx *Table_elementContext) interface{}

	// Visit a parse tree produced by PlSqlParser#object_privilege.
	VisitObject_privilege(ctx *Object_privilegeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#system_privilege.
	VisitSystem_privilege(ctx *System_privilegeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by PlSqlParser#constant_without_variable.
	VisitConstant_without_variable(ctx *Constant_without_variableContext) interface{}

	// Visit a parse tree produced by PlSqlParser#numeric.
	VisitNumeric(ctx *NumericContext) interface{}

	// Visit a parse tree produced by PlSqlParser#numeric_negative.
	VisitNumeric_negative(ctx *Numeric_negativeContext) interface{}

	// Visit a parse tree produced by PlSqlParser#quoted_string.
	VisitQuoted_string(ctx *Quoted_stringContext) interface{}

	// Visit a parse tree produced by PlSqlParser#char_str.
	VisitChar_str(ctx *Char_strContext) interface{}

	// Visit a parse tree produced by PlSqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by PlSqlParser#id_expression.
	VisitId_expression(ctx *Id_expressionContext) interface{}

	// Visit a parse tree produced by PlSqlParser#outer_join_sign.
	VisitOuter_join_sign(ctx *Outer_join_signContext) interface{}

	// Visit a parse tree produced by PlSqlParser#regular_id.
	VisitRegular_id(ctx *Regular_idContext) interface{}

	// Visit a parse tree produced by PlSqlParser#non_reserved_keywords_in_12c.
	VisitNon_reserved_keywords_in_12c(ctx *Non_reserved_keywords_in_12cContext) interface{}

	// Visit a parse tree produced by PlSqlParser#non_reserved_keywords_pre12c.
	VisitNon_reserved_keywords_pre12c(ctx *Non_reserved_keywords_pre12cContext) interface{}

	// Visit a parse tree produced by PlSqlParser#string_function_name.
	VisitString_function_name(ctx *String_function_nameContext) interface{}

	// Visit a parse tree produced by PlSqlParser#numeric_function_name.
	VisitNumeric_function_name(ctx *Numeric_function_nameContext) interface{}
}
