// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	domain "github.com/TovarischSuhov/pos-telegram/domain"
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

func easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels(in *jlexer.Lexer, out *User) {
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
			out.ID = domain.UserID(in.Uint64())
		case "is_bot":
			out.IsBot = bool(in.Bool())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "language_code":
			out.LanguageCode = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"is_bot\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsBot))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"language_code\":"
		out.RawString(prefix)
		out.String(string(in.LanguageCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels(l, v)
}
func easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels1(in *jlexer.Lexer, out *Update) {
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
		case "update_id":
			out.UpdateID = int(in.Int())
		case "message":
			(out.Message).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels1(out *jwriter.Writer, in Update) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"update_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.UpdateID))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		(in.Message).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Update) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Update) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Update) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Update) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels1(l, v)
}
func easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels2(in *jlexer.Lexer, out *Message) {
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
		case "message_id":
			out.MessageID = domain.MessageID(in.Uint64())
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "sender_chat":
			(out.SenderChat).UnmarshalEasyJSON(in)
		case "date":
			out.Date = int64(in.Int64())
		case "chat":
			(out.Chat).UnmarshalEasyJSON(in)
		case "test":
			out.Test = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels2(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.MessageID))
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		(in.From).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"sender_chat\":"
		out.RawString(prefix)
		(in.SenderChat).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Int64(int64(in.Date))
	}
	{
		const prefix string = ",\"chat\":"
		out.RawString(prefix)
		(in.Chat).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"test\":"
		out.RawString(prefix)
		out.String(string(in.Test))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels2(l, v)
}
func easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels3(in *jlexer.Lexer, out *GetUpdateRequest) {
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
		case "offset":
			out.Offset = domain.Offset(in.Uint64())
		case "limit":
			out.Limit = int64(in.Int64())
		case "timeout":
			out.Timeout = int64(in.Int64())
		case "allowed_updates":
			if in.IsNull() {
				in.Skip()
				out.AllowedUpdates = nil
			} else {
				in.Delim('[')
				if out.AllowedUpdates == nil {
					if !in.IsDelim(']') {
						out.AllowedUpdates = make([]string, 0, 4)
					} else {
						out.AllowedUpdates = []string{}
					}
				} else {
					out.AllowedUpdates = (out.AllowedUpdates)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AllowedUpdates = append(out.AllowedUpdates, v1)
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
func easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels3(out *jwriter.Writer, in GetUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Offset))
	}
	if in.Limit != 0 {
		const prefix string = ",\"limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Limit))
	}
	if in.Timeout != 0 {
		const prefix string = ",\"timeout\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Timeout))
	}
	if len(in.AllowedUpdates) != 0 {
		const prefix string = ",\"allowed_updates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.AllowedUpdates {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels3(l, v)
}
func easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels4(in *jlexer.Lexer, out *Chat) {
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
			out.ID = int64(in.Int64())
		case "type":
			out.Type = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels4(out *jwriter.Writer, in Chat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Chat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComTovarischSuhovPosTelegramTelegramModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComTovarischSuhovPosTelegramTelegramModels4(l, v)
}
