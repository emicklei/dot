package dot

import (
	"fmt"
	"strings"
)

type RecordBuilder struct {
	target       Node
	level        int
	nesting      *stack
	currentLabel recordLabel
}

func NewRecordBuilder(n Node) *RecordBuilder {
	return &RecordBuilder{
		target:  n,
		level:   0,
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

func (r *RecordBuilder) AddField(content string) {
	rf := recordField{
		id: recordFieldId{
			content: content,
		},
	}
	r.currentLabel = append(r.currentLabel, rf)
}

func (r *RecordBuilder) FlipWhile(block func()) {
	r.nesting.push(r.currentLabel)
	r.currentLabel = recordLabel{}
	r.level++
	block()
	// currentlabel has zero or more record fields
	top := r.nesting.pop()
	// TODO
	r.currentLabel = top
	r.level--
}

func (r *RecordBuilder) Build() error {
	r.target.Attr("shape", "record")
	r.target.Attr("label", r.Label())
	return nil
}

func (r *RecordBuilder) Label() string {
	buf := new(strings.Builder)
	for i, each := range r.currentLabel {
		if i > 0 {
			buf.WriteString(" | ")
		}
		each.writeOn(buf)
	}
	return buf.String()
}

type stack []recordLabel

func (s *stack) push(r recordLabel) {
	*s = append(*s, r)
}
func (s *stack) pop() recordLabel {
	top := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return top
}
