// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package git

import (
	"time"

	billy "gopkg.in/src-d/go-billy.v4"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage"
)

// Repository is a handful of git.Repository functions reified as an interface to facilitate testing.
type Repository interface {
	CommitObject(h plumbing.Hash) (*object.Commit, error)
	Tags() (storer.ReferenceIter, error)
	Fetch(o *git.FetchOptions) error
}

// Git is a handful of git functions reified as an interface to facilitate testing.
type Git interface {
	Clone(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (Repository, error)
}

// GoGit is the Git implementation for src-d/go-git.
type GoGit struct{}

// Clone creates a clone of a git repository.
func (GoGit) Clone(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (Repository, error) {
	return git.Clone(s, worktree, o)
}

// Revision provides read access to hash and commit time git commit metadata.
type Revision interface {
	GetHash() plumbing.Hash
	GetCommitTime() time.Time
}

// RevisionData is a vanilla data structure implementation of Revision.
type RevisionData struct {
	Hash       plumbing.Hash
	CommitTime time.Time
}

// GetHash returns the hash stored in RevisionData.Hash.
func (d RevisionData) GetHash() plumbing.Hash {
	return d.Hash
}

// GetCommitTime returns the commit time stored in RevisionData.CommitTime.
func (d RevisionData) GetCommitTime() time.Time {
	return d.CommitTime
}
