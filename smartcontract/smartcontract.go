package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// staticOrganizationKey of the smartcontract address in the environment variable
const staticOrganizationKey = "STATIC_ORGANIZATION"

// remoteUrlKey of the blockchain node in the environment variable
const remoteUrlKey = "REMOTE_URL"

// accountKey of the private key in the environment variable
const accountKey = "ACCOUNT"

var privateKey *ecdsa.PrivateKey
var fromAddress common.Address

// Init the configuration for the required parameters to run the smartcontract
func Init(appConfig *configuration.Config) error {
	if !appConfig.Exist(staticOrganizationKey) {
		return fmt.Errorf("missing required parameter in configuration: %s", staticOrganizationKey)
	}
	if !appConfig.Exist(remoteUrlKey) {
		return fmt.Errorf("missing required parameter in configuration: %s", remoteUrlKey)
	}
	if !appConfig.Exist(accountKey) {
		return fmt.Errorf("missing required parameter in configuration: %s", accountKey)
	}

	var err error
	privateKeyHex := appConfig.GetString(accountKey)
	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to create ECDSA string: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	return nil
}

func Client(appConfig *configuration.Config) (*ethclient.Client, error) {
	remoteUrl := appConfig.GetString(remoteUrlKey)

	client, err := ethclient.Dial(remoteUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to dial the remote url: %w", err)
	}

	return client, nil
}

// TransactionOpts returns the transaction parameters derived from the private key.
// It has the account that will sign the transaction, also it will have the nonce and the gas price as well
//
// Important, to call this function after Init, since it won't verify the private key existence
func TransactionOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce for %s: %w", fromAddress.String(), err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get the chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to binding from the private key: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// Instance of the smartcontract. Make sure that you call this after Init()
// since this function won't check the configuration
func Instance(appConfig *configuration.Config) (*StaticOrganization, error) {
	staticOrganizationAddress := appConfig.GetString(staticOrganizationKey)
	address := common.HexToAddress(staticOrganizationAddress)

	client, err := Client(appConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %w", err)
	}

	instance, err := NewStaticOrganization(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create a smartcontract instance")
	}

	return instance, nil
}
