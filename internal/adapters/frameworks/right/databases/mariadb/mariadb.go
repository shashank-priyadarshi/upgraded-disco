package mariadb

type MariaDatabase struct{}

func (rd *MariaDatabase) Create(config, data interface{}) error{return nil}
func(rd *MariaDatabase) Get(config, data interface{}) (interface{}, error){return nil,nil}
func(rd *MariaDatabase) Update(config, data interface{}) error{return nil}
func(rd *MariaDatabase) Delete(config, data interface{}) error{return nil}