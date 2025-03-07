package dynakube

import (
	"context"

	dynatracev1beta1 "github.com/Dynatrace/dynatrace-operator/pkg/api/v1beta1/dynakube"
	"github.com/Dynatrace/dynatrace-operator/pkg/util/logger"
	"github.com/go-logr/logr"
)

const oneagentEnableVolumeStorageEnvVarName = "ONEAGENT_ENABLE_VOLUME_STORAGE"

var log = logger.Factory.GetLogger("dynakube-validation")

type validator func(ctx context.Context, dv *dynakubeValidator, dynakube *dynatracev1beta1.DynaKube) string

var validators = []validator{
	NoApiUrl,
	IsInvalidApiUrl,
	IsThirdGenAPIUrl,
	missingCSIDaemonSet,
	disabledCSIForReadonlyCSIVolume,
	conflictingActiveGateConfiguration,
	exclusiveSyntheticCapability,
	invalidActiveGateCapabilities,
	duplicateActiveGateCapabilities,
	invalidActiveGateProxyUrl,
	conflictingOneAgentConfiguration,
	conflictingNodeSelector,
	conflictingNamespaceSelector,
	conflictingReadOnlyFilesystemAndMultipleOsAgentsOnNode,
	noResourcesAvailable,
	imageFieldSetWithoutCSIFlag,
	conflictingOneAgentVolumeStorageSettings,
	invalidSyntheticNodeType,
	nameViolatesDNS1035,
	nameTooLong,
	namespaceSelectorMatchLabelsViolateLabelSpec,
}

var warnings = []validator{
	deprecatedFeatureFlagFormat,
	missingActiveGateMemoryLimit,
	deprecatedFeatureFlagDisableActiveGateUpdates,
	deprecatedFeatureFlagDisableActiveGateRawImage,
	deprecatedFeatureFlagDisableHostsRequests,
	deprecatedFeatureFlagDisableReadOnlyAgent,
	deprecatedFeatureFlagDisableWebhookReinvocationPolicy,
	deprecatedFeatureFlagDisableMetadataEnrichment,
	ineffectiveReadOnlyHostFsFeatureFlag,
	syntheticPreviewWarning,
	deprecatedFeatureFlag,
}

func SetLogger(logger logr.Logger) {
	log = logger
}
