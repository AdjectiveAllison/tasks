// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel(in *jlexer.Lexer, out *Task) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "links":
			if in.IsNull() {
				in.Skip()
				out.Links = nil
			} else {
				in.Delim('[')
				if out.Links == nil {
					if !in.IsDelim(']') {
						out.Links = make([]string, 0, 4)
					} else {
						out.Links = []string{}
					}
				} else {
					out.Links = (out.Links)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Links = append(out.Links, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "updatedAt":
			out.UpdatedAt = uint64(in.Uint64())
		case "completed":
			out.Completed = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"links\":"
		out.RawString(prefix)
		if in.Links == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Links {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UpdatedAt))
	}
	{
		const prefix string = ",\"completed\":"
		out.RawString(prefix)
		out.Bool(bool(in.Completed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Task) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Task) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel(l, v)
}
func easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel1(in *jlexer.Lexer, out *ListTasksResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]Task, 0, 0)
					} else {
						out.Tasks = []Task{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Task
					(v4).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel1(out *jwriter.Writer, in ListTasksResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tasks\":"
		out.RawString(prefix[1:])
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tasks {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListTasksResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListTasksResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComAdjectiveAllisonTasksAppModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListTasksResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListTasksResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComAdjectiveAllisonTasksAppModel1(l, v)
}