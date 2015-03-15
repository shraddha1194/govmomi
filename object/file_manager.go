/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type FileManager struct {
	c *govmomi.Client
}

func NewFileManager(c *govmomi.Client) *FileManager {
	f := FileManager{
		c: c,
	}

	return &f
}

func (f FileManager) CopyDatastoreFile(sourceName string, sourceDatacenter *Datacenter, destinationName string, destinationDatacenter *Datacenter, force bool) (*Task, error) {
	req := types.CopyDatastoreFile_Task{
		This:            *f.c.ServiceContent.FileManager,
		SourceName:      sourceName,
		DestinationName: destinationName,
		Force:           force,
	}

	if sourceDatacenter != nil {
		ref := sourceDatacenter.Reference()
		req.SourceDatacenter = &ref
	}

	if destinationDatacenter != nil {
		ref := destinationDatacenter.Reference()
		req.DestinationDatacenter = &ref
	}

	res, err := methods.CopyDatastoreFile_Task(context.TODO(), f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

// DeleteDatastoreFile deletes the specified file or folder from the datastore.
func (f FileManager) DeleteDatastoreFile(name string, dc *Datacenter) (*Task, error) {
	req := types.DeleteDatastoreFile_Task{
		This: *f.c.ServiceContent.FileManager,
		Name: name,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	res, err := methods.DeleteDatastoreFile_Task(context.TODO(), f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}

// MakeDirectory creates a folder using the specified name.
func (f FileManager) MakeDirectory(name string, dc *Datacenter, createParentDirectories bool) error {
	req := types.MakeDirectory{
		This: *f.c.ServiceContent.FileManager,
		Name: name,
		CreateParentDirectories: createParentDirectories,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	_, err := methods.MakeDirectory(context.TODO(), f.c, &req)
	return err
}

func (f FileManager) MoveDatastoreFile(sourceName string, sourceDatacenter *Datacenter, destinationName string, destinationDatacenter *Datacenter, force bool) (*Task, error) {
	req := types.MoveDatastoreFile_Task{
		This:            *f.c.ServiceContent.FileManager,
		SourceName:      sourceName,
		DestinationName: destinationName,
		Force:           force,
	}

	if sourceDatacenter != nil {
		ref := sourceDatacenter.Reference()
		req.SourceDatacenter = &ref
	}

	if destinationDatacenter != nil {
		ref := destinationDatacenter.Reference()
		req.DestinationDatacenter = &ref
	}

	res, err := methods.MoveDatastoreFile_Task(context.TODO(), f.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(f.c, res.Returnval), nil
}
