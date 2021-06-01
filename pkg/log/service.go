package log

import (
	"code-boiler/pkg/elasticsearch"
	"context"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

func InsertErrorLog(ctx context.Context, log *LogError) error {
	return elasticsearch.Insert(ctx, INDEX_LOG_ERROR, log)
}

func InsertActivityLog(ctx context.Context, log *LogError) error {
	return elasticsearch.Insert(ctx, INDEX_LOG_ACTIVITY, log)
}

func InsertLoginLog(ctx context.Context, log *LogError) error {
	return elasticsearch.Insert(ctx, INDEX_LOG_LOGIN, log)
}

func Update(ctx context.Context, index, ID string, update map[string]interface{}) error {
	if _, err := elasticsearch.Client.Update().
		Index(index).
		Type("_doc").
		Id(ID).Doc(update).Do(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"Elasticsearch": "cannot insert data",
			"ID":            ID,
			"Index":         index,
			"Data":          update,
		}).Error(err.Error())

		return err
	}
	return nil
}

func Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error) {
	results, err := elasticsearch.Client.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Elasticsearch": "cannot search data",
		}).Error(err.Error())
		return nil, err
	}

	return results, nil
}

func Count(ctx context.Context, index string, searchSource *el.SearchSource) (int64, error) {
	count, err := elasticsearch.Client.Count(index).Query(searchSource).Do(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}
