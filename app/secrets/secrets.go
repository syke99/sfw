package secrets

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"
)

type secrets struct {
	vault *vault.KVv2
}

func NewSecretsStore(mountPath string, url string, secretId string, roleId string) (SecretsStore, error) {
	client, err := vault.NewClient(&vault.Config{
		Address: url,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client: %w", err)
	}

	role, err := approle.NewAppRoleAuth(
		roleId,
		&approle.SecretID{FromString: secretId},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create approle auth role: %w", err)
	}

	token, err := client.Auth().Login(context.Background(), role)
	if err != nil {
		return nil, fmt.Errorf("failed to login to vault instance: %w", err)
	}

	client.SetToken(token.Auth.ClientToken)

	kvV2Client := client.KVv2(mountPath)

	return &secrets{
		vault: kvV2Client,
	}, nil
}

func (s *secrets) Get(secret string) (string, error) {
	scrt, err := s.vault.Get(context.Background(), secret)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret from vault: %w", err)
	}

	return s.getSecretAsString(scrt, secret), nil
}

func (s *secrets) getSecretAsString(scrt *vault.KVSecret, secret string) string {
	return scrt.Data[secret].(string)
}
