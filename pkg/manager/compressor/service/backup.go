//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"reflect"

	gback "github.com/vdaas/vald/apis/grpc/v1/manager/backup"
	"github.com/vdaas/vald/apis/grpc/v1/payload"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/net/grpc"
)

type Backup interface {
	Start(ctx context.Context) (<-chan error, error)
	GetObject(ctx context.Context, uuid string) (*payload.Backup_Compressed_Vector, error)
	GetLocation(ctx context.Context, uuid string) ([]string, error)
	Register(ctx context.Context, vec *payload.Backup_Compressed_Vector) error
	RegisterMultiple(ctx context.Context, vecs *payload.Backup_Compressed_Vectors) error
	Remove(ctx context.Context, uuid string) error
	RemoveMultiple(ctx context.Context, uuids ...string) error
	RegisterIPs(ctx context.Context, uuid string, ips []string) error
	RemoveIPs(ctx context.Context, ips []string) error
}

type backup struct {
	client grpc.Client
}

func NewBackup(opts ...BackupOption) (bu Backup, err error) {
	b := new(backup)
	for _, opt := range append(defaultBackupOpts, opts...) {
		if err := opt(b); err != nil {
			return nil, errors.ErrOptionFailed(err, reflect.ValueOf(opt))
		}
	}

	return b, nil
}

func (b *backup) Start(ctx context.Context) (<-chan error, error) {
	return b.client.StartConnectionMonitor(ctx)
}

func (b *backup) GetObject(ctx context.Context, uuid string) (vec *payload.Backup_Compressed_Vector, err error) {
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		vec, err = gback.NewBackupClient(conn).GetVector(ctx, &payload.Backup_GetVector_Request{
			Uuid: uuid,
		}, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) GetLocation(ctx context.Context, uuid string) (ipList []string, err error) {
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		ips, err := gback.NewBackupClient(conn).Locations(ctx, &payload.Backup_Locations_Request{
			Uuid: uuid,
		}, copts...)
		if err != nil {
			return nil, err
		}
		ipList = ips.GetIp()
		return
	})
	return
}

func (b *backup) Register(ctx context.Context, vec *payload.Backup_Compressed_Vector) (err error) {
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).Register(ctx, vec, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) RegisterMultiple(ctx context.Context, vecs *payload.Backup_Compressed_Vectors) (err error) {
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).RegisterMulti(ctx, vecs, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) Remove(ctx context.Context, uuid string) (err error) {
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).Remove(ctx, &payload.Backup_Remove_Request{
			Uuid: uuid,
		}, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) RemoveMultiple(ctx context.Context, uuids ...string) (err error) {
	req := new(payload.Backup_Remove_RequestMulti)
	req.Uuids = uuids
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).RemoveMulti(ctx, req, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) RegisterIPs(ctx context.Context, uuid string, ips []string) (err error) {
	req := new(payload.Backup_IP_Register_Request)
	req.Uuid = uuid
	req.Ips = ips
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).RegisterIPs(ctx, req, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}

func (b *backup) RemoveIPs(ctx context.Context, ips []string) (err error) {
	req := new(payload.Backup_IP_Remove_Request)
	req.Ips = ips
	_, err = b.client.RoundRobin(ctx, func(ctx context.Context,
		conn *grpc.ClientConn, copts ...grpc.CallOption) (i interface{}, err error) {
		_, err = gback.NewBackupClient(conn).RemoveIPs(ctx, req, copts...)
		if err != nil {
			return nil, err
		}
		return
	})
	return
}
