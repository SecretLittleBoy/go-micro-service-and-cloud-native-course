// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"review-service/internal/data/model"
)

func newReviewAppealInfo(db *gorm.DB, opts ...gen.DOOption) reviewAppealInfo {
	_reviewAppealInfo := reviewAppealInfo{}

	_reviewAppealInfo.reviewAppealInfoDo.UseDB(db, opts...)
	_reviewAppealInfo.reviewAppealInfoDo.UseModel(&model.ReviewAppealInfo{})

	tableName := _reviewAppealInfo.reviewAppealInfoDo.TableName()
	_reviewAppealInfo.ALL = field.NewAsterisk(tableName)
	_reviewAppealInfo.ID = field.NewInt64(tableName, "id")
	_reviewAppealInfo.CreateBy = field.NewString(tableName, "create_by")
	_reviewAppealInfo.UpdateBy = field.NewString(tableName, "update_by")
	_reviewAppealInfo.CreateAt = field.NewTime(tableName, "create_at")
	_reviewAppealInfo.UpdateAt = field.NewTime(tableName, "update_at")
	_reviewAppealInfo.DeleteAt = field.NewTime(tableName, "delete_at")
	_reviewAppealInfo.Version = field.NewInt32(tableName, "version")
	_reviewAppealInfo.AppealID = field.NewInt64(tableName, "appeal_id")
	_reviewAppealInfo.ReviewID = field.NewInt64(tableName, "review_id")
	_reviewAppealInfo.StoreID = field.NewInt64(tableName, "store_id")
	_reviewAppealInfo.Status = field.NewInt32(tableName, "status")
	_reviewAppealInfo.Reason = field.NewString(tableName, "reason")
	_reviewAppealInfo.Content = field.NewString(tableName, "content")
	_reviewAppealInfo.PicInfo = field.NewString(tableName, "pic_info")
	_reviewAppealInfo.VideoInfo = field.NewString(tableName, "video_info")
	_reviewAppealInfo.OpRemarks = field.NewString(tableName, "op_remarks")
	_reviewAppealInfo.OpUser = field.NewString(tableName, "op_user")
	_reviewAppealInfo.ExtJSON = field.NewString(tableName, "ext_json")
	_reviewAppealInfo.CtrlJSON = field.NewString(tableName, "ctrl_json")

	_reviewAppealInfo.fillFieldMap()

	return _reviewAppealInfo
}

type reviewAppealInfo struct {
	reviewAppealInfoDo reviewAppealInfoDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键
	CreateBy  field.String // 创建方标识
	UpdateBy  field.String // 更新方标识
	CreateAt  field.Time   // 创建时间
	UpdateAt  field.Time   // 更新时间
	DeleteAt  field.Time   // 逻辑删除标记
	Version   field.Int32  // 乐观锁标记
	AppealID  field.Int64  // 回复id
	ReviewID  field.Int64  // 评价id
	StoreID   field.Int64  // 店铺id
	Status    field.Int32  // 状态:10待审核；20申诉通过；30申诉驳回
	Reason    field.String // 申诉原因类别
	Content   field.String // 申诉内容描述
	PicInfo   field.String // 媒体信息：图片
	VideoInfo field.String // 媒体信息：视频
	OpRemarks field.String // 运营备注
	OpUser    field.String // 运营者标识
	ExtJSON   field.String // 信息扩展
	CtrlJSON  field.String // 控制扩展

	fieldMap map[string]field.Expr
}

func (r reviewAppealInfo) Table(newTableName string) *reviewAppealInfo {
	r.reviewAppealInfoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reviewAppealInfo) As(alias string) *reviewAppealInfo {
	r.reviewAppealInfoDo.DO = *(r.reviewAppealInfoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reviewAppealInfo) updateTableName(table string) *reviewAppealInfo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreateBy = field.NewString(table, "create_by")
	r.UpdateBy = field.NewString(table, "update_by")
	r.CreateAt = field.NewTime(table, "create_at")
	r.UpdateAt = field.NewTime(table, "update_at")
	r.DeleteAt = field.NewTime(table, "delete_at")
	r.Version = field.NewInt32(table, "version")
	r.AppealID = field.NewInt64(table, "appeal_id")
	r.ReviewID = field.NewInt64(table, "review_id")
	r.StoreID = field.NewInt64(table, "store_id")
	r.Status = field.NewInt32(table, "status")
	r.Reason = field.NewString(table, "reason")
	r.Content = field.NewString(table, "content")
	r.PicInfo = field.NewString(table, "pic_info")
	r.VideoInfo = field.NewString(table, "video_info")
	r.OpRemarks = field.NewString(table, "op_remarks")
	r.OpUser = field.NewString(table, "op_user")
	r.ExtJSON = field.NewString(table, "ext_json")
	r.CtrlJSON = field.NewString(table, "ctrl_json")

	r.fillFieldMap()

	return r
}

func (r *reviewAppealInfo) WithContext(ctx context.Context) IReviewAppealInfoDo {
	return r.reviewAppealInfoDo.WithContext(ctx)
}

func (r reviewAppealInfo) TableName() string { return r.reviewAppealInfoDo.TableName() }

func (r reviewAppealInfo) Alias() string { return r.reviewAppealInfoDo.Alias() }

func (r reviewAppealInfo) Columns(cols ...field.Expr) gen.Columns {
	return r.reviewAppealInfoDo.Columns(cols...)
}

func (r *reviewAppealInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reviewAppealInfo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 19)
	r.fieldMap["id"] = r.ID
	r.fieldMap["create_by"] = r.CreateBy
	r.fieldMap["update_by"] = r.UpdateBy
	r.fieldMap["create_at"] = r.CreateAt
	r.fieldMap["update_at"] = r.UpdateAt
	r.fieldMap["delete_at"] = r.DeleteAt
	r.fieldMap["version"] = r.Version
	r.fieldMap["appeal_id"] = r.AppealID
	r.fieldMap["review_id"] = r.ReviewID
	r.fieldMap["store_id"] = r.StoreID
	r.fieldMap["status"] = r.Status
	r.fieldMap["reason"] = r.Reason
	r.fieldMap["content"] = r.Content
	r.fieldMap["pic_info"] = r.PicInfo
	r.fieldMap["video_info"] = r.VideoInfo
	r.fieldMap["op_remarks"] = r.OpRemarks
	r.fieldMap["op_user"] = r.OpUser
	r.fieldMap["ext_json"] = r.ExtJSON
	r.fieldMap["ctrl_json"] = r.CtrlJSON
}

func (r reviewAppealInfo) clone(db *gorm.DB) reviewAppealInfo {
	r.reviewAppealInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reviewAppealInfo) replaceDB(db *gorm.DB) reviewAppealInfo {
	r.reviewAppealInfoDo.ReplaceDB(db)
	return r
}

type reviewAppealInfoDo struct{ gen.DO }

type IReviewAppealInfoDo interface {
	gen.SubQuery
	Debug() IReviewAppealInfoDo
	WithContext(ctx context.Context) IReviewAppealInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReviewAppealInfoDo
	WriteDB() IReviewAppealInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReviewAppealInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReviewAppealInfoDo
	Not(conds ...gen.Condition) IReviewAppealInfoDo
	Or(conds ...gen.Condition) IReviewAppealInfoDo
	Select(conds ...field.Expr) IReviewAppealInfoDo
	Where(conds ...gen.Condition) IReviewAppealInfoDo
	Order(conds ...field.Expr) IReviewAppealInfoDo
	Distinct(cols ...field.Expr) IReviewAppealInfoDo
	Omit(cols ...field.Expr) IReviewAppealInfoDo
	Join(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo
	Group(cols ...field.Expr) IReviewAppealInfoDo
	Having(conds ...gen.Condition) IReviewAppealInfoDo
	Limit(limit int) IReviewAppealInfoDo
	Offset(offset int) IReviewAppealInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewAppealInfoDo
	Unscoped() IReviewAppealInfoDo
	Create(values ...*model.ReviewAppealInfo) error
	CreateInBatches(values []*model.ReviewAppealInfo, batchSize int) error
	Save(values ...*model.ReviewAppealInfo) error
	First() (*model.ReviewAppealInfo, error)
	Take() (*model.ReviewAppealInfo, error)
	Last() (*model.ReviewAppealInfo, error)
	Find() ([]*model.ReviewAppealInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewAppealInfo, err error)
	FindInBatches(result *[]*model.ReviewAppealInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ReviewAppealInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReviewAppealInfoDo
	Assign(attrs ...field.AssignExpr) IReviewAppealInfoDo
	Joins(fields ...field.RelationField) IReviewAppealInfoDo
	Preload(fields ...field.RelationField) IReviewAppealInfoDo
	FirstOrInit() (*model.ReviewAppealInfo, error)
	FirstOrCreate() (*model.ReviewAppealInfo, error)
	FindByPage(offset int, limit int) (result []*model.ReviewAppealInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReviewAppealInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reviewAppealInfoDo) Debug() IReviewAppealInfoDo {
	return r.withDO(r.DO.Debug())
}

func (r reviewAppealInfoDo) WithContext(ctx context.Context) IReviewAppealInfoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reviewAppealInfoDo) ReadDB() IReviewAppealInfoDo {
	return r.Clauses(dbresolver.Read)
}

func (r reviewAppealInfoDo) WriteDB() IReviewAppealInfoDo {
	return r.Clauses(dbresolver.Write)
}

func (r reviewAppealInfoDo) Session(config *gorm.Session) IReviewAppealInfoDo {
	return r.withDO(r.DO.Session(config))
}

func (r reviewAppealInfoDo) Clauses(conds ...clause.Expression) IReviewAppealInfoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reviewAppealInfoDo) Returning(value interface{}, columns ...string) IReviewAppealInfoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reviewAppealInfoDo) Not(conds ...gen.Condition) IReviewAppealInfoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reviewAppealInfoDo) Or(conds ...gen.Condition) IReviewAppealInfoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reviewAppealInfoDo) Select(conds ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reviewAppealInfoDo) Where(conds ...gen.Condition) IReviewAppealInfoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reviewAppealInfoDo) Order(conds ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reviewAppealInfoDo) Distinct(cols ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reviewAppealInfoDo) Omit(cols ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reviewAppealInfoDo) Join(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reviewAppealInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reviewAppealInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reviewAppealInfoDo) Group(cols ...field.Expr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reviewAppealInfoDo) Having(conds ...gen.Condition) IReviewAppealInfoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reviewAppealInfoDo) Limit(limit int) IReviewAppealInfoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reviewAppealInfoDo) Offset(offset int) IReviewAppealInfoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reviewAppealInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewAppealInfoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reviewAppealInfoDo) Unscoped() IReviewAppealInfoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reviewAppealInfoDo) Create(values ...*model.ReviewAppealInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reviewAppealInfoDo) CreateInBatches(values []*model.ReviewAppealInfo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reviewAppealInfoDo) Save(values ...*model.ReviewAppealInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reviewAppealInfoDo) First() (*model.ReviewAppealInfo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewAppealInfo), nil
	}
}

func (r reviewAppealInfoDo) Take() (*model.ReviewAppealInfo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewAppealInfo), nil
	}
}

func (r reviewAppealInfoDo) Last() (*model.ReviewAppealInfo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewAppealInfo), nil
	}
}

func (r reviewAppealInfoDo) Find() ([]*model.ReviewAppealInfo, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReviewAppealInfo), err
}

func (r reviewAppealInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewAppealInfo, err error) {
	buf := make([]*model.ReviewAppealInfo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reviewAppealInfoDo) FindInBatches(result *[]*model.ReviewAppealInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reviewAppealInfoDo) Attrs(attrs ...field.AssignExpr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reviewAppealInfoDo) Assign(attrs ...field.AssignExpr) IReviewAppealInfoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reviewAppealInfoDo) Joins(fields ...field.RelationField) IReviewAppealInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reviewAppealInfoDo) Preload(fields ...field.RelationField) IReviewAppealInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reviewAppealInfoDo) FirstOrInit() (*model.ReviewAppealInfo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewAppealInfo), nil
	}
}

func (r reviewAppealInfoDo) FirstOrCreate() (*model.ReviewAppealInfo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewAppealInfo), nil
	}
}

func (r reviewAppealInfoDo) FindByPage(offset int, limit int) (result []*model.ReviewAppealInfo, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reviewAppealInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reviewAppealInfoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reviewAppealInfoDo) Delete(models ...*model.ReviewAppealInfo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reviewAppealInfoDo) withDO(do gen.Dao) *reviewAppealInfoDo {
	r.DO = *do.(*gen.DO)
	return r
}
