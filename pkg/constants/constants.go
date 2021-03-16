package constants

import (
	"errors"
)

// The finalizer
const MySQLClusterFinalizer = "moco.cybozu.com/mysqlcluster"

// MySQL user names for MOCO
const (
	// AdminUser is a name of MOCO operator-admin user.
	// This user is a super user especially for creating and granting privileges to other users.
	AdminUser = "moco-admin"

	// AgentUser is a name of MOCO agent user.
	AgentUser = "moco-agent"

	// ReplicationUser is a name of MOCO replicator user.
	ReplicationUser = "moco-repl"

	// CloneDonorUser is a name of MOCO clone-donor user.
	CloneDonorUser = "moco-clone-donor"

	// ReadOnlyUser is a name of MOCO predefined human user with wide read-only rights used for manual operation.
	ReadOnlyUser = "moco-readonly"

	// WritableUser is a name of MOCO predefined human user with wide read/write rights used for manual operation.
	WritableUser = "moco-writable"
)

const (
	// MySQLDataPath is a path for MySQL data dir.
	MySQLDataPath = "/var/lib/mysql"

	// MySQLConfPath is a path for MySQL conf dir.
	MySQLConfPath = "/etc/mysql"

	// MySQLInitConfPath is a path for MySQL conf dir generated by moco-init.
	MySQLInitConfPath = "/etc/mysql-conf.d"

	// MySQLConfName is a filename for MySQL conf.
	MySQLConfName = "my.cnf"

	// RunPath is a path for variable files which concerns MySQLd.
	RunPath = "/run"

	// LogDirPath is a path for /var/log/mysql.
	LogDirPath = "/var/log/mysql"

	// MySQLFilesPath is a path for /var/lib/mysql-files.
	MySQLFilesPath = "/var/lib/mysql-files"

	// MySQLErrorLogName is a filename of error log for MySQL.
	MySQLErrorLogName = "mysql.err"

	// MySQLSlowLogName is a filename of slow query log for MySQL.
	MySQLSlowLogName = "mysql.slow"

	// TmpPath is a path for /tmp.
	TmpPath = "/tmp"

	// MyCnfSecretPath is the path for my.cnf formated credentials for CLI
	MyCnfSecretPath = "/mysql-credentials"

	// MySQLConfTemplatePath is
	MySQLConfTemplatePath = "/etc/mysql_template"

	// DonorPasswordPath is the path to donor user passsword file
	DonorPasswordPath = MySQLDataPath + "/donor-password"

	// AgentPasswordPath is the path to misc user passsword file
	AgentPasswordPath = MySQLDataPath + "/agent-password"

	// InnoDBBufferPoolRatioPercent is the ratio of InnoDB buffer pool size to resource.limits.memory or resource.requests.memory
	// Note that the pool size doesn't set to lower than 128MiB, which is the default innodb_buffer_pool_size value
	InnoDBBufferPoolRatioPercent = 70
)

const (
	// FluentBitConfigName is a filename for fluent-bit conf.
	FluentBitConfigName = "fluent-bit.conf"

	// FluentBitConfigPath is the path for fluent-bit config directory.
	FluentBitConfigPath = "/fluent-bit/etc"
)

const (
	// AgentTokenEnvName is a name of the environment variable of agent token.
	AgentTokenEnvName = "MOCO_AGENT_TOKEN"

	// AgentTokenParam is a name of the param of agent token.
	AgentTokenParam = "token"
)

// environment variables for agent
const (
	// PodNameEnvKey is the environment variable name of a pod name.
	PodNameEnvKey = "POD_NAME"

	// ClusterNameEnvKey is the environment variable name of MySQLCluster name.
	ClusterNameEnvKey = "CLUSTER_NAME"
)

const (
	// ReadOnlyMyCnfKey is the username and password of moco-readonly formated as my.cnf
	ReadOnlyMyCnfKey = ReadOnlyUser + "-my.cnf"

	// WritableMyCnfKey is the username and password or moco-writable formated as my.cnf
	WritableMyCnfKey = WritableUser + "-my.cnf"

	// ReplicationSourcePrimaryHostKey etc. are Secret key for replication source secret
	ReplicationSourcePrimaryHostKey            = "PRIMARY_HOST"
	ReplicationSourcePrimaryUserKey            = "PRIMARY_USER"
	ReplicationSourcePrimaryPasswordKey        = "PRIMARY_PASSWORD"
	ReplicationSourcePrimaryPortKey            = "PRIMARY_PORT"
	ReplicationSourceCloneUserKey              = "CLONE_USER"
	ReplicationSourceClonePasswordKey          = "CLONE_PASSWORD"
	ReplicationSourceInitAfterCloneUserKey     = "INIT_AFTER_CLONE_USER"
	ReplicationSourceInitAfterClonePasswordKey = "INIT_AFTER_CLONE_PASSWORD"
)

const (
	// InitializedClusterIndexField is an index name for Initialized MySQL Clusters
	InitializedClusterIndexField = ".status.conditions.type.initialized"
)

const (
	// PodNamespaceEnvName is a name of the environment variable of a pod namespace.
	PodNamespaceEnvName = "POD_NAMESPACE"
)

type OperationPhase string

const (
	PhaseInitializing    = OperationPhase("initializing")
	PhaseWaitRelayLog    = OperationPhase("wait-relay-log")
	PhaseRestoreInstance = OperationPhase("restoring-instance")
	PhaseCompleted       = OperationPhase("completed")
)

var AllOperationPhases = []OperationPhase{
	PhaseInitializing,
	PhaseWaitRelayLog,
	PhaseRestoreInstance,
	PhaseCompleted,
}

var (
	// ErrConstraintsViolation is returned when the constraints violation occurs
	ErrConstraintsViolation = errors.New("constraints violation occurs")
	// ErrConstraintsRecovered is returned when the constrains recovered but once violated
	ErrConstraintsRecovered = errors.New("constrains recovered but once violated")
	// ErrCannotCompareGTIDs is returned if GTID comparison returns error
	ErrCannotCompareGTIDs = errors.New("cannot compare gtids")
)
