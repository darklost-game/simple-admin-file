// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-file/ent/cloudfiletag"
)

// CloudFileTag is the model entity for the CloudFileTag schema.
type CloudFileTag struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// CloudFileTag's name | 标签名称
	Name string `json:"name,omitempty"`
	// The remark of tag | 标签的备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CloudFileTagQuery when eager-loading is set.
	Edges        CloudFileTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CloudFileTagEdges holds the relations/edges for other nodes in the graph.
type CloudFileTagEdges struct {
	// CloudFiles holds the value of the cloud_files edge.
	CloudFiles []*CloudFile `json:"cloud_files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CloudFilesOrErr returns the CloudFiles value or an error if the edge
// was not loaded in eager-loading.
func (e CloudFileTagEdges) CloudFilesOrErr() ([]*CloudFile, error) {
	if e.loadedTypes[0] {
		return e.CloudFiles, nil
	}
	return nil, &NotLoadedError{edge: "cloud_files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CloudFileTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cloudfiletag.FieldID, cloudfiletag.FieldStatus:
			values[i] = new(sql.NullInt64)
		case cloudfiletag.FieldName, cloudfiletag.FieldRemark:
			values[i] = new(sql.NullString)
		case cloudfiletag.FieldCreatedAt, cloudfiletag.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CloudFileTag fields.
func (cft *CloudFileTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cloudfiletag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cft.ID = uint64(value.Int64)
		case cloudfiletag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cft.CreatedAt = value.Time
			}
		case cloudfiletag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cft.UpdatedAt = value.Time
			}
		case cloudfiletag.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cft.Status = uint8(value.Int64)
			}
		case cloudfiletag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cft.Name = value.String
			}
		case cloudfiletag.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				cft.Remark = value.String
			}
		default:
			cft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CloudFileTag.
// This includes values selected through modifiers, order, etc.
func (cft *CloudFileTag) Value(name string) (ent.Value, error) {
	return cft.selectValues.Get(name)
}

// QueryCloudFiles queries the "cloud_files" edge of the CloudFileTag entity.
func (cft *CloudFileTag) QueryCloudFiles() *CloudFileQuery {
	return NewCloudFileTagClient(cft.config).QueryCloudFiles(cft)
}

// Update returns a builder for updating this CloudFileTag.
// Note that you need to call CloudFileTag.Unwrap() before calling this method if this CloudFileTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (cft *CloudFileTag) Update() *CloudFileTagUpdateOne {
	return NewCloudFileTagClient(cft.config).UpdateOne(cft)
}

// Unwrap unwraps the CloudFileTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cft *CloudFileTag) Unwrap() *CloudFileTag {
	_tx, ok := cft.config.driver.(*txDriver)
	if !ok {
		panic("ent: CloudFileTag is not a transactional entity")
	}
	cft.config.driver = _tx.drv
	return cft
}

// String implements the fmt.Stringer.
func (cft *CloudFileTag) String() string {
	var builder strings.Builder
	builder.WriteString("CloudFileTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cft.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cft.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cft.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cft.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cft.Name)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(cft.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// CloudFileTags is a parsable slice of CloudFileTag.
type CloudFileTags []*CloudFileTag