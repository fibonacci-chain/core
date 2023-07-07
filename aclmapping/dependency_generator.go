package aclmapping

import (
	aclkeeper "github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
	aclbankmapping "github.com/fibonacci-chain/core/aclmapping/bank"
	acldexmapping "github.com/fibonacci-chain/core/aclmapping/dex"
	acloraclemapping "github.com/fibonacci-chain/core/aclmapping/oracle"
	acltokenfactorymapping "github.com/fibonacci-chain/core/aclmapping/tokenfactory"
	aclwasmmapping "github.com/fibonacci-chain/core/aclmapping/wasm"
)

type CustomDependencyGenerator struct{}

func NewCustomDependencyGenerator() CustomDependencyGenerator {
	return CustomDependencyGenerator{}
}

func (customDepGen CustomDependencyGenerator) GetCustomDependencyGenerators() aclkeeper.DependencyGeneratorMap {
	dependencyGeneratorMap := make(aclkeeper.DependencyGeneratorMap)
	wasmDependencyGenerators := aclwasmmapping.NewWasmDependencyGenerator()

	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acldexmapping.GetDexDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(aclbankmapping.GetBankDepedencyGenerator())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acltokenfactorymapping.GetTokenFactoryDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(wasmDependencyGenerators.GetWasmDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acloraclemapping.GetOracleDependencyGenerator())

	return dependencyGeneratorMap
}
