// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package summaries

import "github.com/google/go-github/github"

// Pending is the struct for pending.md.
type Pending struct {
	CheckState

	HostName string // Host environment name
	RunsURL  string // URL for the list of test runs
}

// GetCheckState returns the info needed to update a check.
func (p Pending) GetCheckState() CheckState {
	return p.CheckState
}

// GetSummary executes the template for the data.
func (p Pending) GetSummary() (string, error) {
	return compile(&p, "pending.md")
}

// CancelAction is an action that can be taken to cancel a pending check run.
func CancelAction() *github.CheckRunAction {
	return &github.CheckRunAction{
		Identifier:  "cancel",
		Label:       "Cancel",
		Description: "Cancel this pending check run",
	}
}

// GetActions returns the actions that can be taken by the user.
func (p Pending) GetActions() []*github.CheckRunAction {
	return []*github.CheckRunAction{
		CancelAction(),
	}
}
