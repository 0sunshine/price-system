use db_price_system;

CREATE TABLE table_project_kind (
	project_kind_id varchar(64) not null COMMENT '项目类型id',
	project_name varchar(256) character set 'utf8mb4' not null COMMENT '项目类型名称',
	parent_project_kind_id varchar(64) not null default '' COMMENT '父项目类型id',
	notes varchar(1024) character set 'utf8mb4' not null COMMENT '备注',
	
	PRIMARY KEY (project_kind_id) 
) default character set = 'utf8mb4';


CREATE TABLE table_material_kind (
	material_kind_id bigint not null auto_increment COMMENT '材料类型id',
	project_kind_id varchar(64) not null COMMENT '项目类型id',
	material_name varchar(256) character set 'utf8mb4' not null COMMENT '材料类型名称',
	unit varchar(32) character set 'utf8mb4' not null COMMENT '单位',
	notes varchar(1024) character set 'utf8mb4' not null COMMENT '备注',
	
	PRIMARY KEY (material_kind_id) 
) default character set = 'utf8mb4';


CREATE TABLE table_material_attr (
	attr_id bigint not null auto_increment COMMENT '材料属性id',
	material_kind_id bigint not null COMMENT '材料类型id',
	attr_desc text character set 'utf8mb4' not null COMMENT '材料属性描述',
	PRIMARY KEY (attr_id)
) default character set = 'utf8mb4';
	
	
CREATE TABLE table_material (
	material_id bigint not null auto_increment COMMENT '材料id',
	material_kind_id bigint not null COMMENT '材料类型id',
	price decimal(62,2) not null COMMENT '材料价格',
	extra_price decimal(62,2) not null COMMENT '额外材料价格，如人工',
	
	PRIMARY KEY (material_id) 
) default character set = 'utf8mb4';

	
CREATE TABLE table_material_formation (
	material_id bigint not null COMMENT '材料id',
	attr_id bigint not null COMMENT '材料属性id',
	
	PRIMARY KEY (material_id,attr_id) 
) default character set = 'utf8mb4';

