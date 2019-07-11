package migcommon

const MigrationRegistry string = "openshift.io/migration-registry"

const SwingPVAnnotation string = "openshift.io/swing-pv"

// copy, swing, TODO: others (snapshot, custom, etc.)
const MigrateTypeAnnotation string = "openshift.io/migrate-type"

// target storage class
const MigrateStorageClassAnnotation string = "openshift.io/target-storage-class"
//target access mode
const MigrateAccessModeAnnotation string = "openshift.io/target-access-mode"

//stage, final. Only valid for copy type.
const MigrateCopyPhaseAnnotation string = "openshift.io/migrate-copy-phase"

const MigrateQuiesceAnnotation string = "openshift.io/migrate-quiesce-pods"

const PodStageLabel string = "migration-stage-pod"

// Restic annotations
const ResticRestoreAnnotationPrefix string = "snapshot.velero.io"
const ResticBackupAnnotation string = "backup.velero.io/backup-volumes"
