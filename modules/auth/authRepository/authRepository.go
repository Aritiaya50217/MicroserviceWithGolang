package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	"github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth"
	playerPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerPb"
	grpccon "github.com/Aritiaya50217/MicroserviceWithGolang/pkg/database/grpcCon"
	"github.com/Aritiaya50217/MicroserviceWithGolang/pkg/jwtauth"
	"github.com/Aritiaya50217/MicroserviceWithGolang/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error)
		InsertOnePlayerCredential(pctx context.Context, req *auth.Credentail) (primitive.ObjectID, error)
		FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credentail, error)
		UpdatedOnePlayerCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error
		DeleteOnePlayerCredential(pctx context.Context, credentialId string) (int64, error)
		FindOneAccessToken(pctx context.Context, accessToken string) (*auth.Credentail, error)
		RoleCount(pctx context.Context) (int64, error)
		AccessToken(cfg *config.Config, claims *jwtauth.Claims) string
		RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string
	}
	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

func (r *authRepository) authDbConn(ptcx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}

func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Player().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return nil, errors.New("error: email or password is incorrect")
	}

	return result, nil
}

func (r *authRepository) FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error : gRPC connection failed : %s ", err.Error())
	}
	result, err := conn.Player().FindOnePlayerProfileToRefresh(ctx, req)
	if err != nil {
		log.Fatal("Error : FindOnePlayerProfileToRefresh : %s", err.Error())
		return nil, errors.New("error : player profile not found")
	}
	return result, nil
}

func (r *authRepository) InsertOnePlayerCredential(pctx context.Context, req *auth.Credentail) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	req.CreatedAt = utils.LocalTime()
	req.UpdatedAt = utils.LocalTime()

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error : InsertOnePlayerCredential failed")
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *authRepository) FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credentail, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.authDbConn(ctx)
	col := db.Collection("auth")
	result := new(auth.Credentail)
	if err := col.FindOne(ctx, bson.M{
		"_id": utils.ConvertToObjectId(credentialId),
	}).Decode(result); err != nil {
		log.Printf("Error : FindOnePlayerCredential failed : %s ", err.Error())
		return nil, errors.New("error : find one player credential failed")
	}
	return result, nil
}

func (r *authRepository) UpdatedOnePlayerCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")
	_, err := col.UpdateOne(
		ctx,
		bson.M{
			"_id": utils.ConvertToObjectId(credentialId),
		},
		bson.M{
			"$set": bson.M{
				"player_id":     req.PlayerId,
				"access_token":  req.AccessToken,
				"refresh_token": req.RefreshToken,
				"updated_at":    req.UpdatedAt,
			},
		},
	)
	if err != nil {
		log.Printf("Error : UpdateOnePlayerCredential failed : %s ", err.Error())
		return errors.New("error: player credential not found")
	}
	return nil
}

func (r *authRepository) DeleteOnePlayerCredential(pctx context.Context, credentialId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.DeleteOne(ctx, bson.M{
		"_id": utils.ConvertToObjectId(credentialId),
	})
	if err != nil {
		log.Printf("Error : DeleteOnePlayerCredential failed : %s ", err.Error())
		return -1, errors.New("error : delete player credential faild")
	}
	log.Printf("DeleteOnePlayerCredentail result : %v ", result)
	return result.DeletedCount, nil
}

func (r *authRepository) FindOneAccessToken(pctx context.Context, accessToken string) (*auth.Credentail, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	credential := new(auth.Credentail)
	if err := col.FindOne(ctx, bson.M{
		"access_token": accessToken,
	}).Decode(credential); err != nil {
		log.Printf("Error : FindOneAccessToken failed : %s", err.Error())
		return nil, errors.New("error : access token not found")
	}
	return credential, nil
}

func (r *authRepository) RoleCount(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("roles")
	count, err := col.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error : RoleCount faileed : %s", err.Error())
		return -1, errors.New("Error : roles count failed")
	}
	return count, nil
}

func (r *authRepository) AccessToken(cfg *config.Config, claims *jwtauth.Claims) string {
	return jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		PlayerId: claims.PlayerId,
		RoleCode: int(claims.RoleCode),
	}).SignToken()
}

func (r *authRepository) RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string {
	return jwtauth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtauth.Claims{
		PlayerId: claims.PlayerId,
		RoleCode: int(claims.RoleCode),
	}).SignToken()
}
