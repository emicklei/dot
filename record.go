package dot

import (
	"fmt"
	"strings"
)

// RecordBuilder can build the label of a node with shape "record" or "mrecord".
type RecordBuilder struct {
	target       Node
	shape        string
	nesting      *stack
	currentLabel recordLabel
}

// NewRecordBuilder returns a new RecordBuilder for constructing the label of the node.
func NewRecordBuilder(n Node) *RecordBuilder {
	return &RecordBuilder{
		target:  n,
		shape:   "record",
		nesting: new(stack),
	}
}

type recordLabel []recordField

func (r recordLabel) writeOn(buf *strings.Builder) {
	for i, each := range r {
		if i > 0 {
			buf.WriteRune('|')
		}
		each.writeOn(buf)
	}
}

type recordField struct {
	id recordFieldId
	// or
	nestedLabel *recordLabel
}

func (r recordField) writeOn(buf *strings.Builder) {
	if r.nestedLabel != nil {
		buf.WriteRune('{')
		r.nestedLabel.writeOn(buf)
		buf.WriteRune('}')
		return
	}
	r.id.writeOn(buf)
}

type recordFieldId struct {
	id      string
	content string
}

func (r recordFieldId) writeOn(buf *strings.Builder) {
	if r.id != "" {
		fmt.Fprintf(buf, "<%s> ", r.id)
	}
	buf.WriteString(r.content)
}

// MRecord sets the shape of the node to "mrecord"
func (r *RecordBuilder) MRecord() *RecordBuilder {
	r.shape = "mrecord"
	return r
}

// Field adds a record field
func (r *RecordBuilder) Field(content string) *RecordBuilder {
	rf := recordField{
		id: recordFieldId{
			content: content,
		},
	}
	r.currentLabel = append(r.currentLabel, rf)
	return r
}

// FieldWithId adds a record field with an identifier for connecting edges.
func (r *RecordBuilder) FieldWithId(content, id string) *RecordBuilder {
	rf := recordField{
		id: recordFieldId{
			id:      id,
			content: content,
		},
	}
	r.currentLabel = append(r.currentLabel, rf)
	return r
}

// Nesting will create a nested (layout flipped) list of rlabel.
func (r *RecordBuilder) Nesting(block func()) {
	r.nesting.push(r.currentLabel)
	r.currentLabel = recordLabel{}
	block()
	// currentLabel has fields added by block
	// top of stack has label before block
	top := r.nesting.pop()
	cpy := r.currentLabel[:]
	top = append(top, recordField{
		nestedLabel: &cpy,
	})
	r.currentLabel = top
}

// Build sets the computed label and shape
func (r *RecordBuilder) Build() error {
	r.target.Attr("shape", r.shape)
	r.target.Attr("label", r.Label())
	return nil
}

// Label returns the computed label
func (r *RecordBuilder) Label() string {
	buf := new(strings.Builder)
	for i, each := range r.currentLabel {
		if i > 0 {
			buf.WriteString("|")
		}
		each.writeOn(buf)
	}
	return buf.String()
}

// stack implements a lifo queue for recordLabel instances.
type stack []recordLabel

func (s *stack) push(r recordLabel) {
	*s = append(*s, r)
}
func (s *stack) pop() recordLabel {
	top := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return top
}
