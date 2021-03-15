// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"context"
	"time"

	"github.com/globocom/gitlab-lint/rules"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCollection struct {
	dbName  string
	session *mongo.Client
}

type DB interface {
	Aggregate(d rules.Queryable, pipeline interface{}) ([]bson.M, error)
	BuildSearchQueryFromString(q string, d rules.Queryable) bson.M
	DeleteMany(d rules.Queryable, q bson.M) (*mongo.DeleteResult, error)
	Get(d rules.Queryable, q bson.M, o *options.FindOneOptions) error
	GetAll(d rules.Queryable, q bson.M, o *options.FindOptions) ([]rules.Queryable, error)
	Insert(d rules.Queryable) (*mongo.InsertOneResult, error)
	InsertMany(d rules.Queryable, i []interface{}) (*mongo.InsertManyResult, error)
}

func newDBContext() (context.Context, context.CancelFunc) {
	timeout := viper.GetDuration("db.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func NewMongoSession() (DB, error) {
	log.Debug("[DB] New mongo session")
	dbURI := viper.GetString("mongodb.endpoint")
	ctx, _ := newDBContext()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("[DB] Error on create mongo session: %s", err)
	}
	mongo := &mongoCollection{
		session: client,
		dbName:  viper.GetString("mongodb.name"),
	}
	return mongo, err
}

func (m *mongoCollection) Aggregate(d rules.Queryable, pipeline interface{}) ([]bson.M, error) {
	log.Debug("[DB] Aggregate...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if cursor == nil {
		return nil, err
	}
	var results []bson.M
	ctx, _ = newDBContext()
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *mongoCollection) DeleteMany(d rules.Queryable, q bson.M) (*mongo.DeleteResult, error) {
	log.Debug("[DB] DeleteMany...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	return collection.DeleteMany(ctx, q)
}

func (m *mongoCollection) Insert(d rules.Queryable) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] Insert...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	return collection.InsertOne(ctx, d)
}

func (m *mongoCollection) InsertMany(d rules.Queryable, i []interface{}) (*mongo.InsertManyResult, error) {
	log.Debug("[DB] InsertMany...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	return collection.InsertMany(ctx, i)
}

func (m *mongoCollection) Get(d rules.Queryable, q bson.M, o *options.FindOneOptions) error {
	log.Debug("[DB] Get...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	return collection.FindOne(ctx, q).Decode(d)
}

func (m mongoCollection) GetAll(d rules.Queryable, q bson.M, opts *options.FindOptions) ([]rules.Queryable, error) {
	log.Debug("[DB] GetAll...")
	collection := m.session.Database(m.dbName).Collection(d.GetCollectionName())
	ctx, _ := newDBContext()
	cur, err := collection.Find(ctx, q, opts)
	if err != nil {
		log.Errorf("[DB] Find: %s", err)
		return nil, err
	}

	defer cur.Close(ctx)

	items := []rules.Queryable{}

	for cur.Next(ctx) {
		data := d.Cast()
		if err := cur.Decode(data); err != nil {
			log.Errorf("[DB] Decode: %s", err)
			return nil, err
		}
		items = append(items, data)
	}

	if err := cur.Err(); err != nil {
		log.Errorf("[DB] Cursor: %s", err)
		return nil, err
	}

	return items, nil
}

func (m mongoCollection) BuildSearchQueryFromString(q string, d rules.Queryable) bson.M {
	fields := d.GetSearchableFields()
	if q == "" || fields == nil {
		return nil
	}

	searchRegex := bson.M{"$regex": primitive.Regex{Pattern: q, Options: "i"}}
	queryFields := []bson.M{}
	for _, field := range fields {
		queryFields = append(queryFields, bson.M{field: searchRegex})
	}

	return bson.M{"$or": queryFields}
}
