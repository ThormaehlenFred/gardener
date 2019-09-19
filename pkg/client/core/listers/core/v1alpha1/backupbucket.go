// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupBucketLister helps list BackupBuckets.
type BackupBucketLister interface {
	// List lists all BackupBuckets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BackupBucket, err error)
	// Get retrieves the BackupBucket from the index for a given name.
	Get(name string) (*v1alpha1.BackupBucket, error)
	BackupBucketListerExpansion
}

// backupBucketLister implements the BackupBucketLister interface.
type backupBucketLister struct {
	indexer cache.Indexer
}

// NewBackupBucketLister returns a new BackupBucketLister.
func NewBackupBucketLister(indexer cache.Indexer) BackupBucketLister {
	return &backupBucketLister{indexer: indexer}
}

// List lists all BackupBuckets in the indexer.
func (s *backupBucketLister) List(selector labels.Selector) (ret []*v1alpha1.BackupBucket, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupBucket))
	})
	return ret, err
}

// Get retrieves the BackupBucket from the index for a given name.
func (s *backupBucketLister) Get(name string) (*v1alpha1.BackupBucket, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backupbucket"), name)
	}
	return obj.(*v1alpha1.BackupBucket), nil
}