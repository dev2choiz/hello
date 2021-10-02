package config

import (
	kms "cloud.google.com/go/kms/apiv1"
	"context"
	"encoding/base64"
	"github.com/spf13/viper"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

type Config struct {
	// set by command
	Name           string
	Port           string
	WithImprobable bool
	WithTLS        bool

	// set by viper
	DecryptConf      bool
	AppEnvContext    string
	KmsCryptoKeys    string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresPort     string
}

var Conf = &Config{}

var confInitialized = false

var sensibles = []string {
	"postgres_password",
}

func init()  {
	if confInitialized {
		return
	}

	viper.AutomaticEnv()
	Conf.DecryptConf = viper.GetBool("decrypt_conf")
	Conf.AppEnvContext = viper.GetString("app_env_context")
	Conf.KmsCryptoKeys = viper.GetString("kms_crypto_keys")
	Conf.PostgresHost = viper.GetString("postgres_host")
	Conf.PostgresUser = viper.GetString("postgres_user")
	Conf.PostgresPassword = getString("postgres_password")
	Conf.PostgresDB = viper.GetString("postgres_db")
	Conf.PostgresPort = viper.GetString("postgres_port")

	confInitialized = true
}

// getString return the env var value, decrypting it if needed
func getString(s string) string {
	v := viper.GetString(s)
	if v == "" || !shouldDecrypt(s) {
		return v
	}

	cypherText, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		panic(err)
	}
	v, err = kmsDecrypt(cypherText)
	if err != nil {
		panic(err)
	}
	return v
}

// kmsDecrypt decrypt a value with kms
func kmsDecrypt(cypherText []byte) (string, error) {
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		panic(err)
	}

	req := &kmspb.DecryptRequest{
		Name:       Conf.KmsCryptoKeys,
		Ciphertext: cypherText,
	}
	res, err := client.Decrypt(ctx, req)
	return string(res.Plaintext), err
}

func shouldDecrypt(e string) bool {
	if !Conf.DecryptConf {
		return false
	}

	for _, a := range sensibles {
		if a == e {
			return true
		}
	}
	return false
}
