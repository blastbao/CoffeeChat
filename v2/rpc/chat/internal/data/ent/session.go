// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/internal/data/ent/session"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Updated holds the value of the "updated" field.
	Updated time.Time `json:"updated,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// PeerID holds the value of the "peer_id" field.
	PeerID string `json:"peer_id,omitempty"`
	// SessionType holds the value of the "session_type" field.
	SessionType int8 `json:"session_type,omitempty"`
	// SessionStatus holds the value of the "session_status" field.
	SessionStatus int8 `json:"session_status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID, session.FieldSessionType, session.FieldSessionStatus:
			values[i] = new(sql.NullInt64)
		case session.FieldUserID, session.FieldPeerID:
			values[i] = new(sql.NullString)
		case session.FieldCreated, session.FieldUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int32(value.Int64)
		case session.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				s.Created = value.Time
			}
		case session.FieldUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated", values[i])
			} else if value.Valid {
				s.Updated = value.Time
			}
		case session.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = value.String
			}
		case session.FieldPeerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field peer_id", values[i])
			} else if value.Valid {
				s.PeerID = value.String
			}
		case session.FieldSessionType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field session_type", values[i])
			} else if value.Valid {
				s.SessionType = int8(value.Int64)
			}
		case session.FieldSessionStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field session_status", values[i])
			} else if value.Valid {
				s.SessionStatus = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created=")
	builder.WriteString(s.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated=")
	builder.WriteString(s.Updated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(s.UserID)
	builder.WriteString(", ")
	builder.WriteString("peer_id=")
	builder.WriteString(s.PeerID)
	builder.WriteString(", ")
	builder.WriteString("session_type=")
	builder.WriteString(fmt.Sprintf("%v", s.SessionType))
	builder.WriteString(", ")
	builder.WriteString("session_status=")
	builder.WriteString(fmt.Sprintf("%v", s.SessionStatus))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
