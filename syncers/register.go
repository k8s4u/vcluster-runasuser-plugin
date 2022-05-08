package syncers

import (
	"fmt"

	"github.com/loft-sh/vcluster-sdk/applier"
	"github.com/loft-sh/vcluster-sdk/log"
	"github.com/loft-sh/vcluster-sdk/syncer"
	synccontext "github.com/loft-sh/vcluster-sdk/syncer/context"
)

const (
	RegisterManifestPath = "/manifests/register.yaml"
)

func NewRegisterSyncer(ctx *synccontext.RegisterContext) syncer.Base {
	return &RegisterSyncer{}
}

type RegisterSyncer struct{}

var _ syncer.Base = &RegisterSyncer{}

func (s *RegisterSyncer) Name() string {
	return "runasuser-plugin-init"
}

var _ syncer.Initializer = &RegisterSyncer{}

func (s *RegisterSyncer) Init(ctx *synccontext.RegisterContext) error {
	err := applier.ApplyManifestFile(ctx.VirtualManager.GetConfig(), RegisterManifestPath)
	if err == nil {
		log.New(s.Name()).Infof("Successfully applied manifest %s", RegisterManifestPath)
	} else {
		err = fmt.Errorf("failed to apply manifest %s: %v", RegisterManifestPath, err)
	}
	return err
}
