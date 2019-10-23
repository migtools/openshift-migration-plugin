package migcommon

// MigrationRegistry migration registry
const MigrationRegistry string = "openshift.io/migration-registry"

// SwingPVAnnotation swing PV annotation
const SwingPVAnnotation string = "openshift.io/swing-pv"

// MigrateTypeAnnotation migrate type annotation
// copy, swing, TODO: others (snapshot, custom, etc.)
const MigrateTypeAnnotation string = "openshift.io/migrate-type"

// MigrateStorageClassAnnotation target storage class
const MigrateStorageClassAnnotation string = "openshift.io/target-storage-class"

// MigrateAccessModeAnnotation target access mode
const MigrateAccessModeAnnotation string = "openshift.io/target-access-mode"

// MigrateCopyPhaseAnnotation stage, final. Only valid for copy type.
const MigrateCopyPhaseAnnotation string = "openshift.io/migrate-copy-phase"

// MigrateQuiesceAnnotation migrate quiesce annotation
const MigrateQuiesceAnnotation string = "openshift.io/migrate-quiesce-pods"

// PodStageLabel pod stage label
const PodStageLabel string = "migration-stage-pod"

// PVCSelectedNodeAnnotation kubernetes PVC annotations
const PVCSelectedNodeAnnotation string = "volume.kubernetes.io/selected-node"

// ResticRestoreAnnotationPrefix restic annotation
const ResticRestoreAnnotationPrefix string = "snapshot.velero.io"

// ResticBackupAnnotation restic annotation
const ResticBackupAnnotation string = "backup.velero.io/backup-volumes"

// NamespaceSCCAnnotationMCS namespace SCC annotation
const NamespaceSCCAnnotationMCS string = "openshift.io/sa.scc.mcs"

// NamespaceSCCAnnotationGroups namespace SCC annotation
const NamespaceSCCAnnotationGroups string = "openshift.io/sa.scc.supplemental-groups"

// NamespaceSCCAnnotationUIDRange namespace SCC annotation
const NamespaceSCCAnnotationUIDRange string = "openshift.io/sa.scc.uid-range"
