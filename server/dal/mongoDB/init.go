package mongoDB

import (
	"context"
	ue "userbase-api/server/utils/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const certPath = "cert/mongoCert.pem"

func Setup(ctx context.Context) error {
	var err error
	uri := "mongodb+srv://userapi.djuypob.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=" + certPath
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return ue.NewMongoDbError("Setup failed to connect to client, err=%v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return ue.NewMongoDbError("Setup failed to ping client, err=%v", err)
	}
	return nil
}

func Disconnect(ctx context.Context) error {
	return client.Disconnect(ctx)
}
