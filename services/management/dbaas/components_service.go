// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package dbaas

import (
	"context"

	dbaasv1beta1 "github.com/percona/pmm/api/managementpb/dbaas"
	"github.com/percona/pmm/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

type componentsService struct {
	l                    *logrus.Entry
	db                   *reform.DB
	dbaasClient          dbaasClient
	versionServiceClient versionService
}

// NewComponentsService creates Components Service.
func NewComponentsService(db *reform.DB, dbaasClient dbaasClient, versionServiceClient versionService) dbaasv1beta1.ComponentsServer {
	l := logrus.WithField("component", "components_service")
	return &componentsService{
		l:                    l,
		db:                   db,
		dbaasClient:          dbaasClient,
		versionServiceClient: versionServiceClient,
	}
}

func (c componentsService) GetPSMDBComponents(ctx context.Context, req *dbaasv1beta1.GetPSMDBComponentsRequest) (*dbaasv1beta1.GetPSMDBComponentsResponse, error) {
	params := componentsParams{
		operator:  psmdbOperator,
		dbVersion: req.DbVersion,
	}
	if req.KubernetesClusterName != "" {
		kubernetesCluster, err := models.FindKubernetesClusterByName(c.db.Querier, req.KubernetesClusterName)
		if err != nil {
			return nil, err
		}

		checkResponse, e := c.dbaasClient.CheckKubernetesClusterConnection(ctx, kubernetesCluster.KubeConfig)
		if e != nil {
			return nil, e
		}

		params.operatorVersion = checkResponse.Operators.Psmdb.Version
	}

	versions, err := c.versions(ctx, params, nil)
	if err != nil {
		return nil, err
	}
	return &dbaasv1beta1.GetPSMDBComponentsResponse{Versions: versions}, nil
}

func (c componentsService) GetPXCComponents(ctx context.Context, req *dbaasv1beta1.GetPXCComponentsRequest) (*dbaasv1beta1.GetPXCComponentsResponse, error) {
	var kubernetesCluster *models.KubernetesCluster
	params := componentsParams{
		operator:  pxcOperator,
		dbVersion: req.DbVersion,
	}
	if req.KubernetesClusterName != "" {
		var err error
		kubernetesCluster, err = models.FindKubernetesClusterByName(c.db.Querier, req.KubernetesClusterName)
		if err != nil {
			return nil, err
		}

		checkResponse, e := c.dbaasClient.CheckKubernetesClusterConnection(ctx, kubernetesCluster.KubeConfig)
		if e != nil {
			return nil, e
		}

		params.operatorVersion = checkResponse.Operators.Xtradb.Version
	}

	versions, err := c.versions(ctx, params, kubernetesCluster)
	if err != nil {
		return nil, err
	}
	return &dbaasv1beta1.GetPXCComponentsResponse{Versions: versions}, nil
}

func (c componentsService) SetPSMDBComponents(ctx context.Context, req *dbaasv1beta1.SetPSMDBComponentsRequest) (*dbaasv1beta1.SetPSMDBComponentsResponse, error) {
	err := c.db.InTransaction(func(tx *reform.TX) error {
		kubernetesCluster, e := models.FindKubernetesClusterByName(tx.Querier, req.KubernetesClusterName)
		if e != nil {
			return e
		}

		if req.Mongod != nil {
			kubernetesCluster.Mongod = setComponent(kubernetesCluster.Mongod, req.Mongod)
		}

		e = tx.Save(kubernetesCluster)
		if e != nil {
			return e
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &dbaasv1beta1.SetPSMDBComponentsResponse{}, nil
}

func (c componentsService) SetPXCComponents(ctx context.Context, req *dbaasv1beta1.SetPXCComponentsRequest) (*dbaasv1beta1.SetPXCComponentsResponse, error) {
	err := c.db.InTransaction(func(tx *reform.TX) error {
		kubernetesCluster, e := models.FindKubernetesClusterByName(tx.Querier, req.KubernetesClusterName)
		if e != nil {
			return e
		}

		if req.Pxc != nil {
			kubernetesCluster.PXC = setComponent(kubernetesCluster.PXC, req.Pxc)
		}

		if req.Proxysql != nil {
			kubernetesCluster.ProxySQL = setComponent(kubernetesCluster.ProxySQL, req.Proxysql)
		}

		e = tx.Save(kubernetesCluster)
		if e != nil {
			return e
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &dbaasv1beta1.SetPXCComponentsResponse{}, nil
}

func (c componentsService) versions(ctx context.Context, params componentsParams, cluster *models.KubernetesCluster) ([]*dbaasv1beta1.Version, error) {
	components, err := c.versionServiceClient.Matrix(ctx, params)
	if err != nil {
		return nil, err
	}

	var mongod, pxc, proxySQL *models.Component
	if cluster != nil {
		mongod = cluster.Mongod
		pxc = cluster.PXC
		proxySQL = cluster.ProxySQL
	}

	versions := make([]*dbaasv1beta1.Version, 0, len(components.Versions))
	for _, v := range components.Versions {
		respVersion := &dbaasv1beta1.Version{
			Product:  v.Product,
			Operator: v.Operator,
			Matrix: &dbaasv1beta1.Matrix{
				Mongod:       c.matrix(v.Matrix.Mongod, mongod),
				Pxc:          c.matrix(v.Matrix.Pxc, pxc),
				Pmm:          c.matrix(v.Matrix.Pmm, nil),
				Proxysql:     c.matrix(v.Matrix.Proxysql, proxySQL),
				Haproxy:      c.matrix(v.Matrix.Haproxy, nil),
				Backup:       c.matrix(v.Matrix.Backup, nil),
				Operator:     c.matrix(v.Matrix.Operator, nil),
				LogCollector: c.matrix(v.Matrix.LogCollector, nil),
			},
		}
		versions = append(versions, respVersion)
	}

	return versions, nil
}

func (c componentsService) matrix(m map[string]component, kc *models.Component) map[string]*dbaasv1beta1.Component {
	result := make(map[string]*dbaasv1beta1.Component)

	var lastVersion string
	lastVersionParsed := version.MustParse("0.0.0")
	for v, component := range m {
		result[v] = &dbaasv1beta1.Component{
			ImagePath: component.ImagePath,
			ImageHash: component.ImageHash,
			Status:    component.Status,
			Critical:  component.Critical,
		}
		parsedVersion, err := version.Parse(v)
		if err != nil {
			c.l.Warnf("couldn't parse version %s: %s", v, err.Error())
		}
		if lastVersionParsed.Less(parsedVersion) && component.Status == "recommended" {
			lastVersionParsed = parsedVersion
			lastVersion = v
		}
	}

	defaultVersionSet := false
	if kc != nil {
		if _, ok := result[kc.DefaultVersion]; ok {
			result[kc.DefaultVersion].Default = true
			defaultVersionSet = true
		}
		for _, v := range kc.DisabledVersions {
			if _, ok := result[v]; ok {
				result[v].Disabled = true
			}
		}
	}
	if lastVersion != "" && !defaultVersionSet {
		result[lastVersion].Default = true
	}
	return result
}

func setComponent(kc *models.Component, rc *dbaasv1beta1.SetComponent) *models.Component {
	if kc == nil {
		kc = new(models.Component)
	}
	kc.DefaultVersion = rc.DefaultVersion

	disabledVersions := make(map[string]struct{})
	for _, v := range kc.DisabledVersions {
		disabledVersions[v] = struct{}{}
	}
	for _, v := range rc.GetDisableVersions() {
		disabledVersions[v] = struct{}{}
	}
	for _, v := range rc.GetEnableVersions() {
		delete(disabledVersions, v)
	}
	kc.DisabledVersions = make([]string, 0)
	for v := range disabledVersions {
		kc.DisabledVersions = append(kc.DisabledVersions, v)
	}
	return kc
}
