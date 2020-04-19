// gomuks - A terminal Matrix client written in Go.
// Copyright (C) 2020 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pushrules_test

import (
	"maunium.net/go/mautrix/event"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushCondition_Match_DisplayName(t *testing.T) {
	evt := newFakeEvent(event.EventMessage, event.Content{
		MsgType: event.MsgText,
		Body:    "tulir: test mention",
	})
	evt.Sender = "@someone_else:matrix.org"
	assert.True(t, displaynamePushCondition.Match(displaynameTestRoom, evt))
}

func TestPushCondition_Match_DisplayName_Fail(t *testing.T) {
	evt := newFakeEvent(event.EventMessage, event.Content{
		MsgType: event.MsgText,
		Body:    "not a mention",
	})
	evt.Sender = "@someone_else:matrix.org"
	assert.False(t, displaynamePushCondition.Match(displaynameTestRoom, evt))
}

func TestPushCondition_Match_DisplayName_FailsOnEmptyRoom(t *testing.T) {
	emptyRoom := newFakeRoom(0)
	evt := newFakeEvent(event.EventMessage, event.Content{
		MsgType: event.MsgText,
		Body:    "tulir: this room doesn't have the owner Member available, so it fails.",
	})
	evt.Sender = "@someone_else:matrix.org"
	assert.False(t, displaynamePushCondition.Match(emptyRoom, evt))
}
