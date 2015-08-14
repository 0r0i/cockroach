// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (peter@cockroachlabs.com)

// This code was derived from https://github.com/youtube/vitess.
//
// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package parser

import (
	"bytes"
	"fmt"
)

func (*RenameDatabase) statement() {}

// RenameDatabase represents a RENAME DATABASE statement.
type RenameDatabase struct {
	Name    Name
	NewName Name
}

func (node *RenameDatabase) String() string {
	return fmt.Sprintf("ALTER DATABASE %s RENAME TO %s", node.Name, node.NewName)
}

func (*RenameTable) statement() {}

// RenameTable represents a RENAME TABLE statement.
type RenameTable struct {
	Name     *QualifiedName
	NewName  Name
	IfExists bool
}

func (node *RenameTable) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ALTER TABLE ")
	if node.IfExists {
		_, _ = buf.WriteString("IF EXISTS ")
	}
	_, _ = buf.WriteString(fmt.Sprintf("%s RENAME TO %s", node.Name, node.NewName))
	return buf.String()
}
