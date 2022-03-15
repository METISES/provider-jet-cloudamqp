/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	alarm "github.com/timgchile/provider-jet-cloudamqp/internal/controller/cloudamqp/alarm"
	instance "github.com/timgchile/provider-jet-cloudamqp/internal/controller/cloudamqp/instance"
	notification "github.com/timgchile/provider-jet-cloudamqp/internal/controller/cloudamqp/notification"
	plugin "github.com/timgchile/provider-jet-cloudamqp/internal/controller/cloudamqp/plugin"
	webhook "github.com/timgchile/provider-jet-cloudamqp/internal/controller/cloudamqp/webhook"
	domain "github.com/timgchile/provider-jet-cloudamqp/internal/controller/custom/domain"
	log "github.com/timgchile/provider-jet-cloudamqp/internal/controller/integration/log"
	metric "github.com/timgchile/provider-jet-cloudamqp/internal/controller/integration/metric"
	community "github.com/timgchile/provider-jet-cloudamqp/internal/controller/plugin/community"
	providerconfig "github.com/timgchile/provider-jet-cloudamqp/internal/controller/providerconfig"
	firewall "github.com/timgchile/provider-jet-cloudamqp/internal/controller/security/firewall"
	gcppeering "github.com/timgchile/provider-jet-cloudamqp/internal/controller/vpc/gcppeering"
	peering "github.com/timgchile/provider-jet-cloudamqp/internal/controller/vpc/peering"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alarm.Setup,
		instance.Setup,
		notification.Setup,
		plugin.Setup,
		webhook.Setup,
		domain.Setup,
		log.Setup,
		metric.Setup,
		community.Setup,
		providerconfig.Setup,
		firewall.Setup,
		gcppeering.Setup,
		peering.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
