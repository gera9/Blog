package repositories

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindAll(ctx context.Context, coll *mongo.Collection, limit, offset int, results any) error {
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return err
	}

	err = cursor.All(ctx, results)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	return err
}

func FindById(ctx context.Context, coll *mongo.Collection, id string, result any) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": oid,
	}

	return coll.FindOne(ctx, filter).Decode(result)
}

func CreateReturningId[T any](ctx context.Context, coll *mongo.Collection, document any) (T, error) {
	var zero T

	r, err := coll.InsertOne(ctx, document)
	if err != nil {
		return zero, err
	}

	id, ok := r.InsertedID.(T)
	if !ok {
		return zero, fmt.Errorf("InsertedID is not of the expected type %T", zero)
	}

	return id, nil
}

func Create(ctx context.Context, coll *mongo.Collection, document any) (string, error) {
	id, err := CreateReturningId[bson.ObjectID](ctx, coll, document)
	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

func UpdateById(ctx context.Context, coll *mongo.Collection, id string, document any) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": oid,
	}

	update := bson.M{
		"$set": document,
	}

	ur, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if ur.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func DeleteById(ctx context.Context, coll *mongo.Collection, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": oid,
	}

	_, err = coll.DeleteOne(ctx, filter)

	return err
}
