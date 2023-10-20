package mongodb

type MongoDatabase struct{}

func NewMongoDBInstance(log, config interface{})(*MongoDatabase,error){return &MongoDatabase{}, nil}

func (rd *MongoDatabase) Create(config, data interface{}) error{return nil}
func(rd *MongoDatabase) Get(config, data interface{}) (interface{}, error){return nil,nil}
func(rd *MongoDatabase) Update(config, data interface{}) error{return nil}
func(rd *MongoDatabase) Delete(config, data interface{}) error{return nil}