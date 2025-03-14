package core

type DataRepository interface {
	SaveData(data *Data) error
}
