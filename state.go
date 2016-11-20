package ecsstate

import (
	"github.com/jinzhu/gorm"
)

// StateOps is the interface for refreshing and interacting with the local
// ECS state.
type StateOps interface {
	FindClusterByName(td string) Cluster
	FindLocationsForTaskDefinition(td string) *[]ContainerInstance
	FindTaskDefinition(td string) TaskDefinition
	RefreshClusterState()
	RefreshContainerInstanceState()
	RefreshTaskState()
	DB() *gorm.DB
}
